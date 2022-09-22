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

// BasAccountantMetaData contains all meta data concerning the BasAccountant contract.
var BasAccountantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Allocate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"drawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"allocateProfitAdvanced\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"allocateProfitNormal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"}],\"name\":\"daoWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasAccountantABI is the input ABI used to generate the binding from.
// Deprecated: Use BasAccountantMetaData.ABI instead.
var BasAccountantABI = BasAccountantMetaData.ABI

// BasAccountant is an auto generated Go binding around an Ethereum contract.
type BasAccountant struct {
	BasAccountantCaller     // Read-only binding to the contract
	BasAccountantTransactor // Write-only binding to the contract
	BasAccountantFilterer   // Log filterer for contract events
}

// BasAccountantCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasAccountantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAccountantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasAccountantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAccountantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasAccountantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasAccountantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasAccountantSession struct {
	Contract     *BasAccountant    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasAccountantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasAccountantCallerSession struct {
	Contract *BasAccountantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BasAccountantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasAccountantTransactorSession struct {
	Contract     *BasAccountantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BasAccountantRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasAccountantRaw struct {
	Contract *BasAccountant // Generic contract binding to access the raw methods on
}

// BasAccountantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasAccountantCallerRaw struct {
	Contract *BasAccountantCaller // Generic read-only contract binding to access the raw methods on
}

// BasAccountantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasAccountantTransactorRaw struct {
	Contract *BasAccountantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasAccountant creates a new instance of BasAccountant, bound to a specific deployed contract.
func NewBasAccountant(address common.Address, backend bind.ContractBackend) (*BasAccountant, error) {
	contract, err := bindBasAccountant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasAccountant{BasAccountantCaller: BasAccountantCaller{contract: contract}, BasAccountantTransactor: BasAccountantTransactor{contract: contract}, BasAccountantFilterer: BasAccountantFilterer{contract: contract}}, nil
}

// NewBasAccountantCaller creates a new read-only instance of BasAccountant, bound to a specific deployed contract.
func NewBasAccountantCaller(address common.Address, caller bind.ContractCaller) (*BasAccountantCaller, error) {
	contract, err := bindBasAccountant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasAccountantCaller{contract: contract}, nil
}

// NewBasAccountantTransactor creates a new write-only instance of BasAccountant, bound to a specific deployed contract.
func NewBasAccountantTransactor(address common.Address, transactor bind.ContractTransactor) (*BasAccountantTransactor, error) {
	contract, err := bindBasAccountant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasAccountantTransactor{contract: contract}, nil
}

// NewBasAccountantFilterer creates a new log filterer instance of BasAccountant, bound to a specific deployed contract.
func NewBasAccountantFilterer(address common.Address, filterer bind.ContractFilterer) (*BasAccountantFilterer, error) {
	contract, err := bindBasAccountant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasAccountantFilterer{contract: contract}, nil
}

// bindBasAccountant binds a generic wrapper to an already deployed contract.
func bindBasAccountant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasAccountantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasAccountant *BasAccountantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasAccountant.Contract.BasAccountantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasAccountant *BasAccountantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasAccountant.Contract.BasAccountantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasAccountant *BasAccountantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasAccountant.Contract.BasAccountantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasAccountant *BasAccountantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasAccountant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasAccountant *BasAccountantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasAccountant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasAccountant *BasAccountantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasAccountant.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasAccountant *BasAccountantCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasAccountant.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasAccountant *BasAccountantSession) DAOAddress() (common.Address, error) {
	return _BasAccountant.Contract.DAOAddress(&_BasAccountant.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasAccountant *BasAccountantCallerSession) DAOAddress() (common.Address, error) {
	return _BasAccountant.Contract.DAOAddress(&_BasAccountant.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0xfbfa941f.
//
// Solidity: function ledger(address ) view returns(uint256)
func (_BasAccountant *BasAccountantCaller) Ledger(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasAccountant.contract.Call(opts, &out, "ledger", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ledger is a free data retrieval call binding the contract method 0xfbfa941f.
//
// Solidity: function ledger(address ) view returns(uint256)
func (_BasAccountant *BasAccountantSession) Ledger(arg0 common.Address) (*big.Int, error) {
	return _BasAccountant.Contract.Ledger(&_BasAccountant.CallOpts, arg0)
}

// Ledger is a free data retrieval call binding the contract method 0xfbfa941f.
//
// Solidity: function ledger(address ) view returns(uint256)
func (_BasAccountant *BasAccountantCallerSession) Ledger(arg0 common.Address) (*big.Int, error) {
	return _BasAccountant.Contract.Ledger(&_BasAccountant.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasAccountant *BasAccountantCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasAccountant.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasAccountant *BasAccountantSession) Rel() (common.Address, error) {
	return _BasAccountant.Contract.Rel(&_BasAccountant.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasAccountant *BasAccountantCallerSession) Rel() (common.Address, error) {
	return _BasAccountant.Contract.Rel(&_BasAccountant.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasAccountant *BasAccountantTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasAccountant *BasAccountantSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.ChangeDAO(&_BasAccountant.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasAccountant *BasAccountantTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.ChangeDAO(&_BasAccountant.TransactOpts, newDao)
}

// AllocateProfitAdvanced is a paid mutator transaction binding the contract method 0xd49141db.
//
// Solidity: function allocateProfitAdvanced(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantTransactor) AllocateProfitAdvanced(opts *bind.TransactOpts, receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "allocateProfitAdvanced", receiver, value)
}

// AllocateProfitAdvanced is a paid mutator transaction binding the contract method 0xd49141db.
//
// Solidity: function allocateProfitAdvanced(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantSession) AllocateProfitAdvanced(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.AllocateProfitAdvanced(&_BasAccountant.TransactOpts, receiver, value)
}

// AllocateProfitAdvanced is a paid mutator transaction binding the contract method 0xd49141db.
//
// Solidity: function allocateProfitAdvanced(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantTransactorSession) AllocateProfitAdvanced(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.AllocateProfitAdvanced(&_BasAccountant.TransactOpts, receiver, value)
}

// AllocateProfitNormal is a paid mutator transaction binding the contract method 0xc22e5fbf.
//
// Solidity: function allocateProfitNormal(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantTransactor) AllocateProfitNormal(opts *bind.TransactOpts, receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "allocateProfitNormal", receiver, value)
}

// AllocateProfitNormal is a paid mutator transaction binding the contract method 0xc22e5fbf.
//
// Solidity: function allocateProfitNormal(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantSession) AllocateProfitNormal(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.AllocateProfitNormal(&_BasAccountant.TransactOpts, receiver, value)
}

// AllocateProfitNormal is a paid mutator transaction binding the contract method 0xc22e5fbf.
//
// Solidity: function allocateProfitNormal(address receiver, uint256 value) returns()
func (_BasAccountant *BasAccountantTransactorSession) AllocateProfitNormal(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.AllocateProfitNormal(&_BasAccountant.TransactOpts, receiver, value)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasAccountant *BasAccountantTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasAccountant *BasAccountantSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.ChangeRelation(&_BasAccountant.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasAccountant *BasAccountantTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.ChangeRelation(&_BasAccountant.TransactOpts, new_rel)
}

// DaoWithdraw is a paid mutator transaction binding the contract method 0x0436734c.
//
// Solidity: function daoWithdraw(address to, uint256 no) returns()
func (_BasAccountant *BasAccountantTransactor) DaoWithdraw(opts *bind.TransactOpts, to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "daoWithdraw", to, no)
}

// DaoWithdraw is a paid mutator transaction binding the contract method 0x0436734c.
//
// Solidity: function daoWithdraw(address to, uint256 no) returns()
func (_BasAccountant *BasAccountantSession) DaoWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.DaoWithdraw(&_BasAccountant.TransactOpts, to, no)
}

// DaoWithdraw is a paid mutator transaction binding the contract method 0x0436734c.
//
// Solidity: function daoWithdraw(address to, uint256 no) returns()
func (_BasAccountant *BasAccountantTransactorSession) DaoWithdraw(to common.Address, no *big.Int) (*types.Transaction, error) {
	return _BasAccountant.Contract.DaoWithdraw(&_BasAccountant.TransactOpts, to, no)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address target) returns()
func (_BasAccountant *BasAccountantTransactor) WithdrawTo(opts *bind.TransactOpts, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasAccountant.contract.Transact(opts, "withdrawTo", amount, target)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address target) returns()
func (_BasAccountant *BasAccountantSession) WithdrawTo(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.WithdrawTo(&_BasAccountant.TransactOpts, amount, target)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address target) returns()
func (_BasAccountant *BasAccountantTransactorSession) WithdrawTo(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _BasAccountant.Contract.WithdrawTo(&_BasAccountant.TransactOpts, amount, target)
}

// BasAccountantAllocateIterator is returned from FilterAllocate and is used to iterate over the raw logs and unpacked data for Allocate events raised by the BasAccountant contract.
type BasAccountantAllocateIterator struct {
	Event *BasAccountantAllocate // Event containing the contract specifics and raw log

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
func (it *BasAccountantAllocateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAccountantAllocate)
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
		it.Event = new(BasAccountantAllocate)
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
func (it *BasAccountantAllocateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAccountantAllocateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAccountantAllocate represents a Allocate event raised by the BasAccountant contract.
type BasAccountantAllocate struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllocate is a free log retrieval operation binding the contract event 0x249d8eb76d5a22983620d741de2470148d1a9a26ab923aec4262770690d11ebc.
//
// Solidity: event Allocate(address indexed receiver, uint256 amount)
func (_BasAccountant *BasAccountantFilterer) FilterAllocate(opts *bind.FilterOpts, receiver []common.Address) (*BasAccountantAllocateIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BasAccountant.contract.FilterLogs(opts, "Allocate", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BasAccountantAllocateIterator{contract: _BasAccountant.contract, event: "Allocate", logs: logs, sub: sub}, nil
}

// WatchAllocate is a free log subscription operation binding the contract event 0x249d8eb76d5a22983620d741de2470148d1a9a26ab923aec4262770690d11ebc.
//
// Solidity: event Allocate(address indexed receiver, uint256 amount)
func (_BasAccountant *BasAccountantFilterer) WatchAllocate(opts *bind.WatchOpts, sink chan<- *BasAccountantAllocate, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BasAccountant.contract.WatchLogs(opts, "Allocate", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAccountantAllocate)
				if err := _BasAccountant.contract.UnpackLog(event, "Allocate", log); err != nil {
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

// ParseAllocate is a log parse operation binding the contract event 0x249d8eb76d5a22983620d741de2470148d1a9a26ab923aec4262770690d11ebc.
//
// Solidity: event Allocate(address indexed receiver, uint256 amount)
func (_BasAccountant *BasAccountantFilterer) ParseAllocate(log types.Log) (*BasAccountantAllocate, error) {
	event := new(BasAccountantAllocate)
	if err := _BasAccountant.contract.UnpackLog(event, "Allocate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasAccountantWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BasAccountant contract.
type BasAccountantWithdrawIterator struct {
	Event *BasAccountantWithdraw // Event containing the contract specifics and raw log

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
func (it *BasAccountantWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasAccountantWithdraw)
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
		it.Event = new(BasAccountantWithdraw)
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
func (it *BasAccountantWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasAccountantWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasAccountantWithdraw represents a Withdraw event raised by the BasAccountant contract.
type BasAccountantWithdraw struct {
	Drawer common.Address
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amount, address to)
func (_BasAccountant *BasAccountantFilterer) FilterWithdraw(opts *bind.FilterOpts, drawer []common.Address) (*BasAccountantWithdrawIterator, error) {

	var drawerRule []interface{}
	for _, drawerItem := range drawer {
		drawerRule = append(drawerRule, drawerItem)
	}

	logs, sub, err := _BasAccountant.contract.FilterLogs(opts, "Withdraw", drawerRule)
	if err != nil {
		return nil, err
	}
	return &BasAccountantWithdrawIterator{contract: _BasAccountant.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amount, address to)
func (_BasAccountant *BasAccountantFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BasAccountantWithdraw, drawer []common.Address) (event.Subscription, error) {

	var drawerRule []interface{}
	for _, drawerItem := range drawer {
		drawerRule = append(drawerRule, drawerItem)
	}

	logs, sub, err := _BasAccountant.contract.WatchLogs(opts, "Withdraw", drawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasAccountantWithdraw)
				if err := _BasAccountant.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amount, address to)
func (_BasAccountant *BasAccountantFilterer) ParseWithdraw(log types.Log) (*BasAccountantWithdraw, error) {
	event := new(BasAccountantWithdraw)
	if err := _BasAccountant.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
