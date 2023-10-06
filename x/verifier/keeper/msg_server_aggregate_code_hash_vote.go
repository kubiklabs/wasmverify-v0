package keeper

import (
	"context"

	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AggregateCodeHashVote(goCtx context.Context, msg *types.MsgAggregateCodeHashVote) (*types.MsgAggregateCodeHashVoteResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// derive validator address from tx sender address
	valAddr, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.ValidateValidator(ctx, valAddr); err != nil {
		return nil, err
	}

	params := ms.GetParams(ctx)
	aggregatePrevote, err := ms.GetAggregateExchangeRatePrevote(ctx, valAddr)
	if err != nil {
		return nil, types.ErrNoAggregatePrevote.Wrap(msg.Validator)
	}

	// Check the vote is submitted in the `period == prevote.period+1`
	if (uint64(ctx.BlockHeight()) / params.VotePeriod) != (aggregatePrevote.SubmitBlock/params.VotePeriod)+1 {
		return nil, types.ErrRevealPeriodMissMatch
	}

	exchangeRateTuples, err := types.ParseExchangeRateTuples(msg.ExchangeRates)
	if err != nil {
		return nil, err
	}

	// Verify that the vote hash and prevote hash match
	hash := types.GetAggregateVoteHash(msg.Salt, msg.ExchangeRates, valAddr)
	if aggregatePrevote.Hash != hash.String() {
		return nil, types.ErrVerificationFailed.Wrapf("must be given %s not %s", aggregatePrevote.Hash, hash)
	}

	// Filter out rates which aren't included in the AcceptList
	// This is also needed for slashing; in the end blocker we are checking if validator voted
	// for all required currencies. If they missed some, then we increase their slashing counter.
	filteredTuples := types.ExchangeRateTuples{}
	for _, tuple := range exchangeRateTuples {
		if params.AcceptList.Contains(tuple.Denom) {
			filteredTuples = append(filteredTuples, tuple)
		}
	}

	// Move aggregate prevote to aggregate vote with given exchange rates
	ms.SetAggregateExchangeRateVote(ctx, valAddr, types.NewAggregateExchangeRateVote(filteredTuples, valAddr))
	ms.DeleteAggregateExchangeRatePrevote(ctx, valAddr)

	return &types.MsgAggregateCodeHashVoteResponse{}, nil
}
