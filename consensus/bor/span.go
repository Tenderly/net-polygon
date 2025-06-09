package bor

import (
	"context"

	"github.com/tenderly/net-polygon/common"
	"github.com/tenderly/net-polygon/consensus/bor/heimdall/span"
	"github.com/tenderly/net-polygon/consensus/bor/valset"
	"github.com/tenderly/net-polygon/core"
	"github.com/tenderly/net-polygon/core/state"
	"github.com/tenderly/net-polygon/core/types"
	"github.com/tenderly/net-polygon/rpc"
)

//go:generate mockgen -destination=./span_mock.go -package=bor . Spanner
type Spanner interface {
	GetCurrentSpan(ctx context.Context, headerHash common.Hash) (*span.Span, error)
	GetCurrentValidatorsByHash(ctx context.Context, headerHash common.Hash, blockNumber uint64) ([]*valset.Validator, error)
	GetCurrentValidatorsByBlockNrOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, blockNumber uint64) ([]*valset.Validator, error)
	CommitSpan(ctx context.Context, heimdallSpan span.HeimdallSpan, state *state.StateDB, header *types.Header, chainContext core.ChainContext) error
}
