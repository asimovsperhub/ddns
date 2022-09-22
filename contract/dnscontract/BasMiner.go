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

// BasMinerMetaData contains all meta data concerning the BasMiner contract.
var BasMinerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMiner\",\"type\":\"address\"}],\"name\":\"MinerReplace\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"drawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amout\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllMainNodeAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MainNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MainNodeSize\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"addMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"removeMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldM\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newM\",\"type\":\"address\"}],\"name\":\"replaceMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"}],\"name\":\"subAllocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasMinerABI is the input ABI used to generate the binding from.
// Deprecated: Use BasMinerMetaData.ABI instead.
var BasMinerABI = BasMinerMetaData.ABI

// BasMiner is an auto generated Go binding around an Ethereum contract.
type BasMiner struct {
	BasMinerCaller     // Read-only binding to the contract
	BasMinerTransactor // Write-only binding to the contract
	BasMinerFilterer   // Log filterer for contract events
}

// BasMinerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMinerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMinerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMinerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMinerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMinerSession struct {
	Contract     *BasMiner         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMinerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMinerCallerSession struct {
	Contract *BasMinerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasMinerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMinerTransactorSession struct {
	Contract     *BasMinerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasMinerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMinerRaw struct {
	Contract *BasMiner // Generic contract binding to access the raw methods on
}

// BasMinerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMinerCallerRaw struct {
	Contract *BasMinerCaller // Generic read-only contract binding to access the raw methods on
}

// BasMinerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMinerTransactorRaw struct {
	Contract *BasMinerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMiner creates a new instance of BasMiner, bound to a specific deployed contract.
func NewBasMiner(address common.Address, backend bind.ContractBackend) (*BasMiner, error) {
	contract, err := bindBasMiner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMiner{BasMinerCaller: BasMinerCaller{contract: contract}, BasMinerTransactor: BasMinerTransactor{contract: contract}, BasMinerFilterer: BasMinerFilterer{contract: contract}}, nil
}

// NewBasMinerCaller creates a new read-only instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerCaller(address common.Address, caller bind.ContractCaller) (*BasMinerCaller, error) {
	contract, err := bindBasMiner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMinerCaller{contract: contract}, nil
}

// NewBasMinerTransactor creates a new write-only instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMinerTransactor, error) {
	contract, err := bindBasMiner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMinerTransactor{contract: contract}, nil
}

// NewBasMinerFilterer creates a new log filterer instance of BasMiner, bound to a specific deployed contract.
func NewBasMinerFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMinerFilterer, error) {
	contract, err := bindBasMiner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMinerFilterer{contract: contract}, nil
}

// bindBasMiner binds a generic wrapper to an already deployed contract.
func bindBasMiner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMinerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMiner *BasMinerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMiner.Contract.BasMinerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMiner *BasMinerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.Contract.BasMinerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMiner *BasMinerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMiner.Contract.BasMinerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMiner *BasMinerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMiner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMiner *BasMinerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMiner *BasMinerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMiner.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMiner *BasMinerCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMiner *BasMinerSession) DAOAddress() (common.Address, error) {
	return _BasMiner.Contract.DAOAddress(&_BasMiner.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMiner *BasMinerCallerSession) DAOAddress() (common.Address, error) {
	return _BasMiner.Contract.DAOAddress(&_BasMiner.CallOpts)
}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() view returns(address[])
func (_BasMiner *BasMinerCaller) GetAllMainNodeAddress(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "GetAllMainNodeAddress")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() view returns(address[])
func (_BasMiner *BasMinerSession) GetAllMainNodeAddress() ([]common.Address, error) {
	return _BasMiner.Contract.GetAllMainNodeAddress(&_BasMiner.CallOpts)
}

// GetAllMainNodeAddress is a free data retrieval call binding the contract method 0x179e0f00.
//
// Solidity: function GetAllMainNodeAddress() view returns(address[])
func (_BasMiner *BasMinerCallerSession) GetAllMainNodeAddress() ([]common.Address, error) {
	return _BasMiner.Contract.GetAllMainNodeAddress(&_BasMiner.CallOpts)
}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) view returns(address)
func (_BasMiner *BasMinerCaller) MainNode(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "MainNode", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) view returns(address)
func (_BasMiner *BasMinerSession) MainNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.MainNode(&_BasMiner.CallOpts, arg0)
}

// MainNode is a free data retrieval call binding the contract method 0x57d3a2ba.
//
// Solidity: function MainNode(uint256 ) view returns(address)
func (_BasMiner *BasMinerCallerSession) MainNode(arg0 *big.Int) (common.Address, error) {
	return _BasMiner.Contract.MainNode(&_BasMiner.CallOpts, arg0)
}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() view returns(uint8)
func (_BasMiner *BasMinerCaller) MainNodeSize(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "MainNodeSize")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() view returns(uint8)
func (_BasMiner *BasMinerSession) MainNodeSize() (uint8, error) {
	return _BasMiner.Contract.MainNodeSize(&_BasMiner.CallOpts)
}

// MainNodeSize is a free data retrieval call binding the contract method 0xdea21241.
//
// Solidity: function MainNodeSize() view returns(uint8)
func (_BasMiner *BasMinerCallerSession) MainNodeSize() (uint8, error) {
	return _BasMiner.Contract.MainNodeSize(&_BasMiner.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasMiner *BasMinerCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasMiner *BasMinerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasMiner.Contract.BalanceOf(&_BasMiner.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasMiner *BasMinerCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasMiner.Contract.BalanceOf(&_BasMiner.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_BasMiner *BasMinerCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_BasMiner *BasMinerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _BasMiner.Contract.GetBalance(&_BasMiner.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_BasMiner *BasMinerCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _BasMiner.Contract.GetBalance(&_BasMiner.CallOpts, addr)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMiner *BasMinerCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMiner.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMiner *BasMinerSession) Rel() (common.Address, error) {
	return _BasMiner.Contract.Rel(&_BasMiner.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMiner *BasMinerCallerSession) Rel() (common.Address, error) {
	return _BasMiner.Contract.Rel(&_BasMiner.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMiner *BasMinerTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMiner *BasMinerSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ChangeDAO(&_BasMiner.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMiner *BasMinerTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ChangeDAO(&_BasMiner.TransactOpts, newDao)
}

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerTransactor) AddMiner(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "addMiner", m)
}

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerSession) AddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AddMiner(&_BasMiner.TransactOpts, m)
}

// AddMiner is a paid mutator transaction binding the contract method 0xf3982e5e.
//
// Solidity: function addMiner(address m) returns()
func (_BasMiner *BasMinerTransactorSession) AddMiner(m common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.AddMiner(&_BasMiner.TransactOpts, m)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMiner *BasMinerTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMiner *BasMinerSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ChangeRelation(&_BasMiner.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMiner *BasMinerTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ChangeRelation(&_BasMiner.TransactOpts, new_rel)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactor) RemoveMiner(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "removeMiner", miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerSession) RemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.RemoveMiner(&_BasMiner.TransactOpts, miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x10242590.
//
// Solidity: function removeMiner(address miner) returns()
func (_BasMiner *BasMinerTransactorSession) RemoveMiner(miner common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.RemoveMiner(&_BasMiner.TransactOpts, miner)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactor) ReplaceMiner(opts *bind.TransactOpts, oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "replaceMiner", oldM, newM)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerSession) ReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// ReplaceMiner is a paid mutator transaction binding the contract method 0x0bdf2fd3.
//
// Solidity: function replaceMiner(address oldM, address newM) returns()
func (_BasMiner *BasMinerTransactorSession) ReplaceMiner(oldM common.Address, newM common.Address) (*types.Transaction, error) {
	return _BasMiner.Contract.ReplaceMiner(&_BasMiner.TransactOpts, oldM, newM)
}

// SubAllocate is a paid mutator transaction binding the contract method 0x16c89c1f.
//
// Solidity: function subAllocate(uint256 sum) returns()
func (_BasMiner *BasMinerTransactor) SubAllocate(opts *bind.TransactOpts, sum *big.Int) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "subAllocate", sum)
}

// SubAllocate is a paid mutator transaction binding the contract method 0x16c89c1f.
//
// Solidity: function subAllocate(uint256 sum) returns()
func (_BasMiner *BasMinerSession) SubAllocate(sum *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SubAllocate(&_BasMiner.TransactOpts, sum)
}

// SubAllocate is a paid mutator transaction binding the contract method 0x16c89c1f.
//
// Solidity: function subAllocate(uint256 sum) returns()
func (_BasMiner *BasMinerTransactorSession) SubAllocate(sum *big.Int) (*types.Transaction, error) {
	return _BasMiner.Contract.SubAllocate(&_BasMiner.TransactOpts, sum)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMiner.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerSession) Withdraw() (*types.Transaction, error) {
	return _BasMiner.Contract.Withdraw(&_BasMiner.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BasMiner *BasMinerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BasMiner.Contract.Withdraw(&_BasMiner.TransactOpts)
}

// BasMinerMinerAddIterator is returned from FilterMinerAdd and is used to iterate over the raw logs and unpacked data for MinerAdd events raised by the BasMiner contract.
type BasMinerMinerAddIterator struct {
	Event *BasMinerMinerAdd // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerAdd)
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
		it.Event = new(BasMinerMinerAdd)
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
func (it *BasMinerMinerAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerAdd represents a MinerAdd event raised by the BasMiner contract.
type BasMinerMinerAdd struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerAdd is a free log retrieval operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) FilterMinerAdd(opts *bind.FilterOpts) (*BasMinerMinerAddIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerAdd")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerAddIterator{contract: _BasMiner.contract, event: "MinerAdd", logs: logs, sub: sub}, nil
}

// WatchMinerAdd is a free log subscription operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) WatchMinerAdd(opts *bind.WatchOpts, sink chan<- *BasMinerMinerAdd) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerAdd)
				if err := _BasMiner.contract.UnpackLog(event, "MinerAdd", log); err != nil {
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

// ParseMinerAdd is a log parse operation binding the contract event 0x424346829819678c1f897cd0d980f28004ae4eb06bb95f034a791e0e4e9cebf5.
//
// Solidity: event MinerAdd(address miner)
func (_BasMiner *BasMinerFilterer) ParseMinerAdd(log types.Log) (*BasMinerMinerAdd, error) {
	event := new(BasMinerMinerAdd)
	if err := _BasMiner.contract.UnpackLog(event, "MinerAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMinerMinerRemoveIterator is returned from FilterMinerRemove and is used to iterate over the raw logs and unpacked data for MinerRemove events raised by the BasMiner contract.
type BasMinerMinerRemoveIterator struct {
	Event *BasMinerMinerRemove // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerRemove)
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
		it.Event = new(BasMinerMinerRemove)
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
func (it *BasMinerMinerRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerRemove represents a MinerRemove event raised by the BasMiner contract.
type BasMinerMinerRemove struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRemove is a free log retrieval operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) FilterMinerRemove(opts *bind.FilterOpts) (*BasMinerMinerRemoveIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerRemove")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerRemoveIterator{contract: _BasMiner.contract, event: "MinerRemove", logs: logs, sub: sub}, nil
}

// WatchMinerRemove is a free log subscription operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) WatchMinerRemove(opts *bind.WatchOpts, sink chan<- *BasMinerMinerRemove) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerRemove)
				if err := _BasMiner.contract.UnpackLog(event, "MinerRemove", log); err != nil {
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

// ParseMinerRemove is a log parse operation binding the contract event 0x872f9c238db5a645adcaea90821e7c7884c6e063b9e047fd2658592f7cea397d.
//
// Solidity: event MinerRemove(address miner)
func (_BasMiner *BasMinerFilterer) ParseMinerRemove(log types.Log) (*BasMinerMinerRemove, error) {
	event := new(BasMinerMinerRemove)
	if err := _BasMiner.contract.UnpackLog(event, "MinerRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMinerMinerReplaceIterator is returned from FilterMinerReplace and is used to iterate over the raw logs and unpacked data for MinerReplace events raised by the BasMiner contract.
type BasMinerMinerReplaceIterator struct {
	Event *BasMinerMinerReplace // Event containing the contract specifics and raw log

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
func (it *BasMinerMinerReplaceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerMinerReplace)
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
		it.Event = new(BasMinerMinerReplace)
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
func (it *BasMinerMinerReplaceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerMinerReplaceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerMinerReplace represents a MinerReplace event raised by the BasMiner contract.
type BasMinerMinerReplace struct {
	OldMiner common.Address
	NewMiner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinerReplace is a free log retrieval operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) FilterMinerReplace(opts *bind.FilterOpts) (*BasMinerMinerReplaceIterator, error) {

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "MinerReplace")
	if err != nil {
		return nil, err
	}
	return &BasMinerMinerReplaceIterator{contract: _BasMiner.contract, event: "MinerReplace", logs: logs, sub: sub}, nil
}

// WatchMinerReplace is a free log subscription operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) WatchMinerReplace(opts *bind.WatchOpts, sink chan<- *BasMinerMinerReplace) (event.Subscription, error) {

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "MinerReplace")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerMinerReplace)
				if err := _BasMiner.contract.UnpackLog(event, "MinerReplace", log); err != nil {
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

// ParseMinerReplace is a log parse operation binding the contract event 0xc205c6ca767906b347bde3b64a6a00763decbd4fb0ea515cb8afb411647591ee.
//
// Solidity: event MinerReplace(address oldMiner, address newMiner)
func (_BasMiner *BasMinerFilterer) ParseMinerReplace(log types.Log) (*BasMinerMinerReplace, error) {
	event := new(BasMinerMinerReplace)
	if err := _BasMiner.contract.UnpackLog(event, "MinerReplace", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMinerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BasMiner contract.
type BasMinerWithdrawIterator struct {
	Event *BasMinerWithdraw // Event containing the contract specifics and raw log

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
func (it *BasMinerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMinerWithdraw)
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
		it.Event = new(BasMinerWithdraw)
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
func (it *BasMinerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMinerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMinerWithdraw represents a Withdraw event raised by the BasMiner contract.
type BasMinerWithdraw struct {
	Drawer common.Address
	Amout  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) FilterWithdraw(opts *bind.FilterOpts, drawer []common.Address) (*BasMinerWithdrawIterator, error) {

	var drawerRule []interface{}
	for _, drawerItem := range drawer {
		drawerRule = append(drawerRule, drawerItem)
	}

	logs, sub, err := _BasMiner.contract.FilterLogs(opts, "Withdraw", drawerRule)
	if err != nil {
		return nil, err
	}
	return &BasMinerWithdrawIterator{contract: _BasMiner.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BasMinerWithdraw, drawer []common.Address) (event.Subscription, error) {

	var drawerRule []interface{}
	for _, drawerItem := range drawer {
		drawerRule = append(drawerRule, drawerItem)
	}

	logs, sub, err := _BasMiner.contract.WatchLogs(opts, "Withdraw", drawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMinerWithdraw)
				if err := _BasMiner.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed drawer, uint256 amout)
func (_BasMiner *BasMinerFilterer) ParseWithdraw(log types.Log) (*BasMinerWithdraw, error) {
	event := new(BasMinerWithdraw)
	if err := _BasMiner.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
