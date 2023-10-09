package cli

import (
	"strconv"

	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAggregateCodeHashVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-code-hash-vote [application-id] [salt] [code-hash]",
		Short: "Broadcast message AggregateCodeHashVote",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argApplicationId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argSalt := args[1]
			// argValidator := args[2]
			argCodeHash := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			valAddress := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgAggregateCodeHashVote(
				argApplicationId,
				valAddress,
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
