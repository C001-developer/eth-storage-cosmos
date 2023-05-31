package ethstorage

import (
	"eth-storage/x/ethstorage/keeper"
	"eth-storage/x/ethstorage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the storage
	for _, elem := range genState.StorageList {
		k.AppendStorage(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StorageList = k.GetAllStorage(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
