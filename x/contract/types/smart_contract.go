package types

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
