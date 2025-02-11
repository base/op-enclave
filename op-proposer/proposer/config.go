package proposer

import (
	"github.com/base/op-enclave/op-proposer/flags"
	"github.com/ethereum-optimism/optimism/op-proposer/proposer"
	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	*proposer.CLIConfig
	L2EthRpc            string
	L2Reth              bool
	EnclaveRpc          string
	MinProposalInterval uint64
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		CLIConfig:           proposer.NewConfig(ctx),
		L2EthRpc:            ctx.String(flags.L2EthRpcFlag.Name),
		L2Reth:              ctx.Bool(flags.L2RethFlag.Name),
		EnclaveRpc:          ctx.String(flags.EnclaveRpcFlag.Name),
		MinProposalInterval: ctx.Uint64(flags.MinProposalIntervalFlag.Name),
	}
}
