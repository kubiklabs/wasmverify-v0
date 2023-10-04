package keeper

import (
	"verifier/x/verifier/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) handleResponseBlock(ctx sdk.Context, msg *types.MsgVerifyContract) uint64 {

	// compilationBlocks := k.GetTotalVerificationBlocks(ctx)
	totalReqBlocks := k.GetTotalVerificationBlocks(ctx)

	validationBlock := uint64(0)
	// total contract links uploaded so far, id
	currTotalContract := k.GetContractCount(ctx)

	var lastPendingContract types.PendingContracts
	var found bool = false
	if currTotalContract != 0 {
		lastPendingContract, found = k.GetPendingContracts(ctx, currTotalContract)
	}

	// Pending contract coun has been increased in Appendpendingcontract function
	// totalPending := k.GetPendingContractsCount(ctx)
	// k.SetPendingContractsCount(ctx, totalPending+1)

	if !found || (uint64(ctx.BlockHeight()) > lastPendingContract.GetValidationBlock()) {
		validationBlock = uint64(ctx.BlockHeight()) + compilationBlocks
	} else {
		validationBlock = lastPendingContract.GetValidationBlock() + compilationBlocks
	}

	pendingContract := types.PendingContracts{
		// code id 0 means codeId is not provided yet
		CodeId:          0,
		ContractURL:     msg.ContractURL,
		ValidationBlock: validationBlock,
		ContractURLHash: "",
	}
	id := k.AppendPendingContracts(ctx, pendingContract)

	return id
}
