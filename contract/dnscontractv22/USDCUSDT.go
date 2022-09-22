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

// USDCUSDTMetaData contains all meta data concerning the USDCUSDT contract.
var USDCUSDTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"name\":\"insertIntoRoundInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"roundInfoMap\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// USDCUSDTABI is the input ABI used to generate the binding from.
// Deprecated: Use USDCUSDTMetaData.ABI instead.
var USDCUSDTABI = USDCUSDTMetaData.ABI

// USDCUSDT is an auto generated Go binding around an Ethereum contract.
type USDCUSDT struct {
	USDCUSDTCaller     // Read-only binding to the contract
	USDCUSDTTransactor // Write-only binding to the contract
	USDCUSDTFilterer   // Log filterer for contract events
}

// USDCUSDTCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDCUSDTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCUSDTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDCUSDTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCUSDTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDCUSDTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDCUSDTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDCUSDTSession struct {
	Contract     *USDCUSDT         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDCUSDTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDCUSDTCallerSession struct {
	Contract *USDCUSDTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// USDCUSDTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDCUSDTTransactorSession struct {
	Contract     *USDCUSDTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// USDCUSDTRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDCUSDTRaw struct {
	Contract *USDCUSDT // Generic contract binding to access the raw methods on
}

// USDCUSDTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDCUSDTCallerRaw struct {
	Contract *USDCUSDTCaller // Generic read-only contract binding to access the raw methods on
}

// USDCUSDTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDCUSDTTransactorRaw struct {
	Contract *USDCUSDTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDCUSDT creates a new instance of USDCUSDT, bound to a specific deployed contract.
func NewUSDCUSDT(address common.Address, backend bind.ContractBackend) (*USDCUSDT, error) {
	contract, err := bindUSDCUSDT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCUSDT{USDCUSDTCaller: USDCUSDTCaller{contract: contract}, USDCUSDTTransactor: USDCUSDTTransactor{contract: contract}, USDCUSDTFilterer: USDCUSDTFilterer{contract: contract}}, nil
}

// NewUSDCUSDTCaller creates a new read-only instance of USDCUSDT, bound to a specific deployed contract.
func NewUSDCUSDTCaller(address common.Address, caller bind.ContractCaller) (*USDCUSDTCaller, error) {
	contract, err := bindUSDCUSDT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCUSDTCaller{contract: contract}, nil
}

// NewUSDCUSDTTransactor creates a new write-only instance of USDCUSDT, bound to a specific deployed contract.
func NewUSDCUSDTTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCUSDTTransactor, error) {
	contract, err := bindUSDCUSDT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCUSDTTransactor{contract: contract}, nil
}

// NewUSDCUSDTFilterer creates a new log filterer instance of USDCUSDT, bound to a specific deployed contract.
func NewUSDCUSDTFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCUSDTFilterer, error) {
	contract, err := bindUSDCUSDT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCUSDTFilterer{contract: contract}, nil
}

// bindUSDCUSDT binds a generic wrapper to an already deployed contract.
func bindUSDCUSDT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDCUSDTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCUSDT *USDCUSDTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCUSDT.Contract.USDCUSDTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCUSDT *USDCUSDTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCUSDT.Contract.USDCUSDTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCUSDT *USDCUSDTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCUSDT.Contract.USDCUSDTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDCUSDT *USDCUSDTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCUSDT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDCUSDT *USDCUSDTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCUSDT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDCUSDT *USDCUSDTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCUSDT.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDCUSDT *USDCUSDTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDCUSDT *USDCUSDTSession) Decimals() (uint8, error) {
	return _USDCUSDT.Contract.Decimals(&_USDCUSDT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDCUSDT *USDCUSDTCallerSession) Decimals() (uint8, error) {
	return _USDCUSDT.Contract.Decimals(&_USDCUSDT.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDCUSDT *USDCUSDTCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDCUSDT *USDCUSDTSession) Description() (string, error) {
	return _USDCUSDT.Contract.Description(&_USDCUSDT.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDCUSDT *USDCUSDTCallerSession) Description() (string, error) {
	return _USDCUSDT.Contract.Description(&_USDCUSDT.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "getRoundData", _roundId)

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
func (_USDCUSDT *USDCUSDTSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.GetRoundData(&_USDCUSDT.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.GetRoundData(&_USDCUSDT.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "latestRoundData")

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
func (_USDCUSDT *USDCUSDTSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.LatestRoundData(&_USDCUSDT.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.LatestRoundData(&_USDCUSDT.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_USDCUSDT *USDCUSDTCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_USDCUSDT *USDCUSDTSession) LatestRoundId() (*big.Int, error) {
	return _USDCUSDT.Contract.LatestRoundId(&_USDCUSDT.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_USDCUSDT *USDCUSDTCallerSession) LatestRoundId() (*big.Int, error) {
	return _USDCUSDT.Contract.LatestRoundId(&_USDCUSDT.CallOpts)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCaller) RoundInfoMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "roundInfoMap", arg0)

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
func (_USDCUSDT *USDCUSDTSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.RoundInfoMap(&_USDCUSDT.CallOpts, arg0)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_USDCUSDT *USDCUSDTCallerSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _USDCUSDT.Contract.RoundInfoMap(&_USDCUSDT.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_USDCUSDT *USDCUSDTCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCUSDT.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_USDCUSDT *USDCUSDTSession) Version() (*big.Int, error) {
	return _USDCUSDT.Contract.Version(&_USDCUSDT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_USDCUSDT *USDCUSDTCallerSession) Version() (*big.Int, error) {
	return _USDCUSDT.Contract.Version(&_USDCUSDT.CallOpts)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_USDCUSDT *USDCUSDTTransactor) InsertIntoRoundInfo(opts *bind.TransactOpts, roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _USDCUSDT.contract.Transact(opts, "insertIntoRoundInfo", roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_USDCUSDT *USDCUSDTSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _USDCUSDT.Contract.InsertIntoRoundInfo(&_USDCUSDT.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_USDCUSDT *USDCUSDTTransactorSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _USDCUSDT.Contract.InsertIntoRoundInfo(&_USDCUSDT.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}
