package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	twapquery "github.com/osmosis-labs/osmosis/v13/x/twap/client/queryproto"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// get pools information from Osmosis
	// make request for pools
	connections := k.icqKeeper.IBCKeeper.ConnectionKeeper.GetAllConnections(ctx)
	fmt.Printf("connections = %v\n", connections)
	//connections[0].Counterparty.GetPrefix()
}

// EndBlocker of feeabstraction module
func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// get pools information from Osmosis
	// make request for pools

	// make request for twap
	req := twapquery.ArithmeticTwapToNowRequest{
		PoolId:    1,
		BaseAsset: "uosmo",
	}
	data, err := req.Marshal()
	if err != nil {
		return
	}
	ttl := uint64(1000)

	if err := k.icqKeeper.MakeRequest(ctx,
		types.ModuleName,
		ICQCallbackID_FeeRate,
		host_zone_chain_id,
		host_zone_connection_id,
		types.TWAP_STORE_QUERY_WITH_PROOF,
		data,
		ttl,
	); err != nil {
		return
	}

	if err != nil {
		k.Logger(ctx).Error("failed to make request", "error", err)
	}
}
