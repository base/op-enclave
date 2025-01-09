package enclave

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

func ExecuteStateless(
	ctx context.Context,
	config *params.ChainConfig,
	rollupConfig *rollup.Config,
	l1Origin *types.Header,
	l1Receipts types.Receipts,
	previousBlockTxs []hexutil.Bytes,
	blockHeader *types.Header,
	sequencedTxs []hexutil.Bytes,
	witness *stateless.Witness,
	messageAccount *eth.AccountResult,
) error {
	l1OriginHash := l1Origin.Hash()
	computed := types.DeriveSha(l1Receipts, trie.NewStackTrie(nil))
	if computed != l1Origin.ReceiptHash {
		return errors.New("invalid receipts")
	}

	previousBlockHeader := witness.Headers[0]
	previousBlockHash := previousBlockHeader.Hash()
	if blockHeader.ParentHash != previousBlockHash {
		return errors.New("invalid parent hash")
	}

	// block must only contain deposit transactions if it is outside the sequencer drift
	if len(sequencedTxs) > 0 && blockHeader.Time > l1Origin.Time+maxSequencerDriftFjord {
		return errors.New("l1 origin is too old")
	}

	unmarshalTxs := func(rlp []hexutil.Bytes) (types.Transactions, error) {
		txs := make(types.Transactions, len(rlp))
		for i, tx := range rlp {
			txs[i] = new(types.Transaction)
			if err := txs[i].UnmarshalBinary(tx); err != nil {
				return nil, fmt.Errorf("failed to unmarshal transaction: %w", err)
			}
		}
		return txs, nil
	}
	previousTxs, err := unmarshalTxs(previousBlockTxs)
	if err != nil {
		return err
	}

	previousTxHash := types.DeriveSha(previousTxs, trie.NewStackTrie(nil))
	if previousTxHash != previousBlockHeader.TxHash {
		return errors.New("invalid tx hash")
	}

	previousBlock := types.NewBlockWithHeader(previousBlockHeader).WithBody(types.Body{
		Transactions: previousTxs,
	})

	l2Parent, err := derive.L2BlockToBlockRef(rollupConfig, previousBlock)
	if err != nil {
		return fmt.Errorf("failed to convert L2 block to block ref: %w", err)
	}

	if l2Parent.L1Origin.Hash != l1OriginHash && l2Parent.L1Origin.Hash != l1Origin.ParentHash {
		return errors.New("invalid L1 origin")
	}

	l1Fetcher := NewL1ReceiptsFetcher(l1OriginHash, l1Origin, l1Receipts)
	l2Fetcher := NewL2SystemConfigFetcher(rollupConfig, previousBlockHash, previousBlockHeader, previousTxs)
	attributeBuilder := derive.NewFetchingAttributesBuilder(rollupConfig, l1Fetcher, l2Fetcher)
	payload, err := attributeBuilder.PreparePayloadAttributes(ctx, l2Parent, eth.BlockID{
		Hash:   l1OriginHash,
		Number: l1Origin.Number.Uint64(),
	})
	if err != nil {
		return fmt.Errorf("failed to prepare payload attributes: %w", err)
	}

	// sequencer cannot include manual deposit transactions; otherwise it could mint funds arbitrarily
	txs, err := unmarshalTxs(sequencedTxs)
	if err != nil {
		return err
	}
	for _, tx := range txs {
		if tx.IsDepositTx() {
			return errors.New("sequenced txs cannot include deposits")
		}
	}

	// now add the deposits from L1 (and any from fork upgrades)
	payloadTxs, err := unmarshalTxs(payload.Transactions)
	if err != nil {
		return fmt.Errorf("failed to parse payload transactions: %w", err)
	}
	txs = append(payloadTxs, txs...)

	expectedRoot := blockHeader.Root
	expectedReceiptHash := blockHeader.ReceiptHash
	blockHeader = types.CopyHeader(blockHeader)
	blockHeader.Root = common.Hash{}
	blockHeader.ReceiptHash = common.Hash{}
	block := types.NewBlockWithHeader(blockHeader).WithBody(types.Body{
		Transactions: txs,
	})
	blockHeader.Root, blockHeader.ReceiptHash, err = core.ExecuteStateless(config, block, witness)
	if err != nil {
		return fmt.Errorf("failed to execute stateless: %w", err)
	}
	if blockHeader.Root != expectedRoot {
		return errors.New("invalid state root")
	}
	if blockHeader.ReceiptHash != expectedReceiptHash {
		return errors.New("invalid receipt hash")
	}

	if messageAccount.Address.Cmp(l2ToL1MessagePasserAddress) != 0 {
		return errors.New("invalid message account address")
	}
	if err = messageAccount.Verify(blockHeader.Root); err != nil {
		return fmt.Errorf("failed to verify message account: %w", err)
	}

	return nil
}
