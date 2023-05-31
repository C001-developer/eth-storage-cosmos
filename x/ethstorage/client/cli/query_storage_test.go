package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"eth-storage/testutil/network"
	"eth-storage/testutil/nullify"
	"eth-storage/x/ethstorage/client/cli"
	"eth-storage/x/ethstorage/types"
)

func networkWithStorageObjects(t *testing.T, n int) (*network.Network, []types.Storage) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		storage := types.Storage{
			Address: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			Block:   17380596,
			Slot:    uint64(i),
			Value:   fmt.Sprintf("%d", i),
		}
		nullify.Fill(&storage)
		state.StorageList = append(state.StorageList, storage)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.StorageList
}

func TestShowStorage(t *testing.T) {
	net, objs := networkWithStorageObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{

		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		key  []string
		args []string
		err  error
		obj  types.Storage
	}{

		{
			desc: "found",
			key:  []string{objs[0].Address, fmt.Sprintf("%d", objs[0].Slot), fmt.Sprintf("%d", objs[0].Block)},
			args: common,
			obj:  objs[0],
		},
		{
			desc: "latest found",
			key:  []string{objs[0].Address, fmt.Sprintf("%d", objs[0].Slot)},
			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			key:  []string{"not_found"},
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := append(tc.key, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowStorage(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetStorageResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Storage)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Storage),
				)
			}
		})
	}
}
