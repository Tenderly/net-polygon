package api

import (
	"context"

	"github.com/tenderly/net-polygon/common/hexutil"
	"github.com/tenderly/net-polygon/core/state"
	"github.com/tenderly/net-polygon/internal/ethapi"
	"github.com/tenderly/net-polygon/rpc"
)

//go:generate mockgen -destination=./caller_mock.go -package=api . Caller
type Caller interface {
	Call(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, blockOverrides *ethapi.BlockOverrides) (hexutil.Bytes, error)
	CallWithState(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, state *state.StateDB, overrides *ethapi.StateOverride, blockOverrides *ethapi.BlockOverrides) (hexutil.Bytes, error)
}
