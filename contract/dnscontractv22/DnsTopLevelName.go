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

// DnsTopLevelNameMetaData contains all meta data concerning the DnsTopLevelName contract.
var DnsTopLevelNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"EVChargeTopLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeTopLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintTopLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintTopLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"EvOpen2Reg\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeTopLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeTopLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"MintTopLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintTopLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"}],\"name\":\"Open2Reg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"contractIDnsAccountant\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costContractAddr\",\"outputs\":[{\"internalType\":\"contractIcost\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsSecond\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"erc721Store\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"}],\"name\":\"getErc721Contract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintSwitch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountantAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsSecond_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mintSw_\",\"type\":\"uint8\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DnsTopLevelNameABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsTopLevelNameMetaData.ABI instead.
var DnsTopLevelNameABI = DnsTopLevelNameMetaData.ABI

// DnsTopLevelName is an auto generated Go binding around an Ethereum contract.
type DnsTopLevelName struct {
	DnsTopLevelNameCaller     // Read-only binding to the contract
	DnsTopLevelNameTransactor // Write-only binding to the contract
	DnsTopLevelNameFilterer   // Log filterer for contract events
}

// DnsTopLevelNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsTopLevelNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsTopLevelNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsTopLevelNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsTopLevelNameSession struct {
	Contract     *DnsTopLevelName  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsTopLevelNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsTopLevelNameCallerSession struct {
	Contract *DnsTopLevelNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DnsTopLevelNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsTopLevelNameTransactorSession struct {
	Contract     *DnsTopLevelNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DnsTopLevelNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsTopLevelNameRaw struct {
	Contract *DnsTopLevelName // Generic contract binding to access the raw methods on
}

// DnsTopLevelNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsTopLevelNameCallerRaw struct {
	Contract *DnsTopLevelNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsTopLevelNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsTopLevelNameTransactorRaw struct {
	Contract *DnsTopLevelNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsTopLevelName creates a new instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelName(address common.Address, backend bind.ContractBackend) (*DnsTopLevelName, error) {
	contract, err := bindDnsTopLevelName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelName{DnsTopLevelNameCaller: DnsTopLevelNameCaller{contract: contract}, DnsTopLevelNameTransactor: DnsTopLevelNameTransactor{contract: contract}, DnsTopLevelNameFilterer: DnsTopLevelNameFilterer{contract: contract}}, nil
}

// NewDnsTopLevelNameCaller creates a new read-only instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameCaller(address common.Address, caller bind.ContractCaller) (*DnsTopLevelNameCaller, error) {
	contract, err := bindDnsTopLevelName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameCaller{contract: contract}, nil
}

// NewDnsTopLevelNameTransactor creates a new write-only instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsTopLevelNameTransactor, error) {
	contract, err := bindDnsTopLevelName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameTransactor{contract: contract}, nil
}

// NewDnsTopLevelNameFilterer creates a new log filterer instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsTopLevelNameFilterer, error) {
	contract, err := bindDnsTopLevelName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameFilterer{contract: contract}, nil
}

// bindDnsTopLevelName binds a generic wrapper to an already deployed contract.
func bindDnsTopLevelName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsTopLevelNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsTopLevelName.Contract.DnsTopLevelNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.DnsTopLevelNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.DnsTopLevelNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsTopLevelName *DnsTopLevelNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsTopLevelName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsTopLevelName *DnsTopLevelNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsTopLevelName *DnsTopLevelNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) AccountantC() (common.Address, error) {
	return _DnsTopLevelName.Contract.AccountantC(&_DnsTopLevelName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsTopLevelName.Contract.AccountantC(&_DnsTopLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) CostContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "costContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) CostContractAddr() (common.Address, error) {
	return _DnsTopLevelName.Contract.CostContractAddr(&_DnsTopLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) CostContractAddr() (common.Address, error) {
	return _DnsTopLevelName.Contract.CostContractAddr(&_DnsTopLevelName.CallOpts)
}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) DnsSecond(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "dnsSecond")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) DnsSecond() (common.Address, error) {
	return _DnsTopLevelName.Contract.DnsSecond(&_DnsTopLevelName.CallOpts)
}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) DnsSecond() (common.Address, error) {
	return _DnsTopLevelName.Contract.DnsSecond(&_DnsTopLevelName.CallOpts)
}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) Erc721Store(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "erc721Store", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) Erc721Store(arg0 [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.Erc721Store(&_DnsTopLevelName.CallOpts, arg0)
}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) Erc721Store(arg0 [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.Erc721Store(&_DnsTopLevelName.CallOpts, arg0)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) GetErc721Contract(opts *bind.CallOpts, fatherHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "getErc721Contract", fatherHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.GetErc721Contract(&_DnsTopLevelName.CallOpts, fatherHash)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.GetErc721Contract(&_DnsTopLevelName.CallOpts, fatherHash)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) GetOperator() (common.Address, error) {
	return _DnsTopLevelName.Contract.GetOperator(&_DnsTopLevelName.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) GetOperator() (common.Address, error) {
	return _DnsTopLevelName.Contract.GetOperator(&_DnsTopLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameCaller) MintSwitch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "mintSwitch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameSession) MintSwitch() (uint8, error) {
	return _DnsTopLevelName.Contract.MintSwitch(&_DnsTopLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) MintSwitch() (uint8, error) {
	return _DnsTopLevelName.Contract.MintSwitch(&_DnsTopLevelName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "nameStore", arg0)

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

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameSession) NameStore(arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsTopLevelName.Contract.NameStore(&_DnsTopLevelName.CallOpts, arg0)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) NameStore(arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsTopLevelName.Contract.NameStore(&_DnsTopLevelName.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) Owner() (common.Address, error) {
	return _DnsTopLevelName.Contract.Owner(&_DnsTopLevelName.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) Owner() (common.Address, error) {
	return _DnsTopLevelName.Contract.Owner(&_DnsTopLevelName.CallOpts)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) ChargeTopLevelName(opts *bind.TransactOpts, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "ChargeTopLevelName", nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) ChargeTopLevelName(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelName(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) ChargeTopLevelName(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelName(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) ChargeTopLevelNameBySig(opts *bind.TransactOpts, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "ChargeTopLevelNameBySig", nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) ChargeTopLevelNameBySig(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) ChargeTopLevelNameBySig(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) MintTopLevelName(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "MintTopLevelName", entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) MintTopLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelName(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) MintTopLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelName(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) MintTopLevelNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "MintTopLevelNameBySig", entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) MintTopLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) MintTopLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) Open2Reg(opts *bind.TransactOpts, nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "Open2Reg", nameHash_)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) Open2Reg(nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.Open2Reg(&_DnsTopLevelName.TransactOpts, nameHash_)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) Open2Reg(nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.Open2Reg(&_DnsTopLevelName.TransactOpts, nameHash_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) SetContract(opts *bind.TransactOpts, costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "setContract", costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) SetContract(costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.SetContract(&_DnsTopLevelName.TransactOpts, costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) SetContract(costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.SetContract(&_DnsTopLevelName.TransactOpts, costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.TransferOwnership(&_DnsTopLevelName.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.TransferOwnership(&_DnsTopLevelName.TransactOpts, newOwner)
}

// DnsTopLevelNameEVChargeTopLevelNameIterator is returned from FilterEVChargeTopLevelName and is used to iterate over the raw logs and unpacked data for EVChargeTopLevelName events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEVChargeTopLevelNameIterator struct {
	Event *DnsTopLevelNameEVChargeTopLevelName // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEVChargeTopLevelName)
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
		it.Event = new(DnsTopLevelNameEVChargeTopLevelName)
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
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEVChargeTopLevelName represents a EVChargeTopLevelName event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEVChargeTopLevelName struct {
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEVChargeTopLevelName is a free log retrieval operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEVChargeTopLevelName(opts *bind.FilterOpts) (*DnsTopLevelNameEVChargeTopLevelNameIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EVChargeTopLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEVChargeTopLevelNameIterator{contract: _DnsTopLevelName.contract, event: "EVChargeTopLevelName", logs: logs, sub: sub}, nil
}

// WatchEVChargeTopLevelName is a free log subscription operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEVChargeTopLevelName(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEVChargeTopLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EVChargeTopLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEVChargeTopLevelName)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EVChargeTopLevelName", log); err != nil {
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

// ParseEVChargeTopLevelName is a log parse operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEVChargeTopLevelName(log types.Log) (*DnsTopLevelNameEVChargeTopLevelName, error) {
	event := new(DnsTopLevelNameEVChargeTopLevelName)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EVChargeTopLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvChargeTopLevelNameBySigIterator is returned from FilterEvChargeTopLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvChargeTopLevelNameBySig events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvChargeTopLevelNameBySigIterator struct {
	Event *DnsTopLevelNameEvChargeTopLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvChargeTopLevelNameBySig)
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
		it.Event = new(DnsTopLevelNameEvChargeTopLevelNameBySig)
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
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvChargeTopLevelNameBySig represents a EvChargeTopLevelNameBySig event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvChargeTopLevelNameBySig struct {
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeTopLevelNameBySig is a free log retrieval operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvChargeTopLevelNameBySig(opts *bind.FilterOpts) (*DnsTopLevelNameEvChargeTopLevelNameBySigIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvChargeTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvChargeTopLevelNameBySigIterator{contract: _DnsTopLevelName.contract, event: "EvChargeTopLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvChargeTopLevelNameBySig is a free log subscription operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvChargeTopLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvChargeTopLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvChargeTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvChargeTopLevelNameBySig)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvChargeTopLevelNameBySig", log); err != nil {
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

// ParseEvChargeTopLevelNameBySig is a log parse operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvChargeTopLevelNameBySig(log types.Log) (*DnsTopLevelNameEvChargeTopLevelNameBySig, error) {
	event := new(DnsTopLevelNameEvChargeTopLevelNameBySig)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvChargeTopLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvMintTopLevelNameIterator is returned from FilterEvMintTopLevelName and is used to iterate over the raw logs and unpacked data for EvMintTopLevelName events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameIterator struct {
	Event *DnsTopLevelNameEvMintTopLevelName // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvMintTopLevelName)
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
		it.Event = new(DnsTopLevelNameEvMintTopLevelName)
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
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvMintTopLevelName represents a EvMintTopLevelName event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelName struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintTopLevelName is a free log retrieval operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvMintTopLevelName(opts *bind.FilterOpts) (*DnsTopLevelNameEvMintTopLevelNameIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvMintTopLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvMintTopLevelNameIterator{contract: _DnsTopLevelName.contract, event: "EvMintTopLevelName", logs: logs, sub: sub}, nil
}

// WatchEvMintTopLevelName is a free log subscription operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvMintTopLevelName(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvMintTopLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvMintTopLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvMintTopLevelName)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelName", log); err != nil {
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

// ParseEvMintTopLevelName is a log parse operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvMintTopLevelName(log types.Log) (*DnsTopLevelNameEvMintTopLevelName, error) {
	event := new(DnsTopLevelNameEvMintTopLevelName)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvMintTopLevelNameBySigIterator is returned from FilterEvMintTopLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvMintTopLevelNameBySig events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameBySigIterator struct {
	Event *DnsTopLevelNameEvMintTopLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvMintTopLevelNameBySig)
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
		it.Event = new(DnsTopLevelNameEvMintTopLevelNameBySig)
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
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvMintTopLevelNameBySig represents a EvMintTopLevelNameBySig event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameBySig struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintTopLevelNameBySig is a free log retrieval operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvMintTopLevelNameBySig(opts *bind.FilterOpts) (*DnsTopLevelNameEvMintTopLevelNameBySigIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvMintTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvMintTopLevelNameBySigIterator{contract: _DnsTopLevelName.contract, event: "EvMintTopLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvMintTopLevelNameBySig is a free log subscription operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvMintTopLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvMintTopLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvMintTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvMintTopLevelNameBySig)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelNameBySig", log); err != nil {
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

// ParseEvMintTopLevelNameBySig is a log parse operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvMintTopLevelNameBySig(log types.Log) (*DnsTopLevelNameEvMintTopLevelNameBySig, error) {
	event := new(DnsTopLevelNameEvMintTopLevelNameBySig)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvOpen2RegIterator is returned from FilterEvOpen2Reg and is used to iterate over the raw logs and unpacked data for EvOpen2Reg events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvOpen2RegIterator struct {
	Event *DnsTopLevelNameEvOpen2Reg // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvOpen2RegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvOpen2Reg)
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
		it.Event = new(DnsTopLevelNameEvOpen2Reg)
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
func (it *DnsTopLevelNameEvOpen2RegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvOpen2RegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvOpen2Reg represents a EvOpen2Reg event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvOpen2Reg struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvOpen2Reg is a free log retrieval operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvOpen2Reg(opts *bind.FilterOpts) (*DnsTopLevelNameEvOpen2RegIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvOpen2Reg")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvOpen2RegIterator{contract: _DnsTopLevelName.contract, event: "EvOpen2Reg", logs: logs, sub: sub}, nil
}

// WatchEvOpen2Reg is a free log subscription operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvOpen2Reg(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvOpen2Reg) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvOpen2Reg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvOpen2Reg)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvOpen2Reg", log); err != nil {
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

// ParseEvOpen2Reg is a log parse operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvOpen2Reg(log types.Log) (*DnsTopLevelNameEvOpen2Reg, error) {
	event := new(DnsTopLevelNameEvOpen2Reg)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvOpen2Reg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
