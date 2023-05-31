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
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		StorageCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthstorageKeeper(t)
	ethstorage.InitGenesis(ctx, *k, genesisState)
	got := ethstorage.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StorageList, got.StorageList)
	require.Equal(t, genesisState.StorageCount, got.StorageCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
