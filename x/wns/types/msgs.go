package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgSetMetaData{}
	_ sdk.Msg = &MsgSetOwner{}
	_ sdk.Msg = &MsgSetResolver{}
	_ sdk.Msg = &MsgSetTTL{}
)

// ValidateBasic runs stateless checks on the message
func (msg MsgSetMetaData) ValidateBasic() error {
	name := &DomainName{Name: msg.Name}
	if !name.Valid() {
		return fmt.Errorf("invalid name %v or record %v", msg.Name, msg.Metadata)
	}

	return nil
}

// GetSigners defines whose signature is required
func (msg MsgSetMetaData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetOwner) ValidateBasic() error {
	name := &DomainName{Name: msg.Name}
	if !name.Valid() {
		return fmt.Errorf("invalid name %v", msg.Name)
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgSetOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetResolver) ValidateBasic() error {
	name := &DomainName{Name: msg.Name}
	if !name.Valid() {
		return fmt.Errorf("invalid name %v", msg.Name)
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgSetResolver) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetTTL) ValidateBasic() error {
	name := &DomainName{Name: msg.Name}
	if !name.Valid() {
		return fmt.Errorf("invalid name %v", msg.Name)
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgSetTTL) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}
