package keeper

import (
	"github.com/arran8901/chainlog-platform/x/contract/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSmartContract set a specific smartContract in the store from its index
func (k Keeper) SetSmartContract(ctx sdk.Context, smartContract types.SmartContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))
	b := k.cdc.MustMarshal(&smartContract)
	store.Set(types.SmartContractKey(
		smartContract.Address,
	), b)
}

// GetSmartContract returns a smartContract from its index
func (k Keeper) GetSmartContract(
	ctx sdk.Context,
	address string,

) (val types.SmartContract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))

	b := store.Get(types.SmartContractKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSmartContract removes a smartContract from the store
func (k Keeper) RemoveSmartContract(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))
	store.Delete(types.SmartContractKey(
		address,
	))
}

// GetAllSmartContract returns all smartContract
func (k Keeper) GetAllSmartContract(ctx sdk.Context) (list []types.SmartContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SmartContractKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SmartContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
