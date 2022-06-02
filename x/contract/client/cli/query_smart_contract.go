package cli

import (
	"context"

	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListSmartContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-smart-contract",
		Short: "list all smart-contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSmartContractRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SmartContractAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSmartContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-smart-contract [address]",
		Short: "shows a smart-contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			params := &types.QueryGetSmartContractRequest{
				Address: argAddress,
			}

			res, err := queryClient.SmartContract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
