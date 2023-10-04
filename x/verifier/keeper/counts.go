package keeper

import (
	"encoding/binary"

	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetTotalVerificationBlocks(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	compilationBlockskey := types.KeyPrefix(types.ContractCompilationBlockTimeKey)
	PrevoteBLockskey := types.KeyPrefix(types.PrevoteBlockTimeKey)
	VoteBlockskey := types.KeyPrefix(types.VoteBlockTimeKey)

	bz1 := store.Get(compilationBlockskey)
	bz2 := store.Get(PrevoteBLockskey)
	bz3 := store.Get(VoteBlockskey)

	// Count doesn't exist: no element
	if bz1 == nil && bz2 == nil && bz3 == nil {
		return 0
	}

	// Parse bytes
	return (binary.BigEndian.Uint64(bz1) + binary.BigEndian.Uint64(bz2) + binary.BigEndian.Uint64(bz3))
}

func (k Keeper) SetTotalVerificationBlocks(ctx sdk.Context, compileblocks uint64, prevoteblocks uint64, voteblockks uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	compilationBlockskey := types.KeyPrefix(types.ContractCompilationBlockTimeKey)
	PrevoteBLockskey := types.KeyPrefix(types.PrevoteBlockTimeKey)
	VoteBlockskey := types.KeyPrefix(types.VoteBlockTimeKey)

	bz1 := make([]byte, 8)
	bz2 := make([]byte, 8)
	bz3 := make([]byte, 8)

	binary.BigEndian.PutUint64(bz1, compileblocks)
	binary.BigEndian.PutUint64(bz2, prevoteblocks)
	binary.BigEndian.PutUint64(bz3, voteblockks)

	store.Set(compilationBlockskey, bz1)
	store.Set(PrevoteBLockskey, bz2)
	store.Set(VoteBlockskey, bz3)
}

func (k Keeper) GetContractCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetContractCount(ctx sdk.Context, block uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, block)
	store.Set(byteKey, bz)
}
