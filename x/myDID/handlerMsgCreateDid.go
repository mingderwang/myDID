package myDID

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mingderwang/myDID/x/myDID/types"
	"github.com/mingderwang/myDID/x/myDID/keeper"
)

func handleMsgCreateDid(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateDid) (*sdk.Result, error) {
	var did = types.Did{
		Creator: msg.Creator,
		ID:      msg.ID,
	}
	k.CreateDid(ctx, did)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
