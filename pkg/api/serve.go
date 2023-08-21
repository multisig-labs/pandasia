package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slog"

	"github.com/AbsaOSS/env-binder/env"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/merkle"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"github.com/multisig-labs/pandasia/pkg/syncer"
	"github.com/neilotoole/errgroup"
)

type Config struct {
	// Seconds between job runs
	JobPeriod string `env:"JOB_PERIOD,default=24h"`
	jobPeriod time.Duration
}

var EnvConfig *Config

func init() {
	var err error
	EnvConfig = &Config{}
	if err = env.Bind(EnvConfig); err != nil {
		panic(err)
	}
	EnvConfig.jobPeriod, err = time.ParseDuration(EnvConfig.JobPeriod)
	if err != nil {
		panic(err)
	}
}

type proofResponse struct {
	TreeType string
	Height   int
	Root     string
	Proof    []string
}

func StartHttpServer(dbFileName string, host string, port int, nodeURL string) {
	slog.Info("init", "EnvConfig", EnvConfig)

	ctx, cancel := context.WithCancel(context.Background())
	dbFile, queries := db.OpenDB(dbFileName)

	e := echo.New()
	e.HideBanner = true
	e.Debug = true // Show more detailed errors in json response
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: 4, // ERROR level
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogRemoteIP: true,
		LogError:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"URI", v.URI,
				"status", v.Status,
				"remote_ip", v.RemoteIP,
				"error", v.Error,
			)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})

	e.GET("/trees", func(c echo.Context) error {
		roots, err := queries.ListMerkleRoots(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		return c.JSON(http.StatusOK, roots)
	})

	e.GET("/trees/:root/:addr", func(c echo.Context) error {
		root := c.Param("root")
		addr := c.Param("addr")

		t, err := queries.FindMerkleTreeByRoot(ctx, root)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		proof, err := merkle.GenerateProof(t.Tree, addr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		proofAry := make([]string, len(proof))
		for a, b := range proof {
			proofAry[a] = hexutil.Encode(b)
		}

		r := proofResponse{
			TreeType: t.TreeType,
			Height:   int(t.Height),
			Root:     t.Root,
			Proof:    proofAry,
		}

		return c.JSON(http.StatusOK, r)
	})

	go func() {
		cSigTerm := make(chan os.Signal, 1)
		signal.Notify(cSigTerm, os.Interrupt, syscall.SIGTERM)
		<-cSigTerm
		slog.Info("Sigterm recvd, shutting down...")
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		listenAddr := fmt.Sprintf("%s:%v", host, port)
		return e.Start(listenAddr)
	})

	g.Go(func() error {
		<-gCtx.Done()
		return e.Shutdown(context.Background())
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(EnvConfig.jobPeriod):
				slog.Info("starting sync job")
				err := syncer.SyncPChain(gCtx, queries, nodeURL, nil)
				if err != nil {
					return err
				}

				slog.Info("updating rewards")
				err = syncer.UpdateRewards(gCtx, dbFile)
				if err != nil {
					return err
				}

				height, err := queries.MaxHeight(gCtx)
				if err != nil {
					return err
				}

				slog.Info("loading addrs", "height", height)
				vaddrs, err := merkle.LoadAddrsFromDB(gCtx, queries, pchain.AddValidatorTxId, int(height))
				if err != nil {
					return err
				}

				slog.Info("calculating tree", "len(addrs)", len(vaddrs))
				tree, err := merkle.GenerateTree(vaddrs)
				if err != nil {
					return err
				}

				err = merkle.SaveTreeToDB(gCtx, queries, merkle.TREE_TYPE_VALIDATOR, int(height), tree)
				if err != nil {
					return err
				}
				slog.Info("tree saved to db", "height", height)
			}
		}
	})

	if err := g.Wait(); err != nil {
		slog.Error("server", "msg", err)
	}

}