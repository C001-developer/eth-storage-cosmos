package keeper

import (
	"encoding/binary"

	"eth-storage/x/ethstorage/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// GetFinalizedCount get the last finalized slot
func (k Keeper) GetFinalizedCount(ctx sdk.Context, address string) (blockNumber, slot uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageCountKey))
	iterator := sdk.KVStorePrefixIterator(store, GetBytesFromAddress(address))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		_, blockNumber, slot = GetStorageFromBytes(key)
	}

	return
}

// SetStorageCount set the total number of storage
func (k Keeper) SetFinalizedCount(ctx sdk.Context, address string, blockNumber, slot uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageCountKey))
	iterator := sdk.KVStorePrefixIterator(store, GetBytesFromAddress(address))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		store.Delete(key)
	}

	// Set the current finalized slot
	key := GetStorageKey(address, blockNumber, slot)
	store.Set(key, []byte{})
}

// AppendStorage appends a new storage in the store
func (k Keeper) AppendStorage(
	ctx sdk.Context,
	storage types.Storage,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKey))
	key := GetStorageKey(storage.Address, storage.Block, storage.Slot)
	store.Set(key, []byte(storage.Value))
}

// GetLastCount get the last block number and slot without considering finalized blocks
func (k Keeper) GetLastCount(ctx sdk.Context, address string) (blockNumber, slot uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKey))
	iterator := sdk.KVStoreReversePrefixIterator(store, GetBytesFromAddress(address))

	defer iterator.Close()

	if iterator.Valid() {
		key := iterator.Key()
		_, blockNumber, slot = GetStorageFromBytes(key)
		return
	}

	return
}

// GetStorage returns a storage from
func (k Keeper) GetStorage(ctx sdk.Context, address string, blockNumber, slot uint64) (storage *types.Storage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKey))
	fBlockNumber, fSlot := k.GetFinalizedCount(ctx, address)
	if blockNumber == 0 { // Get the last finalized block
		if slot <= fSlot {
			blockNumber = fBlockNumber
		} else {
			blockNumber = fBlockNumber - 1
		}
	} else {
		if blockNumber > fBlockNumber {
			return nil, false
		}
		if blockNumber == fBlockNumber && slot > fSlot {
			return nil, false
		}
	}

	key := GetStorageKey(address, blockNumber, slot)
	bz := store.Get(key)
	if bz == nil {
		return nil, false
	}
	return &types.Storage{
		Address: address,
		Block:   blockNumber,
		Slot:    slot,
		Value:   string(bz),
	}, true
}

// GetAllStorage returns the list of all storage
func (k Keeper) GetAllStorage(ctx sdk.Context) (list []types.Storage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKey))
	iterator := sdk.KVStorePrefixIterator(store, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, blockNumber, slot := GetStorageFromBytes(key)
		val := store.Get(key)
		list = append(list, types.Storage{
			Address: address,
			Block:   blockNumber,
			Slot:    slot,
			Value:   string(val),
		})
	}

	return
}

// GetBytesFromAddress returns the byte representation of the address
func GetBytesFromAddress(address string) []byte {
	return common.HexToAddress(address).Bytes()
}

// GetStorageIDBytes returns the byte representation of the storage key
func GetStorageKey(address string, blockNumber, slot uint64) []byte {
	bn := make([]byte, 8)
	binary.BigEndian.PutUint64(bn, blockNumber)
	sl := make([]byte, 8)
	binary.BigEndian.PutUint64(sl, slot)
	h := GetBytesFromAddress(address)
	h = append(h, bn...)
	h = append(h, sl...)
	return h
}

// GetStorageFromBytes returns ID in uint64 format from a byte array
func GetStorageFromBytes(bz []byte) (address string, blockNumber uint64, slot uint64) {
	if bz == nil || len(bz) != 36 {
		return
	}
	address = common.BytesToAddress(bz[:20]).Hex()
	blockNumber = binary.BigEndian.Uint64(bz[20:28])
	slot = binary.BigEndian.Uint64(bz[28:36])
	return
}
