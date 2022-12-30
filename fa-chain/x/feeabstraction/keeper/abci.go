package keeper

import (
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	appparams "github.com/Smartosc-Blockchain/fa-chain/app/params"
	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	gammtypes "github.com/osmosis-labs/osmosis/v13/x/gamm/types"
	twapquery "github.com/osmosis-labs/osmosis/v13/x/twap/client/queryproto"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// due to ibc denom hash, we can only accept denom directly from Osmosis
	// is there a way to get all assets on a channel
	k.transferKeeper.IterateDenomTraces(ctx, func(denomTrace ibctransfertypes.DenomTrace) bool {
		// if an ibc denom exists, skip
		if k.HasDenomTrack(ctx, denomTrace.GetBaseDenom()) {
			return true
		}

		// if found out that denom belongs to osmosis channel_id, register denom trace
		if strings.Contains(denomTrace.GetPath(), osmo_juno_channel_id) {
			k.SetDenomTrack(ctx, denomTrace.GetBaseDenom(), denomTrace.IBCDenom())
		}

		return true
	})

	// get pools information from Osmosis
	// make request for pools
	k.IterateDenomTrack(ctx, func(denomOsmo, _ string) bool {
		req := gammtypes.QueryPoolsWithFilterRequest{
			MinLiquidity: sdk.NewCoins(sdk.NewCoin(denomOsmo, sdk.ZeroInt()), sdk.NewCoin(GetIBCDenom(osmo_juno_channel_id, appparams.DefaultBondDenom).IBCDenom(), sdk.ZeroInt())),
			PoolType:     "Balancer",
		}

		data, err := req.Marshal()
		if err != nil {
			k.Logger(ctx).Error("failed to marshall request", "error", err)
			return true
		}
		ttl := uint64(1000)

		if err := k.icqKeeper.MakeRequest(ctx,
			types.ModuleName,
			ICQCallbackID_Pool,
			host_zone_chain_id,
			juno_osmo_connection_id,
			types.POOL_STORE_QUERY,
			data,
			ttl,
		); err != nil {
			k.Logger(ctx).Error("failed to make request", "error", err)
			return true
		}

		return true
	})
}

// EndBlocker of feeabstraction module
func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// update fee
	k.IteratePool(ctx, func(denomOsmo string, poolId uint64) bool {
		// make request for twap
		req := twapquery.ArithmeticTwapToNowRequest{
			PoolId:    poolId,
			BaseAsset: denomOsmo,
		}
		data, err := req.Marshal()
		if err != nil {
			k.Logger(ctx).Error("failed to marshall request", "error", err)
			return true
		}
		ttl := uint64(1000)

		if err := k.icqKeeper.MakeRequest(ctx,
			types.ModuleName,
			ICQCallbackID_FeeRate,
			host_zone_chain_id,
			juno_osmo_connection_id,
			types.TWAP_STORE_QUERY,
			data,
			ttl,
		); err != nil {
			k.Logger(ctx).Error("failed to make request", "error", err)
			return true
		}

		return true
	})
}
