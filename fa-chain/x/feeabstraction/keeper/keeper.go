package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/Smartosc-Blockchain/fa-chain/x/feeabstraction/types"
	icqkeeper "github.com/Smartosc-Blockchain/fa-chain/x/interchainquery/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		icqKeeper  icqkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	icqKeeper icqkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		icqKeeper:  icqKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// fee of non - native compared to ujuno
func (k Keeper) SetFeeRate(ctx sdk.Context, denom string, fee_rate sdk.Dec) error {
	store := ctx.KVStore(k.storeKey)
	data, err := fee_rate.Marshal()
	if err != nil {
		return err
	}
	store.Set(types.GetKeyDenomPrefix(denom), data)

	return nil
}
