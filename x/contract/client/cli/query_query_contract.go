package cli

import (
	"strconv"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-contract [contract-address] [query] [n-derivations]",
		Short: "Query query-contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqContractAddress := args[0]
			reqQuery := args[1]
			reqNDerivations, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQueryContractRequest{

				ContractAddress: reqContractAddress,
				Query:           reqQuery,
				NDerivations:    reqNDerivations,
			}

			res, err := queryClient.QueryContract(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
