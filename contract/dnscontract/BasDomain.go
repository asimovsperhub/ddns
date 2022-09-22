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

// BasDomainMetaData contains all meta data concerning the BasDomain contract.
var BasDomainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rechargeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEnd\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BasDomainABI is the input ABI used to generate the binding from.
// Deprecated: Use BasDomainMetaData.ABI instead.
var BasDomainABI = BasDomainMetaData.ABI

// BasDomain is an auto generated Go binding around an Ethereum contract.
type BasDomain struct {
	BasDomainCaller     // Read-only binding to the contract
	BasDomainTransactor // Write-only binding to the contract
	BasDomainFilterer   // Log filterer for contract events
}

// BasDomainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasDomainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasDomainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasDomainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasDomainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasDomainSession struct {
	Contract     *BasDomain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasDomainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasDomainCallerSession struct {
	Contract *BasDomainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BasDomainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasDomainTransactorSession struct {
	Contract     *BasDomainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasDomainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasDomainRaw struct {
	Contract *BasDomain // Generic contract binding to access the raw methods on
}

// BasDomainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasDomainCallerRaw struct {
	Contract *BasDomainCaller // Generic read-only contract binding to access the raw methods on
}

// BasDomainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasDomainTransactorRaw struct {
	Contract *BasDomainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasDomain creates a new instance of BasDomain, bound to a specific deployed contract.
func NewBasDomain(address common.Address, backend bind.ContractBackend) (*BasDomain, error) {
	contract, err := bindBasDomain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasDomain{BasDomainCaller: BasDomainCaller{contract: contract}, BasDomainTransactor: BasDomainTransactor{contract: contract}, BasDomainFilterer: BasDomainFilterer{contract: contract}}, nil
}

// NewBasDomainCaller creates a new read-only instance of BasDomain, bound to a specific deployed contract.
func NewBasDomainCaller(address common.Address, caller bind.ContractCaller) (*BasDomainCaller, error) {
	contract, err := bindBasDomain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasDomainCaller{contract: contract}, nil
}

// NewBasDomainTransactor creates a new write-only instance of BasDomain, bound to a specific deployed contract.
func NewBasDomainTransactor(address common.Address, transactor bind.ContractTransactor) (*BasDomainTransactor, error) {
	contract, err := bindBasDomain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasDomainTransactor{contract: contract}, nil
}

// NewBasDomainFilterer creates a new log filterer instance of BasDomain, bound to a specific deployed contract.
func NewBasDomainFilterer(address common.Address, filterer bind.ContractFilterer) (*BasDomainFilterer, error) {
	contract, err := bindBasDomain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasDomainFilterer{contract: contract}, nil
}

// bindBasDomain binds a generic wrapper to an already deployed contract.
func bindBasDomain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasDomainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDomain *BasDomainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasDomain.Contract.BasDomainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDomain *BasDomainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDomain.Contract.BasDomainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDomain *BasDomainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDomain.Contract.BasDomainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasDomain *BasDomainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasDomain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasDomain *BasDomainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasDomain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasDomain *BasDomainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasDomain.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomain *BasDomainCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasDomain.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomain *BasDomainSession) DAOAddress() (common.Address, error) {
	return _BasDomain.Contract.DAOAddress(&_BasDomain.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasDomain *BasDomainCallerSession) DAOAddress() (common.Address, error) {
	return _BasDomain.Contract.DAOAddress(&_BasDomain.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomain *BasDomainCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasDomain.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomain *BasDomainSession) Rel() (common.Address, error) {
	return _BasDomain.Contract.Rel(&_BasDomain.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasDomain *BasDomainCallerSession) Rel() (common.Address, error) {
	return _BasDomain.Contract.Rel(&_BasDomain.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomain *BasDomainTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasDomain.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomain *BasDomainSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasDomain.Contract.ChangeDAO(&_BasDomain.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasDomain *BasDomainTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasDomain.Contract.ChangeDAO(&_BasDomain.TransactOpts, newDao)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomain *BasDomainTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasDomain.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomain *BasDomainSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasDomain.Contract.ChangeRelation(&_BasDomain.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasDomain *BasDomainTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasDomain.Contract.ChangeRelation(&_BasDomain.TransactOpts, new_rel)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasDomain *BasDomainTransactor) Recharge(opts *bind.TransactOpts, nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasDomain.contract.Transact(opts, "recharge", nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasDomain *BasDomainSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasDomain.Contract.Recharge(&_BasDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasDomain *BasDomainTransactorSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasDomain.Contract.Recharge(&_BasDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// BasDomainRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the BasDomain contract.
type BasDomainRechargeIterator struct {
	Event *BasDomainRecharge // Event containing the contract specifics and raw log

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
func (it *BasDomainRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasDomainRecharge)
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
		it.Event = new(BasDomainRecharge)
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
func (it *BasDomainRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasDomainRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasDomainRecharge represents a Recharge event raised by the BasDomain contract.
type BasDomainRecharge struct {
	NameHash [32]byte
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasDomain *BasDomainFilterer) FilterRecharge(opts *bind.FilterOpts) (*BasDomainRechargeIterator, error) {

	logs, sub, err := _BasDomain.contract.FilterLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return &BasDomainRechargeIterator{contract: _BasDomain.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasDomain *BasDomainFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *BasDomainRecharge) (event.Subscription, error) {

	logs, sub, err := _BasDomain.contract.WatchLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasDomainRecharge)
				if err := _BasDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
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

// ParseRecharge is a log parse operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasDomain *BasDomainFilterer) ParseRecharge(log types.Log) (*BasDomainRecharge, error) {
	event := new(BasDomainRecharge)
	if err := _BasDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
