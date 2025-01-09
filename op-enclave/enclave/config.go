package enclave

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

const (
	version0 uint64 = 0
)

var (
	l2GenesisBlockBaseFeePerGas = hexutil.Big(*(big.NewInt(1000000000)))
	vaultMinWithdrawalAmount    = mustHexBigFromHex("0x8ac7230489e80000")
)

var chainConfigTemplate params.ChainConfig
var rollupConfigTemplate rollup.Config

func init() {
	deployConfig := DefaultDeployConfig()

	var err error
	chainConfigTemplate, rollupConfigTemplate, err = newChainConfigTemplate(&deployConfig)
	if err != nil {
		panic(err)
	}
}

type ChainConfig struct {
	*params.ChainConfig
	*PerChainConfig
}

func newChainConfigTemplate(cfg *genesis.DeployConfig) (params.ChainConfig, rollup.Config, error) {
	l1StartHeader := &types.Header{
		Time:   1,
		Number: big.NewInt(0),
	}
	g, err := genesis.NewL2Genesis(cfg, l1StartHeader)
	if err != nil {
		return params.ChainConfig{}, rollup.Config{}, err
	}

	cfg.OptimismPortalProxy = common.Address{1}
	cfg.SystemConfigProxy = common.Address{1}
	rollupConfig, err := cfg.RollupConfig(l1StartHeader, common.Hash{}, 0)
	if err != nil {
		return params.ChainConfig{}, rollup.Config{}, err
	}

	return *g.Config, *rollupConfig, nil
}

func NewChainConfig(cfg *PerChainConfig) *ChainConfig {
	cfg.ForceDefaults()
	chainConfig := chainConfigTemplate
	chainConfig.ChainID = cfg.ChainID
	return &ChainConfig{
		ChainConfig:    &chainConfig,
		PerChainConfig: cfg,
	}
}

type PerChainConfig struct {
	ChainID *big.Int `json:"chain_id"`

	Genesis   rollup.Genesis `json:"genesis"`
	BlockTime uint64         `json:"block_time"`

	DepositContractAddress common.Address `json:"deposit_contract_address"`
	L1SystemConfigAddress  common.Address `json:"l1_system_config_address"`
}

func FromRollupConfig(cfg *rollup.Config) *PerChainConfig {
	p := &PerChainConfig{
		ChainID:                cfg.L2ChainID,
		Genesis:                cfg.Genesis,
		BlockTime:              cfg.BlockTime,
		DepositContractAddress: cfg.DepositContractAddress,
		L1SystemConfigAddress:  cfg.L1SystemConfigAddress,
	}
	p.ForceDefaults()
	return p
}

func (p *PerChainConfig) ToRollupConfig() *rollup.Config {
	cfg := rollupConfigTemplate
	cfg.L2ChainID = p.ChainID
	cfg.Genesis = p.Genesis
	cfg.BlockTime = p.BlockTime
	cfg.DepositContractAddress = p.DepositContractAddress
	cfg.L1SystemConfigAddress = p.L1SystemConfigAddress
	return &cfg
}

func (p *PerChainConfig) ForceDefaults() {
	p.BlockTime = 1
	p.Genesis.L2.Number = 0
	p.Genesis.SystemConfig.Overhead = eth.Bytes32{}
}

func (p *PerChainConfig) Hash() common.Hash {
	return crypto.Keccak256Hash(p.MarshalBinary())
}

func (p *PerChainConfig) MarshalBinary() (data []byte) {
	data = binary.BigEndian.AppendUint64(data, version0)
	chainIDBytes := p.ChainID.Bytes()
	data = append(data, make([]byte, 32-len(chainIDBytes))...)
	data = append(data, chainIDBytes...)
	data = append(data, p.Genesis.L1.Hash[:]...)
	data = append(data, p.Genesis.L2.Hash[:]...)
	data = binary.BigEndian.AppendUint64(data, p.Genesis.L2Time)
	data = append(data, p.Genesis.SystemConfig.BatcherAddr.Bytes()...)
	data = append(data, p.Genesis.SystemConfig.Scalar[:]...)
	data = binary.BigEndian.AppendUint64(data, p.Genesis.SystemConfig.GasLimit)
	data = append(data, p.DepositContractAddress.Bytes()...)
	data = append(data, p.L1SystemConfigAddress.Bytes()...)
	return data
}

func DefaultDeployConfig() genesis.DeployConfig {
	return genesis.DeployConfig{
		L2InitializationConfig: genesis.L2InitializationConfig{
			L2GenesisBlockDeployConfig: genesis.L2GenesisBlockDeployConfig{
				L2GenesisBlockGasLimit:      30_000_000,
				L2GenesisBlockBaseFeePerGas: &l2GenesisBlockBaseFeePerGas,
			},
			L2VaultsDeployConfig: genesis.L2VaultsDeployConfig{
				BaseFeeVaultWithdrawalNetwork:            "local",
				L1FeeVaultWithdrawalNetwork:              "local",
				SequencerFeeVaultWithdrawalNetwork:       "local",
				SequencerFeeVaultMinimumWithdrawalAmount: vaultMinWithdrawalAmount,
				BaseFeeVaultMinimumWithdrawalAmount:      vaultMinWithdrawalAmount,
				L1FeeVaultMinimumWithdrawalAmount:        vaultMinWithdrawalAmount,
			},
			GovernanceDeployConfig: genesis.GovernanceDeployConfig{
				EnableGovernance:      true,
				GovernanceTokenSymbol: "OP",
				GovernanceTokenName:   "Optimism",
			},
			GasPriceOracleDeployConfig: genesis.GasPriceOracleDeployConfig{
				GasPriceOracleBaseFeeScalar:     1368,
				GasPriceOracleBlobBaseFeeScalar: 810949,
			},
			EIP1559DeployConfig: genesis.EIP1559DeployConfig{
				EIP1559Denominator:       50,
				EIP1559DenominatorCanyon: 250,
				EIP1559Elasticity:        6,
			},
			UpgradeScheduleDeployConfig: genesis.UpgradeScheduleDeployConfig{
				L2GenesisRegolithTimeOffset: u64UtilPtr(0),
				L2GenesisCanyonTimeOffset:   u64UtilPtr(0),
				L2GenesisDeltaTimeOffset:    u64UtilPtr(0),
				L2GenesisEcotoneTimeOffset:  u64UtilPtr(0),
				L2GenesisFjordTimeOffset:    u64UtilPtr(0),
				L2GenesisGraniteTimeOffset:  u64UtilPtr(0),
				UseInterop:                  false,
			},
			L2CoreDeployConfig: genesis.L2CoreDeployConfig{
				L2ChainID:                 1,
				L2BlockTime:               2,
				FinalizationPeriodSeconds: 12,
				MaxSequencerDrift:         600,
				SequencerWindowSize:       3600,
				ChannelTimeoutBedrock:     300,
				SystemConfigStartBlock:    0,
			},
		},
		FaultProofDeployConfig: genesis.FaultProofDeployConfig{
			FaultGameWithdrawalDelay:        604800,
			PreimageOracleMinProposalSize:   126000,
			PreimageOracleChallengePeriod:   86400,
			ProofMaturityDelaySeconds:       604800,
			DisputeGameFinalityDelaySeconds: 302400,
		},
	}
}

func mustHexBigFromHex(hex string) *hexutil.Big {
	num := hexutil.MustDecodeBig(hex)
	hexBig := hexutil.Big(*num)
	return &hexBig
}

func u64UtilPtr(in uint64) *hexutil.Uint64 {
	util := hexutil.Uint64(in)
	return &util
}
