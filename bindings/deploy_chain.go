// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DeployChainAddressConfiguration is an auto generated low-level Go binding around an user-defined struct.
type DeployChainAddressConfiguration struct {
	Batcher           common.Address
	Proposer          common.Address
	UnsafeBlockSigner common.Address
}

// DeployChainDeployAddresses is an auto generated low-level Go binding around an user-defined struct.
type DeployChainDeployAddresses struct {
	L2OutputOracle               common.Address
	SystemConfig                 common.Address
	OptimismPortal               common.Address
	L1CrossDomainMessenger       common.Address
	L1StandardBridge             common.Address
	L1ERC721Bridge               common.Address
	OptimismMintableERC20Factory common.Address
}

// DeployChainGasConfiguration is an auto generated low-level Go binding around an user-defined struct.
type DeployChainGasConfiguration struct {
	BasefeeScalar     uint32
	BlobbasefeeScalar uint32
	GasLimit          uint64
	GasToken          common.Address
}

// DeployChainGenesisConfiguration is an auto generated low-level Go binding around an user-defined struct.
type DeployChainGenesisConfiguration struct {
	L1Number    uint64
	L2Hash      [32]byte
	L2StateRoot [32]byte
	L2Time      uint64
}

// DeployChainMetaData contains all meta data concerning the DeployChain contract.
var DeployChainMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_proxyAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_superchainConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_protocolVersions\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_PASSER_STORAGE_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateBatchInbox\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.GenesisConfiguration\",\"components\":[{\"name\":\"l1Number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l2Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2StateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2Time\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"gasConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.GasConfiguration\",\"components\":[{\"name\":\"basefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobbasefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"addressConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.AddressConfiguration\",\"components\":[{\"name\":\"batcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unsafeBlockSigner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"proofsEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployAddresses\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.DeployAddresses\",\"components\":[{\"name\":\"l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployProxy\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1CrossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ERC721Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1StandardBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismMintableERC20Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismPortal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolVersions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAddress\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"superchainConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deploy\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"batchInbox\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structDeployChain.DeployAddresses\",\"components\":[{\"name\":\"l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false}]",
	Bin: "0x6101c06040523480156200001257600080fd5b506040516200252d3803806200252d83398101604081905262000035916200009c565b6001600160a01b03998a1660805297891660a05295881660c05293871660e05291861661010052851661012052841661014052831661016052821661018052166101a0526200016b565b80516001600160a01b03811681146200009757600080fd5b919050565b6000806000806000806000806000806101408b8d031215620000bd57600080fd5b620000c88b6200007f565b9950620000d860208c016200007f565b9850620000e860408c016200007f565b9750620000f860608c016200007f565b96506200010860808c016200007f565b95506200011860a08c016200007f565b94506200012860c08c016200007f565b93506200013860e08c016200007f565b9250620001496101008c016200007f565b91506200015a6101208c016200007f565b90509295989b9194979a5092959850565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051612289620002a460003960006102440152600081816101bc015281816110fb015281816114960152818161152b01526115b701526000818161021d0152818161052401528181610bbe01526110270152600081816102a7015281816105f301528181610c8d01526114d2015260008181610280015281816106c201528181610d5c01526116b50152600081816103230152818161067d01528181610d17015261159101526000818161010a0152818161063801528181610cd2015261155f0152600081816101950152818161056901528181610c0301526112ff01526000818161015b015281816105ae01528181610c4801526111260152600081816101f6015281816118020152611a8801526122896000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636624856a11610097578063aabcb26e11610066578063aabcb26e146102c9578063beab4f7e146102fe578063c4e8ddfa1461031e578063dda8804c1461034557600080fd5b80636624856a1461023f57806394e49a1b146102665780639b7d7f0a1461027b578063a7119869146102a257600080fd5b806335e80ab3116100d357806335e80ab3146101b757806336e0909b146101de5780633e47158c146101f15780634d9f15591461021857600080fd5b8063078f29cf146101055780630a49cb0314610156578063113bfd4a1461017d57806333d7e2bd14610190575b600080fd5b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c61018b366004611b1a565b610358565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c6101ec366004611b4a565b6103c8565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b610279610274366004611c72565b610450565b005b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6102f07f8ed4baae3a927be3dea54996b4d5899f8c01e7594bf50b17dc1e741388ce3d1281565b60405190815260200161014d565b61031161030c366004611b4a565b6104da565b60405161014d9190611da5565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61012c610353366004611b1a565b610704565b600080838360405160200161039c92919091825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340190565b6040516020818303038152906040528051906020012090506103be3082610937565b9150505b92915050565b6000805b8215610406576103dd600a84611e45565b6103e882600a611e88565b6103f29190611ec5565b90506103ff600a84611edd565b92506103cc565b60005b81156104335761041a600a83611e45565b60049190911b1761042c600a83611edd565b9150610409565b73ff00000000000000000000000000000000000000179392505050565b600061045b86610b74565b90506000610470878787876000015186610d80565b9050600061047d886103c8565b905061048d868683858789610fda565b815160208301516040518a927f49ea8b4c640f12c7d41cb7b7931d984f226f95ce1d55e1e449ee3d61b877c1ad926104c89286908990611ef1565b60405180910390a25050505050505050565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e00160405280610548847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff16815260200161058d847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff1681526020016105d2847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff168152602001610617847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff16815260200161065c847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff1681526020016106a1847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff1681526020016106e6847f0000000000000000000000000000000000000000000000000000000000000000610704565b73ffffffffffffffffffffffffffffffffffffffff16905292915050565b600080838360405160200161074892919091825260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015260340190565b6040516020818303038152906040528051906020012090506103be30826040517f600661010e565b730000000000000000000000000000000000000000000000008152606092831b60088201527f90915561012a8060295f395ff35f365f600860e5565b80549091805480331433601c8201527f15171560545760045f5f375f5160e01c8063f851a4401460a75780635c60da1b603c8201527f1460a45780638f2839701460b45780633659cfe61460b157634f1ef2861460af605c8201527f575b63204e1c7a60e01b5f5282801560e2576004525f5f60245f845afa3d5f5f607c8201527f3e3d60201416805f510290158402015f875f89895f375f935af43d5f893d6020609c8201527f5260205f523e5f3d89019160a257fd5bf35b50505b505f5260205ff35b5f5b9360bc8201527f915b5050602060045f375f518091559160de57903333602060445f375f51956060dc8201527f64955050604096506054565b5f5ff35b5ffd5b7f360894a13ba1a3210667c82860fc8201527f492db98dca3e2076cc3735a920a3ca505d382bbc7fb53127684a568b3173ae1361011c8201527fb9f8a6016e243e63b6e8ee1178d6a717850b5d61039156ff000000000000000061013c8201523090921b610154830152610168820152610153808220610188830152605591012090565b6040517f600661010e565b730000000000000000000000000000000000000000000000008152606083901b60088201527f90915561012a8060295f395ff35f365f600860e5565b80549091805480331433601c8201527f15171560545760045f5f375f5160e01c8063f851a4401460a75780635c60da1b603c8201527f1460a45780638f2839701460b45780633659cfe61460b157634f1ef2861460af605c8201527f575b63204e1c7a60e01b5f5282801560e2576004525f5f60245f845afa3d5f5f607c8201527f3e3d60201416805f510290158402015f875f89895f375f935af43d5f893d6020609c8201527f5260205f523e5f3d89019160a257fd5bf35b50505b505f5260205ff35b5f5b9360bc8201527f915b5050602060045f375f518091559160de57903333602060445f375f51956060dc8201527f64955050604096506054565b5f5ff35b5ffd5b7f360894a13ba1a3210667c82860fc8201527f492db98dca3e2076cc3735a920a3ca505d382bbc7fb53127684a568b3173ae1361011c8201527fb9f8a6016e243e63b6e8ee1178d6a717850b5d6103915600000000000000000061013c820152600090826101538284f591505073ffffffffffffffffffffffffffffffffffffffff81166103c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f50726f78793a2063726561746532206661696c6564000000000000000000000060448201526064015b60405180910390fd5b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e00160405280610be2847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff168152602001610c27847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff168152602001610c6c847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff168152602001610cb1847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff168152602001610cf6847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff168152602001610d3b847f0000000000000000000000000000000000000000000000000000000000000000610358565b73ffffffffffffffffffffffffffffffffffffffff1681526020016106e6847f0000000000000000000000000000000000000000000000000000000000000000610358565b6040805180820190915260008082526020820152845167ffffffffffffffff164080610e2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4465706c6f79436861696e3a2067656e6573697320626c6f636b68617368206e60448201527f6f7420617661696c61626c6500000000000000000000000000000000000000006064820152608401610b6b565b6000856000015163ffffffff166020876020015163ffffffff16901b60f86001901b171760001b905060008089848a602001518b606001518a878d604001518c604001518d60200151604051602001610f339a9998979695949392919060c09a8b1b7fffffffffffffffff0000000000000000000000000000000000000000000000009081168252600882019a909a526028810198909852604888019690965293881b87166068870152606092831b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811660708801526084870192909252871b90951660a485015290811b841660ac8401521b9091169181019190915260d40190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012060808401835260008085528c840151858401527f8ed4baae3a927be3dea54996b4d5899f8c01e7594bf50b17dc1e741388ce3d1293850193909352908b01516060840152925090610fb99061175c565b60408051808201909152928352602083015250925050505b95945050505050565b815160208084015185519186015160405173ffffffffffffffffffffffffffffffffffffffff90921660248301526044820192909252606481019190915282151560848201526110c591907f00000000000000000000000000000000000000000000000000000000000000009060a401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fb8205148000000000000000000000000000000000000000000000000000000001790526117b8565b60408083015183516020850151925173ffffffffffffffffffffffffffffffffffffffff918216602482015292811660448401527f00000000000000000000000000000000000000000000000000000000000000001660648301526111c5917f0000000000000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526117b8565b60006112f38388606001516040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e00160405280846060015173ffffffffffffffffffffffffffffffffffffffff1681526020018460a0015173ffffffffffffffffffffffffffffffffffffffff168152602001846080015173ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001846040015173ffffffffffffffffffffffffffffffffffffffff1681526020018460c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815250905092915050565b905061146d83602001517f000000000000000000000000000000000000000000000000000000000000000089600001518a602001518a6000015173ffffffffffffffffffffffffffffffffffffffff1660001b8c604001518c604001516113d36040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526301312d008152600a6020820152600891810191909152633b9aca006060820152620f424060808201526fffffffffffffffffffffffffffffffff60a082015290565b8d8f602001518b6040516024016113f299989796959493929190611f8c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fdc7e20a5000000000000000000000000000000000000000000000000000000001790526117b8565b60608301516040808501516020860151915173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301529182166044820152911660648201526114fa91907f00000000000000000000000000000000000000000000000000000000000000009060840161114a565b60808301516060840151602085015160405173ffffffffffffffffffffffffffffffffffffffff92831660248201527f0000000000000000000000000000000000000000000000000000000000000000831660448201529116606482015261158791907f00000000000000000000000000000000000000000000000000000000000000009060840161114a565b6116848360a001517f000000000000000000000000000000000000000000000000000000000000000085606001517f000000000000000000000000000000000000000000000000000000000000000060405160240161160992919073ffffffffffffffffffffffffffffffffffffffff92831681529116602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc955000000000000000000000000000000000000000000000000000000001790526117b8565b60c0830151608084015160405173ffffffffffffffffffffffffffffffffffffffff909116602482015261175391907f000000000000000000000000000000000000000000000000000000000000000090604401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc4d66de8000000000000000000000000000000000000000000000000000000001790526117b8565b50505050505050565b6000816000015182602001518360400151846060015160405160200161179b949392919093845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b6040517f204e1c7a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015284916000917f0000000000000000000000000000000000000000000000000000000000000000169063204e1c7a90602401602060405180830381865afa158015611849573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186d91906120ef565b905073ffffffffffffffffffffffffffffffffffffffff8116611912576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f4465706c6f79436861696e3a20696e76616c696420696d706c656d656e74617460448201527f696f6e00000000000000000000000000000000000000000000000000000000006064820152608401610b6b565b6040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831690634f1ef286906119669084908790600401612143565b6000604051808303816000875af1158015611985573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526119cb91908101906121b1565b506040517f3659cfe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152831690633659cfe690602401600060405180830381600087803b158015611a3557600080fd5b505af1158015611a49573d6000803e3d6000fd5b50506040517f8f28397000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116600483015285169250638f2839709150602401600060405180830381600087803b158015611ad657600080fd5b505af1158015611aea573d6000803e3d6000fd5b505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611b1757600080fd5b50565b60008060408385031215611b2d57600080fd5b823591506020830135611b3f81611af5565b809150509250929050565b600060208284031215611b5c57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715611bb557611bb5611b63565b60405290565b803567ffffffffffffffff81168114611bd357600080fd5b919050565b803563ffffffff81168114611bd357600080fd5b600060608284031215611bfe57600080fd5b6040516060810181811067ffffffffffffffff82111715611c2157611c21611b63565b6040529050808235611c3281611af5565b81526020830135611c4281611af5565b60208201526040830135611c5581611af5565b6040919091015292915050565b80358015158114611bd357600080fd5b60008060008060008587036101a0811215611c8c57600080fd5b8635955060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082011215611cc057600080fd5b611cc8611b92565b611cd460208901611bbb565b81526040880135602082015260608801356040820152611cf660808901611bbb565b6060820152945060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6082011215611d2d57600080fd5b50611d36611b92565b611d4260a08801611bd8565b8152611d5060c08801611bd8565b6020820152611d6160e08801611bbb565b6040820152610100870135611d7581611af5565b60608201529250611d8a876101208801611bec565b9150611d996101808701611c62565b90509295509295909350565b60e081016103c2828473ffffffffffffffffffffffffffffffffffffffff8082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611e5457611e54611e16565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ec057611ec0611e59565b500290565b60008219821115611ed857611ed8611e59565b500190565b600082611eec57611eec611e16565b500490565b8481526020810184905273ffffffffffffffffffffffffffffffffffffffff831660408201526101408101610fd1606083018473ffffffffffffffffffffffffffffffffffffffff8082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b60006102808201905063ffffffff808c168352808b16602084015289604084015267ffffffffffffffff8916606084015273ffffffffffffffffffffffffffffffffffffffff881660808401528087511660a084015260ff60208801511660c084015260ff60408801511660e08401528060608801511661010084015280608088015116610120840152506fffffffffffffffffffffffffffffffff60a08701511661014083015261205761016083018673ffffffffffffffffffffffffffffffffffffffff169052565b73ffffffffffffffffffffffffffffffffffffffff8416610180830152825173ffffffffffffffffffffffffffffffffffffffff9081166101a0840152602084015181166101c0840152604084015181166101e0840152606084015181166102008401526080840151811661022084015260a0840151811661024084015260c0840151166102608301529a9950505050505050505050565b60006020828403121561210157600080fd5b815161210c81611af5565b9392505050565b60005b8381101561212e578181015183820152602001612116565b8381111561213d576000848401525b50505050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000825180604084015261217e816060850160208701612113565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016060019392505050565b6000602082840312156121c357600080fd5b815167ffffffffffffffff808211156121db57600080fd5b818401915084601f8301126121ef57600080fd5b81518181111561220157612201611b63565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561224757612247611b63565b8160405282815287602084870101111561226057600080fd5b612271836020830160208801612113565b97965050505050505056fea164736f6c634300080f000a",
}

// DeployChainABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployChainMetaData.ABI instead.
var DeployChainABI = DeployChainMetaData.ABI

// DeployChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeployChainMetaData.Bin instead.
var DeployChainBin = DeployChainMetaData.Bin

// DeployDeployChain deploys a new Ethereum contract, binding an instance of DeployChain to it.
func DeployDeployChain(auth *bind.TransactOpts, backend bind.ContractBackend, _proxyAdmin common.Address, _optimismPortal common.Address, _systemConfig common.Address, _l1StandardBridge common.Address, _l1ERC721Bridge common.Address, _optimismMintableERC20Factory common.Address, _l1CrossDomainMessenger common.Address, _l2OutputOracle common.Address, _superchainConfig common.Address, _protocolVersions common.Address) (common.Address, *types.Transaction, *DeployChain, error) {
	parsed, err := DeployChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeployChainBin), backend, _proxyAdmin, _optimismPortal, _systemConfig, _l1StandardBridge, _l1ERC721Bridge, _optimismMintableERC20Factory, _l1CrossDomainMessenger, _l2OutputOracle, _superchainConfig, _protocolVersions)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeployChain{DeployChainCaller: DeployChainCaller{contract: contract}, DeployChainTransactor: DeployChainTransactor{contract: contract}, DeployChainFilterer: DeployChainFilterer{contract: contract}}, nil
}

// DeployChain is an auto generated Go binding around an Ethereum contract.
type DeployChain struct {
	DeployChainCaller     // Read-only binding to the contract
	DeployChainTransactor // Write-only binding to the contract
	DeployChainFilterer   // Log filterer for contract events
}

// DeployChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployChainSession struct {
	Contract     *DeployChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployChainCallerSession struct {
	Contract *DeployChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DeployChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployChainTransactorSession struct {
	Contract     *DeployChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DeployChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployChainRaw struct {
	Contract *DeployChain // Generic contract binding to access the raw methods on
}

// DeployChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployChainCallerRaw struct {
	Contract *DeployChainCaller // Generic read-only contract binding to access the raw methods on
}

// DeployChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployChainTransactorRaw struct {
	Contract *DeployChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployChain creates a new instance of DeployChain, bound to a specific deployed contract.
func NewDeployChain(address common.Address, backend bind.ContractBackend) (*DeployChain, error) {
	contract, err := bindDeployChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeployChain{DeployChainCaller: DeployChainCaller{contract: contract}, DeployChainTransactor: DeployChainTransactor{contract: contract}, DeployChainFilterer: DeployChainFilterer{contract: contract}}, nil
}

// NewDeployChainCaller creates a new read-only instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainCaller(address common.Address, caller bind.ContractCaller) (*DeployChainCaller, error) {
	contract, err := bindDeployChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployChainCaller{contract: contract}, nil
}

// NewDeployChainTransactor creates a new write-only instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployChainTransactor, error) {
	contract, err := bindDeployChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployChainTransactor{contract: contract}, nil
}

// NewDeployChainFilterer creates a new log filterer instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployChainFilterer, error) {
	contract, err := bindDeployChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployChainFilterer{contract: contract}, nil
}

// bindDeployChain binds a generic wrapper to an already deployed contract.
func bindDeployChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeployChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployChain *DeployChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployChain.Contract.DeployChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployChain *DeployChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployChain *DeployChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployChain *DeployChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployChain *DeployChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployChain *DeployChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployChain.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEPASSERSTORAGEHASH is a free data retrieval call binding the contract method 0xaabcb26e.
//
// Solidity: function MESSAGE_PASSER_STORAGE_HASH() view returns(bytes32)
func (_DeployChain *DeployChainCaller) MESSAGEPASSERSTORAGEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "MESSAGE_PASSER_STORAGE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGEPASSERSTORAGEHASH is a free data retrieval call binding the contract method 0xaabcb26e.
//
// Solidity: function MESSAGE_PASSER_STORAGE_HASH() view returns(bytes32)
func (_DeployChain *DeployChainSession) MESSAGEPASSERSTORAGEHASH() ([32]byte, error) {
	return _DeployChain.Contract.MESSAGEPASSERSTORAGEHASH(&_DeployChain.CallOpts)
}

// MESSAGEPASSERSTORAGEHASH is a free data retrieval call binding the contract method 0xaabcb26e.
//
// Solidity: function MESSAGE_PASSER_STORAGE_HASH() view returns(bytes32)
func (_DeployChain *DeployChainCallerSession) MESSAGEPASSERSTORAGEHASH() ([32]byte, error) {
	return _DeployChain.Contract.MESSAGEPASSERSTORAGEHASH(&_DeployChain.CallOpts)
}

// CalculateBatchInbox is a free data retrieval call binding the contract method 0x36e0909b.
//
// Solidity: function calculateBatchInbox(uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainCaller) CalculateBatchInbox(opts *bind.CallOpts, chainID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "calculateBatchInbox", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateBatchInbox is a free data retrieval call binding the contract method 0x36e0909b.
//
// Solidity: function calculateBatchInbox(uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainSession) CalculateBatchInbox(chainID *big.Int) (common.Address, error) {
	return _DeployChain.Contract.CalculateBatchInbox(&_DeployChain.CallOpts, chainID)
}

// CalculateBatchInbox is a free data retrieval call binding the contract method 0x36e0909b.
//
// Solidity: function calculateBatchInbox(uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainCallerSession) CalculateBatchInbox(chainID *big.Int) (common.Address, error) {
	return _DeployChain.Contract.CalculateBatchInbox(&_DeployChain.CallOpts, chainID)
}

// DeployAddresses is a free data retrieval call binding the contract method 0xbeab4f7e.
//
// Solidity: function deployAddresses(uint256 chainID) view returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainCaller) DeployAddresses(opts *bind.CallOpts, chainID *big.Int) (DeployChainDeployAddresses, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "deployAddresses", chainID)

	if err != nil {
		return *new(DeployChainDeployAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(DeployChainDeployAddresses)).(*DeployChainDeployAddresses)

	return out0, err

}

// DeployAddresses is a free data retrieval call binding the contract method 0xbeab4f7e.
//
// Solidity: function deployAddresses(uint256 chainID) view returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainSession) DeployAddresses(chainID *big.Int) (DeployChainDeployAddresses, error) {
	return _DeployChain.Contract.DeployAddresses(&_DeployChain.CallOpts, chainID)
}

// DeployAddresses is a free data retrieval call binding the contract method 0xbeab4f7e.
//
// Solidity: function deployAddresses(uint256 chainID) view returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainCallerSession) DeployAddresses(chainID *big.Int) (DeployChainDeployAddresses, error) {
	return _DeployChain.Contract.DeployAddresses(&_DeployChain.CallOpts, chainID)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainCaller) L1CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1CrossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainSession) L1CrossDomainMessenger() (common.Address, error) {
	return _DeployChain.Contract.L1CrossDomainMessenger(&_DeployChain.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _DeployChain.Contract.L1CrossDomainMessenger(&_DeployChain.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainCaller) L1ERC721Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1ERC721Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainSession) L1ERC721Bridge() (common.Address, error) {
	return _DeployChain.Contract.L1ERC721Bridge(&_DeployChain.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1ERC721Bridge() (common.Address, error) {
	return _DeployChain.Contract.L1ERC721Bridge(&_DeployChain.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainCaller) L1StandardBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1StandardBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainSession) L1StandardBridge() (common.Address, error) {
	return _DeployChain.Contract.L1StandardBridge(&_DeployChain.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1StandardBridge() (common.Address, error) {
	return _DeployChain.Contract.L1StandardBridge(&_DeployChain.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_DeployChain *DeployChainCaller) L2OutputOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l2OutputOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_DeployChain *DeployChainSession) L2OutputOracle() (common.Address, error) {
	return _DeployChain.Contract.L2OutputOracle(&_DeployChain.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_DeployChain *DeployChainCallerSession) L2OutputOracle() (common.Address, error) {
	return _DeployChain.Contract.L2OutputOracle(&_DeployChain.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainCaller) OptimismMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "optimismMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _DeployChain.Contract.OptimismMintableERC20Factory(&_DeployChain.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainCallerSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _DeployChain.Contract.OptimismMintableERC20Factory(&_DeployChain.CallOpts)
}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address)
func (_DeployChain *DeployChainCaller) OptimismPortal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "optimismPortal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address)
func (_DeployChain *DeployChainSession) OptimismPortal() (common.Address, error) {
	return _DeployChain.Contract.OptimismPortal(&_DeployChain.CallOpts)
}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address)
func (_DeployChain *DeployChainCallerSession) OptimismPortal() (common.Address, error) {
	return _DeployChain.Contract.OptimismPortal(&_DeployChain.CallOpts)
}

// ProtocolVersions is a free data retrieval call binding the contract method 0x6624856a.
//
// Solidity: function protocolVersions() view returns(address)
func (_DeployChain *DeployChainCaller) ProtocolVersions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "protocolVersions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolVersions is a free data retrieval call binding the contract method 0x6624856a.
//
// Solidity: function protocolVersions() view returns(address)
func (_DeployChain *DeployChainSession) ProtocolVersions() (common.Address, error) {
	return _DeployChain.Contract.ProtocolVersions(&_DeployChain.CallOpts)
}

// ProtocolVersions is a free data retrieval call binding the contract method 0x6624856a.
//
// Solidity: function protocolVersions() view returns(address)
func (_DeployChain *DeployChainCallerSession) ProtocolVersions() (common.Address, error) {
	return _DeployChain.Contract.ProtocolVersions(&_DeployChain.CallOpts)
}

// ProxyAddress is a free data retrieval call binding the contract method 0xdda8804c.
//
// Solidity: function proxyAddress(uint256 chainID, address implementation) view returns(address)
func (_DeployChain *DeployChainCaller) ProxyAddress(opts *bind.CallOpts, chainID *big.Int, implementation common.Address) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "proxyAddress", chainID, implementation)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAddress is a free data retrieval call binding the contract method 0xdda8804c.
//
// Solidity: function proxyAddress(uint256 chainID, address implementation) view returns(address)
func (_DeployChain *DeployChainSession) ProxyAddress(chainID *big.Int, implementation common.Address) (common.Address, error) {
	return _DeployChain.Contract.ProxyAddress(&_DeployChain.CallOpts, chainID, implementation)
}

// ProxyAddress is a free data retrieval call binding the contract method 0xdda8804c.
//
// Solidity: function proxyAddress(uint256 chainID, address implementation) view returns(address)
func (_DeployChain *DeployChainCallerSession) ProxyAddress(chainID *big.Int, implementation common.Address) (common.Address, error) {
	return _DeployChain.Contract.ProxyAddress(&_DeployChain.CallOpts, chainID, implementation)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainSession) ProxyAdmin() (common.Address, error) {
	return _DeployChain.Contract.ProxyAdmin(&_DeployChain.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainCallerSession) ProxyAdmin() (common.Address, error) {
	return _DeployChain.Contract.ProxyAdmin(&_DeployChain.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainCaller) SuperchainConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "superchainConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainSession) SuperchainConfig() (common.Address, error) {
	return _DeployChain.Contract.SuperchainConfig(&_DeployChain.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainCallerSession) SuperchainConfig() (common.Address, error) {
	return _DeployChain.Contract.SuperchainConfig(&_DeployChain.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainCaller) SystemConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "systemConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainSession) SystemConfig() (common.Address, error) {
	return _DeployChain.Contract.SystemConfig(&_DeployChain.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainCallerSession) SystemConfig() (common.Address, error) {
	return _DeployChain.Contract.SystemConfig(&_DeployChain.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x94e49a1b.
//
// Solidity: function deploy(uint256 chainID, (uint64,bytes32,bytes32,uint64) genesisConfig, (uint32,uint32,uint64,address) gasConfig, (address,address,address) addressConfig, bool proofsEnabled) returns()
func (_DeployChain *DeployChainTransactor) Deploy(opts *bind.TransactOpts, chainID *big.Int, genesisConfig DeployChainGenesisConfiguration, gasConfig DeployChainGasConfiguration, addressConfig DeployChainAddressConfiguration, proofsEnabled bool) (*types.Transaction, error) {
	return _DeployChain.contract.Transact(opts, "deploy", chainID, genesisConfig, gasConfig, addressConfig, proofsEnabled)
}

// Deploy is a paid mutator transaction binding the contract method 0x94e49a1b.
//
// Solidity: function deploy(uint256 chainID, (uint64,bytes32,bytes32,uint64) genesisConfig, (uint32,uint32,uint64,address) gasConfig, (address,address,address) addressConfig, bool proofsEnabled) returns()
func (_DeployChain *DeployChainSession) Deploy(chainID *big.Int, genesisConfig DeployChainGenesisConfiguration, gasConfig DeployChainGasConfiguration, addressConfig DeployChainAddressConfiguration, proofsEnabled bool) (*types.Transaction, error) {
	return _DeployChain.Contract.Deploy(&_DeployChain.TransactOpts, chainID, genesisConfig, gasConfig, addressConfig, proofsEnabled)
}

// Deploy is a paid mutator transaction binding the contract method 0x94e49a1b.
//
// Solidity: function deploy(uint256 chainID, (uint64,bytes32,bytes32,uint64) genesisConfig, (uint32,uint32,uint64,address) gasConfig, (address,address,address) addressConfig, bool proofsEnabled) returns()
func (_DeployChain *DeployChainTransactorSession) Deploy(chainID *big.Int, genesisConfig DeployChainGenesisConfiguration, gasConfig DeployChainGasConfiguration, addressConfig DeployChainAddressConfiguration, proofsEnabled bool) (*types.Transaction, error) {
	return _DeployChain.Contract.Deploy(&_DeployChain.TransactOpts, chainID, genesisConfig, gasConfig, addressConfig, proofsEnabled)
}

// DeployProxy is a paid mutator transaction binding the contract method 0x113bfd4a.
//
// Solidity: function deployProxy(uint256 chainID, address implementation) returns(address)
func (_DeployChain *DeployChainTransactor) DeployProxy(opts *bind.TransactOpts, chainID *big.Int, implementation common.Address) (*types.Transaction, error) {
	return _DeployChain.contract.Transact(opts, "deployProxy", chainID, implementation)
}

// DeployProxy is a paid mutator transaction binding the contract method 0x113bfd4a.
//
// Solidity: function deployProxy(uint256 chainID, address implementation) returns(address)
func (_DeployChain *DeployChainSession) DeployProxy(chainID *big.Int, implementation common.Address) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployProxy(&_DeployChain.TransactOpts, chainID, implementation)
}

// DeployProxy is a paid mutator transaction binding the contract method 0x113bfd4a.
//
// Solidity: function deployProxy(uint256 chainID, address implementation) returns(address)
func (_DeployChain *DeployChainTransactorSession) DeployProxy(chainID *big.Int, implementation common.Address) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployProxy(&_DeployChain.TransactOpts, chainID, implementation)
}

// DeployChainDeployIterator is returned from FilterDeploy and is used to iterate over the raw logs and unpacked data for Deploy events raised by the DeployChain contract.
type DeployChainDeployIterator struct {
	Event *DeployChainDeploy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeployChainDeployIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployChainDeploy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeployChainDeploy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeployChainDeployIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployChainDeployIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployChainDeploy represents a Deploy event raised by the DeployChain contract.
type DeployChainDeploy struct {
	ChainID    *big.Int
	ConfigHash [32]byte
	OutputRoot [32]byte
	BatchInbox common.Address
	Addresses  DeployChainDeployAddresses
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeploy is a free log retrieval operation binding the contract event 0x49ea8b4c640f12c7d41cb7b7931d984f226f95ce1d55e1e449ee3d61b877c1ad.
//
// Solidity: event Deploy(uint256 indexed chainID, bytes32 configHash, bytes32 outputRoot, address batchInbox, (address,address,address,address,address,address,address) addresses)
func (_DeployChain *DeployChainFilterer) FilterDeploy(opts *bind.FilterOpts, chainID []*big.Int) (*DeployChainDeployIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _DeployChain.contract.FilterLogs(opts, "Deploy", chainIDRule)
	if err != nil {
		return nil, err
	}
	return &DeployChainDeployIterator{contract: _DeployChain.contract, event: "Deploy", logs: logs, sub: sub}, nil
}

// WatchDeploy is a free log subscription operation binding the contract event 0x49ea8b4c640f12c7d41cb7b7931d984f226f95ce1d55e1e449ee3d61b877c1ad.
//
// Solidity: event Deploy(uint256 indexed chainID, bytes32 configHash, bytes32 outputRoot, address batchInbox, (address,address,address,address,address,address,address) addresses)
func (_DeployChain *DeployChainFilterer) WatchDeploy(opts *bind.WatchOpts, sink chan<- *DeployChainDeploy, chainID []*big.Int) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _DeployChain.contract.WatchLogs(opts, "Deploy", chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployChainDeploy)
				if err := _DeployChain.contract.UnpackLog(event, "Deploy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeploy is a log parse operation binding the contract event 0x49ea8b4c640f12c7d41cb7b7931d984f226f95ce1d55e1e449ee3d61b877c1ad.
//
// Solidity: event Deploy(uint256 indexed chainID, bytes32 configHash, bytes32 outputRoot, address batchInbox, (address,address,address,address,address,address,address) addresses)
func (_DeployChain *DeployChainFilterer) ParseDeploy(log types.Log) (*DeployChainDeploy, error) {
	event := new(DeployChainDeploy)
	if err := _DeployChain.contract.UnpackLog(event, "Deploy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
