// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnscontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WhiteListABI is the input ABI used to generate the binding from.
const WhiteListABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DelWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IsInWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cnt\",\"type\":\"uint256\"}],\"name\":\"SetMintCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WhiteList is an auto generated Go binding around an Ethereum contract.
type WhiteList struct {
	WhiteListCaller     // Read-only binding to the contract
	WhiteListTransactor // Write-only binding to the contract
	WhiteListFilterer   // Log filterer for contract events
}

// WhiteListCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhiteListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhiteListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhiteListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhiteListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhiteListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhiteListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhiteListSession struct {
	Contract     *WhiteList        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhiteListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhiteListCallerSession struct {
	Contract *WhiteListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhiteListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhiteListTransactorSession struct {
	Contract     *WhiteListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhiteListRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhiteListRaw struct {
	Contract *WhiteList // Generic contract binding to access the raw methods on
}

// WhiteListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhiteListCallerRaw struct {
	Contract *WhiteListCaller // Generic read-only contract binding to access the raw methods on
}

// WhiteListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhiteListTransactorRaw struct {
	Contract *WhiteListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhiteList creates a new instance of WhiteList, bound to a specific deployed contract.
func NewWhiteList(address common.Address, backend bind.ContractBackend) (*WhiteList, error) {
	contract, err := bindWhiteList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhiteList{WhiteListCaller: WhiteListCaller{contract: contract}, WhiteListTransactor: WhiteListTransactor{contract: contract}, WhiteListFilterer: WhiteListFilterer{contract: contract}}, nil
}

// NewWhiteListCaller creates a new read-only instance of WhiteList, bound to a specific deployed contract.
func NewWhiteListCaller(address common.Address, caller bind.ContractCaller) (*WhiteListCaller, error) {
	contract, err := bindWhiteList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhiteListCaller{contract: contract}, nil
}

// NewWhiteListTransactor creates a new write-only instance of WhiteList, bound to a specific deployed contract.
func NewWhiteListTransactor(address common.Address, transactor bind.ContractTransactor) (*WhiteListTransactor, error) {
	contract, err := bindWhiteList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhiteListTransactor{contract: contract}, nil
}

// NewWhiteListFilterer creates a new log filterer instance of WhiteList, bound to a specific deployed contract.
func NewWhiteListFilterer(address common.Address, filterer bind.ContractFilterer) (*WhiteListFilterer, error) {
	contract, err := bindWhiteList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhiteListFilterer{contract: contract}, nil
}

// bindWhiteList binds a generic wrapper to an already deployed contract.
func bindWhiteList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhiteListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhiteList *WhiteListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhiteList.Contract.WhiteListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhiteList *WhiteListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhiteList.Contract.WhiteListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhiteList *WhiteListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhiteList.Contract.WhiteListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhiteList *WhiteListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhiteList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhiteList *WhiteListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhiteList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhiteList *WhiteListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhiteList.Contract.contract.Transact(opts, method, params...)
}

// IsInWhiteList is a free data retrieval call binding the contract method 0x7c449570.
//
// Solidity: function IsInWhiteList() view returns(bool)
func (_WhiteList *WhiteListCaller) IsInWhiteList(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WhiteList.contract.Call(opts, &out, "IsInWhiteList")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWhiteList is a free data retrieval call binding the contract method 0x7c449570.
//
// Solidity: function IsInWhiteList() view returns(bool)
func (_WhiteList *WhiteListSession) IsInWhiteList() (bool, error) {
	return _WhiteList.Contract.IsInWhiteList(&_WhiteList.CallOpts)
}

// IsInWhiteList is a free data retrieval call binding the contract method 0x7c449570.
//
// Solidity: function IsInWhiteList() view returns(bool)
func (_WhiteList *WhiteListCallerSession) IsInWhiteList() (bool, error) {
	return _WhiteList.Contract.IsInWhiteList(&_WhiteList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhiteList *WhiteListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhiteList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhiteList *WhiteListSession) Owner() (common.Address, error) {
	return _WhiteList.Contract.Owner(&_WhiteList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhiteList *WhiteListCallerSession) Owner() (common.Address, error) {
	return _WhiteList.Contract.Owner(&_WhiteList.CallOpts)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x55cc02a2.
//
// Solidity: function AddWhiteList() returns()
func (_WhiteList *WhiteListTransactor) AddWhiteList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhiteList.contract.Transact(opts, "AddWhiteList")
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x55cc02a2.
//
// Solidity: function AddWhiteList() returns()
func (_WhiteList *WhiteListSession) AddWhiteList() (*types.Transaction, error) {
	return _WhiteList.Contract.AddWhiteList(&_WhiteList.TransactOpts)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x55cc02a2.
//
// Solidity: function AddWhiteList() returns()
func (_WhiteList *WhiteListTransactorSession) AddWhiteList() (*types.Transaction, error) {
	return _WhiteList.Contract.AddWhiteList(&_WhiteList.TransactOpts)
}

// DelWhiteList is a paid mutator transaction binding the contract method 0x3d046199.
//
// Solidity: function DelWhiteList() returns()
func (_WhiteList *WhiteListTransactor) DelWhiteList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhiteList.contract.Transact(opts, "DelWhiteList")
}

// DelWhiteList is a paid mutator transaction binding the contract method 0x3d046199.
//
// Solidity: function DelWhiteList() returns()
func (_WhiteList *WhiteListSession) DelWhiteList() (*types.Transaction, error) {
	return _WhiteList.Contract.DelWhiteList(&_WhiteList.TransactOpts)
}

// DelWhiteList is a paid mutator transaction binding the contract method 0x3d046199.
//
// Solidity: function DelWhiteList() returns()
func (_WhiteList *WhiteListTransactorSession) DelWhiteList() (*types.Transaction, error) {
	return _WhiteList.Contract.DelWhiteList(&_WhiteList.TransactOpts)
}

// SetMintCount is a paid mutator transaction binding the contract method 0x9942c87d.
//
// Solidity: function SetMintCount(uint256 cnt) returns()
func (_WhiteList *WhiteListTransactor) SetMintCount(opts *bind.TransactOpts, cnt *big.Int) (*types.Transaction, error) {
	return _WhiteList.contract.Transact(opts, "SetMintCount", cnt)
}

// SetMintCount is a paid mutator transaction binding the contract method 0x9942c87d.
//
// Solidity: function SetMintCount(uint256 cnt) returns()
func (_WhiteList *WhiteListSession) SetMintCount(cnt *big.Int) (*types.Transaction, error) {
	return _WhiteList.Contract.SetMintCount(&_WhiteList.TransactOpts, cnt)
}

// SetMintCount is a paid mutator transaction binding the contract method 0x9942c87d.
//
// Solidity: function SetMintCount(uint256 cnt) returns()
func (_WhiteList *WhiteListTransactorSession) SetMintCount(cnt *big.Int) (*types.Transaction, error) {
	return _WhiteList.Contract.SetMintCount(&_WhiteList.TransactOpts, cnt)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhiteList *WhiteListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WhiteList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhiteList *WhiteListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhiteList.Contract.TransferOwnership(&_WhiteList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhiteList *WhiteListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhiteList.Contract.TransferOwnership(&_WhiteList.TransactOpts, newOwner)
}
