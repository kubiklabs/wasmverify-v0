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

func CmdUpdateBlockTime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-block-time [compilation-blocks] [prevote-blocks] [vote-blocks]",
		Short: "Broadcast message UpdateBlockTime",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCompilationBlocks, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argPrevoteBlocks, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argVoteBlocks, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBlockTime(
				clientCtx.GetFromAddress().String(),
				argCompilationBlocks,
				argPrevoteBlocks,
				argVoteBlocks,
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
