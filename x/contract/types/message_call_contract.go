package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCallContract = "call_contract"

var _ sdk.Msg = &MsgCallContract{}

func NewMsgCallContract(creator string, contractAddress string, value string, messageTerm string) *MsgCallContract {
	return &MsgCallContract{
		Creator:         creator,
		ContractAddress: contractAddress,
		Value:           value,
		MessageTerm:     messageTerm,
	}
}

func (msg *MsgCallContract) Route() string {
	return RouterKey
}

func (msg *MsgCallContract) Type() string {
	return TypeMsgCallContract
}

func (msg *MsgCallContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCallContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCallContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
