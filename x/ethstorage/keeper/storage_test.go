package keeper_test

import (
	"testing"

	keepertest "eth-storage/testutil/keeper"
	"eth-storage/testutil/nullify"
	"eth-storage/x/ethstorage/keeper"
	"eth-storage/x/ethstorage/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNStorage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Storage {
	items := make([]types.Storage, n)
	for i := range items {
		items[i].Id = keeper.AppendStorage(ctx, items[i])
	}
	return items
}

func TestStorageGet(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	items := createNStorage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetStorage(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestStorageRemove(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	items := createNStorage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStorage(ctx, item.Id)
		_, found := keeper.GetStorage(ctx, item.Id)
		require.False(t, found)
	}
}

func TestStorageGetAll(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	items := createNStorage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStorage(ctx)),
	)
}

func TestStorageCount(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	items := createNStorage(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetStorageCount(ctx))
}
