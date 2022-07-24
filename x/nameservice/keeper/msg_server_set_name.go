package keeper

import (
	"context"

	"nameservice-nel/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetName(
	goCtx context.Context,
	msg *types.MsgSetName,
) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get name information from the store
	whois, nameFound := k.GetWhois(ctx, msg.Name)
	if !nameFound {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Name doesn't exist")
	}

	// if the message sender does not match the name owner, can't set name
	if msg.Creator != whois.Owner {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	// otherwise, create an updated whois record
	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: msg.Value,
		Owner: whois.Owner,
		Price: whois.Price,
	}

	// write whois information to the store
	k.SetWhois(ctx, newWhois)

	return &types.MsgSetNameResponse{}, nil
}
