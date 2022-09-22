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

// BasDomainConfMetaData contains all meta data concerning the BasDomainConf contract.
var BasDomainConfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DomainConfChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TypeNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"typName\",\"type\":\"string\"}],\"name\":\"addNewType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"clearByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"otherConf\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashArray\",\"type\":\"bytes32[]\"}],\"name\":\"dataMigrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"domainConfData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"typName\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"bca\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"typName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"updateByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasDomainConfABI is the input ABI used to generate the binding from.
// Deprecated: Use BasDomainConfMetaData.ABI instead.
var BasDomainConfABI = BasDomainConfMetaData.ABI

// BasDomainConf is an auto generated Go binding around an Ethereum contract.
type BasDomainConf struct {
	BasDomainConfCaller     // Read-only binding to the contract
	BasDomainConfTransactor // Write-only binding to the contract
	BasDomainConfFilterer   // Log filterer for contract events
}

// BasDomainConfCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasDomainConfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainConfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasDomainConfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainConfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasDomainConfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainConfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasDomainConfSession struct {
	Contract     *BasDomainConf    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasDomainConfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasDomainConfCallerSession struct {
	Contract *BasDomainConfCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BasDomainConfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasDomainConfTransactorSession struct {
	Contract     *BasDomainConfTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BasDomainConfRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasDomainConfRaw struct {
	Contract *BasDomainConf // Generic contract binding to access the raw methods on
}

// BasDomainConfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasDomainConfCallerRaw struct {
	Contract *BasDomainConfCaller // Generic read-only contract binding to access the raw methods on
}

// BasDomainConfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasDomainConfTransactorRaw struct {
	Contract *BasDomainConfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasDomainConf creates a new instance of BasDomainConf, bound to a specific deployed contract.
func NewBasDomainConf(address common.Address, backend bind.ContractBackend) (*BasDomainConf, error) {
	contract, err := bindBasDomainConf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasDomainConf{BasDomainConfCaller: BasDomainConfCaller{contract: contract}, BasDomainConfTransactor: BasDomainConfTransactor{contract: contract}, BasDomainConfFilterer: BasDomainConfFilterer{contract: contract}}, nil
}

// NewBasDomainConfCaller creates a new read-only instance of BasDomainConf, bound to a specific deployed contract.
func NewBasDomainConfCaller(address common.Address, caller bind.ContractCaller) (*BasDomainConfCaller, error) {
	contract, err := bindBasDomainConf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasDomainConfCaller{contract: contract}, nil
}

// NewBasDomainConfTransactor creates a new write-only instance of BasDomainConf, bound to a specific deployed contract.
func NewBasDomainConfTransactor(address common.Address, transactor bind.ContractTransactor) (*BasDomainConfTransactor, error) {
	contract, err := bindBasDomainConf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasDomainConfTransactor{contract: contract}, nil
}

// NewBasDomainConfFilterer creates a new log filterer instance of BasDomainConf, bound to a specific deployed contract.
func NewBasDomainConfFilterer(address common.Address, filterer bind.ContractFilterer) (*BasDomainConfFilterer, error) {
	contract, err := bindBasDomainConf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasDomainConfFilterer{contract: contract}, nil
}

// bindBasDomainConf binds a generic wrapper to an already deployed contract.
func bindBasDomainConf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasDomainConfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDomainConf *BasDomainConfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasDomainConf.Contract.BasDomainConfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDomainConf *BasDomainConfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDomainConf.Contract.BasDomainConfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDomainConf *BasDomainConfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDomainConf.Contract.BasDomainConfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDomainConf *BasDomainConfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasDomainConf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDomainConf *BasDomainConfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDomainConf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDomainConf *BasDomainConfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDomainConf.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomainConf *BasDomainConfCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasDomainConf.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomainConf *BasDomainConfSession) DAOAddress() (common.Address, error) {
	return _BasDomainConf.Contract.DAOAddress(&_BasDomainConf.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomainConf *BasDomainConfCallerSession) DAOAddress() (common.Address, error) {
	return _BasDomainConf.Contract.DAOAddress(&_BasDomainConf.CallOpts)
}

// TypeNames is a free data retrieval call binding the contract method 0xd3c7f131.
//
// Solidity: function TypeNames(uint256 ) view returns(string)
func (_BasDomainConf *BasDomainConfCaller) TypeNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _BasDomainConf.contract.Call(opts, &out, "TypeNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeNames is a free data retrieval call binding the contract method 0xd3c7f131.
//
// Solidity: function TypeNames(uint256 ) view returns(string)
func (_BasDomainConf *BasDomainConfSession) TypeNames(arg0 *big.Int) (string, error) {
	return _BasDomainConf.Contract.TypeNames(&_BasDomainConf.CallOpts, arg0)
}

// TypeNames is a free data retrieval call binding the contract method 0xd3c7f131.
//
// Solidity: function TypeNames(uint256 ) view returns(string)
func (_BasDomainConf *BasDomainConfCallerSession) TypeNames(arg0 *big.Int) (string, error) {
	return _BasDomainConf.Contract.TypeNames(&_BasDomainConf.CallOpts, arg0)
}

// DomainConfData is a free data retrieval call binding the contract method 0xe3db3d7a.
//
// Solidity: function domainConfData(bytes32 , string ) view returns(bytes)
func (_BasDomainConf *BasDomainConfCaller) DomainConfData(opts *bind.CallOpts, arg0 [32]byte, arg1 string) ([]byte, error) {
	var out []interface{}
	err := _BasDomainConf.contract.Call(opts, &out, "domainConfData", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DomainConfData is a free data retrieval call binding the contract method 0xe3db3d7a.
//
// Solidity: function domainConfData(bytes32 , string ) view returns(bytes)
func (_BasDomainConf *BasDomainConfSession) DomainConfData(arg0 [32]byte, arg1 string) ([]byte, error) {
	return _BasDomainConf.Contract.DomainConfData(&_BasDomainConf.CallOpts, arg0, arg1)
}

// DomainConfData is a free data retrieval call binding the contract method 0xe3db3d7a.
//
// Solidity: function domainConfData(bytes32 , string ) view returns(bytes)
func (_BasDomainConf *BasDomainConfCallerSession) DomainConfData(arg0 [32]byte, arg1 string) ([]byte, error) {
	return _BasDomainConf.Contract.DomainConfData(&_BasDomainConf.CallOpts, arg0, arg1)
}

// Query is a free data retrieval call binding the contract method 0x2eef3d65.
//
// Solidity: function query(bytes32 nameHash, string typName) view returns(bytes bca)
func (_BasDomainConf *BasDomainConfCaller) Query(opts *bind.CallOpts, nameHash [32]byte, typName string) ([]byte, error) {
	var out []interface{}
	err := _BasDomainConf.contract.Call(opts, &out, "query", nameHash, typName)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Query is a free data retrieval call binding the contract method 0x2eef3d65.
//
// Solidity: function query(bytes32 nameHash, string typName) view returns(bytes bca)
func (_BasDomainConf *BasDomainConfSession) Query(nameHash [32]byte, typName string) ([]byte, error) {
	return _BasDomainConf.Contract.Query(&_BasDomainConf.CallOpts, nameHash, typName)
}

// Query is a free data retrieval call binding the contract method 0x2eef3d65.
//
// Solidity: function query(bytes32 nameHash, string typName) view returns(bytes bca)
func (_BasDomainConf *BasDomainConfCallerSession) Query(nameHash [32]byte, typName string) ([]byte, error) {
	return _BasDomainConf.Contract.Query(&_BasDomainConf.CallOpts, nameHash, typName)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomainConf *BasDomainConfCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasDomainConf.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomainConf *BasDomainConfSession) Rel() (common.Address, error) {
	return _BasDomainConf.Contract.Rel(&_BasDomainConf.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomainConf *BasDomainConfCallerSession) Rel() (common.Address, error) {
	return _BasDomainConf.Contract.Rel(&_BasDomainConf.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomainConf *BasDomainConfTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomainConf *BasDomainConfSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ChangeDAO(&_BasDomainConf.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ChangeDAO(&_BasDomainConf.TransactOpts, newDao)
}

// AddNewType is a paid mutator transaction binding the contract method 0x268c876d.
//
// Solidity: function addNewType(string typName) returns()
func (_BasDomainConf *BasDomainConfTransactor) AddNewType(opts *bind.TransactOpts, typName string) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "addNewType", typName)
}

// AddNewType is a paid mutator transaction binding the contract method 0x268c876d.
//
// Solidity: function addNewType(string typName) returns()
func (_BasDomainConf *BasDomainConfSession) AddNewType(typName string) (*types.Transaction, error) {
	return _BasDomainConf.Contract.AddNewType(&_BasDomainConf.TransactOpts, typName)
}

// AddNewType is a paid mutator transaction binding the contract method 0x268c876d.
//
// Solidity: function addNewType(string typName) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) AddNewType(typName string) (*types.Transaction, error) {
	return _BasDomainConf.Contract.AddNewType(&_BasDomainConf.TransactOpts, typName)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomainConf *BasDomainConfTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomainConf *BasDomainConfSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ChangeRelation(&_BasDomainConf.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ChangeRelation(&_BasDomainConf.TransactOpts, new_rel)
}

// ClearByOwner is a paid mutator transaction binding the contract method 0x6ba92f55.
//
// Solidity: function clearByOwner(bytes32 nameHash) returns()
func (_BasDomainConf *BasDomainConfTransactor) ClearByOwner(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "clearByOwner", nameHash)
}

// ClearByOwner is a paid mutator transaction binding the contract method 0x6ba92f55.
//
// Solidity: function clearByOwner(bytes32 nameHash) returns()
func (_BasDomainConf *BasDomainConfSession) ClearByOwner(nameHash [32]byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ClearByOwner(&_BasDomainConf.TransactOpts, nameHash)
}

// ClearByOwner is a paid mutator transaction binding the contract method 0x6ba92f55.
//
// Solidity: function clearByOwner(bytes32 nameHash) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) ClearByOwner(nameHash [32]byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.ClearByOwner(&_BasDomainConf.TransactOpts, nameHash)
}

// DataMigrate is a paid mutator transaction binding the contract method 0x8ceb33d3.
//
// Solidity: function dataMigrate(address otherConf, bytes32[] hashArray) returns()
func (_BasDomainConf *BasDomainConfTransactor) DataMigrate(opts *bind.TransactOpts, otherConf common.Address, hashArray [][32]byte) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "dataMigrate", otherConf, hashArray)
}

// DataMigrate is a paid mutator transaction binding the contract method 0x8ceb33d3.
//
// Solidity: function dataMigrate(address otherConf, bytes32[] hashArray) returns()
func (_BasDomainConf *BasDomainConfSession) DataMigrate(otherConf common.Address, hashArray [][32]byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.DataMigrate(&_BasDomainConf.TransactOpts, otherConf, hashArray)
}

// DataMigrate is a paid mutator transaction binding the contract method 0x8ceb33d3.
//
// Solidity: function dataMigrate(address otherConf, bytes32[] hashArray) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) DataMigrate(otherConf common.Address, hashArray [][32]byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.DataMigrate(&_BasDomainConf.TransactOpts, otherConf, hashArray)
}

// UpdateByOwner is a paid mutator transaction binding the contract method 0x7118a4b7.
//
// Solidity: function updateByOwner(bytes32 nameHash, string typName, bytes data) returns()
func (_BasDomainConf *BasDomainConfTransactor) UpdateByOwner(opts *bind.TransactOpts, nameHash [32]byte, typName string, data []byte) (*types.Transaction, error) {
	return _BasDomainConf.contract.Transact(opts, "updateByOwner", nameHash, typName, data)
}

// UpdateByOwner is a paid mutator transaction binding the contract method 0x7118a4b7.
//
// Solidity: function updateByOwner(bytes32 nameHash, string typName, bytes data) returns()
func (_BasDomainConf *BasDomainConfSession) UpdateByOwner(nameHash [32]byte, typName string, data []byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.UpdateByOwner(&_BasDomainConf.TransactOpts, nameHash, typName, data)
}

// UpdateByOwner is a paid mutator transaction binding the contract method 0x7118a4b7.
//
// Solidity: function updateByOwner(bytes32 nameHash, string typName, bytes data) returns()
func (_BasDomainConf *BasDomainConfTransactorSession) UpdateByOwner(nameHash [32]byte, typName string, data []byte) (*types.Transaction, error) {
	return _BasDomainConf.Contract.UpdateByOwner(&_BasDomainConf.TransactOpts, nameHash, typName, data)
}

// BasDomainConfDomainConfChangedIterator is returned from FilterDomainConfChanged and is used to iterate over the raw logs and unpacked data for DomainConfChanged events raised by the BasDomainConf contract.
type BasDomainConfDomainConfChangedIterator struct {
	Event *BasDomainConfDomainConfChanged // Event containing the contract specifics and raw log

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
func (it *BasDomainConfDomainConfChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasDomainConfDomainConfChanged)
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
		it.Event = new(BasDomainConfDomainConfChanged)
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
func (it *BasDomainConfDomainConfChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasDomainConfDomainConfChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasDomainConfDomainConfChanged represents a DomainConfChanged event raised by the BasDomainConf contract.
type BasDomainConfDomainConfChanged struct {
	NameHash [32]byte
	TypeName string
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDomainConfChanged is a free log retrieval operation binding the contract event 0xd259dc7a87c80cf605c627fe050932c20304ef17c55f23b331a24ef6b3fdb1de.
//
// Solidity: event DomainConfChanged(bytes32 nameHash, string typeName, bytes data)
func (_BasDomainConf *BasDomainConfFilterer) FilterDomainConfChanged(opts *bind.FilterOpts) (*BasDomainConfDomainConfChangedIterator, error) {

	logs, sub, err := _BasDomainConf.contract.FilterLogs(opts, "DomainConfChanged")
	if err != nil {
		return nil, err
	}
	return &BasDomainConfDomainConfChangedIterator{contract: _BasDomainConf.contract, event: "DomainConfChanged", logs: logs, sub: sub}, nil
}

// WatchDomainConfChanged is a free log subscription operation binding the contract event 0xd259dc7a87c80cf605c627fe050932c20304ef17c55f23b331a24ef6b3fdb1de.
//
// Solidity: event DomainConfChanged(bytes32 nameHash, string typeName, bytes data)
func (_BasDomainConf *BasDomainConfFilterer) WatchDomainConfChanged(opts *bind.WatchOpts, sink chan<- *BasDomainConfDomainConfChanged) (event.Subscription, error) {

	logs, sub, err := _BasDomainConf.contract.WatchLogs(opts, "DomainConfChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasDomainConfDomainConfChanged)
				if err := _BasDomainConf.contract.UnpackLog(event, "DomainConfChanged", log); err != nil {
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

// ParseDomainConfChanged is a log parse operation binding the contract event 0xd259dc7a87c80cf605c627fe050932c20304ef17c55f23b331a24ef6b3fdb1de.
//
// Solidity: event DomainConfChanged(bytes32 nameHash, string typeName, bytes data)
func (_BasDomainConf *BasDomainConfFilterer) ParseDomainConfChanged(log types.Log) (*BasDomainConfDomainConfChanged, error) {
	event := new(BasDomainConfDomainConfChanged)
	if err := _BasDomainConf.contract.UnpackLog(event, "DomainConfChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
