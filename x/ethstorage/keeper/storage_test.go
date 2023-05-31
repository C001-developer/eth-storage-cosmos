package keeper_test

import (
	"fmt"
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
		items[i].Address = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
		items[i].Block = 17380596
		items[i].Slot = uint64(i)
		items[i].Value = fmt.Sprintf("%d", i)
		keeper.AppendStorage(ctx, items[i])
	}
	return items
}

func TestStorageGet(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	items := createNStorage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetStorage(ctx, item.Address, item.Block, item.Slot)
		require.True(t, found)
		require.Equal(t, item, *got)
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
