package cli

import (
	"context"
	"errors"
	"strconv"

	"eth-storage/x/ethstorage/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-storage [id]",
		Short: "shows a storage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("not enough arguments")
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			slot, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			var blockNumber uint64
			if len(args) > 2 {
				blockNumber, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}
			}
			params := &types.QueryGetStorageRequest{
				Address: args[0],
				Slot:    slot,
				Block:   blockNumber,
			}

			res, err := queryClient.Storage(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
