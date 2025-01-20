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

// TypesOutputProposal is an auto generated low-level Go binding around a user-defined struct.
type TypesOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// OutputOracleMetaData contains all meta data concerning the OutputOracle contract.
var OutputOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_systemConfigGlobal\",\"type\":\"address\",\"internalType\":\"contractSystemConfigGlobal\"},{\"name\":\"_maxOutputCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableProofs\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getL2Output\",\"inputs\":[{\"name\":\"_l2OutputIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputProposal\",\"components\":[{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2OutputAfter\",\"inputs\":[{\"name\":\"_l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputProposal\",\"components\":[{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL2OutputIndexAfter\",\"inputs\":[{\"name\":\"_l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"contractSystemConfigOwnable\"},{\"name\":\"_configHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_genesisOutputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_proofsEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestL2Output\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputProposal\",\"components\":[{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestOutputIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOutputCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextOutputIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofsEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeL2Output\",\"inputs\":[{\"name\":\"_outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l1BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSystemConfigOwnable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfigGlobal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSystemConfigGlobal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutputProposed\",\"inputs\":[{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"l2OutputIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"l1Timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001b9338038062001b93833981016040819052620000349162000266565b6001600160a01b03821660805260a08190526200005560008080806200005d565b5050620002a2565b600054610100900460ff16158080156200007e5750600054600160ff909116105b80620000ae57506200009b306200025760201b620011e61760201c565b158015620000ae575060005460ff166001145b620001165760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff1916600117905580156200013a576000805461ff0019166101001790555b600080546001600160a01b038716620100000262010000600160b01b03199091161781556001858155604080516060810182528681526001600160801b0342811660208301908152928201858152600280549586018155958690529151939094027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810193909355905190518316600160801b029216919091177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf909101556003805483151560ff19909116179055801562000250576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6001600160a01b03163b151590565b600080604083850312156200027a57600080fd5b82516001600160a01b03811681146200029257600080fd5b6020939093015192949293505050565b60805160a0516118af620002e4600039600081816102a40152818161046e0152610b2b0152600081816102cb015281816109db0152610dcb01526118af6000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80639ad84880116100b2578063c885bbb611610081578063cd92b3fe11610066578063cd92b3fe146102c6578063cf8e5cf0146102ed578063e1f1176d1461030057600080fd5b8063c885bbb614610297578063cc23c3811461029f57600080fd5b80639ad848801461021a578063a25ae5571461022d578063a8e4fb901461027c578063b82051481461028457600080fd5b806360df09b2116100ee57806360df09b2146101ec57806369f16eec146101f65780636abcf563146101ff5780637f0064201461020757600080fd5b806325f881721461012057806333d7e2bd146101425780634599c7881461018d57806354fd4d50146101a3575b600080fd5b60035461012d9060ff1681565b60405190151581526020015b60405180910390f35b6000546101689062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610139565b610195610309565b604051908152602001610139565b6101df6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161013991906115da565b6101f4610376565b005b61019560045481565b61019561046a565b61019561021536600461164d565b6104ab565b6101f4610228366004611666565b610641565b61024061023b36600461164d565b610c5a565b60408051825181526020808401516fffffffffffffffffffffffffffffffff908116918301919091529282015190921690820152606001610139565b610168610cee565b6101f4610292366004611723565b610e65565b610240611119565b6101957f000000000000000000000000000000000000000000000000000000000000000081565b6101687f000000000000000000000000000000000000000000000000000000000000000081565b6102406102fb36600461164d565b6111ae565b61019560015481565b6002546000901561035c576002600454815481106103295761032961176d565b906000526020600020906002020160010160109054906101000a90046fffffffffffffffffffffffffffffffff1661035f565b60005b6fffffffffffffffffffffffffffffffff16905090565b61037e610cee565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4f75747075744f7261636c653a206f6e6c79207468652070726f706f7365722060448201527f616464726573732063616e20656e61626c652070726f6f66730000000000000060648201526084015b60405180910390fd5b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60007f0000000000000000000000000000000000000000000000000000000000000000600454600161049c919061179c565b6104a6919061180a565b905090565b60006104b5610309565b82111561056a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f70726f706f736564000000000000000000000000000000000000000000000000608482015260a401610434565b600061057461046a565b6002549091508190600090610589908361179c565b90505b8082101561062b57600060026105a2838561179c565b6105ac919061181e565b6002805491925087916105bf908461180a565b815481106105cf576105cf61176d565b600091825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1610156106215761061a81600161179c565b9250610625565b8091505b5061058c565b600254610638908361180a565b95945050505050565b610649610cee565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610703576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f75747075744f7261636c653a206f6e6c79207468652070726f706f7365722060448201527f616464726573732063616e2070726f706f7365206e6577206f757470757473006064820152608401610434565b61070b610309565b84116107bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f4f75747075744f7261636c653a20626c6f636b206e756d626572206d7573742060448201527f62652067726561746572207468616e2070726576696f75736c792070726f706f60648201527f73656420626c6f636b206e756d62657200000000000000000000000000000000608482015260a401610434565b8461084c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f4f75747075744f7261636c653a204c32206f75747075742070726f706f73616c60448201527f2063616e6e6f7420626520746865207a65726f206861736800000000000000006064820152608401610434565b60035460ff1615610ab2578240806108e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4f75747075744f7261636c653a20626c6f636b68617368206e6f74206176616960448201527f6c61626c650000000000000000000000000000000000000000000000000000006064820152608401610434565b60006002600454815481106108fd576108fd61176d565b6000918252602080832060029092029091015460015460408051938401919091528201859052606082018990526080820181905260a082018a905292506109939060c0016040516020818303038152906040528051906020012086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061120292505050565b6040517f6a73b00b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192507f000000000000000000000000000000000000000000000000000000000000000090911690636a73b00b90602401602060405180830381865afa158015610a24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a489190611832565b610aae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4f75747075744f7261636c653a20696e76616c6964207369676e6174757265006044820152606401610434565b5050505b610aba61046a565b600481905560405142815285919087907fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e29060200160405180910390a4604080516060810182528681526fffffffffffffffffffffffffffffffff42811660208301528616918101919091526002547f00000000000000000000000000000000000000000000000000000000000000001115610be857600280546001810182556000829052825191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810191909155602082015160408301516fffffffffffffffffffffffffffffffff908116700100000000000000000000000000000000029116177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf90910155610c52565b80600260045481548110610bfe57610bfe61176d565b60009182526020918290208351600290920201908155908201516040909201516fffffffffffffffffffffffffffffffff908116700100000000000000000000000000000000029216919091176001909101555b505050505050565b604080516060810182526000808252602082018190529181019190915260028281548110610c8a57610c8a61176d565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b60008054819062010000900473ffffffffffffffffffffffffffffffffffffffff16610d1b576000610dac565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a8e4fb906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dac9190611856565b905073ffffffffffffffffffffffffffffffffffffffff8116610e5d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a8e4fb906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e589190611856565b610e5f565b805b91505090565b600054610100900460ff1615808015610e855750600054600160ff909116105b80610e9f5750303b158015610e9f575060005460ff166001145b610f2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610434565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610f8957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6000805473ffffffffffffffffffffffffffffffffffffffff871662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff9091161781556001858155604080516060810182528681526fffffffffffffffffffffffffffffffff42811660208301908152928201858152600280549586018155958690529151939094027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810193909355905190518316700100000000000000000000000000000000029216919091177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf90910155600380548315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909116179055801561111257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b604080516060810182526000808252602082018190529181019190915260026004548154811061114b5761114b61176d565b600091825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff80821694840194909452700100000000000000000000000000000000900490921691810191909152919050565b604080516060810182526000808252602082018190529181019190915260026111d6836104ab565b81548110610c8a57610c8a61176d565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60008060006112118585611226565b9150915061121e8161126b565b509392505050565b600080825160410361125c5760208301516040840151606085015160001a611250878285856114c2565b94509450505050611264565b506000905060025b9250929050565b600081600481111561127f5761127f611873565b036112875750565b600181600481111561129b5761129b611873565b03611302576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610434565b600281600481111561131657611316611873565b0361137d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610434565b600381600481111561139157611391611873565b0361141e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610434565b600481600481111561143257611432611873565b036114bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610434565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114f957506000905060036115d1565b8460ff16601b1415801561151157508460ff16601c14155b1561152257506000905060046115d1565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611576573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166115ca576000600192509250506115d1565b9150600090505b94509492505050565b600060208083528351808285015260005b81811015611607578581018301518582016040015282016115eb565b81811115611619576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561165f57600080fd5b5035919050565b60008060008060006080868803121561167e57600080fd5b853594506020860135935060408601359250606086013567ffffffffffffffff808211156116ab57600080fd5b818801915088601f8301126116bf57600080fd5b8135818111156116ce57600080fd5b8960208285010111156116e057600080fd5b9699959850939650602001949392505050565b73ffffffffffffffffffffffffffffffffffffffff811681146114bf57600080fd5b80151581146114bf57600080fd5b6000806000806080858703121561173957600080fd5b8435611744816116f3565b93506020850135925060408501359150606085013561176281611715565b939692955090935050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082198211156117d6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611819576118196117db565b500690565b60008261182d5761182d6117db565b500490565b60006020828403121561184457600080fd5b815161184f81611715565b9392505050565b60006020828403121561186857600080fd5b815161184f816116f3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a",
}

// OutputOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OutputOracleMetaData.ABI instead.
var OutputOracleABI = OutputOracleMetaData.ABI

// OutputOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutputOracleMetaData.Bin instead.
var OutputOracleBin = OutputOracleMetaData.Bin

// DeployOutputOracle deploys a new Ethereum contract, binding an instance of OutputOracle to it.
func DeployOutputOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _systemConfigGlobal common.Address, _maxOutputCount *big.Int) (common.Address, *types.Transaction, *OutputOracle, error) {
	parsed, err := OutputOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutputOracleBin), backend, _systemConfigGlobal, _maxOutputCount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutputOracle{OutputOracleCaller: OutputOracleCaller{contract: contract}, OutputOracleTransactor: OutputOracleTransactor{contract: contract}, OutputOracleFilterer: OutputOracleFilterer{contract: contract}}, nil
}

// OutputOracle is an auto generated Go binding around an Ethereum contract.
type OutputOracle struct {
	OutputOracleCaller     // Read-only binding to the contract
	OutputOracleTransactor // Write-only binding to the contract
	OutputOracleFilterer   // Log filterer for contract events
}

// OutputOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutputOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutputOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutputOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutputOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutputOracleSession struct {
	Contract     *OutputOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutputOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutputOracleCallerSession struct {
	Contract *OutputOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OutputOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutputOracleTransactorSession struct {
	Contract     *OutputOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OutputOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutputOracleRaw struct {
	Contract *OutputOracle // Generic contract binding to access the raw methods on
}

// OutputOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutputOracleCallerRaw struct {
	Contract *OutputOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OutputOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutputOracleTransactorRaw struct {
	Contract *OutputOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutputOracle creates a new instance of OutputOracle, bound to a specific deployed contract.
func NewOutputOracle(address common.Address, backend bind.ContractBackend) (*OutputOracle, error) {
	contract, err := bindOutputOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutputOracle{OutputOracleCaller: OutputOracleCaller{contract: contract}, OutputOracleTransactor: OutputOracleTransactor{contract: contract}, OutputOracleFilterer: OutputOracleFilterer{contract: contract}}, nil
}

// NewOutputOracleCaller creates a new read-only instance of OutputOracle, bound to a specific deployed contract.
func NewOutputOracleCaller(address common.Address, caller bind.ContractCaller) (*OutputOracleCaller, error) {
	contract, err := bindOutputOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutputOracleCaller{contract: contract}, nil
}

// NewOutputOracleTransactor creates a new write-only instance of OutputOracle, bound to a specific deployed contract.
func NewOutputOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OutputOracleTransactor, error) {
	contract, err := bindOutputOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutputOracleTransactor{contract: contract}, nil
}

// NewOutputOracleFilterer creates a new log filterer instance of OutputOracle, bound to a specific deployed contract.
func NewOutputOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OutputOracleFilterer, error) {
	contract, err := bindOutputOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutputOracleFilterer{contract: contract}, nil
}

// bindOutputOracle binds a generic wrapper to an already deployed contract.
func bindOutputOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutputOracle *OutputOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutputOracle.Contract.OutputOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutputOracle *OutputOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutputOracle.Contract.OutputOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutputOracle *OutputOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutputOracle.Contract.OutputOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutputOracle *OutputOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutputOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutputOracle *OutputOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutputOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutputOracle *OutputOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutputOracle.Contract.contract.Transact(opts, method, params...)
}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(bytes32)
func (_OutputOracle *OutputOracleCaller) ConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "configHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(bytes32)
func (_OutputOracle *OutputOracleSession) ConfigHash() ([32]byte, error) {
	return _OutputOracle.Contract.ConfigHash(&_OutputOracle.CallOpts)
}

// ConfigHash is a free data retrieval call binding the contract method 0xe1f1176d.
//
// Solidity: function configHash() view returns(bytes32)
func (_OutputOracle *OutputOracleCallerSession) ConfigHash() ([32]byte, error) {
	return _OutputOracle.Contract.ConfigHash(&_OutputOracle.CallOpts)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OutputOracle.Contract.GetL2Output(&_OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCallerSession) GetL2Output(_l2OutputIndex *big.Int) (TypesOutputProposal, error) {
	return _OutputOracle.Contract.GetL2Output(&_OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCaller) GetL2OutputAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "getL2OutputAfter", _l2BlockNumber)

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OutputOracle.Contract.GetL2OutputAfter(&_OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCallerSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesOutputProposal, error) {
	return _OutputOracle.Contract.GetL2OutputAfter(&_OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OutputOracle *OutputOracleCaller) GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "getL2OutputIndexAfter", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OutputOracle *OutputOracleSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OutputOracle.Contract.GetL2OutputIndexAfter(&_OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_OutputOracle *OutputOracleCallerSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _OutputOracle.Contract.GetL2OutputIndexAfter(&_OutputOracle.CallOpts, _l2BlockNumber)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OutputOracle *OutputOracleCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OutputOracle *OutputOracleSession) LatestBlockNumber() (*big.Int, error) {
	return _OutputOracle.Contract.LatestBlockNumber(&_OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_OutputOracle *OutputOracleCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _OutputOracle.Contract.LatestBlockNumber(&_OutputOracle.CallOpts)
}

// LatestL2Output is a free data retrieval call binding the contract method 0xc885bbb6.
//
// Solidity: function latestL2Output() view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCaller) LatestL2Output(opts *bind.CallOpts) (TypesOutputProposal, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "latestL2Output")

	if err != nil {
		return *new(TypesOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesOutputProposal)).(*TypesOutputProposal)

	return out0, err

}

// LatestL2Output is a free data retrieval call binding the contract method 0xc885bbb6.
//
// Solidity: function latestL2Output() view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleSession) LatestL2Output() (TypesOutputProposal, error) {
	return _OutputOracle.Contract.LatestL2Output(&_OutputOracle.CallOpts)
}

// LatestL2Output is a free data retrieval call binding the contract method 0xc885bbb6.
//
// Solidity: function latestL2Output() view returns((bytes32,uint128,uint128))
func (_OutputOracle *OutputOracleCallerSession) LatestL2Output() (TypesOutputProposal, error) {
	return _OutputOracle.Contract.LatestL2Output(&_OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleSession) LatestOutputIndex() (*big.Int, error) {
	return _OutputOracle.Contract.LatestOutputIndex(&_OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _OutputOracle.Contract.LatestOutputIndex(&_OutputOracle.CallOpts)
}

// MaxOutputCount is a free data retrieval call binding the contract method 0xcc23c381.
//
// Solidity: function maxOutputCount() view returns(uint256)
func (_OutputOracle *OutputOracleCaller) MaxOutputCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "maxOutputCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOutputCount is a free data retrieval call binding the contract method 0xcc23c381.
//
// Solidity: function maxOutputCount() view returns(uint256)
func (_OutputOracle *OutputOracleSession) MaxOutputCount() (*big.Int, error) {
	return _OutputOracle.Contract.MaxOutputCount(&_OutputOracle.CallOpts)
}

// MaxOutputCount is a free data retrieval call binding the contract method 0xcc23c381.
//
// Solidity: function maxOutputCount() view returns(uint256)
func (_OutputOracle *OutputOracleCallerSession) MaxOutputCount() (*big.Int, error) {
	return _OutputOracle.Contract.MaxOutputCount(&_OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleSession) NextOutputIndex() (*big.Int, error) {
	return _OutputOracle.Contract.NextOutputIndex(&_OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_OutputOracle *OutputOracleCallerSession) NextOutputIndex() (*big.Int, error) {
	return _OutputOracle.Contract.NextOutputIndex(&_OutputOracle.CallOpts)
}

// ProofsEnabled is a free data retrieval call binding the contract method 0x25f88172.
//
// Solidity: function proofsEnabled() view returns(bool)
func (_OutputOracle *OutputOracleCaller) ProofsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "proofsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProofsEnabled is a free data retrieval call binding the contract method 0x25f88172.
//
// Solidity: function proofsEnabled() view returns(bool)
func (_OutputOracle *OutputOracleSession) ProofsEnabled() (bool, error) {
	return _OutputOracle.Contract.ProofsEnabled(&_OutputOracle.CallOpts)
}

// ProofsEnabled is a free data retrieval call binding the contract method 0x25f88172.
//
// Solidity: function proofsEnabled() view returns(bool)
func (_OutputOracle *OutputOracleCallerSession) ProofsEnabled() (bool, error) {
	return _OutputOracle.Contract.ProofsEnabled(&_OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OutputOracle *OutputOracleCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OutputOracle *OutputOracleSession) Proposer() (common.Address, error) {
	return _OutputOracle.Contract.Proposer(&_OutputOracle.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OutputOracle *OutputOracleCallerSession) Proposer() (common.Address, error) {
	return _OutputOracle.Contract.Proposer(&_OutputOracle.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_OutputOracle *OutputOracleCaller) SystemConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "systemConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_OutputOracle *OutputOracleSession) SystemConfig() (common.Address, error) {
	return _OutputOracle.Contract.SystemConfig(&_OutputOracle.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_OutputOracle *OutputOracleCallerSession) SystemConfig() (common.Address, error) {
	return _OutputOracle.Contract.SystemConfig(&_OutputOracle.CallOpts)
}

// SystemConfigGlobal is a free data retrieval call binding the contract method 0xcd92b3fe.
//
// Solidity: function systemConfigGlobal() view returns(address)
func (_OutputOracle *OutputOracleCaller) SystemConfigGlobal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "systemConfigGlobal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfigGlobal is a free data retrieval call binding the contract method 0xcd92b3fe.
//
// Solidity: function systemConfigGlobal() view returns(address)
func (_OutputOracle *OutputOracleSession) SystemConfigGlobal() (common.Address, error) {
	return _OutputOracle.Contract.SystemConfigGlobal(&_OutputOracle.CallOpts)
}

// SystemConfigGlobal is a free data retrieval call binding the contract method 0xcd92b3fe.
//
// Solidity: function systemConfigGlobal() view returns(address)
func (_OutputOracle *OutputOracleCallerSession) SystemConfigGlobal() (common.Address, error) {
	return _OutputOracle.Contract.SystemConfigGlobal(&_OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OutputOracle *OutputOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OutputOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OutputOracle *OutputOracleSession) Version() (string, error) {
	return _OutputOracle.Contract.Version(&_OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OutputOracle *OutputOracleCallerSession) Version() (string, error) {
	return _OutputOracle.Contract.Version(&_OutputOracle.CallOpts)
}

// EnableProofs is a paid mutator transaction binding the contract method 0x60df09b2.
//
// Solidity: function enableProofs() returns()
func (_OutputOracle *OutputOracleTransactor) EnableProofs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutputOracle.contract.Transact(opts, "enableProofs")
}

// EnableProofs is a paid mutator transaction binding the contract method 0x60df09b2.
//
// Solidity: function enableProofs() returns()
func (_OutputOracle *OutputOracleSession) EnableProofs() (*types.Transaction, error) {
	return _OutputOracle.Contract.EnableProofs(&_OutputOracle.TransactOpts)
}

// EnableProofs is a paid mutator transaction binding the contract method 0x60df09b2.
//
// Solidity: function enableProofs() returns()
func (_OutputOracle *OutputOracleTransactorSession) EnableProofs() (*types.Transaction, error) {
	return _OutputOracle.Contract.EnableProofs(&_OutputOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xb8205148.
//
// Solidity: function initialize(address _systemConfig, bytes32 _configHash, bytes32 _genesisOutputRoot, bool _proofsEnabled) returns()
func (_OutputOracle *OutputOracleTransactor) Initialize(opts *bind.TransactOpts, _systemConfig common.Address, _configHash [32]byte, _genesisOutputRoot [32]byte, _proofsEnabled bool) (*types.Transaction, error) {
	return _OutputOracle.contract.Transact(opts, "initialize", _systemConfig, _configHash, _genesisOutputRoot, _proofsEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xb8205148.
//
// Solidity: function initialize(address _systemConfig, bytes32 _configHash, bytes32 _genesisOutputRoot, bool _proofsEnabled) returns()
func (_OutputOracle *OutputOracleSession) Initialize(_systemConfig common.Address, _configHash [32]byte, _genesisOutputRoot [32]byte, _proofsEnabled bool) (*types.Transaction, error) {
	return _OutputOracle.Contract.Initialize(&_OutputOracle.TransactOpts, _systemConfig, _configHash, _genesisOutputRoot, _proofsEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xb8205148.
//
// Solidity: function initialize(address _systemConfig, bytes32 _configHash, bytes32 _genesisOutputRoot, bool _proofsEnabled) returns()
func (_OutputOracle *OutputOracleTransactorSession) Initialize(_systemConfig common.Address, _configHash [32]byte, _genesisOutputRoot [32]byte, _proofsEnabled bool) (*types.Transaction, error) {
	return _OutputOracle.Contract.Initialize(&_OutputOracle.TransactOpts, _systemConfig, _configHash, _genesisOutputRoot, _proofsEnabled)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9ad84880.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, uint256 _l1BlockNumber, bytes _signature) returns()
func (_OutputOracle *OutputOracleTransactor) ProposeL2Output(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockNumber *big.Int, _signature []byte) (*types.Transaction, error) {
	return _OutputOracle.contract.Transact(opts, "proposeL2Output", _outputRoot, _l2BlockNumber, _l1BlockNumber, _signature)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9ad84880.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, uint256 _l1BlockNumber, bytes _signature) returns()
func (_OutputOracle *OutputOracleSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockNumber *big.Int, _signature []byte) (*types.Transaction, error) {
	return _OutputOracle.Contract.ProposeL2Output(&_OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockNumber, _signature)
}

// ProposeL2Output is a paid mutator transaction binding the contract method 0x9ad84880.
//
// Solidity: function proposeL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, uint256 _l1BlockNumber, bytes _signature) returns()
func (_OutputOracle *OutputOracleTransactorSession) ProposeL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockNumber *big.Int, _signature []byte) (*types.Transaction, error) {
	return _OutputOracle.Contract.ProposeL2Output(&_OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockNumber, _signature)
}

// OutputOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OutputOracle contract.
type OutputOracleInitializedIterator struct {
	Event *OutputOracleInitialized // Event containing the contract specifics and raw log

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
func (it *OutputOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutputOracleInitialized)
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
		it.Event = new(OutputOracleInitialized)
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
func (it *OutputOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutputOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutputOracleInitialized represents a Initialized event raised by the OutputOracle contract.
type OutputOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OutputOracle *OutputOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*OutputOracleInitializedIterator, error) {

	logs, sub, err := _OutputOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OutputOracleInitializedIterator{contract: _OutputOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OutputOracle *OutputOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OutputOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _OutputOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutputOracleInitialized)
				if err := _OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OutputOracle *OutputOracleFilterer) ParseInitialized(log types.Log) (*OutputOracleInitialized, error) {
	event := new(OutputOracleInitialized)
	if err := _OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutputOracleOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the OutputOracle contract.
type OutputOracleOutputProposedIterator struct {
	Event *OutputOracleOutputProposed // Event containing the contract specifics and raw log

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
func (it *OutputOracleOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutputOracleOutputProposed)
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
		it.Event = new(OutputOracleOutputProposed)
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
func (it *OutputOracleOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutputOracleOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutputOracleOutputProposed represents a OutputProposed event raised by the OutputOracle contract.
type OutputOracleOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OutputOracle *OutputOracleFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*OutputOracleOutputProposedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OutputOracle.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &OutputOracleOutputProposedIterator{contract: _OutputOracle.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OutputOracle *OutputOracleFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *OutputOracleOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _OutputOracle.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutputOracleOutputProposed)
				if err := _OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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

// ParseOutputProposed is a log parse operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_OutputOracle *OutputOracleFilterer) ParseOutputProposed(log types.Log) (*OutputOracleOutputProposed, error) {
	event := new(OutputOracleOutputProposed)
	if err := _OutputOracle.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
