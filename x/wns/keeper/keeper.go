package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

type WnsKeeper interface {
	SetRecord(sdk.Context, types.Node, types.Record) error
	IsOwner(sdk.Context, types.Node, string) bool
	SetNodeOwner(sdk.Context, types.Node, string) error
	SetResolver(sdk.Context, types.Node, string) error
	SetTTL(sdk.Context, types.Node, time.Time) error
}

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	memeKey  storetypes.StoreKey

	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper
}

func (k Keeper) SetRecord(ctx sdk.Context, node types.Node, record types.Record) error {
	if node.IsValid() && record.IsValid() {
		store := ctx.KVStore(k.storeKey)
		bz := k.cdc.MustMarshal(record)
		store.Set(types.KeyPrefixNodeToRecord(node), bz)
		return nil
	}

	return fmt.Errorf("either node or record is invalid")
}

func (k Keeper) IsOwner(ctx, sdk.Context, node types.Node, address string) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(node[:])
	var record *types.Record
	k.cdc.MustUnmarshal(bz, record)
	if record == nil {
		return false
	}

	if record.Owner == address {
		return true
	}

	return false
}
