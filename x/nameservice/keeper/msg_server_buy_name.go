package keeper

import (
	"context"

	"nameservice-nel/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyName(
	goCtx context.Context,
	msg *types.MsgBuyName,
) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get a name from the store
	whois, nameFound := k.GetWhois(ctx, msg.Name)

	// set the price at which the name has to be bought if it is not already owned
	minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	// convert price and bid strings to sdk.Coins
	price, _ := sdk.ParseCoinsNormalized(whois.Price)
	bid, _ := sdk.ParseCoinsNormalized(msg.Bid)

	// convert owner and buyer addresses to sdk.AccAddress
	owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	// if name is found in store
	if nameFound {
		// if current price is higher than the bid
		if price.IsAllGT(bid) {
			// throw an error
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		// otherwise, the bid is higher
		// so send tokens from buyer to the owner
		k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
	} else { // name not found in store
		// if minPrice is higher than the bid
		if minPrice.IsAllGT(bid) {
			// throw an error
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		// otherwise, the bid is higher
		// so send tokesn from buyer to the module
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
	}

	// create an updated whois record
	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: whois.Value,
		Price: bid.String(),
		Owner: buyer.String(),
	}

	// write whois information to the store
	k.SetWhois(ctx, newWhois)

	return &types.MsgBuyNameResponse{}, nil
}
