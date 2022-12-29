package keeper

import (
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// params changeable through gov
const (
	host_zone_chain_id      = "test-osmo"
	juno_osmo_connection_id = "connection-0"
	osmo_juno_connection_id = "connection-0"
	juno_osmo_channel_id    = "channel-0"
	osmo_juno_channel_id    = "channel-0"
)

var (
	// one token on Juno is mapped to its equivalence on Osmosis
	// should have been stored in KVStore
	token_pairs = map[string]string{
		"ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518": "uosmo",
	}
)

func GetJunoIBCDenom() transfertypes.DenomTrace {
	sourcePrefix := transfertypes.GetDenomPrefix("transfer", juno_osmo_channel_id)
	prefixedDenom := sourcePrefix + "ujuno"

	return transfertypes.ParseDenomTrace(prefixedDenom)
}
