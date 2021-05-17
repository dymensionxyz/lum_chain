package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgOpenBeam{}

// NewMsgOpenBeam Build a open beam message based on parameters
func NewMsgOpenBeam(id string, creator string, amount int64, secret string, schema string, reward *BeamSchemeReward, review *BeamSchemeReview) *MsgOpenBeam {
	return &MsgOpenBeam{
		Id:      id,
		Creator: creator,
		Amount: amount,
		Secret:  secret,
		Schema: schema,
		Reward:  reward,
		Review:  review,
	}
}

// Route dunno
func (msg MsgOpenBeam) Route() string {
	return RouterKey
}

// Type Return the message type
func (msg MsgOpenBeam) Type() string {
	return "OpenBeam"
}

// GetSigners Return the list of signers for the given message
func (msg *MsgOpenBeam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes Return the generated bytes from the signature
func (msg *MsgOpenBeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Validate the message payload before dispatching to the local kv store
func (msg *MsgOpenBeam) ValidateBasic() error {
	if len(msg.Id) <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid id supplied (%d)", len(msg.Id))
	}

	// Ensure the address is correct and that we are able to acquire it
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	// Validate the secret
	if len(msg.Secret) <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid secret supplied")
	}

	// Validate the schema
	if msg.GetSchema() != BEAM_SCHEMA_REVIEW && msg.GetSchema() != BEAM_SCHEMA_REWARD {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid schema must be review or reward")
	}
	return nil
}
