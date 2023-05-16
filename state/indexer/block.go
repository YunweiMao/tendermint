package indexer

import (
	"context"

	//Query is in ./libs/pubsub/query/query.go
	"github.com/YunweiMao/tendermint/libs/pubsub/query"
	//EventDataNewBlockHeader is in ./types/events.go
	"github.com/YunweiMao/tendermint/types"
)

//go:generate ../../scripts/mockery_generate.sh BlockIndexer

// BlockIndexer defines an interface contract for indexing block events.
type BlockIndexer interface {
	// Has returns true if the given height has been indexed. An error is returned
	// upon database query failure.
	Has(height int64) (bool, error)

	// Index indexes BeginBlock and EndBlock events for a given block by its height.
	Index(types.EventDataNewBlockHeader) error

	// Search performs a query for block heights that match a given BeginBlock
	// and Endblock event search criteria.
	Search(ctx context.Context, q *query.Query) ([]int64, error)
}
