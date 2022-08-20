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

// BasOANNABI is the input ABI used to generate the binding from.
const BasOANNABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllocationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Cost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BROOT_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CUSTOM_PRICE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_CUSTOEMED_SUB_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_CUSTOEMED_SUB_O\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_NORMAL_SUB_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_NORMAL_SUB_O\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_ROOT_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_SELF_SUB_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_SELF_SUB_O\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUB_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reg_root_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_self_sub_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_self_sub_o\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_normal_sub_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_normal_sub_o\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_custom_sub_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_custom_sub_o\",\"type\":\"uint256\"}],\"name\":\"allocationSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"openCustomPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aroot_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"broot_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sub_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"custom_price_gas\",\"type\":\"uint256\"}],\"name\":\"priceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"rechargeRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"rechargeSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"registerSub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BasOANN is an auto generated Go binding around an Ethereum contract.
type BasOANN struct {
	BasOANNCaller     // Read-only binding to the contract
	BasOANNTransactor // Write-only binding to the contract
	BasOANNFilterer   // Log filterer for contract events
}

// BasOANNCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasOANNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasOANNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasOANNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasOANNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasOANNSession struct {
	Contract     *BasOANN          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasOANNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasOANNCallerSession struct {
	Contract *BasOANNCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BasOANNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasOANNTransactorSession struct {
	Contract     *BasOANNTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasOANNRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasOANNRaw struct {
	Contract *BasOANN // Generic contract binding to access the raw methods on
}

// BasOANNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasOANNCallerRaw struct {
	Contract *BasOANNCaller // Generic read-only contract binding to access the raw methods on
}

// BasOANNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasOANNTransactorRaw struct {
	Contract *BasOANNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasOANN creates a new instance of BasOANN, bound to a specific deployed contract.
func NewBasOANN(address common.Address, backend bind.ContractBackend) (*BasOANN, error) {
	contract, err := bindBasOANN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasOANN{BasOANNCaller: BasOANNCaller{contract: contract}, BasOANNTransactor: BasOANNTransactor{contract: contract}, BasOANNFilterer: BasOANNFilterer{contract: contract}}, nil
}

// NewBasOANNCaller creates a new read-only instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNCaller(address common.Address, caller bind.ContractCaller) (*BasOANNCaller, error) {
	contract, err := bindBasOANN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasOANNCaller{contract: contract}, nil
}

// NewBasOANNTransactor creates a new write-only instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNTransactor(address common.Address, transactor bind.ContractTransactor) (*BasOANNTransactor, error) {
	contract, err := bindBasOANN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasOANNTransactor{contract: contract}, nil
}

// NewBasOANNFilterer creates a new log filterer instance of BasOANN, bound to a specific deployed contract.
func NewBasOANNFilterer(address common.Address, filterer bind.ContractFilterer) (*BasOANNFilterer, error) {
	contract, err := bindBasOANN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasOANNFilterer{contract: contract}, nil
}

// bindBasOANN binds a generic wrapper to an already deployed contract.
func bindBasOANN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasOANNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOANN *BasOANNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasOANN.Contract.BasOANNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOANN *BasOANNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOANN.Contract.BasOANNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOANN *BasOANNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOANN.Contract.BasOANNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasOANN *BasOANNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasOANN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasOANN *BasOANNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasOANN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasOANN *BasOANNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasOANN.Contract.contract.Transact(opts, method, params...)
}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNCaller) AROOTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "AROOT_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNSession) AROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.AROOTGAS(&_BasOANN.CallOpts)
}

// AROOTGAS is a free data retrieval call binding the contract method 0x789dba4f.
//
// Solidity: function AROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) AROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.AROOTGAS(&_BasOANN.CallOpts)
}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNCaller) BROOTGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "BROOT_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNSession) BROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.BROOTGAS(&_BasOANN.CallOpts)
}

// BROOTGAS is a free data retrieval call binding the contract method 0xfcb26830.
//
// Solidity: function BROOT_GAS() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) BROOTGAS() (*big.Int, error) {
	return _BasOANN.Contract.BROOTGAS(&_BasOANN.CallOpts)
}

// CUSTOMPRICEGAS is a free data retrieval call binding the contract method 0xef59fef9.
//
// Solidity: function CUSTOM_PRICE_GAS() view returns(uint256)
func (_BasOANN *BasOANNCaller) CUSTOMPRICEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "CUSTOM_PRICE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CUSTOMPRICEGAS is a free data retrieval call binding the contract method 0xef59fef9.
//
// Solidity: function CUSTOM_PRICE_GAS() view returns(uint256)
func (_BasOANN *BasOANNSession) CUSTOMPRICEGAS() (*big.Int, error) {
	return _BasOANN.Contract.CUSTOMPRICEGAS(&_BasOANN.CallOpts)
}

// CUSTOMPRICEGAS is a free data retrieval call binding the contract method 0xef59fef9.
//
// Solidity: function CUSTOM_PRICE_GAS() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) CUSTOMPRICEGAS() (*big.Int, error) {
	return _BasOANN.Contract.CUSTOMPRICEGAS(&_BasOANN.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasOANN *BasOANNCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasOANN *BasOANNSession) DAOAddress() (common.Address, error) {
	return _BasOANN.Contract.DAOAddress(&_BasOANN.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasOANN *BasOANNCallerSession) DAOAddress() (common.Address, error) {
	return _BasOANN.Contract.DAOAddress(&_BasOANN.CallOpts)
}

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() view returns(uint256)
func (_BasOANN *BasOANNCaller) MAXYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "MAX_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() view returns(uint256)
func (_BasOANN *BasOANNSession) MAXYEAR() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEAR(&_BasOANN.CallOpts)
}

// MAXYEAR is a free data retrieval call binding the contract method 0x61c9511e.
//
// Solidity: function MAX_YEAR() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) MAXYEAR() (*big.Int, error) {
	return _BasOANN.Contract.MAXYEAR(&_BasOANN.CallOpts)
}

// REGCUSTOEMEDSUBM is a free data retrieval call binding the contract method 0x66561d29.
//
// Solidity: function REG_CUSTOEMED_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGCUSTOEMEDSUBM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_CUSTOEMED_SUB_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGCUSTOEMEDSUBM is a free data retrieval call binding the contract method 0x66561d29.
//
// Solidity: function REG_CUSTOEMED_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNSession) REGCUSTOEMEDSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGCUSTOEMEDSUBM(&_BasOANN.CallOpts)
}

// REGCUSTOEMEDSUBM is a free data retrieval call binding the contract method 0x66561d29.
//
// Solidity: function REG_CUSTOEMED_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGCUSTOEMEDSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGCUSTOEMEDSUBM(&_BasOANN.CallOpts)
}

// REGCUSTOEMEDSUBO is a free data retrieval call binding the contract method 0xae8b1033.
//
// Solidity: function REG_CUSTOEMED_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGCUSTOEMEDSUBO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_CUSTOEMED_SUB_O")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGCUSTOEMEDSUBO is a free data retrieval call binding the contract method 0xae8b1033.
//
// Solidity: function REG_CUSTOEMED_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNSession) REGCUSTOEMEDSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGCUSTOEMEDSUBO(&_BasOANN.CallOpts)
}

// REGCUSTOEMEDSUBO is a free data retrieval call binding the contract method 0xae8b1033.
//
// Solidity: function REG_CUSTOEMED_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGCUSTOEMEDSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGCUSTOEMEDSUBO(&_BasOANN.CallOpts)
}

// REGNORMALSUBM is a free data retrieval call binding the contract method 0x90cd6925.
//
// Solidity: function REG_NORMAL_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGNORMALSUBM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_NORMAL_SUB_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGNORMALSUBM is a free data retrieval call binding the contract method 0x90cd6925.
//
// Solidity: function REG_NORMAL_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNSession) REGNORMALSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGNORMALSUBM(&_BasOANN.CallOpts)
}

// REGNORMALSUBM is a free data retrieval call binding the contract method 0x90cd6925.
//
// Solidity: function REG_NORMAL_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGNORMALSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGNORMALSUBM(&_BasOANN.CallOpts)
}

// REGNORMALSUBO is a free data retrieval call binding the contract method 0x28835235.
//
// Solidity: function REG_NORMAL_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGNORMALSUBO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_NORMAL_SUB_O")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGNORMALSUBO is a free data retrieval call binding the contract method 0x28835235.
//
// Solidity: function REG_NORMAL_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNSession) REGNORMALSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGNORMALSUBO(&_BasOANN.CallOpts)
}

// REGNORMALSUBO is a free data retrieval call binding the contract method 0x28835235.
//
// Solidity: function REG_NORMAL_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGNORMALSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGNORMALSUBO(&_BasOANN.CallOpts)
}

// REGROOTM is a free data retrieval call binding the contract method 0x4255284a.
//
// Solidity: function REG_ROOT_M() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGROOTM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_ROOT_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGROOTM is a free data retrieval call binding the contract method 0x4255284a.
//
// Solidity: function REG_ROOT_M() view returns(uint256)
func (_BasOANN *BasOANNSession) REGROOTM() (*big.Int, error) {
	return _BasOANN.Contract.REGROOTM(&_BasOANN.CallOpts)
}

// REGROOTM is a free data retrieval call binding the contract method 0x4255284a.
//
// Solidity: function REG_ROOT_M() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGROOTM() (*big.Int, error) {
	return _BasOANN.Contract.REGROOTM(&_BasOANN.CallOpts)
}

// REGSELFSUBM is a free data retrieval call binding the contract method 0x06b7d622.
//
// Solidity: function REG_SELF_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGSELFSUBM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_SELF_SUB_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGSELFSUBM is a free data retrieval call binding the contract method 0x06b7d622.
//
// Solidity: function REG_SELF_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNSession) REGSELFSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGSELFSUBM(&_BasOANN.CallOpts)
}

// REGSELFSUBM is a free data retrieval call binding the contract method 0x06b7d622.
//
// Solidity: function REG_SELF_SUB_M() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGSELFSUBM() (*big.Int, error) {
	return _BasOANN.Contract.REGSELFSUBM(&_BasOANN.CallOpts)
}

// REGSELFSUBO is a free data retrieval call binding the contract method 0x84806d00.
//
// Solidity: function REG_SELF_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCaller) REGSELFSUBO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "REG_SELF_SUB_O")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGSELFSUBO is a free data retrieval call binding the contract method 0x84806d00.
//
// Solidity: function REG_SELF_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNSession) REGSELFSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGSELFSUBO(&_BasOANN.CallOpts)
}

// REGSELFSUBO is a free data retrieval call binding the contract method 0x84806d00.
//
// Solidity: function REG_SELF_SUB_O() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) REGSELFSUBO() (*big.Int, error) {
	return _BasOANN.Contract.REGSELFSUBO(&_BasOANN.CallOpts)
}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() view returns(uint256)
func (_BasOANN *BasOANNCaller) SUBGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "SUB_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() view returns(uint256)
func (_BasOANN *BasOANNSession) SUBGAS() (*big.Int, error) {
	return _BasOANN.Contract.SUBGAS(&_BasOANN.CallOpts)
}

// SUBGAS is a free data retrieval call binding the contract method 0xfa1826bb.
//
// Solidity: function SUB_GAS() view returns(uint256)
func (_BasOANN *BasOANNCallerSession) SUBGAS() (*big.Int, error) {
	return _BasOANN.Contract.SUBGAS(&_BasOANN.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasOANN *BasOANNCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasOANN.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasOANN *BasOANNSession) Rel() (common.Address, error) {
	return _BasOANN.Contract.Rel(&_BasOANN.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasOANN *BasOANNCallerSession) Rel() (common.Address, error) {
	return _BasOANN.Contract.Rel(&_BasOANN.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasOANN *BasOANNTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasOANN *BasOANNSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeDAO(&_BasOANN.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasOANN *BasOANNTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeDAO(&_BasOANN.TransactOpts, newDao)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0x63d9691c.
//
// Solidity: function allocationSetting(uint256 reg_root_m, uint256 reg_self_sub_m, uint256 reg_self_sub_o, uint256 reg_normal_sub_m, uint256 reg_normal_sub_o, uint256 reg_custom_sub_m, uint256 reg_custom_sub_o) returns()
func (_BasOANN *BasOANNTransactor) AllocationSetting(opts *bind.TransactOpts, reg_root_m *big.Int, reg_self_sub_m *big.Int, reg_self_sub_o *big.Int, reg_normal_sub_m *big.Int, reg_normal_sub_o *big.Int, reg_custom_sub_m *big.Int, reg_custom_sub_o *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "allocationSetting", reg_root_m, reg_self_sub_m, reg_self_sub_o, reg_normal_sub_m, reg_normal_sub_o, reg_custom_sub_m, reg_custom_sub_o)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0x63d9691c.
//
// Solidity: function allocationSetting(uint256 reg_root_m, uint256 reg_self_sub_m, uint256 reg_self_sub_o, uint256 reg_normal_sub_m, uint256 reg_normal_sub_o, uint256 reg_custom_sub_m, uint256 reg_custom_sub_o) returns()
func (_BasOANN *BasOANNSession) AllocationSetting(reg_root_m *big.Int, reg_self_sub_m *big.Int, reg_self_sub_o *big.Int, reg_normal_sub_m *big.Int, reg_normal_sub_o *big.Int, reg_custom_sub_m *big.Int, reg_custom_sub_o *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.AllocationSetting(&_BasOANN.TransactOpts, reg_root_m, reg_self_sub_m, reg_self_sub_o, reg_normal_sub_m, reg_normal_sub_o, reg_custom_sub_m, reg_custom_sub_o)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0x63d9691c.
//
// Solidity: function allocationSetting(uint256 reg_root_m, uint256 reg_self_sub_m, uint256 reg_self_sub_o, uint256 reg_normal_sub_m, uint256 reg_normal_sub_o, uint256 reg_custom_sub_m, uint256 reg_custom_sub_o) returns()
func (_BasOANN *BasOANNTransactorSession) AllocationSetting(reg_root_m *big.Int, reg_self_sub_m *big.Int, reg_self_sub_o *big.Int, reg_normal_sub_m *big.Int, reg_normal_sub_o *big.Int, reg_custom_sub_m *big.Int, reg_custom_sub_o *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.AllocationSetting(&_BasOANN.TransactOpts, reg_root_m, reg_self_sub_m, reg_self_sub_o, reg_normal_sub_m, reg_normal_sub_o, reg_custom_sub_m, reg_custom_sub_o)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasOANN *BasOANNTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasOANN *BasOANNSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeRelation(&_BasOANN.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasOANN *BasOANNTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasOANN.Contract.ChangeRelation(&_BasOANN.TransactOpts, new_rel)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0xcdc78d36.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactor) OpenCustomPrice(opts *bind.TransactOpts, nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "openCustomPrice", nameHash, price)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0xcdc78d36.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNSession) OpenCustomPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomPrice(&_BasOANN.TransactOpts, nameHash, price)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0xcdc78d36.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 price) returns()
func (_BasOANN *BasOANNTransactorSession) OpenCustomPrice(nameHash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.OpenCustomPrice(&_BasOANN.TransactOpts, nameHash, price)
}

// PriceSetting is a paid mutator transaction binding the contract method 0xdc4b4714.
//
// Solidity: function priceSetting(uint256 max_year, uint256 aroot_gas, uint256 broot_gas, uint256 sub_gas, uint256 custom_price_gas) returns()
func (_BasOANN *BasOANNTransactor) PriceSetting(opts *bind.TransactOpts, max_year *big.Int, aroot_gas *big.Int, broot_gas *big.Int, sub_gas *big.Int, custom_price_gas *big.Int) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "priceSetting", max_year, aroot_gas, broot_gas, sub_gas, custom_price_gas)
}

// PriceSetting is a paid mutator transaction binding the contract method 0xdc4b4714.
//
// Solidity: function priceSetting(uint256 max_year, uint256 aroot_gas, uint256 broot_gas, uint256 sub_gas, uint256 custom_price_gas) returns()
func (_BasOANN *BasOANNSession) PriceSetting(max_year *big.Int, aroot_gas *big.Int, broot_gas *big.Int, sub_gas *big.Int, custom_price_gas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.PriceSetting(&_BasOANN.TransactOpts, max_year, aroot_gas, broot_gas, sub_gas, custom_price_gas)
}

// PriceSetting is a paid mutator transaction binding the contract method 0xdc4b4714.
//
// Solidity: function priceSetting(uint256 max_year, uint256 aroot_gas, uint256 broot_gas, uint256 sub_gas, uint256 custom_price_gas) returns()
func (_BasOANN *BasOANNTransactorSession) PriceSetting(max_year *big.Int, aroot_gas *big.Int, broot_gas *big.Int, sub_gas *big.Int, custom_price_gas *big.Int) (*types.Transaction, error) {
	return _BasOANN.Contract.PriceSetting(&_BasOANN.TransactOpts, max_year, aroot_gas, broot_gas, sub_gas, custom_price_gas)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RechargeRoot(opts *bind.TransactOpts, nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "rechargeRoot", nameHash, durationInYear)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RechargeRoot(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeRoot(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeRoot is a paid mutator transaction binding the contract method 0x79607e1f.
//
// Solidity: function rechargeRoot(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RechargeRoot(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeRoot(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RechargeSub(opts *bind.TransactOpts, nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "rechargeSub", nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RechargeSub(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeSub(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RechargeSub is a paid mutator transaction binding the contract method 0x6f6c774a.
//
// Solidity: function rechargeSub(bytes32 nameHash, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RechargeSub(nameHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RechargeSub(&_BasOANN.TransactOpts, nameHash, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustom, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RegisterRoot(opts *bind.TransactOpts, name []byte, isOpen bool, isCustom bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "registerRoot", name, isOpen, isCustom, cusPrice, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustom, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RegisterRoot(name []byte, isOpen bool, isCustom bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterRoot(&_BasOANN.TransactOpts, name, isOpen, isCustom, cusPrice, durationInYear)
}

// RegisterRoot is a paid mutator transaction binding the contract method 0x70a1495c.
//
// Solidity: function registerRoot(bytes name, bool isOpen, bool isCustom, uint256 cusPrice, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RegisterRoot(name []byte, isOpen bool, isCustom bool, cusPrice *big.Int, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterRoot(&_BasOANN.TransactOpts, name, isOpen, isCustom, cusPrice, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactor) RegisterSub(opts *bind.TransactOpts, rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.contract.Transact(opts, "registerSub", rName, sName, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNSession) RegisterSub(rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterSub(&_BasOANN.TransactOpts, rName, sName, durationInYear)
}

// RegisterSub is a paid mutator transaction binding the contract method 0xee9a512a.
//
// Solidity: function registerSub(bytes rName, bytes sName, uint8 durationInYear) returns()
func (_BasOANN *BasOANNTransactorSession) RegisterSub(rName []byte, sName []byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasOANN.Contract.RegisterSub(&_BasOANN.TransactOpts, rName, sName, durationInYear)
}

// BasOANNAllocationChangedIterator is returned from FilterAllocationChanged and is used to iterate over the raw logs and unpacked data for AllocationChanged events raised by the BasOANN contract.
type BasOANNAllocationChangedIterator struct {
	Event *BasOANNAllocationChanged // Event containing the contract specifics and raw log

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
func (it *BasOANNAllocationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOANNAllocationChanged)
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
		it.Event = new(BasOANNAllocationChanged)
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
func (it *BasOANNAllocationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOANNAllocationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOANNAllocationChanged represents a AllocationChanged event raised by the BasOANN contract.
type BasOANNAllocationChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllocationChanged is a free log retrieval operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasOANN *BasOANNFilterer) FilterAllocationChanged(opts *bind.FilterOpts) (*BasOANNAllocationChangedIterator, error) {

	logs, sub, err := _BasOANN.contract.FilterLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return &BasOANNAllocationChangedIterator{contract: _BasOANN.contract, event: "AllocationChanged", logs: logs, sub: sub}, nil
}

// WatchAllocationChanged is a free log subscription operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasOANN *BasOANNFilterer) WatchAllocationChanged(opts *bind.WatchOpts, sink chan<- *BasOANNAllocationChanged) (event.Subscription, error) {

	logs, sub, err := _BasOANN.contract.WatchLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOANNAllocationChanged)
				if err := _BasOANN.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
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

// ParseAllocationChanged is a log parse operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasOANN *BasOANNFilterer) ParseAllocationChanged(log types.Log) (*BasOANNAllocationChanged, error) {
	event := new(BasOANNAllocationChanged)
	if err := _BasOANN.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasOANNCostIterator is returned from FilterCost and is used to iterate over the raw logs and unpacked data for Cost events raised by the BasOANN contract.
type BasOANNCostIterator struct {
	Event *BasOANNCost // Event containing the contract specifics and raw log

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
func (it *BasOANNCostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOANNCost)
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
		it.Event = new(BasOANNCost)
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
func (it *BasOANNCostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOANNCostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOANNCost represents a Cost event raised by the BasOANN contract.
type BasOANNCost struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCost is a free log retrieval operation binding the contract event 0xfadea4513efa7fdb5c1017ae621bedcc55738e10a9dffbeff3c44776b7946acc.
//
// Solidity: event Cost(uint256 amount)
func (_BasOANN *BasOANNFilterer) FilterCost(opts *bind.FilterOpts) (*BasOANNCostIterator, error) {

	logs, sub, err := _BasOANN.contract.FilterLogs(opts, "Cost")
	if err != nil {
		return nil, err
	}
	return &BasOANNCostIterator{contract: _BasOANN.contract, event: "Cost", logs: logs, sub: sub}, nil
}

// WatchCost is a free log subscription operation binding the contract event 0xfadea4513efa7fdb5c1017ae621bedcc55738e10a9dffbeff3c44776b7946acc.
//
// Solidity: event Cost(uint256 amount)
func (_BasOANN *BasOANNFilterer) WatchCost(opts *bind.WatchOpts, sink chan<- *BasOANNCost) (event.Subscription, error) {

	logs, sub, err := _BasOANN.contract.WatchLogs(opts, "Cost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOANNCost)
				if err := _BasOANN.contract.UnpackLog(event, "Cost", log); err != nil {
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

// ParseCost is a log parse operation binding the contract event 0xfadea4513efa7fdb5c1017ae621bedcc55738e10a9dffbeff3c44776b7946acc.
//
// Solidity: event Cost(uint256 amount)
func (_BasOANN *BasOANNFilterer) ParseCost(log types.Log) (*BasOANNCost, error) {
	event := new(BasOANNCost)
	if err := _BasOANN.contract.UnpackLog(event, "Cost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasOANNPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the BasOANN contract.
type BasOANNPriceChangedIterator struct {
	Event *BasOANNPriceChanged // Event containing the contract specifics and raw log

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
func (it *BasOANNPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasOANNPriceChanged)
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
		it.Event = new(BasOANNPriceChanged)
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
func (it *BasOANNPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasOANNPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasOANNPriceChanged represents a PriceChanged event raised by the BasOANN contract.
type BasOANNPriceChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0x2160b3562aa2451b5ffd926270b68f67c99d1a3256bf7dd566c03788b0415b62.
//
// Solidity: event PriceChanged()
func (_BasOANN *BasOANNFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*BasOANNPriceChangedIterator, error) {

	logs, sub, err := _BasOANN.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &BasOANNPriceChangedIterator{contract: _BasOANN.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0x2160b3562aa2451b5ffd926270b68f67c99d1a3256bf7dd566c03788b0415b62.
//
// Solidity: event PriceChanged()
func (_BasOANN *BasOANNFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *BasOANNPriceChanged) (event.Subscription, error) {

	logs, sub, err := _BasOANN.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasOANNPriceChanged)
				if err := _BasOANN.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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

// ParsePriceChanged is a log parse operation binding the contract event 0x2160b3562aa2451b5ffd926270b68f67c99d1a3256bf7dd566c03788b0415b62.
//
// Solidity: event PriceChanged()
func (_BasOANN *BasOANNFilterer) ParsePriceChanged(log types.Log) (*BasOANNPriceChanged, error) {
	event := new(BasOANNPriceChanged)
	if err := _BasOANN.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
