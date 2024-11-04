package flags

import (
	"github.com/ethereum-optimism/optimism/op-proposer/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/urfave/cli/v2"
)

func prefixEnvVar(name string) []string {
	return opservice.PrefixEnvVar(flags.EnvVarPrefix, name)
}

var (
	L2EthRpcFlag = &cli.StringFlag{
		Name:     "l2-eth-rpc",
		Usage:    "HTTP provider URL for L2",
		EnvVars:  prefixEnvVar("L2_ETH_RPC"),
		Required: true,
	}
	L2RethFlag = &cli.BoolFlag{
		Name:     "l2-reth",
		Usage:    "Is the L2 HTTP provider running Reth?",
		EnvVars:  prefixEnvVar("L2_RETH"),
		Value:    false,
		Required: false,
	}
	EnclaveRpcFlag = &cli.StringFlag{
		Name:     "enclave-rpc",
		Usage:    "HTTP provider URL for the enclave service",
		EnvVars:  prefixEnvVar("ENCLAVE_RPC"),
		Required: true,
	}
	MinProposalIntervalFlag = &cli.Uint64Flag{
		Name:    "min-proposal-interval",
		Usage:   "Minimum time between proposals (in L2 blocks)",
		EnvVars: prefixEnvVar("MIN_PROPOSAL_INTERVAL"),
		Value:   600,
	}
)

var requiredFlags = []cli.Flag{
	L2EthRpcFlag,
	L2RethFlag,
	EnclaveRpcFlag,
	MinProposalIntervalFlag,
}

func init() {
	Flags = append(requiredFlags, flags.Flags...)
}

var Flags []cli.Flag
