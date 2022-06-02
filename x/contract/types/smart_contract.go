package types

import (
	"crypto/sha256"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CodeSizeLimit = 65536
)

func ValidateCodeBasic(code string) error {
	// Check contract does not exceed size limit
	if len(code) > CodeSizeLimit {
		return ErrContractCodeTooLarge
	}

	return nil
}

func GenerateContractAddress(creator sdk.AccAddress, creatorSequence uint64) sdk.AccAddress {
	// Convert creatorSequence to []byte
	creatorSequenceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(creatorSequenceBytes, creatorSequence)

	// Concatenate creator address' raw bytes and creatorSequenceBytes
	// Then hash the result to give the contract address' raw bytes
	var contractAddressByteArr [sha256.Size]byte = sha256.Sum256(append(creator.Bytes(), creatorSequenceBytes...))
	var contractAddressBytes []byte = contractAddressByteArr[:]

	// Address should be valid
	if err := sdk.VerifyAddressFormat(contractAddressBytes); err != nil {
		panic(err)
	}
	// Convert raw address bytes to sdk.AccAddress and return
	return sdk.AccAddress(contractAddressBytes)
}
