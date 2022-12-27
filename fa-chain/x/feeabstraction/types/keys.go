package types

const (
	// ModuleName defines the module name
	ModuleName = "fachain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_fachain"

	TWAP_STORE_QUERY_WITH_PROOF = "/osmosis.twap.v1beta1.Query/ArithmeticTwapToNow"
)

var (
	KeyPrefixDenom = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetKeyDenomPrefix(denom string) []byte {
	return append(KeyPrefixDenom, []byte(denom)...)
}
