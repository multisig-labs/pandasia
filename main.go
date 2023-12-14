package main

import (
	"bufio"
	"context"
	"embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"runtime/debug"

	"github.com/jxskiss/mcli"
	"github.com/multisig-labs/pandasia/pkg/api"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/merkle"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"github.com/multisig-labs/pandasia/pkg/syncer"
	"github.com/multisig-labs/pandasia/pkg/version"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/slog"
	_ "modernc.org/sqlite"
)

//go:embed public/*
var webContent embed.FS

func main() {
	defer handlePanic()
	mcli.Add("sync-pchain", syncPchainCmd, "Sync the P-Chain to the db")
	mcli.Add("generate-tree", generateTreeCmd, "Generate the Merkle Tree at current height and save to DB")
	mcli.Add("generate-tree-stdin", generateTreeFromStdinCmd, "Generate the Merkle Tree from list of addr on stdin")
	mcli.Add("serve", serveApiCmd, "Start API server")
	mcli.Add("verify-tree", verifyTreeCmd, "Verify an entire merkle tree")
	mcli.Add("addrs", addrsForHeight, "Output list of addrs that would be included at a specific height")
	mcli.AddHelp()
	mcli.AddCompletion()
	mcli.Run()
}

func syncPchainCmd() {
	args := struct {
		NodeURL     string `cli:"--node-url, Avalanche node URL" default:"http://localhost:9650"`
		DbFile      string `cli:"--db, SQLite database file name" default:"pandasia.db"`
		StartHeight string `cli:"--start-height, Start height to begin fetch" default:"last"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	_, queries := db.OpenDB(args.DbFile)

	var bar *progressbar.ProgressBar
	f := func(tot int, n int) {
		if bar == nil {
			bar = progressbar.NewOptions(tot,
				progressbar.OptionSetWriter(os.Stderr),
				progressbar.OptionSetRenderBlankState(true),
				progressbar.OptionEnableColorCodes(true),
				progressbar.OptionShowCount(),
				progressbar.OptionShowIts(),
				progressbar.OptionThrottle(1000*time.Millisecond),
				progressbar.OptionSetDescription("[cyan]Syncing P-chain...[reset]"),
				progressbar.OptionOnCompletion(func() {
					fmt.Fprint(os.Stderr, "\n")
				}),
			)
		}
		bar.Add(n)
	}

	if args.StartHeight == "last" {
		slog.Info("Syncing recent transactions")
		err := syncer.SyncPChainRecent(ctx, queries, args.NodeURL, f)
		handleError(err)
	} else {
		startHeight, err := strconv.ParseInt(args.StartHeight, 10, 64)
		handleError(err)

		slog.Info("Syncing transactions from", "blockHeight", startHeight)
		err = syncer.SyncPChain(ctx, queries, args.NodeURL, startHeight, f)
		handleError(err)
	}
}

func serveApiCmd() {
	args := struct {
		DbFile       string `cli:"--db, SQLite database file name" default:"pandasia.db"`
		NodeURL      string `cli:"--node-url, Avalanche node URL" default:"http://localhost:9650"`
		Host         string `cli:"--host, host" default:"0.0.0.0"`
		Port         int    `cli:"--port, port" default:"8000"`
		PandasiaAddr string `cli:"--pandasia-addr, contract addr"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	api.StartHttpServer(args.DbFile, args.Host, args.Port, args.NodeURL, webContent, args.PandasiaAddr)
}

// Generate a Merkle tree from txs table at current height, and store tree JSON in DB
func generateTreeCmd() {
	args := struct {
		DbFile string `cli:"--db, SQLite database file name" default:"pandasia.db"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	dbFile, queries := db.OpenDB(args.DbFile)

	slog.Info("updating rewards")
	err := syncer.UpdateRewards(ctx, dbFile)
	handleError(err)

	height, err := queries.MaxHeight(ctx)
	handleError(err)

	slog.Info("loading addrs", "height", height)
	vaddrs, err := merkle.LoadAddrsFromDB(ctx, queries, pchain.AddValidatorTxId, int(height))
	handleError(err)

	slog.Info("calculating tree", "len(addrs)", len(vaddrs))
	tree, err := merkle.GenerateTree(vaddrs)
	handleError(err)

	err = merkle.SaveTreeToDB(ctx, queries, merkle.TREE_TYPE_VALIDATOR, int(height), tree, "")
	handleError(err)
	slog.Info("saved tree to db", "height", height)
}

// Generate a merkle tree from addresses piped in via stdin
// TODO this is just for testing, need a route to do this for real via API
func generateTreeFromStdinCmd() {
	args := struct {
		DbFile      string `cli:"--db, SQLite database file name" default:"pandasia.db"`
		Description string `cli:"--desc, description of the merkle tree"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	_, queries := db.OpenDB(args.DbFile)

	scanner := bufio.NewScanner(os.Stdin)
	vaddrs := []merkle.ValidatorAddress{}

	for scanner.Scan() {
		vaddrs = append(vaddrs, merkle.ValidatorAddress{AddrHex: scanner.Text()})
	}
	handleError(scanner.Err())

	slog.Info("calculating tree", "len(addrs)", len(vaddrs))
	tree, err := merkle.GenerateTree(vaddrs)
	handleError(err)

	err = merkle.SaveTreeToDB(ctx, queries, merkle.TREE_TYPE_CUSTOM, 0, tree, args.Description)
	handleError(err)
	slog.Info("saved tree to db", "desc", args.Description)

}

func verifyTreeCmd() {
	args := struct {
		Root   string `cli:"--root, merkle root to verify"`
		DbFile string `cli:"--db, SQLite database file name" default:"pandasia.db"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	_, queries := db.OpenDB(args.DbFile)
	ok, err := merkle.VerifyTree(ctx, queries, args.Root)
	handleError(err)
	slog.Info("Verification", "ok", ok)
}

// For debugging, grab all addresses from DB that would be included in a Merkle tree at height
func addrsForHeight() {
	args := struct {
		DbFile string `cli:"--db, SQLite database file name" default:"pandasia.db"`
		Height int    `cli:"--height, Height at which to grab addrs"`
		TxType int    `cli:"--tx-type, validator=12 delegator=14" default:"12"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	_, queries := db.OpenDB(args.DbFile)
	vaddrs, err := merkle.LoadAddrsFromDB(ctx, queries, args.TxType, args.Height)
	handleError(err)
	for _, v := range vaddrs {
		fmt.Printf("%s %s\n", v.Addr, v.AddrHex)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handlePanic() {
	if panicPayload := recover(); panicPayload != nil {
		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "            Fatal error. Sorry! You found a bug.")
		fmt.Fprintln(os.Stderr, "    Please copy all of this info into an issue at")
		fmt.Fprintln(os.Stderr, "     https://github.com/multisig-labs/pandasia")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintf(os.Stderr, "Version:           %s\n", version.Version)
		fmt.Fprintf(os.Stderr, "Build Date:        %s\n", version.BuildDate)
		fmt.Fprintf(os.Stderr, "Git Commit:        %s\n", version.GitCommit)
		fmt.Fprintf(os.Stderr, "Go Version:        %s\n", version.GoVersion)
		fmt.Fprintf(os.Stderr, "OS / Arch:         %s\n", version.OsArch)
		fmt.Fprintf(os.Stderr, "Panic:             %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
		os.Exit(1)
	}
}
