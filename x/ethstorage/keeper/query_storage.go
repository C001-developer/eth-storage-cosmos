package keeper

import (
	"context"

	"eth-storage/x/ethstorage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Storage(goCtx context.Context, req *types.QueryGetStorageRequest) (*types.QueryGetStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetStorage(ctx, req.Address, req.Block, req.Slot)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetStorageResponse{Storage: types.Storage{
		Address: req.Address,
		Block:   req.Block,
		Slot:    req.Slot,
		Value:   val,
	}}, nil
}
