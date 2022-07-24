package keeper

import (
	"nameservice-nel/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
