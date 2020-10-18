package myDID

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mingderwang/myDID/x/myDID/types"
	"github.com/mingderwang/myDID/x/myDID/keeper"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUser) (*sdk.Result, error) {
	var user = types.User{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
    	Email: msg.Email,
	}
	k.CreateUser(ctx, user)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
