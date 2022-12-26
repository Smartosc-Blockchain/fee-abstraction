package fachain_test

import (
	"testing"

	keepertest "github.com/Smartosc-Blockchain/fa-chain/testutil/keeper"
	"github.com/Smartosc-Blockchain/fa-chain/testutil/nullify"
	fachain "github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction"
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FachainKeeper(t)
	fachain.InitGenesis(ctx, *k, genesisState)
	got := fachain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
