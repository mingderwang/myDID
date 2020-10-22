package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/mingderwang/myDID/x/myDID/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for myDID clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListDid:
			return listDid(ctx, k)
		case types.QueryGetDid:
			return getDid(ctx, path[1:], k)
		case types.QueryListUser:
			return listUser(ctx, k)
		case types.QueryGetUser:
			return getUser(ctx, path[1:], k)
		case types.QueryListPost:
			return listPost(ctx, k)
		case types.QueryGetPost:
			return getPost(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown myDID query endpoint")
		}
	}
}
