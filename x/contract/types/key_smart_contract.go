package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SmartContractKeyPrefix is the prefix to retrieve all SmartContract
	SmartContractKeyPrefix = "SmartContract/value/"
)

// SmartContractKey returns the store key to retrieve a SmartContract from the index fields
func SmartContractKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
