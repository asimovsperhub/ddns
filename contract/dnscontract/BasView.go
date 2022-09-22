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

// BasViewMetaData contains all meta data concerning the BasView contract.
var BasViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rel_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDao\",\"type\":\"address\"}],\"name\":\"ChangeDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ErrorCode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_rel\",\"type\":\"address\"}],\"name\":\"changeRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isCustom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cusPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"durationInYear\",\"type\":\"uint8\"}],\"name\":\"checkRootRegistry\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sName\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"durationInYear\",\"type\":\"uint256\"}],\"name\":\"checkSubRegistry\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"domainIsWild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOANNParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"MAX_YEAR\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AROOT_GAS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"BROOT_GAS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SUB_GAS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CUSTOM_PRICE_GAS\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"queryDomainConfigs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"A\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"AAAA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"MX\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"BlockChain\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"IOTA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"CName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"MXBCA\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"queryDomainEmailInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"openToPublic\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"queryDomainInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRoot\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rIsOpen\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rIsCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rIsRare\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rCusPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isMarketOrder\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mailHash\",\"type\":\"bytes32\"}],\"name\":\"queryEmailInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"domainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"aliasName\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bcAddress\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"queryOrderInfo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rel\",\"outputs\":[{\"internalType\":\"contractBasRelations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"setErrorCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasViewABI is the input ABI used to generate the binding from.
// Deprecated: Use BasViewMetaData.ABI instead.
var BasViewABI = BasViewMetaData.ABI

// BasView is an auto generated Go binding around an Ethereum contract.
type BasView struct {
	BasViewCaller     // Read-only binding to the contract
	BasViewTransactor // Write-only binding to the contract
	BasViewFilterer   // Log filterer for contract events
}

// BasViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasViewSession struct {
	Contract     *BasView          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasViewCallerSession struct {
	Contract *BasViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BasViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasViewTransactorSession struct {
	Contract     *BasViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasViewRaw struct {
	Contract *BasView // Generic contract binding to access the raw methods on
}

// BasViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasViewCallerRaw struct {
	Contract *BasViewCaller // Generic read-only contract binding to access the raw methods on
}

// BasViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasViewTransactorRaw struct {
	Contract *BasViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasView creates a new instance of BasView, bound to a specific deployed contract.
func NewBasView(address common.Address, backend bind.ContractBackend) (*BasView, error) {
	contract, err := bindBasView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasView{BasViewCaller: BasViewCaller{contract: contract}, BasViewTransactor: BasViewTransactor{contract: contract}, BasViewFilterer: BasViewFilterer{contract: contract}}, nil
}

// NewBasViewCaller creates a new read-only instance of BasView, bound to a specific deployed contract.
func NewBasViewCaller(address common.Address, caller bind.ContractCaller) (*BasViewCaller, error) {
	contract, err := bindBasView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasViewCaller{contract: contract}, nil
}

// NewBasViewTransactor creates a new write-only instance of BasView, bound to a specific deployed contract.
func NewBasViewTransactor(address common.Address, transactor bind.ContractTransactor) (*BasViewTransactor, error) {
	contract, err := bindBasView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasViewTransactor{contract: contract}, nil
}

// NewBasViewFilterer creates a new log filterer instance of BasView, bound to a specific deployed contract.
func NewBasViewFilterer(address common.Address, filterer bind.ContractFilterer) (*BasViewFilterer, error) {
	contract, err := bindBasView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasViewFilterer{contract: contract}, nil
}

// bindBasView binds a generic wrapper to an already deployed contract.
func bindBasView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasView *BasViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasView.Contract.BasViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasView *BasViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasView.Contract.BasViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasView *BasViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasView.Contract.BasViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasView *BasViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasView *BasViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasView *BasViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasView.Contract.contract.Transact(opts, method, params...)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasView *BasViewCaller) DAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "DAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasView *BasViewSession) DAOAddress() (common.Address, error) {
	return _BasView.Contract.DAOAddress(&_BasView.CallOpts)
}

// DAOAddress is a free data retrieval call binding the contract method 0xd392eab1.
//
// Solidity: function DAOAddress() view returns(address)
func (_BasView *BasViewCallerSession) DAOAddress() (common.Address, error) {
	return _BasView.Contract.DAOAddress(&_BasView.CallOpts)
}

// ErrorCode is a free data retrieval call binding the contract method 0x5d05fa53.
//
// Solidity: function ErrorCode(uint8 ) view returns(string)
func (_BasView *BasViewCaller) ErrorCode(opts *bind.CallOpts, arg0 uint8) (string, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "ErrorCode", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ErrorCode is a free data retrieval call binding the contract method 0x5d05fa53.
//
// Solidity: function ErrorCode(uint8 ) view returns(string)
func (_BasView *BasViewSession) ErrorCode(arg0 uint8) (string, error) {
	return _BasView.Contract.ErrorCode(&_BasView.CallOpts, arg0)
}

// ErrorCode is a free data retrieval call binding the contract method 0x5d05fa53.
//
// Solidity: function ErrorCode(uint8 ) view returns(string)
func (_BasView *BasViewCallerSession) ErrorCode(arg0 uint8) (string, error) {
	return _BasView.Contract.ErrorCode(&_BasView.CallOpts, arg0)
}

// CheckRootRegistry is a free data retrieval call binding the contract method 0xa84ccbf4.
//
// Solidity: function checkRootRegistry(bytes name, bool isCustom, uint256 cusPrice, uint8 durationInYear) view returns(uint8, bool, uint256)
func (_BasView *BasViewCaller) CheckRootRegistry(opts *bind.CallOpts, name []byte, isCustom bool, cusPrice *big.Int, durationInYear uint8) (uint8, bool, *big.Int, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "checkRootRegistry", name, isCustom, cusPrice, durationInYear)

	if err != nil {
		return *new(uint8), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CheckRootRegistry is a free data retrieval call binding the contract method 0xa84ccbf4.
//
// Solidity: function checkRootRegistry(bytes name, bool isCustom, uint256 cusPrice, uint8 durationInYear) view returns(uint8, bool, uint256)
func (_BasView *BasViewSession) CheckRootRegistry(name []byte, isCustom bool, cusPrice *big.Int, durationInYear uint8) (uint8, bool, *big.Int, error) {
	return _BasView.Contract.CheckRootRegistry(&_BasView.CallOpts, name, isCustom, cusPrice, durationInYear)
}

// CheckRootRegistry is a free data retrieval call binding the contract method 0xa84ccbf4.
//
// Solidity: function checkRootRegistry(bytes name, bool isCustom, uint256 cusPrice, uint8 durationInYear) view returns(uint8, bool, uint256)
func (_BasView *BasViewCallerSession) CheckRootRegistry(name []byte, isCustom bool, cusPrice *big.Int, durationInYear uint8) (uint8, bool, *big.Int, error) {
	return _BasView.Contract.CheckRootRegistry(&_BasView.CallOpts, name, isCustom, cusPrice, durationInYear)
}

// CheckSubRegistry is a free data retrieval call binding the contract method 0xd13b893f.
//
// Solidity: function checkSubRegistry(bytes rName, bytes sName, uint256 durationInYear) view returns(uint8, uint256)
func (_BasView *BasViewCaller) CheckSubRegistry(opts *bind.CallOpts, rName []byte, sName []byte, durationInYear *big.Int) (uint8, *big.Int, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "checkSubRegistry", rName, sName, durationInYear)

	if err != nil {
		return *new(uint8), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CheckSubRegistry is a free data retrieval call binding the contract method 0xd13b893f.
//
// Solidity: function checkSubRegistry(bytes rName, bytes sName, uint256 durationInYear) view returns(uint8, uint256)
func (_BasView *BasViewSession) CheckSubRegistry(rName []byte, sName []byte, durationInYear *big.Int) (uint8, *big.Int, error) {
	return _BasView.Contract.CheckSubRegistry(&_BasView.CallOpts, rName, sName, durationInYear)
}

// CheckSubRegistry is a free data retrieval call binding the contract method 0xd13b893f.
//
// Solidity: function checkSubRegistry(bytes rName, bytes sName, uint256 durationInYear) view returns(uint8, uint256)
func (_BasView *BasViewCallerSession) CheckSubRegistry(rName []byte, sName []byte, durationInYear *big.Int) (uint8, *big.Int, error) {
	return _BasView.Contract.CheckSubRegistry(&_BasView.CallOpts, rName, sName, durationInYear)
}

// DomainIsWild is a free data retrieval call binding the contract method 0x256dec96.
//
// Solidity: function domainIsWild(bytes32 hash) view returns(bool)
func (_BasView *BasViewCaller) DomainIsWild(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "domainIsWild", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DomainIsWild is a free data retrieval call binding the contract method 0x256dec96.
//
// Solidity: function domainIsWild(bytes32 hash) view returns(bool)
func (_BasView *BasViewSession) DomainIsWild(hash [32]byte) (bool, error) {
	return _BasView.Contract.DomainIsWild(&_BasView.CallOpts, hash)
}

// DomainIsWild is a free data retrieval call binding the contract method 0x256dec96.
//
// Solidity: function domainIsWild(bytes32 hash) view returns(bool)
func (_BasView *BasViewCallerSession) DomainIsWild(hash [32]byte) (bool, error) {
	return _BasView.Contract.DomainIsWild(&_BasView.CallOpts, hash)
}

// GetOANNParams is a free data retrieval call binding the contract method 0xff32cd20.
//
// Solidity: function getOANNParams() view returns(uint256 MAX_YEAR, uint256 AROOT_GAS, uint256 BROOT_GAS, uint256 SUB_GAS, uint256 CUSTOM_PRICE_GAS)
func (_BasView *BasViewCaller) GetOANNParams(opts *bind.CallOpts) (struct {
	MAXYEAR        *big.Int
	AROOTGAS       *big.Int
	BROOTGAS       *big.Int
	SUBGAS         *big.Int
	CUSTOMPRICEGAS *big.Int
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "getOANNParams")

	outstruct := new(struct {
		MAXYEAR        *big.Int
		AROOTGAS       *big.Int
		BROOTGAS       *big.Int
		SUBGAS         *big.Int
		CUSTOMPRICEGAS *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MAXYEAR = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AROOTGAS = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BROOTGAS = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SUBGAS = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CUSTOMPRICEGAS = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOANNParams is a free data retrieval call binding the contract method 0xff32cd20.
//
// Solidity: function getOANNParams() view returns(uint256 MAX_YEAR, uint256 AROOT_GAS, uint256 BROOT_GAS, uint256 SUB_GAS, uint256 CUSTOM_PRICE_GAS)
func (_BasView *BasViewSession) GetOANNParams() (struct {
	MAXYEAR        *big.Int
	AROOTGAS       *big.Int
	BROOTGAS       *big.Int
	SUBGAS         *big.Int
	CUSTOMPRICEGAS *big.Int
}, error) {
	return _BasView.Contract.GetOANNParams(&_BasView.CallOpts)
}

// GetOANNParams is a free data retrieval call binding the contract method 0xff32cd20.
//
// Solidity: function getOANNParams() view returns(uint256 MAX_YEAR, uint256 AROOT_GAS, uint256 BROOT_GAS, uint256 SUB_GAS, uint256 CUSTOM_PRICE_GAS)
func (_BasView *BasViewCallerSession) GetOANNParams() (struct {
	MAXYEAR        *big.Int
	AROOTGAS       *big.Int
	BROOTGAS       *big.Int
	SUBGAS         *big.Int
	CUSTOMPRICEGAS *big.Int
}, error) {
	return _BasView.Contract.GetOANNParams(&_BasView.CallOpts)
}

// QueryDomainConfigs is a free data retrieval call binding the contract method 0x4f69b3a2.
//
// Solidity: function queryDomainConfigs(bytes32 nameHash) view returns(bytes A, bytes AAAA, bytes MX, bytes BlockChain, bytes IOTA, bytes CName, bytes MXBCA)
func (_BasView *BasViewCaller) QueryDomainConfigs(opts *bind.CallOpts, nameHash [32]byte) (struct {
	A          []byte
	AAAA       []byte
	MX         []byte
	BlockChain []byte
	IOTA       []byte
	CName      []byte
	MXBCA      []byte
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "queryDomainConfigs", nameHash)

	outstruct := new(struct {
		A          []byte
		AAAA       []byte
		MX         []byte
		BlockChain []byte
		IOTA       []byte
		CName      []byte
		MXBCA      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AAAA = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.MX = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.BlockChain = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.IOTA = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.CName = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.MXBCA = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// QueryDomainConfigs is a free data retrieval call binding the contract method 0x4f69b3a2.
//
// Solidity: function queryDomainConfigs(bytes32 nameHash) view returns(bytes A, bytes AAAA, bytes MX, bytes BlockChain, bytes IOTA, bytes CName, bytes MXBCA)
func (_BasView *BasViewSession) QueryDomainConfigs(nameHash [32]byte) (struct {
	A          []byte
	AAAA       []byte
	MX         []byte
	BlockChain []byte
	IOTA       []byte
	CName      []byte
	MXBCA      []byte
}, error) {
	return _BasView.Contract.QueryDomainConfigs(&_BasView.CallOpts, nameHash)
}

// QueryDomainConfigs is a free data retrieval call binding the contract method 0x4f69b3a2.
//
// Solidity: function queryDomainConfigs(bytes32 nameHash) view returns(bytes A, bytes AAAA, bytes MX, bytes BlockChain, bytes IOTA, bytes CName, bytes MXBCA)
func (_BasView *BasViewCallerSession) QueryDomainConfigs(nameHash [32]byte) (struct {
	A          []byte
	AAAA       []byte
	MX         []byte
	BlockChain []byte
	IOTA       []byte
	CName      []byte
	MXBCA      []byte
}, error) {
	return _BasView.Contract.QueryDomainConfigs(&_BasView.CallOpts, nameHash)
}

// QueryDomainEmailInfo is a free data retrieval call binding the contract method 0x1f766f14.
//
// Solidity: function queryDomainEmailInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isActive, bool openToPublic)
func (_BasView *BasViewCaller) QueryDomainEmailInfo(opts *bind.CallOpts, nameHash [32]byte) (struct {
	Name         []byte
	Owner        common.Address
	Expiration   *big.Int
	IsActive     bool
	OpenToPublic bool
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "queryDomainEmailInfo", nameHash)

	outstruct := new(struct {
		Name         []byte
		Owner        common.Address
		Expiration   *big.Int
		IsActive     bool
		OpenToPublic bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expiration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.OpenToPublic = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// QueryDomainEmailInfo is a free data retrieval call binding the contract method 0x1f766f14.
//
// Solidity: function queryDomainEmailInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isActive, bool openToPublic)
func (_BasView *BasViewSession) QueryDomainEmailInfo(nameHash [32]byte) (struct {
	Name         []byte
	Owner        common.Address
	Expiration   *big.Int
	IsActive     bool
	OpenToPublic bool
}, error) {
	return _BasView.Contract.QueryDomainEmailInfo(&_BasView.CallOpts, nameHash)
}

// QueryDomainEmailInfo is a free data retrieval call binding the contract method 0x1f766f14.
//
// Solidity: function queryDomainEmailInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isActive, bool openToPublic)
func (_BasView *BasViewCallerSession) QueryDomainEmailInfo(nameHash [32]byte) (struct {
	Name         []byte
	Owner        common.Address
	Expiration   *big.Int
	IsActive     bool
	OpenToPublic bool
}, error) {
	return _BasView.Contract.QueryDomainEmailInfo(&_BasView.CallOpts, nameHash)
}

// QueryDomainInfo is a free data retrieval call binding the contract method 0x3870d91e.
//
// Solidity: function queryDomainInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isRoot, bool rIsOpen, bool rIsCustom, bool rIsRare, uint256 rCusPrice, bytes32 sRootHash, bool isMarketOrder)
func (_BasView *BasViewCaller) QueryDomainInfo(opts *bind.CallOpts, nameHash [32]byte) (struct {
	Name          []byte
	Owner         common.Address
	Expiration    *big.Int
	IsRoot        bool
	RIsOpen       bool
	RIsCustom     bool
	RIsRare       bool
	RCusPrice     *big.Int
	SRootHash     [32]byte
	IsMarketOrder bool
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "queryDomainInfo", nameHash)

	outstruct := new(struct {
		Name          []byte
		Owner         common.Address
		Expiration    *big.Int
		IsRoot        bool
		RIsOpen       bool
		RIsCustom     bool
		RIsRare       bool
		RCusPrice     *big.Int
		SRootHash     [32]byte
		IsMarketOrder bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expiration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsRoot = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.RIsOpen = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.RIsCustom = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.RIsRare = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.RCusPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.SRootHash = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.IsMarketOrder = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// QueryDomainInfo is a free data retrieval call binding the contract method 0x3870d91e.
//
// Solidity: function queryDomainInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isRoot, bool rIsOpen, bool rIsCustom, bool rIsRare, uint256 rCusPrice, bytes32 sRootHash, bool isMarketOrder)
func (_BasView *BasViewSession) QueryDomainInfo(nameHash [32]byte) (struct {
	Name          []byte
	Owner         common.Address
	Expiration    *big.Int
	IsRoot        bool
	RIsOpen       bool
	RIsCustom     bool
	RIsRare       bool
	RCusPrice     *big.Int
	SRootHash     [32]byte
	IsMarketOrder bool
}, error) {
	return _BasView.Contract.QueryDomainInfo(&_BasView.CallOpts, nameHash)
}

// QueryDomainInfo is a free data retrieval call binding the contract method 0x3870d91e.
//
// Solidity: function queryDomainInfo(bytes32 nameHash) view returns(bytes name, address owner, uint256 expiration, bool isRoot, bool rIsOpen, bool rIsCustom, bool rIsRare, uint256 rCusPrice, bytes32 sRootHash, bool isMarketOrder)
func (_BasView *BasViewCallerSession) QueryDomainInfo(nameHash [32]byte) (struct {
	Name          []byte
	Owner         common.Address
	Expiration    *big.Int
	IsRoot        bool
	RIsOpen       bool
	RIsCustom     bool
	RIsRare       bool
	RCusPrice     *big.Int
	SRootHash     [32]byte
	IsMarketOrder bool
}, error) {
	return _BasView.Contract.QueryDomainInfo(&_BasView.CallOpts, nameHash)
}

// QueryEmailInfo is a free data retrieval call binding the contract method 0xa806e95c.
//
// Solidity: function queryEmailInfo(bytes32 mailHash) view returns(address owner, uint256 expiration, bytes32 domainHash, bool isValid, bytes aliasName, bytes bcAddress)
func (_BasView *BasViewCaller) QueryEmailInfo(opts *bind.CallOpts, mailHash [32]byte) (struct {
	Owner      common.Address
	Expiration *big.Int
	DomainHash [32]byte
	IsValid    bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "queryEmailInfo", mailHash)

	outstruct := new(struct {
		Owner      common.Address
		Expiration *big.Int
		DomainHash [32]byte
		IsValid    bool
		AliasName  []byte
		BcAddress  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DomainHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IsValid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.AliasName = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.BcAddress = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// QueryEmailInfo is a free data retrieval call binding the contract method 0xa806e95c.
//
// Solidity: function queryEmailInfo(bytes32 mailHash) view returns(address owner, uint256 expiration, bytes32 domainHash, bool isValid, bytes aliasName, bytes bcAddress)
func (_BasView *BasViewSession) QueryEmailInfo(mailHash [32]byte) (struct {
	Owner      common.Address
	Expiration *big.Int
	DomainHash [32]byte
	IsValid    bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	return _BasView.Contract.QueryEmailInfo(&_BasView.CallOpts, mailHash)
}

// QueryEmailInfo is a free data retrieval call binding the contract method 0xa806e95c.
//
// Solidity: function queryEmailInfo(bytes32 mailHash) view returns(address owner, uint256 expiration, bytes32 domainHash, bool isValid, bytes aliasName, bytes bcAddress)
func (_BasView *BasViewCallerSession) QueryEmailInfo(mailHash [32]byte) (struct {
	Owner      common.Address
	Expiration *big.Int
	DomainHash [32]byte
	IsValid    bool
	AliasName  []byte
	BcAddress  []byte
}, error) {
	return _BasView.Contract.QueryEmailInfo(&_BasView.CallOpts, mailHash)
}

// QueryOrderInfo is a free data retrieval call binding the contract method 0xd8d0db68.
//
// Solidity: function queryOrderInfo(address seller, bytes32 nameHash) view returns(bytes name, uint256 price, bool isValid)
func (_BasView *BasViewCaller) QueryOrderInfo(opts *bind.CallOpts, seller common.Address, nameHash [32]byte) (struct {
	Name    []byte
	Price   *big.Int
	IsValid bool
}, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "queryOrderInfo", seller, nameHash)

	outstruct := new(struct {
		Name    []byte
		Price   *big.Int
		IsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// QueryOrderInfo is a free data retrieval call binding the contract method 0xd8d0db68.
//
// Solidity: function queryOrderInfo(address seller, bytes32 nameHash) view returns(bytes name, uint256 price, bool isValid)
func (_BasView *BasViewSession) QueryOrderInfo(seller common.Address, nameHash [32]byte) (struct {
	Name    []byte
	Price   *big.Int
	IsValid bool
}, error) {
	return _BasView.Contract.QueryOrderInfo(&_BasView.CallOpts, seller, nameHash)
}

// QueryOrderInfo is a free data retrieval call binding the contract method 0xd8d0db68.
//
// Solidity: function queryOrderInfo(address seller, bytes32 nameHash) view returns(bytes name, uint256 price, bool isValid)
func (_BasView *BasViewCallerSession) QueryOrderInfo(seller common.Address, nameHash [32]byte) (struct {
	Name    []byte
	Price   *big.Int
	IsValid bool
}, error) {
	return _BasView.Contract.QueryOrderInfo(&_BasView.CallOpts, seller, nameHash)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasView *BasViewCaller) Rel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasView.contract.Call(opts, &out, "rel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasView *BasViewSession) Rel() (common.Address, error) {
	return _BasView.Contract.Rel(&_BasView.CallOpts)
}

// Rel is a free data retrieval call binding the contract method 0xce26e78a.
//
// Solidity: function rel() view returns(address)
func (_BasView *BasViewCallerSession) Rel() (common.Address, error) {
	return _BasView.Contract.Rel(&_BasView.CallOpts)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasView *BasViewTransactor) ChangeDAO(opts *bind.TransactOpts, newDao common.Address) (*types.Transaction, error) {
	return _BasView.contract.Transact(opts, "ChangeDAO", newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasView *BasViewSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasView.Contract.ChangeDAO(&_BasView.TransactOpts, newDao)
}

// ChangeDAO is a paid mutator transaction binding the contract method 0x8a42876b.
//
// Solidity: function ChangeDAO(address newDao) returns()
func (_BasView *BasViewTransactorSession) ChangeDAO(newDao common.Address) (*types.Transaction, error) {
	return _BasView.Contract.ChangeDAO(&_BasView.TransactOpts, newDao)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasView *BasViewTransactor) ChangeRelation(opts *bind.TransactOpts, new_rel common.Address) (*types.Transaction, error) {
	return _BasView.contract.Transact(opts, "changeRelation", new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasView *BasViewSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasView.Contract.ChangeRelation(&_BasView.TransactOpts, new_rel)
}

// ChangeRelation is a paid mutator transaction binding the contract method 0x57b29ef4.
//
// Solidity: function changeRelation(address new_rel) returns()
func (_BasView *BasViewTransactorSession) ChangeRelation(new_rel common.Address) (*types.Transaction, error) {
	return _BasView.Contract.ChangeRelation(&_BasView.TransactOpts, new_rel)
}

// SetErrorCode is a paid mutator transaction binding the contract method 0x4411c158.
//
// Solidity: function setErrorCode(uint8 index, string reason) returns()
func (_BasView *BasViewTransactor) SetErrorCode(opts *bind.TransactOpts, index uint8, reason string) (*types.Transaction, error) {
	return _BasView.contract.Transact(opts, "setErrorCode", index, reason)
}

// SetErrorCode is a paid mutator transaction binding the contract method 0x4411c158.
//
// Solidity: function setErrorCode(uint8 index, string reason) returns()
func (_BasView *BasViewSession) SetErrorCode(index uint8, reason string) (*types.Transaction, error) {
	return _BasView.Contract.SetErrorCode(&_BasView.TransactOpts, index, reason)
}

// SetErrorCode is a paid mutator transaction binding the contract method 0x4411c158.
//
// Solidity: function setErrorCode(uint8 index, string reason) returns()
func (_BasView *BasViewTransactorSession) SetErrorCode(index uint8, reason string) (*types.Transaction, error) {
	return _BasView.Contract.SetErrorCode(&_BasView.TransactOpts, index, reason)
}
