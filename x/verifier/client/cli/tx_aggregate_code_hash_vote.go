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

func CmdAggregateCodeHashVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-code-hash-vote [salt] [code-hash]",
		Short: "Broadcast message AggregateCodeHashVote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSalt := args[0]
			argCodeHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAggregateCodeHashVote(
				clientCtx.GetFromAddress().String(),
				argSalt,
				argCodeHash,
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
