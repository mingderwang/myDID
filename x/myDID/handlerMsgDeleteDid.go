package myDID

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/myDID/x/myDID/types"
	"github.com/mingderwang/myDID/x/myDID/keeper"
)

// Handle a message to delete name
func handleMsgDeleteDid(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteDid) (*sdk.Result, error) {
	if !k.DidExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetDidOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteDid(ctx, msg.ID)
	return &sdk.Result{}, nil
}
