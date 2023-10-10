package verifier

import (
	"fmt"
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
	if currentpendingContractId == 0 {
		return nil
	}

	// currentPendingContractFinalVerificationBlock, err := k.GetContractFinalVerificationTime(currentpendingContractId)
	currentPendingContractInfo, found := k.GetContractInfo(ctx, currentpendingContractId)

	if !found {
		return types.ErrContractInfoNotFound
	}

	if uint64(ctx.BlockHeight()) == currentPendingContractInfo.AssignedVerificationBlockHeight {
		finalizeHashForthisApplicationId := calculateStakeWeightedFinalHash(ctx, k)
		if finalizeHashForthisApplicationId == "" {
			return types.ErrInvalidHash
		}
		currentPendingContractInfo.OffchainCodeHash = finalizeHashForthisApplicationId

		k.SetContractInfo(ctx, currentPendingContractInfo)
	}

	return nil
}

// Also handle the slashing for different hash provided
func calculateStakeWeightedFinalHash(ctx sdk.Context, k keeper.Keeper) string {
	// Build claim map over all validators in active set
	fmt.Println("Started calculateweightedhash")
	validatorDataMap := make(map[string]types.ValidatorData)
	powerReduction := k.StakingKeeper.PowerReduction(ctx)
	// Calculate total validator power
	var totalBondedPower int64
	for _, v := range k.StakingKeeper.GetBondedValidatorsByPower(ctx) {
		// operator is the validator itself in our case
		addr := v.GetOperator()
		power := v.GetConsensusPower(powerReduction)
		totalBondedPower += power
		validatorDataMap[addr.String()] = types.NewValidatorData(power, addr)
	}

	// Here we have a map consisting all the validators and the voting power of the validator
	// Now create a  map iterate over all the votes and store the hash in the map with hash-> total power received on this hash
	// We can set a threshold for the voting to be done or that much of voting atleast received from that vote
	// Save that hash in the state and we will the validators who have not voted and can be slashed
	voteOnHashMap := make(map[string]uint64)
	finalizeHashhandler := func(voterAddr sdk.ValAddress, vote types.CodeHashVote) bool {
		_, ok := validatorDataMap[vote.Voter]
		if ok {
			voteOnHashMap[vote.CodeHash] += uint64(validatorDataMap[vote.Voter].Power)
		}
		return false
	}
	k.IterateAggregateCodeHashVotes(ctx, finalizeHashhandler)

	fmt.Println("Iterated on all the prevotes and calculated hash and their weight")

	var finalHash string = ""
	var associatedHashWt uint64 = 0
	// finalize average by dividing by total stake, i.e. total weights
	// What top two  weight is equal, then ambigious
	for hash, weight := range voteOnHashMap {
		if uint64(associatedHashWt) < weight {
			associatedHashWt = weight
			finalHash = hash
		}
	}
	// Slash all the validators who didn't vote

	// remove all the prevote and vote
	k.ClearVotes(ctx)
	fmt.Println("Cleared votes")

	return finalHash
}
