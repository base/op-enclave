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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_proxyAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_superchainConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_protocolVersions\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_PASSER_STORAGE_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateBatchInbox\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.GenesisConfiguration\",\"components\":[{\"name\":\"l1Number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l2Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2StateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2Time\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"gasConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.GasConfiguration\",\"components\":[{\"name\":\"basefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobbasefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"addressConfig\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.AddressConfiguration\",\"components\":[{\"name\":\"batcher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unsafeBlockSigner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"proofsEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployAddresses\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.DeployAddresses\",\"components\":[{\"name\":\"l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1CrossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ERC721Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1StandardBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismMintableERC20Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismPortal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolVersions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAddress\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"superchainConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deploy\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"configHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"batchInbox\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"addresses\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structDeployChain.DeployAddresses\",\"components\":[{\"name\":\"l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101c06040523480156200001257600080fd5b506040516200230a3803806200230a8339810160408190526200003591620002bb565b62000040336200016d565b6001600160a01b038a166200005457600080fd5b6001600160a01b0389166200006857600080fd5b6001600160a01b0388166200007c57600080fd5b6001600160a01b0387166200009057600080fd5b6001600160a01b038616620000a457600080fd5b6001600160a01b038516620000b857600080fd5b6001600160a01b038416620000cc57600080fd5b6001600160a01b038316620000e057600080fd5b6001600160a01b038216620000f457600080fd5b6001600160a01b0381166200010857600080fd5b6001600160a01b03808b1660805289811660a05288811660c05287811660e052868116610100528581166101205284811661014052838116610160528281166101805281166101a0526200015c8b620001bd565b50505050505050505050506200039e565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b620001c762000240565b6001600160a01b038116620002325760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6200023d816200016d565b50565b6000546001600160a01b031633146200029c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000229565b565b80516001600160a01b0381168114620002b657600080fd5b919050565b60008060008060008060008060008060006101608c8e031215620002de57600080fd5b620002e98c6200029e565b9a50620002f960208d016200029e565b99506200030960408d016200029e565b98506200031960608d016200029e565b97506200032960808d016200029e565b96506200033960a08d016200029e565b95506200034960c08d016200029e565b94506200035960e08d016200029e565b93506200036a6101008d016200029e565b92506200037b6101208d016200029e565b91506200038c6101408d016200029e565b90509295989b509295989b9093969950565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051611e64620004a660003960006102670152600081816101df01528181611066015281816113770152818161144001526114fa015260008181610240015281816106fe0152610b4b0152600081816102f0015281816107cd0152610c1a0152600081816102c90152818161089c0152610ce901526000818161036c015281816108570152610ca4015260008181610140015281816108120152610c5f0152600081816101b8015281816107430152610b90015260008181610191015281816107880152610bd50152600081816102190152818161040c01526115fb0152611e646000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80638da5cb5b116100b2578063aabcb26e11610081578063c4e8ddfa11610066578063c4e8ddfa14610367578063d655a76f1461038e578063f2fde38b146103a157600080fd5b8063aabcb26e14610312578063beab4f7e1461034757600080fd5b80638da5cb5b1461029357806394e49a1b146102b15780639b7d7f0a146102c4578063a7119869146102eb57600080fd5b8063380cb000116101095780634d9f1559116100ee5780634d9f15591461023b5780636624856a14610262578063715018a61461028957600080fd5b8063380cb000146102015780633e47158c1461021457600080fd5b8063078f29cf1461013b5780630a49cb031461018c57806333d7e2bd146101b357806335e80ab3146101da575b600080fd5b6101627f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b61016261020f366004611913565b6103b4565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6102916105e0565b005b60005473ffffffffffffffffffffffffffffffffffffffff16610162565b6102916102bf366004611a5f565b6105f4565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b6103397f8ed4baae3a927be3dea54996b4d5899f8c01e7594bf50b17dc1e741388ce3d1281565b604051908152602001610183565b61035a610355366004611b90565b610688565b6040516101839190611ba9565b6101627f000000000000000000000000000000000000000000000000000000000000000081565b61016261039c366004611c1a565b6108e0565b6102916103af366004611c3e565b610923565b6040517f600661011c565b730000000000000000000000000000000000000000000000008152606083811b60088301527f9055730000000000000000000000000000000000000000000000000000000000601c8301527f0000000000000000000000000000000000000000000000000000000000000000811b601f8301527f905561012280603f5f395ff35f365f600860dd565b805490918054803314331560338301527f171560545760045f5f375f5160e01c8063f851a4401460a25780635c60da1b1460538301527f609f5780638f2839701460af5780633659cfe61460ac57634f1ef2861460aa5760738301527f5b63204e1c7a60e01b5f52826004525f5f60245f845afa3d5f5f3e3d6020141660938301527f805f510290158402015f875f89895f375f935af43d5f893d60205260205f523e60b38301527f5f3d890191609d57fd5bf35b50505b505f5260205ff35b5f5b93915b5050602060d38301527f60045f375f518091559160d957903333602060445f375f51956064955050604060f38301527f96506054565b5f5ff35b7f360894a13ba1a3210667c828492db98dca3e2076cc6101138301527f3735a920a3ca505d382bbc7fb53127684a568b3173ae13b9f8a6016e243e63b66101338301527fe8ee1178d6a717850b5d61039156ff000000000000000000000000000000000061015383015230901b610162820152610176810182905261016180822061019683015260559101206000905b90505b92915050565b6105e86109df565b6105f26000610a60565b565b6105fc6109df565b600061060786610ad5565b9050600061061c878787876000015186610d0e565b9050600061062b6000896108e0565b905061063b868683858789610f68565b815160208301516040518a927f49ea8b4c640f12c7d41cb7b7931d984f226f95ce1d55e1e449ee3d61b877c1ad926106769286908990611c59565b60405180910390a25050505050505050565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526000826040516020016106d691815260200190565b6040516020818303038152906040528051906020012090506040518060e001604052806107237f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1681526020016107687f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1681526020016107ad7f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1681526020016107f27f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1681526020016108377f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff16815260200161087c7f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1681526020016108c17f0000000000000000000000000000000000000000000000000000000000000000846103b4565b73ffffffffffffffffffffffffffffffffffffffff1690529392505050565b60006068826040516020016108f791815260200190565b6040516020818303038152906040528051906020012060001c901c60988460ff16901b17905092915050565b61092b6109df565b73ffffffffffffffffffffffffffffffffffffffff81166109d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6109dc81610a60565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109ca565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152600082604051602001610b2391815260200190565b6040516020818303038152906040528051906020012090506040518060e00160405280610b707f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff168152602001610bb57f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff168152602001610bfa7f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff168152602001610c3f7f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff168152602001610c847f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff168152602001610cc97f0000000000000000000000000000000000000000000000000000000000000000846115f3565b73ffffffffffffffffffffffffffffffffffffffff1681526020016108c17f0000000000000000000000000000000000000000000000000000000000000000846115f3565b6040805180820190915260008082526020820152845167ffffffffffffffff164080610dbc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4465706c6f79436861696e3a2067656e6573697320626c6f636b68617368206e60448201527f6f7420617661696c61626c65000000000000000000000000000000000000000060648201526084016109ca565b6000856000015163ffffffff166020876020015163ffffffff16901b60f86001901b171760001b905060008089848a602001518b606001518a878d604001518c604001518d60200151604051602001610ec19a9998979695949392919060c09a8b1b7fffffffffffffffff0000000000000000000000000000000000000000000000009081168252600882019a909a526028810198909852604888019690965293881b87166068870152606092831b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090811660708801526084870192909252871b90951660a485015290811b841660ac8401521b9091169181019190915260d40190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012060808401835260008085528c840151858401527f8ed4baae3a927be3dea54996b4d5899f8c01e7594bf50b17dc1e741388ce3d1293850193909352908b01516060840152925090610f4790611620565b60408051808201909152928352602083015250925050505b95945050505050565b81516020808401518551918601516040517fb820514800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92831660048201526024810193909352604483015283151560648301529091169063b820514890608401600060405180830381600087803b158015610ff557600080fd5b505af1158015611009573d6000803e3d6000fd5b505050506040828101518351602085015192517fc0c53b8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff918216600482015292811660248401527f000000000000000000000000000000000000000000000000000000000000000081166044840152169063c0c53b8b90606401600060405180830381600087803b1580156110b157600080fd5b505af11580156110c5573d6000803e3d6000fd5b5050505060006111f78388606001516040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e00160405280846060015173ffffffffffffffffffffffffffffffffffffffff1681526020018460a0015173ffffffffffffffffffffffffffffffffffffffff168152602001846080015173ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001846040015173ffffffffffffffffffffffffffffffffffffffff1681526020018460c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815250905092915050565b9050826020015173ffffffffffffffffffffffffffffffffffffffff1663dc7e20a588600001518960200151896000015173ffffffffffffffffffffffffffffffffffffffff1660001b8b604001518b604001516112ce6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526301312d008152600a6020820152600891810191909152633b9aca006060820152620f424060808201526fffffffffffffffffffffffffffffffff60a082015290565b8c8e602001518a6040518a63ffffffff1660e01b81526004016112f999989796959493929190611cf4565b600060405180830381600087803b15801561131357600080fd5b505af1158015611327573d6000803e3d6000fd5b505050506060830151604080850151602086015191517fc0c53b8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166004830152918216602482015291811660448301529091169063c0c53b8b90606401600060405180830381600087803b1580156113d457600080fd5b505af11580156113e8573d6000803e3d6000fd5b505050506080830151606084015160208501516040517fc0c53b8b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92831660048201527f000000000000000000000000000000000000000000000000000000000000000083166024820152908216604482015291169063c0c53b8b90606401600060405180830381600087803b15801561149457600080fd5b505af11580156114a8573d6000803e3d6000fd5b50505060a084015160608501516040517f485cc95500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527f0000000000000000000000000000000000000000000000000000000000000000821660248201529116915063485cc95590604401600060405180830381600087803b15801561154757600080fd5b505af115801561155b573d6000803e3d6000fd5b50505060c084015160808501516040517fc4d66de800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529116915063c4d66de890602401600060405180830381600087803b1580156115d257600080fd5b505af11580156115e6573d6000803e3d6000fd5b5050505050505050505050565b60006105d7837f00000000000000000000000000000000000000000000000000000000000000008461167c565b6000816000015182602001518360400151846060015160405160200161165f949392919093845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b6040517f600661011c565b730000000000000000000000000000000000000000000000008152606084811b60088301527f9055730000000000000000000000000000000000000000000000000000000000601c83015283901b601f8201527f905561012280603f5f395ff35f365f600860dd565b805490918054803314331560338201527f171560545760045f5f375f5160e01c8063f851a4401460a25780635c60da1b1460538201527f609f5780638f2839701460af5780633659cfe61460ac57634f1ef2861460aa5760738201527f5b63204e1c7a60e01b5f52826004525f5f60245f845afa3d5f5f3e3d6020141660938201527f805f510290158402015f875f89895f375f935af43d5f893d60205260205f523e60b38201527f5f3d890191609d57fd5bf35b50505b505f5260205ff35b5f5b93915b5050602060d38201527f60045f375f518091559160d957903333602060445f375f51956064955050604060f38201527f96506054565b5f5ff35b7f360894a13ba1a3210667c828492db98dca3e2076cc6101138201527f3735a920a3ca505d382bbc7fb53127684a568b3173ae13b9f8a6016e243e63b66101338201527fe8ee1178d6a717850b5d61039156000000000000000000000000000000000000610153820152600090826101618284f591505073ffffffffffffffffffffffffffffffffffffffff81166118e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f50726f78793a2063726561746532206661696c6564000000000000000000000060448201526064016109ca565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461190e57600080fd5b919050565b6000806040838503121561192657600080fd5b61192f836118ea565b946020939093013593505050565b6040516080810167ffffffffffffffff81118282101715611987577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803567ffffffffffffffff8116811461190e57600080fd5b803563ffffffff8116811461190e57600080fd5b6000606082840312156119cb57600080fd5b6040516060810181811067ffffffffffffffff82111715611a15577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052905080611a24836118ea565b8152611a32602084016118ea565b6020820152611a43604084016118ea565b60408201525092915050565b8035801515811461190e57600080fd5b60008060008060008587036101a0811215611a7957600080fd5b8635955060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082011215611aad57600080fd5b611ab561193d565b611ac16020890161198d565b81526040880135602082015260608801356040820152611ae36080890161198d565b6060820152945060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6082011215611b1a57600080fd5b50611b2361193d565b611b2f60a088016119a5565b8152611b3d60c088016119a5565b6020820152611b4e60e0880161198d565b6040820152611b6061010088016118ea565b60608201529250611b758761012088016119b9565b9150611b846101808701611a4f565b90509295509295909350565b600060208284031215611ba257600080fd5b5035919050565b60e081016105da828473ffffffffffffffffffffffffffffffffffffffff8082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b60008060408385031215611c2d57600080fd5b823560ff8116811461192f57600080fd5b600060208284031215611c5057600080fd5b6105d7826118ea565b8481526020810184905273ffffffffffffffffffffffffffffffffffffffff831660408201526101408101610f5f606083018473ffffffffffffffffffffffffffffffffffffffff8082511683528060208301511660208401528060408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b60006102808201905063ffffffff808c168352808b16602084015289604084015267ffffffffffffffff8916606084015273ffffffffffffffffffffffffffffffffffffffff881660808401528087511660a084015260ff60208801511660c084015260ff60408801511660e08401528060608801511661010084015280608088015116610120840152506fffffffffffffffffffffffffffffffff60a087015116610140830152611dbf61016083018673ffffffffffffffffffffffffffffffffffffffff169052565b73ffffffffffffffffffffffffffffffffffffffff8416610180830152825173ffffffffffffffffffffffffffffffffffffffff9081166101a0840152602084015181166101c0840152604084015181166101e0840152606084015181166102008401526080840151811661022084015260a0840151811661024084015260c0840151166102608301529a995050505050505050505056fea164736f6c634300080f000a",
}

// DeployChainABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployChainMetaData.ABI instead.
var DeployChainABI = DeployChainMetaData.ABI

// DeployChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeployChainMetaData.Bin instead.
var DeployChainBin = DeployChainMetaData.Bin

// DeployDeployChain deploys a new Ethereum contract, binding an instance of DeployChain to it.
func DeployDeployChain(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _proxyAdmin common.Address, _optimismPortal common.Address, _systemConfig common.Address, _l1StandardBridge common.Address, _l1ERC721Bridge common.Address, _optimismMintableERC20Factory common.Address, _l1CrossDomainMessenger common.Address, _l2OutputOracle common.Address, _superchainConfig common.Address, _protocolVersions common.Address) (common.Address, *types.Transaction, *DeployChain, error) {
	parsed, err := DeployChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeployChainBin), backend, _owner, _proxyAdmin, _optimismPortal, _systemConfig, _l1StandardBridge, _l1ERC721Bridge, _optimismMintableERC20Factory, _l1CrossDomainMessenger, _l2OutputOracle, _superchainConfig, _protocolVersions)
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

// CalculateBatchInbox is a free data retrieval call binding the contract method 0xd655a76f.
//
// Solidity: function calculateBatchInbox(uint8 version, uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainCaller) CalculateBatchInbox(opts *bind.CallOpts, version uint8, chainID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "calculateBatchInbox", version, chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateBatchInbox is a free data retrieval call binding the contract method 0xd655a76f.
//
// Solidity: function calculateBatchInbox(uint8 version, uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainSession) CalculateBatchInbox(version uint8, chainID *big.Int) (common.Address, error) {
	return _DeployChain.Contract.CalculateBatchInbox(&_DeployChain.CallOpts, version, chainID)
}

// CalculateBatchInbox is a free data retrieval call binding the contract method 0xd655a76f.
//
// Solidity: function calculateBatchInbox(uint8 version, uint256 chainID) pure returns(address)
func (_DeployChain *DeployChainCallerSession) CalculateBatchInbox(version uint8, chainID *big.Int) (common.Address, error) {
	return _DeployChain.Contract.CalculateBatchInbox(&_DeployChain.CallOpts, version, chainID)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeployChain *DeployChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeployChain *DeployChainSession) Owner() (common.Address, error) {
	return _DeployChain.Contract.Owner(&_DeployChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeployChain *DeployChainCallerSession) Owner() (common.Address, error) {
	return _DeployChain.Contract.Owner(&_DeployChain.CallOpts)
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

// ProxyAddress is a free data retrieval call binding the contract method 0x380cb000.
//
// Solidity: function proxyAddress(address proxy, bytes32 salt) view returns(address)
func (_DeployChain *DeployChainCaller) ProxyAddress(opts *bind.CallOpts, proxy common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "proxyAddress", proxy, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAddress is a free data retrieval call binding the contract method 0x380cb000.
//
// Solidity: function proxyAddress(address proxy, bytes32 salt) view returns(address)
func (_DeployChain *DeployChainSession) ProxyAddress(proxy common.Address, salt [32]byte) (common.Address, error) {
	return _DeployChain.Contract.ProxyAddress(&_DeployChain.CallOpts, proxy, salt)
}

// ProxyAddress is a free data retrieval call binding the contract method 0x380cb000.
//
// Solidity: function proxyAddress(address proxy, bytes32 salt) view returns(address)
func (_DeployChain *DeployChainCallerSession) ProxyAddress(proxy common.Address, salt [32]byte) (common.Address, error) {
	return _DeployChain.Contract.ProxyAddress(&_DeployChain.CallOpts, proxy, salt)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeployChain *DeployChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeployChain *DeployChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeployChain.Contract.RenounceOwnership(&_DeployChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeployChain *DeployChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeployChain.Contract.RenounceOwnership(&_DeployChain.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeployChain *DeployChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DeployChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeployChain *DeployChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeployChain.Contract.TransferOwnership(&_DeployChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeployChain *DeployChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeployChain.Contract.TransferOwnership(&_DeployChain.TransactOpts, newOwner)
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

// DeployChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DeployChain contract.
type DeployChainOwnershipTransferredIterator struct {
	Event *DeployChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeployChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployChainOwnershipTransferred)
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
		it.Event = new(DeployChainOwnershipTransferred)
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
func (it *DeployChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployChainOwnershipTransferred represents a OwnershipTransferred event raised by the DeployChain contract.
type DeployChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeployChain *DeployChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeployChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeployChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeployChainOwnershipTransferredIterator{contract: _DeployChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeployChain *DeployChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeployChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeployChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployChainOwnershipTransferred)
				if err := _DeployChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeployChain *DeployChainFilterer) ParseOwnershipTransferred(log types.Log) (*DeployChainOwnershipTransferred, error) {
	event := new(DeployChainOwnershipTransferred)
	if err := _DeployChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
