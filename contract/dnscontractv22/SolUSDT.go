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

// SolUSDTMetaData contains all meta data concerning the SolUSDT contract.
var SolUSDTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"name\":\"insertIntoRoundInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"roundInfoMap\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SolUSDTABI is the input ABI used to generate the binding from.
// Deprecated: Use SolUSDTMetaData.ABI instead.
var SolUSDTABI = SolUSDTMetaData.ABI

// SolUSDT is an auto generated Go binding around an Ethereum contract.
type SolUSDT struct {
	SolUSDTCaller     // Read-only binding to the contract
	SolUSDTTransactor // Write-only binding to the contract
	SolUSDTFilterer   // Log filterer for contract events
}

// SolUSDTCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolUSDTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolUSDTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolUSDTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolUSDTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolUSDTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolUSDTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolUSDTSession struct {
	Contract     *SolUSDT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolUSDTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolUSDTCallerSession struct {
	Contract *SolUSDTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SolUSDTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolUSDTTransactorSession struct {
	Contract     *SolUSDTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SolUSDTRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolUSDTRaw struct {
	Contract *SolUSDT // Generic contract binding to access the raw methods on
}

// SolUSDTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolUSDTCallerRaw struct {
	Contract *SolUSDTCaller // Generic read-only contract binding to access the raw methods on
}

// SolUSDTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolUSDTTransactorRaw struct {
	Contract *SolUSDTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolUSDT creates a new instance of SolUSDT, bound to a specific deployed contract.
func NewSolUSDT(address common.Address, backend bind.ContractBackend) (*SolUSDT, error) {
	contract, err := bindSolUSDT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolUSDT{SolUSDTCaller: SolUSDTCaller{contract: contract}, SolUSDTTransactor: SolUSDTTransactor{contract: contract}, SolUSDTFilterer: SolUSDTFilterer{contract: contract}}, nil
}

// NewSolUSDTCaller creates a new read-only instance of SolUSDT, bound to a specific deployed contract.
func NewSolUSDTCaller(address common.Address, caller bind.ContractCaller) (*SolUSDTCaller, error) {
	contract, err := bindSolUSDT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolUSDTCaller{contract: contract}, nil
}

// NewSolUSDTTransactor creates a new write-only instance of SolUSDT, bound to a specific deployed contract.
func NewSolUSDTTransactor(address common.Address, transactor bind.ContractTransactor) (*SolUSDTTransactor, error) {
	contract, err := bindSolUSDT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolUSDTTransactor{contract: contract}, nil
}

// NewSolUSDTFilterer creates a new log filterer instance of SolUSDT, bound to a specific deployed contract.
func NewSolUSDTFilterer(address common.Address, filterer bind.ContractFilterer) (*SolUSDTFilterer, error) {
	contract, err := bindSolUSDT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolUSDTFilterer{contract: contract}, nil
}

// bindSolUSDT binds a generic wrapper to an already deployed contract.
func bindSolUSDT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolUSDTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolUSDT *SolUSDTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolUSDT.Contract.SolUSDTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolUSDT *SolUSDTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolUSDT.Contract.SolUSDTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolUSDT *SolUSDTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolUSDT.Contract.SolUSDTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolUSDT *SolUSDTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolUSDT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolUSDT *SolUSDTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolUSDT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolUSDT *SolUSDTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolUSDT.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolUSDT *SolUSDTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolUSDT *SolUSDTSession) Decimals() (uint8, error) {
	return _SolUSDT.Contract.Decimals(&_SolUSDT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolUSDT *SolUSDTCallerSession) Decimals() (uint8, error) {
	return _SolUSDT.Contract.Decimals(&_SolUSDT.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_SolUSDT *SolUSDTCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_SolUSDT *SolUSDTSession) Description() (string, error) {
	return _SolUSDT.Contract.Description(&_SolUSDT.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_SolUSDT *SolUSDTCallerSession) Description() (string, error) {
	return _SolUSDT.Contract.Description(&_SolUSDT.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "getRoundData", _roundId)

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
func (_SolUSDT *SolUSDTSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.GetRoundData(&_SolUSDT.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.GetRoundData(&_SolUSDT.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "latestRoundData")

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
func (_SolUSDT *SolUSDTSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.LatestRoundData(&_SolUSDT.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.LatestRoundData(&_SolUSDT.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_SolUSDT *SolUSDTCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_SolUSDT *SolUSDTSession) LatestRoundId() (*big.Int, error) {
	return _SolUSDT.Contract.LatestRoundId(&_SolUSDT.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_SolUSDT *SolUSDTCallerSession) LatestRoundId() (*big.Int, error) {
	return _SolUSDT.Contract.LatestRoundId(&_SolUSDT.CallOpts)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCaller) RoundInfoMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "roundInfoMap", arg0)

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
func (_SolUSDT *SolUSDTSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.RoundInfoMap(&_SolUSDT.CallOpts, arg0)
}

// RoundInfoMap is a free data retrieval call binding the contract method 0xc339d6d1.
//
// Solidity: function roundInfoMap(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_SolUSDT *SolUSDTCallerSession) RoundInfoMap(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _SolUSDT.Contract.RoundInfoMap(&_SolUSDT.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_SolUSDT *SolUSDTCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolUSDT.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_SolUSDT *SolUSDTSession) Version() (*big.Int, error) {
	return _SolUSDT.Contract.Version(&_SolUSDT.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_SolUSDT *SolUSDTCallerSession) Version() (*big.Int, error) {
	return _SolUSDT.Contract.Version(&_SolUSDT.CallOpts)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_SolUSDT *SolUSDTTransactor) InsertIntoRoundInfo(opts *bind.TransactOpts, roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _SolUSDT.contract.Transact(opts, "insertIntoRoundInfo", roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_SolUSDT *SolUSDTSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _SolUSDT.Contract.InsertIntoRoundInfo(&_SolUSDT.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}

// InsertIntoRoundInfo is a paid mutator transaction binding the contract method 0x7e22168b.
//
// Solidity: function insertIntoRoundInfo(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) returns()
func (_SolUSDT *SolUSDTTransactorSession) InsertIntoRoundInfo(roundId *big.Int, answer *big.Int, startedAt *big.Int, updatedAt *big.Int, answeredInRound *big.Int) (*types.Transaction, error) {
	return _SolUSDT.Contract.InsertIntoRoundInfo(&_SolUSDT.TransactOpts, roundId, answer, startedAt, updatedAt, answeredInRound)
}
