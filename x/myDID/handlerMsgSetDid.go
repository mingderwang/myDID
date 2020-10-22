package myDID

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mingderwang/myDID/x/myDID/types"
	"github.com/mingderwang/myDID/x/myDID/keeper"
)

func handleMsgSetDid(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetDid) (*sdk.Result, error) {
	var did = types.Did{
		Creator: msg.Creator,
		ID:      msg.ID,
	}
	if !msg.Creator.Equals(k.GetDidOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetDid(ctx, did)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
