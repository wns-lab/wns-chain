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
	prefixNameToMetaData = iota + 1
	prefixAddressToName
	paramsKey
)

// KVStore key prefixes
var (
	KeyPrefixNameToMetaData = []byte{prefixNameToMetaData}
	KeyPrefixAddressToName  = []byte{prefixAddressToName}

	// ParamsKey is the key for params.
	ParamsKey = []byte{paramsKey}
)

// this line is used by starport scaffolding # ibc/keys/port

// NameToMetaDataKey defines the store key for name to record mapping
func NameToMetaDataKey(name string) []byte {
	return append(KeyPrefixNameToMetaData, []byte(name)...)
}

// AddressToNameKey defines the store key for account address to domain name mapping
func AddressToNameKey(address string) []byte {
	return append(KeyPrefixAddressToName, address...)
}
