package keeper

import (
	"eth-storage/x/ethstorage/types"
)

var _ types.QueryServer = Keeper{}
