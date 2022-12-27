package keeper

import (
	"context"

	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) FeeRate(goCtx context.Context, req *types.QueryFeeRateRequest) (*types.QueryFeeRateResponse, error) {
	// get newest fee rate
	feeRate := sdk.Coin{}

	return &types.QueryFeeRateResponse{FeeRate: feeRate}, nil
}
