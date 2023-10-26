package keeper

import (
	"verifier/x/verifier/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVerifiedMetaData set a specific verifiedMetaData in the store from its index
func (k Keeper) SetVerifiedMetaData(ctx sdk.Context, verifiedMetaData types.VerifiedMetaData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedMetaDataKeyPrefix))
	b := k.cdc.MustMarshal(&verifiedMetaData)
	store.Set(types.VerifiedMetaDataKey(
		verifiedMetaData.CodeId,
	), b)
}

// GetVerifiedMetaData returns a verifiedMetaData from its index
func (k Keeper) GetVerifiedMetaData(
	ctx sdk.Context,
	codeId uint64,

) (val types.VerifiedMetaData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedMetaDataKeyPrefix))

	b := store.Get(types.VerifiedMetaDataKey(
		codeId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVerifiedMetaData removes a verifiedMetaData from the store
func (k Keeper) RemoveVerifiedMetaData(
	ctx sdk.Context,
	codeId uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedMetaDataKeyPrefix))
	store.Delete(types.VerifiedMetaDataKey(
		codeId,
	))
}

// GetAllVerifiedMetaData returns all verifiedMetaData
func (k Keeper) GetAllVerifiedMetaData(ctx sdk.Context) (list []types.VerifiedMetaData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedMetaDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerifiedMetaData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
