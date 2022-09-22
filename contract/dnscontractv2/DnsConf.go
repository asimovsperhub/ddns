// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc

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

// DnsConfMetaData contains all meta data concerning the DnsConf contract.
var DnsConfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"EvAddMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tName\",\"type\":\"bytes\"}],\"name\":\"EvAddType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"EvUpdateMap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tName\",\"type\":\"bytes\"}],\"name\":\"RemoveMap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"tName_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"v_\",\"type\":\"bytes\"}],\"name\":\"addMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"tName_\",\"type\":\"bytes\"}],\"name\":\"addType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTypes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"tName_\",\"type\":\"bytes\"}],\"name\":\"getMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"tName_\",\"type\":\"bytes\"}],\"name\":\"removeMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferDaoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"tName_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"v_\",\"type\":\"bytes\"}],\"name\":\"updateMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DnsConfABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsConfMetaData.ABI instead.
var DnsConfABI = DnsConfMetaData.ABI

// DnsConf is an auto generated Go binding around an Ethereum contract.
type DnsConf struct {
	DnsConfCaller     // Read-only binding to the contract
	DnsConfTransactor // Write-only binding to the contract
	DnsConfFilterer   // Log filterer for contract events
}

// DnsConfCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsConfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsConfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsConfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsConfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsConfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsConfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsConfSession struct {
	Contract     *DnsConf          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsConfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsConfCallerSession struct {
	Contract *DnsConfCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DnsConfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsConfTransactorSession struct {
	Contract     *DnsConfTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DnsConfRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsConfRaw struct {
	Contract *DnsConf // Generic contract binding to access the raw methods on
}

// DnsConfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsConfCallerRaw struct {
	Contract *DnsConfCaller // Generic read-only contract binding to access the raw methods on
}

// DnsConfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsConfTransactorRaw struct {
	Contract *DnsConfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsConf creates a new instance of DnsConf, bound to a specific deployed contract.
func NewDnsConf(address common.Address, backend bind.ContractBackend) (*DnsConf, error) {
	contract, err := bindDnsConf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsConf{DnsConfCaller: DnsConfCaller{contract: contract}, DnsConfTransactor: DnsConfTransactor{contract: contract}, DnsConfFilterer: DnsConfFilterer{contract: contract}}, nil
}

// NewDnsConfCaller creates a new read-only instance of DnsConf, bound to a specific deployed contract.
func NewDnsConfCaller(address common.Address, caller bind.ContractCaller) (*DnsConfCaller, error) {
	contract, err := bindDnsConf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsConfCaller{contract: contract}, nil
}

// NewDnsConfTransactor creates a new write-only instance of DnsConf, bound to a specific deployed contract.
func NewDnsConfTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsConfTransactor, error) {
	contract, err := bindDnsConf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsConfTransactor{contract: contract}, nil
}

// NewDnsConfFilterer creates a new log filterer instance of DnsConf, bound to a specific deployed contract.
func NewDnsConfFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsConfFilterer, error) {
	contract, err := bindDnsConf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsConfFilterer{contract: contract}, nil
}

// bindDnsConf binds a generic wrapper to an already deployed contract.
func bindDnsConf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsConfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsConf *DnsConfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsConf.Contract.DnsConfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsConf *DnsConfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsConf.Contract.DnsConfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsConf *DnsConfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsConf.Contract.DnsConfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsConf *DnsConfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsConf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsConf *DnsConfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsConf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsConf *DnsConfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsConf.Contract.contract.Transact(opts, method, params...)
}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes)
func (_DnsConf *DnsConfCaller) GetAllTypes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DnsConf.contract.Call(opts, &out, "getAllTypes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes)
func (_DnsConf *DnsConfSession) GetAllTypes() ([]byte, error) {
	return _DnsConf.Contract.GetAllTypes(&_DnsConf.CallOpts)
}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes)
func (_DnsConf *DnsConfCallerSession) GetAllTypes() ([]byte, error) {
	return _DnsConf.Contract.GetAllTypes(&_DnsConf.CallOpts)
}

// GetMap is a free data retrieval call binding the contract method 0xf5c5e40d.
//
// Solidity: function getMap(bytes32 nameHash_, bytes tName_) view returns(bool, bytes)
func (_DnsConf *DnsConfCaller) GetMap(opts *bind.CallOpts, nameHash_ [32]byte, tName_ []byte) (bool, []byte, error) {
	var out []interface{}
	err := _DnsConf.contract.Call(opts, &out, "getMap", nameHash_, tName_)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetMap is a free data retrieval call binding the contract method 0xf5c5e40d.
//
// Solidity: function getMap(bytes32 nameHash_, bytes tName_) view returns(bool, bytes)
func (_DnsConf *DnsConfSession) GetMap(nameHash_ [32]byte, tName_ []byte) (bool, []byte, error) {
	return _DnsConf.Contract.GetMap(&_DnsConf.CallOpts, nameHash_, tName_)
}

// GetMap is a free data retrieval call binding the contract method 0xf5c5e40d.
//
// Solidity: function getMap(bytes32 nameHash_, bytes tName_) view returns(bool, bytes)
func (_DnsConf *DnsConfCallerSession) GetMap(nameHash_ [32]byte, tName_ []byte) (bool, []byte, error) {
	return _DnsConf.Contract.GetMap(&_DnsConf.CallOpts, nameHash_, tName_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsConf *DnsConfCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsConf.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsConf *DnsConfSession) Owner() (common.Address, error) {
	return _DnsConf.Contract.Owner(&_DnsConf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsConf *DnsConfCallerSession) Owner() (common.Address, error) {
	return _DnsConf.Contract.Owner(&_DnsConf.CallOpts)
}

// AddMap is a paid mutator transaction binding the contract method 0xdbba53cd.
//
// Solidity: function addMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfTransactor) AddMap(opts *bind.TransactOpts, nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.contract.Transact(opts, "addMap", nameHash_, tName_, v_)
}

// AddMap is a paid mutator transaction binding the contract method 0xdbba53cd.
//
// Solidity: function addMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfSession) AddMap(nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.AddMap(&_DnsConf.TransactOpts, nameHash_, tName_, v_)
}

// AddMap is a paid mutator transaction binding the contract method 0xdbba53cd.
//
// Solidity: function addMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfTransactorSession) AddMap(nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.AddMap(&_DnsConf.TransactOpts, nameHash_, tName_, v_)
}

// AddType is a paid mutator transaction binding the contract method 0xd5efae42.
//
// Solidity: function addType(bytes tName_) returns()
func (_DnsConf *DnsConfTransactor) AddType(opts *bind.TransactOpts, tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.contract.Transact(opts, "addType", tName_)
}

// AddType is a paid mutator transaction binding the contract method 0xd5efae42.
//
// Solidity: function addType(bytes tName_) returns()
func (_DnsConf *DnsConfSession) AddType(tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.AddType(&_DnsConf.TransactOpts, tName_)
}

// AddType is a paid mutator transaction binding the contract method 0xd5efae42.
//
// Solidity: function addType(bytes tName_) returns()
func (_DnsConf *DnsConfTransactorSession) AddType(tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.AddType(&_DnsConf.TransactOpts, tName_)
}

// RemoveMap is a paid mutator transaction binding the contract method 0x5e845629.
//
// Solidity: function removeMap(bytes32 nameHash_, bytes tName_) returns()
func (_DnsConf *DnsConfTransactor) RemoveMap(opts *bind.TransactOpts, nameHash_ [32]byte, tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.contract.Transact(opts, "removeMap", nameHash_, tName_)
}

// RemoveMap is a paid mutator transaction binding the contract method 0x5e845629.
//
// Solidity: function removeMap(bytes32 nameHash_, bytes tName_) returns()
func (_DnsConf *DnsConfSession) RemoveMap(nameHash_ [32]byte, tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.RemoveMap(&_DnsConf.TransactOpts, nameHash_, tName_)
}

// RemoveMap is a paid mutator transaction binding the contract method 0x5e845629.
//
// Solidity: function removeMap(bytes32 nameHash_, bytes tName_) returns()
func (_DnsConf *DnsConfTransactorSession) RemoveMap(nameHash_ [32]byte, tName_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.RemoveMap(&_DnsConf.TransactOpts, nameHash_, tName_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsConf *DnsConfTransactor) TransferDaoOwner(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsConf.contract.Transact(opts, "transferDaoOwner", newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsConf *DnsConfSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsConf.Contract.TransferDaoOwner(&_DnsConf.TransactOpts, newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsConf *DnsConfTransactorSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsConf.Contract.TransferDaoOwner(&_DnsConf.TransactOpts, newOwner_)
}

// UpdateMap is a paid mutator transaction binding the contract method 0xbafb7fa5.
//
// Solidity: function updateMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfTransactor) UpdateMap(opts *bind.TransactOpts, nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.contract.Transact(opts, "updateMap", nameHash_, tName_, v_)
}

// UpdateMap is a paid mutator transaction binding the contract method 0xbafb7fa5.
//
// Solidity: function updateMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfSession) UpdateMap(nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.UpdateMap(&_DnsConf.TransactOpts, nameHash_, tName_, v_)
}

// UpdateMap is a paid mutator transaction binding the contract method 0xbafb7fa5.
//
// Solidity: function updateMap(bytes32 nameHash_, bytes tName_, bytes v_) returns()
func (_DnsConf *DnsConfTransactorSession) UpdateMap(nameHash_ [32]byte, tName_ []byte, v_ []byte) (*types.Transaction, error) {
	return _DnsConf.Contract.UpdateMap(&_DnsConf.TransactOpts, nameHash_, tName_, v_)
}

// DnsConfEvAddMapIterator is returned from FilterEvAddMap and is used to iterate over the raw logs and unpacked data for EvAddMap events raised by the DnsConf contract.
type DnsConfEvAddMapIterator struct {
	Event *DnsConfEvAddMap // Event containing the contract specifics and raw log

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
func (it *DnsConfEvAddMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsConfEvAddMap)
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
		it.Event = new(DnsConfEvAddMap)
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
func (it *DnsConfEvAddMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsConfEvAddMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsConfEvAddMap represents a EvAddMap event raised by the DnsConf contract.
type DnsConfEvAddMap struct {
	NameHash [32]byte
	TName    []byte
	Val      []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvAddMap is a free log retrieval operation binding the contract event 0x17b631082163894d6881d8b1fb86f651fd750b9811dc14a0bd7b650141dae761.
//
// Solidity: event EvAddMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) FilterEvAddMap(opts *bind.FilterOpts) (*DnsConfEvAddMapIterator, error) {

	logs, sub, err := _DnsConf.contract.FilterLogs(opts, "EvAddMap")
	if err != nil {
		return nil, err
	}
	return &DnsConfEvAddMapIterator{contract: _DnsConf.contract, event: "EvAddMap", logs: logs, sub: sub}, nil
}

// WatchEvAddMap is a free log subscription operation binding the contract event 0x17b631082163894d6881d8b1fb86f651fd750b9811dc14a0bd7b650141dae761.
//
// Solidity: event EvAddMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) WatchEvAddMap(opts *bind.WatchOpts, sink chan<- *DnsConfEvAddMap) (event.Subscription, error) {

	logs, sub, err := _DnsConf.contract.WatchLogs(opts, "EvAddMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsConfEvAddMap)
				if err := _DnsConf.contract.UnpackLog(event, "EvAddMap", log); err != nil {
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

// ParseEvAddMap is a log parse operation binding the contract event 0x17b631082163894d6881d8b1fb86f651fd750b9811dc14a0bd7b650141dae761.
//
// Solidity: event EvAddMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) ParseEvAddMap(log types.Log) (*DnsConfEvAddMap, error) {
	event := new(DnsConfEvAddMap)
	if err := _DnsConf.contract.UnpackLog(event, "EvAddMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsConfEvAddTypeIterator is returned from FilterEvAddType and is used to iterate over the raw logs and unpacked data for EvAddType events raised by the DnsConf contract.
type DnsConfEvAddTypeIterator struct {
	Event *DnsConfEvAddType // Event containing the contract specifics and raw log

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
func (it *DnsConfEvAddTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsConfEvAddType)
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
		it.Event = new(DnsConfEvAddType)
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
func (it *DnsConfEvAddTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsConfEvAddTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsConfEvAddType represents a EvAddType event raised by the DnsConf contract.
type DnsConfEvAddType struct {
	TName []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEvAddType is a free log retrieval operation binding the contract event 0x9c917bc14c0f3464ccdc83f1b5b8e09c3a5a945b07198123f13e2e24967d661a.
//
// Solidity: event EvAddType(bytes tName)
func (_DnsConf *DnsConfFilterer) FilterEvAddType(opts *bind.FilterOpts) (*DnsConfEvAddTypeIterator, error) {

	logs, sub, err := _DnsConf.contract.FilterLogs(opts, "EvAddType")
	if err != nil {
		return nil, err
	}
	return &DnsConfEvAddTypeIterator{contract: _DnsConf.contract, event: "EvAddType", logs: logs, sub: sub}, nil
}

// WatchEvAddType is a free log subscription operation binding the contract event 0x9c917bc14c0f3464ccdc83f1b5b8e09c3a5a945b07198123f13e2e24967d661a.
//
// Solidity: event EvAddType(bytes tName)
func (_DnsConf *DnsConfFilterer) WatchEvAddType(opts *bind.WatchOpts, sink chan<- *DnsConfEvAddType) (event.Subscription, error) {

	logs, sub, err := _DnsConf.contract.WatchLogs(opts, "EvAddType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsConfEvAddType)
				if err := _DnsConf.contract.UnpackLog(event, "EvAddType", log); err != nil {
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

// ParseEvAddType is a log parse operation binding the contract event 0x9c917bc14c0f3464ccdc83f1b5b8e09c3a5a945b07198123f13e2e24967d661a.
//
// Solidity: event EvAddType(bytes tName)
func (_DnsConf *DnsConfFilterer) ParseEvAddType(log types.Log) (*DnsConfEvAddType, error) {
	event := new(DnsConfEvAddType)
	if err := _DnsConf.contract.UnpackLog(event, "EvAddType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsConfEvUpdateMapIterator is returned from FilterEvUpdateMap and is used to iterate over the raw logs and unpacked data for EvUpdateMap events raised by the DnsConf contract.
type DnsConfEvUpdateMapIterator struct {
	Event *DnsConfEvUpdateMap // Event containing the contract specifics and raw log

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
func (it *DnsConfEvUpdateMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsConfEvUpdateMap)
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
		it.Event = new(DnsConfEvUpdateMap)
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
func (it *DnsConfEvUpdateMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsConfEvUpdateMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsConfEvUpdateMap represents a EvUpdateMap event raised by the DnsConf contract.
type DnsConfEvUpdateMap struct {
	NameHash [32]byte
	TName    []byte
	Val      []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvUpdateMap is a free log retrieval operation binding the contract event 0xab4795f698e3711ce9fa243099c732d478d5e67b00916e70d335a7e00f779257.
//
// Solidity: event EvUpdateMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) FilterEvUpdateMap(opts *bind.FilterOpts) (*DnsConfEvUpdateMapIterator, error) {

	logs, sub, err := _DnsConf.contract.FilterLogs(opts, "EvUpdateMap")
	if err != nil {
		return nil, err
	}
	return &DnsConfEvUpdateMapIterator{contract: _DnsConf.contract, event: "EvUpdateMap", logs: logs, sub: sub}, nil
}

// WatchEvUpdateMap is a free log subscription operation binding the contract event 0xab4795f698e3711ce9fa243099c732d478d5e67b00916e70d335a7e00f779257.
//
// Solidity: event EvUpdateMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) WatchEvUpdateMap(opts *bind.WatchOpts, sink chan<- *DnsConfEvUpdateMap) (event.Subscription, error) {

	logs, sub, err := _DnsConf.contract.WatchLogs(opts, "EvUpdateMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsConfEvUpdateMap)
				if err := _DnsConf.contract.UnpackLog(event, "EvUpdateMap", log); err != nil {
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

// ParseEvUpdateMap is a log parse operation binding the contract event 0xab4795f698e3711ce9fa243099c732d478d5e67b00916e70d335a7e00f779257.
//
// Solidity: event EvUpdateMap(bytes32 nameHash, bytes tName, bytes val)
func (_DnsConf *DnsConfFilterer) ParseEvUpdateMap(log types.Log) (*DnsConfEvUpdateMap, error) {
	event := new(DnsConfEvUpdateMap)
	if err := _DnsConf.contract.UnpackLog(event, "EvUpdateMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsConfRemoveMapIterator is returned from FilterRemoveMap and is used to iterate over the raw logs and unpacked data for RemoveMap events raised by the DnsConf contract.
type DnsConfRemoveMapIterator struct {
	Event *DnsConfRemoveMap // Event containing the contract specifics and raw log

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
func (it *DnsConfRemoveMapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsConfRemoveMap)
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
		it.Event = new(DnsConfRemoveMap)
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
func (it *DnsConfRemoveMapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsConfRemoveMapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsConfRemoveMap represents a RemoveMap event raised by the DnsConf contract.
type DnsConfRemoveMap struct {
	NameHash [32]byte
	TName    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveMap is a free log retrieval operation binding the contract event 0x740990bd46b9ca68c748868acf65f139ea9d5d151c948bf497199e13d66a1732.
//
// Solidity: event RemoveMap(bytes32 nameHash, bytes tName)
func (_DnsConf *DnsConfFilterer) FilterRemoveMap(opts *bind.FilterOpts) (*DnsConfRemoveMapIterator, error) {

	logs, sub, err := _DnsConf.contract.FilterLogs(opts, "RemoveMap")
	if err != nil {
		return nil, err
	}
	return &DnsConfRemoveMapIterator{contract: _DnsConf.contract, event: "RemoveMap", logs: logs, sub: sub}, nil
}

// WatchRemoveMap is a free log subscription operation binding the contract event 0x740990bd46b9ca68c748868acf65f139ea9d5d151c948bf497199e13d66a1732.
//
// Solidity: event RemoveMap(bytes32 nameHash, bytes tName)
func (_DnsConf *DnsConfFilterer) WatchRemoveMap(opts *bind.WatchOpts, sink chan<- *DnsConfRemoveMap) (event.Subscription, error) {

	logs, sub, err := _DnsConf.contract.WatchLogs(opts, "RemoveMap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsConfRemoveMap)
				if err := _DnsConf.contract.UnpackLog(event, "RemoveMap", log); err != nil {
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

// ParseRemoveMap is a log parse operation binding the contract event 0x740990bd46b9ca68c748868acf65f139ea9d5d151c948bf497199e13d66a1732.
//
// Solidity: event RemoveMap(bytes32 nameHash, bytes tName)
func (_DnsConf *DnsConfFilterer) ParseRemoveMap(log types.Log) (*DnsConfRemoveMap, error) {
	event := new(DnsConfRemoveMap)
	if err := _DnsConf.contract.UnpackLog(event, "RemoveMap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
