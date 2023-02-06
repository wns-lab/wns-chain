package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
	SetAccount(sdk.Context, authtypes.AccountI)
	GetModuleAddress(name string) sdk.AccAddress
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
