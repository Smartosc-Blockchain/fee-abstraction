package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	appparams "github.com/Smartosc-Blockchain/fa-chain/app/params"
)

func init() {
	appparams.SetAddressPrefixes()
}

// Initializes a new StrideApp without IBC functionality
func InitTestApp(initChain bool) *App {
	db := dbm.NewMemDB()
	codec := MakeEncodingConfig()
	app := New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		5,
		codec,
		simapp.EmptyAppOptions{},
	)
	if initChain {
		genesisState := NewDefaultGenesisState(codec.Marshaler)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

// Initializes a new Stride App casted as a TestingApp for IBC support
func InitStrideIBCTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	app := InitTestApp(false)
	return ibctesting.TestingApp(app), NewDefaultGenesisState(app.appCodec)
}

// initializes a new Osmosis App casted as a TestingApp for IBC support
