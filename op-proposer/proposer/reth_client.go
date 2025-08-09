package proposer

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/sources/caching"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/ethclient"
)

type rethClient struct {
	ethClient
}

func NewRethClient(client *ethclient.Client, metrics caching.Metrics) Client {
	ethClient := newClient(client, metrics)
	return &rethClient{
		ethClient: *ethClient,
	}
}

func (e *rethClient) ExecutionWitness(ctx context.Context, hash common.Hash) (*stateless.ExecutionWitness, error) {
	var witness stateless.ExecutionWitness
	if err := e.client.Client().CallContext(ctx, &witness, "debug_executionWitness", hash.Hex()); err != nil {
		return nil, err
	}

	// reth doesn't return required headers (for BLOCKHASH), so eagerly populate them all:
	header, err := e.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	parentHash := header.ParentHash
	history := int(min(256, header.Number.Uint64()))
	for i := 0; i < history; i++ {
		parent, err := e.HeaderByHash(ctx, parentHash)
		if err != nil {
			return nil, err
		}
		witness.Headers = append(witness.Headers, parent)
		parentHash = parent.ParentHash
	}

	return &witness, nil
}
