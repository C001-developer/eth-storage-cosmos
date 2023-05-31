package keeper

import (
	"encoding/binary"
	"encoding/hex"

	"eth-storage/x/ethstorage/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
func (k Keeper) GetStorage(ctx sdk.Context, address string, blockNumber, slot uint64) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKey))
	if blockNumber == 0 { // Get the last finalized block
		blockNumber, _ = k.GetFinalizedCount(ctx, address)
	}
	key := GetStorageKey(address, blockNumber, slot)
	bz := store.Get(key)
	if bz == nil {
		return val, false
	}
	return string(bz), true
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
	if len(address) >= 2 && address[:2] == "0x" {
		address = address[2:]
	}
	h, _ := hex.DecodeString(address)
	return h
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
	if bz == nil || len(bz) != 48 {
		return
	}
	address = "0x" + hex.EncodeToString(bz[:32])
	blockNumber = binary.BigEndian.Uint64(bz[32:40])
	slot = binary.BigEndian.Uint64(bz[40:48])
	return
}
