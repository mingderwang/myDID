package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteDid{}

type MsgDeleteDid struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteDid(id string, creator sdk.AccAddress) MsgDeleteDid {
  return MsgDeleteDid{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteDid) Route() string {
  return RouterKey
}

func (msg MsgDeleteDid) Type() string {
  return "DeleteDid"
}

func (msg MsgDeleteDid) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteDid) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteDid) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}