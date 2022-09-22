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

// BasMailMetaData contains all meta data concerning the BasMail contract.
var BasMailMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"AbandMail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extendTime\",\"type\":\"uint256\"}],\"name\":\"MailRecharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bcAddress\",\"type\":\"bytes\"}],\"name\":\"MailUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewMail\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"MailRecords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bcAddress\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"abandon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"}],\"name\":\"domainHashOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"}],\"name\":\"newEmail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"extendTime\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bcAddress\",\"type\":\"bytes\"}],\"name\":\"updateMail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasMailABI is the input ABI used to generate the binding from.
// Deprecated: Use BasMailMetaData.ABI instead.
var BasMailABI = BasMailMetaData.ABI

// BasMail is an auto generated Go binding around an Ethereum contract.
type BasMail struct {
	BasMailCaller     // Read-only binding to the contract
	BasMailTransactor // Write-only binding to the contract
	BasMailFilterer   // Log filterer for contract events
}

// BasMailCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMailCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMailTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMailFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMailSession struct {
	Contract     *BasMail          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMailCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMailCallerSession struct {
	Contract *BasMailCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BasMailTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMailTransactorSession struct {
	Contract     *BasMailTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasMailRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMailRaw struct {
	Contract *BasMail // Generic contract binding to access the raw methods on
}

// BasMailCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMailCallerRaw struct {
	Contract *BasMailCaller // Generic read-only contract binding to access the raw methods on
}

// BasMailTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMailTransactorRaw struct {
	Contract *BasMailTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMail creates a new instance of BasMail, bound to a specific deployed contract.
func NewBasMail(address common.Address, backend bind.ContractBackend) (*BasMail, error) {
	contract, err := bindBasMail(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMail{BasMailCaller: BasMailCaller{contract: contract}, BasMailTransactor: BasMailTransactor{contract: contract}, BasMailFilterer: BasMailFilterer{contract: contract}}, nil
}

// NewBasMailCaller creates a new read-only instance of BasMail, bound to a specific deployed contract.
func NewBasMailCaller(address common.Address, caller bind.ContractCaller) (*BasMailCaller, error) {
	contract, err := bindBasMail(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMailCaller{contract: contract}, nil
}

// NewBasMailTransactor creates a new write-only instance of BasMail, bound to a specific deployed contract.
func NewBasMailTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMailTransactor, error) {
	contract, err := bindBasMail(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMailTransactor{contract: contract}, nil
}

// NewBasMailFilterer creates a new log filterer instance of BasMail, bound to a specific deployed contract.
func NewBasMailFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMailFilterer, error) {
	contract, err := bindBasMail(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMailFilterer{contract: contract}, nil
}

// bindBasMail binds a generic wrapper to an already deployed contract.
func bindBasMail(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMailABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMail *BasMailRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMail.Contract.BasMailCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMail *BasMailRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMail.Contract.BasMailTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMail *BasMailRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMail.Contract.BasMailTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMail *BasMailCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMail.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMail *BasMailTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMail.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMail *BasMailTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMail.Contract.contract.Transact(opts, method, params...)
}

// MailRecords is a free data retrieval call binding the contract method 0x62288e2a.
//
// Solidity: function MailRecords(bytes32 ) view returns(bytes32 domainHash, bytes32 mailHash, bool valid, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailCaller) MailRecords(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DomainHash [32]byte
	MailHash   [32]byte
	Valid      bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	var out []interface{}
	err := _BasMail.contract.Call(opts, &out, "MailRecords", arg0)

	outstruct := new(struct {
		DomainHash [32]byte
		MailHash   [32]byte
		Valid      bool
		AliasName  []byte
		BcAddress  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DomainHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.MailHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Valid = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.AliasName = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.BcAddress = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// MailRecords is a free data retrieval call binding the contract method 0x62288e2a.
//
// Solidity: function MailRecords(bytes32 ) view returns(bytes32 domainHash, bytes32 mailHash, bool valid, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailSession) MailRecords(arg0 [32]byte) (struct {
	DomainHash [32]byte
	MailHash   [32]byte
	Valid      bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	return _BasMail.Contract.MailRecords(&_BasMail.CallOpts, arg0)
}

// MailRecords is a free data retrieval call binding the contract method 0x62288e2a.
//
// Solidity: function MailRecords(bytes32 ) view returns(bytes32 domainHash, bytes32 mailHash, bool valid, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailCallerSession) MailRecords(arg0 [32]byte) (struct {
	DomainHash [32]byte
	MailHash   [32]byte
	Valid      bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	return _BasMail.Contract.MailRecords(&_BasMail.CallOpts, arg0)
}

// DomainHashOf is a free data retrieval call binding the contract method 0x1f5d78ad.
//
// Solidity: function domainHashOf(bytes32 mailHash) view returns(bytes32)
func (_BasMail *BasMailCaller) DomainHashOf(opts *bind.CallOpts, mailHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BasMail.contract.Call(opts, &out, "domainHashOf", mailHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainHashOf is a free data retrieval call binding the contract method 0x1f5d78ad.
//
// Solidity: function domainHashOf(bytes32 mailHash) view returns(bytes32)
func (_BasMail *BasMailSession) DomainHashOf(mailHash [32]byte) ([32]byte, error) {
	return _BasMail.Contract.DomainHashOf(&_BasMail.CallOpts, mailHash)
}

// DomainHashOf is a free data retrieval call binding the contract method 0x1f5d78ad.
//
// Solidity: function domainHashOf(bytes32 mailHash) view returns(bytes32)
func (_BasMail *BasMailCallerSession) DomainHashOf(mailHash [32]byte) ([32]byte, error) {
	return _BasMail.Contract.DomainHashOf(&_BasMail.CallOpts, mailHash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMail *BasMailCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMail.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMail *BasMailSession) Rel() (common.Address, error) {
	return _BasMail.Contract.Rel(&_BasMail.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMail *BasMailCallerSession) Rel() (common.Address, error) {
	return _BasMail.Contract.Rel(&_BasMail.CallOpts)
}

// Abandon is a paid mutator transaction binding the contract method 0xa4242091.
//
// Solidity: function abandon(bytes32 hash) returns()
func (_BasMail *BasMailTransactor) Abandon(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _BasMail.contract.Transact(opts, "abandon", hash)
}

// Abandon is a paid mutator transaction binding the contract method 0xa4242091.
//
// Solidity: function abandon(bytes32 hash) returns()
func (_BasMail *BasMailSession) Abandon(hash [32]byte) (*types.Transaction, error) {
	return _BasMail.Contract.Abandon(&_BasMail.TransactOpts, hash)
}

// Abandon is a paid mutator transaction binding the contract method 0xa4242091.
//
// Solidity: function abandon(bytes32 hash) returns()
func (_BasMail *BasMailTransactorSession) Abandon(hash [32]byte) (*types.Transaction, error) {
	return _BasMail.Contract.Abandon(&_BasMail.TransactOpts, hash)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMail *BasMailTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasMail.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMail *BasMailSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMail.Contract.ChangeRelation(&_BasMail.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMail *BasMailTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMail.Contract.ChangeRelation(&_BasMail.TransactOpts, new_rel)
}

// NewEmail is a paid mutator transaction binding the contract method 0x756c11d8.
//
// Solidity: function newEmail(bytes32 domainHash, bytes32 mailHash, address owner, uint256 expire, bytes aliasName) returns()
func (_BasMail *BasMailTransactor) NewEmail(opts *bind.TransactOpts, domainHash [32]byte, mailHash [32]byte, owner common.Address, expire *big.Int, aliasName []byte) (*types.Transaction, error) {
	return _BasMail.contract.Transact(opts, "newEmail", domainHash, mailHash, owner, expire, aliasName)
}

// NewEmail is a paid mutator transaction binding the contract method 0x756c11d8.
//
// Solidity: function newEmail(bytes32 domainHash, bytes32 mailHash, address owner, uint256 expire, bytes aliasName) returns()
func (_BasMail *BasMailSession) NewEmail(domainHash [32]byte, mailHash [32]byte, owner common.Address, expire *big.Int, aliasName []byte) (*types.Transaction, error) {
	return _BasMail.Contract.NewEmail(&_BasMail.TransactOpts, domainHash, mailHash, owner, expire, aliasName)
}

// NewEmail is a paid mutator transaction binding the contract method 0x756c11d8.
//
// Solidity: function newEmail(bytes32 domainHash, bytes32 mailHash, address owner, uint256 expire, bytes aliasName) returns()
func (_BasMail *BasMailTransactorSession) NewEmail(domainHash [32]byte, mailHash [32]byte, owner common.Address, expire *big.Int, aliasName []byte) (*types.Transaction, error) {
	return _BasMail.Contract.NewEmail(&_BasMail.TransactOpts, domainHash, mailHash, owner, expire, aliasName)
}

// Recharge is a paid mutator transaction binding the contract method 0x8133aff0.
//
// Solidity: function recharge(bytes32 mailHash, uint256 extendTime) returns()
func (_BasMail *BasMailTransactor) Recharge(opts *bind.TransactOpts, mailHash [32]byte, extendTime *big.Int) (*types.Transaction, error) {
	return _BasMail.contract.Transact(opts, "recharge", mailHash, extendTime)
}

// Recharge is a paid mutator transaction binding the contract method 0x8133aff0.
//
// Solidity: function recharge(bytes32 mailHash, uint256 extendTime) returns()
func (_BasMail *BasMailSession) Recharge(mailHash [32]byte, extendTime *big.Int) (*types.Transaction, error) {
	return _BasMail.Contract.Recharge(&_BasMail.TransactOpts, mailHash, extendTime)
}

// Recharge is a paid mutator transaction binding the contract method 0x8133aff0.
//
// Solidity: function recharge(bytes32 mailHash, uint256 extendTime) returns()
func (_BasMail *BasMailTransactorSession) Recharge(mailHash [32]byte, extendTime *big.Int) (*types.Transaction, error) {
	return _BasMail.Contract.Recharge(&_BasMail.TransactOpts, mailHash, extendTime)
}

// UpdateMail is a paid mutator transaction binding the contract method 0x2ae48f5e.
//
// Solidity: function updateMail(bytes32 mailHash, bytes aliasName, bytes bcAddress) returns()
func (_BasMail *BasMailTransactor) UpdateMail(opts *bind.TransactOpts, mailHash [32]byte, aliasName []byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasMail.contract.Transact(opts, "updateMail", mailHash, aliasName, bcAddress)
}

// UpdateMail is a paid mutator transaction binding the contract method 0x2ae48f5e.
//
// Solidity: function updateMail(bytes32 mailHash, bytes aliasName, bytes bcAddress) returns()
func (_BasMail *BasMailSession) UpdateMail(mailHash [32]byte, aliasName []byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasMail.Contract.UpdateMail(&_BasMail.TransactOpts, mailHash, aliasName, bcAddress)
}

// UpdateMail is a paid mutator transaction binding the contract method 0x2ae48f5e.
//
// Solidity: function updateMail(bytes32 mailHash, bytes aliasName, bytes bcAddress) returns()
func (_BasMail *BasMailTransactorSession) UpdateMail(mailHash [32]byte, aliasName []byte, bcAddress []byte) (*types.Transaction, error) {
	return _BasMail.Contract.UpdateMail(&_BasMail.TransactOpts, mailHash, aliasName, bcAddress)
}

// BasMailAbandMailIterator is returned from FilterAbandMail and is used to iterate over the raw logs and unpacked data for AbandMail events raised by the BasMail contract.
type BasMailAbandMailIterator struct {
	Event *BasMailAbandMail // Event containing the contract specifics and raw log

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
func (it *BasMailAbandMailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailAbandMail)
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
		it.Event = new(BasMailAbandMail)
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
func (it *BasMailAbandMailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailAbandMailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailAbandMail represents a AbandMail event raised by the BasMail contract.
type BasMailAbandMail struct {
	DomainHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAbandMail is a free log retrieval operation binding the contract event 0x0175419262f4cada22c8d14a6739ac30cea1b70b6c1568c903a7bfad75aeeb98.
//
// Solidity: event AbandMail(bytes32 domainHash)
func (_BasMail *BasMailFilterer) FilterAbandMail(opts *bind.FilterOpts) (*BasMailAbandMailIterator, error) {

	logs, sub, err := _BasMail.contract.FilterLogs(opts, "AbandMail")
	if err != nil {
		return nil, err
	}
	return &BasMailAbandMailIterator{contract: _BasMail.contract, event: "AbandMail", logs: logs, sub: sub}, nil
}

// WatchAbandMail is a free log subscription operation binding the contract event 0x0175419262f4cada22c8d14a6739ac30cea1b70b6c1568c903a7bfad75aeeb98.
//
// Solidity: event AbandMail(bytes32 domainHash)
func (_BasMail *BasMailFilterer) WatchAbandMail(opts *bind.WatchOpts, sink chan<- *BasMailAbandMail) (event.Subscription, error) {

	logs, sub, err := _BasMail.contract.WatchLogs(opts, "AbandMail")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailAbandMail)
				if err := _BasMail.contract.UnpackLog(event, "AbandMail", log); err != nil {
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

// ParseAbandMail is a log parse operation binding the contract event 0x0175419262f4cada22c8d14a6739ac30cea1b70b6c1568c903a7bfad75aeeb98.
//
// Solidity: event AbandMail(bytes32 domainHash)
func (_BasMail *BasMailFilterer) ParseAbandMail(log types.Log) (*BasMailAbandMail, error) {
	event := new(BasMailAbandMail)
	if err := _BasMail.contract.UnpackLog(event, "AbandMail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailMailRechargedIterator is returned from FilterMailRecharged and is used to iterate over the raw logs and unpacked data for MailRecharged events raised by the BasMail contract.
type BasMailMailRechargedIterator struct {
	Event *BasMailMailRecharged // Event containing the contract specifics and raw log

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
func (it *BasMailMailRechargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailMailRecharged)
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
		it.Event = new(BasMailMailRecharged)
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
func (it *BasMailMailRechargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailMailRechargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailMailRecharged represents a MailRecharged event raised by the BasMail contract.
type BasMailMailRecharged struct {
	DomainHash [32]byte
	ExtendTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMailRecharged is a free log retrieval operation binding the contract event 0x1c0239827918e02705374caee2ddd0ca4e87be1d332dd4d81940ff0aa7f5b2c4.
//
// Solidity: event MailRecharged(bytes32 domainHash, uint256 extendTime)
func (_BasMail *BasMailFilterer) FilterMailRecharged(opts *bind.FilterOpts) (*BasMailMailRechargedIterator, error) {

	logs, sub, err := _BasMail.contract.FilterLogs(opts, "MailRecharged")
	if err != nil {
		return nil, err
	}
	return &BasMailMailRechargedIterator{contract: _BasMail.contract, event: "MailRecharged", logs: logs, sub: sub}, nil
}

// WatchMailRecharged is a free log subscription operation binding the contract event 0x1c0239827918e02705374caee2ddd0ca4e87be1d332dd4d81940ff0aa7f5b2c4.
//
// Solidity: event MailRecharged(bytes32 domainHash, uint256 extendTime)
func (_BasMail *BasMailFilterer) WatchMailRecharged(opts *bind.WatchOpts, sink chan<- *BasMailMailRecharged) (event.Subscription, error) {

	logs, sub, err := _BasMail.contract.WatchLogs(opts, "MailRecharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailMailRecharged)
				if err := _BasMail.contract.UnpackLog(event, "MailRecharged", log); err != nil {
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

// ParseMailRecharged is a log parse operation binding the contract event 0x1c0239827918e02705374caee2ddd0ca4e87be1d332dd4d81940ff0aa7f5b2c4.
//
// Solidity: event MailRecharged(bytes32 domainHash, uint256 extendTime)
func (_BasMail *BasMailFilterer) ParseMailRecharged(log types.Log) (*BasMailMailRecharged, error) {
	event := new(BasMailMailRecharged)
	if err := _BasMail.contract.UnpackLog(event, "MailRecharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailMailUpdateIterator is returned from FilterMailUpdate and is used to iterate over the raw logs and unpacked data for MailUpdate events raised by the BasMail contract.
type BasMailMailUpdateIterator struct {
	Event *BasMailMailUpdate // Event containing the contract specifics and raw log

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
func (it *BasMailMailUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailMailUpdate)
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
		it.Event = new(BasMailMailUpdate)
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
func (it *BasMailMailUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailMailUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailMailUpdate represents a MailUpdate event raised by the BasMail contract.
type BasMailMailUpdate struct {
	DomainHash [32]byte
	AliasName  []byte
	BcAddress  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMailUpdate is a free log retrieval operation binding the contract event 0x572ecefef376a39b01928600d236f6e11cbff1bdb9d30a5dd7d72327c9a07c05.
//
// Solidity: event MailUpdate(bytes32 domainHash, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailFilterer) FilterMailUpdate(opts *bind.FilterOpts) (*BasMailMailUpdateIterator, error) {

	logs, sub, err := _BasMail.contract.FilterLogs(opts, "MailUpdate")
	if err != nil {
		return nil, err
	}
	return &BasMailMailUpdateIterator{contract: _BasMail.contract, event: "MailUpdate", logs: logs, sub: sub}, nil
}

// WatchMailUpdate is a free log subscription operation binding the contract event 0x572ecefef376a39b01928600d236f6e11cbff1bdb9d30a5dd7d72327c9a07c05.
//
// Solidity: event MailUpdate(bytes32 domainHash, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailFilterer) WatchMailUpdate(opts *bind.WatchOpts, sink chan<- *BasMailMailUpdate) (event.Subscription, error) {

	logs, sub, err := _BasMail.contract.WatchLogs(opts, "MailUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailMailUpdate)
				if err := _BasMail.contract.UnpackLog(event, "MailUpdate", log); err != nil {
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

// ParseMailUpdate is a log parse operation binding the contract event 0x572ecefef376a39b01928600d236f6e11cbff1bdb9d30a5dd7d72327c9a07c05.
//
// Solidity: event MailUpdate(bytes32 domainHash, bytes aliasName, bytes bcAddress)
func (_BasMail *BasMailFilterer) ParseMailUpdate(log types.Log) (*BasMailMailUpdate, error) {
	event := new(BasMailMailUpdate)
	if err := _BasMail.contract.UnpackLog(event, "MailUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailNewMailIterator is returned from FilterNewMail and is used to iterate over the raw logs and unpacked data for NewMail events raised by the BasMail contract.
type BasMailNewMailIterator struct {
	Event *BasMailNewMail // Event containing the contract specifics and raw log

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
func (it *BasMailNewMailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailNewMail)
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
		it.Event = new(BasMailNewMail)
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
func (it *BasMailNewMailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailNewMailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailNewMail represents a NewMail event raised by the BasMail contract.
type BasMailNewMail struct {
	DomainHash [32]byte
	NameHash   [32]byte
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewMail is a free log retrieval operation binding the contract event 0x20f3794f70576de98bb3758dc95010f495b687345fa10628b0e4016e20a42607.
//
// Solidity: event NewMail(bytes32 domainHash, bytes32 nameHash, address owner)
func (_BasMail *BasMailFilterer) FilterNewMail(opts *bind.FilterOpts) (*BasMailNewMailIterator, error) {

	logs, sub, err := _BasMail.contract.FilterLogs(opts, "NewMail")
	if err != nil {
		return nil, err
	}
	return &BasMailNewMailIterator{contract: _BasMail.contract, event: "NewMail", logs: logs, sub: sub}, nil
}

// WatchNewMail is a free log subscription operation binding the contract event 0x20f3794f70576de98bb3758dc95010f495b687345fa10628b0e4016e20a42607.
//
// Solidity: event NewMail(bytes32 domainHash, bytes32 nameHash, address owner)
func (_BasMail *BasMailFilterer) WatchNewMail(opts *bind.WatchOpts, sink chan<- *BasMailNewMail) (event.Subscription, error) {

	logs, sub, err := _BasMail.contract.WatchLogs(opts, "NewMail")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailNewMail)
				if err := _BasMail.contract.UnpackLog(event, "NewMail", log); err != nil {
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

// ParseNewMail is a log parse operation binding the contract event 0x20f3794f70576de98bb3758dc95010f495b687345fa10628b0e4016e20a42607.
//
// Solidity: event NewMail(bytes32 domainHash, bytes32 nameHash, address owner)
func (_BasMail *BasMailFilterer) ParseNewMail(log types.Log) (*BasMailNewMail, error) {
	event := new(BasMailNewMail)
	if err := _BasMail.contract.UnpackLog(event, "NewMail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
