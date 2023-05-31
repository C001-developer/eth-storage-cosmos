package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// FetchStorage fetches the storage from the ethereum blockchain
func (k *Keeper) FetchStorage(ctx sdk.Context, address string, blockNumber, slot uint64) (string, bool) {
	return "", false
}
