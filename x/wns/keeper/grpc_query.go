package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func (q Querier) Metadata(goCtx context.Context, req *types.QueryMetaDataRequest) (*types.QueryMetaDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	metadata, err := q.GetMetaData(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &types.QueryMetaDataResponse{Metadata: metadata}, nil
}
