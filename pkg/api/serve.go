package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/merkle"
)

type proofResponse struct {
	TreeType string
	Height   int
	Root     string
	Proof    []string
}

func StartHttpServer(dbfile string, host string, port int) {
	e := echo.New()
	e.HideBanner = true
	e.Debug = true // Show more detailed errors in json response
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: 4, // ERROR level
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "ok")
	})

	e.GET("/trees", func(c echo.Context) error {
		ctx := context.Background()
		_, queries := db.OpenDB(dbfile)
		roots, err := queries.ListMerkleRoots(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		return c.JSON(http.StatusOK, roots)
	})

	e.GET("/trees/:root/:addr", func(c echo.Context) error {
		root := c.Param("root")
		addr := c.Param("addr")

		ctx := context.Background()
		_, queries := db.OpenDB(dbfile)

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

	// Start up HTTP server
	listenAddr := fmt.Sprintf("%s:%v", host, port)
	if err := e.Start(listenAddr); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
