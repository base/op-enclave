package proposer

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-service/sources/caching"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/ethclient"
)

// see https://github.com/alloy-rs/alloy/blob/main/crates/rpc-types-debug/src/debug.rs
type RethExecutionWitness struct {
	State map[common.Hash]hexutil.Bytes
	Codes map[common.Hash]hexutil.Bytes
	Keys  map[common.Hash]hexutil.Bytes
}

type rethClient struct {
	ethClient
}

func NewRethClient(client *ethclient.Client, metrics caching.Metrics) Client {
	ethClient := newClient(client, metrics)
	return &rethClient{
		ethClient: *ethClient,
	}
}

func (e *rethClient) ExecutionWitness(ctx context.Context, hash common.Hash) ([]byte, error) {
	header, err := e.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	var witness RethExecutionWitness
	// TODO it would be nice if reth accepted a tx hash parameter here
	if err := e.client.Client().CallContext(ctx, &witness, "debug_executionWitness", "0x"+header.Number.Text(16)); err != nil {
		return nil, err
	}

	w := &stateless.Witness{
		Codes: make(map[string]struct{}),
		State: make(map[string]struct{}),
	}
	for _, code := range witness.Codes {
		w.Codes[string(code)] = struct{}{}
	}
	for _, state := range witness.State {
		w.State[string(state)] = struct{}{}
	}

	// reth doesn't return required headers (for BLOCKHASH), so eagerly populate them all:
	parentHash := header.ParentHash
	history := int(min(256, header.Number.Uint64()))
	for i := 0; i < history; i++ {
		parent, err := e.HeaderByHash(ctx, parentHash)
		if err != nil {
			return nil, err
		}
		w.Headers = append(w.Headers, parent)
		parentHash = parent.ParentHash
	}

	var buf bytes.Buffer
	if err = w.EncodeRLP(&buf); err != nil {
		return nil, fmt.Errorf("failed to encode witness: %w", err)
	}
	return buf.Bytes(), nil
}
