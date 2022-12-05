package types

const (
	// ModuleName defines the module name
	ModuleName = "wns"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wns"

	// this line is used by starport scaffolding # ibc/keys/name
)

// prefix bytes for the cronos persistent store
const (
	prefixNodeToRecord = iota + 1
	prefixAddressToName
	paramsKey
)

// KVStore key prefixes
var (
	KeyPrefixNodeToRecord  = []byte{prefixNodeToRecord}
	KeyPrefixAddressToName = []byte{prefixAddressToName}

	// ParamsKey is the key for params.
	ParamsKey = []byte{paramsKey}
)

// this line is used by starport scaffolding # ibc/keys/port

// NodeToRecordKey defines the store key for node to record mapping
func NodeToRecordKey(node Node) []byte {
	return append(KeyPrefixNodeToRecord, node[:]...)
}

// AddressToNameKey defines the store key for account address to domain name mapping
func AddressToNameKey(address string) []byte {
	return append(KeyPrefixAddressToName, address...)
}
