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

// BasMailManagerMetaData contains all meta data concerning the BasMailManager contract.
var BasMailManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllocationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Cost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MailServerActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MailServerInactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"MailServerOpenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"UpdateConf\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIC_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MAIL_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPEN_ACTION_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_MAIL_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOP_DOMAIN_M\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOP_DOMAIN_O\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"basic_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"top_domain_m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"top_domain_o\",\"type\":\"uint256\"}],\"name\":\"allocationSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"}],\"name\":\"changeConfByDaoProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"cRootServe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cRootPub\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"subServe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"subPub\",\"type\":\"bool\"}],\"name\":\"changeSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mailConfigs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"}],\"name\":\"openMailService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_mail_year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open_action_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reg_mail_gas\",\"type\":\"uint256\"}],\"name\":\"priceSetting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"}],\"name\":\"registerMail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"}],\"name\":\"removeMailServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"}],\"name\":\"setPublic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasMailManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BasMailManagerMetaData.ABI instead.
var BasMailManagerABI = BasMailManagerMetaData.ABI

// BasMailManager is an auto generated Go binding around an Ethereum contract.
type BasMailManager struct {
	BasMailManagerCaller     // Read-only binding to the contract
	BasMailManagerTransactor // Write-only binding to the contract
	BasMailManagerFilterer   // Log filterer for contract events
}

// BasMailManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasMailManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasMailManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasMailManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasMailManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasMailManagerSession struct {
	Contract     *BasMailManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasMailManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasMailManagerCallerSession struct {
	Contract *BasMailManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BasMailManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasMailManagerTransactorSession struct {
	Contract     *BasMailManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BasMailManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasMailManagerRaw struct {
	Contract *BasMailManager // Generic contract binding to access the raw methods on
}

// BasMailManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasMailManagerCallerRaw struct {
	Contract *BasMailManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BasMailManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasMailManagerTransactorRaw struct {
	Contract *BasMailManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasMailManager creates a new instance of BasMailManager, bound to a specific deployed contract.
func NewBasMailManager(address common.Address, backend bind.ContractBackend) (*BasMailManager, error) {
	contract, err := bindBasMailManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasMailManager{BasMailManagerCaller: BasMailManagerCaller{contract: contract}, BasMailManagerTransactor: BasMailManagerTransactor{contract: contract}, BasMailManagerFilterer: BasMailManagerFilterer{contract: contract}}, nil
}

// NewBasMailManagerCaller creates a new read-only instance of BasMailManager, bound to a specific deployed contract.
func NewBasMailManagerCaller(address common.Address, caller bind.ContractCaller) (*BasMailManagerCaller, error) {
	contract, err := bindBasMailManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasMailManagerCaller{contract: contract}, nil
}

// NewBasMailManagerTransactor creates a new write-only instance of BasMailManager, bound to a specific deployed contract.
func NewBasMailManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BasMailManagerTransactor, error) {
	contract, err := bindBasMailManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasMailManagerTransactor{contract: contract}, nil
}

// NewBasMailManagerFilterer creates a new log filterer instance of BasMailManager, bound to a specific deployed contract.
func NewBasMailManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BasMailManagerFilterer, error) {
	contract, err := bindBasMailManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasMailManagerFilterer{contract: contract}, nil
}

// bindBasMailManager binds a generic wrapper to an already deployed contract.
func bindBasMailManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasMailManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMailManager *BasMailManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMailManager.Contract.BasMailManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMailManager *BasMailManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMailManager.Contract.BasMailManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMailManager *BasMailManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMailManager.Contract.BasMailManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasMailManager *BasMailManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasMailManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasMailManager *BasMailManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasMailManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasMailManager *BasMailManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasMailManager.Contract.contract.Transact(opts, method, params...)
}

// BASICM is a free data retrieval call binding the contract method 0x9cb155ac.
//
// Solidity: function BASIC_M() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) BASICM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "BASIC_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASICM is a free data retrieval call binding the contract method 0x9cb155ac.
//
// Solidity: function BASIC_M() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) BASICM() (*big.Int, error) {
	return _BasMailManager.Contract.BASICM(&_BasMailManager.CallOpts)
}

// BASICM is a free data retrieval call binding the contract method 0x9cb155ac.
//
// Solidity: function BASIC_M() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) BASICM() (*big.Int, error) {
	return _BasMailManager.Contract.BASICM(&_BasMailManager.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMailManager *BasMailManagerCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMailManager *BasMailManagerSession) DAOAddress() (common.Address, error) {
	return _BasMailManager.Contract.DAOAddress(&_BasMailManager.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasMailManager *BasMailManagerCallerSession) DAOAddress() (common.Address, error) {
	return _BasMailManager.Contract.DAOAddress(&_BasMailManager.CallOpts)
}

// MAXMAILYEAR is a free data retrieval call binding the contract method 0xe466682d.
//
// Solidity: function MAX_MAIL_YEAR() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) MAXMAILYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "MAX_MAIL_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMAILYEAR is a free data retrieval call binding the contract method 0xe466682d.
//
// Solidity: function MAX_MAIL_YEAR() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) MAXMAILYEAR() (*big.Int, error) {
	return _BasMailManager.Contract.MAXMAILYEAR(&_BasMailManager.CallOpts)
}

// MAXMAILYEAR is a free data retrieval call binding the contract method 0xe466682d.
//
// Solidity: function MAX_MAIL_YEAR() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) MAXMAILYEAR() (*big.Int, error) {
	return _BasMailManager.Contract.MAXMAILYEAR(&_BasMailManager.CallOpts)
}

// OPENACTIONGAS is a free data retrieval call binding the contract method 0xa9015a83.
//
// Solidity: function OPEN_ACTION_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) OPENACTIONGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "OPEN_ACTION_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPENACTIONGAS is a free data retrieval call binding the contract method 0xa9015a83.
//
// Solidity: function OPEN_ACTION_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) OPENACTIONGAS() (*big.Int, error) {
	return _BasMailManager.Contract.OPENACTIONGAS(&_BasMailManager.CallOpts)
}

// OPENACTIONGAS is a free data retrieval call binding the contract method 0xa9015a83.
//
// Solidity: function OPEN_ACTION_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) OPENACTIONGAS() (*big.Int, error) {
	return _BasMailManager.Contract.OPENACTIONGAS(&_BasMailManager.CallOpts)
}

// REGMAILGAS is a free data retrieval call binding the contract method 0x48f63694.
//
// Solidity: function REG_MAIL_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) REGMAILGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "REG_MAIL_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGMAILGAS is a free data retrieval call binding the contract method 0x48f63694.
//
// Solidity: function REG_MAIL_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) REGMAILGAS() (*big.Int, error) {
	return _BasMailManager.Contract.REGMAILGAS(&_BasMailManager.CallOpts)
}

// REGMAILGAS is a free data retrieval call binding the contract method 0x48f63694.
//
// Solidity: function REG_MAIL_GAS() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) REGMAILGAS() (*big.Int, error) {
	return _BasMailManager.Contract.REGMAILGAS(&_BasMailManager.CallOpts)
}

// TOPDOMAINM is a free data retrieval call binding the contract method 0x270563ab.
//
// Solidity: function TOP_DOMAIN_M() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) TOPDOMAINM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "TOP_DOMAIN_M")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOPDOMAINM is a free data retrieval call binding the contract method 0x270563ab.
//
// Solidity: function TOP_DOMAIN_M() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) TOPDOMAINM() (*big.Int, error) {
	return _BasMailManager.Contract.TOPDOMAINM(&_BasMailManager.CallOpts)
}

// TOPDOMAINM is a free data retrieval call binding the contract method 0x270563ab.
//
// Solidity: function TOP_DOMAIN_M() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) TOPDOMAINM() (*big.Int, error) {
	return _BasMailManager.Contract.TOPDOMAINM(&_BasMailManager.CallOpts)
}

// TOPDOMAINO is a free data retrieval call binding the contract method 0x6dacf424.
//
// Solidity: function TOP_DOMAIN_O() view returns(uint256)
func (_BasMailManager *BasMailManagerCaller) TOPDOMAINO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "TOP_DOMAIN_O")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOPDOMAINO is a free data retrieval call binding the contract method 0x6dacf424.
//
// Solidity: function TOP_DOMAIN_O() view returns(uint256)
func (_BasMailManager *BasMailManagerSession) TOPDOMAINO() (*big.Int, error) {
	return _BasMailManager.Contract.TOPDOMAINO(&_BasMailManager.CallOpts)
}

// TOPDOMAINO is a free data retrieval call binding the contract method 0x6dacf424.
//
// Solidity: function TOP_DOMAIN_O() view returns(uint256)
func (_BasMailManager *BasMailManagerCallerSession) TOPDOMAINO() (*big.Int, error) {
	return _BasMailManager.Contract.TOPDOMAINO(&_BasMailManager.CallOpts)
}

// MailConfigs is a free data retrieval call binding the contract method 0x3ce8d657.
//
// Solidity: function mailConfigs(bytes32 ) view returns(bool active, bool openToPublic)
func (_BasMailManager *BasMailManagerCaller) MailConfigs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Active       bool
	OpenToPublic bool
}, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "mailConfigs", arg0)

	outstruct := new(struct {
		Active       bool
		OpenToPublic bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.OpenToPublic = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// MailConfigs is a free data retrieval call binding the contract method 0x3ce8d657.
//
// Solidity: function mailConfigs(bytes32 ) view returns(bool active, bool openToPublic)
func (_BasMailManager *BasMailManagerSession) MailConfigs(arg0 [32]byte) (struct {
	Active       bool
	OpenToPublic bool
}, error) {
	return _BasMailManager.Contract.MailConfigs(&_BasMailManager.CallOpts, arg0)
}

// MailConfigs is a free data retrieval call binding the contract method 0x3ce8d657.
//
// Solidity: function mailConfigs(bytes32 ) view returns(bool active, bool openToPublic)
func (_BasMailManager *BasMailManagerCallerSession) MailConfigs(arg0 [32]byte) (struct {
	Active       bool
	OpenToPublic bool
}, error) {
	return _BasMailManager.Contract.MailConfigs(&_BasMailManager.CallOpts, arg0)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMailManager *BasMailManagerCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasMailManager.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMailManager *BasMailManagerSession) Rel() (common.Address, error) {
	return _BasMailManager.Contract.Rel(&_BasMailManager.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasMailManager *BasMailManagerCallerSession) Rel() (common.Address, error) {
	return _BasMailManager.Contract.Rel(&_BasMailManager.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMailManager *BasMailManagerTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMailManager *BasMailManagerSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeDAO(&_BasMailManager.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasMailManager *BasMailManagerTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeDAO(&_BasMailManager.TransactOpts, newDao)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0xe8a9c0ce.
//
// Solidity: function allocationSetting(uint256 basic_m, uint256 top_domain_m, uint256 top_domain_o) returns()
func (_BasMailManager *BasMailManagerTransactor) AllocationSetting(opts *bind.TransactOpts, basic_m *big.Int, top_domain_m *big.Int, top_domain_o *big.Int) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "allocationSetting", basic_m, top_domain_m, top_domain_o)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0xe8a9c0ce.
//
// Solidity: function allocationSetting(uint256 basic_m, uint256 top_domain_m, uint256 top_domain_o) returns()
func (_BasMailManager *BasMailManagerSession) AllocationSetting(basic_m *big.Int, top_domain_m *big.Int, top_domain_o *big.Int) (*types.Transaction, error) {
	return _BasMailManager.Contract.AllocationSetting(&_BasMailManager.TransactOpts, basic_m, top_domain_m, top_domain_o)
}

// AllocationSetting is a paid mutator transaction binding the contract method 0xe8a9c0ce.
//
// Solidity: function allocationSetting(uint256 basic_m, uint256 top_domain_m, uint256 top_domain_o) returns()
func (_BasMailManager *BasMailManagerTransactorSession) AllocationSetting(basic_m *big.Int, top_domain_m *big.Int, top_domain_o *big.Int) (*types.Transaction, error) {
	return _BasMailManager.Contract.AllocationSetting(&_BasMailManager.TransactOpts, basic_m, top_domain_m, top_domain_o)
}

// ChangeConfByDaoProposal is a paid mutator transaction binding the contract method 0xa42e7344.
//
// Solidity: function changeConfByDaoProposal(bytes32 nameHash, bool active, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactor) ChangeConfByDaoProposal(opts *bind.TransactOpts, nameHash [32]byte, active bool, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "changeConfByDaoProposal", nameHash, active, openToPublic)
}

// ChangeConfByDaoProposal is a paid mutator transaction binding the contract method 0xa42e7344.
//
// Solidity: function changeConfByDaoProposal(bytes32 nameHash, bool active, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerSession) ChangeConfByDaoProposal(nameHash [32]byte, active bool, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeConfByDaoProposal(&_BasMailManager.TransactOpts, nameHash, active, openToPublic)
}

// ChangeConfByDaoProposal is a paid mutator transaction binding the contract method 0xa42e7344.
//
// Solidity: function changeConfByDaoProposal(bytes32 nameHash, bool active, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactorSession) ChangeConfByDaoProposal(nameHash [32]byte, active bool, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeConfByDaoProposal(&_BasMailManager.TransactOpts, nameHash, active, openToPublic)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMailManager *BasMailManagerTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMailManager *BasMailManagerSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeRelation(&_BasMailManager.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasMailManager *BasMailManagerTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeRelation(&_BasMailManager.TransactOpts, new_rel)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x61546b55.
//
// Solidity: function changeSettings(bool cRootServe, bool cRootPub, bool subServe, bool subPub) returns()
func (_BasMailManager *BasMailManagerTransactor) ChangeSettings(opts *bind.TransactOpts, cRootServe bool, cRootPub bool, subServe bool, subPub bool) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "changeSettings", cRootServe, cRootPub, subServe, subPub)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x61546b55.
//
// Solidity: function changeSettings(bool cRootServe, bool cRootPub, bool subServe, bool subPub) returns()
func (_BasMailManager *BasMailManagerSession) ChangeSettings(cRootServe bool, cRootPub bool, subServe bool, subPub bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeSettings(&_BasMailManager.TransactOpts, cRootServe, cRootPub, subServe, subPub)
}

// ChangeSettings is a paid mutator transaction binding the contract method 0x61546b55.
//
// Solidity: function changeSettings(bool cRootServe, bool cRootPub, bool subServe, bool subPub) returns()
func (_BasMailManager *BasMailManagerTransactorSession) ChangeSettings(cRootServe bool, cRootPub bool, subServe bool, subPub bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.ChangeSettings(&_BasMailManager.TransactOpts, cRootServe, cRootPub, subServe, subPub)
}

// OpenMailService is a paid mutator transaction binding the contract method 0xc7b2eb2e.
//
// Solidity: function openMailService(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactor) OpenMailService(opts *bind.TransactOpts, domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "openMailService", domainHash, openToPublic)
}

// OpenMailService is a paid mutator transaction binding the contract method 0xc7b2eb2e.
//
// Solidity: function openMailService(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerSession) OpenMailService(domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.OpenMailService(&_BasMailManager.TransactOpts, domainHash, openToPublic)
}

// OpenMailService is a paid mutator transaction binding the contract method 0xc7b2eb2e.
//
// Solidity: function openMailService(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactorSession) OpenMailService(domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.OpenMailService(&_BasMailManager.TransactOpts, domainHash, openToPublic)
}

// PriceSetting is a paid mutator transaction binding the contract method 0x8c7d0076.
//
// Solidity: function priceSetting(uint256 max_mail_year, uint256 open_action_gas, uint256 reg_mail_gas) returns()
func (_BasMailManager *BasMailManagerTransactor) PriceSetting(opts *bind.TransactOpts, max_mail_year *big.Int, open_action_gas *big.Int, reg_mail_gas *big.Int) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "priceSetting", max_mail_year, open_action_gas, reg_mail_gas)
}

// PriceSetting is a paid mutator transaction binding the contract method 0x8c7d0076.
//
// Solidity: function priceSetting(uint256 max_mail_year, uint256 open_action_gas, uint256 reg_mail_gas) returns()
func (_BasMailManager *BasMailManagerSession) PriceSetting(max_mail_year *big.Int, open_action_gas *big.Int, reg_mail_gas *big.Int) (*types.Transaction, error) {
	return _BasMailManager.Contract.PriceSetting(&_BasMailManager.TransactOpts, max_mail_year, open_action_gas, reg_mail_gas)
}

// PriceSetting is a paid mutator transaction binding the contract method 0x8c7d0076.
//
// Solidity: function priceSetting(uint256 max_mail_year, uint256 open_action_gas, uint256 reg_mail_gas) returns()
func (_BasMailManager *BasMailManagerTransactorSession) PriceSetting(max_mail_year *big.Int, open_action_gas *big.Int, reg_mail_gas *big.Int) (*types.Transaction, error) {
	return _BasMailManager.Contract.PriceSetting(&_BasMailManager.TransactOpts, max_mail_year, open_action_gas, reg_mail_gas)
}

// Recharge is a paid mutator transaction binding the contract method 0x227b5f24.
//
// Solidity: function recharge(bytes32 mailHash, uint8 durationInYear) returns()
func (_BasMailManager *BasMailManagerTransactor) Recharge(opts *bind.TransactOpts, mailHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "recharge", mailHash, durationInYear)
}

// Recharge is a paid mutator transaction binding the contract method 0x227b5f24.
//
// Solidity: function recharge(bytes32 mailHash, uint8 durationInYear) returns()
func (_BasMailManager *BasMailManagerSession) Recharge(mailHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasMailManager.Contract.Recharge(&_BasMailManager.TransactOpts, mailHash, durationInYear)
}

// Recharge is a paid mutator transaction binding the contract method 0x227b5f24.
//
// Solidity: function recharge(bytes32 mailHash, uint8 durationInYear) returns()
func (_BasMailManager *BasMailManagerTransactorSession) Recharge(mailHash [32]byte, durationInYear uint8) (*types.Transaction, error) {
	return _BasMailManager.Contract.Recharge(&_BasMailManager.TransactOpts, mailHash, durationInYear)
}

// RegisterMail is a paid mutator transaction binding the contract method 0x63d67d85.
//
// Solidity: function registerMail(bytes32 domainHash, bytes32 mailHash, uint8 durationInYear, bytes aliasName) returns()
func (_BasMailManager *BasMailManagerTransactor) RegisterMail(opts *bind.TransactOpts, domainHash [32]byte, mailHash [32]byte, durationInYear uint8, aliasName []byte) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "registerMail", domainHash, mailHash, durationInYear, aliasName)
}

// RegisterMail is a paid mutator transaction binding the contract method 0x63d67d85.
//
// Solidity: function registerMail(bytes32 domainHash, bytes32 mailHash, uint8 durationInYear, bytes aliasName) returns()
func (_BasMailManager *BasMailManagerSession) RegisterMail(domainHash [32]byte, mailHash [32]byte, durationInYear uint8, aliasName []byte) (*types.Transaction, error) {
	return _BasMailManager.Contract.RegisterMail(&_BasMailManager.TransactOpts, domainHash, mailHash, durationInYear, aliasName)
}

// RegisterMail is a paid mutator transaction binding the contract method 0x63d67d85.
//
// Solidity: function registerMail(bytes32 domainHash, bytes32 mailHash, uint8 durationInYear, bytes aliasName) returns()
func (_BasMailManager *BasMailManagerTransactorSession) RegisterMail(domainHash [32]byte, mailHash [32]byte, durationInYear uint8, aliasName []byte) (*types.Transaction, error) {
	return _BasMailManager.Contract.RegisterMail(&_BasMailManager.TransactOpts, domainHash, mailHash, durationInYear, aliasName)
}

// RemoveMailServer is a paid mutator transaction binding the contract method 0xe7d3bd2a.
//
// Solidity: function removeMailServer(bytes32 domainHash) returns()
func (_BasMailManager *BasMailManagerTransactor) RemoveMailServer(opts *bind.TransactOpts, domainHash [32]byte) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "removeMailServer", domainHash)
}

// RemoveMailServer is a paid mutator transaction binding the contract method 0xe7d3bd2a.
//
// Solidity: function removeMailServer(bytes32 domainHash) returns()
func (_BasMailManager *BasMailManagerSession) RemoveMailServer(domainHash [32]byte) (*types.Transaction, error) {
	return _BasMailManager.Contract.RemoveMailServer(&_BasMailManager.TransactOpts, domainHash)
}

// RemoveMailServer is a paid mutator transaction binding the contract method 0xe7d3bd2a.
//
// Solidity: function removeMailServer(bytes32 domainHash) returns()
func (_BasMailManager *BasMailManagerTransactorSession) RemoveMailServer(domainHash [32]byte) (*types.Transaction, error) {
	return _BasMailManager.Contract.RemoveMailServer(&_BasMailManager.TransactOpts, domainHash)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactor) SetPublic(opts *bind.TransactOpts, domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.contract.Transact(opts, "setPublic", domainHash, openToPublic)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerSession) SetPublic(domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.SetPublic(&_BasMailManager.TransactOpts, domainHash, openToPublic)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 domainHash, bool openToPublic) returns()
func (_BasMailManager *BasMailManagerTransactorSession) SetPublic(domainHash [32]byte, openToPublic bool) (*types.Transaction, error) {
	return _BasMailManager.Contract.SetPublic(&_BasMailManager.TransactOpts, domainHash, openToPublic)
}

// BasMailManagerAllocationChangedIterator is returned from FilterAllocationChanged and is used to iterate over the raw logs and unpacked data for AllocationChanged events raised by the BasMailManager contract.
type BasMailManagerAllocationChangedIterator struct {
	Event *BasMailManagerAllocationChanged // Event containing the contract specifics and raw log

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
func (it *BasMailManagerAllocationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerAllocationChanged)
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
		it.Event = new(BasMailManagerAllocationChanged)
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
func (it *BasMailManagerAllocationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerAllocationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerAllocationChanged represents a AllocationChanged event raised by the BasMailManager contract.
type BasMailManagerAllocationChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllocationChanged is a free log retrieval operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasMailManager *BasMailManagerFilterer) FilterAllocationChanged(opts *bind.FilterOpts) (*BasMailManagerAllocationChangedIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerAllocationChangedIterator{contract: _BasMailManager.contract, event: "AllocationChanged", logs: logs, sub: sub}, nil
}

// WatchAllocationChanged is a free log subscription operation binding the contract event 0x315c884dd6c5858ebe0d5c16dc4eed2fc70df851669a1dbd221990202eaa733c.
//
// Solidity: event AllocationChanged()
func (_BasMailManager *BasMailManagerFilterer) WatchAllocationChanged(opts *bind.WatchOpts, sink chan<- *BasMailManagerAllocationChanged) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "AllocationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerAllocationChanged)
				if err := _BasMailManager.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
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
func (_BasMailManager *BasMailManagerFilterer) ParseAllocationChanged(log types.Log) (*BasMailManagerAllocationChanged, error) {
	event := new(BasMailManagerAllocationChanged)
	if err := _BasMailManager.contract.UnpackLog(event, "AllocationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerCostIterator is returned from FilterCost and is used to iterate over the raw logs and unpacked data for Cost events raised by the BasMailManager contract.
type BasMailManagerCostIterator struct {
	Event *BasMailManagerCost // Event containing the contract specifics and raw log

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
func (it *BasMailManagerCostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerCost)
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
		it.Event = new(BasMailManagerCost)
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
func (it *BasMailManagerCostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerCostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerCost represents a Cost event raised by the BasMailManager contract.
type BasMailManagerCost struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCost is a free log retrieval operation binding the contract event 0xfadea4513efa7fdb5c1017ae621bedcc55738e10a9dffbeff3c44776b7946acc.
//
// Solidity: event Cost(uint256 amount)
func (_BasMailManager *BasMailManagerFilterer) FilterCost(opts *bind.FilterOpts) (*BasMailManagerCostIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "Cost")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerCostIterator{contract: _BasMailManager.contract, event: "Cost", logs: logs, sub: sub}, nil
}

// WatchCost is a free log subscription operation binding the contract event 0xfadea4513efa7fdb5c1017ae621bedcc55738e10a9dffbeff3c44776b7946acc.
//
// Solidity: event Cost(uint256 amount)
func (_BasMailManager *BasMailManagerFilterer) WatchCost(opts *bind.WatchOpts, sink chan<- *BasMailManagerCost) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "Cost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerCost)
				if err := _BasMailManager.contract.UnpackLog(event, "Cost", log); err != nil {
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
func (_BasMailManager *BasMailManagerFilterer) ParseCost(log types.Log) (*BasMailManagerCost, error) {
	event := new(BasMailManagerCost)
	if err := _BasMailManager.contract.UnpackLog(event, "Cost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerMailServerActiveIterator is returned from FilterMailServerActive and is used to iterate over the raw logs and unpacked data for MailServerActive events raised by the BasMailManager contract.
type BasMailManagerMailServerActiveIterator struct {
	Event *BasMailManagerMailServerActive // Event containing the contract specifics and raw log

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
func (it *BasMailManagerMailServerActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerMailServerActive)
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
		it.Event = new(BasMailManagerMailServerActive)
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
func (it *BasMailManagerMailServerActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerMailServerActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerMailServerActive represents a MailServerActive event raised by the BasMailManager contract.
type BasMailManagerMailServerActive struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMailServerActive is a free log retrieval operation binding the contract event 0x34f72c0ba8132aa857d0d43acb46231a2147febdee86a394989c7eec8314b4ec.
//
// Solidity: event MailServerActive()
func (_BasMailManager *BasMailManagerFilterer) FilterMailServerActive(opts *bind.FilterOpts) (*BasMailManagerMailServerActiveIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "MailServerActive")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerMailServerActiveIterator{contract: _BasMailManager.contract, event: "MailServerActive", logs: logs, sub: sub}, nil
}

// WatchMailServerActive is a free log subscription operation binding the contract event 0x34f72c0ba8132aa857d0d43acb46231a2147febdee86a394989c7eec8314b4ec.
//
// Solidity: event MailServerActive()
func (_BasMailManager *BasMailManagerFilterer) WatchMailServerActive(opts *bind.WatchOpts, sink chan<- *BasMailManagerMailServerActive) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "MailServerActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerMailServerActive)
				if err := _BasMailManager.contract.UnpackLog(event, "MailServerActive", log); err != nil {
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

// ParseMailServerActive is a log parse operation binding the contract event 0x34f72c0ba8132aa857d0d43acb46231a2147febdee86a394989c7eec8314b4ec.
//
// Solidity: event MailServerActive()
func (_BasMailManager *BasMailManagerFilterer) ParseMailServerActive(log types.Log) (*BasMailManagerMailServerActive, error) {
	event := new(BasMailManagerMailServerActive)
	if err := _BasMailManager.contract.UnpackLog(event, "MailServerActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerMailServerInactiveIterator is returned from FilterMailServerInactive and is used to iterate over the raw logs and unpacked data for MailServerInactive events raised by the BasMailManager contract.
type BasMailManagerMailServerInactiveIterator struct {
	Event *BasMailManagerMailServerInactive // Event containing the contract specifics and raw log

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
func (it *BasMailManagerMailServerInactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerMailServerInactive)
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
		it.Event = new(BasMailManagerMailServerInactive)
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
func (it *BasMailManagerMailServerInactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerMailServerInactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerMailServerInactive represents a MailServerInactive event raised by the BasMailManager contract.
type BasMailManagerMailServerInactive struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMailServerInactive is a free log retrieval operation binding the contract event 0xa650c54c64f25bb1db2f1f1730a9ead305eb30c95bda607fdc1648eb3c948880.
//
// Solidity: event MailServerInactive()
func (_BasMailManager *BasMailManagerFilterer) FilterMailServerInactive(opts *bind.FilterOpts) (*BasMailManagerMailServerInactiveIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "MailServerInactive")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerMailServerInactiveIterator{contract: _BasMailManager.contract, event: "MailServerInactive", logs: logs, sub: sub}, nil
}

// WatchMailServerInactive is a free log subscription operation binding the contract event 0xa650c54c64f25bb1db2f1f1730a9ead305eb30c95bda607fdc1648eb3c948880.
//
// Solidity: event MailServerInactive()
func (_BasMailManager *BasMailManagerFilterer) WatchMailServerInactive(opts *bind.WatchOpts, sink chan<- *BasMailManagerMailServerInactive) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "MailServerInactive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerMailServerInactive)
				if err := _BasMailManager.contract.UnpackLog(event, "MailServerInactive", log); err != nil {
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

// ParseMailServerInactive is a log parse operation binding the contract event 0xa650c54c64f25bb1db2f1f1730a9ead305eb30c95bda607fdc1648eb3c948880.
//
// Solidity: event MailServerInactive()
func (_BasMailManager *BasMailManagerFilterer) ParseMailServerInactive(log types.Log) (*BasMailManagerMailServerInactive, error) {
	event := new(BasMailManagerMailServerInactive)
	if err := _BasMailManager.contract.UnpackLog(event, "MailServerInactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerMailServerOpenChangedIterator is returned from FilterMailServerOpenChanged and is used to iterate over the raw logs and unpacked data for MailServerOpenChanged events raised by the BasMailManager contract.
type BasMailManagerMailServerOpenChangedIterator struct {
	Event *BasMailManagerMailServerOpenChanged // Event containing the contract specifics and raw log

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
func (it *BasMailManagerMailServerOpenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerMailServerOpenChanged)
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
		it.Event = new(BasMailManagerMailServerOpenChanged)
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
func (it *BasMailManagerMailServerOpenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerMailServerOpenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerMailServerOpenChanged represents a MailServerOpenChanged event raised by the BasMailManager contract.
type BasMailManagerMailServerOpenChanged struct {
	DomainHash [32]byte
	IsOpen     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMailServerOpenChanged is a free log retrieval operation binding the contract event 0x1292a5f7457bce2e724f29ea382307fdd692ec671dd45ef1e746b50a52896322.
//
// Solidity: event MailServerOpenChanged(bytes32 domainHash, bool isOpen)
func (_BasMailManager *BasMailManagerFilterer) FilterMailServerOpenChanged(opts *bind.FilterOpts) (*BasMailManagerMailServerOpenChangedIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "MailServerOpenChanged")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerMailServerOpenChangedIterator{contract: _BasMailManager.contract, event: "MailServerOpenChanged", logs: logs, sub: sub}, nil
}

// WatchMailServerOpenChanged is a free log subscription operation binding the contract event 0x1292a5f7457bce2e724f29ea382307fdd692ec671dd45ef1e746b50a52896322.
//
// Solidity: event MailServerOpenChanged(bytes32 domainHash, bool isOpen)
func (_BasMailManager *BasMailManagerFilterer) WatchMailServerOpenChanged(opts *bind.WatchOpts, sink chan<- *BasMailManagerMailServerOpenChanged) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "MailServerOpenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerMailServerOpenChanged)
				if err := _BasMailManager.contract.UnpackLog(event, "MailServerOpenChanged", log); err != nil {
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

// ParseMailServerOpenChanged is a log parse operation binding the contract event 0x1292a5f7457bce2e724f29ea382307fdd692ec671dd45ef1e746b50a52896322.
//
// Solidity: event MailServerOpenChanged(bytes32 domainHash, bool isOpen)
func (_BasMailManager *BasMailManagerFilterer) ParseMailServerOpenChanged(log types.Log) (*BasMailManagerMailServerOpenChanged, error) {
	event := new(BasMailManagerMailServerOpenChanged)
	if err := _BasMailManager.contract.UnpackLog(event, "MailServerOpenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the BasMailManager contract.
type BasMailManagerPriceChangedIterator struct {
	Event *BasMailManagerPriceChanged // Event containing the contract specifics and raw log

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
func (it *BasMailManagerPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerPriceChanged)
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
		it.Event = new(BasMailManagerPriceChanged)
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
func (it *BasMailManagerPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerPriceChanged represents a PriceChanged event raised by the BasMailManager contract.
type BasMailManagerPriceChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0x2160b3562aa2451b5ffd926270b68f67c99d1a3256bf7dd566c03788b0415b62.
//
// Solidity: event PriceChanged()
func (_BasMailManager *BasMailManagerFilterer) FilterPriceChanged(opts *bind.FilterOpts) (*BasMailManagerPriceChangedIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerPriceChangedIterator{contract: _BasMailManager.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0x2160b3562aa2451b5ffd926270b68f67c99d1a3256bf7dd566c03788b0415b62.
//
// Solidity: event PriceChanged()
func (_BasMailManager *BasMailManagerFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *BasMailManagerPriceChanged) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "PriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerPriceChanged)
				if err := _BasMailManager.contract.UnpackLog(event, "PriceChanged", log); err != nil {
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
func (_BasMailManager *BasMailManagerFilterer) ParsePriceChanged(log types.Log) (*BasMailManagerPriceChanged, error) {
	event := new(BasMailManagerPriceChanged)
	if err := _BasMailManager.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasMailManagerUpdateConfIterator is returned from FilterUpdateConf and is used to iterate over the raw logs and unpacked data for UpdateConf events raised by the BasMailManager contract.
type BasMailManagerUpdateConfIterator struct {
	Event *BasMailManagerUpdateConf // Event containing the contract specifics and raw log

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
func (it *BasMailManagerUpdateConfIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasMailManagerUpdateConf)
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
		it.Event = new(BasMailManagerUpdateConf)
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
func (it *BasMailManagerUpdateConfIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasMailManagerUpdateConfIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasMailManagerUpdateConf represents a UpdateConf event raised by the BasMailManager contract.
type BasMailManagerUpdateConf struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateConf is a free log retrieval operation binding the contract event 0xce8b2fa5dec0c17ad624de493c8f9cd8c59c687f042f00e557f4b76b5dcf47e9.
//
// Solidity: event UpdateConf()
func (_BasMailManager *BasMailManagerFilterer) FilterUpdateConf(opts *bind.FilterOpts) (*BasMailManagerUpdateConfIterator, error) {

	logs, sub, err := _BasMailManager.contract.FilterLogs(opts, "UpdateConf")
	if err != nil {
		return nil, err
	}
	return &BasMailManagerUpdateConfIterator{contract: _BasMailManager.contract, event: "UpdateConf", logs: logs, sub: sub}, nil
}

// WatchUpdateConf is a free log subscription operation binding the contract event 0xce8b2fa5dec0c17ad624de493c8f9cd8c59c687f042f00e557f4b76b5dcf47e9.
//
// Solidity: event UpdateConf()
func (_BasMailManager *BasMailManagerFilterer) WatchUpdateConf(opts *bind.WatchOpts, sink chan<- *BasMailManagerUpdateConf) (event.Subscription, error) {

	logs, sub, err := _BasMailManager.contract.WatchLogs(opts, "UpdateConf")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasMailManagerUpdateConf)
				if err := _BasMailManager.contract.UnpackLog(event, "UpdateConf", log); err != nil {
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

// ParseUpdateConf is a log parse operation binding the contract event 0xce8b2fa5dec0c17ad624de493c8f9cd8c59c687f042f00e557f4b76b5dcf47e9.
//
// Solidity: event UpdateConf()
func (_BasMailManager *BasMailManagerFilterer) ParseUpdateConf(log types.Log) (*BasMailManagerUpdateConf, error) {
	event := new(BasMailManagerUpdateConf)
	if err := _BasMailManager.contract.UnpackLog(event, "UpdateConf", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
