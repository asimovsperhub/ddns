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

// DnsSecondLevelNameMetaData contains all meta data concerning the DnsSecondLevelName contract.
var DnsSecondLevelNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cost_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"acct_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsTop_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireSubName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"level2FatherHash\",\"type\":\"bytes32\"}],\"name\":\"EvAddSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeSecondLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeSecondLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"level2FatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"topHash\",\"type\":\"bytes32\"}],\"name\":\"EvDelSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintSecondLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintSecondLevelNameBySig\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireSubName_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"level2FatherHash_\",\"type\":\"bytes32\"}],\"name\":\"AddSubName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeSecondLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeSecondLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"level2FatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topHash_\",\"type\":\"bytes32\"}],\"name\":\"DelSubName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"MintSecondLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintSecondLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"contractIDnsAccountant\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costContractAddr\",\"outputs\":[{\"internalType\":\"contractIcost\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsTop\",\"outputs\":[{\"internalType\":\"contractIDnsTopLevelName\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintSwitch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cost_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"acct_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsTop_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mintSw_\",\"type\":\"uint8\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subNameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DnsSecondLevelNameABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsSecondLevelNameMetaData.ABI instead.
var DnsSecondLevelNameABI = DnsSecondLevelNameMetaData.ABI

// DnsSecondLevelName is an auto generated Go binding around an Ethereum contract.
type DnsSecondLevelName struct {
	DnsSecondLevelNameCaller     // Read-only binding to the contract
	DnsSecondLevelNameTransactor // Write-only binding to the contract
	DnsSecondLevelNameFilterer   // Log filterer for contract events
}

// DnsSecondLevelNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsSecondLevelNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsSecondLevelNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsSecondLevelNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsSecondLevelNameSession struct {
	Contract     *DnsSecondLevelName // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DnsSecondLevelNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsSecondLevelNameCallerSession struct {
	Contract *DnsSecondLevelNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DnsSecondLevelNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsSecondLevelNameTransactorSession struct {
	Contract     *DnsSecondLevelNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DnsSecondLevelNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsSecondLevelNameRaw struct {
	Contract *DnsSecondLevelName // Generic contract binding to access the raw methods on
}

// DnsSecondLevelNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsSecondLevelNameCallerRaw struct {
	Contract *DnsSecondLevelNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsSecondLevelNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsSecondLevelNameTransactorRaw struct {
	Contract *DnsSecondLevelNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsSecondLevelName creates a new instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelName(address common.Address, backend bind.ContractBackend) (*DnsSecondLevelName, error) {
	contract, err := bindDnsSecondLevelName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelName{DnsSecondLevelNameCaller: DnsSecondLevelNameCaller{contract: contract}, DnsSecondLevelNameTransactor: DnsSecondLevelNameTransactor{contract: contract}, DnsSecondLevelNameFilterer: DnsSecondLevelNameFilterer{contract: contract}}, nil
}

// NewDnsSecondLevelNameCaller creates a new read-only instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameCaller(address common.Address, caller bind.ContractCaller) (*DnsSecondLevelNameCaller, error) {
	contract, err := bindDnsSecondLevelName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameCaller{contract: contract}, nil
}

// NewDnsSecondLevelNameTransactor creates a new write-only instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsSecondLevelNameTransactor, error) {
	contract, err := bindDnsSecondLevelName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameTransactor{contract: contract}, nil
}

// NewDnsSecondLevelNameFilterer creates a new log filterer instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsSecondLevelNameFilterer, error) {
	contract, err := bindDnsSecondLevelName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameFilterer{contract: contract}, nil
}

// bindDnsSecondLevelName binds a generic wrapper to an already deployed contract.
func bindDnsSecondLevelName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsSecondLevelNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSecondLevelName *DnsSecondLevelNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSecondLevelName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) AccountantC() (common.Address, error) {
	return _DnsSecondLevelName.Contract.AccountantC(&_DnsSecondLevelName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsSecondLevelName.Contract.AccountantC(&_DnsSecondLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) CostContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "costContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) CostContractAddr() (common.Address, error) {
	return _DnsSecondLevelName.Contract.CostContractAddr(&_DnsSecondLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) CostContractAddr() (common.Address, error) {
	return _DnsSecondLevelName.Contract.CostContractAddr(&_DnsSecondLevelName.CallOpts)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) DnsTop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "dnsTop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) DnsTop() (common.Address, error) {
	return _DnsSecondLevelName.Contract.DnsTop(&_DnsSecondLevelName.CallOpts)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) DnsTop() (common.Address, error) {
	return _DnsSecondLevelName.Contract.DnsTop(&_DnsSecondLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) MintSwitch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "mintSwitch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSwitch() (uint8, error) {
	return _DnsSecondLevelName.Contract.MintSwitch(&_DnsSecondLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) MintSwitch() (uint8, error) {
	return _DnsSecondLevelName.Contract.MintSwitch(&_DnsSecondLevelName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "nameStore", arg0, arg1)

	outstruct := new(struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EntireName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ExpireTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) NameStore(arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsSecondLevelName.Contract.NameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) NameStore(arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsSecondLevelName.Contract.NameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) SubNameStore(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (string, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "subNameStore", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) SubNameStore(arg0 [32]byte, arg1 [32]byte) (string, error) {
	return _DnsSecondLevelName.Contract.SubNameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) SubNameStore(arg0 [32]byte, arg1 [32]byte) (string, error) {
	return _DnsSecondLevelName.Contract.SubNameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) AddSubName(opts *bind.TransactOpts, entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "AddSubName", entireSubName_, level2FatherHash_)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) AddSubName(entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.AddSubName(&_DnsSecondLevelName.TransactOpts, entireSubName_, level2FatherHash_)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) AddSubName(entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.AddSubName(&_DnsSecondLevelName.TransactOpts, entireSubName_, level2FatherHash_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) ChargeSecondLevelName(opts *bind.TransactOpts, fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "ChargeSecondLevelName", fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) ChargeSecondLevelName(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelName(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) ChargeSecondLevelName(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelName(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) ChargeSecondLevelNameBySig(opts *bind.TransactOpts, fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "ChargeSecondLevelNameBySig", fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) ChargeSecondLevelNameBySig(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) ChargeSecondLevelNameBySig(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) DelSubName(opts *bind.TransactOpts, nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "DelSubName", nameHash_, level2FatherHash_, topHash_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) DelSubName(nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DelSubName(&_DnsSecondLevelName.TransactOpts, nameHash_, level2FatherHash_, topHash_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) DelSubName(nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DelSubName(&_DnsSecondLevelName.TransactOpts, nameHash_, level2FatherHash_, topHash_)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) MintSecondLevelName(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "MintSecondLevelName", entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSecondLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelName(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) MintSecondLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelName(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) MintSecondLevelNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "MintSecondLevelNameBySig", entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSecondLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) MintSecondLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) SetContract(opts *bind.TransactOpts, cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "setContract", cost_, acct_, dnsTop_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) SetContract(cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.SetContract(&_DnsSecondLevelName.TransactOpts, cost_, acct_, dnsTop_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) SetContract(cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.SetContract(&_DnsSecondLevelName.TransactOpts, cost_, acct_, dnsTop_, mintSw_)
}

// DnsSecondLevelNameEvAddSubNameIterator is returned from FilterEvAddSubName and is used to iterate over the raw logs and unpacked data for EvAddSubName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvAddSubNameIterator struct {
	Event *DnsSecondLevelNameEvAddSubName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvAddSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvAddSubName)
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
		it.Event = new(DnsSecondLevelNameEvAddSubName)
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
func (it *DnsSecondLevelNameEvAddSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvAddSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvAddSubName represents a EvAddSubName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvAddSubName struct {
	EntireSubName    string
	Level2FatherHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEvAddSubName is a free log retrieval operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvAddSubName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvAddSubNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvAddSubNameIterator{contract: _DnsSecondLevelName.contract, event: "EvAddSubName", logs: logs, sub: sub}, nil
}

// WatchEvAddSubName is a free log subscription operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvAddSubName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvAddSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvAddSubName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
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

// ParseEvAddSubName is a log parse operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvAddSubName(log types.Log) (*DnsSecondLevelNameEvAddSubName, error) {
	event := new(DnsSecondLevelNameEvAddSubName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameIterator is returned from FilterEvChargeSecondLevelName and is used to iterate over the raw logs and unpacked data for EvChargeSecondLevelName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameIterator struct {
	Event *DnsSecondLevelNameEvChargeSecondLevelName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvChargeSecondLevelName)
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
		it.Event = new(DnsSecondLevelNameEvChargeSecondLevelName)
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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvChargeSecondLevelName represents a EvChargeSecondLevelName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelName struct {
	FatherHash [32]byte
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeSecondLevelName is a free log retrieval operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvChargeSecondLevelName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvChargeSecondLevelNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvChargeSecondLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvChargeSecondLevelNameIterator{contract: _DnsSecondLevelName.contract, event: "EvChargeSecondLevelName", logs: logs, sub: sub}, nil
}

// WatchEvChargeSecondLevelName is a free log subscription operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvChargeSecondLevelName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvChargeSecondLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvChargeSecondLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvChargeSecondLevelName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelName", log); err != nil {
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

// ParseEvChargeSecondLevelName is a log parse operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvChargeSecondLevelName(log types.Log) (*DnsSecondLevelNameEvChargeSecondLevelName, error) {
	event := new(DnsSecondLevelNameEvChargeSecondLevelName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator is returned from FilterEvChargeSecondLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvChargeSecondLevelNameBySig events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator struct {
	Event *DnsSecondLevelNameEvChargeSecondLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
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
		it.Event = new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameBySig represents a EvChargeSecondLevelNameBySig event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameBySig struct {
	FatherHash [32]byte
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeSecondLevelNameBySig is a free log retrieval operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvChargeSecondLevelNameBySig(opts *bind.FilterOpts) (*DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvChargeSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator{contract: _DnsSecondLevelName.contract, event: "EvChargeSecondLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvChargeSecondLevelNameBySig is a free log subscription operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvChargeSecondLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvChargeSecondLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvChargeSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelNameBySig", log); err != nil {
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

// ParseEvChargeSecondLevelNameBySig is a log parse operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvChargeSecondLevelNameBySig(log types.Log) (*DnsSecondLevelNameEvChargeSecondLevelNameBySig, error) {
	event := new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvDelSubNameIterator is returned from FilterEvDelSubName and is used to iterate over the raw logs and unpacked data for EvDelSubName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvDelSubNameIterator struct {
	Event *DnsSecondLevelNameEvDelSubName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvDelSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvDelSubName)
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
		it.Event = new(DnsSecondLevelNameEvDelSubName)
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
func (it *DnsSecondLevelNameEvDelSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvDelSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvDelSubName represents a EvDelSubName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvDelSubName struct {
	NameHash         [32]byte
	Level2FatherHash [32]byte
	TopHash          [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEvDelSubName is a free log retrieval operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvDelSubName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvDelSubNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvDelSubNameIterator{contract: _DnsSecondLevelName.contract, event: "EvDelSubName", logs: logs, sub: sub}, nil
}

// WatchEvDelSubName is a free log subscription operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvDelSubName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvDelSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvDelSubName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
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

// ParseEvDelSubName is a log parse operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvDelSubName(log types.Log) (*DnsSecondLevelNameEvDelSubName, error) {
	event := new(DnsSecondLevelNameEvDelSubName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvMintSecondLevelNameIterator is returned from FilterEvMintSecondLevelName and is used to iterate over the raw logs and unpacked data for EvMintSecondLevelName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameIterator struct {
	Event *DnsSecondLevelNameEvMintSecondLevelName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvMintSecondLevelName)
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
		it.Event = new(DnsSecondLevelNameEvMintSecondLevelName)
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
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvMintSecondLevelName represents a EvMintSecondLevelName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelName struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintSecondLevelName is a free log retrieval operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvMintSecondLevelName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvMintSecondLevelNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvMintSecondLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvMintSecondLevelNameIterator{contract: _DnsSecondLevelName.contract, event: "EvMintSecondLevelName", logs: logs, sub: sub}, nil
}

// WatchEvMintSecondLevelName is a free log subscription operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvMintSecondLevelName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvMintSecondLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvMintSecondLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvMintSecondLevelName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelName", log); err != nil {
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

// ParseEvMintSecondLevelName is a log parse operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvMintSecondLevelName(log types.Log) (*DnsSecondLevelNameEvMintSecondLevelName, error) {
	event := new(DnsSecondLevelNameEvMintSecondLevelName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvMintSecondLevelNameBySigIterator is returned from FilterEvMintSecondLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvMintSecondLevelNameBySig events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameBySigIterator struct {
	Event *DnsSecondLevelNameEvMintSecondLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
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
		it.Event = new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
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
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvMintSecondLevelNameBySig represents a EvMintSecondLevelNameBySig event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameBySig struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintSecondLevelNameBySig is a free log retrieval operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvMintSecondLevelNameBySig(opts *bind.FilterOpts) (*DnsSecondLevelNameEvMintSecondLevelNameBySigIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvMintSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvMintSecondLevelNameBySigIterator{contract: _DnsSecondLevelName.contract, event: "EvMintSecondLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvMintSecondLevelNameBySig is a free log subscription operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvMintSecondLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvMintSecondLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvMintSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelNameBySig", log); err != nil {
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

// ParseEvMintSecondLevelNameBySig is a log parse operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvMintSecondLevelNameBySig(log types.Log) (*DnsSecondLevelNameEvMintSecondLevelNameBySig, error) {
	event := new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
