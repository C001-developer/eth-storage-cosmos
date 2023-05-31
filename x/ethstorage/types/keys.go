package types

const (
	// ModuleName defines the module name
	ModuleName = "ethstorage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ethstorage"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	StorageKey      = "Storage/value/"
	StorageCountKey = "Storage/count/"
)
