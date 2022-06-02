package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "contract"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_contract"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Keys for smart contract mapping
const (
	// SmartContractKeyPrefix is the prefix for a stored smart contract
	SmartContractKeyPrefix = "SmartContract/"
)

// SmartContractStoreKey returns the store key for a given smart contract address
func SmartContractStoreKey(address sdk.AccAddress) []byte {
	return append([]byte(SmartContractKeyPrefix), address.Bytes()...)
}
