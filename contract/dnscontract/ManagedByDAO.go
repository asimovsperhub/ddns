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

// ManagedByDAOMetaData contains all meta data concerning the ManagedByDAO contract.
var ManagedByDAOMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ManagedByDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use ManagedByDAOMetaData.ABI instead.
var ManagedByDAOABI = ManagedByDAOMetaData.ABI

// ManagedByDAO is an auto generated Go binding around an Ethereum contract.
type ManagedByDAO struct {
	ManagedByDAOCaller     // Read-only binding to the contract
	ManagedByDAOTransactor // Write-only binding to the contract
	ManagedByDAOFilterer   // Log filterer for contract events
}

// ManagedByDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagedByDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagedByDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagedByDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagedByDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagedByDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagedByDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagedByDAOSession struct {
	Contract     *ManagedByDAO     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagedByDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagedByDAOCallerSession struct {
	Contract *ManagedByDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ManagedByDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagedByDAOTransactorSession struct {
	Contract     *ManagedByDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ManagedByDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagedByDAORaw struct {
	Contract *ManagedByDAO // Generic contract binding to access the raw methods on
}

// ManagedByDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagedByDAOCallerRaw struct {
	Contract *ManagedByDAOCaller // Generic read-only contract binding to access the raw methods on
}

// ManagedByDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagedByDAOTransactorRaw struct {
	Contract *ManagedByDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManagedByDAO creates a new instance of ManagedByDAO, bound to a specific deployed contract.
func NewManagedByDAO(address common.Address, backend bind.ContractBackend) (*ManagedByDAO, error) {
	contract, err := bindManagedByDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ManagedByDAO{ManagedByDAOCaller: ManagedByDAOCaller{contract: contract}, ManagedByDAOTransactor: ManagedByDAOTransactor{contract: contract}, ManagedByDAOFilterer: ManagedByDAOFilterer{contract: contract}}, nil
}

// NewManagedByDAOCaller creates a new read-only instance of ManagedByDAO, bound to a specific deployed contract.
func NewManagedByDAOCaller(address common.Address, caller bind.ContractCaller) (*ManagedByDAOCaller, error) {
	contract, err := bindManagedByDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagedByDAOCaller{contract: contract}, nil
}

// NewManagedByDAOTransactor creates a new write-only instance of ManagedByDAO, bound to a specific deployed contract.
func NewManagedByDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagedByDAOTransactor, error) {
	contract, err := bindManagedByDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagedByDAOTransactor{contract: contract}, nil
}

// NewManagedByDAOFilterer creates a new log filterer instance of ManagedByDAO, bound to a specific deployed contract.
func NewManagedByDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagedByDAOFilterer, error) {
	contract, err := bindManagedByDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagedByDAOFilterer{contract: contract}, nil
}

// bindManagedByDAO binds a generic wrapper to an already deployed contract.
func bindManagedByDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagedByDAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagedByDAO *ManagedByDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagedByDAO.Contract.ManagedByDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagedByDAO *ManagedByDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.ManagedByDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagedByDAO *ManagedByDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.ManagedByDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ManagedByDAO *ManagedByDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ManagedByDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ManagedByDAO *ManagedByDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ManagedByDAO *ManagedByDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_ManagedByDAO *ManagedByDAOCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ManagedByDAO.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_ManagedByDAO *ManagedByDAOSession) DAOAddress() (common.Address, error) {
	return _ManagedByDAO.Contract.DAOAddress(&_ManagedByDAO.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_ManagedByDAO *ManagedByDAOCallerSession) DAOAddress() (common.Address, error) {
	return _ManagedByDAO.Contract.DAOAddress(&_ManagedByDAO.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_ManagedByDAO *ManagedByDAOTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _ManagedByDAO.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_ManagedByDAO *ManagedByDAOSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.ChangeDAO(&_ManagedByDAO.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_ManagedByDAO *ManagedByDAOTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _ManagedByDAO.Contract.ChangeDAO(&_ManagedByDAO.TransactOpts, newDao)
}
