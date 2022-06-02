package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateContract = "create_contract"

var _ sdk.Msg = &MsgCreateContract{}

func NewMsgCreateContract(creator string, code string, value string) *MsgCreateContract {
	return &MsgCreateContract{
		Creator: creator,
		Code:    code,
		Value:   value,
	}
}

func (msg *MsgCreateContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateContract) Type() string {
	return TypeMsgCreateContract
}

func (msg *MsgCreateContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Parse and validate coins
	value, err := sdk.ParseCoinsNormalized(msg.Value)
	if err != nil || !value.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins value: %s", msg.Value)
	}

	// Validate contract code (check size limit)
	if err = ValidateCodeBasic(msg.Code); err != nil {
		return err
	}

	return nil
}
