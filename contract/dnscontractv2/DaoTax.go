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

// DaoTaxMetaData contains all meta data concerning the DaoTax contract.
var DaoTaxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerC_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"taxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percentBase\",\"type\":\"uint256\"}],\"name\":\"EvAddTax\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"EvAddTaxLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"EvDelTaxLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"taxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultTax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percentBase\",\"type\":\"uint256\"}],\"name\":\"EvUpdateTax\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"taxType_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"defaultTax_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentBase_\",\"type\":\"uint256\"}],\"name\":\"AddTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tax_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent_\",\"type\":\"uint256\"}],\"name\":\"AddTaxLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"name\":\"DelTaxLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"taxType_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"defaultTax_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentBase_\",\"type\":\"uint256\"}],\"name\":\"UpdateTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"name\":\"tax\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferDaoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaoTaxABI is the input ABI used to generate the binding from.
// Deprecated: Use DaoTaxMetaData.ABI instead.
var DaoTaxABI = DaoTaxMetaData.ABI

// DaoTax is an auto generated Go binding around an Ethereum contract.
type DaoTax struct {
	DaoTaxCaller     // Read-only binding to the contract
	DaoTaxTransactor // Write-only binding to the contract
	DaoTaxFilterer   // Log filterer for contract events
}

// DaoTaxCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaoTaxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTaxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaoTaxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTaxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaoTaxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaoTaxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaoTaxSession struct {
	Contract     *DaoTax           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoTaxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaoTaxCallerSession struct {
	Contract *DaoTaxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DaoTaxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaoTaxTransactorSession struct {
	Contract     *DaoTaxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaoTaxRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaoTaxRaw struct {
	Contract *DaoTax // Generic contract binding to access the raw methods on
}

// DaoTaxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaoTaxCallerRaw struct {
	Contract *DaoTaxCaller // Generic read-only contract binding to access the raw methods on
}

// DaoTaxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaoTaxTransactorRaw struct {
	Contract *DaoTaxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaoTax creates a new instance of DaoTax, bound to a specific deployed contract.
func NewDaoTax(address common.Address, backend bind.ContractBackend) (*DaoTax, error) {
	contract, err := bindDaoTax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaoTax{DaoTaxCaller: DaoTaxCaller{contract: contract}, DaoTaxTransactor: DaoTaxTransactor{contract: contract}, DaoTaxFilterer: DaoTaxFilterer{contract: contract}}, nil
}

// NewDaoTaxCaller creates a new read-only instance of DaoTax, bound to a specific deployed contract.
func NewDaoTaxCaller(address common.Address, caller bind.ContractCaller) (*DaoTaxCaller, error) {
	contract, err := bindDaoTax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaoTaxCaller{contract: contract}, nil
}

// NewDaoTaxTransactor creates a new write-only instance of DaoTax, bound to a specific deployed contract.
func NewDaoTaxTransactor(address common.Address, transactor bind.ContractTransactor) (*DaoTaxTransactor, error) {
	contract, err := bindDaoTax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaoTaxTransactor{contract: contract}, nil
}

// NewDaoTaxFilterer creates a new log filterer instance of DaoTax, bound to a specific deployed contract.
func NewDaoTaxFilterer(address common.Address, filterer bind.ContractFilterer) (*DaoTaxFilterer, error) {
	contract, err := bindDaoTax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaoTaxFilterer{contract: contract}, nil
}

// bindDaoTax binds a generic wrapper to an already deployed contract.
func bindDaoTax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaoTaxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoTax *DaoTaxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaoTax.Contract.DaoTaxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoTax *DaoTaxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoTax.Contract.DaoTaxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoTax *DaoTaxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoTax.Contract.DaoTaxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaoTax *DaoTaxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaoTax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaoTax *DaoTaxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaoTax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaoTax *DaoTaxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaoTax.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoTax *DaoTaxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DaoTax.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoTax *DaoTaxSession) Name() (string, error) {
	return _DaoTax.Contract.Name(&_DaoTax.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DaoTax *DaoTaxCallerSession) Name() (string, error) {
	return _DaoTax.Contract.Name(&_DaoTax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DaoTax *DaoTaxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaoTax.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DaoTax *DaoTaxSession) Owner() (common.Address, error) {
	return _DaoTax.Contract.Owner(&_DaoTax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DaoTax *DaoTaxCallerSession) Owner() (common.Address, error) {
	return _DaoTax.Contract.Owner(&_DaoTax.CallOpts)
}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DaoTax *DaoTaxCaller) PaymentC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaoTax.contract.Call(opts, &out, "paymentC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DaoTax *DaoTaxSession) PaymentC() (common.Address, error) {
	return _DaoTax.Contract.PaymentC(&_DaoTax.CallOpts)
}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DaoTax *DaoTaxCallerSession) PaymentC() (common.Address, error) {
	return _DaoTax.Contract.PaymentC(&_DaoTax.CallOpts)
}

// Tax is a free data retrieval call binding the contract method 0xd9b57789.
//
// Solidity: function tax(address erc20Addr_, uint256 len_) view returns(uint8, uint256, uint256, uint256)
func (_DaoTax *DaoTaxCaller) Tax(opts *bind.CallOpts, erc20Addr_ common.Address, len_ *big.Int) (uint8, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _DaoTax.contract.Call(opts, &out, "tax", erc20Addr_, len_)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// Tax is a free data retrieval call binding the contract method 0xd9b57789.
//
// Solidity: function tax(address erc20Addr_, uint256 len_) view returns(uint8, uint256, uint256, uint256)
func (_DaoTax *DaoTaxSession) Tax(erc20Addr_ common.Address, len_ *big.Int) (uint8, *big.Int, *big.Int, *big.Int, error) {
	return _DaoTax.Contract.Tax(&_DaoTax.CallOpts, erc20Addr_, len_)
}

// Tax is a free data retrieval call binding the contract method 0xd9b57789.
//
// Solidity: function tax(address erc20Addr_, uint256 len_) view returns(uint8, uint256, uint256, uint256)
func (_DaoTax *DaoTaxCallerSession) Tax(erc20Addr_ common.Address, len_ *big.Int) (uint8, *big.Int, *big.Int, *big.Int, error) {
	return _DaoTax.Contract.Tax(&_DaoTax.CallOpts, erc20Addr_, len_)
}

// AddTax is a paid mutator transaction binding the contract method 0xaef0c801.
//
// Solidity: function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxTransactor) AddTax(opts *bind.TransactOpts, erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.contract.Transact(opts, "AddTax", erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// AddTax is a paid mutator transaction binding the contract method 0xaef0c801.
//
// Solidity: function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxSession) AddTax(erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.AddTax(&_DaoTax.TransactOpts, erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// AddTax is a paid mutator transaction binding the contract method 0xaef0c801.
//
// Solidity: function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxTransactorSession) AddTax(erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.AddTax(&_DaoTax.TransactOpts, erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// AddTaxLen is a paid mutator transaction binding the contract method 0x5512ee53.
//
// Solidity: function AddTaxLen(address erc20Addr_, uint256 len_, uint256 tax_, uint256 percent_) returns()
func (_DaoTax *DaoTaxTransactor) AddTaxLen(opts *bind.TransactOpts, erc20Addr_ common.Address, len_ *big.Int, tax_ *big.Int, percent_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.contract.Transact(opts, "AddTaxLen", erc20Addr_, len_, tax_, percent_)
}

// AddTaxLen is a paid mutator transaction binding the contract method 0x5512ee53.
//
// Solidity: function AddTaxLen(address erc20Addr_, uint256 len_, uint256 tax_, uint256 percent_) returns()
func (_DaoTax *DaoTaxSession) AddTaxLen(erc20Addr_ common.Address, len_ *big.Int, tax_ *big.Int, percent_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.AddTaxLen(&_DaoTax.TransactOpts, erc20Addr_, len_, tax_, percent_)
}

// AddTaxLen is a paid mutator transaction binding the contract method 0x5512ee53.
//
// Solidity: function AddTaxLen(address erc20Addr_, uint256 len_, uint256 tax_, uint256 percent_) returns()
func (_DaoTax *DaoTaxTransactorSession) AddTaxLen(erc20Addr_ common.Address, len_ *big.Int, tax_ *big.Int, percent_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.AddTaxLen(&_DaoTax.TransactOpts, erc20Addr_, len_, tax_, percent_)
}

// DelTaxLen is a paid mutator transaction binding the contract method 0x8c0447de.
//
// Solidity: function DelTaxLen(address erc20Addr_, uint256 len_) returns()
func (_DaoTax *DaoTaxTransactor) DelTaxLen(opts *bind.TransactOpts, erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.contract.Transact(opts, "DelTaxLen", erc20Addr_, len_)
}

// DelTaxLen is a paid mutator transaction binding the contract method 0x8c0447de.
//
// Solidity: function DelTaxLen(address erc20Addr_, uint256 len_) returns()
func (_DaoTax *DaoTaxSession) DelTaxLen(erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.DelTaxLen(&_DaoTax.TransactOpts, erc20Addr_, len_)
}

// DelTaxLen is a paid mutator transaction binding the contract method 0x8c0447de.
//
// Solidity: function DelTaxLen(address erc20Addr_, uint256 len_) returns()
func (_DaoTax *DaoTaxTransactorSession) DelTaxLen(erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.DelTaxLen(&_DaoTax.TransactOpts, erc20Addr_, len_)
}

// UpdateTax is a paid mutator transaction binding the contract method 0x9343dee1.
//
// Solidity: function UpdateTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxTransactor) UpdateTax(opts *bind.TransactOpts, erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.contract.Transact(opts, "UpdateTax", erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// UpdateTax is a paid mutator transaction binding the contract method 0x9343dee1.
//
// Solidity: function UpdateTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxSession) UpdateTax(erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.UpdateTax(&_DaoTax.TransactOpts, erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// UpdateTax is a paid mutator transaction binding the contract method 0x9343dee1.
//
// Solidity: function UpdateTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_, uint256 percentBase_) returns()
func (_DaoTax *DaoTaxTransactorSession) UpdateTax(erc20Addr_ common.Address, taxType_ uint8, defaultTax_ *big.Int, percent_ *big.Int, percentBase_ *big.Int) (*types.Transaction, error) {
	return _DaoTax.Contract.UpdateTax(&_DaoTax.TransactOpts, erc20Addr_, taxType_, defaultTax_, percent_, percentBase_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DaoTax *DaoTaxTransactor) TransferDaoOwner(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DaoTax.contract.Transact(opts, "transferDaoOwner", newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DaoTax *DaoTaxSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DaoTax.Contract.TransferDaoOwner(&_DaoTax.TransactOpts, newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DaoTax *DaoTaxTransactorSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DaoTax.Contract.TransferDaoOwner(&_DaoTax.TransactOpts, newOwner_)
}

// DaoTaxEvAddTaxIterator is returned from FilterEvAddTax and is used to iterate over the raw logs and unpacked data for EvAddTax events raised by the DaoTax contract.
type DaoTaxEvAddTaxIterator struct {
	Event *DaoTaxEvAddTax // Event containing the contract specifics and raw log

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
func (it *DaoTaxEvAddTaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoTaxEvAddTax)
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
		it.Event = new(DaoTaxEvAddTax)
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
func (it *DaoTaxEvAddTaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoTaxEvAddTaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoTaxEvAddTax represents a EvAddTax event raised by the DaoTax contract.
type DaoTaxEvAddTax struct {
	Erc20Addr   common.Address
	TaxType     uint8
	DefaultTax  *big.Int
	Percent     *big.Int
	PercentBase *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEvAddTax is a free log retrieval operation binding the contract event 0x9b81cd48415e4a818c0933789f7eabbd5a59e0a2eab5043b939c892c7e5c8daf.
//
// Solidity: event EvAddTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) FilterEvAddTax(opts *bind.FilterOpts) (*DaoTaxEvAddTaxIterator, error) {

	logs, sub, err := _DaoTax.contract.FilterLogs(opts, "EvAddTax")
	if err != nil {
		return nil, err
	}
	return &DaoTaxEvAddTaxIterator{contract: _DaoTax.contract, event: "EvAddTax", logs: logs, sub: sub}, nil
}

// WatchEvAddTax is a free log subscription operation binding the contract event 0x9b81cd48415e4a818c0933789f7eabbd5a59e0a2eab5043b939c892c7e5c8daf.
//
// Solidity: event EvAddTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) WatchEvAddTax(opts *bind.WatchOpts, sink chan<- *DaoTaxEvAddTax) (event.Subscription, error) {

	logs, sub, err := _DaoTax.contract.WatchLogs(opts, "EvAddTax")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoTaxEvAddTax)
				if err := _DaoTax.contract.UnpackLog(event, "EvAddTax", log); err != nil {
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

// ParseEvAddTax is a log parse operation binding the contract event 0x9b81cd48415e4a818c0933789f7eabbd5a59e0a2eab5043b939c892c7e5c8daf.
//
// Solidity: event EvAddTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) ParseEvAddTax(log types.Log) (*DaoTaxEvAddTax, error) {
	event := new(DaoTaxEvAddTax)
	if err := _DaoTax.contract.UnpackLog(event, "EvAddTax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoTaxEvAddTaxLenIterator is returned from FilterEvAddTaxLen and is used to iterate over the raw logs and unpacked data for EvAddTaxLen events raised by the DaoTax contract.
type DaoTaxEvAddTaxLenIterator struct {
	Event *DaoTaxEvAddTaxLen // Event containing the contract specifics and raw log

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
func (it *DaoTaxEvAddTaxLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoTaxEvAddTaxLen)
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
		it.Event = new(DaoTaxEvAddTaxLen)
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
func (it *DaoTaxEvAddTaxLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoTaxEvAddTaxLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoTaxEvAddTaxLen represents a EvAddTaxLen event raised by the DaoTax contract.
type DaoTaxEvAddTaxLen struct {
	Erc20Addr common.Address
	Len       *big.Int
	Tax       *big.Int
	Percent   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvAddTaxLen is a free log retrieval operation binding the contract event 0x429dac2bd0938e8c970df7bba1f17b3658c3d98aa2a955e364ec6d71070d3cad.
//
// Solidity: event EvAddTaxLen(address erc20Addr, uint256 len, uint256 tax, uint256 percent)
func (_DaoTax *DaoTaxFilterer) FilterEvAddTaxLen(opts *bind.FilterOpts) (*DaoTaxEvAddTaxLenIterator, error) {

	logs, sub, err := _DaoTax.contract.FilterLogs(opts, "EvAddTaxLen")
	if err != nil {
		return nil, err
	}
	return &DaoTaxEvAddTaxLenIterator{contract: _DaoTax.contract, event: "EvAddTaxLen", logs: logs, sub: sub}, nil
}

// WatchEvAddTaxLen is a free log subscription operation binding the contract event 0x429dac2bd0938e8c970df7bba1f17b3658c3d98aa2a955e364ec6d71070d3cad.
//
// Solidity: event EvAddTaxLen(address erc20Addr, uint256 len, uint256 tax, uint256 percent)
func (_DaoTax *DaoTaxFilterer) WatchEvAddTaxLen(opts *bind.WatchOpts, sink chan<- *DaoTaxEvAddTaxLen) (event.Subscription, error) {

	logs, sub, err := _DaoTax.contract.WatchLogs(opts, "EvAddTaxLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoTaxEvAddTaxLen)
				if err := _DaoTax.contract.UnpackLog(event, "EvAddTaxLen", log); err != nil {
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

// ParseEvAddTaxLen is a log parse operation binding the contract event 0x429dac2bd0938e8c970df7bba1f17b3658c3d98aa2a955e364ec6d71070d3cad.
//
// Solidity: event EvAddTaxLen(address erc20Addr, uint256 len, uint256 tax, uint256 percent)
func (_DaoTax *DaoTaxFilterer) ParseEvAddTaxLen(log types.Log) (*DaoTaxEvAddTaxLen, error) {
	event := new(DaoTaxEvAddTaxLen)
	if err := _DaoTax.contract.UnpackLog(event, "EvAddTaxLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoTaxEvDelTaxLenIterator is returned from FilterEvDelTaxLen and is used to iterate over the raw logs and unpacked data for EvDelTaxLen events raised by the DaoTax contract.
type DaoTaxEvDelTaxLenIterator struct {
	Event *DaoTaxEvDelTaxLen // Event containing the contract specifics and raw log

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
func (it *DaoTaxEvDelTaxLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoTaxEvDelTaxLen)
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
		it.Event = new(DaoTaxEvDelTaxLen)
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
func (it *DaoTaxEvDelTaxLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoTaxEvDelTaxLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoTaxEvDelTaxLen represents a EvDelTaxLen event raised by the DaoTax contract.
type DaoTaxEvDelTaxLen struct {
	Erc20Addr common.Address
	Len       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvDelTaxLen is a free log retrieval operation binding the contract event 0xdf67b60dfe7bb038a11487226561a702e63d004eb6a394b82f89be6e8e4b116a.
//
// Solidity: event EvDelTaxLen(address erc20Addr, uint256 len)
func (_DaoTax *DaoTaxFilterer) FilterEvDelTaxLen(opts *bind.FilterOpts) (*DaoTaxEvDelTaxLenIterator, error) {

	logs, sub, err := _DaoTax.contract.FilterLogs(opts, "EvDelTaxLen")
	if err != nil {
		return nil, err
	}
	return &DaoTaxEvDelTaxLenIterator{contract: _DaoTax.contract, event: "EvDelTaxLen", logs: logs, sub: sub}, nil
}

// WatchEvDelTaxLen is a free log subscription operation binding the contract event 0xdf67b60dfe7bb038a11487226561a702e63d004eb6a394b82f89be6e8e4b116a.
//
// Solidity: event EvDelTaxLen(address erc20Addr, uint256 len)
func (_DaoTax *DaoTaxFilterer) WatchEvDelTaxLen(opts *bind.WatchOpts, sink chan<- *DaoTaxEvDelTaxLen) (event.Subscription, error) {

	logs, sub, err := _DaoTax.contract.WatchLogs(opts, "EvDelTaxLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoTaxEvDelTaxLen)
				if err := _DaoTax.contract.UnpackLog(event, "EvDelTaxLen", log); err != nil {
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

// ParseEvDelTaxLen is a log parse operation binding the contract event 0xdf67b60dfe7bb038a11487226561a702e63d004eb6a394b82f89be6e8e4b116a.
//
// Solidity: event EvDelTaxLen(address erc20Addr, uint256 len)
func (_DaoTax *DaoTaxFilterer) ParseEvDelTaxLen(log types.Log) (*DaoTaxEvDelTaxLen, error) {
	event := new(DaoTaxEvDelTaxLen)
	if err := _DaoTax.contract.UnpackLog(event, "EvDelTaxLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaoTaxEvUpdateTaxIterator is returned from FilterEvUpdateTax and is used to iterate over the raw logs and unpacked data for EvUpdateTax events raised by the DaoTax contract.
type DaoTaxEvUpdateTaxIterator struct {
	Event *DaoTaxEvUpdateTax // Event containing the contract specifics and raw log

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
func (it *DaoTaxEvUpdateTaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaoTaxEvUpdateTax)
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
		it.Event = new(DaoTaxEvUpdateTax)
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
func (it *DaoTaxEvUpdateTaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaoTaxEvUpdateTaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaoTaxEvUpdateTax represents a EvUpdateTax event raised by the DaoTax contract.
type DaoTaxEvUpdateTax struct {
	Erc20Addr   common.Address
	TaxType     uint8
	DefaultTax  *big.Int
	Percent     *big.Int
	PercentBase *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEvUpdateTax is a free log retrieval operation binding the contract event 0x4e1e5853ab0da09d208861b9fc3b3b0800b3ddfbc243c040d66d0d23158b5a24.
//
// Solidity: event EvUpdateTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) FilterEvUpdateTax(opts *bind.FilterOpts) (*DaoTaxEvUpdateTaxIterator, error) {

	logs, sub, err := _DaoTax.contract.FilterLogs(opts, "EvUpdateTax")
	if err != nil {
		return nil, err
	}
	return &DaoTaxEvUpdateTaxIterator{contract: _DaoTax.contract, event: "EvUpdateTax", logs: logs, sub: sub}, nil
}

// WatchEvUpdateTax is a free log subscription operation binding the contract event 0x4e1e5853ab0da09d208861b9fc3b3b0800b3ddfbc243c040d66d0d23158b5a24.
//
// Solidity: event EvUpdateTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) WatchEvUpdateTax(opts *bind.WatchOpts, sink chan<- *DaoTaxEvUpdateTax) (event.Subscription, error) {

	logs, sub, err := _DaoTax.contract.WatchLogs(opts, "EvUpdateTax")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaoTaxEvUpdateTax)
				if err := _DaoTax.contract.UnpackLog(event, "EvUpdateTax", log); err != nil {
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

// ParseEvUpdateTax is a log parse operation binding the contract event 0x4e1e5853ab0da09d208861b9fc3b3b0800b3ddfbc243c040d66d0d23158b5a24.
//
// Solidity: event EvUpdateTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase)
func (_DaoTax *DaoTaxFilterer) ParseEvUpdateTax(log types.Log) (*DaoTaxEvUpdateTax, error) {
	event := new(DaoTaxEvUpdateTax)
	if err := _DaoTax.contract.UnpackLog(event, "EvUpdateTax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
