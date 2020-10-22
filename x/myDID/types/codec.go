package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateDid{}, "myDID/CreateDid", nil)
		cdc.RegisterConcrete(MsgSetDid{}, "myDID/SetDid", nil)
		cdc.RegisterConcrete(MsgDeleteDid{}, "myDID/DeleteDid", nil)
		cdc.RegisterConcrete(MsgCreateUser{}, "myDID/CreateUser", nil)
		cdc.RegisterConcrete(MsgSetUser{}, "myDID/SetUser", nil)
		cdc.RegisterConcrete(MsgDeleteUser{}, "myDID/DeleteUser", nil)
		cdc.RegisterConcrete(MsgCreatePost{}, "myDID/CreatePost", nil)
		cdc.RegisterConcrete(MsgSetPost{}, "myDID/SetPost", nil)
		cdc.RegisterConcrete(MsgDeletePost{}, "myDID/DeletePost", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
