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

// BasSubDomainMetaData contains all meta data concerning the BasSubDomain contract.
var BasSubDomainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"NewSubDomain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"SubDataReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"SubUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Sub\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"totalName\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNameByHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"hasDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rechargeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEnd\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"subdomain\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"newExpire\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"applicant\",\"type\":\"address\"}],\"name\":\"replaceOrCreate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"verifySub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BasSubDomainABI is the input ABI used to generate the binding from.
// Deprecated: Use BasSubDomainMetaData.ABI instead.
var BasSubDomainABI = BasSubDomainMetaData.ABI

// BasSubDomain is an auto generated Go binding around an Ethereum contract.
type BasSubDomain struct {
	BasSubDomainCaller     // Read-only binding to the contract
	BasSubDomainTransactor // Write-only binding to the contract
	BasSubDomainFilterer   // Log filterer for contract events
}

// BasSubDomainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasSubDomainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSubDomainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasSubDomainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSubDomainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasSubDomainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasSubDomainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasSubDomainSession struct {
	Contract     *BasSubDomain     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasSubDomainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasSubDomainCallerSession struct {
	Contract *BasSubDomainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BasSubDomainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasSubDomainTransactorSession struct {
	Contract     *BasSubDomainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BasSubDomainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasSubDomainRaw struct {
	Contract *BasSubDomain // Generic contract binding to access the raw methods on
}

// BasSubDomainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasSubDomainCallerRaw struct {
	Contract *BasSubDomainCaller // Generic read-only contract binding to access the raw methods on
}

// BasSubDomainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasSubDomainTransactorRaw struct {
	Contract *BasSubDomainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasSubDomain creates a new instance of BasSubDomain, bound to a specific deployed contract.
func NewBasSubDomain(address common.Address, backend bind.ContractBackend) (*BasSubDomain, error) {
	contract, err := bindBasSubDomain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasSubDomain{BasSubDomainCaller: BasSubDomainCaller{contract: contract}, BasSubDomainTransactor: BasSubDomainTransactor{contract: contract}, BasSubDomainFilterer: BasSubDomainFilterer{contract: contract}}, nil
}

// NewBasSubDomainCaller creates a new read-only instance of BasSubDomain, bound to a specific deployed contract.
func NewBasSubDomainCaller(address common.Address, caller bind.ContractCaller) (*BasSubDomainCaller, error) {
	contract, err := bindBasSubDomain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasSubDomainCaller{contract: contract}, nil
}

// NewBasSubDomainTransactor creates a new write-only instance of BasSubDomain, bound to a specific deployed contract.
func NewBasSubDomainTransactor(address common.Address, transactor bind.ContractTransactor) (*BasSubDomainTransactor, error) {
	contract, err := bindBasSubDomain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasSubDomainTransactor{contract: contract}, nil
}

// NewBasSubDomainFilterer creates a new log filterer instance of BasSubDomain, bound to a specific deployed contract.
func NewBasSubDomainFilterer(address common.Address, filterer bind.ContractFilterer) (*BasSubDomainFilterer, error) {
	contract, err := bindBasSubDomain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasSubDomainFilterer{contract: contract}, nil
}

// bindBasSubDomain binds a generic wrapper to an already deployed contract.
func bindBasSubDomain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasSubDomainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasSubDomain *BasSubDomainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasSubDomain.Contract.BasSubDomainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasSubDomain *BasSubDomainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasSubDomain.Contract.BasSubDomainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasSubDomain *BasSubDomainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasSubDomain.Contract.BasSubDomainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasSubDomain *BasSubDomainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasSubDomain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasSubDomain *BasSubDomainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasSubDomain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasSubDomain *BasSubDomainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasSubDomain.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasSubDomain *BasSubDomainCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasSubDomain *BasSubDomainSession) DAOAddress() (common.Address, error) {
	return _BasSubDomain.Contract.DAOAddress(&_BasSubDomain.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasSubDomain *BasSubDomainCallerSession) DAOAddress() (common.Address, error) {
	return _BasSubDomain.Contract.DAOAddress(&_BasSubDomain.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) view returns(bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainCaller) Sub(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "Sub", arg0)

	outstruct := new(struct {
		TotalName []byte
		RootHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalName = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.RootHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) view returns(bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainSession) Sub(arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	return _BasSubDomain.Contract.Sub(&_BasSubDomain.CallOpts, arg0)
}

// Sub is a free data retrieval call binding the contract method 0x333cf44c.
//
// Solidity: function Sub(bytes32 ) view returns(bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainCallerSession) Sub(arg0 [32]byte) (struct {
	TotalName []byte
	RootHash  [32]byte
}, error) {
	return _BasSubDomain.Contract.Sub(&_BasSubDomain.CallOpts, arg0)
}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasSubDomain *BasSubDomainCaller) GetNameByHash(opts *bind.CallOpts, hash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "getNameByHash", hash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasSubDomain *BasSubDomainSession) GetNameByHash(hash [32]byte) ([]byte, error) {
	return _BasSubDomain.Contract.GetNameByHash(&_BasSubDomain.CallOpts, hash)
}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasSubDomain *BasSubDomainCallerSession) GetNameByHash(hash [32]byte) ([]byte, error) {
	return _BasSubDomain.Contract.GetNameByHash(&_BasSubDomain.CallOpts, hash)
}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasSubDomain *BasSubDomainCaller) HasDomain(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "hasDomain", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasSubDomain *BasSubDomainSession) HasDomain(hash [32]byte) (bool, error) {
	return _BasSubDomain.Contract.HasDomain(&_BasSubDomain.CallOpts, hash)
}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasSubDomain *BasSubDomainCallerSession) HasDomain(hash [32]byte) (bool, error) {
	return _BasSubDomain.Contract.HasDomain(&_BasSubDomain.CallOpts, hash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasSubDomain *BasSubDomainCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasSubDomain *BasSubDomainSession) Rel() (common.Address, error) {
	return _BasSubDomain.Contract.Rel(&_BasSubDomain.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasSubDomain *BasSubDomainCallerSession) Rel() (common.Address, error) {
	return _BasSubDomain.Contract.Rel(&_BasSubDomain.CallOpts)
}

// VerifySub is a free data retrieval call binding the contract method 0x4e99d93e.
//
// Solidity: function verifySub(bytes name) pure returns(bool)
func (_BasSubDomain *BasSubDomainCaller) VerifySub(opts *bind.CallOpts, name []byte) (bool, error) {
	var out []interface{}
	err := _BasSubDomain.contract.Call(opts, &out, "verifySub", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySub is a free data retrieval call binding the contract method 0x4e99d93e.
//
// Solidity: function verifySub(bytes name) pure returns(bool)
func (_BasSubDomain *BasSubDomainSession) VerifySub(name []byte) (bool, error) {
	return _BasSubDomain.Contract.VerifySub(&_BasSubDomain.CallOpts, name)
}

// VerifySub is a free data retrieval call binding the contract method 0x4e99d93e.
//
// Solidity: function verifySub(bytes name) pure returns(bool)
func (_BasSubDomain *BasSubDomainCallerSession) VerifySub(name []byte) (bool, error) {
	return _BasSubDomain.Contract.VerifySub(&_BasSubDomain.CallOpts, name)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasSubDomain *BasSubDomainTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasSubDomain.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasSubDomain *BasSubDomainSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ChangeDAO(&_BasSubDomain.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasSubDomain *BasSubDomainTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ChangeDAO(&_BasSubDomain.TransactOpts, newDao)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasSubDomain *BasSubDomainTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasSubDomain.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasSubDomain *BasSubDomainSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ChangeRelation(&_BasSubDomain.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasSubDomain *BasSubDomainTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ChangeRelation(&_BasSubDomain.TransactOpts, new_rel)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasSubDomain *BasSubDomainTransactor) Recharge(opts *bind.TransactOpts, nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasSubDomain.contract.Transact(opts, "recharge", nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasSubDomain *BasSubDomainSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasSubDomain.Contract.Recharge(&_BasSubDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasSubDomain *BasSubDomainTransactorSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasSubDomain.Contract.Recharge(&_BasSubDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x6a9d2b89.
//
// Solidity: function replaceOrCreate(bytes subdomain, bytes32 rootHash, uint256 newExpire, address applicant) returns()
func (_BasSubDomain *BasSubDomainTransactor) ReplaceOrCreate(opts *bind.TransactOpts, subdomain []byte, rootHash [32]byte, newExpire *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasSubDomain.contract.Transact(opts, "replaceOrCreate", subdomain, rootHash, newExpire, applicant)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x6a9d2b89.
//
// Solidity: function replaceOrCreate(bytes subdomain, bytes32 rootHash, uint256 newExpire, address applicant) returns()
func (_BasSubDomain *BasSubDomainSession) ReplaceOrCreate(subdomain []byte, rootHash [32]byte, newExpire *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ReplaceOrCreate(&_BasSubDomain.TransactOpts, subdomain, rootHash, newExpire, applicant)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x6a9d2b89.
//
// Solidity: function replaceOrCreate(bytes subdomain, bytes32 rootHash, uint256 newExpire, address applicant) returns()
func (_BasSubDomain *BasSubDomainTransactorSession) ReplaceOrCreate(subdomain []byte, rootHash [32]byte, newExpire *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasSubDomain.Contract.ReplaceOrCreate(&_BasSubDomain.TransactOpts, subdomain, rootHash, newExpire, applicant)
}

// BasSubDomainNewSubDomainIterator is returned from FilterNewSubDomain and is used to iterate over the raw logs and unpacked data for NewSubDomain events raised by the BasSubDomain contract.
type BasSubDomainNewSubDomainIterator struct {
	Event *BasSubDomainNewSubDomain // Event containing the contract specifics and raw log

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
func (it *BasSubDomainNewSubDomainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasSubDomainNewSubDomain)
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
		it.Event = new(BasSubDomainNewSubDomain)
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
func (it *BasSubDomainNewSubDomainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasSubDomainNewSubDomainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasSubDomainNewSubDomain represents a NewSubDomain event raised by the BasSubDomain contract.
type BasSubDomainNewSubDomain struct {
	NameHash  [32]byte
	TotalName []byte
	RootHash  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewSubDomain is a free log retrieval operation binding the contract event 0x19082335ea9754093d3eb8cb128a00b859749b65ad1d7dfc60cc5d83cc883b72.
//
// Solidity: event NewSubDomain(bytes32 indexed nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) FilterNewSubDomain(opts *bind.FilterOpts, nameHash [][32]byte) (*BasSubDomainNewSubDomainIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasSubDomain.contract.FilterLogs(opts, "NewSubDomain", nameHashRule)
	if err != nil {
		return nil, err
	}
	return &BasSubDomainNewSubDomainIterator{contract: _BasSubDomain.contract, event: "NewSubDomain", logs: logs, sub: sub}, nil
}

// WatchNewSubDomain is a free log subscription operation binding the contract event 0x19082335ea9754093d3eb8cb128a00b859749b65ad1d7dfc60cc5d83cc883b72.
//
// Solidity: event NewSubDomain(bytes32 indexed nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) WatchNewSubDomain(opts *bind.WatchOpts, sink chan<- *BasSubDomainNewSubDomain, nameHash [][32]byte) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasSubDomain.contract.WatchLogs(opts, "NewSubDomain", nameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasSubDomainNewSubDomain)
				if err := _BasSubDomain.contract.UnpackLog(event, "NewSubDomain", log); err != nil {
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

// ParseNewSubDomain is a log parse operation binding the contract event 0x19082335ea9754093d3eb8cb128a00b859749b65ad1d7dfc60cc5d83cc883b72.
//
// Solidity: event NewSubDomain(bytes32 indexed nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) ParseNewSubDomain(log types.Log) (*BasSubDomainNewSubDomain, error) {
	event := new(BasSubDomainNewSubDomain)
	if err := _BasSubDomain.contract.UnpackLog(event, "NewSubDomain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasSubDomainRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the BasSubDomain contract.
type BasSubDomainRechargeIterator struct {
	Event *BasSubDomainRecharge // Event containing the contract specifics and raw log

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
func (it *BasSubDomainRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasSubDomainRecharge)
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
		it.Event = new(BasSubDomainRecharge)
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
func (it *BasSubDomainRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasSubDomainRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasSubDomainRecharge represents a Recharge event raised by the BasSubDomain contract.
type BasSubDomainRecharge struct {
	NameHash [32]byte
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasSubDomain *BasSubDomainFilterer) FilterRecharge(opts *bind.FilterOpts) (*BasSubDomainRechargeIterator, error) {

	logs, sub, err := _BasSubDomain.contract.FilterLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return &BasSubDomainRechargeIterator{contract: _BasSubDomain.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasSubDomain *BasSubDomainFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *BasSubDomainRecharge) (event.Subscription, error) {

	logs, sub, err := _BasSubDomain.contract.WatchLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasSubDomainRecharge)
				if err := _BasSubDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
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

// ParseRecharge is a log parse operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasSubDomain *BasSubDomainFilterer) ParseRecharge(log types.Log) (*BasSubDomainRecharge, error) {
	event := new(BasSubDomainRecharge)
	if err := _BasSubDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasSubDomainSubDataReplacedIterator is returned from FilterSubDataReplaced and is used to iterate over the raw logs and unpacked data for SubDataReplaced events raised by the BasSubDomain contract.
type BasSubDomainSubDataReplacedIterator struct {
	Event *BasSubDomainSubDataReplaced // Event containing the contract specifics and raw log

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
func (it *BasSubDomainSubDataReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasSubDomainSubDataReplaced)
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
		it.Event = new(BasSubDomainSubDataReplaced)
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
func (it *BasSubDomainSubDataReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasSubDomainSubDataReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasSubDomainSubDataReplaced represents a SubDataReplaced event raised by the BasSubDomain contract.
type BasSubDomainSubDataReplaced struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubDataReplaced is a free log retrieval operation binding the contract event 0xc887b1cdecd42cef29eb320b6d34477a573f2cc3011242cf421670ef70ee41c3.
//
// Solidity: event SubDataReplaced(bytes32 nameHash)
func (_BasSubDomain *BasSubDomainFilterer) FilterSubDataReplaced(opts *bind.FilterOpts) (*BasSubDomainSubDataReplacedIterator, error) {

	logs, sub, err := _BasSubDomain.contract.FilterLogs(opts, "SubDataReplaced")
	if err != nil {
		return nil, err
	}
	return &BasSubDomainSubDataReplacedIterator{contract: _BasSubDomain.contract, event: "SubDataReplaced", logs: logs, sub: sub}, nil
}

// WatchSubDataReplaced is a free log subscription operation binding the contract event 0xc887b1cdecd42cef29eb320b6d34477a573f2cc3011242cf421670ef70ee41c3.
//
// Solidity: event SubDataReplaced(bytes32 nameHash)
func (_BasSubDomain *BasSubDomainFilterer) WatchSubDataReplaced(opts *bind.WatchOpts, sink chan<- *BasSubDomainSubDataReplaced) (event.Subscription, error) {

	logs, sub, err := _BasSubDomain.contract.WatchLogs(opts, "SubDataReplaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasSubDomainSubDataReplaced)
				if err := _BasSubDomain.contract.UnpackLog(event, "SubDataReplaced", log); err != nil {
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

// ParseSubDataReplaced is a log parse operation binding the contract event 0xc887b1cdecd42cef29eb320b6d34477a573f2cc3011242cf421670ef70ee41c3.
//
// Solidity: event SubDataReplaced(bytes32 nameHash)
func (_BasSubDomain *BasSubDomainFilterer) ParseSubDataReplaced(log types.Log) (*BasSubDomainSubDataReplaced, error) {
	event := new(BasSubDomainSubDataReplaced)
	if err := _BasSubDomain.contract.UnpackLog(event, "SubDataReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasSubDomainSubUpdateIterator is returned from FilterSubUpdate and is used to iterate over the raw logs and unpacked data for SubUpdate events raised by the BasSubDomain contract.
type BasSubDomainSubUpdateIterator struct {
	Event *BasSubDomainSubUpdate // Event containing the contract specifics and raw log

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
func (it *BasSubDomainSubUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasSubDomainSubUpdate)
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
		it.Event = new(BasSubDomainSubUpdate)
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
func (it *BasSubDomainSubUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasSubDomainSubUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasSubDomainSubUpdate represents a SubUpdate event raised by the BasSubDomain contract.
type BasSubDomainSubUpdate struct {
	NameHash  [32]byte
	TotalName []byte
	RootHash  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubUpdate is a free log retrieval operation binding the contract event 0xb6f50a6e1d60b4bd38ab7675686333c3c71f3163e00d5fd6869810b68303ca08.
//
// Solidity: event SubUpdate(bytes32 nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) FilterSubUpdate(opts *bind.FilterOpts) (*BasSubDomainSubUpdateIterator, error) {

	logs, sub, err := _BasSubDomain.contract.FilterLogs(opts, "SubUpdate")
	if err != nil {
		return nil, err
	}
	return &BasSubDomainSubUpdateIterator{contract: _BasSubDomain.contract, event: "SubUpdate", logs: logs, sub: sub}, nil
}

// WatchSubUpdate is a free log subscription operation binding the contract event 0xb6f50a6e1d60b4bd38ab7675686333c3c71f3163e00d5fd6869810b68303ca08.
//
// Solidity: event SubUpdate(bytes32 nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) WatchSubUpdate(opts *bind.WatchOpts, sink chan<- *BasSubDomainSubUpdate) (event.Subscription, error) {

	logs, sub, err := _BasSubDomain.contract.WatchLogs(opts, "SubUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasSubDomainSubUpdate)
				if err := _BasSubDomain.contract.UnpackLog(event, "SubUpdate", log); err != nil {
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

// ParseSubUpdate is a log parse operation binding the contract event 0xb6f50a6e1d60b4bd38ab7675686333c3c71f3163e00d5fd6869810b68303ca08.
//
// Solidity: event SubUpdate(bytes32 nameHash, bytes totalName, bytes32 rootHash)
func (_BasSubDomain *BasSubDomainFilterer) ParseSubUpdate(log types.Log) (*BasSubDomainSubUpdate, error) {
	event := new(BasSubDomainSubUpdate)
	if err := _BasSubDomain.contract.UnpackLog(event, "SubUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
