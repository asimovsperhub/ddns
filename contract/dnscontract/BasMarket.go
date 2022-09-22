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

// BasMarketMetaData contains all meta data concerning the BasMarket contract.
var BasMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SellRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SoldBySell\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AT_LEAST_REMAIN_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"AddToSells\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BuyFromSells\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ChangeSellPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DomainSellOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"RemoveSellOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDays\",\"type\":\"uint256\"}],\"name\":\"changeAtLeastDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ordersCountsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"ordersOf\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BasMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use BasMarketMetaData.ABI instead.
var BasMarketABI = BasMarketMetaData.ABI

// BasMarket is an auto generated Go binding around an Ethereum contract.
type BasMarket struct {
	BasMarketCaller     // Read-only binding to the contract
	BasMarketTransactor // Write-only binding to the contract
	BasMarketFilterer   // Log filterer for contract events
}

// BasMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMarketSession struct {
	Contract     *BasMarket        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMarketCallerSession struct {
	Contract *BasMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BasMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMarketTransactorSession struct {
	Contract     *BasMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMarketRaw struct {
	Contract *BasMarket // Generic contract binding to access the raw methods on
}

// BasMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMarketCallerRaw struct {
	Contract *BasMarketCaller // Generic read-only contract binding to access the raw methods on
}

// BasMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMarketTransactorRaw struct {
	Contract *BasMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMarket creates a new instance of BasMarket, bound to a specific deployed contract.
func NewBasMarket(address common.Address, backend bind.ContractBackend) (*BasMarket, error) {
	contract, err := bindBasMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMarket{BasMarketCaller: BasMarketCaller{contract: contract}, BasMarketTransactor: BasMarketTransactor{contract: contract}, BasMarketFilterer: BasMarketFilterer{contract: contract}}, nil
}

// NewBasMarketCaller creates a new read-only instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketCaller(address common.Address, caller bind.ContractCaller) (*BasMarketCaller, error) {
	contract, err := bindBasMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMarketCaller{contract: contract}, nil
}

// NewBasMarketTransactor creates a new write-only instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMarketTransactor, error) {
	contract, err := bindBasMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMarketTransactor{contract: contract}, nil
}

// NewBasMarketFilterer creates a new log filterer instance of BasMarket, bound to a specific deployed contract.
func NewBasMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMarketFilterer, error) {
	contract, err := bindBasMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMarketFilterer{contract: contract}, nil
}

// bindBasMarket binds a generic wrapper to an already deployed contract.
func bindBasMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMarket *BasMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMarket.Contract.BasMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMarket *BasMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMarket.Contract.BasMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMarket *BasMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMarket.Contract.BasMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMarket *BasMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMarket *BasMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMarket *BasMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMarket.Contract.contract.Transact(opts, method, params...)
}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() view returns(uint256)
func (_BasMarket *BasMarketCaller) ATLEASTREMAINTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "AT_LEAST_REMAIN_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() view returns(uint256)
func (_BasMarket *BasMarketSession) ATLEASTREMAINTIME() (*big.Int, error) {
	return _BasMarket.Contract.ATLEASTREMAINTIME(&_BasMarket.CallOpts)
}

// ATLEASTREMAINTIME is a free data retrieval call binding the contract method 0x1d6a1ed0.
//
// Solidity: function AT_LEAST_REMAIN_TIME() view returns(uint256)
func (_BasMarket *BasMarketCallerSession) ATLEASTREMAINTIME() (*big.Int, error) {
	return _BasMarket.Contract.ATLEASTREMAINTIME(&_BasMarket.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMarket *BasMarketCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMarket *BasMarketSession) DAOAddress() (common.Address, error) {
	return _BasMarket.Contract.DAOAddress(&_BasMarket.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMarket *BasMarketCallerSession) DAOAddress() (common.Address, error) {
	return _BasMarket.Contract.DAOAddress(&_BasMarket.CallOpts)
}

// DomainSellOrders is a free data retrieval call binding the contract method 0x4c6606a3.
//
// Solidity: function DomainSellOrders(address , bytes32 ) view returns(uint256)
func (_BasMarket *BasMarketCaller) DomainSellOrders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "DomainSellOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DomainSellOrders is a free data retrieval call binding the contract method 0x4c6606a3.
//
// Solidity: function DomainSellOrders(address , bytes32 ) view returns(uint256)
func (_BasMarket *BasMarketSession) DomainSellOrders(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _BasMarket.Contract.DomainSellOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// DomainSellOrders is a free data retrieval call binding the contract method 0x4c6606a3.
//
// Solidity: function DomainSellOrders(address , bytes32 ) view returns(uint256)
func (_BasMarket *BasMarketCallerSession) DomainSellOrders(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _BasMarket.Contract.DomainSellOrders(&_BasMarket.CallOpts, arg0, arg1)
}

// OrdersCountsOf is a free data retrieval call binding the contract method 0xb1f45a90.
//
// Solidity: function ordersCountsOf() view returns(uint256)
func (_BasMarket *BasMarketCaller) OrdersCountsOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "ordersCountsOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrdersCountsOf is a free data retrieval call binding the contract method 0xb1f45a90.
//
// Solidity: function ordersCountsOf() view returns(uint256)
func (_BasMarket *BasMarketSession) OrdersCountsOf() (*big.Int, error) {
	return _BasMarket.Contract.OrdersCountsOf(&_BasMarket.CallOpts)
}

// OrdersCountsOf is a free data retrieval call binding the contract method 0xb1f45a90.
//
// Solidity: function ordersCountsOf() view returns(uint256)
func (_BasMarket *BasMarketCallerSession) OrdersCountsOf() (*big.Int, error) {
	return _BasMarket.Contract.OrdersCountsOf(&_BasMarket.CallOpts)
}

// OrdersOf is a free data retrieval call binding the contract method 0x65f6c15e.
//
// Solidity: function ordersOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasMarket *BasMarketCaller) OrdersOf(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "ordersOf", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// OrdersOf is a free data retrieval call binding the contract method 0x65f6c15e.
//
// Solidity: function ordersOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasMarket *BasMarketSession) OrdersOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasMarket.Contract.OrdersOf(&_BasMarket.CallOpts, start, end)
}

// OrdersOf is a free data retrieval call binding the contract method 0x65f6c15e.
//
// Solidity: function ordersOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasMarket *BasMarketCallerSession) OrdersOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasMarket.Contract.OrdersOf(&_BasMarket.CallOpts, start, end)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMarket *BasMarketCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMarket.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMarket *BasMarketSession) Rel() (common.Address, error) {
	return _BasMarket.Contract.Rel(&_BasMarket.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMarket *BasMarketCallerSession) Rel() (common.Address, error) {
	return _BasMarket.Contract.Rel(&_BasMarket.CallOpts)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactor) AddToSells(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "AddToSells", nameHash, price)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketSession) AddToSells(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToSells(&_BasMarket.TransactOpts, nameHash, price)
}

// AddToSells is a paid mutator transaction binding the contract method 0x75a69f19.
//
// Solidity: function AddToSells(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactorSession) AddToSells(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.AddToSells(&_BasMarket.TransactOpts, nameHash, price)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketTransactor) BuyFromSells(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "BuyFromSells", nameHash, owner)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketSession) BuyFromSells(nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.BuyFromSells(&_BasMarket.TransactOpts, nameHash, owner)
}

// BuyFromSells is a paid mutator transaction binding the contract method 0x3a2ef923.
//
// Solidity: function BuyFromSells(bytes32 nameHash, address owner) returns()
func (_BasMarket *BasMarketTransactorSession) BuyFromSells(nameHash [32]byte, owner common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.BuyFromSells(&_BasMarket.TransactOpts, nameHash, owner)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMarket *BasMarketTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMarket *BasMarketSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeDAO(&_BasMarket.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeDAO(&_BasMarket.TransactOpts, newDao)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactor) ChangeSellPrice(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "ChangeSellPrice", nameHash, price)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketSession) ChangeSellPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeSellPrice(&_BasMarket.TransactOpts, nameHash, price)
}

// ChangeSellPrice is a paid mutator transaction binding the contract method 0x1d4edd20.
//
// Solidity: function ChangeSellPrice(bytes32 nameHash, uint256 price) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeSellPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeSellPrice(&_BasMarket.TransactOpts, nameHash, price)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactor) RemoveSellOrder(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "RemoveSellOrder", nameHash)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketSession) RemoveSellOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveSellOrder(&_BasMarket.TransactOpts, nameHash)
}

// RemoveSellOrder is a paid mutator transaction binding the contract method 0xe53427ca.
//
// Solidity: function RemoveSellOrder(bytes32 nameHash) returns()
func (_BasMarket *BasMarketTransactorSession) RemoveSellOrder(nameHash [32]byte) (*types.Transaction, error) {
	return _BasMarket.Contract.RemoveSellOrder(&_BasMarket.TransactOpts, nameHash)
}

// ChangeAtLeastDays is a paid mutator transaction binding the contract method 0xd8e21021.
//
// Solidity: function changeAtLeastDays(uint256 newDays) returns()
func (_BasMarket *BasMarketTransactor) ChangeAtLeastDays(opts *bind.TransactOpts, newDays *big.Int) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "changeAtLeastDays", newDays)
}

// ChangeAtLeastDays is a paid mutator transaction binding the contract method 0xd8e21021.
//
// Solidity: function changeAtLeastDays(uint256 newDays) returns()
func (_BasMarket *BasMarketSession) ChangeAtLeastDays(newDays *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeAtLeastDays(&_BasMarket.TransactOpts, newDays)
}

// ChangeAtLeastDays is a paid mutator transaction binding the contract method 0xd8e21021.
//
// Solidity: function changeAtLeastDays(uint256 newDays) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeAtLeastDays(newDays *big.Int) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeAtLeastDays(&_BasMarket.TransactOpts, newDays)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMarket *BasMarketTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasMarket.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMarket *BasMarketSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeRelation(&_BasMarket.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMarket *BasMarketTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMarket.Contract.ChangeRelation(&_BasMarket.TransactOpts, new_rel)
}

// BasMarketSellAddedIterator is returned from FilterSellAdded and is used to iterate over the raw logs and unpacked data for SellAdded events raised by the BasMarket contract.
type BasMarketSellAddedIterator struct {
	Event *BasMarketSellAdded // Event containing the contract specifics and raw log

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
func (it *BasMarketSellAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellAdded)
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
		it.Event = new(BasMarketSellAdded)
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
func (it *BasMarketSellAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellAdded represents a SellAdded event raised by the BasMarket contract.
type BasMarketSellAdded struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellAdded is a free log retrieval operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 indexed nameHash, address indexed operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) FilterSellAdded(opts *bind.FilterOpts, nameHash [][32]byte, operator []common.Address, price []*big.Int) (*BasMarketSellAddedIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellAdded", nameHashRule, operatorRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &BasMarketSellAddedIterator{contract: _BasMarket.contract, event: "SellAdded", logs: logs, sub: sub}, nil
}

// WatchSellAdded is a free log subscription operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 indexed nameHash, address indexed operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) WatchSellAdded(opts *bind.WatchOpts, sink chan<- *BasMarketSellAdded, nameHash [][32]byte, operator []common.Address, price []*big.Int) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellAdded", nameHashRule, operatorRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellAdded)
				if err := _BasMarket.contract.UnpackLog(event, "SellAdded", log); err != nil {
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

// ParseSellAdded is a log parse operation binding the contract event 0x10f7d9b00b30f2c94496f80f416e54d140fcd3c6040b5440ae55c39a3294086c.
//
// Solidity: event SellAdded(bytes32 indexed nameHash, address indexed operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) ParseSellAdded(log types.Log) (*BasMarketSellAdded, error) {
	event := new(BasMarketSellAdded)
	if err := _BasMarket.contract.UnpackLog(event, "SellAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMarketSellChangedIterator is returned from FilterSellChanged and is used to iterate over the raw logs and unpacked data for SellChanged events raised by the BasMarket contract.
type BasMarketSellChangedIterator struct {
	Event *BasMarketSellChanged // Event containing the contract specifics and raw log

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
func (it *BasMarketSellChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellChanged)
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
		it.Event = new(BasMarketSellChanged)
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
func (it *BasMarketSellChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellChanged represents a SellChanged event raised by the BasMarket contract.
type BasMarketSellChanged struct {
	NameHash [32]byte
	Operator common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellChanged is a free log retrieval operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 indexed nameHash, address operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) FilterSellChanged(opts *bind.FilterOpts, nameHash [][32]byte, price []*big.Int) (*BasMarketSellChangedIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellChanged", nameHashRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &BasMarketSellChangedIterator{contract: _BasMarket.contract, event: "SellChanged", logs: logs, sub: sub}, nil
}

// WatchSellChanged is a free log subscription operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 indexed nameHash, address operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) WatchSellChanged(opts *bind.WatchOpts, sink chan<- *BasMarketSellChanged, nameHash [][32]byte, price []*big.Int) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellChanged", nameHashRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellChanged)
				if err := _BasMarket.contract.UnpackLog(event, "SellChanged", log); err != nil {
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

// ParseSellChanged is a log parse operation binding the contract event 0xe2fbe11a69dd3397016a8ac5cc0d982dd57cc1586c3400f8e3a3167d0dcd3634.
//
// Solidity: event SellChanged(bytes32 indexed nameHash, address operator, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) ParseSellChanged(log types.Log) (*BasMarketSellChanged, error) {
	event := new(BasMarketSellChanged)
	if err := _BasMarket.contract.UnpackLog(event, "SellChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMarketSellRemovedIterator is returned from FilterSellRemoved and is used to iterate over the raw logs and unpacked data for SellRemoved events raised by the BasMarket contract.
type BasMarketSellRemovedIterator struct {
	Event *BasMarketSellRemoved // Event containing the contract specifics and raw log

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
func (it *BasMarketSellRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSellRemoved)
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
		it.Event = new(BasMarketSellRemoved)
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
func (it *BasMarketSellRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSellRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSellRemoved represents a SellRemoved event raised by the BasMarket contract.
type BasMarketSellRemoved struct {
	NameHash [32]byte
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSellRemoved is a free log retrieval operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 indexed nameHash, address operator)
func (_BasMarket *BasMarketFilterer) FilterSellRemoved(opts *bind.FilterOpts, nameHash [][32]byte) (*BasMarketSellRemovedIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SellRemoved", nameHashRule)
	if err != nil {
		return nil, err
	}
	return &BasMarketSellRemovedIterator{contract: _BasMarket.contract, event: "SellRemoved", logs: logs, sub: sub}, nil
}

// WatchSellRemoved is a free log subscription operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 indexed nameHash, address operator)
func (_BasMarket *BasMarketFilterer) WatchSellRemoved(opts *bind.WatchOpts, sink chan<- *BasMarketSellRemoved, nameHash [][32]byte) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SellRemoved", nameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSellRemoved)
				if err := _BasMarket.contract.UnpackLog(event, "SellRemoved", log); err != nil {
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

// ParseSellRemoved is a log parse operation binding the contract event 0xd369082d239e5af317337ed91a37711a5809487a596e485189541c3219a35a0a.
//
// Solidity: event SellRemoved(bytes32 indexed nameHash, address operator)
func (_BasMarket *BasMarketFilterer) ParseSellRemoved(log types.Log) (*BasMarketSellRemoved, error) {
	event := new(BasMarketSellRemoved)
	if err := _BasMarket.contract.UnpackLog(event, "SellRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMarketSoldBySellIterator is returned from FilterSoldBySell and is used to iterate over the raw logs and unpacked data for SoldBySell events raised by the BasMarket contract.
type BasMarketSoldBySellIterator struct {
	Event *BasMarketSoldBySell // Event containing the contract specifics and raw log

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
func (it *BasMarketSoldBySellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMarketSoldBySell)
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
		it.Event = new(BasMarketSoldBySell)
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
func (it *BasMarketSoldBySellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMarketSoldBySellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMarketSoldBySell represents a SoldBySell event raised by the BasMarket contract.
type BasMarketSoldBySell struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSoldBySell is a free log retrieval operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 indexed nameHash, address from, address indexed to, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) FilterSoldBySell(opts *bind.FilterOpts, nameHash [][32]byte, to []common.Address, price []*big.Int) (*BasMarketSoldBySellIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.FilterLogs(opts, "SoldBySell", nameHashRule, toRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &BasMarketSoldBySellIterator{contract: _BasMarket.contract, event: "SoldBySell", logs: logs, sub: sub}, nil
}

// WatchSoldBySell is a free log subscription operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 indexed nameHash, address from, address indexed to, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) WatchSoldBySell(opts *bind.WatchOpts, sink chan<- *BasMarketSoldBySell, nameHash [][32]byte, to []common.Address, price []*big.Int) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _BasMarket.contract.WatchLogs(opts, "SoldBySell", nameHashRule, toRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMarketSoldBySell)
				if err := _BasMarket.contract.UnpackLog(event, "SoldBySell", log); err != nil {
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

// ParseSoldBySell is a log parse operation binding the contract event 0x9c0f845844791139cd02a1448f071f31346e7d1fee8cafc98925d507dc4d6244.
//
// Solidity: event SoldBySell(bytes32 indexed nameHash, address from, address indexed to, uint256 indexed price)
func (_BasMarket *BasMarketFilterer) ParseSoldBySell(log types.Log) (*BasMarketSoldBySell, error) {
	event := new(BasMarketSoldBySell)
	if err := _BasMarket.contract.UnpackLog(event, "SoldBySell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
