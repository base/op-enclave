package withdrawals

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/base/op-enclave/bindings"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type ProofClient interface {
	GetProof(context.Context, common.Address, []string, *big.Int) (*gethclient.AccountResult, error)
}

type EthClient interface {
	TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error)
	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
}

type OutputOracle interface {
	LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error)
	GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error)
	LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error)
}

type Portal interface {
	ProveAndFinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx bindings.TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof bindings.TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error)
}

func WaitForOutputBlock(ctx context.Context, outputOracle *bindings.OutputOracle, blockNumber *big.Int, pollInterval time.Duration) (*big.Int, error) {
	for {
		l2OutputBlock, err := outputOracle.LatestBlockNumber(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		if l2OutputBlock.Cmp(blockNumber) >= 0 {
			return l2OutputBlock, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(pollInterval):
		}
	}
}

func ProveAndFinalizeWithdrawals(ctx context.Context, l2ProofCl ProofClient, l2Client EthClient, opts *bind.TransactOpts, outputOracle OutputOracle, portal Portal, withdrawalTxHash common.Hash, l2OutputBlock *big.Int) ([]*types.Transaction, error) {
	l2OutputIndex, err := outputOracle.LatestOutputIndex(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("error getting output index: %w", err)
	}

	receipt, err := l2Client.TransactionReceipt(ctx, withdrawalTxHash)
	if err != nil {
		return nil, fmt.Errorf("error getting withdrawal transaction receipt: %w", err)
	}
	evs, err := withdrawals.ParseMessagesPassed(receipt)
	if err != nil {
		return nil, fmt.Errorf("error parsing withdrawal logs: %w", err)
	}

	header, err := l2Client.HeaderByNumber(ctx, l2OutputBlock)
	if err != nil {
		return nil, fmt.Errorf("error requesting block header: %w", err)
	}

	txs := make([]*types.Transaction, len(evs))
	for i, ev := range evs {
		withdrawal, err := withdrawals.ProveWithdrawalParametersForEvent(ctx, l2ProofCl, ev, header, l2OutputIndex)
		if err != nil {
			return nil, fmt.Errorf("error generating withdrawal proof: %w", err)
		}

		outputRootProof := bindings.TypesOutputRootProof{
			Version:                  withdrawal.OutputRootProof.Version,
			StateRoot:                withdrawal.OutputRootProof.StateRoot,
			MessagePasserStorageRoot: withdrawal.OutputRootProof.MessagePasserStorageRoot,
			LatestBlockhash:          withdrawal.OutputRootProof.LatestBlockhash,
		}

		txs[i], err = portal.ProveAndFinalizeWithdrawalTransaction(
			opts,
			bindings.TypesWithdrawalTransaction{
				Nonce:    withdrawal.Nonce,
				Sender:   withdrawal.Sender,
				Target:   withdrawal.Target,
				Value:    withdrawal.Value,
				GasLimit: withdrawal.GasLimit,
				Data:     withdrawal.Data,
			},
			withdrawal.L2OutputIndex,
			outputRootProof,
			withdrawal.WithdrawalProof,
		)
		if err != nil {
			return nil, fmt.Errorf("error submitting withdrawal tx: %w", err)
		}
	}

	return txs, nil
}

func WaitForReceipt(ctx context.Context, client EthClient, txHash common.Hash, pollInterval time.Duration) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(pollInterval):
			}
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("unsuccessful receipt status")
		}
		return receipt, nil
	}
}
