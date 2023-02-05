package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genesis *types.GenesisState) {
	if genesis == nil {
		panic("nil genesis pointer")
	}
	for _, name := range genesis.Names {
		if !name.Valid() {
			panic("not valid genesis domain names")
		}

		err := k.SetMetaData(ctx, name.Name, name.Metadata)
		if err != nil {
			panic(err)
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var names []*types.DomainName
	k.IterateDomainNames(ctx, func(name *types.DomainName) bool {
		names = append(names, name)
		return false
	})

	return &types.GenesisState{Names: names}
}
