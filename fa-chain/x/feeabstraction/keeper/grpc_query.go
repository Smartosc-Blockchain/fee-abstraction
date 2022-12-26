package keeper

import (
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
)

var _ types.QueryServer = Keeper{}
