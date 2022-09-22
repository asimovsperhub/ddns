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

// DnsCostMetaData contains all meta data concerning the DnsCost contract.
var DnsCostMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"top_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"tax_\",\"type\":\"uint256[8]\"}],\"name\":\"SetSecondLevelNameTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"price_\",\"type\":\"uint256[8]\"}],\"name\":\"SetTopLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chainLinkContact_\",\"type\":\"address\"}],\"name\":\"addChainLinkAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"}],\"name\":\"addPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseDecimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chainLinkRateAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"delPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPayments\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"}],\"name\":\"getAllSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSecondLevelNameTax\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getExchangeWithId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getLatestExchangeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"paymentIsExist\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"priceIdIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dnstop_\",\"type\":\"address\"}],\"name\":\"setRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"prices_\",\"type\":\"uint256[8]\"}],\"name\":\"setSecondLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdtTokenAddr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimal_\",\"type\":\"uint8\"}],\"name\":\"setUsdtAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DnsCostABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsCostMetaData.ABI instead.
var DnsCostABI = DnsCostMetaData.ABI

// DnsCost is an auto generated Go binding around an Ethereum contract.
type DnsCost struct {
	DnsCostCaller     // Read-only binding to the contract
	DnsCostTransactor // Write-only binding to the contract
	DnsCostFilterer   // Log filterer for contract events
}

// DnsCostCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsCostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsCostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsCostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsCostSession struct {
	Contract     *DnsCost          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsCostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsCostCallerSession struct {
	Contract *DnsCostCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DnsCostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsCostTransactorSession struct {
	Contract     *DnsCostTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DnsCostRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsCostRaw struct {
	Contract *DnsCost // Generic contract binding to access the raw methods on
}

// DnsCostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsCostCallerRaw struct {
	Contract *DnsCostCaller // Generic read-only contract binding to access the raw methods on
}

// DnsCostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsCostTransactorRaw struct {
	Contract *DnsCostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsCost creates a new instance of DnsCost, bound to a specific deployed contract.
func NewDnsCost(address common.Address, backend bind.ContractBackend) (*DnsCost, error) {
	contract, err := bindDnsCost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsCost{DnsCostCaller: DnsCostCaller{contract: contract}, DnsCostTransactor: DnsCostTransactor{contract: contract}, DnsCostFilterer: DnsCostFilterer{contract: contract}}, nil
}

// NewDnsCostCaller creates a new read-only instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostCaller(address common.Address, caller bind.ContractCaller) (*DnsCostCaller, error) {
	contract, err := bindDnsCost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsCostCaller{contract: contract}, nil
}

// NewDnsCostTransactor creates a new write-only instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsCostTransactor, error) {
	contract, err := bindDnsCost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsCostTransactor{contract: contract}, nil
}

// NewDnsCostFilterer creates a new log filterer instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsCostFilterer, error) {
	contract, err := bindDnsCost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsCostFilterer{contract: contract}, nil
}

// bindDnsCost binds a generic wrapper to an already deployed contract.
func bindDnsCost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsCostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsCost *DnsCostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsCost.Contract.DnsCostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsCost *DnsCostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsCost.Contract.DnsCostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsCost *DnsCostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsCost.Contract.DnsCostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsCost *DnsCostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsCost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsCost *DnsCostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsCost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsCost *DnsCostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsCost.Contract.contract.Transact(opts, method, params...)
}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) SLPrice(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "SLPrice", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) SLPrice(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLPrice(&_DnsCost.CallOpts, arg0, arg1)
}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) SLPrice(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLPrice(&_DnsCost.CallOpts, arg0, arg1)
}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) SLTax(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "SLTax", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) SLTax(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLTax(&_DnsCost.CallOpts, arg0)
}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) SLTax(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLTax(&_DnsCost.CallOpts, arg0)
}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) TLPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "TLPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) TLPrice(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.TLPrice(&_DnsCost.CallOpts, arg0)
}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) TLPrice(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.TLPrice(&_DnsCost.CallOpts, arg0)
}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostCaller) BaseDecimal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "baseDecimal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostSession) BaseDecimal() (uint8, error) {
	return _DnsCost.Contract.BaseDecimal(&_DnsCost.CallOpts)
}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostCallerSession) BaseDecimal() (uint8, error) {
	return _DnsCost.Contract.BaseDecimal(&_DnsCost.CallOpts)
}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostCaller) ChainLinkRateAddr(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "chainLinkRateAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostSession) ChainLinkRateAddr(arg0 common.Address) (common.Address, error) {
	return _DnsCost.Contract.ChainLinkRateAddr(&_DnsCost.CallOpts, arg0)
}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostCallerSession) ChainLinkRateAddr(arg0 common.Address) (common.Address, error) {
	return _DnsCost.Contract.ChainLinkRateAddr(&_DnsCost.CallOpts, arg0)
}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostCaller) GetAllPayments(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllPayments")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostSession) GetAllPayments() ([]byte, error) {
	return _DnsCost.Contract.GetAllPayments(&_DnsCost.CallOpts)
}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostCallerSession) GetAllPayments() ([]byte, error) {
	return _DnsCost.Contract.GetAllPayments(&_DnsCost.CallOpts)
}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllSecondLevelNamePrice(opts *bind.CallOpts, ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllSecondLevelNamePrice", ownerNameHash_)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllSecondLevelNamePrice(ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNamePrice(&_DnsCost.CallOpts, ownerNameHash_)
}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllSecondLevelNamePrice(ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNamePrice(&_DnsCost.CallOpts, ownerNameHash_)
}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllSecondLevelNameTax(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllSecondLevelNameTax")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllSecondLevelNameTax() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNameTax(&_DnsCost.CallOpts)
}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllSecondLevelNameTax() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNameTax(&_DnsCost.CallOpts)
}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllTopLevelNamePrice(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllTopLevelNamePrice")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllTopLevelNamePrice() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllTopLevelNamePrice(&_DnsCost.CallOpts)
}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllTopLevelNamePrice() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllTopLevelNamePrice(&_DnsCost.CallOpts)
}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostCaller) GetExchangeWithId(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getExchangeWithId", erc20Addr_, lastPriceId_, price_)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostSession) GetExchangeWithId(erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetExchangeWithId(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, price_)
}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetExchangeWithId(erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetExchangeWithId(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, price_)
}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostCaller) GetLatestExchangeValue(opts *bind.CallOpts, erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getLatestExchangeValue", erc20Addr_, price_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostSession) GetLatestExchangeValue(erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	return _DnsCost.Contract.GetLatestExchangeValue(&_DnsCost.CallOpts, erc20Addr_, price_)
}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostCallerSession) GetLatestExchangeValue(erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	return _DnsCost.Contract.GetLatestExchangeValue(&_DnsCost.CallOpts, erc20Addr_, price_)
}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostCaller) GetPaymentsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getPaymentsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostSession) GetPaymentsCount() (*big.Int, error) {
	return _DnsCost.Contract.GetPaymentsCount(&_DnsCost.CallOpts)
}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostCallerSession) GetPaymentsCount() (*big.Int, error) {
	return _DnsCost.Contract.GetPaymentsCount(&_DnsCost.CallOpts)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_DnsCost *DnsCostCaller) GetSecondLevelNamePrice(opts *bind.CallOpts, fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getSecondLevelNamePrice", fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_DnsCost *DnsCostSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _DnsCost.Contract.GetSecondLevelNamePrice(&_DnsCost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _DnsCost.Contract.GetSecondLevelNamePrice(&_DnsCost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_DnsCost *DnsCostCaller) GetTopLevelNamePrice(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getTopLevelNamePrice", erc20Addr_, lastPriceId_, nameLength_, year_)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_DnsCost *DnsCostSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetTopLevelNamePrice(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetTopLevelNamePrice(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostCaller) PaymentIsExist(opts *bind.CallOpts, erc20Addr_ common.Address) (uint8, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentIsExist", erc20Addr_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostSession) PaymentIsExist(erc20Addr_ common.Address) (uint8, error) {
	return _DnsCost.Contract.PaymentIsExist(&_DnsCost.CallOpts, erc20Addr_)
}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostCallerSession) PaymentIsExist(erc20Addr_ common.Address) (uint8, error) {
	return _DnsCost.Contract.PaymentIsExist(&_DnsCost.CallOpts, erc20Addr_)
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostCaller) PaymentList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostSession) PaymentList(arg0 *big.Int) (common.Address, error) {
	return _DnsCost.Contract.PaymentList(&_DnsCost.CallOpts, arg0)
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostCallerSession) PaymentList(arg0 *big.Int) (common.Address, error) {
	return _DnsCost.Contract.PaymentList(&_DnsCost.CallOpts, arg0)
}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostCaller) PaymentMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentMap", arg0)

	outstruct := new(struct {
		Decimal uint8
		Enable  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimal = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Enable = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostSession) PaymentMap(arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	return _DnsCost.Contract.PaymentMap(&_DnsCost.CallOpts, arg0)
}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostCallerSession) PaymentMap(arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	return _DnsCost.Contract.PaymentMap(&_DnsCost.CallOpts, arg0)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostCaller) PriceIdIsValid(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "priceIdIsValid", erc20Addr_, lastPriceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _DnsCost.Contract.PriceIdIsValid(&_DnsCost.CallOpts, erc20Addr_, lastPriceId)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostCallerSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _DnsCost.Contract.PriceIdIsValid(&_DnsCost.CallOpts, erc20Addr_, lastPriceId)
}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostCaller) UsdtAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "usdtAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostSession) UsdtAddr() (common.Address, error) {
	return _DnsCost.Contract.UsdtAddr(&_DnsCost.CallOpts)
}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostCallerSession) UsdtAddr() (common.Address, error) {
	return _DnsCost.Contract.UsdtAddr(&_DnsCost.CallOpts)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostTransactor) SetSecondLevelNameTax(opts *bind.TransactOpts, tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "SetSecondLevelNameTax", tax_)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostSession) SetSecondLevelNameTax(tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNameTax(&_DnsCost.TransactOpts, tax_)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostTransactorSession) SetSecondLevelNameTax(tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNameTax(&_DnsCost.TransactOpts, tax_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostTransactor) SetTopLevelNamePrice(opts *bind.TransactOpts, price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "SetTopLevelNamePrice", price_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostSession) SetTopLevelNamePrice(price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetTopLevelNamePrice(&_DnsCost.TransactOpts, price_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostTransactorSession) SetTopLevelNamePrice(price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetTopLevelNamePrice(&_DnsCost.TransactOpts, price_)
}

// AddChainLinkAddr is a paid mutator transaction binding the contract method 0xbffe2b60.
//
// Solidity: function addChainLinkAddr(address erc20Addr_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostTransactor) AddChainLinkAddr(opts *bind.TransactOpts, erc20Addr_ common.Address, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "addChainLinkAddr", erc20Addr_, chainLinkContact_)
}

// AddChainLinkAddr is a paid mutator transaction binding the contract method 0xbffe2b60.
//
// Solidity: function addChainLinkAddr(address erc20Addr_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostSession) AddChainLinkAddr(erc20Addr_ common.Address, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.AddChainLinkAddr(&_DnsCost.TransactOpts, erc20Addr_, chainLinkContact_)
}

// AddChainLinkAddr is a paid mutator transaction binding the contract method 0xbffe2b60.
//
// Solidity: function addChainLinkAddr(address erc20Addr_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostTransactorSession) AddChainLinkAddr(erc20Addr_ common.Address, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.AddChainLinkAddr(&_DnsCost.TransactOpts, erc20Addr_, chainLinkContact_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3600a5df.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_) returns()
func (_DnsCost *DnsCostTransactor) AddPayment(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "addPayment", erc20Addr_, decimal_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3600a5df.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_) returns()
func (_DnsCost *DnsCostSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.AddPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3600a5df.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_) returns()
func (_DnsCost *DnsCostTransactorSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.AddPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_)
}

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostTransactor) DelPayment(opts *bind.TransactOpts, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "delPayment", erc20Addr_)
}

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostSession) DelPayment(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.DelPayment(&_DnsCost.TransactOpts, erc20Addr_)
}

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostTransactorSession) DelPayment(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.DelPayment(&_DnsCost.TransactOpts, erc20Addr_)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostTransactor) SetPayment(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setPayment", erc20Addr_, decimal_, enable)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostSession) SetPayment(erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.Contract.SetPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, enable)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostTransactorSession) SetPayment(erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.Contract.SetPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, enable)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostTransactor) SetRelation(opts *bind.TransactOpts, dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setRelation", dnstop_)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostSession) SetRelation(dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.SetRelation(&_DnsCost.TransactOpts, dnstop_)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostTransactorSession) SetRelation(dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.SetRelation(&_DnsCost.TransactOpts, dnstop_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostTransactor) SetSecondLevelNamePrice(opts *bind.TransactOpts, ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setSecondLevelNamePrice", ownerNameHash_, prices_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostSession) SetSecondLevelNamePrice(ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNamePrice(&_DnsCost.TransactOpts, ownerNameHash_, prices_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostTransactorSession) SetSecondLevelNamePrice(ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNamePrice(&_DnsCost.TransactOpts, ownerNameHash_, prices_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostTransactor) SetUsdtAddr(opts *bind.TransactOpts, usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setUsdtAddr", usdtTokenAddr_, baseDecimal_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostSession) SetUsdtAddr(usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.SetUsdtAddr(&_DnsCost.TransactOpts, usdtTokenAddr_, baseDecimal_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostTransactorSession) SetUsdtAddr(usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.SetUsdtAddr(&_DnsCost.TransactOpts, usdtTokenAddr_, baseDecimal_)
}
