package pchain

import (
	"runtime"
	"testing"

	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	t.Logf("%d", runtime.NumCPU())
	t.Fatal()
	blks, err := FetchBlocks("http://100.83.243.106:9650", 0, 10000)
	require.NoError(t, err)
	t.Logf("Blocks: %d", len(blks))
	// spew.Dump(blks)

	t.Fatal()
}

func Test2(t *testing.T) {

	blks, err := SuperFetchBlocks("http://100.83.243.106:9650", 0, 10000)
	require.NoError(t, err)
	t.Logf("Blocks: %d", len(blks))
	// spew.Dump(blks)

	t.Fatal()
}

func Test3(t *testing.T) {
	block_16 := "0x000000000000b12e01beb62231fb0eaea0e0d1a1b232132c5170a8a95c5ab63eac727f5c8d3200000000000000100000000e0000000100000000000000000000000000000000000000000000000000000000000000000000000121e67317cbc4be2aeb00677ad6462778a8f52274b9d605df2591b23027a87dff000000160000000060bd61800000000700000a3586233a00000000000000000000000001000000013babddb84ae529d688d2750199b861c23ddddd80000000010000000000000000000000000000000000000000000000000000000000000000000003b821e67317cbc4be2aeb00677ad6462778a8f52274b9d605df2591b23027a87dff000000150000000060bd61800000000500000a3b5840f40000000001000000000000000400000000d1a56f61f9985df2a737e246e8c0783b153d6a9c000000005f6962c4000000005f7d3d3000000005d21dba000000000121e67317cbc4be2aeb00677ad6462778a8f52274b9d605df2591b23027a87dff000000160000000060bd61800000000700000005d21dba00000000000000000000000001000000013babddb84ae529d688d2750199b861c23ddddd800000000b000000000000000000000001000000013babddb84ae529d688d2750199b861c23ddddd8000000001000000090000000109a8e63464561d2d629941851488551dcfdbb55fc68b72ff2771b8adacec8eb87ed4d7d435bbe7b8c536608a04b0da3d8b7a3c021ecdf4ffb218d46f508f23f600a44a0ff2"
	b, _ := formatting.Decode(formatting.Hex, block_16)
	blk, _ := NewBlock(b)
	spew.Dump(blk)
	t.Fatal()
}
