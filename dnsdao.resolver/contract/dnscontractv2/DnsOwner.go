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

// DnsOwnersMetaData contains all meta data concerning the DnsOwners contract.
var DnsOwnersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"EvAddAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nameOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"EvUpdateOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nameOwner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nameOwner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"entireName_\",\"type\":\"bytes\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nameOwner_\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"addAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractAddrs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dnsOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"dnsOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"entireName\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"}],\"name\":\"getEntireName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"getOwnerByContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"}],\"name\":\"getOwnerByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"}],\"name\":\"getUDIDSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"operatorIsAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DnsOwnersABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsOwnersMetaData.ABI instead.
var DnsOwnersABI = DnsOwnersMetaData.ABI

// DnsOwners is an auto generated Go binding around an Ethereum contract.
type DnsOwners struct {
	DnsOwnersCaller     // Read-only binding to the contract
	DnsOwnersTransactor // Write-only binding to the contract
	DnsOwnersFilterer   // Log filterer for contract events
}

// DnsOwnersCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsOwnersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsOwnersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsOwnersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsOwnersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsOwnersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsOwnersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsOwnersSession struct {
	Contract     *DnsOwners        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsOwnersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsOwnersCallerSession struct {
	Contract *DnsOwnersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DnsOwnersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsOwnersTransactorSession struct {
	Contract     *DnsOwnersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DnsOwnersRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsOwnersRaw struct {
	Contract *DnsOwners // Generic contract binding to access the raw methods on
}

// DnsOwnersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsOwnersCallerRaw struct {
	Contract *DnsOwnersCaller // Generic read-only contract binding to access the raw methods on
}

// DnsOwnersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsOwnersTransactorRaw struct {
	Contract *DnsOwnersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsOwners creates a new instance of DnsOwners, bound to a specific deployed contract.
func NewDnsOwners(address common.Address, backend bind.ContractBackend) (*DnsOwners, error) {
	contract, err := bindDnsOwners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsOwners{DnsOwnersCaller: DnsOwnersCaller{contract: contract}, DnsOwnersTransactor: DnsOwnersTransactor{contract: contract}, DnsOwnersFilterer: DnsOwnersFilterer{contract: contract}}, nil
}

// NewDnsOwnersCaller creates a new read-only instance of DnsOwners, bound to a specific deployed contract.
func NewDnsOwnersCaller(address common.Address, caller bind.ContractCaller) (*DnsOwnersCaller, error) {
	contract, err := bindDnsOwners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsOwnersCaller{contract: contract}, nil
}

// NewDnsOwnersTransactor creates a new write-only instance of DnsOwners, bound to a specific deployed contract.
func NewDnsOwnersTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsOwnersTransactor, error) {
	contract, err := bindDnsOwners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsOwnersTransactor{contract: contract}, nil
}

// NewDnsOwnersFilterer creates a new log filterer instance of DnsOwners, bound to a specific deployed contract.
func NewDnsOwnersFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsOwnersFilterer, error) {
	contract, err := bindDnsOwners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsOwnersFilterer{contract: contract}, nil
}

// bindDnsOwners binds a generic wrapper to an already deployed contract.
func bindDnsOwners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsOwnersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsOwners *DnsOwnersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsOwners.Contract.DnsOwnersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsOwners *DnsOwnersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsOwners.Contract.DnsOwnersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsOwners *DnsOwnersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsOwners.Contract.DnsOwnersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsOwners *DnsOwnersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsOwners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsOwners *DnsOwnersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsOwners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsOwners *DnsOwnersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsOwners.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(bool)
func (_DnsOwners *DnsOwnersCaller) Allowance(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "allowance", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(bool)
func (_DnsOwners *DnsOwnersSession) Allowance(arg0 common.Address) (bool, error) {
	return _DnsOwners.Contract.Allowance(&_DnsOwners.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(bool)
func (_DnsOwners *DnsOwnersCallerSession) Allowance(arg0 common.Address) (bool, error) {
	return _DnsOwners.Contract.Allowance(&_DnsOwners.CallOpts, arg0)
}

// ContractAddrs is a free data retrieval call binding the contract method 0x117c2111.
//
// Solidity: function contractAddrs(address ) view returns(bytes32)
func (_DnsOwners *DnsOwnersCaller) ContractAddrs(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "contractAddrs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddrs is a free data retrieval call binding the contract method 0x117c2111.
//
// Solidity: function contractAddrs(address ) view returns(bytes32)
func (_DnsOwners *DnsOwnersSession) ContractAddrs(arg0 common.Address) ([32]byte, error) {
	return _DnsOwners.Contract.ContractAddrs(&_DnsOwners.CallOpts, arg0)
}

// ContractAddrs is a free data retrieval call binding the contract method 0x117c2111.
//
// Solidity: function contractAddrs(address ) view returns(bytes32)
func (_DnsOwners *DnsOwnersCallerSession) ContractAddrs(arg0 common.Address) ([32]byte, error) {
	return _DnsOwners.Contract.ContractAddrs(&_DnsOwners.CallOpts, arg0)
}

// DnsOwners is a free data retrieval call binding the contract method 0x854e49ee.
//
// Solidity: function dnsOwners(bytes32 ) view returns(address dnsOwner, uint256 tokenId, bytes entireName)
func (_DnsOwners *DnsOwnersCaller) DnsOwners(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DnsOwner   common.Address
	TokenId    *big.Int
	EntireName []byte
}, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "dnsOwners", arg0)

	outstruct := new(struct {
		DnsOwner   common.Address
		TokenId    *big.Int
		EntireName []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DnsOwner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EntireName = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DnsOwners is a free data retrieval call binding the contract method 0x854e49ee.
//
// Solidity: function dnsOwners(bytes32 ) view returns(address dnsOwner, uint256 tokenId, bytes entireName)
func (_DnsOwners *DnsOwnersSession) DnsOwners(arg0 [32]byte) (struct {
	DnsOwner   common.Address
	TokenId    *big.Int
	EntireName []byte
}, error) {
	return _DnsOwners.Contract.DnsOwners(&_DnsOwners.CallOpts, arg0)
}

// DnsOwners is a free data retrieval call binding the contract method 0x854e49ee.
//
// Solidity: function dnsOwners(bytes32 ) view returns(address dnsOwner, uint256 tokenId, bytes entireName)
func (_DnsOwners *DnsOwnersCallerSession) DnsOwners(arg0 [32]byte) (struct {
	DnsOwner   common.Address
	TokenId    *big.Int
	EntireName []byte
}, error) {
	return _DnsOwners.Contract.DnsOwners(&_DnsOwners.CallOpts, arg0)
}

// GetEntireName is a free data retrieval call binding the contract method 0x6e0a1453.
//
// Solidity: function getEntireName(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersCaller) GetEntireName(opts *bind.CallOpts, nameHash_ [32]byte) (string, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "getEntireName", nameHash_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEntireName is a free data retrieval call binding the contract method 0x6e0a1453.
//
// Solidity: function getEntireName(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersSession) GetEntireName(nameHash_ [32]byte) (string, error) {
	return _DnsOwners.Contract.GetEntireName(&_DnsOwners.CallOpts, nameHash_)
}

// GetEntireName is a free data retrieval call binding the contract method 0x6e0a1453.
//
// Solidity: function getEntireName(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersCallerSession) GetEntireName(nameHash_ [32]byte) (string, error) {
	return _DnsOwners.Contract.GetEntireName(&_DnsOwners.CallOpts, nameHash_)
}

// GetOwnerByContract is a free data retrieval call binding the contract method 0x2ce75d18.
//
// Solidity: function getOwnerByContract(address contractAddr_) view returns(address)
func (_DnsOwners *DnsOwnersCaller) GetOwnerByContract(opts *bind.CallOpts, contractAddr_ common.Address) (common.Address, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "getOwnerByContract", contractAddr_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerByContract is a free data retrieval call binding the contract method 0x2ce75d18.
//
// Solidity: function getOwnerByContract(address contractAddr_) view returns(address)
func (_DnsOwners *DnsOwnersSession) GetOwnerByContract(contractAddr_ common.Address) (common.Address, error) {
	return _DnsOwners.Contract.GetOwnerByContract(&_DnsOwners.CallOpts, contractAddr_)
}

// GetOwnerByContract is a free data retrieval call binding the contract method 0x2ce75d18.
//
// Solidity: function getOwnerByContract(address contractAddr_) view returns(address)
func (_DnsOwners *DnsOwnersCallerSession) GetOwnerByContract(contractAddr_ common.Address) (common.Address, error) {
	return _DnsOwners.Contract.GetOwnerByContract(&_DnsOwners.CallOpts, contractAddr_)
}

// GetOwnerByHash is a free data retrieval call binding the contract method 0x549c0f30.
//
// Solidity: function getOwnerByHash(bytes32 nameHash_) view returns(address)
func (_DnsOwners *DnsOwnersCaller) GetOwnerByHash(opts *bind.CallOpts, nameHash_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "getOwnerByHash", nameHash_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerByHash is a free data retrieval call binding the contract method 0x549c0f30.
//
// Solidity: function getOwnerByHash(bytes32 nameHash_) view returns(address)
func (_DnsOwners *DnsOwnersSession) GetOwnerByHash(nameHash_ [32]byte) (common.Address, error) {
	return _DnsOwners.Contract.GetOwnerByHash(&_DnsOwners.CallOpts, nameHash_)
}

// GetOwnerByHash is a free data retrieval call binding the contract method 0x549c0f30.
//
// Solidity: function getOwnerByHash(bytes32 nameHash_) view returns(address)
func (_DnsOwners *DnsOwnersCallerSession) GetOwnerByHash(nameHash_ [32]byte) (common.Address, error) {
	return _DnsOwners.Contract.GetOwnerByHash(&_DnsOwners.CallOpts, nameHash_)
}

// GetUDIDSymbol is a free data retrieval call binding the contract method 0x8a6343f9.
//
// Solidity: function getUDIDSymbol(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersCaller) GetUDIDSymbol(opts *bind.CallOpts, nameHash_ [32]byte) (string, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "getUDIDSymbol", nameHash_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUDIDSymbol is a free data retrieval call binding the contract method 0x8a6343f9.
//
// Solidity: function getUDIDSymbol(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersSession) GetUDIDSymbol(nameHash_ [32]byte) (string, error) {
	return _DnsOwners.Contract.GetUDIDSymbol(&_DnsOwners.CallOpts, nameHash_)
}

// GetUDIDSymbol is a free data retrieval call binding the contract method 0x8a6343f9.
//
// Solidity: function getUDIDSymbol(bytes32 nameHash_) view returns(string)
func (_DnsOwners *DnsOwnersCallerSession) GetUDIDSymbol(nameHash_ [32]byte) (string, error) {
	return _DnsOwners.Contract.GetUDIDSymbol(&_DnsOwners.CallOpts, nameHash_)
}

// OperatorIsAllowance is a free data retrieval call binding the contract method 0x226297ad.
//
// Solidity: function operatorIsAllowance(address contractAddr_) view returns(bool)
func (_DnsOwners *DnsOwnersCaller) OperatorIsAllowance(opts *bind.CallOpts, contractAddr_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "operatorIsAllowance", contractAddr_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorIsAllowance is a free data retrieval call binding the contract method 0x226297ad.
//
// Solidity: function operatorIsAllowance(address contractAddr_) view returns(bool)
func (_DnsOwners *DnsOwnersSession) OperatorIsAllowance(contractAddr_ common.Address) (bool, error) {
	return _DnsOwners.Contract.OperatorIsAllowance(&_DnsOwners.CallOpts, contractAddr_)
}

// OperatorIsAllowance is a free data retrieval call binding the contract method 0x226297ad.
//
// Solidity: function operatorIsAllowance(address contractAddr_) view returns(bool)
func (_DnsOwners *DnsOwnersCallerSession) OperatorIsAllowance(contractAddr_ common.Address) (bool, error) {
	return _DnsOwners.Contract.OperatorIsAllowance(&_DnsOwners.CallOpts, contractAddr_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsOwners *DnsOwnersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsOwners.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsOwners *DnsOwnersSession) Owner() (common.Address, error) {
	return _DnsOwners.Contract.Owner(&_DnsOwners.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsOwners *DnsOwnersCallerSession) Owner() (common.Address, error) {
	return _DnsOwners.Contract.Owner(&_DnsOwners.CallOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x2c782d5a.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, address contractAddr_) returns()
func (_DnsOwners *DnsOwnersTransactor) UpdateOwner(opts *bind.TransactOpts, nameHash_ [32]byte, nameOwner_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.contract.Transact(opts, "UpdateOwner", nameHash_, nameOwner_, contractAddr_)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x2c782d5a.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, address contractAddr_) returns()
func (_DnsOwners *DnsOwnersSession) UpdateOwner(nameHash_ [32]byte, nameOwner_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner(&_DnsOwners.TransactOpts, nameHash_, nameOwner_, contractAddr_)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x2c782d5a.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, address contractAddr_) returns()
func (_DnsOwners *DnsOwnersTransactorSession) UpdateOwner(nameHash_ [32]byte, nameOwner_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner(&_DnsOwners.TransactOpts, nameHash_, nameOwner_, contractAddr_)
}

// UpdateOwner0 is a paid mutator transaction binding the contract method 0xdd78fe0d.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, uint256 tokenId_, bytes entireName_) returns()
func (_DnsOwners *DnsOwnersTransactor) UpdateOwner0(opts *bind.TransactOpts, nameHash_ [32]byte, nameOwner_ common.Address, tokenId_ *big.Int, entireName_ []byte) (*types.Transaction, error) {
	return _DnsOwners.contract.Transact(opts, "UpdateOwner0", nameHash_, nameOwner_, tokenId_, entireName_)
}

// UpdateOwner0 is a paid mutator transaction binding the contract method 0xdd78fe0d.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, uint256 tokenId_, bytes entireName_) returns()
func (_DnsOwners *DnsOwnersSession) UpdateOwner0(nameHash_ [32]byte, nameOwner_ common.Address, tokenId_ *big.Int, entireName_ []byte) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner0(&_DnsOwners.TransactOpts, nameHash_, nameOwner_, tokenId_, entireName_)
}

// UpdateOwner0 is a paid mutator transaction binding the contract method 0xdd78fe0d.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_, uint256 tokenId_, bytes entireName_) returns()
func (_DnsOwners *DnsOwnersTransactorSession) UpdateOwner0(nameHash_ [32]byte, nameOwner_ common.Address, tokenId_ *big.Int, entireName_ []byte) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner0(&_DnsOwners.TransactOpts, nameHash_, nameOwner_, tokenId_, entireName_)
}

// UpdateOwner1 is a paid mutator transaction binding the contract method 0xe8f8f4a9.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_) returns()
func (_DnsOwners *DnsOwnersTransactor) UpdateOwner1(opts *bind.TransactOpts, nameHash_ [32]byte, nameOwner_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.contract.Transact(opts, "UpdateOwner1", nameHash_, nameOwner_)
}

// UpdateOwner1 is a paid mutator transaction binding the contract method 0xe8f8f4a9.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_) returns()
func (_DnsOwners *DnsOwnersSession) UpdateOwner1(nameHash_ [32]byte, nameOwner_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner1(&_DnsOwners.TransactOpts, nameHash_, nameOwner_)
}

// UpdateOwner1 is a paid mutator transaction binding the contract method 0xe8f8f4a9.
//
// Solidity: function UpdateOwner(bytes32 nameHash_, address nameOwner_) returns()
func (_DnsOwners *DnsOwnersTransactorSession) UpdateOwner1(nameHash_ [32]byte, nameOwner_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.UpdateOwner1(&_DnsOwners.TransactOpts, nameHash_, nameOwner_)
}

// AddAllowance is a paid mutator transaction binding the contract method 0xe94f1292.
//
// Solidity: function addAllowance(address user_) returns()
func (_DnsOwners *DnsOwnersTransactor) AddAllowance(opts *bind.TransactOpts, user_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.contract.Transact(opts, "addAllowance", user_)
}

// AddAllowance is a paid mutator transaction binding the contract method 0xe94f1292.
//
// Solidity: function addAllowance(address user_) returns()
func (_DnsOwners *DnsOwnersSession) AddAllowance(user_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.AddAllowance(&_DnsOwners.TransactOpts, user_)
}

// AddAllowance is a paid mutator transaction binding the contract method 0xe94f1292.
//
// Solidity: function addAllowance(address user_) returns()
func (_DnsOwners *DnsOwnersTransactorSession) AddAllowance(user_ common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.AddAllowance(&_DnsOwners.TransactOpts, user_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsOwners *DnsOwnersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsOwners.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsOwners *DnsOwnersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.TransferOwnership(&_DnsOwners.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsOwners *DnsOwnersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsOwners.Contract.TransferOwnership(&_DnsOwners.TransactOpts, newOwner)
}

// DnsOwnersEvAddAllowanceIterator is returned from FilterEvAddAllowance and is used to iterate over the raw logs and unpacked data for EvAddAllowance events raised by the DnsOwners contract.
type DnsOwnersEvAddAllowanceIterator struct {
	Event *DnsOwnersEvAddAllowance // Event containing the contract specifics and raw log

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
func (it *DnsOwnersEvAddAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsOwnersEvAddAllowance)
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
		it.Event = new(DnsOwnersEvAddAllowance)
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
func (it *DnsOwnersEvAddAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsOwnersEvAddAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsOwnersEvAddAllowance represents a EvAddAllowance event raised by the DnsOwners contract.
type DnsOwnersEvAddAllowance struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvAddAllowance is a free log retrieval operation binding the contract event 0xbc0da13b49e39e0b99202daf10ae08633dbc60b8a51fa88897e80a6febf4006d.
//
// Solidity: event EvAddAllowance(address user)
func (_DnsOwners *DnsOwnersFilterer) FilterEvAddAllowance(opts *bind.FilterOpts) (*DnsOwnersEvAddAllowanceIterator, error) {

	logs, sub, err := _DnsOwners.contract.FilterLogs(opts, "EvAddAllowance")
	if err != nil {
		return nil, err
	}
	return &DnsOwnersEvAddAllowanceIterator{contract: _DnsOwners.contract, event: "EvAddAllowance", logs: logs, sub: sub}, nil
}

// WatchEvAddAllowance is a free log subscription operation binding the contract event 0xbc0da13b49e39e0b99202daf10ae08633dbc60b8a51fa88897e80a6febf4006d.
//
// Solidity: event EvAddAllowance(address user)
func (_DnsOwners *DnsOwnersFilterer) WatchEvAddAllowance(opts *bind.WatchOpts, sink chan<- *DnsOwnersEvAddAllowance) (event.Subscription, error) {

	logs, sub, err := _DnsOwners.contract.WatchLogs(opts, "EvAddAllowance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsOwnersEvAddAllowance)
				if err := _DnsOwners.contract.UnpackLog(event, "EvAddAllowance", log); err != nil {
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

// ParseEvAddAllowance is a log parse operation binding the contract event 0xbc0da13b49e39e0b99202daf10ae08633dbc60b8a51fa88897e80a6febf4006d.
//
// Solidity: event EvAddAllowance(address user)
func (_DnsOwners *DnsOwnersFilterer) ParseEvAddAllowance(log types.Log) (*DnsOwnersEvAddAllowance, error) {
	event := new(DnsOwnersEvAddAllowance)
	if err := _DnsOwners.contract.UnpackLog(event, "EvAddAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsOwnersEvUpdateOwnerIterator is returned from FilterEvUpdateOwner and is used to iterate over the raw logs and unpacked data for EvUpdateOwner events raised by the DnsOwners contract.
type DnsOwnersEvUpdateOwnerIterator struct {
	Event *DnsOwnersEvUpdateOwner // Event containing the contract specifics and raw log

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
func (it *DnsOwnersEvUpdateOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsOwnersEvUpdateOwner)
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
		it.Event = new(DnsOwnersEvUpdateOwner)
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
func (it *DnsOwnersEvUpdateOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsOwnersEvUpdateOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsOwnersEvUpdateOwner represents a EvUpdateOwner event raised by the DnsOwners contract.
type DnsOwnersEvUpdateOwner struct {
	NameHash     [32]byte
	NameOwner    common.Address
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvUpdateOwner is a free log retrieval operation binding the contract event 0x5773251697990866ab1ede04ca89d5141f280c313225d194e3a1f195b79d4c0e.
//
// Solidity: event EvUpdateOwner(bytes32 nameHash, address nameOwner, address contractAddr)
func (_DnsOwners *DnsOwnersFilterer) FilterEvUpdateOwner(opts *bind.FilterOpts) (*DnsOwnersEvUpdateOwnerIterator, error) {

	logs, sub, err := _DnsOwners.contract.FilterLogs(opts, "EvUpdateOwner")
	if err != nil {
		return nil, err
	}
	return &DnsOwnersEvUpdateOwnerIterator{contract: _DnsOwners.contract, event: "EvUpdateOwner", logs: logs, sub: sub}, nil
}

// WatchEvUpdateOwner is a free log subscription operation binding the contract event 0x5773251697990866ab1ede04ca89d5141f280c313225d194e3a1f195b79d4c0e.
//
// Solidity: event EvUpdateOwner(bytes32 nameHash, address nameOwner, address contractAddr)
func (_DnsOwners *DnsOwnersFilterer) WatchEvUpdateOwner(opts *bind.WatchOpts, sink chan<- *DnsOwnersEvUpdateOwner) (event.Subscription, error) {

	logs, sub, err := _DnsOwners.contract.WatchLogs(opts, "EvUpdateOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsOwnersEvUpdateOwner)
				if err := _DnsOwners.contract.UnpackLog(event, "EvUpdateOwner", log); err != nil {
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

// ParseEvUpdateOwner is a log parse operation binding the contract event 0x5773251697990866ab1ede04ca89d5141f280c313225d194e3a1f195b79d4c0e.
//
// Solidity: event EvUpdateOwner(bytes32 nameHash, address nameOwner, address contractAddr)
func (_DnsOwners *DnsOwnersFilterer) ParseEvUpdateOwner(log types.Log) (*DnsOwnersEvUpdateOwner, error) {
	event := new(DnsOwnersEvUpdateOwner)
	if err := _DnsOwners.contract.UnpackLog(event, "EvUpdateOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
