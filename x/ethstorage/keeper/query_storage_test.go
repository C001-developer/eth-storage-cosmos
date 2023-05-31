package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "eth-storage/testutil/keeper"
	"eth-storage/testutil/nullify"
	"eth-storage/x/ethstorage/types"
)

func TestStorageQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EthstorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStorage(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetStorageRequest
		response *types.QueryGetStorageResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStorageRequest{Address: msgs[0].Address, Block: msgs[0].Block, Slot: msgs[0].Slot},
			response: &types.QueryGetStorageResponse{Storage: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetStorageRequest{Address: msgs[1].Address, Block: msgs[1].Block, Slot: msgs[1].Slot},
			response: &types.QueryGetStorageResponse{Storage: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetStorageRequest{Address: msgs[1].Address, Block: msgs[1].Block, Slot: msgs[1].Slot + 1},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Storage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
