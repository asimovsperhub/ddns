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

// BasExpiredOwnershipABI is the input ABI used to generate the binding from.
const BasExpiredOwnershipABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Takeover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsCountsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"assetsOf\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"extend\",\"type\":\"uint256\"}],\"name\":\"extendTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"newOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ownerOfWithExpire\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"takeover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"updateByDaoProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BasExpiredOwnership is an auto generated Go binding around an Ethereum contract.
type BasExpiredOwnership struct {
	BasExpiredOwnershipCaller     // Read-only binding to the contract
	BasExpiredOwnershipTransactor // Write-only binding to the contract
	BasExpiredOwnershipFilterer   // Log filterer for contract events
}

// BasExpiredOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasExpiredOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasExpiredOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasExpiredOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasExpiredOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasExpiredOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasExpiredOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasExpiredOwnershipSession struct {
	Contract     *BasExpiredOwnership // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasExpiredOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasExpiredOwnershipCallerSession struct {
	Contract *BasExpiredOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BasExpiredOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasExpiredOwnershipTransactorSession struct {
	Contract     *BasExpiredOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BasExpiredOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasExpiredOwnershipRaw struct {
	Contract *BasExpiredOwnership // Generic contract binding to access the raw methods on
}

// BasExpiredOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasExpiredOwnershipCallerRaw struct {
	Contract *BasExpiredOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// BasExpiredOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasExpiredOwnershipTransactorRaw struct {
	Contract *BasExpiredOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasExpiredOwnership creates a new instance of BasExpiredOwnership, bound to a specific deployed contract.
func NewBasExpiredOwnership(address common.Address, backend bind.ContractBackend) (*BasExpiredOwnership, error) {
	contract, err := bindBasExpiredOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnership{BasExpiredOwnershipCaller: BasExpiredOwnershipCaller{contract: contract}, BasExpiredOwnershipTransactor: BasExpiredOwnershipTransactor{contract: contract}, BasExpiredOwnershipFilterer: BasExpiredOwnershipFilterer{contract: contract}}, nil
}

// NewBasExpiredOwnershipCaller creates a new read-only instance of BasExpiredOwnership, bound to a specific deployed contract.
func NewBasExpiredOwnershipCaller(address common.Address, caller bind.ContractCaller) (*BasExpiredOwnershipCaller, error) {
	contract, err := bindBasExpiredOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipCaller{contract: contract}, nil
}

// NewBasExpiredOwnershipTransactor creates a new write-only instance of BasExpiredOwnership, bound to a specific deployed contract.
func NewBasExpiredOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*BasExpiredOwnershipTransactor, error) {
	contract, err := bindBasExpiredOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipTransactor{contract: contract}, nil
}

// NewBasExpiredOwnershipFilterer creates a new log filterer instance of BasExpiredOwnership, bound to a specific deployed contract.
func NewBasExpiredOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*BasExpiredOwnershipFilterer, error) {
	contract, err := bindBasExpiredOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipFilterer{contract: contract}, nil
}

// bindBasExpiredOwnership binds a generic wrapper to an already deployed contract.
func bindBasExpiredOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasExpiredOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasExpiredOwnership *BasExpiredOwnershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasExpiredOwnership.Contract.BasExpiredOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasExpiredOwnership *BasExpiredOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.BasExpiredOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasExpiredOwnership *BasExpiredOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.BasExpiredOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasExpiredOwnership *BasExpiredOwnershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasExpiredOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) DAOAddress() (common.Address, error) {
	return _BasExpiredOwnership.Contract.DAOAddress(&_BasExpiredOwnership.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) DAOAddress() (common.Address, error) {
	return _BasExpiredOwnership.Contract.DAOAddress(&_BasExpiredOwnership.CallOpts)
}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) AssetsCountsOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "assetsCountsOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) AssetsCountsOf() (*big.Int, error) {
	return _BasExpiredOwnership.Contract.AssetsCountsOf(&_BasExpiredOwnership.CallOpts)
}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) AssetsCountsOf() (*big.Int, error) {
	return _BasExpiredOwnership.Contract.AssetsCountsOf(&_BasExpiredOwnership.CallOpts)
}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) AssetsOf(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "assetsOf", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasExpiredOwnership *BasExpiredOwnershipSession) AssetsOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasExpiredOwnership.Contract.AssetsOf(&_BasExpiredOwnership.CallOpts, start, end)
}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) AssetsOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasExpiredOwnership.Contract.AssetsOf(&_BasExpiredOwnership.CallOpts, start, end)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) OwnerOf(opts *bind.CallOpts, nameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "ownerOf", nameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasExpiredOwnership.Contract.OwnerOf(&_BasExpiredOwnership.CallOpts, nameHash)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasExpiredOwnership.Contract.OwnerOf(&_BasExpiredOwnership.CallOpts, nameHash)
}

// OwnerOfWithExpire is a free data retrieval call binding the contract method 0x33eeb127.
//
// Solidity: function ownerOfWithExpire(bytes32 nameHash) view returns(address, uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) OwnerOfWithExpire(opts *bind.CallOpts, nameHash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "ownerOfWithExpire", nameHash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// OwnerOfWithExpire is a free data retrieval call binding the contract method 0x33eeb127.
//
// Solidity: function ownerOfWithExpire(bytes32 nameHash) view returns(address, uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) OwnerOfWithExpire(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasExpiredOwnership.Contract.OwnerOfWithExpire(&_BasExpiredOwnership.CallOpts, nameHash)
}

// OwnerOfWithExpire is a free data retrieval call binding the contract method 0x33eeb127.
//
// Solidity: function ownerOfWithExpire(bytes32 nameHash) view returns(address, uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) OwnerOfWithExpire(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasExpiredOwnership.Contract.OwnerOfWithExpire(&_BasExpiredOwnership.CallOpts, nameHash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasExpiredOwnership.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) Rel() (common.Address, error) {
	return _BasExpiredOwnership.Contract.Rel(&_BasExpiredOwnership.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasExpiredOwnership *BasExpiredOwnershipCallerSession) Rel() (common.Address, error) {
	return _BasExpiredOwnership.Contract.Rel(&_BasExpiredOwnership.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ChangeDAO(&_BasExpiredOwnership.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ChangeDAO(&_BasExpiredOwnership.TransactOpts, newDao)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ChangeRelation(&_BasExpiredOwnership.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ChangeRelation(&_BasExpiredOwnership.TransactOpts, new_rel)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) ExtendTime(opts *bind.TransactOpts, nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "extendTime", nameHash, extend)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipSession) ExtendTime(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ExtendTime(&_BasExpiredOwnership.TransactOpts, nameHash, extend)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) ExtendTime(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.ExtendTime(&_BasExpiredOwnership.TransactOpts, nameHash, extend)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) NewOwnership(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "newOwnership", nameHash, owner, expire)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipSession) NewOwnership(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.NewOwnership(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) NewOwnership(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.NewOwnership(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) Takeover(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "takeover", nameHash, owner, expire)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipSession) Takeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.Takeover(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) Takeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.Takeover(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactor) UpdateByDaoProposal(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.contract.Transact(opts, "updateByDaoProposal", nameHash, owner, expire)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipSession) UpdateByDaoProposal(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.UpdateByDaoProposal(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasExpiredOwnership *BasExpiredOwnershipTransactorSession) UpdateByDaoProposal(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasExpiredOwnership.Contract.UpdateByDaoProposal(&_BasExpiredOwnership.TransactOpts, nameHash, owner, expire)
}

// BasExpiredOwnershipAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipAddIterator struct {
	Event *BasExpiredOwnershipAdd // Event containing the contract specifics and raw log

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
func (it *BasExpiredOwnershipAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasExpiredOwnershipAdd)
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
		it.Event = new(BasExpiredOwnershipAdd)
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
func (it *BasExpiredOwnershipAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasExpiredOwnershipAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasExpiredOwnershipAdd represents a Add event raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipAdd struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) FilterAdd(opts *bind.FilterOpts) (*BasExpiredOwnershipAddIterator, error) {

	logs, sub, err := _BasExpiredOwnership.contract.FilterLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipAddIterator{contract: _BasExpiredOwnership.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *BasExpiredOwnershipAdd) (event.Subscription, error) {

	logs, sub, err := _BasExpiredOwnership.contract.WatchLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasExpiredOwnershipAdd)
				if err := _BasExpiredOwnership.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) ParseAdd(log types.Log) (*BasExpiredOwnershipAdd, error) {
	event := new(BasExpiredOwnershipAdd)
	if err := _BasExpiredOwnership.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasExpiredOwnershipExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipExtendIterator struct {
	Event *BasExpiredOwnershipExtend // Event containing the contract specifics and raw log

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
func (it *BasExpiredOwnershipExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasExpiredOwnershipExtend)
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
		it.Event = new(BasExpiredOwnershipExtend)
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
func (it *BasExpiredOwnershipExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasExpiredOwnershipExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasExpiredOwnershipExtend represents a Extend event raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipExtend struct {
	NameHash [32]byte
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) FilterExtend(opts *bind.FilterOpts) (*BasExpiredOwnershipExtendIterator, error) {

	logs, sub, err := _BasExpiredOwnership.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipExtendIterator{contract: _BasExpiredOwnership.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *BasExpiredOwnershipExtend) (event.Subscription, error) {

	logs, sub, err := _BasExpiredOwnership.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasExpiredOwnershipExtend)
				if err := _BasExpiredOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) ParseExtend(log types.Log) (*BasExpiredOwnershipExtend, error) {
	event := new(BasExpiredOwnershipExtend)
	if err := _BasExpiredOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasExpiredOwnershipTakeoverIterator is returned from FilterTakeover and is used to iterate over the raw logs and unpacked data for Takeover events raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipTakeoverIterator struct {
	Event *BasExpiredOwnershipTakeover // Event containing the contract specifics and raw log

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
func (it *BasExpiredOwnershipTakeoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasExpiredOwnershipTakeover)
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
		it.Event = new(BasExpiredOwnershipTakeover)
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
func (it *BasExpiredOwnershipTakeoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasExpiredOwnershipTakeoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasExpiredOwnershipTakeover represents a Takeover event raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipTakeover struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTakeover is a free log retrieval operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) FilterTakeover(opts *bind.FilterOpts) (*BasExpiredOwnershipTakeoverIterator, error) {

	logs, sub, err := _BasExpiredOwnership.contract.FilterLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipTakeoverIterator{contract: _BasExpiredOwnership.contract, event: "Takeover", logs: logs, sub: sub}, nil
}

// WatchTakeover is a free log subscription operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) WatchTakeover(opts *bind.WatchOpts, sink chan<- *BasExpiredOwnershipTakeover) (event.Subscription, error) {

	logs, sub, err := _BasExpiredOwnership.contract.WatchLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasExpiredOwnershipTakeover)
				if err := _BasExpiredOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
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

// ParseTakeover is a log parse operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) ParseTakeover(log types.Log) (*BasExpiredOwnershipTakeover, error) {
	event := new(BasExpiredOwnershipTakeover)
	if err := _BasExpiredOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasExpiredOwnershipUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipUpdateIterator struct {
	Event *BasExpiredOwnershipUpdate // Event containing the contract specifics and raw log

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
func (it *BasExpiredOwnershipUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasExpiredOwnershipUpdate)
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
		it.Event = new(BasExpiredOwnershipUpdate)
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
func (it *BasExpiredOwnershipUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasExpiredOwnershipUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasExpiredOwnershipUpdate represents a Update event raised by the BasExpiredOwnership contract.
type BasExpiredOwnershipUpdate struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) FilterUpdate(opts *bind.FilterOpts) (*BasExpiredOwnershipUpdateIterator, error) {

	logs, sub, err := _BasExpiredOwnership.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &BasExpiredOwnershipUpdateIterator{contract: _BasExpiredOwnership.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *BasExpiredOwnershipUpdate) (event.Subscription, error) {

	logs, sub, err := _BasExpiredOwnership.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasExpiredOwnershipUpdate)
				if err := _BasExpiredOwnership.contract.UnpackLog(event, "Update", log); err != nil {
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

// ParseUpdate is a log parse operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasExpiredOwnership *BasExpiredOwnershipFilterer) ParseUpdate(log types.Log) (*BasExpiredOwnershipUpdate, error) {
	event := new(BasExpiredOwnershipUpdate)
	if err := _BasExpiredOwnership.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
