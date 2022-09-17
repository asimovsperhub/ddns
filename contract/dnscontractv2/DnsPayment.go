// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc

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

// DnsPaymentABI is the input ABI used to generate the binding from.
const DnsPaymentABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerC_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"addPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsPaymentName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"getPayment\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"getPaymentByAddr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferDaoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"}],\"name\":\"updateDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"updateName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DnsPayment is an auto generated Go binding around an Ethereum contract.
type DnsPayment struct {
	DnsPaymentCaller     // Read-only binding to the contract
	DnsPaymentTransactor // Write-only binding to the contract
	DnsPaymentFilterer   // Log filterer for contract events
}

// DnsPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsPaymentSession struct {
	Contract     *DnsPayment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsPaymentCallerSession struct {
	Contract *DnsPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DnsPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsPaymentTransactorSession struct {
	Contract     *DnsPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DnsPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsPaymentRaw struct {
	Contract *DnsPayment // Generic contract binding to access the raw methods on
}

// DnsPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsPaymentCallerRaw struct {
	Contract *DnsPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// DnsPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsPaymentTransactorRaw struct {
	Contract *DnsPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsPayment creates a new instance of DnsPayment, bound to a specific deployed contract.
func NewDnsPayment(address common.Address, backend bind.ContractBackend) (*DnsPayment, error) {
	contract, err := bindDnsPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsPayment{DnsPaymentCaller: DnsPaymentCaller{contract: contract}, DnsPaymentTransactor: DnsPaymentTransactor{contract: contract}, DnsPaymentFilterer: DnsPaymentFilterer{contract: contract}}, nil
}

// NewDnsPaymentCaller creates a new read-only instance of DnsPayment, bound to a specific deployed contract.
func NewDnsPaymentCaller(address common.Address, caller bind.ContractCaller) (*DnsPaymentCaller, error) {
	contract, err := bindDnsPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsPaymentCaller{contract: contract}, nil
}

// NewDnsPaymentTransactor creates a new write-only instance of DnsPayment, bound to a specific deployed contract.
func NewDnsPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsPaymentTransactor, error) {
	contract, err := bindDnsPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsPaymentTransactor{contract: contract}, nil
}

// NewDnsPaymentFilterer creates a new log filterer instance of DnsPayment, bound to a specific deployed contract.
func NewDnsPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsPaymentFilterer, error) {
	contract, err := bindDnsPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsPaymentFilterer{contract: contract}, nil
}

// bindDnsPayment binds a generic wrapper to an already deployed contract.
func bindDnsPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsPayment *DnsPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsPayment.Contract.DnsPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsPayment *DnsPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsPayment.Contract.DnsPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsPayment *DnsPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsPayment.Contract.DnsPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsPayment *DnsPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsPayment *DnsPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsPayment *DnsPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsPayment.Contract.contract.Transact(opts, method, params...)
}

// DnsPaymentName is a free data retrieval call binding the contract method 0x5a9aaaef.
//
// Solidity: function dnsPaymentName() view returns(string)
func (_DnsPayment *DnsPaymentCaller) DnsPaymentName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsPayment.contract.Call(opts, &out, "dnsPaymentName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DnsPaymentName is a free data retrieval call binding the contract method 0x5a9aaaef.
//
// Solidity: function dnsPaymentName() view returns(string)
func (_DnsPayment *DnsPaymentSession) DnsPaymentName() (string, error) {
	return _DnsPayment.Contract.DnsPaymentName(&_DnsPayment.CallOpts)
}

// DnsPaymentName is a free data retrieval call binding the contract method 0x5a9aaaef.
//
// Solidity: function dnsPaymentName() view returns(string)
func (_DnsPayment *DnsPaymentCallerSession) DnsPaymentName() (string, error) {
	return _DnsPayment.Contract.DnsPaymentName(&_DnsPayment.CallOpts)
}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 index_) view returns(string, string, address, uint8, bool)
func (_DnsPayment *DnsPaymentCaller) GetPayment(opts *bind.CallOpts, index_ *big.Int) (string, string, common.Address, uint8, bool, error) {
	var out []interface{}
	err := _DnsPayment.contract.Call(opts, &out, "getPayment", index_)

	if err != nil {
		return *new(string), *new(string), *new(common.Address), *new(uint8), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 index_) view returns(string, string, address, uint8, bool)
func (_DnsPayment *DnsPaymentSession) GetPayment(index_ *big.Int) (string, string, common.Address, uint8, bool, error) {
	return _DnsPayment.Contract.GetPayment(&_DnsPayment.CallOpts, index_)
}

// GetPayment is a free data retrieval call binding the contract method 0x3280a836.
//
// Solidity: function getPayment(uint256 index_) view returns(string, string, address, uint8, bool)
func (_DnsPayment *DnsPaymentCallerSession) GetPayment(index_ *big.Int) (string, string, common.Address, uint8, bool, error) {
	return _DnsPayment.Contract.GetPayment(&_DnsPayment.CallOpts, index_)
}

// GetPaymentByAddr is a free data retrieval call binding the contract method 0x32cd750e.
//
// Solidity: function getPaymentByAddr(address erc20Addr_) view returns(string, string, uint8, bool)
func (_DnsPayment *DnsPaymentCaller) GetPaymentByAddr(opts *bind.CallOpts, erc20Addr_ common.Address) (string, string, uint8, bool, error) {
	var out []interface{}
	err := _DnsPayment.contract.Call(opts, &out, "getPaymentByAddr", erc20Addr_)

	if err != nil {
		return *new(string), *new(string), *new(uint8), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetPaymentByAddr is a free data retrieval call binding the contract method 0x32cd750e.
//
// Solidity: function getPaymentByAddr(address erc20Addr_) view returns(string, string, uint8, bool)
func (_DnsPayment *DnsPaymentSession) GetPaymentByAddr(erc20Addr_ common.Address) (string, string, uint8, bool, error) {
	return _DnsPayment.Contract.GetPaymentByAddr(&_DnsPayment.CallOpts, erc20Addr_)
}

// GetPaymentByAddr is a free data retrieval call binding the contract method 0x32cd750e.
//
// Solidity: function getPaymentByAddr(address erc20Addr_) view returns(string, string, uint8, bool)
func (_DnsPayment *DnsPaymentCallerSession) GetPaymentByAddr(erc20Addr_ common.Address) (string, string, uint8, bool, error) {
	return _DnsPayment.Contract.GetPaymentByAddr(&_DnsPayment.CallOpts, erc20Addr_)
}

// GetPaymentCount is a free data retrieval call binding the contract method 0xaf32a03f.
//
// Solidity: function getPaymentCount() view returns(uint256)
func (_DnsPayment *DnsPaymentCaller) GetPaymentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsPayment.contract.Call(opts, &out, "getPaymentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentCount is a free data retrieval call binding the contract method 0xaf32a03f.
//
// Solidity: function getPaymentCount() view returns(uint256)
func (_DnsPayment *DnsPaymentSession) GetPaymentCount() (*big.Int, error) {
	return _DnsPayment.Contract.GetPaymentCount(&_DnsPayment.CallOpts)
}

// GetPaymentCount is a free data retrieval call binding the contract method 0xaf32a03f.
//
// Solidity: function getPaymentCount() view returns(uint256)
func (_DnsPayment *DnsPaymentCallerSession) GetPaymentCount() (*big.Int, error) {
	return _DnsPayment.Contract.GetPaymentCount(&_DnsPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPayment *DnsPaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsPayment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPayment *DnsPaymentSession) Owner() (common.Address, error) {
	return _DnsPayment.Contract.Owner(&_DnsPayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPayment *DnsPaymentCallerSession) Owner() (common.Address, error) {
	return _DnsPayment.Contract.Owner(&_DnsPayment.CallOpts)
}

// AddPayment is a paid mutator transaction binding the contract method 0x8f2a1aaf.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentTransactor) AddPayment(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal_ uint8, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "addPayment", erc20Addr_, decimal_, name_, symbol_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x8f2a1aaf.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.Contract.AddPayment(&_DnsPayment.TransactOpts, erc20Addr_, decimal_, name_, symbol_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x8f2a1aaf.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentTransactorSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.Contract.AddPayment(&_DnsPayment.TransactOpts, erc20Addr_, decimal_, name_, symbol_)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentTransactor) Disable(opts *bind.TransactOpts, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "disable", erc20Addr_)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentSession) Disable(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.Disable(&_DnsPayment.TransactOpts, erc20Addr_)
}

// Disable is a paid mutator transaction binding the contract method 0xe6c09edf.
//
// Solidity: function disable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentTransactorSession) Disable(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.Disable(&_DnsPayment.TransactOpts, erc20Addr_)
}

// Enable is a paid mutator transaction binding the contract method 0x5bfa1b68.
//
// Solidity: function enable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentTransactor) Enable(opts *bind.TransactOpts, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "enable", erc20Addr_)
}

// Enable is a paid mutator transaction binding the contract method 0x5bfa1b68.
//
// Solidity: function enable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentSession) Enable(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.Enable(&_DnsPayment.TransactOpts, erc20Addr_)
}

// Enable is a paid mutator transaction binding the contract method 0x5bfa1b68.
//
// Solidity: function enable(address erc20Addr_) returns()
func (_DnsPayment *DnsPaymentTransactorSession) Enable(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.Enable(&_DnsPayment.TransactOpts, erc20Addr_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPayment *DnsPaymentTransactor) TransferDaoOwner(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "transferDaoOwner", newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPayment *DnsPaymentSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.TransferDaoOwner(&_DnsPayment.TransactOpts, newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPayment *DnsPaymentTransactorSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPayment.Contract.TransferDaoOwner(&_DnsPayment.TransactOpts, newOwner_)
}

// UpdateDecimals is a paid mutator transaction binding the contract method 0xa03d2c19.
//
// Solidity: function updateDecimals(address erc20Addr_, uint8 decimal) returns()
func (_DnsPayment *DnsPaymentTransactor) UpdateDecimals(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal uint8) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "updateDecimals", erc20Addr_, decimal)
}

// UpdateDecimals is a paid mutator transaction binding the contract method 0xa03d2c19.
//
// Solidity: function updateDecimals(address erc20Addr_, uint8 decimal) returns()
func (_DnsPayment *DnsPaymentSession) UpdateDecimals(erc20Addr_ common.Address, decimal uint8) (*types.Transaction, error) {
	return _DnsPayment.Contract.UpdateDecimals(&_DnsPayment.TransactOpts, erc20Addr_, decimal)
}

// UpdateDecimals is a paid mutator transaction binding the contract method 0xa03d2c19.
//
// Solidity: function updateDecimals(address erc20Addr_, uint8 decimal) returns()
func (_DnsPayment *DnsPaymentTransactorSession) UpdateDecimals(erc20Addr_ common.Address, decimal uint8) (*types.Transaction, error) {
	return _DnsPayment.Contract.UpdateDecimals(&_DnsPayment.TransactOpts, erc20Addr_, decimal)
}

// UpdateName is a paid mutator transaction binding the contract method 0xc14e639b.
//
// Solidity: function updateName(address erc20Addr_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentTransactor) UpdateName(opts *bind.TransactOpts, erc20Addr_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.contract.Transact(opts, "updateName", erc20Addr_, name_, symbol_)
}

// UpdateName is a paid mutator transaction binding the contract method 0xc14e639b.
//
// Solidity: function updateName(address erc20Addr_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentSession) UpdateName(erc20Addr_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.Contract.UpdateName(&_DnsPayment.TransactOpts, erc20Addr_, name_, symbol_)
}

// UpdateName is a paid mutator transaction binding the contract method 0xc14e639b.
//
// Solidity: function updateName(address erc20Addr_, string name_, string symbol_) returns()
func (_DnsPayment *DnsPaymentTransactorSession) UpdateName(erc20Addr_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsPayment.Contract.UpdateName(&_DnsPayment.TransactOpts, erc20Addr_, name_, symbol_)
}
