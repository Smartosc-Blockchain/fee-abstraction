package keeper_test

import (
	"testing"

	testkeeper "github.com/Smartosc-Blockchain/fa-chain/testutil/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FachainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
