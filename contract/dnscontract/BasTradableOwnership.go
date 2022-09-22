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

// BasTradableOwnershipMetaData contains all meta data concerning the BasTradableOwnership contract.
var BasTradableOwnershipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Takeover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"TransferFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Update\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetsCountsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"assetsOf\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"extend\",\"type\":\"uint256\"}],\"name\":\"extendTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"newOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ownerOfWithExpire\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"takeover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"}],\"name\":\"updateByDaoProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasTradableOwnershipABI is the input ABI used to generate the binding from.
// Deprecated: Use BasTradableOwnershipMetaData.ABI instead.
var BasTradableOwnershipABI = BasTradableOwnershipMetaData.ABI

// BasTradableOwnership is an auto generated Go binding around an Ethereum contract.
type BasTradableOwnership struct {
	BasTradableOwnershipCaller     // Read-only binding to the contract
	BasTradableOwnershipTransactor // Write-only binding to the contract
	BasTradableOwnershipFilterer   // Log filterer for contract events
}

// BasTradableOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasTradableOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasTradableOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasTradableOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasTradableOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasTradableOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasTradableOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasTradableOwnershipSession struct {
	Contract     *BasTradableOwnership // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasTradableOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasTradableOwnershipCallerSession struct {
	Contract *BasTradableOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BasTradableOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasTradableOwnershipTransactorSession struct {
	Contract     *BasTradableOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BasTradableOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasTradableOwnershipRaw struct {
	Contract *BasTradableOwnership // Generic contract binding to access the raw methods on
}

// BasTradableOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasTradableOwnershipCallerRaw struct {
	Contract *BasTradableOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// BasTradableOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasTradableOwnershipTransactorRaw struct {
	Contract *BasTradableOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasTradableOwnership creates a new instance of BasTradableOwnership, bound to a specific deployed contract.
func NewBasTradableOwnership(address common.Address, backend bind.ContractBackend) (*BasTradableOwnership, error) {
	contract, err := bindBasTradableOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnership{BasTradableOwnershipCaller: BasTradableOwnershipCaller{contract: contract}, BasTradableOwnershipTransactor: BasTradableOwnershipTransactor{contract: contract}, BasTradableOwnershipFilterer: BasTradableOwnershipFilterer{contract: contract}}, nil
}

// NewBasTradableOwnershipCaller creates a new read-only instance of BasTradableOwnership, bound to a specific deployed contract.
func NewBasTradableOwnershipCaller(address common.Address, caller bind.ContractCaller) (*BasTradableOwnershipCaller, error) {
	contract, err := bindBasTradableOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipCaller{contract: contract}, nil
}

// NewBasTradableOwnershipTransactor creates a new write-only instance of BasTradableOwnership, bound to a specific deployed contract.
func NewBasTradableOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*BasTradableOwnershipTransactor, error) {
	contract, err := bindBasTradableOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipTransactor{contract: contract}, nil
}

// NewBasTradableOwnershipFilterer creates a new log filterer instance of BasTradableOwnership, bound to a specific deployed contract.
func NewBasTradableOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*BasTradableOwnershipFilterer, error) {
	contract, err := bindBasTradableOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipFilterer{contract: contract}, nil
}

// bindBasTradableOwnership binds a generic wrapper to an already deployed contract.
func bindBasTradableOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasTradableOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasTradableOwnership *BasTradableOwnershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasTradableOwnership.Contract.BasTradableOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasTradableOwnership *BasTradableOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.BasTradableOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasTradableOwnership *BasTradableOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.BasTradableOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasTradableOwnership *BasTradableOwnershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasTradableOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasTradableOwnership *BasTradableOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasTradableOwnership *BasTradableOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipSession) DAOAddress() (common.Address, error) {
	return _BasTradableOwnership.Contract.DAOAddress(&_BasTradableOwnership.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) DAOAddress() (common.Address, error) {
	return _BasTradableOwnership.Contract.DAOAddress(&_BasTradableOwnership.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCaller) Allowance(opts *bind.CallOpts, owner common.Address, nameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "allowance", owner, nameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasTradableOwnership.Contract.Allowance(&_BasTradableOwnership.CallOpts, owner, nameHash)
}

// Allowance is a free data retrieval call binding the contract method 0xbcf9fcad.
//
// Solidity: function allowance(address owner, bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) Allowance(owner common.Address, nameHash [32]byte) (common.Address, error) {
	return _BasTradableOwnership.Contract.Allowance(&_BasTradableOwnership.CallOpts, owner, nameHash)
}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipCaller) AssetsCountsOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "assetsCountsOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipSession) AssetsCountsOf() (*big.Int, error) {
	return _BasTradableOwnership.Contract.AssetsCountsOf(&_BasTradableOwnership.CallOpts)
}

// AssetsCountsOf is a free data retrieval call binding the contract method 0x8549ddd1.
//
// Solidity: function assetsCountsOf() view returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) AssetsCountsOf() (*big.Int, error) {
	return _BasTradableOwnership.Contract.AssetsCountsOf(&_BasTradableOwnership.CallOpts)
}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasTradableOwnership *BasTradableOwnershipCaller) AssetsOf(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "assetsOf", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasTradableOwnership *BasTradableOwnershipSession) AssetsOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasTradableOwnership.Contract.AssetsOf(&_BasTradableOwnership.CallOpts, start, end)
}

// AssetsOf is a free data retrieval call binding the contract method 0xd2a03b51.
//
// Solidity: function assetsOf(uint256 start, uint256 end) view returns(bytes32[])
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) AssetsOf(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _BasTradableOwnership.Contract.AssetsOf(&_BasTradableOwnership.CallOpts, start, end)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCaller) OwnerOf(opts *bind.CallOpts, nameHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "ownerOf", nameHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasTradableOwnership.Contract.OwnerOf(&_BasTradableOwnership.CallOpts, nameHash)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(bytes32 nameHash) view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) OwnerOf(nameHash [32]byte) (common.Address, error) {
	return _BasTradableOwnership.Contract.OwnerOf(&_BasTradableOwnership.CallOpts, nameHash)
}

// OwnerOfWithExpire is a free data retrieval call binding the contract method 0x33eeb127.
//
// Solidity: function ownerOfWithExpire(bytes32 nameHash) view returns(address, uint256)
func (_BasTradableOwnership *BasTradableOwnershipCaller) OwnerOfWithExpire(opts *bind.CallOpts, nameHash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "ownerOfWithExpire", nameHash)

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
func (_BasTradableOwnership *BasTradableOwnershipSession) OwnerOfWithExpire(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasTradableOwnership.Contract.OwnerOfWithExpire(&_BasTradableOwnership.CallOpts, nameHash)
}

// OwnerOfWithExpire is a free data retrieval call binding the contract method 0x33eeb127.
//
// Solidity: function ownerOfWithExpire(bytes32 nameHash) view returns(address, uint256)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) OwnerOfWithExpire(nameHash [32]byte) (common.Address, *big.Int, error) {
	return _BasTradableOwnership.Contract.OwnerOfWithExpire(&_BasTradableOwnership.CallOpts, nameHash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasTradableOwnership.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipSession) Rel() (common.Address, error) {
	return _BasTradableOwnership.Contract.Rel(&_BasTradableOwnership.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasTradableOwnership *BasTradableOwnershipCallerSession) Rel() (common.Address, error) {
	return _BasTradableOwnership.Contract.Rel(&_BasTradableOwnership.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ChangeDAO(&_BasTradableOwnership.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ChangeDAO(&_BasTradableOwnership.TransactOpts, newDao)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) Approve(opts *bind.TransactOpts, nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "approve", nameHash, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) Approve(nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Approve(&_BasTradableOwnership.TransactOpts, nameHash, spender)
}

// Approve is a paid mutator transaction binding the contract method 0xb2e2c99b.
//
// Solidity: function approve(bytes32 nameHash, address spender) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) Approve(nameHash [32]byte, spender common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Approve(&_BasTradableOwnership.TransactOpts, nameHash, spender)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ChangeRelation(&_BasTradableOwnership.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ChangeRelation(&_BasTradableOwnership.TransactOpts, new_rel)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipTransactor) ExtendTime(opts *bind.TransactOpts, nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "extendTime", nameHash, extend)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipSession) ExtendTime(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ExtendTime(&_BasTradableOwnership.TransactOpts, nameHash, extend)
}

// ExtendTime is a paid mutator transaction binding the contract method 0xc3f7beda.
//
// Solidity: function extendTime(bytes32 nameHash, uint256 extend) returns(uint256)
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) ExtendTime(nameHash [32]byte, extend *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.ExtendTime(&_BasTradableOwnership.TransactOpts, nameHash, extend)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) NewOwnership(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "newOwnership", nameHash, owner, expire)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) NewOwnership(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.NewOwnership(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// NewOwnership is a paid mutator transaction binding the contract method 0x4f9a1fbb.
//
// Solidity: function newOwnership(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) NewOwnership(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.NewOwnership(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) Revoke(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "revoke", nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) Revoke(nameHash [32]byte) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Revoke(&_BasTradableOwnership.TransactOpts, nameHash)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 nameHash) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) Revoke(nameHash [32]byte) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Revoke(&_BasTradableOwnership.TransactOpts, nameHash)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) Takeover(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "takeover", nameHash, owner, expire)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) Takeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Takeover(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// Takeover is a paid mutator transaction binding the contract method 0xb9566b10.
//
// Solidity: function takeover(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) Takeover(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Takeover(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) Transfer(opts *bind.TransactOpts, nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "transfer", nameHash, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) Transfer(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Transfer(&_BasTradableOwnership.TransactOpts, nameHash, to)
}

// Transfer is a paid mutator transaction binding the contract method 0x79ce9fac.
//
// Solidity: function transfer(bytes32 nameHash, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) Transfer(nameHash [32]byte, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.Transfer(&_BasTradableOwnership.TransactOpts, nameHash, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) TransferFrom(opts *bind.TransactOpts, nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "transferFrom", nameHash, from, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) TransferFrom(nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.TransferFrom(&_BasTradableOwnership.TransactOpts, nameHash, from, to)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x82d7bd27.
//
// Solidity: function transferFrom(bytes32 nameHash, address from, address to) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) TransferFrom(nameHash [32]byte, from common.Address, to common.Address) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.TransferFrom(&_BasTradableOwnership.TransactOpts, nameHash, from, to)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactor) UpdateByDaoProposal(opts *bind.TransactOpts, nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.contract.Transact(opts, "updateByDaoProposal", nameHash, owner, expire)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipSession) UpdateByDaoProposal(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.UpdateByDaoProposal(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// UpdateByDaoProposal is a paid mutator transaction binding the contract method 0x18eb3a50.
//
// Solidity: function updateByDaoProposal(bytes32 nameHash, address owner, uint256 expire) returns()
func (_BasTradableOwnership *BasTradableOwnershipTransactorSession) UpdateByDaoProposal(nameHash [32]byte, owner common.Address, expire *big.Int) (*types.Transaction, error) {
	return _BasTradableOwnership.Contract.UpdateByDaoProposal(&_BasTradableOwnership.TransactOpts, nameHash, owner, expire)
}

// BasTradableOwnershipAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the BasTradableOwnership contract.
type BasTradableOwnershipAddIterator struct {
	Event *BasTradableOwnershipAdd // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipAdd)
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
		it.Event = new(BasTradableOwnershipAdd)
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
func (it *BasTradableOwnershipAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipAdd represents a Add event raised by the BasTradableOwnership contract.
type BasTradableOwnershipAdd struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterAdd(opts *bind.FilterOpts) (*BasTradableOwnershipAddIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipAddIterator{contract: _BasTradableOwnership.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x5f17f6377b8024fbe396282bbccde46c3c704613385bd2f4b4b3ff955fd56a35.
//
// Solidity: event Add(bytes32 nameHash, address owner)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipAdd) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipAdd)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Add", log); err != nil {
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
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseAdd(log types.Log) (*BasTradableOwnershipAdd, error) {
	event := new(BasTradableOwnershipAdd)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BasTradableOwnership contract.
type BasTradableOwnershipApprovalIterator struct {
	Event *BasTradableOwnershipApproval // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipApproval)
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
		it.Event = new(BasTradableOwnershipApproval)
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
func (it *BasTradableOwnershipApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipApproval represents a Approval event raised by the BasTradableOwnership contract.
type BasTradableOwnershipApproval struct {
	From     common.Address
	To       common.Address
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x36a9e0c1da9cdc6d8f4bd4cb261f9ad6a45eb1641a557ead7530fbeff9a26336.
//
// Solidity: event Approval(address from, address to, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterApproval(opts *bind.FilterOpts) (*BasTradableOwnershipApprovalIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipApprovalIterator{contract: _BasTradableOwnership.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x36a9e0c1da9cdc6d8f4bd4cb261f9ad6a45eb1641a557ead7530fbeff9a26336.
//
// Solidity: event Approval(address from, address to, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipApproval) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipApproval)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x36a9e0c1da9cdc6d8f4bd4cb261f9ad6a45eb1641a557ead7530fbeff9a26336.
//
// Solidity: event Approval(address from, address to, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseApproval(log types.Log) (*BasTradableOwnershipApproval, error) {
	event := new(BasTradableOwnershipApproval)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the BasTradableOwnership contract.
type BasTradableOwnershipExtendIterator struct {
	Event *BasTradableOwnershipExtend // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipExtend)
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
		it.Event = new(BasTradableOwnershipExtend)
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
func (it *BasTradableOwnershipExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipExtend represents a Extend event raised by the BasTradableOwnership contract.
type BasTradableOwnershipExtend struct {
	NameHash [32]byte
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterExtend(opts *bind.FilterOpts) (*BasTradableOwnershipExtendIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipExtendIterator{contract: _BasTradableOwnership.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0xe898863b26adf3af54e82709ac1a76136eaaa5f3a7d9b790ce6d539d508e75ff.
//
// Solidity: event Extend(bytes32 nameHash, uint256 time)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipExtend) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipExtend)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
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
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseExtend(log types.Log) (*BasTradableOwnershipExtend, error) {
	event := new(BasTradableOwnershipExtend)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the BasTradableOwnership contract.
type BasTradableOwnershipRevokeIterator struct {
	Event *BasTradableOwnershipRevoke // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipRevoke)
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
		it.Event = new(BasTradableOwnershipRevoke)
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
func (it *BasTradableOwnershipRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipRevoke represents a Revoke event raised by the BasTradableOwnership contract.
type BasTradableOwnershipRevoke struct {
	From     common.Address
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address from, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterRevoke(opts *bind.FilterOpts) (*BasTradableOwnershipRevokeIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipRevokeIterator{contract: _BasTradableOwnership.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address from, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipRevoke) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Revoke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipRevoke)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0xc7fb647e59b18047309aa15aad418e5d7ca96d173ad704f1031a2c3d7591734b.
//
// Solidity: event Revoke(address from, bytes32 nameHash)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseRevoke(log types.Log) (*BasTradableOwnershipRevoke, error) {
	event := new(BasTradableOwnershipRevoke)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipTakeoverIterator is returned from FilterTakeover and is used to iterate over the raw logs and unpacked data for Takeover events raised by the BasTradableOwnership contract.
type BasTradableOwnershipTakeoverIterator struct {
	Event *BasTradableOwnershipTakeover // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipTakeoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipTakeover)
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
		it.Event = new(BasTradableOwnershipTakeover)
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
func (it *BasTradableOwnershipTakeoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipTakeoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipTakeover represents a Takeover event raised by the BasTradableOwnership contract.
type BasTradableOwnershipTakeover struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTakeover is a free log retrieval operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterTakeover(opts *bind.FilterOpts) (*BasTradableOwnershipTakeoverIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipTakeoverIterator{contract: _BasTradableOwnership.contract, event: "Takeover", logs: logs, sub: sub}, nil
}

// WatchTakeover is a free log subscription operation binding the contract event 0xae53ff61a227b196db64e2dea26e80d455bb2d90fd10372db78b44ace9cb0f62.
//
// Solidity: event Takeover(bytes32 nameHash, address from, address to)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchTakeover(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipTakeover) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Takeover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipTakeover)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
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
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseTakeover(log types.Log) (*BasTradableOwnershipTakeover, error) {
	event := new(BasTradableOwnershipTakeover)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Takeover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasTradableOwnership contract.
type BasTradableOwnershipTransferIterator struct {
	Event *BasTradableOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipTransfer)
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
		it.Event = new(BasTradableOwnershipTransfer)
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
func (it *BasTradableOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipTransfer represents a Transfer event raised by the BasTradableOwnership contract.
type BasTradableOwnershipTransfer struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*BasTradableOwnershipTransferIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipTransferIterator{contract: _BasTradableOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipTransfer)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x57972c0bcb114a9f52d3501767c95a65e93901ff48da6677c7399b5593b4e999.
//
// Solidity: event Transfer(bytes32 nameHash, address from, address to)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseTransfer(log types.Log) (*BasTradableOwnershipTransfer, error) {
	event := new(BasTradableOwnershipTransfer)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipTransferFromIterator is returned from FilterTransferFrom and is used to iterate over the raw logs and unpacked data for TransferFrom events raised by the BasTradableOwnership contract.
type BasTradableOwnershipTransferFromIterator struct {
	Event *BasTradableOwnershipTransferFrom // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipTransferFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipTransferFrom)
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
		it.Event = new(BasTradableOwnershipTransferFrom)
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
func (it *BasTradableOwnershipTransferFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipTransferFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipTransferFrom represents a TransferFrom event raised by the BasTradableOwnership contract.
type BasTradableOwnershipTransferFrom struct {
	NameHash [32]byte
	From     common.Address
	To       common.Address
	By       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferFrom is a free log retrieval operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterTransferFrom(opts *bind.FilterOpts) (*BasTradableOwnershipTransferFromIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "TransferFrom")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipTransferFromIterator{contract: _BasTradableOwnership.contract, event: "TransferFrom", logs: logs, sub: sub}, nil
}

// WatchTransferFrom is a free log subscription operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchTransferFrom(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipTransferFrom) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "TransferFrom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipTransferFrom)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "TransferFrom", log); err != nil {
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

// ParseTransferFrom is a log parse operation binding the contract event 0x3f8c833a27e9521421e46d68c5a31e603cc465fc4bbff2226b98d964252ebb3f.
//
// Solidity: event TransferFrom(bytes32 nameHash, address from, address to, address by)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseTransferFrom(log types.Log) (*BasTradableOwnershipTransferFrom, error) {
	event := new(BasTradableOwnershipTransferFrom)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "TransferFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasTradableOwnershipUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the BasTradableOwnership contract.
type BasTradableOwnershipUpdateIterator struct {
	Event *BasTradableOwnershipUpdate // Event containing the contract specifics and raw log

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
func (it *BasTradableOwnershipUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasTradableOwnershipUpdate)
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
		it.Event = new(BasTradableOwnershipUpdate)
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
func (it *BasTradableOwnershipUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasTradableOwnershipUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasTradableOwnershipUpdate represents a Update event raised by the BasTradableOwnership contract.
type BasTradableOwnershipUpdate struct {
	NameHash [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) FilterUpdate(opts *bind.FilterOpts) (*BasTradableOwnershipUpdateIterator, error) {

	logs, sub, err := _BasTradableOwnership.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &BasTradableOwnershipUpdateIterator{contract: _BasTradableOwnership.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xb2b848c23b147f2d3c11308c74cb0ff81a30b54628e5c73c4d70bf623eb98f72.
//
// Solidity: event Update(bytes32 nameHash, address owner)
func (_BasTradableOwnership *BasTradableOwnershipFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *BasTradableOwnershipUpdate) (event.Subscription, error) {

	logs, sub, err := _BasTradableOwnership.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasTradableOwnershipUpdate)
				if err := _BasTradableOwnership.contract.UnpackLog(event, "Update", log); err != nil {
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
func (_BasTradableOwnership *BasTradableOwnershipFilterer) ParseUpdate(log types.Log) (*BasTradableOwnershipUpdate, error) {
	event := new(BasTradableOwnershipUpdate)
	if err := _BasTradableOwnership.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
