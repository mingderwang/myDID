package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateDid{}

type MsgCreateDid struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgCreateDid(creator sdk.AccAddress) MsgCreateDid {
  return MsgCreateDid{
    ID: uuid.New().String(),
		Creator: creator,
	}
}

func (msg MsgCreateDid) Route() string {
  return RouterKey
}

func (msg MsgCreateDid) Type() string {
  return "CreateDid"
}

func (msg MsgCreateDid) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateDid) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateDid) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}