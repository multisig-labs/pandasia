package api

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/exp/slog"

	"github.com/AbsaOSS/env-binder/env"
	"github.com/ava-labs/avalanchego/utils/cb58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/multisig-labs/pandasia/pkg/contracts"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/merkle"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"github.com/multisig-labs/pandasia/pkg/syncer"
	"github.com/neilotoole/errgroup"
)

type Config struct {
	// Seconds between job runs
	JobPeriod string `env:"JOB_PERIOD,default=24h"`
	// Serve content from public not embedded
	ServeEmbedded bool `env:"SERVE_EMBEDDED,default=true"`
	jobPeriod     time.Duration
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

type pagination struct {
	Limit  uint64
	Offset uint64
}

type treeResponse struct {
	TreeType string
	Height   int
	Root     string
	Addrs    []string
}

type proofResponse struct {
	TreeType string
	Height   int
	Root     string
	Proof    []string
	SigV     string
	SigR     string
	SigS     string
}

type sigResponse struct {
	SigV string
	SigR string
	SigS string
}

type findPchainResponse struct {
	Exists bool `json:"exists"`
}

type addAddrParams struct {
	Addrs  []string `json:"addrs"`
	Height int      `json:"height"`
}

func StartHttpServer(dbFileName string, host string, port int, nodeURL string, webContent fs.FS, pandasiaAddr string) {
	slog.Info("init", "EnvConfig", EnvConfig)

	ctx, cancel := context.WithCancel(context.Background())
	dbFile, queries := db.OpenDB(dbFileName)
	g, gCtx := errgroup.WithContext(ctx)

	e := echo.New()
	e.HideBanner = true
	e.Debug = true // Show more detailed errors in json response

	authTokenFromEnv, ok := os.LookupEnv("PANDASIA_AUTHTOKEN")
	if ok {
		log.Info("Using PANDASIA_AUTHTOKEN from env")
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
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

	e.GET("/health", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})

	e.GET("/check_pchain_addr/:root", func(c echo.Context) error {
		root := c.Param("root")
		addr := c.QueryParam("addr")

		// TODO check addr format

		t, err := queries.FindMerkleTreeByRoot(ctx, root)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		p := db.FindPchainAddrParams{
			RewardsAddr: sql.NullString{addr, true},
			TypeID:      12,
			Height:      int64(t.Height),
		}

		cnt, err := queries.FindPchainAddr(ctx, p)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		result := findPchainResponse{Exists: cnt > 0}
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/trees", func(c echo.Context) error {
		roots, err := queries.ListMerkleRoots(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		return c.JSON(http.StatusOK, roots)
	})

	e.GET("/trees/:root", func(c echo.Context) error {
		root := c.Param("root")
		t, err := queries.FindMerkleTreeByRoot(ctx, root)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		addrs, err := merkle.LoadAddrsFromTree(t.Tree)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		r := treeResponse{
			TreeType: t.TreeType,
			Height:   int(t.Height),
			Root:     t.Root,
			Addrs:    addrs,
		}

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/sync", func(c echo.Context) error {
		authToken := c.QueryParam("token")
		if authToken != "" && subtle.ConstantTimeCompare([]byte(authTokenFromEnv), []byte(authToken)) == 1 {
			// spawn a new goroutine to do the work
			// TODO improve this to not allow multiple calls
			go func() {
				err := updatePChain(gCtx, dbFile, queries, nodeURL)
				if err != nil {
					slog.Error("Error updating pchain", err)
				}
			}()

			return c.JSON(http.StatusOK, "success")
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("invalid or missing token"))
		}
	})

	e.GET("/sync_all", func(c echo.Context) error {
		authToken := c.QueryParam("token")
		startHeightString := c.QueryParam("startHeight")

		startHeight, err := strconv.ParseInt(startHeightString, 10, 64)
		if err != nil {
			slog.Error("Unable to convert startHeight to int %w", err)
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("unable to parse startHeight %w", err))
		}

		if authToken != "" && subtle.ConstantTimeCompare([]byte(authTokenFromEnv), []byte(authToken)) == 1 {
			go func() {
				err := syncer.SyncPChain(gCtx, queries, nodeURL, int64(startHeight), nil)
				if err != nil {
					slog.Error("Error sycning pchain", err)
				}
			}()

			return c.JSON(http.StatusOK, "success")
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("invalid or missing token"))
		}
	})

	// TODO: Remove before production
	// e.POST("/debug/add-addresses", func(c echo.Context) error {
	// 	var addrs addAddrParams
	// 	err := c.Bind(&addrs)
	// 	if err != nil {
	// 		slog.Error("Error parsing arguments", err)
	// 		return c.String(http.StatusBadRequest, err.Error())
	// 	}

	// 	vaddrs := []merkle.ValidatorAddress{}
	// 	for _, addr := range addrs.Addrs {
	// 		_, _, addrBytes, err := address.Parse(addr)
	// 		if err != nil {
	// 			slog.Error("Error parsing address", err)
	// 			return c.JSON(http.StatusInternalServerError, err.Error())
	// 		}
	// 		vaddr := merkle.ValidatorAddress{Addr: addr, AddrHex: common.BytesToAddress(addrBytes).Hex()}
	// 		vaddrs = append(vaddrs, vaddr)
	// 	}

	// 	tree, err := merkle.GenerateTree(vaddrs)
	// 	if err != nil {
	// 		slog.Error("Error generating tree", err.Error())
	// 		return c.JSON(http.StatusInternalServerError, err.Error())
	// 	}

	// 	err = merkle.SaveTreeToDB(ctx, queries, merkle.TREE_TYPE_VALIDATOR, addrs.Height, tree, "debug")
	// 	if err != nil {
	// 		slog.Error("Error saving tree to databse", err.Error())
	// 		return c.JSON(http.StatusInternalServerError, err.Error())
	// 	}

	// 	return c.JSON(http.StatusOK, "success")
	// })

	e.GET("/signature/:sig", func(c echo.Context) error {
		sig := c.Param("sig")

		sigBytes, err := cb58.Decode(sig)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		r := sigResponse{
			SigR: fmt.Sprintf("0x%x", sigBytes[0:32]),
			SigS: fmt.Sprintf("0x%x", sigBytes[32:64]),
			SigV: fmt.Sprintf("0x%x", sigBytes[64:]),
		}

		return c.JSON(http.StatusOK, r)
	})

	// /proof/:root?addr=P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww&sig=24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx
	e.GET("/proof/:root", func(c echo.Context) error {
		root := c.Param("root")
		addr := c.QueryParam("addr")
		sig := c.QueryParam("sig")

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

		if sig != "" {
			sigBytes, err := cb58.Decode(sig)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
			}
			r.SigR = fmt.Sprintf("0x%x", sigBytes[0:32])
			r.SigS = fmt.Sprintf("0x%x", sigBytes[32:64])
			r.SigV = fmt.Sprintf("0x%x", sigBytes[64:])
		}

		return c.JSON(http.StatusOK, r)
	})

	// If we have a deployed addr, set up some routes to view contract data
	if pandasiaAddr != "" {
		slog.Info("Creating routes for /airdrops", "addr", pandasiaAddr)
		ctrcs, err := contracts.NewContracts(nodeURL, pandasiaAddr)
		if err != nil {
			panic(err)
		}

		e.GET("/airdrops", func(c echo.Context) error {
			p := new(pagination)
			if err := c.Bind(p); err != nil {
				return err
			}

			airdrops, err := ctrcs.Pandasia.GetAirdrops(nil, p.Offset, p.Limit)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
			}

			return c.JSON(http.StatusOK, airdrops)
		})
	}

	// Serve static files in public/
	// Basically "cd" into the /public folder of the embedded content
	contentSub, err := fs.Sub(webContent, "public")
	if err != nil {
		panic(err)
	}
	if !EnvConfig.ServeEmbedded {
		slog.Warn("Ignoring embedded content, serving from /public")
		contentSub = os.DirFS("./public")
	}
	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(contentSub))))

	// Setup Go Routines

	go func() {
		cSigTerm := make(chan os.Signal, 1)
		signal.Notify(cSigTerm, os.Interrupt, syscall.SIGTERM)
		<-cSigTerm
		slog.Info("Sigterm recvd, shutting down...")
		cancel()
	}()

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
				slog.Info("Update PChain routine context done")
				return nil
			case <-time.After(EnvConfig.jobPeriod):
				err := updatePChain(gCtx, dbFile, queries, nodeURL)
				if err != nil {
					slog.Error("Error in update PChain timer", "err::", err)
					return err
				}
			}
		}
	})

	if err := g.Wait(); err != nil {
		slog.Error("server", "msg", err)
	}

}

func updatePChain(gCtx context.Context, dbFile *sql.DB, queries *db.Queries, nodeURL string) error {
	slog.Info("starting sync job")

	counter := 0
	f := func(tot int, n int) {
		if counter%10 == 0 {
			slog.Info("sync_pchain", "tot::", tot, "n::", counter)
		}
		counter++
	}

	err := syncer.SyncPChainRecent(gCtx, queries, nodeURL, f)
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

	err = merkle.SaveTreeToDB(gCtx, queries, merkle.TREE_TYPE_VALIDATOR, int(height), tree, "")
	if err != nil {
		return err
	}
	slog.Info("tree saved to db", "height", height)
	return nil
}
