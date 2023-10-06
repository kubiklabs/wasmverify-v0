package cli

import (
	"strconv"

	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAggregateCodeHashPrevote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-code-hash-prevote [hash]",
		Short: "Broadcast message AggregateCodeHashPrevote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argApplicationId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argHash := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAggregateCodeHashPrevote(
				argApplicationId,
				clientCtx.GetFromAddress().String(),
				argHash,
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
