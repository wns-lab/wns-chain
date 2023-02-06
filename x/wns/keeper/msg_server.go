package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wns-lab/wns-chain/x/wns/types"
)

var _ types.MsgServer = MsgServer{}

type MsgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) MsgServer {
	return MsgServer{Keeper: keeper}
}

func (ms MsgServer) SetMetadata(goCtx context.Context, msg *types.MsgSetMetaData) (*types.SetMetaDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !ms.IsOwner(ctx, msg.Name, msg.Sender) {
		return nil, fmt.Errorf("sender %v is not the name owner", msg.Sender)
	}

	err := ms.Keeper.SetMetaData(ctx, msg.Name, msg.Metadata)
	if err != nil {
		return nil, err
	}

	return &types.SetMetaDataResponse{}, nil
}

func (ms MsgServer) SetOwner(goCtx context.Context, msg *types.MsgSetOwner) (*types.SetOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	re, err := ms.GetMetaData(ctx, msg.Name)
	if err != nil {
		return nil, err
	}

	if msg.Sender != re.Owner {
		return nil, fmt.Errorf("sender %v is not the name owner", msg.Sender)
	}

	metadata := &types.MetaData{
		Owner:    msg.Owner,
		Resolver: re.Resolver,
		TTL:      re.TTL,
	}
	err = ms.Keeper.SetMetaData(ctx, msg.Name, metadata)
	if err != nil {
		return nil, err
	}

	return &types.SetOwnerResponse{}, nil
}

func (ms MsgServer) SetResolver(goCtx context.Context, msg *types.MsgSetResolver) (*types.SetResolverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	re, err := ms.GetMetaData(ctx, msg.Name)
	if err != nil {
		return nil, err
	}

	if msg.Sender != re.Owner {
		return nil, fmt.Errorf("sender %v is not the name owner", msg.Sender)
	}

	metadata := &types.MetaData{
		Owner:    re.Owner,
		Resolver: msg.Resolver,
		TTL:      re.TTL,
	}
	err = ms.Keeper.SetMetaData(ctx, msg.Name, metadata)
	if err != nil {
		return nil, err
	}

	return &types.SetResolverResponse{}, nil
}

func (ms MsgServer) SetTTL(goCtx context.Context, msg *types.MsgSetTTL) (*types.SetTTLResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	re, err := ms.GetMetaData(ctx, msg.Name)
	if err != nil {
		return nil, err
	}

	if msg.Sender != re.Owner {
		return nil, fmt.Errorf("sender %v is not the name owner", msg.Sender)
	}

	metadata := &types.MetaData{
		Owner:    re.Owner,
		Resolver: re.Resolver,
		TTL:      msg.TTL,
	}
	err = ms.Keeper.SetMetaData(ctx, msg.Name, metadata)
	if err != nil {
		return nil, err
	}

	return &types.SetTTLResponse{}, nil
}
