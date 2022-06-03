package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/contract module sentinel errors
var (
	ErrContractCodeTooLarge     = sdkerrors.Register(ModuleName, 1100, fmt.Sprintf("smart contract code exceeds size limit of %dB", CodeSizeLimit))
	ErrContractNotFound         = sdkerrors.Register(ModuleName, 1101, "no contract found")
	ErrContractMessageException = sdkerrors.Register(ModuleName, 1102, "exception during message call")
)
