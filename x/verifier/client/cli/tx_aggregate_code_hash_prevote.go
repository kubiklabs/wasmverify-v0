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

func CmdAggregateCodeHashPrevote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-code-hash-prevote [application-id] [Codehash]",
		Short: "Broadcast message AggregateCodeHashPrevote",
		Args:  cobra.ExactArgs(2),
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
			valAddress := sdk.ValAddress(clientCtx.GetFromAddress())

			msg := types.NewMsgAggregateCodeHashPrevote(
				argApplicationId,
				valAddress,
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
