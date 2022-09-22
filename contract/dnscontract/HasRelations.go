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

// HasRelationsMetaData contains all meta data concerning the HasRelations contract.
var HasRelationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HasRelationsABI is the input ABI used to generate the binding from.
// Deprecated: Use HasRelationsMetaData.ABI instead.
var HasRelationsABI = HasRelationsMetaData.ABI

// HasRelations is an auto generated Go binding around an Ethereum contract.
type HasRelations struct {
	HasRelationsCaller     // Read-only binding to the contract
	HasRelationsTransactor // Write-only binding to the contract
	HasRelationsFilterer   // Log filterer for contract events
}

// HasRelationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasRelationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasRelationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasRelationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasRelationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasRelationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasRelationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasRelationsSession struct {
	Contract     *HasRelations     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasRelationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasRelationsCallerSession struct {
	Contract *HasRelationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HasRelationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasRelationsTransactorSession struct {
	Contract     *HasRelationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HasRelationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasRelationsRaw struct {
	Contract *HasRelations // Generic contract binding to access the raw methods on
}

// HasRelationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasRelationsCallerRaw struct {
	Contract *HasRelationsCaller // Generic read-only contract binding to access the raw methods on
}

// HasRelationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasRelationsTransactorRaw struct {
	Contract *HasRelationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasRelations creates a new instance of HasRelations, bound to a specific deployed contract.
func NewHasRelations(address common.Address, backend bind.ContractBackend) (*HasRelations, error) {
	contract, err := bindHasRelations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasRelations{HasRelationsCaller: HasRelationsCaller{contract: contract}, HasRelationsTransactor: HasRelationsTransactor{contract: contract}, HasRelationsFilterer: HasRelationsFilterer{contract: contract}}, nil
}

// NewHasRelationsCaller creates a new read-only instance of HasRelations, bound to a specific deployed contract.
func NewHasRelationsCaller(address common.Address, caller bind.ContractCaller) (*HasRelationsCaller, error) {
	contract, err := bindHasRelations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasRelationsCaller{contract: contract}, nil
}

// NewHasRelationsTransactor creates a new write-only instance of HasRelations, bound to a specific deployed contract.
func NewHasRelationsTransactor(address common.Address, transactor bind.ContractTransactor) (*HasRelationsTransactor, error) {
	contract, err := bindHasRelations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasRelationsTransactor{contract: contract}, nil
}

// NewHasRelationsFilterer creates a new log filterer instance of HasRelations, bound to a specific deployed contract.
func NewHasRelationsFilterer(address common.Address, filterer bind.ContractFilterer) (*HasRelationsFilterer, error) {
	contract, err := bindHasRelations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasRelationsFilterer{contract: contract}, nil
}

// bindHasRelations binds a generic wrapper to an already deployed contract.
func bindHasRelations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasRelationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasRelations *HasRelationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HasRelations.Contract.HasRelationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasRelations *HasRelationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasRelations.Contract.HasRelationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasRelations *HasRelationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasRelations.Contract.HasRelationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasRelations *HasRelationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HasRelations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasRelations *HasRelationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasRelations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasRelations *HasRelationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasRelations.Contract.contract.Transact(opts, method, params...)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_HasRelations *HasRelationsCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HasRelations.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_HasRelations *HasRelationsSession) Rel() (common.Address, error) {
	return _HasRelations.Contract.Rel(&_HasRelations.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_HasRelations *HasRelationsCallerSession) Rel() (common.Address, error) {
	return _HasRelations.Contract.Rel(&_HasRelations.CallOpts)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_HasRelations *HasRelationsTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _HasRelations.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_HasRelations *HasRelationsSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _HasRelations.Contract.ChangeRelation(&_HasRelations.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_HasRelations *HasRelationsTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _HasRelations.Contract.ChangeRelation(&_HasRelations.TransactOpts, new_rel)
}
