package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Smartosc-Blockchain/fa-chain/testutil/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FachainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
