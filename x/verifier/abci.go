package verifier

import (
	"time"

	"verifier/x/verifier/keeper"
	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// 1. check if it is the last block of current pending contract
	// 2. If yes, then calculate the final hash value and then update it in the state
	// 3. Check for slashing things for wrong hash provided
	// 4. Decrease the couter of the pending contract

	currentpendingContractId := k.GetCurrentPendingContractId(ctx)
	currentPendingContractFinalVerificationBlock, err := k.GetContractFinalVerificationTime(currentpendingContractId)

	if err != nil {
		return err
	}

	if uint64(ctx.BlockHeight()) == currentPendingContractFinalVerificationBlock {

	}

	// params := k.GetParams(ctx)
	// if k.IsPeriodLastBlock(ctx, params.VotePeriod) {
	// 	if err := CalcPrices(ctx, params, k); err != nil {
	// 		return err
	// 	}
	// }

	// // Slash oracle providers who missed voting over the threshold and
	// // reset miss counters of all validators at the last block of slash window
	// if k.IsPeriodLastBlock(ctx, params.SlashWindow) {
	// 	k.SlashAndResetMissCounters(ctx)
	// }

	// k.PruneAllPrices(ctx)

	return nil
}
