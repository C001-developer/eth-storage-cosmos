package keeper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FetchStorage fetches the storage from the ethereum blockchain
func (k *Keeper) FetchStorage(ctx sdk.Context, address string, blockNumber, slot uint64) (string, bool) {
	var keyBuf common.Hash
	big.NewInt(int64(slot)).FillBytes(keyBuf[:])
	res, err := k.ethClient.StorageAt(ctx, common.HexToAddress(address), keyBuf, big.NewInt(int64(blockNumber)))
	if err != nil {
		if err == rpc.ErrNoResult {
			return "", true
		}
		return "", false
	}
	return common.BytesToHash(res).String(), false
}
