package ethstorage

import (
	"eth-storage/x/ethstorage/keeper"
	"eth-storage/x/ethstorage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	for _, address := range params.Addresses {
		// since the previous block is finalized, we can set the last block number and slot as finalized
		blockNumber, slot := k.GetLastCount(ctx, address)
		if blockNumber != 0 {
			k.SetFinalizedCount(ctx, address, blockNumber, slot)
		} else {
			blockNumber = params.FromBlock
			slot = 0
		}
		// fetch the storage slot from the ethereum blockchain
		for count := params.MaxCount; count > 0; count-- {
			slot++
			value, over := k.FetchStorage(ctx, address, blockNumber, slot)
			if over {
				break
			}
			if len(value) == 0 {
				slot = 0
				blockNumber++
				continue
			}
			// append the storage slot to the store
			k.AppendStorage(ctx, types.Storage{
				Address: address,
				Block:   blockNumber,
				Slot:    slot,
				Value:   value,
			})
		}
	}
}
