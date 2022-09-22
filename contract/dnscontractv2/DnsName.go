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

// DnsNameMetaData contains all meta data concerning the DnsName contract.
var DnsNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransferOwner\",\"type\":\"bool\"}],\"name\":\"EvChargeDnsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"}],\"name\":\"EvMintDnsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721Addr\",\"type\":\"address\"}],\"name\":\"EvNewSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"EvOpenToReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"transfer_\",\"type\":\"bool\"}],\"name\":\"ChargeName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"passId_\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gNameId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"id2Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"entireName\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openToReg\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"erc721Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"passCardUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taxC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountantC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sigUser_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"open_\",\"type\":\"bool\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"open_\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"setRegistered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DnsNameABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsNameMetaData.ABI instead.
var DnsNameABI = DnsNameMetaData.ABI

// DnsName is an auto generated Go binding around an Ethereum contract.
type DnsName struct {
	DnsNameCaller     // Read-only binding to the contract
	DnsNameTransactor // Write-only binding to the contract
	DnsNameFilterer   // Log filterer for contract events
}

// DnsNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsNameSession struct {
	Contract     *DnsName          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsNameCallerSession struct {
	Contract *DnsNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DnsNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsNameTransactorSession struct {
	Contract     *DnsNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DnsNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsNameRaw struct {
	Contract *DnsName // Generic contract binding to access the raw methods on
}

// DnsNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsNameCallerRaw struct {
	Contract *DnsNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsNameTransactorRaw struct {
	Contract *DnsNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsName creates a new instance of DnsName, bound to a specific deployed contract.
func NewDnsName(address common.Address, backend bind.ContractBackend) (*DnsName, error) {
	contract, err := bindDnsName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsName{DnsNameCaller: DnsNameCaller{contract: contract}, DnsNameTransactor: DnsNameTransactor{contract: contract}, DnsNameFilterer: DnsNameFilterer{contract: contract}}, nil
}

// NewDnsNameCaller creates a new read-only instance of DnsName, bound to a specific deployed contract.
func NewDnsNameCaller(address common.Address, caller bind.ContractCaller) (*DnsNameCaller, error) {
	contract, err := bindDnsName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameCaller{contract: contract}, nil
}

// NewDnsNameTransactor creates a new write-only instance of DnsName, bound to a specific deployed contract.
func NewDnsNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsNameTransactor, error) {
	contract, err := bindDnsName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameTransactor{contract: contract}, nil
}

// NewDnsNameFilterer creates a new log filterer instance of DnsName, bound to a specific deployed contract.
func NewDnsNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsNameFilterer, error) {
	contract, err := bindDnsName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsNameFilterer{contract: contract}, nil
}

// bindDnsName binds a generic wrapper to an already deployed contract.
func bindDnsName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsName *DnsNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsName.Contract.DnsNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsName *DnsNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsName.Contract.DnsNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsName *DnsNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsName.Contract.DnsNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsName *DnsNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsName *DnsNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsName *DnsNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsName *DnsNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsName *DnsNameSession) AccountantC() (common.Address, error) {
	return _DnsName.Contract.AccountantC(&_DnsName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsName *DnsNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsName.Contract.AccountantC(&_DnsName.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsName *DnsNameCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsName *DnsNameSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsName.Contract.BalanceOf(&_DnsName.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsName *DnsNameCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsName.Contract.BalanceOf(&_DnsName.CallOpts, owner)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsName *DnsNameCaller) GNameId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "gNameId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsName *DnsNameSession) GNameId() (*big.Int, error) {
	return _DnsName.Contract.GNameId(&_DnsName.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsName *DnsNameCallerSession) GNameId() (*big.Int, error) {
	return _DnsName.Contract.GNameId(&_DnsName.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsName.Contract.GetApproved(&_DnsName.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsName.Contract.GetApproved(&_DnsName.CallOpts, tokenId)
}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsName *DnsNameCaller) Id2Hash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "id2Hash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsName *DnsNameSession) Id2Hash(arg0 *big.Int) ([32]byte, error) {
	return _DnsName.Contract.Id2Hash(&_DnsName.CallOpts, arg0)
}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsName *DnsNameCallerSession) Id2Hash(arg0 *big.Int) ([32]byte, error) {
	return _DnsName.Contract.Id2Hash(&_DnsName.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsName *DnsNameCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "isApprovedForAll", owner, operator_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsName *DnsNameSession) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsName.Contract.IsApprovedForAll(&_DnsName.CallOpts, owner, operator_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsName *DnsNameCallerSession) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsName.Contract.IsApprovedForAll(&_DnsName.CallOpts, owner, operator_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsName *DnsNameCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsName *DnsNameSession) Name() (string, error) {
	return _DnsName.Contract.Name(&_DnsName.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsName *DnsNameCallerSession) Name() (string, error) {
	return _DnsName.Contract.Name(&_DnsName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, bool openToReg, address erc721Addr, uint256 expireTime, uint256 tokenId)
func (_DnsName *DnsNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EntireName []byte
	OpenToReg  bool
	Erc721Addr common.Address
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "nameStore", arg0)

	outstruct := new(struct {
		EntireName []byte
		OpenToReg  bool
		Erc721Addr common.Address
		ExpireTime *big.Int
		TokenId    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EntireName = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.OpenToReg = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Erc721Addr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ExpireTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, bool openToReg, address erc721Addr, uint256 expireTime, uint256 tokenId)
func (_DnsName *DnsNameSession) NameStore(arg0 [32]byte) (struct {
	EntireName []byte
	OpenToReg  bool
	Erc721Addr common.Address
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	return _DnsName.Contract.NameStore(&_DnsName.CallOpts, arg0)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, bool openToReg, address erc721Addr, uint256 expireTime, uint256 tokenId)
func (_DnsName *DnsNameCallerSession) NameStore(arg0 [32]byte) (struct {
	EntireName []byte
	OpenToReg  bool
	Erc721Addr common.Address
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	return _DnsName.Contract.NameStore(&_DnsName.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsName *DnsNameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsName *DnsNameSession) Owner() (common.Address, error) {
	return _DnsName.Contract.Owner(&_DnsName.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsName *DnsNameCallerSession) Owner() (common.Address, error) {
	return _DnsName.Contract.Owner(&_DnsName.CallOpts)
}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsName *DnsNameCaller) OwnerC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "ownerC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsName *DnsNameSession) OwnerC() (common.Address, error) {
	return _DnsName.Contract.OwnerC(&_DnsName.CallOpts)
}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsName *DnsNameCallerSession) OwnerC() (common.Address, error) {
	return _DnsName.Contract.OwnerC(&_DnsName.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsName.Contract.OwnerOf(&_DnsName.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsName *DnsNameCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsName.Contract.OwnerOf(&_DnsName.CallOpts, tokenId)
}

// PassCardUsed is a free data retrieval call binding the contract method 0x3be10505.
//
// Solidity: function passCardUsed(uint32 ) view returns(bool)
func (_DnsName *DnsNameCaller) PassCardUsed(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "passCardUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PassCardUsed is a free data retrieval call binding the contract method 0x3be10505.
//
// Solidity: function passCardUsed(uint32 ) view returns(bool)
func (_DnsName *DnsNameSession) PassCardUsed(arg0 uint32) (bool, error) {
	return _DnsName.Contract.PassCardUsed(&_DnsName.CallOpts, arg0)
}

// PassCardUsed is a free data retrieval call binding the contract method 0x3be10505.
//
// Solidity: function passCardUsed(uint32 ) view returns(bool)
func (_DnsName *DnsNameCallerSession) PassCardUsed(arg0 uint32) (bool, error) {
	return _DnsName.Contract.PassCardUsed(&_DnsName.CallOpts, arg0)
}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsName *DnsNameCaller) PriceC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "priceC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsName *DnsNameSession) PriceC() (common.Address, error) {
	return _DnsName.Contract.PriceC(&_DnsName.CallOpts)
}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsName *DnsNameCallerSession) PriceC() (common.Address, error) {
	return _DnsName.Contract.PriceC(&_DnsName.CallOpts)
}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsName *DnsNameCaller) SigUser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "sigUser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsName *DnsNameSession) SigUser() (common.Address, error) {
	return _DnsName.Contract.SigUser(&_DnsName.CallOpts)
}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsName *DnsNameCallerSession) SigUser() (common.Address, error) {
	return _DnsName.Contract.SigUser(&_DnsName.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsName *DnsNameCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsName *DnsNameSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsName.Contract.SupportsInterface(&_DnsName.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsName *DnsNameCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsName.Contract.SupportsInterface(&_DnsName.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsName *DnsNameCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsName *DnsNameSession) Symbol() (string, error) {
	return _DnsName.Contract.Symbol(&_DnsName.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsName *DnsNameCallerSession) Symbol() (string, error) {
	return _DnsName.Contract.Symbol(&_DnsName.CallOpts)
}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsName *DnsNameCaller) TaxC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "taxC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsName *DnsNameSession) TaxC() (common.Address, error) {
	return _DnsName.Contract.TaxC(&_DnsName.CallOpts)
}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsName *DnsNameCallerSession) TaxC() (common.Address, error) {
	return _DnsName.Contract.TaxC(&_DnsName.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsName *DnsNameCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DnsName.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsName *DnsNameSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsName.Contract.TokenURI(&_DnsName.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsName *DnsNameCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsName.Contract.TokenURI(&_DnsName.CallOpts, tokenId)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsName *DnsNameTransactor) ChargeName(opts *bind.TransactOpts, erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "ChargeName", erc20Addr_, nameHash_, year_, transfer_)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsName *DnsNameSession) ChargeName(erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsName.Contract.ChargeName(&_DnsName.TransactOpts, erc20Addr_, nameHash_, year_, transfer_)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsName *DnsNameTransactorSession) ChargeName(erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsName.Contract.ChargeName(&_DnsName.TransactOpts, erc20Addr_, nameHash_, year_, transfer_)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsName *DnsNameTransactor) MintNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "MintNameBySig", entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsName *DnsNameSession) MintNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsName.Contract.MintNameBySig(&_DnsName.TransactOpts, entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsName *DnsNameTransactorSession) MintNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsName.Contract.MintNameBySig(&_DnsName.TransactOpts, entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsName *DnsNameSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.Approve(&_DnsName.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.Approve(&_DnsName.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.SafeTransferFrom(&_DnsName.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.SafeTransferFrom(&_DnsName.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsName *DnsNameTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsName *DnsNameSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsName.Contract.SafeTransferFrom0(&_DnsName.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsName *DnsNameTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsName.Contract.SafeTransferFrom0(&_DnsName.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsName *DnsNameTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "setApprovalForAll", operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsName *DnsNameSession) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsName.Contract.SetApprovalForAll(&_DnsName.TransactOpts, operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsName *DnsNameTransactorSession) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsName.Contract.SetApprovalForAll(&_DnsName.TransactOpts, operator_, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsName *DnsNameTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsName *DnsNameSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsName.Contract.SetBaseUri(&_DnsName.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsName *DnsNameTransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsName.Contract.SetBaseUri(&_DnsName.TransactOpts, baseUri_)
}

// SetContract is a paid mutator transaction binding the contract method 0xac6dd5aa.
//
// Solidity: function setContract(address taxC_, address accountantC_, address priceC_, address sigUser_, bool open_) returns()
func (_DnsName *DnsNameTransactor) SetContract(opts *bind.TransactOpts, taxC_ common.Address, accountantC_ common.Address, priceC_ common.Address, sigUser_ common.Address, open_ bool) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "setContract", taxC_, accountantC_, priceC_, sigUser_, open_)
}

// SetContract is a paid mutator transaction binding the contract method 0xac6dd5aa.
//
// Solidity: function setContract(address taxC_, address accountantC_, address priceC_, address sigUser_, bool open_) returns()
func (_DnsName *DnsNameSession) SetContract(taxC_ common.Address, accountantC_ common.Address, priceC_ common.Address, sigUser_ common.Address, open_ bool) (*types.Transaction, error) {
	return _DnsName.Contract.SetContract(&_DnsName.TransactOpts, taxC_, accountantC_, priceC_, sigUser_, open_)
}

// SetContract is a paid mutator transaction binding the contract method 0xac6dd5aa.
//
// Solidity: function setContract(address taxC_, address accountantC_, address priceC_, address sigUser_, bool open_) returns()
func (_DnsName *DnsNameTransactorSession) SetContract(taxC_ common.Address, accountantC_ common.Address, priceC_ common.Address, sigUser_ common.Address, open_ bool) (*types.Transaction, error) {
	return _DnsName.Contract.SetContract(&_DnsName.TransactOpts, taxC_, accountantC_, priceC_, sigUser_, open_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsName *DnsNameTransactor) SetName(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "setName", name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsName *DnsNameSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsName.Contract.SetName(&_DnsName.TransactOpts, name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsName *DnsNameTransactorSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsName.Contract.SetName(&_DnsName.TransactOpts, name_, symbol_)
}

// SetRegistered is a paid mutator transaction binding the contract method 0x9f311652.
//
// Solidity: function setRegistered(bool open_, bytes32 hash_) returns()
func (_DnsName *DnsNameTransactor) SetRegistered(opts *bind.TransactOpts, open_ bool, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "setRegistered", open_, hash_)
}

// SetRegistered is a paid mutator transaction binding the contract method 0x9f311652.
//
// Solidity: function setRegistered(bool open_, bytes32 hash_) returns()
func (_DnsName *DnsNameSession) SetRegistered(open_ bool, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsName.Contract.SetRegistered(&_DnsName.TransactOpts, open_, hash_)
}

// SetRegistered is a paid mutator transaction binding the contract method 0x9f311652.
//
// Solidity: function setRegistered(bool open_, bytes32 hash_) returns()
func (_DnsName *DnsNameTransactorSession) SetRegistered(open_ bool, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsName.Contract.SetRegistered(&_DnsName.TransactOpts, open_, hash_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.TransferFrom(&_DnsName.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsName *DnsNameTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsName.Contract.TransferFrom(&_DnsName.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_DnsName *DnsNameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsName.contract.Transact(opts, "transferOwnership", newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_DnsName *DnsNameSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsName.Contract.TransferOwnership(&_DnsName.TransactOpts, newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_DnsName *DnsNameTransactorSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsName.Contract.TransferOwnership(&_DnsName.TransactOpts, newOwner_)
}

// DnsNameApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DnsName contract.
type DnsNameApprovalIterator struct {
	Event *DnsNameApproval // Event containing the contract specifics and raw log

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
func (it *DnsNameApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameApproval)
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
		it.Event = new(DnsNameApproval)
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
func (it *DnsNameApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameApproval represents a Approval event raised by the DnsName contract.
type DnsNameApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsName *DnsNameFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DnsNameApprovalIterator, error) {

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

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameApprovalIterator{contract: _DnsName.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsName *DnsNameFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DnsNameApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameApproval)
				if err := _DnsName.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DnsName *DnsNameFilterer) ParseApproval(log types.Log) (*DnsNameApproval, error) {
	event := new(DnsNameApproval)
	if err := _DnsName.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DnsName contract.
type DnsNameApprovalForAllIterator struct {
	Event *DnsNameApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DnsNameApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameApprovalForAll)
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
		it.Event = new(DnsNameApprovalForAll)
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
func (it *DnsNameApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameApprovalForAll represents a ApprovalForAll event raised by the DnsName contract.
type DnsNameApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsName *DnsNameFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DnsNameApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameApprovalForAllIterator{contract: _DnsName.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsName *DnsNameFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DnsNameApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameApprovalForAll)
				if err := _DnsName.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DnsName *DnsNameFilterer) ParseApprovalForAll(log types.Log) (*DnsNameApprovalForAll, error) {
	event := new(DnsNameApprovalForAll)
	if err := _DnsName.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameEvChargeDnsNameIterator is returned from FilterEvChargeDnsName and is used to iterate over the raw logs and unpacked data for EvChargeDnsName events raised by the DnsName contract.
type DnsNameEvChargeDnsNameIterator struct {
	Event *DnsNameEvChargeDnsName // Event containing the contract specifics and raw log

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
func (it *DnsNameEvChargeDnsNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameEvChargeDnsName)
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
		it.Event = new(DnsNameEvChargeDnsName)
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
func (it *DnsNameEvChargeDnsNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameEvChargeDnsNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameEvChargeDnsName represents a EvChargeDnsName event raised by the DnsName contract.
type DnsNameEvChargeDnsName struct {
	Erc20Addr       common.Address
	NameHash        [32]byte
	Year            uint8
	IsTransferOwner bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEvChargeDnsName is a free log retrieval operation binding the contract event 0xa146427f2a4d6c1132cea8182b37e4dede940d131a8688d6654c8ca33f4ee200.
//
// Solidity: event EvChargeDnsName(address erc20Addr, bytes32 nameHash, uint8 year, bool isTransferOwner)
func (_DnsName *DnsNameFilterer) FilterEvChargeDnsName(opts *bind.FilterOpts) (*DnsNameEvChargeDnsNameIterator, error) {

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "EvChargeDnsName")
	if err != nil {
		return nil, err
	}
	return &DnsNameEvChargeDnsNameIterator{contract: _DnsName.contract, event: "EvChargeDnsName", logs: logs, sub: sub}, nil
}

// WatchEvChargeDnsName is a free log subscription operation binding the contract event 0xa146427f2a4d6c1132cea8182b37e4dede940d131a8688d6654c8ca33f4ee200.
//
// Solidity: event EvChargeDnsName(address erc20Addr, bytes32 nameHash, uint8 year, bool isTransferOwner)
func (_DnsName *DnsNameFilterer) WatchEvChargeDnsName(opts *bind.WatchOpts, sink chan<- *DnsNameEvChargeDnsName) (event.Subscription, error) {

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "EvChargeDnsName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameEvChargeDnsName)
				if err := _DnsName.contract.UnpackLog(event, "EvChargeDnsName", log); err != nil {
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

// ParseEvChargeDnsName is a log parse operation binding the contract event 0xa146427f2a4d6c1132cea8182b37e4dede940d131a8688d6654c8ca33f4ee200.
//
// Solidity: event EvChargeDnsName(address erc20Addr, bytes32 nameHash, uint8 year, bool isTransferOwner)
func (_DnsName *DnsNameFilterer) ParseEvChargeDnsName(log types.Log) (*DnsNameEvChargeDnsName, error) {
	event := new(DnsNameEvChargeDnsName)
	if err := _DnsName.contract.UnpackLog(event, "EvChargeDnsName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameEvMintDnsNameIterator is returned from FilterEvMintDnsName and is used to iterate over the raw logs and unpacked data for EvMintDnsName events raised by the DnsName contract.
type DnsNameEvMintDnsNameIterator struct {
	Event *DnsNameEvMintDnsName // Event containing the contract specifics and raw log

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
func (it *DnsNameEvMintDnsNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameEvMintDnsName)
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
		it.Event = new(DnsNameEvMintDnsName)
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
func (it *DnsNameEvMintDnsNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameEvMintDnsNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameEvMintDnsName represents a EvMintDnsName event raised by the DnsName contract.
type DnsNameEvMintDnsName struct {
	Erc20Addr  common.Address
	EntireName string
	Year       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintDnsName is a free log retrieval operation binding the contract event 0x4f2c78fc8de8637c4a58660c3a4d34121009ba566bf6caa5c6629de7e4ff43b4.
//
// Solidity: event EvMintDnsName(address erc20Addr, string entireName, uint8 year)
func (_DnsName *DnsNameFilterer) FilterEvMintDnsName(opts *bind.FilterOpts) (*DnsNameEvMintDnsNameIterator, error) {

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "EvMintDnsName")
	if err != nil {
		return nil, err
	}
	return &DnsNameEvMintDnsNameIterator{contract: _DnsName.contract, event: "EvMintDnsName", logs: logs, sub: sub}, nil
}

// WatchEvMintDnsName is a free log subscription operation binding the contract event 0x4f2c78fc8de8637c4a58660c3a4d34121009ba566bf6caa5c6629de7e4ff43b4.
//
// Solidity: event EvMintDnsName(address erc20Addr, string entireName, uint8 year)
func (_DnsName *DnsNameFilterer) WatchEvMintDnsName(opts *bind.WatchOpts, sink chan<- *DnsNameEvMintDnsName) (event.Subscription, error) {

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "EvMintDnsName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameEvMintDnsName)
				if err := _DnsName.contract.UnpackLog(event, "EvMintDnsName", log); err != nil {
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

// ParseEvMintDnsName is a log parse operation binding the contract event 0x4f2c78fc8de8637c4a58660c3a4d34121009ba566bf6caa5c6629de7e4ff43b4.
//
// Solidity: event EvMintDnsName(address erc20Addr, string entireName, uint8 year)
func (_DnsName *DnsNameFilterer) ParseEvMintDnsName(log types.Log) (*DnsNameEvMintDnsName, error) {
	event := new(DnsNameEvMintDnsName)
	if err := _DnsName.contract.UnpackLog(event, "EvMintDnsName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameEvNewSubNameIterator is returned from FilterEvNewSubName and is used to iterate over the raw logs and unpacked data for EvNewSubName events raised by the DnsName contract.
type DnsNameEvNewSubNameIterator struct {
	Event *DnsNameEvNewSubName // Event containing the contract specifics and raw log

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
func (it *DnsNameEvNewSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameEvNewSubName)
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
		it.Event = new(DnsNameEvNewSubName)
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
func (it *DnsNameEvNewSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameEvNewSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameEvNewSubName represents a EvNewSubName event raised by the DnsName contract.
type DnsNameEvNewSubName struct {
	Hash       [32]byte
	TokenId    *big.Int
	Erc721Addr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvNewSubName is a free log retrieval operation binding the contract event 0x2cbacbceb0827ebe6029c5ed759a7006a5cd8fe8b401d05cec8c63d23fc961be.
//
// Solidity: event EvNewSubName(bytes32 hash, uint256 tokenId, address erc721Addr)
func (_DnsName *DnsNameFilterer) FilterEvNewSubName(opts *bind.FilterOpts) (*DnsNameEvNewSubNameIterator, error) {

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "EvNewSubName")
	if err != nil {
		return nil, err
	}
	return &DnsNameEvNewSubNameIterator{contract: _DnsName.contract, event: "EvNewSubName", logs: logs, sub: sub}, nil
}

// WatchEvNewSubName is a free log subscription operation binding the contract event 0x2cbacbceb0827ebe6029c5ed759a7006a5cd8fe8b401d05cec8c63d23fc961be.
//
// Solidity: event EvNewSubName(bytes32 hash, uint256 tokenId, address erc721Addr)
func (_DnsName *DnsNameFilterer) WatchEvNewSubName(opts *bind.WatchOpts, sink chan<- *DnsNameEvNewSubName) (event.Subscription, error) {

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "EvNewSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameEvNewSubName)
				if err := _DnsName.contract.UnpackLog(event, "EvNewSubName", log); err != nil {
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

// ParseEvNewSubName is a log parse operation binding the contract event 0x2cbacbceb0827ebe6029c5ed759a7006a5cd8fe8b401d05cec8c63d23fc961be.
//
// Solidity: event EvNewSubName(bytes32 hash, uint256 tokenId, address erc721Addr)
func (_DnsName *DnsNameFilterer) ParseEvNewSubName(log types.Log) (*DnsNameEvNewSubName, error) {
	event := new(DnsNameEvNewSubName)
	if err := _DnsName.contract.UnpackLog(event, "EvNewSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameEvOpenToRegIterator is returned from FilterEvOpenToReg and is used to iterate over the raw logs and unpacked data for EvOpenToReg events raised by the DnsName contract.
type DnsNameEvOpenToRegIterator struct {
	Event *DnsNameEvOpenToReg // Event containing the contract specifics and raw log

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
func (it *DnsNameEvOpenToRegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameEvOpenToReg)
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
		it.Event = new(DnsNameEvOpenToReg)
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
func (it *DnsNameEvOpenToRegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameEvOpenToRegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameEvOpenToReg represents a EvOpenToReg event raised by the DnsName contract.
type DnsNameEvOpenToReg struct {
	Open bool
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvOpenToReg is a free log retrieval operation binding the contract event 0x5d17e42084f2cbfe9035a150ff9e06538a04e326946f6c2f5a74174b30316797.
//
// Solidity: event EvOpenToReg(bool open, bytes32 hash)
func (_DnsName *DnsNameFilterer) FilterEvOpenToReg(opts *bind.FilterOpts) (*DnsNameEvOpenToRegIterator, error) {

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "EvOpenToReg")
	if err != nil {
		return nil, err
	}
	return &DnsNameEvOpenToRegIterator{contract: _DnsName.contract, event: "EvOpenToReg", logs: logs, sub: sub}, nil
}

// WatchEvOpenToReg is a free log subscription operation binding the contract event 0x5d17e42084f2cbfe9035a150ff9e06538a04e326946f6c2f5a74174b30316797.
//
// Solidity: event EvOpenToReg(bool open, bytes32 hash)
func (_DnsName *DnsNameFilterer) WatchEvOpenToReg(opts *bind.WatchOpts, sink chan<- *DnsNameEvOpenToReg) (event.Subscription, error) {

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "EvOpenToReg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameEvOpenToReg)
				if err := _DnsName.contract.UnpackLog(event, "EvOpenToReg", log); err != nil {
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

// ParseEvOpenToReg is a log parse operation binding the contract event 0x5d17e42084f2cbfe9035a150ff9e06538a04e326946f6c2f5a74174b30316797.
//
// Solidity: event EvOpenToReg(bool open, bytes32 hash)
func (_DnsName *DnsNameFilterer) ParseEvOpenToReg(log types.Log) (*DnsNameEvOpenToReg, error) {
	event := new(DnsNameEvOpenToReg)
	if err := _DnsName.contract.UnpackLog(event, "EvOpenToReg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DnsName contract.
type DnsNameTransferIterator struct {
	Event *DnsNameTransfer // Event containing the contract specifics and raw log

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
func (it *DnsNameTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameTransfer)
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
		it.Event = new(DnsNameTransfer)
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
func (it *DnsNameTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameTransfer represents a Transfer event raised by the DnsName contract.
type DnsNameTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsName *DnsNameFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DnsNameTransferIterator, error) {

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

	logs, sub, err := _DnsName.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameTransferIterator{contract: _DnsName.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsName *DnsNameFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DnsNameTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DnsName.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameTransfer)
				if err := _DnsName.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DnsName *DnsNameFilterer) ParseTransfer(log types.Log) (*DnsNameTransfer, error) {
	event := new(DnsNameTransfer)
	if err := _DnsName.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
