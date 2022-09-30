// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc22

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

// DnsNameErc721MetaData contains all meta data concerning the DnsNameErc721 contract.
var DnsNameErc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"}],\"name\":\"DnsExtendExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DnsTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"idOwner_\",\"type\":\"address\"}],\"name\":\"MintId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SigUserAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721Owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gNameId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nameIdInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sigUser_\",\"type\":\"address\"}],\"name\":\"setSigUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"erc721Owner_\",\"type\":\"address\"}],\"name\":\"transferErc721Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DnsNameErc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsNameErc721MetaData.ABI instead.
var DnsNameErc721ABI = DnsNameErc721MetaData.ABI

// DnsNameErc721 is an auto generated Go binding around an Ethereum contract.
type DnsNameErc721 struct {
	DnsNameErc721Caller     // Read-only binding to the contract
	DnsNameErc721Transactor // Write-only binding to the contract
	DnsNameErc721Filterer   // Log filterer for contract events
}

// DnsNameErc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type DnsNameErc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsNameErc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsNameErc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsNameErc721Session struct {
	Contract     *DnsNameErc721    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsNameErc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsNameErc721CallerSession struct {
	Contract *DnsNameErc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DnsNameErc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsNameErc721TransactorSession struct {
	Contract     *DnsNameErc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DnsNameErc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type DnsNameErc721Raw struct {
	Contract *DnsNameErc721 // Generic contract binding to access the raw methods on
}

// DnsNameErc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsNameErc721CallerRaw struct {
	Contract *DnsNameErc721Caller // Generic read-only contract binding to access the raw methods on
}

// DnsNameErc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsNameErc721TransactorRaw struct {
	Contract *DnsNameErc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsNameErc721 creates a new instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721(address common.Address, backend bind.ContractBackend) (*DnsNameErc721, error) {
	contract, err := bindDnsNameErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721{DnsNameErc721Caller: DnsNameErc721Caller{contract: contract}, DnsNameErc721Transactor: DnsNameErc721Transactor{contract: contract}, DnsNameErc721Filterer: DnsNameErc721Filterer{contract: contract}}, nil
}

// NewDnsNameErc721Caller creates a new read-only instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Caller(address common.Address, caller bind.ContractCaller) (*DnsNameErc721Caller, error) {
	contract, err := bindDnsNameErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Caller{contract: contract}, nil
}

// NewDnsNameErc721Transactor creates a new write-only instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*DnsNameErc721Transactor, error) {
	contract, err := bindDnsNameErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Transactor{contract: contract}, nil
}

// NewDnsNameErc721Filterer creates a new log filterer instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*DnsNameErc721Filterer, error) {
	contract, err := bindDnsNameErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Filterer{contract: contract}, nil
}

// bindDnsNameErc721 binds a generic wrapper to an already deployed contract.
func bindDnsNameErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsNameErc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsNameErc721 *DnsNameErc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsNameErc721.Contract.DnsNameErc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsNameErc721 *DnsNameErc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsNameErc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsNameErc721 *DnsNameErc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsNameErc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsNameErc721 *DnsNameErc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsNameErc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsNameErc721 *DnsNameErc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsNameErc721 *DnsNameErc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.contract.Transact(opts, method, params...)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) SigUserAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "SigUserAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) SigUserAddr() (common.Address, error) {
	return _DnsNameErc721.Contract.SigUserAddr(&_DnsNameErc721.CallOpts)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) SigUserAddr() (common.Address, error) {
	return _DnsNameErc721.Contract.SigUserAddr(&_DnsNameErc721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsNameErc721.Contract.BalanceOf(&_DnsNameErc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsNameErc721.Contract.BalanceOf(&_DnsNameErc721.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721Caller) BaseUri(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721Session) BaseUri() ([]byte, error) {
	return _DnsNameErc721.Contract.BaseUri(&_DnsNameErc721.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721CallerSession) BaseUri() ([]byte, error) {
	return _DnsNameErc721.Contract.BaseUri(&_DnsNameErc721.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) Erc721Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "erc721Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) Erc721Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Erc721Owner(&_DnsNameErc721.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Erc721Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Erc721Owner(&_DnsNameErc721.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Caller) GNameId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "gNameId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) GNameId() (*big.Int, error) {
	return _DnsNameErc721.Contract.GNameId(&_DnsNameErc721.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721CallerSession) GNameId() (*big.Int, error) {
	return _DnsNameErc721.Contract.GNameId(&_DnsNameErc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.GetApproved(&_DnsNameErc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.GetApproved(&_DnsNameErc721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Session) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsNameErc721.Contract.IsApprovedForAll(&_DnsNameErc721.CallOpts, owner, operator_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721CallerSession) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsNameErc721.Contract.IsApprovedForAll(&_DnsNameErc721.CallOpts, owner, operator_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) Name() (string, error) {
	return _DnsNameErc721.Contract.Name(&_DnsNameErc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Name() (string, error) {
	return _DnsNameErc721.Contract.Name(&_DnsNameErc721.CallOpts)
}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721Caller) NameIdInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "nameIdInfo", arg0)

	outstruct := new(struct {
		ExpireTime uint32
		NameHash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpireTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.NameHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721Session) NameIdInfo(arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	return _DnsNameErc721.Contract.NameIdInfo(&_DnsNameErc721.CallOpts, arg0)
}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721CallerSession) NameIdInfo(arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	return _DnsNameErc721.Contract.NameIdInfo(&_DnsNameErc721.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Owner(&_DnsNameErc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Owner(&_DnsNameErc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.OwnerOf(&_DnsNameErc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.OwnerOf(&_DnsNameErc721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsNameErc721.Contract.SupportsInterface(&_DnsNameErc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsNameErc721.Contract.SupportsInterface(&_DnsNameErc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) Symbol() (string, error) {
	return _DnsNameErc721.Contract.Symbol(&_DnsNameErc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Symbol() (string, error) {
	return _DnsNameErc721.Contract.Symbol(&_DnsNameErc721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsNameErc721.Contract.TokenURI(&_DnsNameErc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsNameErc721.Contract.TokenURI(&_DnsNameErc721.CallOpts, tokenId)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) DnsExtendExpire(opts *bind.TransactOpts, tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "DnsExtendExpire", tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsExtendExpire(&_DnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsExtendExpire(&_DnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) DnsTransfer(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "DnsTransfer", to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsTransfer(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsTransfer(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Transactor) MintId(opts *bind.TransactOpts, hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "MintId", hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.MintId(&_DnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721TransactorSession) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.MintId(&_DnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.Approve(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.Approve(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom0(&_DnsNameErc721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom0(&_DnsNameErc721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setApprovalForAll", operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetApprovalForAll(&_DnsNameErc721.TransactOpts, operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetApprovalForAll(&_DnsNameErc721.TransactOpts, operator_, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetBaseUri(&_DnsNameErc721.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetBaseUri(&_DnsNameErc721.TransactOpts, baseUri_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetName(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setName", name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetName(&_DnsNameErc721.TransactOpts, name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetName(&_DnsNameErc721.TransactOpts, name_, symbol_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetSigUser(opts *bind.TransactOpts, sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setSigUser", sigUser_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetSigUser(sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetSigUser(&_DnsNameErc721.TransactOpts, sigUser_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetSigUser(sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetSigUser(&_DnsNameErc721.TransactOpts, sigUser_)
}

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) TransferErc721Owner(opts *bind.TransactOpts, erc721Owner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "transferErc721Owner", erc721Owner_)
}

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) TransferErc721Owner(erc721Owner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferErc721Owner(&_DnsNameErc721.TransactOpts, erc721Owner_)
}

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) TransferErc721Owner(erc721Owner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferErc721Owner(&_DnsNameErc721.TransactOpts, erc721Owner_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferOwnership(&_DnsNameErc721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferOwnership(&_DnsNameErc721.TransactOpts, newOwner)
}

// DnsNameErc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalIterator struct {
	Event *DnsNameErc721Approval // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721Approval)
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
		it.Event = new(DnsNameErc721Approval)
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
func (it *DnsNameErc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721Approval represents a Approval event raised by the DnsNameErc721 contract.
type DnsNameErc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DnsNameErc721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721ApprovalIterator{contract: _DnsNameErc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DnsNameErc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721Approval)
				if err := _DnsNameErc721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseApproval(log types.Log) (*DnsNameErc721Approval, error) {
	event := new(DnsNameErc721Approval)
	if err := _DnsNameErc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameErc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalForAllIterator struct {
	Event *DnsNameErc721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721ApprovalForAll)
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
		it.Event = new(DnsNameErc721ApprovalForAll)
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
func (it *DnsNameErc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721ApprovalForAll represents a ApprovalForAll event raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DnsNameErc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721ApprovalForAllIterator{contract: _DnsNameErc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DnsNameErc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721ApprovalForAll)
				if err := _DnsNameErc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseApprovalForAll(log types.Log) (*DnsNameErc721ApprovalForAll, error) {
	event := new(DnsNameErc721ApprovalForAll)
	if err := _DnsNameErc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameErc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DnsNameErc721 contract.
type DnsNameErc721TransferIterator struct {
	Event *DnsNameErc721Transfer // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721Transfer)
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
		it.Event = new(DnsNameErc721Transfer)
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
func (it *DnsNameErc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721Transfer represents a Transfer event raised by the DnsNameErc721 contract.
type DnsNameErc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DnsNameErc721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721TransferIterator{contract: _DnsNameErc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DnsNameErc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721Transfer)
				if err := _DnsNameErc721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseTransfer(log types.Log) (*DnsNameErc721Transfer, error) {
	event := new(DnsNameErc721Transfer)
	if err := _DnsNameErc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
