package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"verifier/x/verifier/types"

	"github.com/spf13/cast"
)

func CmdListVerifiedMetaData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-verified-meta-data",
		Short: "list all verifiedMetaData",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVerifiedMetaDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VerifiedMetaDataAll(cmd.Context(), params)
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

func CmdShowVerifiedMetaData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-verified-meta-data [code-id]",
		Short: "shows a verifiedContractMetaData",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argCodeId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetVerifiedMetaDataRequest{
				CodeId: argCodeId,
			}

			res, err := queryClient.VerifiedMetaData(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
