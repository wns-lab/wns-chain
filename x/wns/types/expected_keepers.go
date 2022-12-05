package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AuthKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI
	SetAccount(sdk.Context, types.AccountI)
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
