// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc

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

// DnsPriceABI is the input ABI used to generate the binding from.
const DnsPriceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentC_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultPrice\",\"type\":\"uint256\"}],\"name\":\"EvAddPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EvAddPriceLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EvAddPriceLenArray\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"EvDelPriceLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"EvRemovePrice\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaultPrice_\",\"type\":\"uint256\"}],\"name\":\"AddPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"AddPriceLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaultPrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"lenArr_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"priceArr_\",\"type\":\"uint256[]\"}],\"name\":\"AddPriceLenArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"name\":\"DelPriceLen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"Erc20IsSupport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"name\":\"Price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256[8]\",\"name\":\"lenArr_\",\"type\":\"uint256[8]\"}],\"name\":\"PriceGet\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"len_\",\"type\":\"uint256\"}],\"name\":\"PriceGetOne\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"PriceInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"RemovePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"namePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"defaultPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferDaoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DnsPrice is an auto generated Go binding around an Ethereum contract.
type DnsPrice struct {
	DnsPriceCaller     // Read-only binding to the contract
	DnsPriceTransactor // Write-only binding to the contract
	DnsPriceFilterer   // Log filterer for contract events
}

// DnsPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsPriceSession struct {
	Contract     *DnsPrice         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsPriceCallerSession struct {
	Contract *DnsPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DnsPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsPriceTransactorSession struct {
	Contract     *DnsPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DnsPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsPriceRaw struct {
	Contract *DnsPrice // Generic contract binding to access the raw methods on
}

// DnsPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsPriceCallerRaw struct {
	Contract *DnsPriceCaller // Generic read-only contract binding to access the raw methods on
}

// DnsPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsPriceTransactorRaw struct {
	Contract *DnsPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsPrice creates a new instance of DnsPrice, bound to a specific deployed contract.
func NewDnsPrice(address common.Address, backend bind.ContractBackend) (*DnsPrice, error) {
	contract, err := bindDnsPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsPrice{DnsPriceCaller: DnsPriceCaller{contract: contract}, DnsPriceTransactor: DnsPriceTransactor{contract: contract}, DnsPriceFilterer: DnsPriceFilterer{contract: contract}}, nil
}

// NewDnsPriceCaller creates a new read-only instance of DnsPrice, bound to a specific deployed contract.
func NewDnsPriceCaller(address common.Address, caller bind.ContractCaller) (*DnsPriceCaller, error) {
	contract, err := bindDnsPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsPriceCaller{contract: contract}, nil
}

// NewDnsPriceTransactor creates a new write-only instance of DnsPrice, bound to a specific deployed contract.
func NewDnsPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsPriceTransactor, error) {
	contract, err := bindDnsPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsPriceTransactor{contract: contract}, nil
}

// NewDnsPriceFilterer creates a new log filterer instance of DnsPrice, bound to a specific deployed contract.
func NewDnsPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsPriceFilterer, error) {
	contract, err := bindDnsPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsPriceFilterer{contract: contract}, nil
}

// bindDnsPrice binds a generic wrapper to an already deployed contract.
func bindDnsPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsPriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsPrice *DnsPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsPrice.Contract.DnsPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsPrice *DnsPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsPrice.Contract.DnsPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsPrice *DnsPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsPrice.Contract.DnsPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsPrice *DnsPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsPrice *DnsPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsPrice *DnsPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsPrice.Contract.contract.Transact(opts, method, params...)
}

// Erc20IsSupport is a free data retrieval call binding the contract method 0x2a61750f.
//
// Solidity: function Erc20IsSupport(bytes32 hash_, address erc20Addr_) view returns(bool)
func (_DnsPrice *DnsPriceCaller) Erc20IsSupport(opts *bind.CallOpts, hash_ [32]byte, erc20Addr_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "Erc20IsSupport", hash_, erc20Addr_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Erc20IsSupport is a free data retrieval call binding the contract method 0x2a61750f.
//
// Solidity: function Erc20IsSupport(bytes32 hash_, address erc20Addr_) view returns(bool)
func (_DnsPrice *DnsPriceSession) Erc20IsSupport(hash_ [32]byte, erc20Addr_ common.Address) (bool, error) {
	return _DnsPrice.Contract.Erc20IsSupport(&_DnsPrice.CallOpts, hash_, erc20Addr_)
}

// Erc20IsSupport is a free data retrieval call binding the contract method 0x2a61750f.
//
// Solidity: function Erc20IsSupport(bytes32 hash_, address erc20Addr_) view returns(bool)
func (_DnsPrice *DnsPriceCallerSession) Erc20IsSupport(hash_ [32]byte, erc20Addr_ common.Address) (bool, error) {
	return _DnsPrice.Contract.Erc20IsSupport(&_DnsPrice.CallOpts, hash_, erc20Addr_)
}

// Price is a free data retrieval call binding the contract method 0x18ddc403.
//
// Solidity: function Price(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(uint256)
func (_DnsPrice *DnsPriceCaller) Price(opts *bind.CallOpts, hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "Price", hash_, erc20Addr_, len_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x18ddc403.
//
// Solidity: function Price(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(uint256)
func (_DnsPrice *DnsPriceSession) Price(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*big.Int, error) {
	return _DnsPrice.Contract.Price(&_DnsPrice.CallOpts, hash_, erc20Addr_, len_)
}

// Price is a free data retrieval call binding the contract method 0x18ddc403.
//
// Solidity: function Price(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(uint256)
func (_DnsPrice *DnsPriceCallerSession) Price(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*big.Int, error) {
	return _DnsPrice.Contract.Price(&_DnsPrice.CallOpts, hash_, erc20Addr_, len_)
}

// PriceGet is a free data retrieval call binding the contract method 0xe7ff0d15.
//
// Solidity: function PriceGet(bytes32 hash_, address erc20Addr_, uint256[8] lenArr_) view returns(uint256[8])
func (_DnsPrice *DnsPriceCaller) PriceGet(opts *bind.CallOpts, hash_ [32]byte, erc20Addr_ common.Address, lenArr_ [8]*big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "PriceGet", hash_, erc20Addr_, lenArr_)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// PriceGet is a free data retrieval call binding the contract method 0xe7ff0d15.
//
// Solidity: function PriceGet(bytes32 hash_, address erc20Addr_, uint256[8] lenArr_) view returns(uint256[8])
func (_DnsPrice *DnsPriceSession) PriceGet(hash_ [32]byte, erc20Addr_ common.Address, lenArr_ [8]*big.Int) ([8]*big.Int, error) {
	return _DnsPrice.Contract.PriceGet(&_DnsPrice.CallOpts, hash_, erc20Addr_, lenArr_)
}

// PriceGet is a free data retrieval call binding the contract method 0xe7ff0d15.
//
// Solidity: function PriceGet(bytes32 hash_, address erc20Addr_, uint256[8] lenArr_) view returns(uint256[8])
func (_DnsPrice *DnsPriceCallerSession) PriceGet(hash_ [32]byte, erc20Addr_ common.Address, lenArr_ [8]*big.Int) ([8]*big.Int, error) {
	return _DnsPrice.Contract.PriceGet(&_DnsPrice.CallOpts, hash_, erc20Addr_, lenArr_)
}

// PriceGetOne is a free data retrieval call binding the contract method 0xaa13993d.
//
// Solidity: function PriceGetOne(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(bool, uint256)
func (_DnsPrice *DnsPriceCaller) PriceGetOne(opts *bind.CallOpts, hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "PriceGetOne", hash_, erc20Addr_, len_)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PriceGetOne is a free data retrieval call binding the contract method 0xaa13993d.
//
// Solidity: function PriceGetOne(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(bool, uint256)
func (_DnsPrice *DnsPriceSession) PriceGetOne(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (bool, *big.Int, error) {
	return _DnsPrice.Contract.PriceGetOne(&_DnsPrice.CallOpts, hash_, erc20Addr_, len_)
}

// PriceGetOne is a free data retrieval call binding the contract method 0xaa13993d.
//
// Solidity: function PriceGetOne(bytes32 hash_, address erc20Addr_, uint256 len_) view returns(bool, uint256)
func (_DnsPrice *DnsPriceCallerSession) PriceGetOne(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (bool, *big.Int, error) {
	return _DnsPrice.Contract.PriceGetOne(&_DnsPrice.CallOpts, hash_, erc20Addr_, len_)
}

// NamePrice is a free data retrieval call binding the contract method 0x691ac382.
//
// Solidity: function namePrice(bytes32 , address ) view returns(uint256 defaultPrice)
func (_DnsPrice *DnsPriceCaller) NamePrice(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "namePrice", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NamePrice is a free data retrieval call binding the contract method 0x691ac382.
//
// Solidity: function namePrice(bytes32 , address ) view returns(uint256 defaultPrice)
func (_DnsPrice *DnsPriceSession) NamePrice(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _DnsPrice.Contract.NamePrice(&_DnsPrice.CallOpts, arg0, arg1)
}

// NamePrice is a free data retrieval call binding the contract method 0x691ac382.
//
// Solidity: function namePrice(bytes32 , address ) view returns(uint256 defaultPrice)
func (_DnsPrice *DnsPriceCallerSession) NamePrice(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _DnsPrice.Contract.NamePrice(&_DnsPrice.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPrice *DnsPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPrice *DnsPriceSession) Owner() (common.Address, error) {
	return _DnsPrice.Contract.Owner(&_DnsPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsPrice *DnsPriceCallerSession) Owner() (common.Address, error) {
	return _DnsPrice.Contract.Owner(&_DnsPrice.CallOpts)
}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DnsPrice *DnsPriceCaller) PaymentC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsPrice.contract.Call(opts, &out, "paymentC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DnsPrice *DnsPriceSession) PaymentC() (common.Address, error) {
	return _DnsPrice.Contract.PaymentC(&_DnsPrice.CallOpts)
}

// PaymentC is a free data retrieval call binding the contract method 0x9c09a102.
//
// Solidity: function paymentC() view returns(address)
func (_DnsPrice *DnsPriceCallerSession) PaymentC() (common.Address, error) {
	return _DnsPrice.Contract.PaymentC(&_DnsPrice.CallOpts)
}

// AddPrice is a paid mutator transaction binding the contract method 0x3dcdef0d.
//
// Solidity: function AddPrice(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_) returns()
func (_DnsPrice *DnsPriceTransactor) AddPrice(opts *bind.TransactOpts, hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "AddPrice", hash_, erc20Addr_, defaultPrice_)
}

// AddPrice is a paid mutator transaction binding the contract method 0x3dcdef0d.
//
// Solidity: function AddPrice(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_) returns()
func (_DnsPrice *DnsPriceSession) AddPrice(hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPrice(&_DnsPrice.TransactOpts, hash_, erc20Addr_, defaultPrice_)
}

// AddPrice is a paid mutator transaction binding the contract method 0x3dcdef0d.
//
// Solidity: function AddPrice(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_) returns()
func (_DnsPrice *DnsPriceTransactorSession) AddPrice(hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPrice(&_DnsPrice.TransactOpts, hash_, erc20Addr_, defaultPrice_)
}

// AddPriceLen is a paid mutator transaction binding the contract method 0xb8a09c73.
//
// Solidity: function AddPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_, uint256 price_) returns()
func (_DnsPrice *DnsPriceTransactor) AddPriceLen(opts *bind.TransactOpts, hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "AddPriceLen", hash_, erc20Addr_, len_, price_)
}

// AddPriceLen is a paid mutator transaction binding the contract method 0xb8a09c73.
//
// Solidity: function AddPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_, uint256 price_) returns()
func (_DnsPrice *DnsPriceSession) AddPriceLen(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPriceLen(&_DnsPrice.TransactOpts, hash_, erc20Addr_, len_, price_)
}

// AddPriceLen is a paid mutator transaction binding the contract method 0xb8a09c73.
//
// Solidity: function AddPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_, uint256 price_) returns()
func (_DnsPrice *DnsPriceTransactorSession) AddPriceLen(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int, price_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPriceLen(&_DnsPrice.TransactOpts, hash_, erc20Addr_, len_, price_)
}

// AddPriceLenArray is a paid mutator transaction binding the contract method 0xf22ae031.
//
// Solidity: function AddPriceLenArray(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_, uint256[] lenArr_, uint256[] priceArr_) returns()
func (_DnsPrice *DnsPriceTransactor) AddPriceLenArray(opts *bind.TransactOpts, hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int, lenArr_ []*big.Int, priceArr_ []*big.Int) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "AddPriceLenArray", hash_, erc20Addr_, defaultPrice_, lenArr_, priceArr_)
}

// AddPriceLenArray is a paid mutator transaction binding the contract method 0xf22ae031.
//
// Solidity: function AddPriceLenArray(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_, uint256[] lenArr_, uint256[] priceArr_) returns()
func (_DnsPrice *DnsPriceSession) AddPriceLenArray(hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int, lenArr_ []*big.Int, priceArr_ []*big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPriceLenArray(&_DnsPrice.TransactOpts, hash_, erc20Addr_, defaultPrice_, lenArr_, priceArr_)
}

// AddPriceLenArray is a paid mutator transaction binding the contract method 0xf22ae031.
//
// Solidity: function AddPriceLenArray(bytes32 hash_, address erc20Addr_, uint256 defaultPrice_, uint256[] lenArr_, uint256[] priceArr_) returns()
func (_DnsPrice *DnsPriceTransactorSession) AddPriceLenArray(hash_ [32]byte, erc20Addr_ common.Address, defaultPrice_ *big.Int, lenArr_ []*big.Int, priceArr_ []*big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.AddPriceLenArray(&_DnsPrice.TransactOpts, hash_, erc20Addr_, defaultPrice_, lenArr_, priceArr_)
}

// DelPriceLen is a paid mutator transaction binding the contract method 0x23777e25.
//
// Solidity: function DelPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_) returns()
func (_DnsPrice *DnsPriceTransactor) DelPriceLen(opts *bind.TransactOpts, hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "DelPriceLen", hash_, erc20Addr_, len_)
}

// DelPriceLen is a paid mutator transaction binding the contract method 0x23777e25.
//
// Solidity: function DelPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_) returns()
func (_DnsPrice *DnsPriceSession) DelPriceLen(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.DelPriceLen(&_DnsPrice.TransactOpts, hash_, erc20Addr_, len_)
}

// DelPriceLen is a paid mutator transaction binding the contract method 0x23777e25.
//
// Solidity: function DelPriceLen(bytes32 hash_, address erc20Addr_, uint256 len_) returns()
func (_DnsPrice *DnsPriceTransactorSession) DelPriceLen(hash_ [32]byte, erc20Addr_ common.Address, len_ *big.Int) (*types.Transaction, error) {
	return _DnsPrice.Contract.DelPriceLen(&_DnsPrice.TransactOpts, hash_, erc20Addr_, len_)
}

// PriceInit is a paid mutator transaction binding the contract method 0x31dbec04.
//
// Solidity: function PriceInit(bytes32 hash_) returns()
func (_DnsPrice *DnsPriceTransactor) PriceInit(opts *bind.TransactOpts, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "PriceInit", hash_)
}

// PriceInit is a paid mutator transaction binding the contract method 0x31dbec04.
//
// Solidity: function PriceInit(bytes32 hash_) returns()
func (_DnsPrice *DnsPriceSession) PriceInit(hash_ [32]byte) (*types.Transaction, error) {
	return _DnsPrice.Contract.PriceInit(&_DnsPrice.TransactOpts, hash_)
}

// PriceInit is a paid mutator transaction binding the contract method 0x31dbec04.
//
// Solidity: function PriceInit(bytes32 hash_) returns()
func (_DnsPrice *DnsPriceTransactorSession) PriceInit(hash_ [32]byte) (*types.Transaction, error) {
	return _DnsPrice.Contract.PriceInit(&_DnsPrice.TransactOpts, hash_)
}

// RemovePrice is a paid mutator transaction binding the contract method 0x071ba617.
//
// Solidity: function RemovePrice(bytes32 hash_, address erc20Addr_) returns()
func (_DnsPrice *DnsPriceTransactor) RemovePrice(opts *bind.TransactOpts, hash_ [32]byte, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "RemovePrice", hash_, erc20Addr_)
}

// RemovePrice is a paid mutator transaction binding the contract method 0x071ba617.
//
// Solidity: function RemovePrice(bytes32 hash_, address erc20Addr_) returns()
func (_DnsPrice *DnsPriceSession) RemovePrice(hash_ [32]byte, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.Contract.RemovePrice(&_DnsPrice.TransactOpts, hash_, erc20Addr_)
}

// RemovePrice is a paid mutator transaction binding the contract method 0x071ba617.
//
// Solidity: function RemovePrice(bytes32 hash_, address erc20Addr_) returns()
func (_DnsPrice *DnsPriceTransactorSession) RemovePrice(hash_ [32]byte, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.Contract.RemovePrice(&_DnsPrice.TransactOpts, hash_, erc20Addr_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPrice *DnsPriceTransactor) TransferDaoOwner(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.contract.Transact(opts, "transferDaoOwner", newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPrice *DnsPriceSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.Contract.TransferDaoOwner(&_DnsPrice.TransactOpts, newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsPrice *DnsPriceTransactorSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsPrice.Contract.TransferDaoOwner(&_DnsPrice.TransactOpts, newOwner_)
}

// DnsPriceEvAddPriceIterator is returned from FilterEvAddPrice and is used to iterate over the raw logs and unpacked data for EvAddPrice events raised by the DnsPrice contract.
type DnsPriceEvAddPriceIterator struct {
	Event *DnsPriceEvAddPrice // Event containing the contract specifics and raw log

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
func (it *DnsPriceEvAddPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsPriceEvAddPrice)
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
		it.Event = new(DnsPriceEvAddPrice)
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
func (it *DnsPriceEvAddPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsPriceEvAddPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsPriceEvAddPrice represents a EvAddPrice event raised by the DnsPrice contract.
type DnsPriceEvAddPrice struct {
	NameHash     [32]byte
	Erc20Addr    common.Address
	DefaultPrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvAddPrice is a free log retrieval operation binding the contract event 0x14718b88b3d0cc8bca875969434759d1055149c43057e62b3b21ba7a740ffcdd.
//
// Solidity: event EvAddPrice(bytes32 nameHash, address erc20Addr, uint256 defaultPrice)
func (_DnsPrice *DnsPriceFilterer) FilterEvAddPrice(opts *bind.FilterOpts) (*DnsPriceEvAddPriceIterator, error) {

	logs, sub, err := _DnsPrice.contract.FilterLogs(opts, "EvAddPrice")
	if err != nil {
		return nil, err
	}
	return &DnsPriceEvAddPriceIterator{contract: _DnsPrice.contract, event: "EvAddPrice", logs: logs, sub: sub}, nil
}

// WatchEvAddPrice is a free log subscription operation binding the contract event 0x14718b88b3d0cc8bca875969434759d1055149c43057e62b3b21ba7a740ffcdd.
//
// Solidity: event EvAddPrice(bytes32 nameHash, address erc20Addr, uint256 defaultPrice)
func (_DnsPrice *DnsPriceFilterer) WatchEvAddPrice(opts *bind.WatchOpts, sink chan<- *DnsPriceEvAddPrice) (event.Subscription, error) {

	logs, sub, err := _DnsPrice.contract.WatchLogs(opts, "EvAddPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsPriceEvAddPrice)
				if err := _DnsPrice.contract.UnpackLog(event, "EvAddPrice", log); err != nil {
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

// ParseEvAddPrice is a log parse operation binding the contract event 0x14718b88b3d0cc8bca875969434759d1055149c43057e62b3b21ba7a740ffcdd.
//
// Solidity: event EvAddPrice(bytes32 nameHash, address erc20Addr, uint256 defaultPrice)
func (_DnsPrice *DnsPriceFilterer) ParseEvAddPrice(log types.Log) (*DnsPriceEvAddPrice, error) {
	event := new(DnsPriceEvAddPrice)
	if err := _DnsPrice.contract.UnpackLog(event, "EvAddPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsPriceEvAddPriceLenIterator is returned from FilterEvAddPriceLen and is used to iterate over the raw logs and unpacked data for EvAddPriceLen events raised by the DnsPrice contract.
type DnsPriceEvAddPriceLenIterator struct {
	Event *DnsPriceEvAddPriceLen // Event containing the contract specifics and raw log

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
func (it *DnsPriceEvAddPriceLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsPriceEvAddPriceLen)
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
		it.Event = new(DnsPriceEvAddPriceLen)
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
func (it *DnsPriceEvAddPriceLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsPriceEvAddPriceLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsPriceEvAddPriceLen represents a EvAddPriceLen event raised by the DnsPrice contract.
type DnsPriceEvAddPriceLen struct {
	NameHash  [32]byte
	Erc20Addr common.Address
	Len       *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvAddPriceLen is a free log retrieval operation binding the contract event 0x3d5394a70d243ed7877e411e2017691917cd6155a1f4ec33696d91cd4ff61a43.
//
// Solidity: event EvAddPriceLen(bytes32 nameHash, address erc20Addr, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) FilterEvAddPriceLen(opts *bind.FilterOpts) (*DnsPriceEvAddPriceLenIterator, error) {

	logs, sub, err := _DnsPrice.contract.FilterLogs(opts, "EvAddPriceLen")
	if err != nil {
		return nil, err
	}
	return &DnsPriceEvAddPriceLenIterator{contract: _DnsPrice.contract, event: "EvAddPriceLen", logs: logs, sub: sub}, nil
}

// WatchEvAddPriceLen is a free log subscription operation binding the contract event 0x3d5394a70d243ed7877e411e2017691917cd6155a1f4ec33696d91cd4ff61a43.
//
// Solidity: event EvAddPriceLen(bytes32 nameHash, address erc20Addr, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) WatchEvAddPriceLen(opts *bind.WatchOpts, sink chan<- *DnsPriceEvAddPriceLen) (event.Subscription, error) {

	logs, sub, err := _DnsPrice.contract.WatchLogs(opts, "EvAddPriceLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsPriceEvAddPriceLen)
				if err := _DnsPrice.contract.UnpackLog(event, "EvAddPriceLen", log); err != nil {
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

// ParseEvAddPriceLen is a log parse operation binding the contract event 0x3d5394a70d243ed7877e411e2017691917cd6155a1f4ec33696d91cd4ff61a43.
//
// Solidity: event EvAddPriceLen(bytes32 nameHash, address erc20Addr, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) ParseEvAddPriceLen(log types.Log) (*DnsPriceEvAddPriceLen, error) {
	event := new(DnsPriceEvAddPriceLen)
	if err := _DnsPrice.contract.UnpackLog(event, "EvAddPriceLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsPriceEvAddPriceLenArrayIterator is returned from FilterEvAddPriceLenArray and is used to iterate over the raw logs and unpacked data for EvAddPriceLenArray events raised by the DnsPrice contract.
type DnsPriceEvAddPriceLenArrayIterator struct {
	Event *DnsPriceEvAddPriceLenArray // Event containing the contract specifics and raw log

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
func (it *DnsPriceEvAddPriceLenArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsPriceEvAddPriceLenArray)
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
		it.Event = new(DnsPriceEvAddPriceLenArray)
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
func (it *DnsPriceEvAddPriceLenArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsPriceEvAddPriceLenArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsPriceEvAddPriceLenArray represents a EvAddPriceLenArray event raised by the DnsPrice contract.
type DnsPriceEvAddPriceLenArray struct {
	NameHash     [32]byte
	Erc20Addr    common.Address
	DefaultPrice *big.Int
	Len          *big.Int
	Price        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvAddPriceLenArray is a free log retrieval operation binding the contract event 0xa47b2d894fa9200a3f0a13162acf2e46d2b96f59fbd8aa1b43a554df2b3c21a9.
//
// Solidity: event EvAddPriceLenArray(bytes32 nameHash, address erc20Addr, uint256 defaultPrice, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) FilterEvAddPriceLenArray(opts *bind.FilterOpts) (*DnsPriceEvAddPriceLenArrayIterator, error) {

	logs, sub, err := _DnsPrice.contract.FilterLogs(opts, "EvAddPriceLenArray")
	if err != nil {
		return nil, err
	}
	return &DnsPriceEvAddPriceLenArrayIterator{contract: _DnsPrice.contract, event: "EvAddPriceLenArray", logs: logs, sub: sub}, nil
}

// WatchEvAddPriceLenArray is a free log subscription operation binding the contract event 0xa47b2d894fa9200a3f0a13162acf2e46d2b96f59fbd8aa1b43a554df2b3c21a9.
//
// Solidity: event EvAddPriceLenArray(bytes32 nameHash, address erc20Addr, uint256 defaultPrice, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) WatchEvAddPriceLenArray(opts *bind.WatchOpts, sink chan<- *DnsPriceEvAddPriceLenArray) (event.Subscription, error) {

	logs, sub, err := _DnsPrice.contract.WatchLogs(opts, "EvAddPriceLenArray")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsPriceEvAddPriceLenArray)
				if err := _DnsPrice.contract.UnpackLog(event, "EvAddPriceLenArray", log); err != nil {
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

// ParseEvAddPriceLenArray is a log parse operation binding the contract event 0xa47b2d894fa9200a3f0a13162acf2e46d2b96f59fbd8aa1b43a554df2b3c21a9.
//
// Solidity: event EvAddPriceLenArray(bytes32 nameHash, address erc20Addr, uint256 defaultPrice, uint256 len, uint256 price)
func (_DnsPrice *DnsPriceFilterer) ParseEvAddPriceLenArray(log types.Log) (*DnsPriceEvAddPriceLenArray, error) {
	event := new(DnsPriceEvAddPriceLenArray)
	if err := _DnsPrice.contract.UnpackLog(event, "EvAddPriceLenArray", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsPriceEvDelPriceLenIterator is returned from FilterEvDelPriceLen and is used to iterate over the raw logs and unpacked data for EvDelPriceLen events raised by the DnsPrice contract.
type DnsPriceEvDelPriceLenIterator struct {
	Event *DnsPriceEvDelPriceLen // Event containing the contract specifics and raw log

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
func (it *DnsPriceEvDelPriceLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsPriceEvDelPriceLen)
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
		it.Event = new(DnsPriceEvDelPriceLen)
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
func (it *DnsPriceEvDelPriceLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsPriceEvDelPriceLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsPriceEvDelPriceLen represents a EvDelPriceLen event raised by the DnsPrice contract.
type DnsPriceEvDelPriceLen struct {
	NameHash  [32]byte
	Erc20Addr common.Address
	Len       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvDelPriceLen is a free log retrieval operation binding the contract event 0x9a36ac0b064295f7f284858cf951e2e4735000fa437106d3a6c98190fdcd8151.
//
// Solidity: event EvDelPriceLen(bytes32 nameHash, address erc20Addr, uint256 len)
func (_DnsPrice *DnsPriceFilterer) FilterEvDelPriceLen(opts *bind.FilterOpts) (*DnsPriceEvDelPriceLenIterator, error) {

	logs, sub, err := _DnsPrice.contract.FilterLogs(opts, "EvDelPriceLen")
	if err != nil {
		return nil, err
	}
	return &DnsPriceEvDelPriceLenIterator{contract: _DnsPrice.contract, event: "EvDelPriceLen", logs: logs, sub: sub}, nil
}

// WatchEvDelPriceLen is a free log subscription operation binding the contract event 0x9a36ac0b064295f7f284858cf951e2e4735000fa437106d3a6c98190fdcd8151.
//
// Solidity: event EvDelPriceLen(bytes32 nameHash, address erc20Addr, uint256 len)
func (_DnsPrice *DnsPriceFilterer) WatchEvDelPriceLen(opts *bind.WatchOpts, sink chan<- *DnsPriceEvDelPriceLen) (event.Subscription, error) {

	logs, sub, err := _DnsPrice.contract.WatchLogs(opts, "EvDelPriceLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsPriceEvDelPriceLen)
				if err := _DnsPrice.contract.UnpackLog(event, "EvDelPriceLen", log); err != nil {
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

// ParseEvDelPriceLen is a log parse operation binding the contract event 0x9a36ac0b064295f7f284858cf951e2e4735000fa437106d3a6c98190fdcd8151.
//
// Solidity: event EvDelPriceLen(bytes32 nameHash, address erc20Addr, uint256 len)
func (_DnsPrice *DnsPriceFilterer) ParseEvDelPriceLen(log types.Log) (*DnsPriceEvDelPriceLen, error) {
	event := new(DnsPriceEvDelPriceLen)
	if err := _DnsPrice.contract.UnpackLog(event, "EvDelPriceLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsPriceEvRemovePriceIterator is returned from FilterEvRemovePrice and is used to iterate over the raw logs and unpacked data for EvRemovePrice events raised by the DnsPrice contract.
type DnsPriceEvRemovePriceIterator struct {
	Event *DnsPriceEvRemovePrice // Event containing the contract specifics and raw log

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
func (it *DnsPriceEvRemovePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsPriceEvRemovePrice)
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
		it.Event = new(DnsPriceEvRemovePrice)
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
func (it *DnsPriceEvRemovePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsPriceEvRemovePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsPriceEvRemovePrice represents a EvRemovePrice event raised by the DnsPrice contract.
type DnsPriceEvRemovePrice struct {
	NameHash  [32]byte
	Erc20Addr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvRemovePrice is a free log retrieval operation binding the contract event 0xf716ac1e31a12e82d9009844fe6546f66a61932474b13644543b2aa1b46ec1bd.
//
// Solidity: event EvRemovePrice(bytes32 nameHash, address erc20Addr)
func (_DnsPrice *DnsPriceFilterer) FilterEvRemovePrice(opts *bind.FilterOpts) (*DnsPriceEvRemovePriceIterator, error) {

	logs, sub, err := _DnsPrice.contract.FilterLogs(opts, "EvRemovePrice")
	if err != nil {
		return nil, err
	}
	return &DnsPriceEvRemovePriceIterator{contract: _DnsPrice.contract, event: "EvRemovePrice", logs: logs, sub: sub}, nil
}

// WatchEvRemovePrice is a free log subscription operation binding the contract event 0xf716ac1e31a12e82d9009844fe6546f66a61932474b13644543b2aa1b46ec1bd.
//
// Solidity: event EvRemovePrice(bytes32 nameHash, address erc20Addr)
func (_DnsPrice *DnsPriceFilterer) WatchEvRemovePrice(opts *bind.WatchOpts, sink chan<- *DnsPriceEvRemovePrice) (event.Subscription, error) {

	logs, sub, err := _DnsPrice.contract.WatchLogs(opts, "EvRemovePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsPriceEvRemovePrice)
				if err := _DnsPrice.contract.UnpackLog(event, "EvRemovePrice", log); err != nil {
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

// ParseEvRemovePrice is a log parse operation binding the contract event 0xf716ac1e31a12e82d9009844fe6546f66a61932474b13644543b2aa1b46ec1bd.
//
// Solidity: event EvRemovePrice(bytes32 nameHash, address erc20Addr)
func (_DnsPrice *DnsPriceFilterer) ParseEvRemovePrice(log types.Log) (*DnsPriceEvRemovePrice, error) {
	event := new(DnsPriceEvRemovePrice)
	if err := _DnsPrice.contract.UnpackLog(event, "EvRemovePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
