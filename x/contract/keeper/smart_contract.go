package keeper

import (
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// smartContractStore returns the prefix store that contains smart contracts
func smartContractStore(k Keeper, ctx sdk.Context) (contractStore prefix.Store) {
	// Get module KVStore from the MultiStore
	moduleStore := ctx.KVStore(k.storeKey)

	// Get the subset of the module store with prefix SmartContractKeyPrefix
	contractStore = prefix.NewStore(moduleStore, types.KeyPrefix(types.SmartContractKeyPrefix))

	return
}

// GetSmartContract returns the smart contract with the given address from the contract store
func (k Keeper) GetSmartContract(ctx sdk.Context, address sdk.AccAddress) (contract types.SmartContract, found bool) {
	contractStore := smartContractStore(k, ctx)

	var contractBytes []byte = contractStore.Get(types.SmartContractStoreKey(address))

	if contractBytes == nil {
		return contract, false
	}

	k.cdc.MustUnmarshal(contractBytes, &contract)
	return contract, true
}

// SetSmartContract sets the given smart contract at the given address in the contract store
func (k Keeper) SetSmartContract(ctx sdk.Context, address sdk.AccAddress, contract types.SmartContract) {
	contractStore := smartContractStore(k, ctx)

	var contractBytes []byte = k.cdc.MustMarshal(&contract)
	contractStore.Set(types.SmartContractStoreKey(address), contractBytes)
}
