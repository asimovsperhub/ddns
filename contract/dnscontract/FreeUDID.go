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

// FreeUDIDABI is the input ABI used to generate the binding from.
const FreeUDIDABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LastTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tm\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveUDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetAmountPerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TransferUDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UDIDToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FreeUDID is an auto generated Go binding around an Ethereum contract.
type FreeUDID struct {
	FreeUDIDCaller     // Read-only binding to the contract
	FreeUDIDTransactor // Write-only binding to the contract
	FreeUDIDFilterer   // Log filterer for contract events
}

// FreeUDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreeUDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeUDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreeUDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeUDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreeUDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeUDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreeUDIDSession struct {
	Contract     *FreeUDID         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreeUDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreeUDIDCallerSession struct {
	Contract *FreeUDIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FreeUDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreeUDIDTransactorSession struct {
	Contract     *FreeUDIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FreeUDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreeUDIDRaw struct {
	Contract *FreeUDID // Generic contract binding to access the raw methods on
}

// FreeUDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreeUDIDCallerRaw struct {
	Contract *FreeUDIDCaller // Generic read-only contract binding to access the raw methods on
}

// FreeUDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreeUDIDTransactorRaw struct {
	Contract *FreeUDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreeUDID creates a new instance of FreeUDID, bound to a specific deployed contract.
func NewFreeUDID(address common.Address, backend bind.ContractBackend) (*FreeUDID, error) {
	contract, err := bindFreeUDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreeUDID{FreeUDIDCaller: FreeUDIDCaller{contract: contract}, FreeUDIDTransactor: FreeUDIDTransactor{contract: contract}, FreeUDIDFilterer: FreeUDIDFilterer{contract: contract}}, nil
}

// NewFreeUDIDCaller creates a new read-only instance of FreeUDID, bound to a specific deployed contract.
func NewFreeUDIDCaller(address common.Address, caller bind.ContractCaller) (*FreeUDIDCaller, error) {
	contract, err := bindFreeUDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreeUDIDCaller{contract: contract}, nil
}

// NewFreeUDIDTransactor creates a new write-only instance of FreeUDID, bound to a specific deployed contract.
func NewFreeUDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*FreeUDIDTransactor, error) {
	contract, err := bindFreeUDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreeUDIDTransactor{contract: contract}, nil
}

// NewFreeUDIDFilterer creates a new log filterer instance of FreeUDID, bound to a specific deployed contract.
func NewFreeUDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*FreeUDIDFilterer, error) {
	contract, err := bindFreeUDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreeUDIDFilterer{contract: contract}, nil
}

// bindFreeUDID binds a generic wrapper to an already deployed contract.
func bindFreeUDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FreeUDIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeUDID *FreeUDIDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeUDID.Contract.FreeUDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeUDID *FreeUDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeUDID.Contract.FreeUDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeUDID *FreeUDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeUDID.Contract.FreeUDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeUDID *FreeUDIDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeUDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeUDID *FreeUDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeUDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeUDID *FreeUDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeUDID.Contract.contract.Transact(opts, method, params...)
}

// LastTime is a free data retrieval call binding the contract method 0x1a049cf8.
//
// Solidity: function LastTime() view returns(uint256 tm)
func (_FreeUDID *FreeUDIDCaller) LastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeUDID.contract.Call(opts, &out, "LastTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTime is a free data retrieval call binding the contract method 0x1a049cf8.
//
// Solidity: function LastTime() view returns(uint256 tm)
func (_FreeUDID *FreeUDIDSession) LastTime() (*big.Int, error) {
	return _FreeUDID.Contract.LastTime(&_FreeUDID.CallOpts)
}

// LastTime is a free data retrieval call binding the contract method 0x1a049cf8.
//
// Solidity: function LastTime() view returns(uint256 tm)
func (_FreeUDID *FreeUDIDCallerSession) LastTime() (*big.Int, error) {
	return _FreeUDID.Contract.LastTime(&_FreeUDID.CallOpts)
}

// UDIDToken is a free data retrieval call binding the contract method 0x7d993f7c.
//
// Solidity: function UDIDToken() view returns(address)
func (_FreeUDID *FreeUDIDCaller) UDIDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeUDID.contract.Call(opts, &out, "UDIDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UDIDToken is a free data retrieval call binding the contract method 0x7d993f7c.
//
// Solidity: function UDIDToken() view returns(address)
func (_FreeUDID *FreeUDIDSession) UDIDToken() (common.Address, error) {
	return _FreeUDID.Contract.UDIDToken(&_FreeUDID.CallOpts)
}

// UDIDToken is a free data retrieval call binding the contract method 0x7d993f7c.
//
// Solidity: function UDIDToken() view returns(address)
func (_FreeUDID *FreeUDIDCallerSession) UDIDToken() (common.Address, error) {
	return _FreeUDID.Contract.UDIDToken(&_FreeUDID.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeUDID *FreeUDIDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeUDID.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeUDID *FreeUDIDSession) Owner() (common.Address, error) {
	return _FreeUDID.Contract.Owner(&_FreeUDID.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeUDID *FreeUDIDCallerSession) Owner() (common.Address, error) {
	return _FreeUDID.Contract.Owner(&_FreeUDID.CallOpts)
}

// RetrieveUDID is a paid mutator transaction binding the contract method 0x4ccb280a.
//
// Solidity: function RetrieveUDID() returns()
func (_FreeUDID *FreeUDIDTransactor) RetrieveUDID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeUDID.contract.Transact(opts, "RetrieveUDID")
}

// RetrieveUDID is a paid mutator transaction binding the contract method 0x4ccb280a.
//
// Solidity: function RetrieveUDID() returns()
func (_FreeUDID *FreeUDIDSession) RetrieveUDID() (*types.Transaction, error) {
	return _FreeUDID.Contract.RetrieveUDID(&_FreeUDID.TransactOpts)
}

// RetrieveUDID is a paid mutator transaction binding the contract method 0x4ccb280a.
//
// Solidity: function RetrieveUDID() returns()
func (_FreeUDID *FreeUDIDTransactorSession) RetrieveUDID() (*types.Transaction, error) {
	return _FreeUDID.Contract.RetrieveUDID(&_FreeUDID.TransactOpts)
}

// SetAmountPerDay is a paid mutator transaction binding the contract method 0x4ed03837.
//
// Solidity: function SetAmountPerDay(uint256 amount) returns()
func (_FreeUDID *FreeUDIDTransactor) SetAmountPerDay(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FreeUDID.contract.Transact(opts, "SetAmountPerDay", amount)
}

// SetAmountPerDay is a paid mutator transaction binding the contract method 0x4ed03837.
//
// Solidity: function SetAmountPerDay(uint256 amount) returns()
func (_FreeUDID *FreeUDIDSession) SetAmountPerDay(amount *big.Int) (*types.Transaction, error) {
	return _FreeUDID.Contract.SetAmountPerDay(&_FreeUDID.TransactOpts, amount)
}

// SetAmountPerDay is a paid mutator transaction binding the contract method 0x4ed03837.
//
// Solidity: function SetAmountPerDay(uint256 amount) returns()
func (_FreeUDID *FreeUDIDTransactorSession) SetAmountPerDay(amount *big.Int) (*types.Transaction, error) {
	return _FreeUDID.Contract.SetAmountPerDay(&_FreeUDID.TransactOpts, amount)
}

// TransferUDID is a paid mutator transaction binding the contract method 0x9afe62de.
//
// Solidity: function TransferUDID() returns()
func (_FreeUDID *FreeUDIDTransactor) TransferUDID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeUDID.contract.Transact(opts, "TransferUDID")
}

// TransferUDID is a paid mutator transaction binding the contract method 0x9afe62de.
//
// Solidity: function TransferUDID() returns()
func (_FreeUDID *FreeUDIDSession) TransferUDID() (*types.Transaction, error) {
	return _FreeUDID.Contract.TransferUDID(&_FreeUDID.TransactOpts)
}

// TransferUDID is a paid mutator transaction binding the contract method 0x9afe62de.
//
// Solidity: function TransferUDID() returns()
func (_FreeUDID *FreeUDIDTransactorSession) TransferUDID() (*types.Transaction, error) {
	return _FreeUDID.Contract.TransferUDID(&_FreeUDID.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeUDID *FreeUDIDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FreeUDID.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeUDID *FreeUDIDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeUDID.Contract.TransferOwnership(&_FreeUDID.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeUDID *FreeUDIDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeUDID.Contract.TransferOwnership(&_FreeUDID.TransactOpts, newOwner)
}
