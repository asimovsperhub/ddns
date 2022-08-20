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

// BasRootDomainABI is the input ABI used to generate the binding from.
const BasRootDomainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"ClosedCustomPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"}],\"name\":\"NewRootDomain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"}],\"name\":\"OpenCustomPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"}],\"name\":\"RootDataReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"RootPublicChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"}],\"name\":\"RootUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RARE_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Root\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"length\",\"type\":\"uint8\"}],\"name\":\"changeRareLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"classifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"classifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"closeCustomPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getNameByHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"hasDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"isRare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"openCustomPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rechargeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEnd\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rootName\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expire\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"customPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"applicant\",\"type\":\"address\"}],\"name\":\"replaceOrCreate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"setPublic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"verifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"verifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BasRootDomain is an auto generated Go binding around an Ethereum contract.
type BasRootDomain struct {
	BasRootDomainCaller     // Read-only binding to the contract
	BasRootDomainTransactor // Write-only binding to the contract
	BasRootDomainFilterer   // Log filterer for contract events
}

// BasRootDomainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasRootDomainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRootDomainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasRootDomainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRootDomainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasRootDomainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasRootDomainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasRootDomainSession struct {
	Contract     *BasRootDomain    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasRootDomainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasRootDomainCallerSession struct {
	Contract *BasRootDomainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BasRootDomainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasRootDomainTransactorSession struct {
	Contract     *BasRootDomainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BasRootDomainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasRootDomainRaw struct {
	Contract *BasRootDomain // Generic contract binding to access the raw methods on
}

// BasRootDomainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasRootDomainCallerRaw struct {
	Contract *BasRootDomainCaller // Generic read-only contract binding to access the raw methods on
}

// BasRootDomainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasRootDomainTransactorRaw struct {
	Contract *BasRootDomainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasRootDomain creates a new instance of BasRootDomain, bound to a specific deployed contract.
func NewBasRootDomain(address common.Address, backend bind.ContractBackend) (*BasRootDomain, error) {
	contract, err := bindBasRootDomain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasRootDomain{BasRootDomainCaller: BasRootDomainCaller{contract: contract}, BasRootDomainTransactor: BasRootDomainTransactor{contract: contract}, BasRootDomainFilterer: BasRootDomainFilterer{contract: contract}}, nil
}

// NewBasRootDomainCaller creates a new read-only instance of BasRootDomain, bound to a specific deployed contract.
func NewBasRootDomainCaller(address common.Address, caller bind.ContractCaller) (*BasRootDomainCaller, error) {
	contract, err := bindBasRootDomain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasRootDomainCaller{contract: contract}, nil
}

// NewBasRootDomainTransactor creates a new write-only instance of BasRootDomain, bound to a specific deployed contract.
func NewBasRootDomainTransactor(address common.Address, transactor bind.ContractTransactor) (*BasRootDomainTransactor, error) {
	contract, err := bindBasRootDomain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasRootDomainTransactor{contract: contract}, nil
}

// NewBasRootDomainFilterer creates a new log filterer instance of BasRootDomain, bound to a specific deployed contract.
func NewBasRootDomainFilterer(address common.Address, filterer bind.ContractFilterer) (*BasRootDomainFilterer, error) {
	contract, err := bindBasRootDomain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasRootDomainFilterer{contract: contract}, nil
}

// bindBasRootDomain binds a generic wrapper to an already deployed contract.
func bindBasRootDomain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasRootDomainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRootDomain *BasRootDomainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasRootDomain.Contract.BasRootDomainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRootDomain *BasRootDomainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRootDomain.Contract.BasRootDomainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRootDomain *BasRootDomainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRootDomain.Contract.BasRootDomainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasRootDomain *BasRootDomainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasRootDomain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasRootDomain *BasRootDomainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasRootDomain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasRootDomain *BasRootDomainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasRootDomain.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRootDomain *BasRootDomainCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRootDomain *BasRootDomainSession) DAOAddress() (common.Address, error) {
	return _BasRootDomain.Contract.DAOAddress(&_BasRootDomain.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasRootDomain *BasRootDomainCallerSession) DAOAddress() (common.Address, error) {
	return _BasRootDomain.Contract.DAOAddress(&_BasRootDomain.CallOpts)
}

// RARELENGTH is a free data retrieval call binding the contract method 0x547fccdb.
//
// Solidity: function RARE_LENGTH() view returns(uint8)
func (_BasRootDomain *BasRootDomainCaller) RARELENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "RARE_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RARELENGTH is a free data retrieval call binding the contract method 0x547fccdb.
//
// Solidity: function RARE_LENGTH() view returns(uint8)
func (_BasRootDomain *BasRootDomainSession) RARELENGTH() (uint8, error) {
	return _BasRootDomain.Contract.RARELENGTH(&_BasRootDomain.CallOpts)
}

// RARELENGTH is a free data retrieval call binding the contract method 0x547fccdb.
//
// Solidity: function RARE_LENGTH() view returns(uint8)
func (_BasRootDomain *BasRootDomainCallerSession) RARELENGTH() (uint8, error) {
	return _BasRootDomain.Contract.RARELENGTH(&_BasRootDomain.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) view returns(bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainCaller) Root(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RootName     []byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
}, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "Root", arg0)

	outstruct := new(struct {
		RootName     []byte
		OpenToPublic bool
		IsCustom     bool
		CustomPrice  *big.Int
	})

	outstruct.RootName = out[0].([]byte)
	outstruct.OpenToPublic = out[1].(bool)
	outstruct.IsCustom = out[2].(bool)
	outstruct.CustomPrice = out[3].(*big.Int)

	return *outstruct, err

}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) view returns(bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainSession) Root(arg0 [32]byte) (struct {
	RootName     []byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
}, error) {
	return _BasRootDomain.Contract.Root(&_BasRootDomain.CallOpts, arg0)
}

// Root is a free data retrieval call binding the contract method 0x83b8202f.
//
// Solidity: function Root(bytes32 ) view returns(bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainCallerSession) Root(arg0 [32]byte) (struct {
	RootName     []byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
}, error) {
	return _BasRootDomain.Contract.Root(&_BasRootDomain.CallOpts, arg0)
}

// ClassifyRoot is a free data retrieval call binding the contract method 0x8ec2c758.
//
// Solidity: function classifyRoot(bytes32 nameHash) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainCaller) ClassifyRoot(opts *bind.CallOpts, nameHash [32]byte) (bool, bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "classifyRoot", nameHash)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// ClassifyRoot is a free data retrieval call binding the contract method 0x8ec2c758.
//
// Solidity: function classifyRoot(bytes32 nameHash) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainSession) ClassifyRoot(nameHash [32]byte) (bool, bool, error) {
	return _BasRootDomain.Contract.ClassifyRoot(&_BasRootDomain.CallOpts, nameHash)
}

// ClassifyRoot is a free data retrieval call binding the contract method 0x8ec2c758.
//
// Solidity: function classifyRoot(bytes32 nameHash) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainCallerSession) ClassifyRoot(nameHash [32]byte) (bool, bool, error) {
	return _BasRootDomain.Contract.ClassifyRoot(&_BasRootDomain.CallOpts, nameHash)
}

// ClassifyRoot0 is a free data retrieval call binding the contract method 0xaeb69767.
//
// Solidity: function classifyRoot(bytes name) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainCaller) ClassifyRoot0(opts *bind.CallOpts, name []byte) (bool, bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "classifyRoot0", name)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// ClassifyRoot0 is a free data retrieval call binding the contract method 0xaeb69767.
//
// Solidity: function classifyRoot(bytes name) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainSession) ClassifyRoot0(name []byte) (bool, bool, error) {
	return _BasRootDomain.Contract.ClassifyRoot0(&_BasRootDomain.CallOpts, name)
}

// ClassifyRoot0 is a free data retrieval call binding the contract method 0xaeb69767.
//
// Solidity: function classifyRoot(bytes name) view returns(bool, bool)
func (_BasRootDomain *BasRootDomainCallerSession) ClassifyRoot0(name []byte) (bool, bool, error) {
	return _BasRootDomain.Contract.ClassifyRoot0(&_BasRootDomain.CallOpts, name)
}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasRootDomain *BasRootDomainCaller) GetNameByHash(opts *bind.CallOpts, hash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "getNameByHash", hash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasRootDomain *BasRootDomainSession) GetNameByHash(hash [32]byte) ([]byte, error) {
	return _BasRootDomain.Contract.GetNameByHash(&_BasRootDomain.CallOpts, hash)
}

// GetNameByHash is a free data retrieval call binding the contract method 0xf03590bf.
//
// Solidity: function getNameByHash(bytes32 hash) view returns(bytes)
func (_BasRootDomain *BasRootDomainCallerSession) GetNameByHash(hash [32]byte) ([]byte, error) {
	return _BasRootDomain.Contract.GetNameByHash(&_BasRootDomain.CallOpts, hash)
}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasRootDomain *BasRootDomainCaller) HasDomain(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "hasDomain", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasRootDomain *BasRootDomainSession) HasDomain(hash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.HasDomain(&_BasRootDomain.CallOpts, hash)
}

// HasDomain is a free data retrieval call binding the contract method 0xd33d3edc.
//
// Solidity: function hasDomain(bytes32 hash) view returns(bool)
func (_BasRootDomain *BasRootDomainCallerSession) HasDomain(hash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.HasDomain(&_BasRootDomain.CallOpts, hash)
}

// IsRare is a free data retrieval call binding the contract method 0xa6aff7c7.
//
// Solidity: function isRare(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainCaller) IsRare(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "isRare", nameHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRare is a free data retrieval call binding the contract method 0xa6aff7c7.
//
// Solidity: function isRare(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainSession) IsRare(nameHash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.IsRare(&_BasRootDomain.CallOpts, nameHash)
}

// IsRare is a free data retrieval call binding the contract method 0xa6aff7c7.
//
// Solidity: function isRare(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainCallerSession) IsRare(nameHash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.IsRare(&_BasRootDomain.CallOpts, nameHash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasRootDomain *BasRootDomainCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasRootDomain *BasRootDomainSession) Rel() (common.Address, error) {
	return _BasRootDomain.Contract.Rel(&_BasRootDomain.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasRootDomain *BasRootDomainCallerSession) Rel() (common.Address, error) {
	return _BasRootDomain.Contract.Rel(&_BasRootDomain.CallOpts)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_BasRootDomain *BasRootDomainCaller) VerifyRoot(opts *bind.CallOpts, name []byte) (bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "verifyRoot", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_BasRootDomain *BasRootDomainSession) VerifyRoot(name []byte) (bool, error) {
	return _BasRootDomain.Contract.VerifyRoot(&_BasRootDomain.CallOpts, name)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_BasRootDomain *BasRootDomainCallerSession) VerifyRoot(name []byte) (bool, error) {
	return _BasRootDomain.Contract.VerifyRoot(&_BasRootDomain.CallOpts, name)
}

// VerifyRoot0 is a free data retrieval call binding the contract method 0x83363bcf.
//
// Solidity: function verifyRoot(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainCaller) VerifyRoot0(opts *bind.CallOpts, nameHash [32]byte) (bool, error) {
	var out []interface{}
	err := _BasRootDomain.contract.Call(opts, &out, "verifyRoot0", nameHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRoot0 is a free data retrieval call binding the contract method 0x83363bcf.
//
// Solidity: function verifyRoot(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainSession) VerifyRoot0(nameHash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.VerifyRoot0(&_BasRootDomain.CallOpts, nameHash)
}

// VerifyRoot0 is a free data retrieval call binding the contract method 0x83363bcf.
//
// Solidity: function verifyRoot(bytes32 nameHash) view returns(bool)
func (_BasRootDomain *BasRootDomainCallerSession) VerifyRoot0(nameHash [32]byte) (bool, error) {
	return _BasRootDomain.Contract.VerifyRoot0(&_BasRootDomain.CallOpts, nameHash)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRootDomain *BasRootDomainTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRootDomain *BasRootDomainSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeDAO(&_BasRootDomain.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeDAO(&_BasRootDomain.TransactOpts, newDao)
}

// ChangeRareLength is a paid mutator transaction binding the contract method 0x8a79e6b5.
//
// Solidity: function changeRareLength(uint8 length) returns()
func (_BasRootDomain *BasRootDomainTransactor) ChangeRareLength(opts *bind.TransactOpts, length uint8) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "changeRareLength", length)
}

// ChangeRareLength is a paid mutator transaction binding the contract method 0x8a79e6b5.
//
// Solidity: function changeRareLength(uint8 length) returns()
func (_BasRootDomain *BasRootDomainSession) ChangeRareLength(length uint8) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeRareLength(&_BasRootDomain.TransactOpts, length)
}

// ChangeRareLength is a paid mutator transaction binding the contract method 0x8a79e6b5.
//
// Solidity: function changeRareLength(uint8 length) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) ChangeRareLength(length uint8) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeRareLength(&_BasRootDomain.TransactOpts, length)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasRootDomain *BasRootDomainTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasRootDomain *BasRootDomainSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeRelation(&_BasRootDomain.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ChangeRelation(&_BasRootDomain.TransactOpts, new_rel)
}

// CloseCustomPrice is a paid mutator transaction binding the contract method 0x818294ab.
//
// Solidity: function closeCustomPrice(bytes32 nameHash) returns()
func (_BasRootDomain *BasRootDomainTransactor) CloseCustomPrice(opts *bind.TransactOpts, nameHash [32]byte) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "closeCustomPrice", nameHash)
}

// CloseCustomPrice is a paid mutator transaction binding the contract method 0x818294ab.
//
// Solidity: function closeCustomPrice(bytes32 nameHash) returns()
func (_BasRootDomain *BasRootDomainSession) CloseCustomPrice(nameHash [32]byte) (*types.Transaction, error) {
	return _BasRootDomain.Contract.CloseCustomPrice(&_BasRootDomain.TransactOpts, nameHash)
}

// CloseCustomPrice is a paid mutator transaction binding the contract method 0x818294ab.
//
// Solidity: function closeCustomPrice(bytes32 nameHash) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) CloseCustomPrice(nameHash [32]byte) (*types.Transaction, error) {
	return _BasRootDomain.Contract.CloseCustomPrice(&_BasRootDomain.TransactOpts, nameHash)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0x85ccba48.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 customPrice, address operator) returns()
func (_BasRootDomain *BasRootDomainTransactor) OpenCustomPrice(opts *bind.TransactOpts, nameHash [32]byte, customPrice *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "openCustomPrice", nameHash, customPrice, operator)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0x85ccba48.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 customPrice, address operator) returns()
func (_BasRootDomain *BasRootDomainSession) OpenCustomPrice(nameHash [32]byte, customPrice *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.OpenCustomPrice(&_BasRootDomain.TransactOpts, nameHash, customPrice, operator)
}

// OpenCustomPrice is a paid mutator transaction binding the contract method 0x85ccba48.
//
// Solidity: function openCustomPrice(bytes32 nameHash, uint256 customPrice, address operator) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) OpenCustomPrice(nameHash [32]byte, customPrice *big.Int, operator common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.OpenCustomPrice(&_BasRootDomain.TransactOpts, nameHash, customPrice, operator)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasRootDomain *BasRootDomainTransactor) Recharge(opts *bind.TransactOpts, nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "recharge", nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasRootDomain *BasRootDomainSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasRootDomain.Contract.Recharge(&_BasRootDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// Recharge is a paid mutator transaction binding the contract method 0xbd69144b.
//
// Solidity: function recharge(bytes32 nameHash, uint256 rechargeTime, uint256 maxEnd) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) Recharge(nameHash [32]byte, rechargeTime *big.Int, maxEnd *big.Int) (*types.Transaction, error) {
	return _BasRootDomain.Contract.Recharge(&_BasRootDomain.TransactOpts, nameHash, rechargeTime, maxEnd)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x361be604.
//
// Solidity: function replaceOrCreate(bytes rootName, uint256 expire, bool openToPublic, bool isCustom, uint256 customPrice, address applicant) returns()
func (_BasRootDomain *BasRootDomainTransactor) ReplaceOrCreate(opts *bind.TransactOpts, rootName []byte, expire *big.Int, openToPublic bool, isCustom bool, customPrice *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "replaceOrCreate", rootName, expire, openToPublic, isCustom, customPrice, applicant)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x361be604.
//
// Solidity: function replaceOrCreate(bytes rootName, uint256 expire, bool openToPublic, bool isCustom, uint256 customPrice, address applicant) returns()
func (_BasRootDomain *BasRootDomainSession) ReplaceOrCreate(rootName []byte, expire *big.Int, openToPublic bool, isCustom bool, customPrice *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ReplaceOrCreate(&_BasRootDomain.TransactOpts, rootName, expire, openToPublic, isCustom, customPrice, applicant)
}

// ReplaceOrCreate is a paid mutator transaction binding the contract method 0x361be604.
//
// Solidity: function replaceOrCreate(bytes rootName, uint256 expire, bool openToPublic, bool isCustom, uint256 customPrice, address applicant) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) ReplaceOrCreate(rootName []byte, expire *big.Int, openToPublic bool, isCustom bool, customPrice *big.Int, applicant common.Address) (*types.Transaction, error) {
	return _BasRootDomain.Contract.ReplaceOrCreate(&_BasRootDomain.TransactOpts, rootName, expire, openToPublic, isCustom, customPrice, applicant)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 nameHash, bool isOpen) returns()
func (_BasRootDomain *BasRootDomainTransactor) SetPublic(opts *bind.TransactOpts, nameHash [32]byte, isOpen bool) (*types.Transaction, error) {
	return _BasRootDomain.contract.Transact(opts, "setPublic", nameHash, isOpen)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 nameHash, bool isOpen) returns()
func (_BasRootDomain *BasRootDomainSession) SetPublic(nameHash [32]byte, isOpen bool) (*types.Transaction, error) {
	return _BasRootDomain.Contract.SetPublic(&_BasRootDomain.TransactOpts, nameHash, isOpen)
}

// SetPublic is a paid mutator transaction binding the contract method 0x92c7e0ba.
//
// Solidity: function setPublic(bytes32 nameHash, bool isOpen) returns()
func (_BasRootDomain *BasRootDomainTransactorSession) SetPublic(nameHash [32]byte, isOpen bool) (*types.Transaction, error) {
	return _BasRootDomain.Contract.SetPublic(&_BasRootDomain.TransactOpts, nameHash, isOpen)
}

// BasRootDomainClosedCustomPriceIterator is returned from FilterClosedCustomPrice and is used to iterate over the raw logs and unpacked data for ClosedCustomPrice events raised by the BasRootDomain contract.
type BasRootDomainClosedCustomPriceIterator struct {
	Event *BasRootDomainClosedCustomPrice // Event containing the contract specifics and raw log

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
func (it *BasRootDomainClosedCustomPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainClosedCustomPrice)
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
		it.Event = new(BasRootDomainClosedCustomPrice)
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
func (it *BasRootDomainClosedCustomPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainClosedCustomPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainClosedCustomPrice represents a ClosedCustomPrice event raised by the BasRootDomain contract.
type BasRootDomainClosedCustomPrice struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClosedCustomPrice is a free log retrieval operation binding the contract event 0xb465b26060ed6bba1cefc97652c65086fbda851a04238bbc87b000854cbecb43.
//
// Solidity: event ClosedCustomPrice(bytes32 nameHash)
func (_BasRootDomain *BasRootDomainFilterer) FilterClosedCustomPrice(opts *bind.FilterOpts) (*BasRootDomainClosedCustomPriceIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "ClosedCustomPrice")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainClosedCustomPriceIterator{contract: _BasRootDomain.contract, event: "ClosedCustomPrice", logs: logs, sub: sub}, nil
}

// WatchClosedCustomPrice is a free log subscription operation binding the contract event 0xb465b26060ed6bba1cefc97652c65086fbda851a04238bbc87b000854cbecb43.
//
// Solidity: event ClosedCustomPrice(bytes32 nameHash)
func (_BasRootDomain *BasRootDomainFilterer) WatchClosedCustomPrice(opts *bind.WatchOpts, sink chan<- *BasRootDomainClosedCustomPrice) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "ClosedCustomPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainClosedCustomPrice)
				if err := _BasRootDomain.contract.UnpackLog(event, "ClosedCustomPrice", log); err != nil {
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

// ParseClosedCustomPrice is a log parse operation binding the contract event 0xb465b26060ed6bba1cefc97652c65086fbda851a04238bbc87b000854cbecb43.
//
// Solidity: event ClosedCustomPrice(bytes32 nameHash)
func (_BasRootDomain *BasRootDomainFilterer) ParseClosedCustomPrice(log types.Log) (*BasRootDomainClosedCustomPrice, error) {
	event := new(BasRootDomainClosedCustomPrice)
	if err := _BasRootDomain.contract.UnpackLog(event, "ClosedCustomPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainNewRootDomainIterator is returned from FilterNewRootDomain and is used to iterate over the raw logs and unpacked data for NewRootDomain events raised by the BasRootDomain contract.
type BasRootDomainNewRootDomainIterator struct {
	Event *BasRootDomainNewRootDomain // Event containing the contract specifics and raw log

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
func (it *BasRootDomainNewRootDomainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainNewRootDomain)
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
		it.Event = new(BasRootDomainNewRootDomain)
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
func (it *BasRootDomainNewRootDomainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainNewRootDomainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainNewRootDomain represents a NewRootDomain event raised by the BasRootDomain contract.
type BasRootDomainNewRootDomain struct {
	NameHash     [32]byte
	RootName     []byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewRootDomain is a free log retrieval operation binding the contract event 0xbbbb1eb26970a83df45e3eef1991dc1ae6680e3627794c05a9eabb8ecaecc5da.
//
// Solidity: event NewRootDomain(bytes32 indexed nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) FilterNewRootDomain(opts *bind.FilterOpts, nameHash [][32]byte) (*BasRootDomainNewRootDomainIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "NewRootDomain", nameHashRule)
	if err != nil {
		return nil, err
	}
	return &BasRootDomainNewRootDomainIterator{contract: _BasRootDomain.contract, event: "NewRootDomain", logs: logs, sub: sub}, nil
}

// WatchNewRootDomain is a free log subscription operation binding the contract event 0xbbbb1eb26970a83df45e3eef1991dc1ae6680e3627794c05a9eabb8ecaecc5da.
//
// Solidity: event NewRootDomain(bytes32 indexed nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) WatchNewRootDomain(opts *bind.WatchOpts, sink chan<- *BasRootDomainNewRootDomain, nameHash [][32]byte) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "NewRootDomain", nameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainNewRootDomain)
				if err := _BasRootDomain.contract.UnpackLog(event, "NewRootDomain", log); err != nil {
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

// ParseNewRootDomain is a log parse operation binding the contract event 0xbbbb1eb26970a83df45e3eef1991dc1ae6680e3627794c05a9eabb8ecaecc5da.
//
// Solidity: event NewRootDomain(bytes32 indexed nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) ParseNewRootDomain(log types.Log) (*BasRootDomainNewRootDomain, error) {
	event := new(BasRootDomainNewRootDomain)
	if err := _BasRootDomain.contract.UnpackLog(event, "NewRootDomain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainOpenCustomPriceIterator is returned from FilterOpenCustomPrice and is used to iterate over the raw logs and unpacked data for OpenCustomPrice events raised by the BasRootDomain contract.
type BasRootDomainOpenCustomPriceIterator struct {
	Event *BasRootDomainOpenCustomPrice // Event containing the contract specifics and raw log

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
func (it *BasRootDomainOpenCustomPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainOpenCustomPrice)
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
		it.Event = new(BasRootDomainOpenCustomPrice)
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
func (it *BasRootDomainOpenCustomPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainOpenCustomPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainOpenCustomPrice represents a OpenCustomPrice event raised by the BasRootDomain contract.
type BasRootDomainOpenCustomPrice struct {
	NameHash    [32]byte
	CustomPrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpenCustomPrice is a free log retrieval operation binding the contract event 0xd03330f8e5de2f023ce8d5a62af572a9f8963ad68c2a218a90be044647ef662b.
//
// Solidity: event OpenCustomPrice(bytes32 nameHash, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) FilterOpenCustomPrice(opts *bind.FilterOpts) (*BasRootDomainOpenCustomPriceIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "OpenCustomPrice")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainOpenCustomPriceIterator{contract: _BasRootDomain.contract, event: "OpenCustomPrice", logs: logs, sub: sub}, nil
}

// WatchOpenCustomPrice is a free log subscription operation binding the contract event 0xd03330f8e5de2f023ce8d5a62af572a9f8963ad68c2a218a90be044647ef662b.
//
// Solidity: event OpenCustomPrice(bytes32 nameHash, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) WatchOpenCustomPrice(opts *bind.WatchOpts, sink chan<- *BasRootDomainOpenCustomPrice) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "OpenCustomPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainOpenCustomPrice)
				if err := _BasRootDomain.contract.UnpackLog(event, "OpenCustomPrice", log); err != nil {
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

// ParseOpenCustomPrice is a log parse operation binding the contract event 0xd03330f8e5de2f023ce8d5a62af572a9f8963ad68c2a218a90be044647ef662b.
//
// Solidity: event OpenCustomPrice(bytes32 nameHash, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) ParseOpenCustomPrice(log types.Log) (*BasRootDomainOpenCustomPrice, error) {
	event := new(BasRootDomainOpenCustomPrice)
	if err := _BasRootDomain.contract.UnpackLog(event, "OpenCustomPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the BasRootDomain contract.
type BasRootDomainRechargeIterator struct {
	Event *BasRootDomainRecharge // Event containing the contract specifics and raw log

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
func (it *BasRootDomainRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainRecharge)
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
		it.Event = new(BasRootDomainRecharge)
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
func (it *BasRootDomainRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainRecharge represents a Recharge event raised by the BasRootDomain contract.
type BasRootDomainRecharge struct {
	NameHash [32]byte
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasRootDomain *BasRootDomainFilterer) FilterRecharge(opts *bind.FilterOpts) (*BasRootDomainRechargeIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainRechargeIterator{contract: _BasRootDomain.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0xfa34f6f082bf5f2661b70be0dc544cb62a2d7b7888e2078f2d1d00b5eadf97c6.
//
// Solidity: event Recharge(bytes32 nameHash, uint256 duration)
func (_BasRootDomain *BasRootDomainFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *BasRootDomainRecharge) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "Recharge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainRecharge)
				if err := _BasRootDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
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
func (_BasRootDomain *BasRootDomainFilterer) ParseRecharge(log types.Log) (*BasRootDomainRecharge, error) {
	event := new(BasRootDomainRecharge)
	if err := _BasRootDomain.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainRootDataReplacedIterator is returned from FilterRootDataReplaced and is used to iterate over the raw logs and unpacked data for RootDataReplaced events raised by the BasRootDomain contract.
type BasRootDomainRootDataReplacedIterator struct {
	Event *BasRootDomainRootDataReplaced // Event containing the contract specifics and raw log

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
func (it *BasRootDomainRootDataReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainRootDataReplaced)
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
		it.Event = new(BasRootDomainRootDataReplaced)
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
func (it *BasRootDomainRootDataReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainRootDataReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainRootDataReplaced represents a RootDataReplaced event raised by the BasRootDomain contract.
type BasRootDomainRootDataReplaced struct {
	NameHash     [32]byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRootDataReplaced is a free log retrieval operation binding the contract event 0x60062390b20aaea2d79a828e1c94dd0145b7736e924d061a8a5f013a068fae5e.
//
// Solidity: event RootDataReplaced(bytes32 nameHash, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) FilterRootDataReplaced(opts *bind.FilterOpts) (*BasRootDomainRootDataReplacedIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "RootDataReplaced")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainRootDataReplacedIterator{contract: _BasRootDomain.contract, event: "RootDataReplaced", logs: logs, sub: sub}, nil
}

// WatchRootDataReplaced is a free log subscription operation binding the contract event 0x60062390b20aaea2d79a828e1c94dd0145b7736e924d061a8a5f013a068fae5e.
//
// Solidity: event RootDataReplaced(bytes32 nameHash, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) WatchRootDataReplaced(opts *bind.WatchOpts, sink chan<- *BasRootDomainRootDataReplaced) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "RootDataReplaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainRootDataReplaced)
				if err := _BasRootDomain.contract.UnpackLog(event, "RootDataReplaced", log); err != nil {
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

// ParseRootDataReplaced is a log parse operation binding the contract event 0x60062390b20aaea2d79a828e1c94dd0145b7736e924d061a8a5f013a068fae5e.
//
// Solidity: event RootDataReplaced(bytes32 nameHash, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) ParseRootDataReplaced(log types.Log) (*BasRootDomainRootDataReplaced, error) {
	event := new(BasRootDomainRootDataReplaced)
	if err := _BasRootDomain.contract.UnpackLog(event, "RootDataReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainRootPublicChangedIterator is returned from FilterRootPublicChanged and is used to iterate over the raw logs and unpacked data for RootPublicChanged events raised by the BasRootDomain contract.
type BasRootDomainRootPublicChangedIterator struct {
	Event *BasRootDomainRootPublicChanged // Event containing the contract specifics and raw log

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
func (it *BasRootDomainRootPublicChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainRootPublicChanged)
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
		it.Event = new(BasRootDomainRootPublicChanged)
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
func (it *BasRootDomainRootPublicChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainRootPublicChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainRootPublicChanged represents a RootPublicChanged event raised by the BasRootDomain contract.
type BasRootDomainRootPublicChanged struct {
	NameHash [32]byte
	IsOpen   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRootPublicChanged is a free log retrieval operation binding the contract event 0xd3a42ab11ab9f61994898fcfd96b3dbd3bcb4f02919e13f597006df4af4bce1a.
//
// Solidity: event RootPublicChanged(bytes32 nameHash, bool isOpen)
func (_BasRootDomain *BasRootDomainFilterer) FilterRootPublicChanged(opts *bind.FilterOpts) (*BasRootDomainRootPublicChangedIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "RootPublicChanged")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainRootPublicChangedIterator{contract: _BasRootDomain.contract, event: "RootPublicChanged", logs: logs, sub: sub}, nil
}

// WatchRootPublicChanged is a free log subscription operation binding the contract event 0xd3a42ab11ab9f61994898fcfd96b3dbd3bcb4f02919e13f597006df4af4bce1a.
//
// Solidity: event RootPublicChanged(bytes32 nameHash, bool isOpen)
func (_BasRootDomain *BasRootDomainFilterer) WatchRootPublicChanged(opts *bind.WatchOpts, sink chan<- *BasRootDomainRootPublicChanged) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "RootPublicChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainRootPublicChanged)
				if err := _BasRootDomain.contract.UnpackLog(event, "RootPublicChanged", log); err != nil {
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

// ParseRootPublicChanged is a log parse operation binding the contract event 0xd3a42ab11ab9f61994898fcfd96b3dbd3bcb4f02919e13f597006df4af4bce1a.
//
// Solidity: event RootPublicChanged(bytes32 nameHash, bool isOpen)
func (_BasRootDomain *BasRootDomainFilterer) ParseRootPublicChanged(log types.Log) (*BasRootDomainRootPublicChanged, error) {
	event := new(BasRootDomainRootPublicChanged)
	if err := _BasRootDomain.contract.UnpackLog(event, "RootPublicChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasRootDomainRootUpdateIterator is returned from FilterRootUpdate and is used to iterate over the raw logs and unpacked data for RootUpdate events raised by the BasRootDomain contract.
type BasRootDomainRootUpdateIterator struct {
	Event *BasRootDomainRootUpdate // Event containing the contract specifics and raw log

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
func (it *BasRootDomainRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasRootDomainRootUpdate)
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
		it.Event = new(BasRootDomainRootUpdate)
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
func (it *BasRootDomainRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasRootDomainRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasRootDomainRootUpdate represents a RootUpdate event raised by the BasRootDomain contract.
type BasRootDomainRootUpdate struct {
	NameHash     [32]byte
	RootName     []byte
	OpenToPublic bool
	IsCustom     bool
	CustomPrice  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRootUpdate is a free log retrieval operation binding the contract event 0x5e3a0d25ee4ea2b575d3c4115cf54f580e1a1e5f54f7ddca0ab350ae4ad149eb.
//
// Solidity: event RootUpdate(bytes32 nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) FilterRootUpdate(opts *bind.FilterOpts) (*BasRootDomainRootUpdateIterator, error) {

	logs, sub, err := _BasRootDomain.contract.FilterLogs(opts, "RootUpdate")
	if err != nil {
		return nil, err
	}
	return &BasRootDomainRootUpdateIterator{contract: _BasRootDomain.contract, event: "RootUpdate", logs: logs, sub: sub}, nil
}

// WatchRootUpdate is a free log subscription operation binding the contract event 0x5e3a0d25ee4ea2b575d3c4115cf54f580e1a1e5f54f7ddca0ab350ae4ad149eb.
//
// Solidity: event RootUpdate(bytes32 nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) WatchRootUpdate(opts *bind.WatchOpts, sink chan<- *BasRootDomainRootUpdate) (event.Subscription, error) {

	logs, sub, err := _BasRootDomain.contract.WatchLogs(opts, "RootUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasRootDomainRootUpdate)
				if err := _BasRootDomain.contract.UnpackLog(event, "RootUpdate", log); err != nil {
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

// ParseRootUpdate is a log parse operation binding the contract event 0x5e3a0d25ee4ea2b575d3c4115cf54f580e1a1e5f54f7ddca0ab350ae4ad149eb.
//
// Solidity: event RootUpdate(bytes32 nameHash, bytes rootName, bool openToPublic, bool isCustom, uint256 customPrice)
func (_BasRootDomain *BasRootDomainFilterer) ParseRootUpdate(log types.Log) (*BasRootDomainRootUpdate, error) {
	event := new(BasRootDomainRootUpdate)
	if err := _BasRootDomain.contract.UnpackLog(event, "RootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
