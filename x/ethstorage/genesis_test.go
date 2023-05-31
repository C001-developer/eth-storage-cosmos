package ethstorage_test

import (
	"testing"

	keepertest "eth-storage/testutil/keeper"
	"eth-storage/testutil/nullify"
	"eth-storage/x/ethstorage"
	"eth-storage/x/ethstorage/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StorageList: []types.Storage{
			{
				Address: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				Block:   17380596,
				Slot:    0,
				Value:   "0x00000000",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthstorageKeeper(t)
	ethstorage.InitGenesis(ctx, *k, genesisState)
	got := ethstorage.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StorageList, got.StorageList)
	// this line is used by starport scaffolding # genesis/test/assert
}
