package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"

	"runtime/debug"

	"github.com/jxskiss/mcli"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/syncer"
	"github.com/multisig-labs/pandasia/pkg/version"
	_ "modernc.org/sqlite"
)

func main() {
	defer handlePanic()
	mcli.Add("sync-pchain", syncPchainCmd, "Sync the P-Chain to the db")

	mcli.Add("generate", generateCmd, "Generate the Merkle Tree")
	mcli.AddHelp()
	mcli.AddCompletion()
	mcli.Run()
}

func syncPchainCmd() {
	args := struct {
		NodeURL string `cli:"--node-url, Avalanche node URL" default:"http://localhost:9650"`
		DbFile  string `cli:"--db, SQLite database file name" default:"pandasia.db"`
	}{}
	mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	ctx := context.Background()
	_, queries := openDB(args.DbFile)
	err := syncer.SyncPChain(ctx, queries, args.NodeURL)
	handleError(err)
}

func generateCmd() {
	// args := struct {
	// 	DbFile string `cli:"--db, SQLite database file name" default:"pandasia.db"`
	// }{}
	// mcli.Parse(&args, mcli.WithErrorHandling(flag.ExitOnError))

	// ctx := context.Background()
	// _, queries := openDB(args.DbFile)
	// height, err := queries.MaxHeight(ctx)
	// handleError(err)
	// vaddrs, err := loadDataFromDB(ctx, queries, pchain.AddValidatorTxId, int(height))
	// handleError(err)

	// cfg := mt.Config{SortSiblingPairs: true, DisableLeafHashing: true, Mode: mt.ModeProofGenAndTreeBuild}
	// tree, err := mt.New(&cfg, vaddrs)
	// handleError(err)

	// // rootArgs := db.CreateMerkleRootAndReturnIdParams{
	// // 	Height: height,
	// // 	Type:   "validator",
	// // 	Root:   fmt.Sprintf("0x%x", tree.Root),
	// // }

	// // rootId, err := queries.CreateMerkleRootAndReturnId(ctx, rootArgs)
	// // handleError(err)

	// rootHash := tree.Root
	// fmt.Printf("Merkle Root: %x\n", rootHash)

	// proofs := tree.Proofs
	// for i := 0; i < len(proofs); i++ {
	// 	_, err := tree.Verify(vaddrs[i], proofs[i])
	// 	handleError(err)
	// 	fmt.Printf("Value: %x  Proof: %x", vaddrs[i], proofs[i])
	// 	// v, _ := vaddrs[i].(*validatorAddress)
	// 	// proofArgs := db.CreateMerkleProofParams{
	// 	// 	MerkleRootID: rootId,
	// 	// 	Paddy:         v.Addr,
	// 	// 	Data:         fmt.Sprintf("0x%x", v.Data),
	// 	// 	Proof:        fmt.Sprintf("0x%x", v.Proof),
	// 	// }
	// }
}

func loadAddrsFromDB(ctx context.Context, queries *db.Queries, txType int, height int) ([]string, error) {
	vaddrs := []string{}
	// args := db.FindAddrsForMerkleTreeParams{
	// 	TypeID: int64(txType),
	// 	Height: int64(height),
	// }
	// addrs, err := queries.FindAddrsForMerkleTree(ctx, args)
	// if err != nil {
	// 	return nil, err
	// }
	// for _, addr := range addrs {
	// 	// Parse P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww into just the address bytes
	// 	_, _, addrBytes, err := address.Parse(addr.String)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	vaddr := validatorAddress{Addr: addr.String, Data: addrBytes}
	// 	vaddrs = append(vaddrs, &vaddr)
	// }
	return vaddrs, nil
}

func openDB(dbFileName string) (*sql.DB, *db.Queries) {
	dbFile, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		panic(err)
	}
	_, err = dbFile.Exec("PRAGMA optimize;PRAGMA foreign_keys=ON;PRAGMA journal_mode=WAL;")
	if err != nil {
		panic(err)
	}
	return dbFile, db.New(dbFile)
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
