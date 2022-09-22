// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlink

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
)

// EthUsdtMetaData contains all meta data concerning the EthUsdt contract.
var EthUsdtMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"name\":\"insertIntoRoundInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"roundInfoMap\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060086000806101000a81548160ff021916908360ff1602179055506040518060400160405280600881526020017f6574682d7573647400000000000000000000000000000000000000000000000081525060019080519060200190610077929190610085565b506001600281905550610189565b82805461009190610128565b90600052602060002090601f0160209004810192826100b357600085556100fa565b82601f106100cc57805160ff19168380011785556100fa565b828001600101855582156100fa579182015b828111156100f95782518255916020019190600101906100de565b5b509050610107919061010b565b5090565b5b8082111561012457600081600090555060010161010c565b5090565b6000600282049050600182168061014057607f821691505b602082108114156101545761015361015a565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6109b9806101986000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637e22168b1161005b5780637e22168b146101055780639a6fc8f514610121578063c339d6d114610155578063feaf968c1461018957610088565b806311a8f4131461008d578063313ce567146100ab57806354fd4d50146100c95780637284e416146100e7575b600080fd5b6100956101ab565b6040516100a291906107b8565b60405180910390f35b6100b36101c7565b6040516100c09190610826565b60405180910390f35b6100d16101dd565b6040516100de919061079d565b60405180910390f35b6100ef6101e7565b6040516100fc919061077b565b60405180910390f35b61011f600480360381019061011a919061068b565b610279565b005b61013b6004803603810190610136919061065e565b61042d565b60405161014c9594939291906107d3565b60405180910390f35b61016f600480360381019061016a919061065e565b610451565b6040516101809594939291906107d3565b60405180910390f35b6101916104b3565b6040516101a29594939291906107d3565b60405180910390f35b600460009054906101000a900469ffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b6000600254905090565b6060600180546101f6906108c7565b80601f0160208091040260200160405190810160405280929190818152602001828054610222906108c7565b801561026f5780601f106102445761010080835404028352916020019161026f565b820191906000526020600020905b81548152906001019060200180831161025257829003601f168201915b5050505050905090565b600460009054906101000a900469ffffffffffffffffffff1669ffffffffffffffffffff168569ffffffffffffffffffff1614156102b657610426565b600460009054906101000a900469ffffffffffffffffffff1669ffffffffffffffffffff168569ffffffffffffffffffff16111561031c5784600460006101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff1602179055505b6040518060a001604052808669ffffffffffffffffffff1681526020018581526020018481526020018381526020018269ffffffffffffffffffff1681525060036000600460009054906101000a900469ffffffffffffffffffff1669ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff1602179055509050505b5050505050565b600080600080600061043e866104ed565b9450945094509450945091939590929450565b60036020528060005260406000206000915090508060000160009054906101000a900469ffffffffffffffffffff16908060010154908060020154908060030154908060040160009054906101000a900469ffffffffffffffffffff16905085565b60008060008060006104dc600460009054906101000a900469ffffffffffffffffffff166104ed565b945094509450945094509091929394565b6000806000806000600360008769ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900469ffffffffffffffffffff16600360008869ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060010154600360008969ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060020154600360008a69ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060030154600360008b69ffffffffffffffffffff1669ffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900469ffffffffffffffffffff169450945094509450945091939590929450565b60008135905061062e8161093e565b92915050565b60008135905061064381610955565b92915050565b6000813590506106588161096c565b92915050565b60006020828403121561067457610673610928565b5b600061068284828501610649565b91505092915050565b600080600080600060a086880312156106a7576106a6610928565b5b60006106b588828901610649565b95505060206106c68882890161061f565b94505060406106d788828901610634565b93505060606106e888828901610634565b92505060806106f988828901610649565b9150509295509295909350565b61070f8161085d565b82525050565b600061072082610841565b61072a818561084c565b935061073a818560208601610894565b6107438161092d565b840191505092915050565b61075781610867565b82525050565b6107668161087e565b82525050565b61077581610871565b82525050565b600060208201905081810360008301526107958184610715565b905092915050565b60006020820190506107b2600083018461074e565b92915050565b60006020820190506107cd600083018461075d565b92915050565b600060a0820190506107e8600083018861075d565b6107f56020830187610706565b610802604083018661074e565b61080f606083018561074e565b61081c608083018461075d565b9695505050505050565b600060208201905061083b600083018461076c565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b60005b838110156108b2578082015181840152602081019050610897565b838111156108c1576000848401525b50505050565b600060028204905060018216806108df57607f821691505b602082108114156108f3576108f26108f9565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b6109478161085d565b811461095257600080fd5b50565b61095e81610867565b811461096957600080fd5b50565b6109758161087e565b811461098057600080fd5b5056fea2646970667358221220ee2245efcfd1bdd852b686458b0ca882216838769977ba8c9721a36b14c2b29964736f6c63430008060033",
}

// EthUsdtABI is the input ABI used to generate the binding from.
// Deprecated: Use EthUsdtMetaData.ABI instead.
var EthUsdtABI = EthUsdtMetaData.ABI

// EthUsdtBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthUsdtMetaData.Bin instead.
var EthUsdtBin = EthUsdtMetaData.Bin

// DeployEthUsdt deploys a new Ethereum contract, binding an instance of EthUsdt to it.
func DeployEthUsdt(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthUsdt, error) {
	parsed, err := EthUsdtMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthUsdtBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthUsdt{EthUsdtCaller: EthUsdtCaller{contract: contract}, EthUsdtTransactor: EthUsdtTransactor{contract: contract}, EthUsdtFilterer: EthUsdtFilterer{contract: contract}}, nil
}

// EthUsdt is an auto generated Go binding around an Ethereum contract.
type EthUsdt struct {
	EthUsdtCaller     // Read-only binding to the contract
	EthUsdtTransactor // Write-only binding to the contract
	EthUsdtFilterer   // Log filterer for contract events
}

// EthUsdtCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthUsdtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthUsdtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthUsdtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthUsdtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthUsdtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthUsdtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthUsdtSession struct {
	Contract     *EthUsdt          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthUsdtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthUsdtCallerSession struct {
	Contract *EthUsdtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EthUsdtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthUsdtTransactorSession struct {
	Contract     *EthUsdtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EthUsdtRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthUsdtRaw struct {
	Contract *EthUsdt // Generic contract binding to access the raw methods on
}

// EthUsdtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthUsdtCallerRaw struct {
	Contract *EthUsdtCaller // Generic read-only contract binding to access the raw methods on
}

// EthUsdtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthUsdtTransactorRaw struct {
	Contract *EthUsdtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthUsdt creates a new instance of EthUsdt, bound to a specific deployed contract.
func NewEthUsdt(address common.Address, backend bind.ContractBackend) (*EthUsdt, error) {
	contract, err := bindEthUsdt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthUsdt{EthUsdtCaller: EthUsdtCaller{contract: contract}, EthUsdtTransactor: EthUsdtTransactor{contract: contract}, EthUsdtFilterer: EthUsdtFilterer{contract: contract}}, nil
}

// NewEthUsdtCaller creates a new read-only instance of EthUsdt, bound to a specific deployed contract.
func NewEthUsdtCaller(address common.Address, caller bind.ContractCaller) (*EthUsdtCaller, error) {
	contract, err := bindEthUsdt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthUsdtCaller{contract: contract}, nil
}

// NewEthUsdtTransactor creates a new write-only instance of EthUsdt, bound to a specific deployed contract.
func NewEthUsdtTransactor(address common.Address, transactor bind.ContractTransactor) (*EthUsdtTransactor, error) {
	contract, err := bindEthUsdt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthUsdtTransactor{contract: contract}, nil
}

// NewEthUsdtFilterer creates a new log filterer instance of EthUsdt, bound to a specific deployed contract.
func NewEthUsdtFilterer(address common.Address, filterer bind.ContractFilterer) (*EthUsdtFilterer, error) {
	contract, err := bindEthUsdt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthUsdtFilterer{contract: contract}, nil
}

// bindEthUsdt binds a generic wrapper to an already deployed contract.
func bindEthUsdt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthUsdtABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthUsdt *EthUsdtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthUsdt.Contract.EthUsdtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthUsdt *EthUsdtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthUsdt.Contract.EthUsdtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthUsdt *EthUsdtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthUsdt.Contract.EthUsdtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthUsdt *EthUsdtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthUsdt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthUsdt *EthUsdtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthUsdt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthUsdt *EthUsdtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthUsdt.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthUsdt *EthUsdtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthUsdt *EthUsdtSession) Decimals() (uint8, error) {
	return _EthUsdt.Contract.Decimals(&_EthUsdt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthUsdt *EthUsdtCallerSession) Decimals() (uint8, error) {
	return _EthUsdt.Contract.Decimals(&_EthUsdt.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthUsdt *EthUsdtCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthUsdt *EthUsdtSession) Description() (string, error) {
	return _EthUsdt.Contract.Description(&_EthUsdt.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthUsdt *EthUsdtCallerSession) Description() (string, error) {
	return _EthUsdt.Contract.Description(&_EthUsdt.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.GetRoundData(&_EthUsdt.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.GetRoundData(&_EthUsdt.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.LatestRoundData(&_EthUsdt.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.LatestRoundData(&_EthUsdt.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_EthUsdt *EthUsdtCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_EthUsdt *EthUsdtSession) LatestRoundId() (*big.Int, error) {
	return _EthUsdt.Contract.LatestRoundId(&_EthUsdt.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_EthUsdt *EthUsdtCallerSession) LatestRoundId() (*big.Int, error) {
	return _EthUsdt.Contract.LatestRoundId(&_EthUsdt.CallOpts)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCaller) RoundInfoMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "roundInfoMap", arg0)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.RoundInfoMap(&_EthUsdt.CallOpts, arg0)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthUsdt *EthUsdtCallerSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthUsdt.Contract.RoundInfoMap(&_EthUsdt.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthUsdt *EthUsdtCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthUsdt.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthUsdt *EthUsdtSession) Version() (*big.Int, error) {
	return _EthUsdt.Contract.Version(&_EthUsdt.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthUsdt *EthUsdtCallerSession) Version() (*big.Int, error) {
	return _EthUsdt.Contract.Version(&_EthUsdt.CallOpts)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_EthUsdt *EthUsdtTransactor) InsertIntoRoundInfo(opts *bind.TransactOpts, roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _EthUsdt.contract.Transact(opts, "insertIntoRoundInfo", roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_EthUsdt *EthUsdtSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _EthUsdt.Contract.InsertIntoRoundInfo(&_EthUsdt.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_EthUsdt *EthUsdtTransactorSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _EthUsdt.Contract.InsertIntoRoundInfo(&_EthUsdt.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}