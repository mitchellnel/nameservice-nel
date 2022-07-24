package keeper

import (
	"context"

	"nameservice-nel/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteName(
	goCtx context.Context,
	msg *types.MsgDeleteName,
) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// try getting name information from the store
	whois, nameFound := k.GetWhois(ctx, msg.Name)
	if !nameFound {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Name doesn't exist")
	}

	// if the message sender does not match the name owner, can't delete name
	if msg.Creator != whois.Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	// otherwise, remove the name information from the store
	k.RemoveWhois(ctx, msg.Name)

	return &types.MsgDeleteNameResponse{}, nil
}
