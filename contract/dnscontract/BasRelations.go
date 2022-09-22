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

// BasRelationsMetaData contains all meta data concerning the BasRelations contract.
var BasRelationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pre\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SwitchOn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"addCanByModified\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canBeModified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelationAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"conf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecursor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mail\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mm\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oann\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc_addr\",\"type\":\"address\"}],\"name\":\"setAccAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conf_addr\",\"type\":\"address\"}],\"name\":\"setConfAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"exo_addr\",\"type\":\"address\"}],\"name\":\"setExoAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mail_addr\",\"type\":\"address\"}],\"name\":\"setMailAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mm_addr\",\"type\":\"address\"}],\"name\":\"setMailManagerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_addr\",\"type\":\"address\"}],\"name\":\"setMarketAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner_addr\",\"type\":\"address\"}],\"name\":\"setMinerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oann_addr\",\"type\":\"address\"}],\"name\":\"setOANNAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"root_addr\",\"type\":\"address\"}],\"name\":\"setRootAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sub_addr\",\"type\":\"address\"}],\"name\":\"setSubAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_addr\",\"type\":\"address\"}],\"name\":\"setTokenAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tro_addr\",\"type\":\"address\"}],\"name\":\"setTroAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tro\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BasRelationsABI is the input ABI used to generate the binding from.
// Deprecated: Use BasRelationsMetaData.ABI instead.
var BasRelationsABI = BasRelationsMetaData.ABI

// BasRelations is an auto generated Go binding around an Ethereum contract.
type BasRelations struct {
	BasRelationsCaller     // Read-only binding to the contract
	BasRelationsTransactor // Write-only binding to the contract
	BasRelationsFilterer   // Log filterer for contract events
}

// BasRelationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasRelationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRelationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasRelationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRelationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasRelationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRelationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasRelationsSession struct {
	Contract     *BasRelations     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasRelationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasRelationsCallerSession struct {
	Contract *BasRelationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BasRelationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasRelationsTransactorSession struct {
	Contract     *BasRelationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BasRelationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasRelationsRaw struct {
	Contract *BasRelations // Generic contract binding to access the raw methods on
}

// BasRelationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasRelationsCallerRaw struct {
	Contract *BasRelationsCaller // Generic read-only contract binding to access the raw methods on
}

// BasRelationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasRelationsTransactorRaw struct {
	Contract *BasRelationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasRelations creates a new instance of BasRelations, bound to a specific deployed contract.
func NewBasRelations(address common.Address, backend bind.ContractBackend) (*BasRelations, error) {
	contract, err := bindBasRelations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasRelations{BasRelationsCaller: BasRelationsCaller{contract: contract}, BasRelationsTransactor: BasRelationsTransactor{contract: contract}, BasRelationsFilterer: BasRelationsFilterer{contract: contract}}, nil
}

// NewBasRelationsCaller creates a new read-only instance of BasRelations, bound to a specific deployed contract.
func NewBasRelationsCaller(address common.Address, caller bind.ContractCaller) (*BasRelationsCaller, error) {
	contract, err := bindBasRelations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasRelationsCaller{contract: contract}, nil
}

// NewBasRelationsTransactor creates a new write-only instance of BasRelations, bound to a specific deployed contract.
func NewBasRelationsTransactor(address common.Address, transactor bind.ContractTransactor) (*BasRelationsTransactor, error) {
	contract, err := bindBasRelations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasRelationsTransactor{contract: contract}, nil
}

// NewBasRelationsFilterer creates a new log filterer instance of BasRelations, bound to a specific deployed contract.
func NewBasRelationsFilterer(address common.Address, filterer bind.ContractFilterer) (*BasRelationsFilterer, error) {
	contract, err := bindBasRelations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasRelationsFilterer{contract: contract}, nil
}

// bindBasRelations binds a generic wrapper to an already deployed contract.
func bindBasRelations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasRelationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRelations *BasRelationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasRelations.Contract.BasRelationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRelations *BasRelationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRelations.Contract.BasRelationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRelations *BasRelationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRelations.Contract.BasRelationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRelations *BasRelationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasRelations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRelations *BasRelationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRelations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRelations *BasRelationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRelations.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRelations *BasRelationsCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRelations *BasRelationsSession) DAOAddress() (common.Address, error) {
	return _BasRelations.Contract.DAOAddress(&_BasRelations.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRelations *BasRelationsCallerSession) DAOAddress() (common.Address, error) {
	return _BasRelations.Contract.DAOAddress(&_BasRelations.CallOpts)
}

// Acc is a free data retrieval call binding the contract method 0x7da1365e.
//
// Solidity: function acc() view returns(address)
func (_BasRelations *BasRelationsCaller) Acc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "acc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acc is a free data retrieval call binding the contract method 0x7da1365e.
//
// Solidity: function acc() view returns(address)
func (_BasRelations *BasRelationsSession) Acc() (common.Address, error) {
	return _BasRelations.Contract.Acc(&_BasRelations.CallOpts)
}

// Acc is a free data retrieval call binding the contract method 0x7da1365e.
//
// Solidity: function acc() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Acc() (common.Address, error) {
	return _BasRelations.Contract.Acc(&_BasRelations.CallOpts)
}

// CanBeModified is a free data retrieval call binding the contract method 0xad9ecd0f.
//
// Solidity: function canBeModified(address , address ) view returns(bool)
func (_BasRelations *BasRelationsCaller) CanBeModified(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "canBeModified", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBeModified is a free data retrieval call binding the contract method 0xad9ecd0f.
//
// Solidity: function canBeModified(address , address ) view returns(bool)
func (_BasRelations *BasRelationsSession) CanBeModified(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _BasRelations.Contract.CanBeModified(&_BasRelations.CallOpts, arg0, arg1)
}

// CanBeModified is a free data retrieval call binding the contract method 0xad9ecd0f.
//
// Solidity: function canBeModified(address , address ) view returns(bool)
func (_BasRelations *BasRelationsCallerSession) CanBeModified(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _BasRelations.Contract.CanBeModified(&_BasRelations.CallOpts, arg0, arg1)
}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(address)
func (_BasRelations *BasRelationsCaller) Conf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "conf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(address)
func (_BasRelations *BasRelationsSession) Conf() (common.Address, error) {
	return _BasRelations.Contract.Conf(&_BasRelations.CallOpts)
}

// Conf is a free data retrieval call binding the contract method 0x17792729.
//
// Solidity: function conf() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Conf() (common.Address, error) {
	return _BasRelations.Contract.Conf(&_BasRelations.CallOpts)
}

// Exo is a free data retrieval call binding the contract method 0xa53a9062.
//
// Solidity: function exo() view returns(address)
func (_BasRelations *BasRelationsCaller) Exo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "exo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exo is a free data retrieval call binding the contract method 0xa53a9062.
//
// Solidity: function exo() view returns(address)
func (_BasRelations *BasRelationsSession) Exo() (common.Address, error) {
	return _BasRelations.Contract.Exo(&_BasRelations.CallOpts)
}

// Exo is a free data retrieval call binding the contract method 0xa53a9062.
//
// Solidity: function exo() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Exo() (common.Address, error) {
	return _BasRelations.Contract.Exo(&_BasRelations.CallOpts)
}

// GetPrecursor is a free data retrieval call binding the contract method 0x5c840a4c.
//
// Solidity: function getPrecursor() view returns(address)
func (_BasRelations *BasRelationsCaller) GetPrecursor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "getPrecursor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPrecursor is a free data retrieval call binding the contract method 0x5c840a4c.
//
// Solidity: function getPrecursor() view returns(address)
func (_BasRelations *BasRelationsSession) GetPrecursor() (common.Address, error) {
	return _BasRelations.Contract.GetPrecursor(&_BasRelations.CallOpts)
}

// GetPrecursor is a free data retrieval call binding the contract method 0x5c840a4c.
//
// Solidity: function getPrecursor() view returns(address)
func (_BasRelations *BasRelationsCallerSession) GetPrecursor() (common.Address, error) {
	return _BasRelations.Contract.GetPrecursor(&_BasRelations.CallOpts)
}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_BasRelations *BasRelationsCaller) Mail(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "mail")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_BasRelations *BasRelationsSession) Mail() (common.Address, error) {
	return _BasRelations.Contract.Mail(&_BasRelations.CallOpts)
}

// Mail is a free data retrieval call binding the contract method 0x9a68308f.
//
// Solidity: function mail() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Mail() (common.Address, error) {
	return _BasRelations.Contract.Mail(&_BasRelations.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_BasRelations *BasRelationsCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_BasRelations *BasRelationsSession) Market() (common.Address, error) {
	return _BasRelations.Contract.Market(&_BasRelations.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Market() (common.Address, error) {
	return _BasRelations.Contract.Market(&_BasRelations.CallOpts)
}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_BasRelations *BasRelationsCaller) Miner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "miner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_BasRelations *BasRelationsSession) Miner() (common.Address, error) {
	return _BasRelations.Contract.Miner(&_BasRelations.CallOpts)
}

// Miner is a free data retrieval call binding the contract method 0x349dc329.
//
// Solidity: function miner() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Miner() (common.Address, error) {
	return _BasRelations.Contract.Miner(&_BasRelations.CallOpts)
}

// Mm is a free data retrieval call binding the contract method 0xe90eaa51.
//
// Solidity: function mm() view returns(address)
func (_BasRelations *BasRelationsCaller) Mm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "mm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mm is a free data retrieval call binding the contract method 0xe90eaa51.
//
// Solidity: function mm() view returns(address)
func (_BasRelations *BasRelationsSession) Mm() (common.Address, error) {
	return _BasRelations.Contract.Mm(&_BasRelations.CallOpts)
}

// Mm is a free data retrieval call binding the contract method 0xe90eaa51.
//
// Solidity: function mm() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Mm() (common.Address, error) {
	return _BasRelations.Contract.Mm(&_BasRelations.CallOpts)
}

// Oann is a free data retrieval call binding the contract method 0xc8815d07.
//
// Solidity: function oann() view returns(address)
func (_BasRelations *BasRelationsCaller) Oann(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "oann")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oann is a free data retrieval call binding the contract method 0xc8815d07.
//
// Solidity: function oann() view returns(address)
func (_BasRelations *BasRelationsSession) Oann() (common.Address, error) {
	return _BasRelations.Contract.Oann(&_BasRelations.CallOpts)
}

// Oann is a free data retrieval call binding the contract method 0xc8815d07.
//
// Solidity: function oann() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Oann() (common.Address, error) {
	return _BasRelations.Contract.Oann(&_BasRelations.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BasRelations *BasRelationsCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BasRelations *BasRelationsSession) Root() (common.Address, error) {
	return _BasRelations.Contract.Root(&_BasRelations.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Root() (common.Address, error) {
	return _BasRelations.Contract.Root(&_BasRelations.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() view returns(address)
func (_BasRelations *BasRelationsCaller) Sub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "sub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() view returns(address)
func (_BasRelations *BasRelationsSession) Sub() (common.Address, error) {
	return _BasRelations.Contract.Sub(&_BasRelations.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xc54124be.
//
// Solidity: function sub() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Sub() (common.Address, error) {
	return _BasRelations.Contract.Sub(&_BasRelations.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BasRelations *BasRelationsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BasRelations *BasRelationsSession) Token() (common.Address, error) {
	return _BasRelations.Contract.Token(&_BasRelations.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Token() (common.Address, error) {
	return _BasRelations.Contract.Token(&_BasRelations.CallOpts)
}

// Tro is a free data retrieval call binding the contract method 0x3e0b59e7.
//
// Solidity: function tro() view returns(address)
func (_BasRelations *BasRelationsCaller) Tro(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRelations.contract.Call(opts, &out, "tro")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tro is a free data retrieval call binding the contract method 0x3e0b59e7.
//
// Solidity: function tro() view returns(address)
func (_BasRelations *BasRelationsSession) Tro() (common.Address, error) {
	return _BasRelations.Contract.Tro(&_BasRelations.CallOpts)
}

// Tro is a free data retrieval call binding the contract method 0x3e0b59e7.
//
// Solidity: function tro() view returns(address)
func (_BasRelations *BasRelationsCallerSession) Tro() (common.Address, error) {
	return _BasRelations.Contract.Tro(&_BasRelations.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRelations *BasRelationsTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRelations *BasRelationsSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.ChangeDAO(&_BasRelations.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRelations *BasRelationsTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.ChangeDAO(&_BasRelations.TransactOpts, newDao)
}

// SwitchOn is a paid mutator transaction binding the contract method 0x2bd4c8ad.
//
// Solidity: function SwitchOn() returns()
func (_BasRelations *BasRelationsTransactor) SwitchOn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "SwitchOn")
}

// SwitchOn is a paid mutator transaction binding the contract method 0x2bd4c8ad.
//
// Solidity: function SwitchOn() returns()
func (_BasRelations *BasRelationsSession) SwitchOn() (*types.Transaction, error) {
	return _BasRelations.Contract.SwitchOn(&_BasRelations.TransactOpts)
}

// SwitchOn is a paid mutator transaction binding the contract method 0x2bd4c8ad.
//
// Solidity: function SwitchOn() returns()
func (_BasRelations *BasRelationsTransactorSession) SwitchOn() (*types.Transaction, error) {
	return _BasRelations.Contract.SwitchOn(&_BasRelations.TransactOpts)
}

// AddCanByModified is a paid mutator transaction binding the contract method 0xce018134.
//
// Solidity: function addCanByModified(address data, address caller) returns()
func (_BasRelations *BasRelationsTransactor) AddCanByModified(opts *bind.TransactOpts, data common.Address, caller common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "addCanByModified", data, caller)
}

// AddCanByModified is a paid mutator transaction binding the contract method 0xce018134.
//
// Solidity: function addCanByModified(address data, address caller) returns()
func (_BasRelations *BasRelationsSession) AddCanByModified(data common.Address, caller common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.AddCanByModified(&_BasRelations.TransactOpts, data, caller)
}

// AddCanByModified is a paid mutator transaction binding the contract method 0xce018134.
//
// Solidity: function addCanByModified(address data, address caller) returns()
func (_BasRelations *BasRelationsTransactorSession) AddCanByModified(data common.Address, caller common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.AddCanByModified(&_BasRelations.TransactOpts, data, caller)
}

// ChangeRelationAll is a paid mutator transaction binding the contract method 0x41ce484a.
//
// Solidity: function changeRelationAll(address new_rel) returns()
func (_BasRelations *BasRelationsTransactor) ChangeRelationAll(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "changeRelationAll", new_rel)
}

// ChangeRelationAll is a paid mutator transaction binding the contract method 0x41ce484a.
//
// Solidity: function changeRelationAll(address new_rel) returns()
func (_BasRelations *BasRelationsSession) ChangeRelationAll(new_rel common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.ChangeRelationAll(&_BasRelations.TransactOpts, new_rel)
}

// ChangeRelationAll is a paid mutator transaction binding the contract method 0x41ce484a.
//
// Solidity: function changeRelationAll(address new_rel) returns()
func (_BasRelations *BasRelationsTransactorSession) ChangeRelationAll(new_rel common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.ChangeRelationAll(&_BasRelations.TransactOpts, new_rel)
}

// SetAccAddr is a paid mutator transaction binding the contract method 0x6efa26a2.
//
// Solidity: function setAccAddr(address acc_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetAccAddr(opts *bind.TransactOpts, acc_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setAccAddr", acc_addr)
}

// SetAccAddr is a paid mutator transaction binding the contract method 0x6efa26a2.
//
// Solidity: function setAccAddr(address acc_addr) returns()
func (_BasRelations *BasRelationsSession) SetAccAddr(acc_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetAccAddr(&_BasRelations.TransactOpts, acc_addr)
}

// SetAccAddr is a paid mutator transaction binding the contract method 0x6efa26a2.
//
// Solidity: function setAccAddr(address acc_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetAccAddr(acc_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetAccAddr(&_BasRelations.TransactOpts, acc_addr)
}

// SetConfAddr is a paid mutator transaction binding the contract method 0x1606d1de.
//
// Solidity: function setConfAddr(address conf_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetConfAddr(opts *bind.TransactOpts, conf_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setConfAddr", conf_addr)
}

// SetConfAddr is a paid mutator transaction binding the contract method 0x1606d1de.
//
// Solidity: function setConfAddr(address conf_addr) returns()
func (_BasRelations *BasRelationsSession) SetConfAddr(conf_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetConfAddr(&_BasRelations.TransactOpts, conf_addr)
}

// SetConfAddr is a paid mutator transaction binding the contract method 0x1606d1de.
//
// Solidity: function setConfAddr(address conf_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetConfAddr(conf_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetConfAddr(&_BasRelations.TransactOpts, conf_addr)
}

// SetExoAddr is a paid mutator transaction binding the contract method 0x8fc04abc.
//
// Solidity: function setExoAddr(address exo_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetExoAddr(opts *bind.TransactOpts, exo_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setExoAddr", exo_addr)
}

// SetExoAddr is a paid mutator transaction binding the contract method 0x8fc04abc.
//
// Solidity: function setExoAddr(address exo_addr) returns()
func (_BasRelations *BasRelationsSession) SetExoAddr(exo_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetExoAddr(&_BasRelations.TransactOpts, exo_addr)
}

// SetExoAddr is a paid mutator transaction binding the contract method 0x8fc04abc.
//
// Solidity: function setExoAddr(address exo_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetExoAddr(exo_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetExoAddr(&_BasRelations.TransactOpts, exo_addr)
}

// SetMailAddr is a paid mutator transaction binding the contract method 0x17020d68.
//
// Solidity: function setMailAddr(address mail_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetMailAddr(opts *bind.TransactOpts, mail_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setMailAddr", mail_addr)
}

// SetMailAddr is a paid mutator transaction binding the contract method 0x17020d68.
//
// Solidity: function setMailAddr(address mail_addr) returns()
func (_BasRelations *BasRelationsSession) SetMailAddr(mail_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMailAddr(&_BasRelations.TransactOpts, mail_addr)
}

// SetMailAddr is a paid mutator transaction binding the contract method 0x17020d68.
//
// Solidity: function setMailAddr(address mail_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetMailAddr(mail_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMailAddr(&_BasRelations.TransactOpts, mail_addr)
}

// SetMailManagerAddr is a paid mutator transaction binding the contract method 0x903c002d.
//
// Solidity: function setMailManagerAddr(address mm_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetMailManagerAddr(opts *bind.TransactOpts, mm_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setMailManagerAddr", mm_addr)
}

// SetMailManagerAddr is a paid mutator transaction binding the contract method 0x903c002d.
//
// Solidity: function setMailManagerAddr(address mm_addr) returns()
func (_BasRelations *BasRelationsSession) SetMailManagerAddr(mm_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMailManagerAddr(&_BasRelations.TransactOpts, mm_addr)
}

// SetMailManagerAddr is a paid mutator transaction binding the contract method 0x903c002d.
//
// Solidity: function setMailManagerAddr(address mm_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetMailManagerAddr(mm_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMailManagerAddr(&_BasRelations.TransactOpts, mm_addr)
}

// SetMarketAddr is a paid mutator transaction binding the contract method 0x51d31822.
//
// Solidity: function setMarketAddr(address market_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetMarketAddr(opts *bind.TransactOpts, market_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setMarketAddr", market_addr)
}

// SetMarketAddr is a paid mutator transaction binding the contract method 0x51d31822.
//
// Solidity: function setMarketAddr(address market_addr) returns()
func (_BasRelations *BasRelationsSession) SetMarketAddr(market_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMarketAddr(&_BasRelations.TransactOpts, market_addr)
}

// SetMarketAddr is a paid mutator transaction binding the contract method 0x51d31822.
//
// Solidity: function setMarketAddr(address market_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetMarketAddr(market_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMarketAddr(&_BasRelations.TransactOpts, market_addr)
}

// SetMinerAddr is a paid mutator transaction binding the contract method 0xc3859a90.
//
// Solidity: function setMinerAddr(address miner_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetMinerAddr(opts *bind.TransactOpts, miner_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setMinerAddr", miner_addr)
}

// SetMinerAddr is a paid mutator transaction binding the contract method 0xc3859a90.
//
// Solidity: function setMinerAddr(address miner_addr) returns()
func (_BasRelations *BasRelationsSession) SetMinerAddr(miner_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMinerAddr(&_BasRelations.TransactOpts, miner_addr)
}

// SetMinerAddr is a paid mutator transaction binding the contract method 0xc3859a90.
//
// Solidity: function setMinerAddr(address miner_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetMinerAddr(miner_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetMinerAddr(&_BasRelations.TransactOpts, miner_addr)
}

// SetOANNAddr is a paid mutator transaction binding the contract method 0x335e1f4a.
//
// Solidity: function setOANNAddr(address oann_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetOANNAddr(opts *bind.TransactOpts, oann_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setOANNAddr", oann_addr)
}

// SetOANNAddr is a paid mutator transaction binding the contract method 0x335e1f4a.
//
// Solidity: function setOANNAddr(address oann_addr) returns()
func (_BasRelations *BasRelationsSession) SetOANNAddr(oann_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetOANNAddr(&_BasRelations.TransactOpts, oann_addr)
}

// SetOANNAddr is a paid mutator transaction binding the contract method 0x335e1f4a.
//
// Solidity: function setOANNAddr(address oann_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetOANNAddr(oann_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetOANNAddr(&_BasRelations.TransactOpts, oann_addr)
}

// SetRootAddr is a paid mutator transaction binding the contract method 0x202089b3.
//
// Solidity: function setRootAddr(address root_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetRootAddr(opts *bind.TransactOpts, root_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setRootAddr", root_addr)
}

// SetRootAddr is a paid mutator transaction binding the contract method 0x202089b3.
//
// Solidity: function setRootAddr(address root_addr) returns()
func (_BasRelations *BasRelationsSession) SetRootAddr(root_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetRootAddr(&_BasRelations.TransactOpts, root_addr)
}

// SetRootAddr is a paid mutator transaction binding the contract method 0x202089b3.
//
// Solidity: function setRootAddr(address root_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetRootAddr(root_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetRootAddr(&_BasRelations.TransactOpts, root_addr)
}

// SetSubAddr is a paid mutator transaction binding the contract method 0xbcddeffe.
//
// Solidity: function setSubAddr(address sub_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetSubAddr(opts *bind.TransactOpts, sub_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setSubAddr", sub_addr)
}

// SetSubAddr is a paid mutator transaction binding the contract method 0xbcddeffe.
//
// Solidity: function setSubAddr(address sub_addr) returns()
func (_BasRelations *BasRelationsSession) SetSubAddr(sub_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetSubAddr(&_BasRelations.TransactOpts, sub_addr)
}

// SetSubAddr is a paid mutator transaction binding the contract method 0xbcddeffe.
//
// Solidity: function setSubAddr(address sub_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetSubAddr(sub_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetSubAddr(&_BasRelations.TransactOpts, sub_addr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x2ebd1e28.
//
// Solidity: function setTokenAddr(address token_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetTokenAddr(opts *bind.TransactOpts, token_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setTokenAddr", token_addr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x2ebd1e28.
//
// Solidity: function setTokenAddr(address token_addr) returns()
func (_BasRelations *BasRelationsSession) SetTokenAddr(token_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetTokenAddr(&_BasRelations.TransactOpts, token_addr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x2ebd1e28.
//
// Solidity: function setTokenAddr(address token_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetTokenAddr(token_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetTokenAddr(&_BasRelations.TransactOpts, token_addr)
}

// SetTroAddr is a paid mutator transaction binding the contract method 0x77ae9298.
//
// Solidity: function setTroAddr(address tro_addr) returns()
func (_BasRelations *BasRelationsTransactor) SetTroAddr(opts *bind.TransactOpts, tro_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.contract.Transact(opts, "setTroAddr", tro_addr)
}

// SetTroAddr is a paid mutator transaction binding the contract method 0x77ae9298.
//
// Solidity: function setTroAddr(address tro_addr) returns()
func (_BasRelations *BasRelationsSession) SetTroAddr(tro_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetTroAddr(&_BasRelations.TransactOpts, tro_addr)
}

// SetTroAddr is a paid mutator transaction binding the contract method 0x77ae9298.
//
// Solidity: function setTroAddr(address tro_addr) returns()
func (_BasRelations *BasRelationsTransactorSession) SetTroAddr(tro_addr common.Address) (*types.Transaction, error) {
	return _BasRelations.Contract.SetTroAddr(&_BasRelations.TransactOpts, tro_addr)
}
