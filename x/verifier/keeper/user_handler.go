package keeper

import (
	"encoding/binary"
	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) handleApplication(ctx sdk.Context, msg *types.MsgApplyVerifyApplication) uint64 {

	// compilationBlocks := k.GetTotalVerificationBlocks(ctx)
	totalVerificationBlocksReq := k.GetTotalVerificationBlocks(ctx)

	finalverificationBlockHeight := uint64(0)

	// total contract links uploaded so far, id
	TotalContract := k.GetContractCount(ctx)

	var lastContractPending types.ContractInfo
	var found bool = false
	if TotalContract != 0 {
		lastContractPending, found = k.GetContractInfo(ctx, TotalContract)
	}

	if !found || (uint64(ctx.BlockHeight()) > lastContractPending.GetAssignedVerificationBlockHeight()) {
		finalverificationBlockHeight = uint64(ctx.BlockHeight()) + totalVerificationBlocksReq
	} else {
		finalverificationBlockHeight = lastContractPending.GetAssignedVerificationBlockHeight() + totalVerificationBlocksReq
	}

	ContractInfo := types.ContractInfo{
		// code id 0 means codeId is not provided yet
		CodeId:                          0,
		OffchainCodeUrl:                 msg.OffchainCodeUrl,
		AssignedVerificationBlockHeight: finalverificationBlockHeight,
		OffchainCodeHash:                "",
	}
	id := k.AppendPendingContracts(ctx, ContractInfo)

	return id
}

// SetPendingContracts set a specific pendingContracts in the store
func (k Keeper) SetContractInfo(ctx sdk.Context, pendingContracts types.ContractInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractInfoKey))
	b := k.cdc.MustMarshal(&pendingContracts)
	store.Set(GetPendingContractsIDBytes(pendingContracts.Id), b)
}

// GetPendingContracts returns a pendingContracts from its id
func (k Keeper) GetContractInfo(ctx sdk.Context, id uint64) (val types.ContractInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractInfoKey))
	b := store.Get(GetPendingContractsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetPendingContractsCount get the total number of pendingContracts
func (k Keeper) GetPendingContractsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PendingContractsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPendingContractsCount set the total number of pendingContracts
func (k Keeper) SetPendingContractsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PendingContractsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPendingContracts appends a pendingContracts in the store with a new id and update the count
func (k Keeper) AppendPendingContracts(
	ctx sdk.Context,
	pendingContracts types.ContractInfo,
) uint64 {
	// Create the pendingContracts
	count := k.GetContractCount(ctx)
	pendingCount := k.GetPendingContractsCount(ctx)

	// Set the ID of the appended value
	pendingContracts.Id = count + 1

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractInfoKey))
	appendedValue := k.cdc.MustMarshal(&pendingContracts)
	store.Set(GetPendingContractsIDBytes(pendingContracts.Id), appendedValue)

	// Update pendingContracts count
	k.SetPendingContractsCount(ctx, pendingCount+1)
	k.SetContractCount(ctx, count+1)

	return count + 1
}

// get final block verification time or the contract by its id
func (k Keeper) GetContractFinalVerificationTime(
	ctx sdk.Context,
	id uint64,
) (uint64, error) {
	contractInfo, found := k.GetContractInfo(ctx, id)
	if !found {
		return 0, types.ErrContractInfoNotFound
	}

	return contractInfo.AssignedVerificationBlockHeight, nil
}

// Returns last block till validator can prevote for the contract
func (k Keeper) GetContractPrevoteTime(
	ctx sdk.Context,
	id uint64,
) (uint64, error) {
	contractInfo, found := k.GetContractInfo(ctx, id)
	if !found {
		return 0, types.ErrContractInfoNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	VoteBlockskey := types.KeyPrefix(types.VoteBlockTimeKey)

	bz1 := store.Get(VoteBlockskey)

	// Count doesn't exist: no element
	if bz1 == nil {
		return 0, types.ErrVoteBlockTimeNotFound
	}

	// Parse bytes
	return (contractInfo.AssignedVerificationBlockHeight - binary.BigEndian.Uint64(bz1)), nil

}

// GetPendingContractsIDBytes returns the byte representation of the ID
func GetPendingContractsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPendingContractsIDFromBytes returns ID in uint64 format from a byte array
func GetPendingContractsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
