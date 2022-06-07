package keeper

import (
	"context"

	chainlog "github.com/arran8901/chainlog-lang"
	"github.com/arran8901/chainlog-platform/x/contract/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryContract(goCtx context.Context, req *types.QueryQueryContractRequest) (*types.QueryQueryContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse address
	address, err := sdk.AccAddressFromBech32(req.ContractAddress)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address: %s (%s)", req.ContractAddress, err.Error())
	}

	// Attempt to get contract from store
	// Return an error if not found
	contract, found := k.GetSmartContract(ctx, address)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrContractNotFound, "no contract found at address %s", req.ContractAddress)
	}

	// Get contract account's balance
	contractBalance := k.bankKeeper.GetBalance(ctx, address, types.SmartContractCoinDenom)

	// Initialise new Chainlog interpreter and load code and dynamic KB
	i := chainlog.NewInterpreter()
	i.ConsultWithDynamicKB(contract.Code, chainlog.ParseDynamicKBLogicProgram(contract.DynamicKb))

	// Prepare response, initially unsuccessful and with empty derivations
	response := &types.QueryQueryContractResponse{
		Successful:  false,
		Derivations: make([]*types.QueryContractDerivation, 0),
	}

	// Construct query context
	queryCtx := chainlog.QueryContext{
		Time:    ctx.BlockTime().Unix(),
		Balance: contractBalance.Amount.Uint64(),
	}

	// Submit query to interpreter
	// Continue searching for derivations until reaching the given nDerivations or
	// until there are no more solutions to the query
	// If an exception occurs at any point, return an error
	itr, err := i.QueryWithContext(req.Query, &queryCtx)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrContractQueryException, err.Error())
	}
	for i := uint64(0); i < req.NDerivations; i++ {
		derivation, err := itr.Next()
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrContractQueryException, err.Error())
		}
		if !derivation.Successful {
			break
		}
		response.Successful = true
		response.Derivations = append(response.Derivations, &types.QueryContractDerivation{Unifications: derivation.Unifications})
	}

	return response, nil
}
