package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"verifier/x/verifier/types"
)

var _ = strconv.Itoa(0)

func CmdFinalVerification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "final-verification [application-id] [code-id]",
		Short: "Broadcast message FinalVerification",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argApplicationId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argCodeId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFinalVerification(
				clientCtx.GetFromAddress().String(),
				argApplicationId,
				argCodeId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
