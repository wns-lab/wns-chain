package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

var _ WnsKeeper = Keeper{}

type WnsKeeper interface {
	SetMetaData(sdk.Context, string, *types.MetaData) error
	GetMetaData(sdk.Context, string) (*types.MetaData, error)
	IsOwner(sdk.Context, string, string) bool
}

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	memeKey  storetypes.StoreKey

	authKeeper types.AuthKeeper
	bankKeeper types.BankKeeper
}

func (k Keeper) SetMetaData(ctx sdk.Context, name string, metadata *types.MetaData) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(metadata)
	store.Set(types.NameToMetaDataKey(name), bz)
	return nil
}

func (k Keeper) GetMetaData(ctx sdk.Context, name string) (*types.MetaData, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NameToMetaDataKey(name))
	if bz == nil {
		return nil, fmt.Errorf("record not found for name %v", name)
	}

	var metadata *types.MetaData
	k.cdc.MustUnmarshal(bz, metadata)
	return metadata, nil
}

func (k Keeper) IsOwner(ctx sdk.Context, name string, address string) bool {
	metadata, err := k.GetMetaData(ctx, name)
	if err != nil {
		return false
	}

	if metadata.Owner == address {
		return true
	}

	return false
}
