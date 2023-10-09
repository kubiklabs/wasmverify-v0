package keeper

import (
	"fmt"

	"verifier/x/verifier/types"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		StakingKeeper types.StakingKeeper
		wasmKeeper    types.WasmKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	StakingKeeper types.StakingKeeper,
	wasmKeeper types.WasmKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		StakingKeeper: StakingKeeper,
		wasmKeeper:    wasmKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ValidateValidator validate if the vote sender is validator
func (k Keeper) ValidateValidator(ctx sdk.Context, vAddr sdk.ValAddress) error {
	// check that the given validator exists
	if val := k.StakingKeeper.Validator(ctx, vAddr); val == nil || !val.IsBonded() {
		return stakingtypes.ErrNoValidatorFound.Wrapf("validator %s is not in active set", vAddr)
	}
	return nil
}

///////////////			PREVOTE	PREVOTE	PREVOTE	PREVOTE	PREVOTE		/////////////////////////
///////////////			PREVOTE	PREVOTE	PREVOTE	PREVOTE	PREVOTE		/////////////////////////
///////////////			PREVOTE	PREVOTE	PREVOTE	PREVOTE	PREVOTE		/////////////////////////

// GetAggregateCodeHashPrevote retrieves an oracle prevote(which is submitted hash) from the store.
func (k Keeper) GetAggregateCodeHashPrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) (types.CodeHashPreVote, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(ConcatBytes(0, types.KeyPrefixAggregateCodeHashPrevote, address.MustLengthPrefix(voter)))
	if bz == nil {
		return types.CodeHashPreVote{}, types.ErrNoAggregatePrevote.Wrap(voter.String())
	}

	var aggregatePrevote types.CodeHashPreVote
	k.cdc.MustUnmarshal(bz, &aggregatePrevote)

	return aggregatePrevote, nil
}

// HasAggregateExchangeRatePrevote checks if a validator has an existing prevote.
func (k Keeper) HasAggregateCodeHashPrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(ConcatBytes(0, types.KeyPrefixAggregateCodeHashPrevote, address.MustLengthPrefix(voter)))
}

// SetAggregateExchangeRatePrevote set an oracle aggregate prevote to the store.
func (k Keeper) SetAggregateCodeHashPrevote(
	ctx sdk.Context,
	voter sdk.ValAddress,
	prevote types.CodeHashPreVote,
) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&prevote)
	store.Set(ConcatBytes(0, types.KeyPrefixAggregateCodeHashPrevote, address.MustLengthPrefix(voter)), bz)
}

// DeleteAggregateExchangeRatePrevote deletes an oracle prevote from the store.
func (k Keeper) DeleteAggregateCodeHashPrevote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ConcatBytes(0, types.KeyPrefixAggregateCodeHashPrevote, address.MustLengthPrefix(voter)))
}

// IterateAggregateExchangeRatePrevotes iterates rate over prevotes in the store
func (k Keeper) IterateAggregateCodeHashPrevotes(
	ctx sdk.Context,
	handler func(sdk.ValAddress, types.CodeHashPreVote) bool,
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixAggregateCodeHashPrevote)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])
		var aggregatePrevote types.CodeHashPreVote
		k.cdc.MustUnmarshal(iter.Value(), &aggregatePrevote)

		if handler(voterAddr, aggregatePrevote) {
			break
		}
	}
}

///////////////			VOTE	VOTE	VOTE	VOTE	VOTE		/////////////////////////
///////////////			VOTE	VOTE	VOTE	VOTE	VOTE		/////////////////////////
///////////////			VOTE	VOTE	VOTE	VOTE	VOTE		/////////////////////////

// GetAggregateCodeHashVote retrieves vote from the store.
func (k Keeper) GetAggregateCodeHashVote(
	ctx sdk.Context,
	voter sdk.ValAddress,
) (types.CodeHashVote, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(ConcatBytes(0, types.KeyPrefixAggregateCodeHashVote, address.MustLengthPrefix(voter)))
	if bz == nil {
		return types.CodeHashVote{}, types.ErrNoAggregateVote.Wrap(voter.String())
	}

	var aggregateVote types.CodeHashVote
	k.cdc.MustUnmarshal(bz, &aggregateVote)

	return aggregateVote, nil
}

// SetAggregateExchangeRateVote adds an oracle aggregate vote to the store.
func (k Keeper) SetAggregateCodeHashVote(
	ctx sdk.Context,
	voter sdk.ValAddress,
	vote types.CodeHashVote,
) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&vote)
	store.Set(ConcatBytes(0, types.KeyPrefixAggregateCodeHashVote, address.MustLengthPrefix(voter)), bz)
}

// DeleteAggregateExchangeRateVote deletes an oracle prevote from the store.
func (k Keeper) DeleteAggregateCodeHashVote(ctx sdk.Context, voter sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(ConcatBytes(0, types.KeyPrefixAggregateCodeHashVote, address.MustLengthPrefix(voter)))
}

type IterateCodeHashVote = func(
	voterAddr sdk.ValAddress,
	aggregateVote types.CodeHashVote,
) (stop bool)

// IterateAggregateCodeHashVotes iterates rate over votes in the store.
func (k Keeper) IterateAggregateCodeHashVotes(
	ctx sdk.Context,
	handler IterateCodeHashVote,
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixAggregateCodeHashVote)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		voterAddr := sdk.ValAddress(iter.Key()[2:])
		var aggregateVote types.CodeHashVote
		k.cdc.MustUnmarshal(iter.Value(), &aggregateVote)

		if handler(voterAddr, aggregateVote) {
			break
		}
	}
}

//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////

// ClearVotes clears all tallied prevotes and votes from the store.
func (k Keeper) ClearVotes(ctx sdk.Context) {

	k.IterateAggregateCodeHashPrevotes(
		ctx,
		func(voterAddr sdk.ValAddress, aggPrevote types.CodeHashPreVote) bool {

			k.DeleteAggregateCodeHashPrevote(ctx, voterAddr)

			return false
		},
	)

	// clear all aggregate votes
	k.IterateAggregateCodeHashVotes(
		ctx,
		func(voterAddr sdk.ValAddress, _ types.CodeHashVote) bool {
			k.DeleteAggregateCodeHashVote(ctx, voterAddr)
			return false
		},
	)
}

// ConcatBytes creates a new slice by merging list of bytes and leaving empty amount of margin
// bytes at the end
func ConcatBytes(margin int, bzs ...[]byte) []byte {
	l := 0
	for _, bz := range bzs {
		l += len(bz)
	}
	out := make([]byte, l+margin)
	offset := 0
	for _, bz := range bzs {
		copy(out[offset:], bz)
		offset += len(bz)
	}
	return out
}
