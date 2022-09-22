// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnscontract

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

// BasSetMetaData contains all meta data concerning the BasSet contract.
var BasSetMetaData = &bind.MetaData{
	ABI: "[]",
}

// BasSetABI is the input ABI used to generate the binding from.
// Deprecated: Use BasSetMetaData.ABI instead.
var BasSetABI = BasSetMetaData.ABI

// BasSet is an auto generated Go binding around an Ethereum contract.
type BasSet struct {
	BasSetCaller     // Read-only binding to the contract
	BasSetTransactor // Write-only binding to the contract
	BasSetFilterer   // Log filterer for contract events
}

// BasSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasSetSession struct {
	Contract     *BasSet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasSetCallerSession struct {
	Contract *BasSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BasSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasSetTransactorSession struct {
	Contract     *BasSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasSetRaw struct {
	Contract *BasSet // Generic contract binding to access the raw methods on
}

// BasSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasSetCallerRaw struct {
	Contract *BasSetCaller // Generic read-only contract binding to access the raw methods on
}

// BasSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasSetTransactorRaw struct {
	Contract *BasSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasSet creates a new instance of BasSet, bound to a specific deployed contract.
func NewBasSet(address common.Address, backend bind.ContractBackend) (*BasSet, error) {
	contract, err := bindBasSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasSet{BasSetCaller: BasSetCaller{contract: contract}, BasSetTransactor: BasSetTransactor{contract: contract}, BasSetFilterer: BasSetFilterer{contract: contract}}, nil
}

// NewBasSetCaller creates a new read-only instance of BasSet, bound to a specific deployed contract.
func NewBasSetCaller(address common.Address, caller bind.ContractCaller) (*BasSetCaller, error) {
	contract, err := bindBasSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasSetCaller{contract: contract}, nil
}

// NewBasSetTransactor creates a new write-only instance of BasSet, bound to a specific deployed contract.
func NewBasSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BasSetTransactor, error) {
	contract, err := bindBasSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasSetTransactor{contract: contract}, nil
}

// NewBasSetFilterer creates a new log filterer instance of BasSet, bound to a specific deployed contract.
func NewBasSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BasSetFilterer, error) {
	contract, err := bindBasSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasSetFilterer{contract: contract}, nil
}

// bindBasSet binds a generic wrapper to an already deployed contract.
func bindBasSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasSet *BasSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasSet.Contract.BasSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasSet *BasSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasSet.Contract.BasSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasSet *BasSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasSet.Contract.BasSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasSet *BasSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasSet *BasSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasSet *BasSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasSet.Contract.contract.Transact(opts, method, params...)
}
