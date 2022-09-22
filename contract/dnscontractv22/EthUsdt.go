// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc22

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
}

// EthUsdtABI is the input ABI used to generate the binding from.
// Deprecated: Use EthUsdtMetaData.ABI instead.
var EthUsdtABI = EthUsdtMetaData.ABI

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
