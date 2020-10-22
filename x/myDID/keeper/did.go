package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/myDID/x/myDID/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateDid creates a did
func (k Keeper) CreateDid(ctx sdk.Context, did types.Did) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.DidPrefix + did.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(did)
	store.Set(key, value)
}

// GetDid returns the did information
func (k Keeper) GetDid(ctx sdk.Context, key string) (types.Did, error) {
	store := ctx.KVStore(k.storeKey)
	var did types.Did
	byteKey := []byte(types.DidPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &did)
	if err != nil {
		return did, err
	}
	return did, nil
}

// SetDid sets a did
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) {
	didKey := did.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(did)
	key := []byte(types.DidPrefix + didKey)
	store.Set(key, bz)
}

// DeleteDid deletes a did
func (k Keeper) DeleteDid(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.DidPrefix + key))
}

//
// Functions used by querier
//

func listDid(ctx sdk.Context, k Keeper) ([]byte, error) {
	var didList []types.Did
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.DidPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var did types.Did
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &did)
		didList = append(didList, did)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, didList)
	return res, nil
}

func getDid(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	did, err := k.GetDid(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, did)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetDidOwner(ctx sdk.Context, key string) sdk.AccAddress {
	did, err := k.GetDid(ctx, key)
	if err != nil {
		return nil
	}
	return did.Creator
}


// Check if the key exists in the store
func (k Keeper) DidExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.DidPrefix + key))
}
