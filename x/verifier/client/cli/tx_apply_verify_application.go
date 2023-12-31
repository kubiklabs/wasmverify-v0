package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"verifier/x/verifier/types"
)

var _ = strconv.Itoa(0)

func CmdApplyVerifyApplication() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply-verify-application [offchain-code-url]",
		Short: "Broadcast message ApplyVerifyApplication",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOffchainCodeUrl := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApplyVerifyApplication(
				clientCtx.GetFromAddress().String(),
				argOffchainCodeUrl,
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
