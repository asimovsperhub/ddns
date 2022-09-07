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

// DnsSubNameABI is the input ABI used to generate the binding from.
const DnsSubNameABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"taxC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountC_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nameOwner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerC_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"}],\"name\":\"EvAddSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transfer\",\"type\":\"bool\"}],\"name\":\"EvChargeSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"}],\"name\":\"EvDelSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"}],\"name\":\"EvMintSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"transfer_\",\"type\":\"bool\"}],\"name\":\"ChargeName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"MintName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"passId_\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721Owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fatherAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fatherHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gNameId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"id2Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"entireName\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sigUser_\",\"type\":\"address\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DnsSubName is an auto generated Go binding around an Ethereum contract.
type DnsSubName struct {
	DnsSubNameCaller     // Read-only binding to the contract
	DnsSubNameTransactor // Write-only binding to the contract
	DnsSubNameFilterer   // Log filterer for contract events
}

// DnsSubNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsSubNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSubNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsSubNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSubNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsSubNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSubNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsSubNameSession struct {
	Contract     *DnsSubName       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsSubNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsSubNameCallerSession struct {
	Contract *DnsSubNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DnsSubNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsSubNameTransactorSession struct {
	Contract     *DnsSubNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DnsSubNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsSubNameRaw struct {
	Contract *DnsSubName // Generic contract binding to access the raw methods on
}

// DnsSubNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsSubNameCallerRaw struct {
	Contract *DnsSubNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsSubNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsSubNameTransactorRaw struct {
	Contract *DnsSubNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsSubName creates a new instance of DnsSubName, bound to a specific deployed contract.
func NewDnsSubName(address common.Address, backend bind.ContractBackend) (*DnsSubName, error) {
	contract, err := bindDnsSubName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsSubName{DnsSubNameCaller: DnsSubNameCaller{contract: contract}, DnsSubNameTransactor: DnsSubNameTransactor{contract: contract}, DnsSubNameFilterer: DnsSubNameFilterer{contract: contract}}, nil
}

// NewDnsSubNameCaller creates a new read-only instance of DnsSubName, bound to a specific deployed contract.
func NewDnsSubNameCaller(address common.Address, caller bind.ContractCaller) (*DnsSubNameCaller, error) {
	contract, err := bindDnsSubName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameCaller{contract: contract}, nil
}

// NewDnsSubNameTransactor creates a new write-only instance of DnsSubName, bound to a specific deployed contract.
func NewDnsSubNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsSubNameTransactor, error) {
	contract, err := bindDnsSubName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameTransactor{contract: contract}, nil
}

// NewDnsSubNameFilterer creates a new log filterer instance of DnsSubName, bound to a specific deployed contract.
func NewDnsSubNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsSubNameFilterer, error) {
	contract, err := bindDnsSubName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameFilterer{contract: contract}, nil
}

// bindDnsSubName binds a generic wrapper to an already deployed contract.
func bindDnsSubName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsSubNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSubName *DnsSubNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSubName.Contract.DnsSubNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSubName *DnsSubNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSubName.Contract.DnsSubNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSubName *DnsSubNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSubName.Contract.DnsSubNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSubName *DnsSubNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSubName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSubName *DnsSubNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSubName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSubName *DnsSubNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSubName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSubName *DnsSubNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSubName *DnsSubNameSession) AccountantC() (common.Address, error) {
	return _DnsSubName.Contract.AccountantC(&_DnsSubName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsSubName.Contract.AccountantC(&_DnsSubName.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsSubName *DnsSubNameCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsSubName *DnsSubNameSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsSubName.Contract.BalanceOf(&_DnsSubName.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsSubName *DnsSubNameCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsSubName.Contract.BalanceOf(&_DnsSubName.CallOpts, owner)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsSubName *DnsSubNameCaller) Erc721Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "erc721Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsSubName *DnsSubNameSession) Erc721Owner() (common.Address, error) {
	return _DnsSubName.Contract.Erc721Owner(&_DnsSubName.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) Erc721Owner() (common.Address, error) {
	return _DnsSubName.Contract.Erc721Owner(&_DnsSubName.CallOpts)
}

// FatherAddr is a free data retrieval call binding the contract method 0xb6bde99f.
//
// Solidity: function fatherAddr() view returns(address)
func (_DnsSubName *DnsSubNameCaller) FatherAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "fatherAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FatherAddr is a free data retrieval call binding the contract method 0xb6bde99f.
//
// Solidity: function fatherAddr() view returns(address)
func (_DnsSubName *DnsSubNameSession) FatherAddr() (common.Address, error) {
	return _DnsSubName.Contract.FatherAddr(&_DnsSubName.CallOpts)
}

// FatherAddr is a free data retrieval call binding the contract method 0xb6bde99f.
//
// Solidity: function fatherAddr() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) FatherAddr() (common.Address, error) {
	return _DnsSubName.Contract.FatherAddr(&_DnsSubName.CallOpts)
}

// FatherHash is a free data retrieval call binding the contract method 0x9ca50127.
//
// Solidity: function fatherHash() view returns(bytes32)
func (_DnsSubName *DnsSubNameCaller) FatherHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "fatherHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FatherHash is a free data retrieval call binding the contract method 0x9ca50127.
//
// Solidity: function fatherHash() view returns(bytes32)
func (_DnsSubName *DnsSubNameSession) FatherHash() ([32]byte, error) {
	return _DnsSubName.Contract.FatherHash(&_DnsSubName.CallOpts)
}

// FatherHash is a free data retrieval call binding the contract method 0x9ca50127.
//
// Solidity: function fatherHash() view returns(bytes32)
func (_DnsSubName *DnsSubNameCallerSession) FatherHash() ([32]byte, error) {
	return _DnsSubName.Contract.FatherHash(&_DnsSubName.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsSubName *DnsSubNameCaller) GNameId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "gNameId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsSubName *DnsSubNameSession) GNameId() (*big.Int, error) {
	return _DnsSubName.Contract.GNameId(&_DnsSubName.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsSubName *DnsSubNameCallerSession) GNameId() (*big.Int, error) {
	return _DnsSubName.Contract.GNameId(&_DnsSubName.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsSubName.Contract.GetApproved(&_DnsSubName.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsSubName.Contract.GetApproved(&_DnsSubName.CallOpts, tokenId)
}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsSubName *DnsSubNameCaller) Id2Hash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "id2Hash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsSubName *DnsSubNameSession) Id2Hash(arg0 *big.Int) ([32]byte, error) {
	return _DnsSubName.Contract.Id2Hash(&_DnsSubName.CallOpts, arg0)
}

// Id2Hash is a free data retrieval call binding the contract method 0xfed15642.
//
// Solidity: function id2Hash(uint256 ) view returns(bytes32)
func (_DnsSubName *DnsSubNameCallerSession) Id2Hash(arg0 *big.Int) ([32]byte, error) {
	return _DnsSubName.Contract.Id2Hash(&_DnsSubName.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DnsSubName *DnsSubNameCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DnsSubName *DnsSubNameSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DnsSubName.Contract.IsApprovedForAll(&_DnsSubName.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DnsSubName *DnsSubNameCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DnsSubName.Contract.IsApprovedForAll(&_DnsSubName.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsSubName *DnsSubNameCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsSubName *DnsSubNameSession) Name() (string, error) {
	return _DnsSubName.Contract.Name(&_DnsSubName.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsSubName *DnsSubNameCallerSession) Name() (string, error) {
	return _DnsSubName.Contract.Name(&_DnsSubName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, uint256 expireTime, uint256 tokenId)
func (_DnsSubName *DnsSubNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EntireName []byte
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "nameStore", arg0)

	outstruct := new(struct {
		EntireName []byte
		ExpireTime *big.Int
		TokenId    *big.Int
	})

	outstruct.EntireName = out[0].([]byte)
	outstruct.ExpireTime = out[1].(*big.Int)
	outstruct.TokenId = out[2].(*big.Int)

	return *outstruct, err

}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, uint256 expireTime, uint256 tokenId)
func (_DnsSubName *DnsSubNameSession) NameStore(arg0 [32]byte) (struct {
	EntireName []byte
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	return _DnsSubName.Contract.NameStore(&_DnsSubName.CallOpts, arg0)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(bytes entireName, uint256 expireTime, uint256 tokenId)
func (_DnsSubName *DnsSubNameCallerSession) NameStore(arg0 [32]byte) (struct {
	EntireName []byte
	ExpireTime *big.Int
	TokenId    *big.Int
}, error) {
	return _DnsSubName.Contract.NameStore(&_DnsSubName.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsSubName *DnsSubNameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsSubName *DnsSubNameSession) Owner() (common.Address, error) {
	return _DnsSubName.Contract.Owner(&_DnsSubName.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) Owner() (common.Address, error) {
	return _DnsSubName.Contract.Owner(&_DnsSubName.CallOpts)
}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsSubName *DnsSubNameCaller) OwnerC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "ownerC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsSubName *DnsSubNameSession) OwnerC() (common.Address, error) {
	return _DnsSubName.Contract.OwnerC(&_DnsSubName.CallOpts)
}

// OwnerC is a free data retrieval call binding the contract method 0xc9b9a4bc.
//
// Solidity: function ownerC() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) OwnerC() (common.Address, error) {
	return _DnsSubName.Contract.OwnerC(&_DnsSubName.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsSubName.Contract.OwnerOf(&_DnsSubName.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsSubName.Contract.OwnerOf(&_DnsSubName.CallOpts, tokenId)
}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsSubName *DnsSubNameCaller) PriceC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "priceC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsSubName *DnsSubNameSession) PriceC() (common.Address, error) {
	return _DnsSubName.Contract.PriceC(&_DnsSubName.CallOpts)
}

// PriceC is a free data retrieval call binding the contract method 0xe3474ccf.
//
// Solidity: function priceC() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) PriceC() (common.Address, error) {
	return _DnsSubName.Contract.PriceC(&_DnsSubName.CallOpts)
}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsSubName *DnsSubNameCaller) SigUser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "sigUser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsSubName *DnsSubNameSession) SigUser() (common.Address, error) {
	return _DnsSubName.Contract.SigUser(&_DnsSubName.CallOpts)
}

// SigUser is a free data retrieval call binding the contract method 0x9ff00963.
//
// Solidity: function sigUser() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) SigUser() (common.Address, error) {
	return _DnsSubName.Contract.SigUser(&_DnsSubName.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsSubName *DnsSubNameCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsSubName *DnsSubNameSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsSubName.Contract.SupportsInterface(&_DnsSubName.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsSubName *DnsSubNameCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsSubName.Contract.SupportsInterface(&_DnsSubName.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsSubName *DnsSubNameCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsSubName *DnsSubNameSession) Symbol() (string, error) {
	return _DnsSubName.Contract.Symbol(&_DnsSubName.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsSubName *DnsSubNameCallerSession) Symbol() (string, error) {
	return _DnsSubName.Contract.Symbol(&_DnsSubName.CallOpts)
}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsSubName *DnsSubNameCaller) TaxC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "taxC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsSubName *DnsSubNameSession) TaxC() (common.Address, error) {
	return _DnsSubName.Contract.TaxC(&_DnsSubName.CallOpts)
}

// TaxC is a free data retrieval call binding the contract method 0x4a0e4d2a.
//
// Solidity: function taxC() view returns(address)
func (_DnsSubName *DnsSubNameCallerSession) TaxC() (common.Address, error) {
	return _DnsSubName.Contract.TaxC(&_DnsSubName.CallOpts)
}

// TaxPrice is a free data retrieval call binding the contract method 0xeda5745c.
//
// Solidity: function taxPrice() view returns(uint256)
func (_DnsSubName *DnsSubNameCaller) TaxPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "taxPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxPrice is a free data retrieval call binding the contract method 0xeda5745c.
//
// Solidity: function taxPrice() view returns(uint256)
func (_DnsSubName *DnsSubNameSession) TaxPrice() (*big.Int, error) {
	return _DnsSubName.Contract.TaxPrice(&_DnsSubName.CallOpts)
}

// TaxPrice is a free data retrieval call binding the contract method 0xeda5745c.
//
// Solidity: function taxPrice() view returns(uint256)
func (_DnsSubName *DnsSubNameCallerSession) TaxPrice() (*big.Int, error) {
	return _DnsSubName.Contract.TaxPrice(&_DnsSubName.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsSubName *DnsSubNameCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DnsSubName.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsSubName *DnsSubNameSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsSubName.Contract.TokenURI(&_DnsSubName.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsSubName *DnsSubNameCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsSubName.Contract.TokenURI(&_DnsSubName.CallOpts, tokenId)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsSubName *DnsSubNameTransactor) ChargeName(opts *bind.TransactOpts, erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "ChargeName", erc20Addr_, nameHash_, year_, transfer_)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsSubName *DnsSubNameSession) ChargeName(erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsSubName.Contract.ChargeName(&_DnsSubName.TransactOpts, erc20Addr_, nameHash_, year_, transfer_)
}

// ChargeName is a paid mutator transaction binding the contract method 0xc70f923c.
//
// Solidity: function ChargeName(address erc20Addr_, bytes32 nameHash_, uint8 year_, bool transfer_) payable returns()
func (_DnsSubName *DnsSubNameTransactorSession) ChargeName(erc20Addr_ common.Address, nameHash_ [32]byte, year_ uint8, transfer_ bool) (*types.Transaction, error) {
	return _DnsSubName.Contract.ChargeName(&_DnsSubName.TransactOpts, erc20Addr_, nameHash_, year_, transfer_)
}

// MintName is a paid mutator transaction binding the contract method 0xcd9481af.
//
// Solidity: function MintName(address erc20Addr_, string entireName_, uint8 year_) payable returns()
func (_DnsSubName *DnsSubNameTransactor) MintName(opts *bind.TransactOpts, erc20Addr_ common.Address, entireName_ string, year_ uint8) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "MintName", erc20Addr_, entireName_, year_)
}

// MintName is a paid mutator transaction binding the contract method 0xcd9481af.
//
// Solidity: function MintName(address erc20Addr_, string entireName_, uint8 year_) payable returns()
func (_DnsSubName *DnsSubNameSession) MintName(erc20Addr_ common.Address, entireName_ string, year_ uint8) (*types.Transaction, error) {
	return _DnsSubName.Contract.MintName(&_DnsSubName.TransactOpts, erc20Addr_, entireName_, year_)
}

// MintName is a paid mutator transaction binding the contract method 0xcd9481af.
//
// Solidity: function MintName(address erc20Addr_, string entireName_, uint8 year_) payable returns()
func (_DnsSubName *DnsSubNameTransactorSession) MintName(erc20Addr_ common.Address, entireName_ string, year_ uint8) (*types.Transaction, error) {
	return _DnsSubName.Contract.MintName(&_DnsSubName.TransactOpts, erc20Addr_, entireName_, year_)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsSubName *DnsSubNameTransactor) MintNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "MintNameBySig", entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsSubName *DnsSubNameSession) MintNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsSubName.Contract.MintNameBySig(&_DnsSubName.TransactOpts, entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// MintNameBySig is a paid mutator transaction binding the contract method 0xe0c52d70.
//
// Solidity: function MintNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint256 price_, uint32 passId_, bytes sig) payable returns()
func (_DnsSubName *DnsSubNameTransactorSession) MintNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, price_ *big.Int, passId_ uint32, sig []byte) (*types.Transaction, error) {
	return _DnsSubName.Contract.MintNameBySig(&_DnsSubName.TransactOpts, entireName_, year_, erc20Addr_, price_, passId_, sig)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.Approve(&_DnsSubName.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.Approve(&_DnsSubName.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.SafeTransferFrom(&_DnsSubName.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.SafeTransferFrom(&_DnsSubName.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsSubName *DnsSubNameTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsSubName *DnsSubNameSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsSubName.Contract.SafeTransferFrom0(&_DnsSubName.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsSubName *DnsSubNameTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsSubName.Contract.SafeTransferFrom0(&_DnsSubName.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DnsSubName *DnsSubNameTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DnsSubName *DnsSubNameSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetApprovalForAll(&_DnsSubName.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DnsSubName *DnsSubNameTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetApprovalForAll(&_DnsSubName.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xb470e164.
//
// Solidity: function setBaseUri(string baseUri_, address sigUser_) returns()
func (_DnsSubName *DnsSubNameTransactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string, sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "setBaseUri", baseUri_, sigUser_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xb470e164.
//
// Solidity: function setBaseUri(string baseUri_, address sigUser_) returns()
func (_DnsSubName *DnsSubNameSession) SetBaseUri(baseUri_ string, sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetBaseUri(&_DnsSubName.TransactOpts, baseUri_, sigUser_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xb470e164.
//
// Solidity: function setBaseUri(string baseUri_, address sigUser_) returns()
func (_DnsSubName *DnsSubNameTransactorSession) SetBaseUri(baseUri_ string, sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetBaseUri(&_DnsSubName.TransactOpts, baseUri_, sigUser_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsSubName *DnsSubNameTransactor) SetName(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "setName", name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsSubName *DnsSubNameSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetName(&_DnsSubName.TransactOpts, name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsSubName *DnsSubNameTransactorSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsSubName.Contract.SetName(&_DnsSubName.TransactOpts, name_, symbol_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.TransferFrom(&_DnsSubName.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsSubName *DnsSubNameTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsSubName.Contract.TransferFrom(&_DnsSubName.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsSubName *DnsSubNameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsSubName.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsSubName *DnsSubNameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsSubName.Contract.TransferOwnership(&_DnsSubName.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsSubName *DnsSubNameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsSubName.Contract.TransferOwnership(&_DnsSubName.TransactOpts, newOwner)
}

// DnsSubNameApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DnsSubName contract.
type DnsSubNameApprovalIterator struct {
	Event *DnsSubNameApproval // Event containing the contract specifics and raw log

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
func (it *DnsSubNameApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameApproval)
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
		it.Event = new(DnsSubNameApproval)
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
func (it *DnsSubNameApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameApproval represents a Approval event raised by the DnsSubName contract.
type DnsSubNameApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DnsSubNameApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameApprovalIterator{contract: _DnsSubName.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DnsSubNameApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameApproval)
				if err := _DnsSubName.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) ParseApproval(log types.Log) (*DnsSubNameApproval, error) {
	event := new(DnsSubNameApproval)
	if err := _DnsSubName.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DnsSubName contract.
type DnsSubNameApprovalForAllIterator struct {
	Event *DnsSubNameApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DnsSubNameApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameApprovalForAll)
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
		it.Event = new(DnsSubNameApprovalForAll)
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
func (it *DnsSubNameApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameApprovalForAll represents a ApprovalForAll event raised by the DnsSubName contract.
type DnsSubNameApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsSubName *DnsSubNameFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DnsSubNameApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameApprovalForAllIterator{contract: _DnsSubName.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsSubName *DnsSubNameFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DnsSubNameApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameApprovalForAll)
				if err := _DnsSubName.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsSubName *DnsSubNameFilterer) ParseApprovalForAll(log types.Log) (*DnsSubNameApprovalForAll, error) {
	event := new(DnsSubNameApprovalForAll)
	if err := _DnsSubName.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameEvAddSubNameIterator is returned from FilterEvAddSubName and is used to iterate over the raw logs and unpacked data for EvAddSubName events raised by the DnsSubName contract.
type DnsSubNameEvAddSubNameIterator struct {
	Event *DnsSubNameEvAddSubName // Event containing the contract specifics and raw log

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
func (it *DnsSubNameEvAddSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameEvAddSubName)
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
		it.Event = new(DnsSubNameEvAddSubName)
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
func (it *DnsSubNameEvAddSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameEvAddSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameEvAddSubName represents a EvAddSubName event raised by the DnsSubName contract.
type DnsSubNameEvAddSubName struct {
	Hash       [32]byte
	EntireName string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvAddSubName is a free log retrieval operation binding the contract event 0x76153889eb916e697088a14740d8b847b318607ce85a19cec53251d68cfe967d.
//
// Solidity: event EvAddSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) FilterEvAddSubName(opts *bind.FilterOpts) (*DnsSubNameEvAddSubNameIterator, error) {

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSubNameEvAddSubNameIterator{contract: _DnsSubName.contract, event: "EvAddSubName", logs: logs, sub: sub}, nil
}

// WatchEvAddSubName is a free log subscription operation binding the contract event 0x76153889eb916e697088a14740d8b847b318607ce85a19cec53251d68cfe967d.
//
// Solidity: event EvAddSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) WatchEvAddSubName(opts *bind.WatchOpts, sink chan<- *DnsSubNameEvAddSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameEvAddSubName)
				if err := _DnsSubName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
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

// ParseEvAddSubName is a log parse operation binding the contract event 0x76153889eb916e697088a14740d8b847b318607ce85a19cec53251d68cfe967d.
//
// Solidity: event EvAddSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) ParseEvAddSubName(log types.Log) (*DnsSubNameEvAddSubName, error) {
	event := new(DnsSubNameEvAddSubName)
	if err := _DnsSubName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameEvChargeSubNameIterator is returned from FilterEvChargeSubName and is used to iterate over the raw logs and unpacked data for EvChargeSubName events raised by the DnsSubName contract.
type DnsSubNameEvChargeSubNameIterator struct {
	Event *DnsSubNameEvChargeSubName // Event containing the contract specifics and raw log

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
func (it *DnsSubNameEvChargeSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameEvChargeSubName)
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
		it.Event = new(DnsSubNameEvChargeSubName)
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
func (it *DnsSubNameEvChargeSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameEvChargeSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameEvChargeSubName represents a EvChargeSubName event raised by the DnsSubName contract.
type DnsSubNameEvChargeSubName struct {
	Erc20Addr common.Address
	NameHash  [32]byte
	Year      uint8
	Transfer  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvChargeSubName is a free log retrieval operation binding the contract event 0xa132a141e1df56071ad0482b51adbb87a64ac2d795ae6a93a16fc46d96388c84.
//
// Solidity: event EvChargeSubName(address erc20Addr, bytes32 nameHash, uint8 year, bool transfer)
func (_DnsSubName *DnsSubNameFilterer) FilterEvChargeSubName(opts *bind.FilterOpts) (*DnsSubNameEvChargeSubNameIterator, error) {

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "EvChargeSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSubNameEvChargeSubNameIterator{contract: _DnsSubName.contract, event: "EvChargeSubName", logs: logs, sub: sub}, nil
}

// WatchEvChargeSubName is a free log subscription operation binding the contract event 0xa132a141e1df56071ad0482b51adbb87a64ac2d795ae6a93a16fc46d96388c84.
//
// Solidity: event EvChargeSubName(address erc20Addr, bytes32 nameHash, uint8 year, bool transfer)
func (_DnsSubName *DnsSubNameFilterer) WatchEvChargeSubName(opts *bind.WatchOpts, sink chan<- *DnsSubNameEvChargeSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "EvChargeSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameEvChargeSubName)
				if err := _DnsSubName.contract.UnpackLog(event, "EvChargeSubName", log); err != nil {
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

// ParseEvChargeSubName is a log parse operation binding the contract event 0xa132a141e1df56071ad0482b51adbb87a64ac2d795ae6a93a16fc46d96388c84.
//
// Solidity: event EvChargeSubName(address erc20Addr, bytes32 nameHash, uint8 year, bool transfer)
func (_DnsSubName *DnsSubNameFilterer) ParseEvChargeSubName(log types.Log) (*DnsSubNameEvChargeSubName, error) {
	event := new(DnsSubNameEvChargeSubName)
	if err := _DnsSubName.contract.UnpackLog(event, "EvChargeSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameEvDelSubNameIterator is returned from FilterEvDelSubName and is used to iterate over the raw logs and unpacked data for EvDelSubName events raised by the DnsSubName contract.
type DnsSubNameEvDelSubNameIterator struct {
	Event *DnsSubNameEvDelSubName // Event containing the contract specifics and raw log

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
func (it *DnsSubNameEvDelSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameEvDelSubName)
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
		it.Event = new(DnsSubNameEvDelSubName)
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
func (it *DnsSubNameEvDelSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameEvDelSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameEvDelSubName represents a EvDelSubName event raised by the DnsSubName contract.
type DnsSubNameEvDelSubName struct {
	Hash       [32]byte
	EntireName string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvDelSubName is a free log retrieval operation binding the contract event 0x61093f005e4ab7e3d32e6919e7bc279639e2a1a2137ec674aa4695b9f01a791b.
//
// Solidity: event EvDelSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) FilterEvDelSubName(opts *bind.FilterOpts) (*DnsSubNameEvDelSubNameIterator, error) {

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSubNameEvDelSubNameIterator{contract: _DnsSubName.contract, event: "EvDelSubName", logs: logs, sub: sub}, nil
}

// WatchEvDelSubName is a free log subscription operation binding the contract event 0x61093f005e4ab7e3d32e6919e7bc279639e2a1a2137ec674aa4695b9f01a791b.
//
// Solidity: event EvDelSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) WatchEvDelSubName(opts *bind.WatchOpts, sink chan<- *DnsSubNameEvDelSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameEvDelSubName)
				if err := _DnsSubName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
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

// ParseEvDelSubName is a log parse operation binding the contract event 0x61093f005e4ab7e3d32e6919e7bc279639e2a1a2137ec674aa4695b9f01a791b.
//
// Solidity: event EvDelSubName(bytes32 hash, string entireName)
func (_DnsSubName *DnsSubNameFilterer) ParseEvDelSubName(log types.Log) (*DnsSubNameEvDelSubName, error) {
	event := new(DnsSubNameEvDelSubName)
	if err := _DnsSubName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameEvMintSubNameIterator is returned from FilterEvMintSubName and is used to iterate over the raw logs and unpacked data for EvMintSubName events raised by the DnsSubName contract.
type DnsSubNameEvMintSubNameIterator struct {
	Event *DnsSubNameEvMintSubName // Event containing the contract specifics and raw log

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
func (it *DnsSubNameEvMintSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameEvMintSubName)
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
		it.Event = new(DnsSubNameEvMintSubName)
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
func (it *DnsSubNameEvMintSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameEvMintSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameEvMintSubName represents a EvMintSubName event raised by the DnsSubName contract.
type DnsSubNameEvMintSubName struct {
	Erc20Addr  common.Address
	EntireName string
	Year       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintSubName is a free log retrieval operation binding the contract event 0x35319353f0dd31d02aed296e3700217f1dad39ae7dc87a3958dd62045cbf3e75.
//
// Solidity: event EvMintSubName(address erc20Addr, string entireName, uint8 year)
func (_DnsSubName *DnsSubNameFilterer) FilterEvMintSubName(opts *bind.FilterOpts) (*DnsSubNameEvMintSubNameIterator, error) {

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "EvMintSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSubNameEvMintSubNameIterator{contract: _DnsSubName.contract, event: "EvMintSubName", logs: logs, sub: sub}, nil
}

// WatchEvMintSubName is a free log subscription operation binding the contract event 0x35319353f0dd31d02aed296e3700217f1dad39ae7dc87a3958dd62045cbf3e75.
//
// Solidity: event EvMintSubName(address erc20Addr, string entireName, uint8 year)
func (_DnsSubName *DnsSubNameFilterer) WatchEvMintSubName(opts *bind.WatchOpts, sink chan<- *DnsSubNameEvMintSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "EvMintSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameEvMintSubName)
				if err := _DnsSubName.contract.UnpackLog(event, "EvMintSubName", log); err != nil {
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

// ParseEvMintSubName is a log parse operation binding the contract event 0x35319353f0dd31d02aed296e3700217f1dad39ae7dc87a3958dd62045cbf3e75.
//
// Solidity: event EvMintSubName(address erc20Addr, string entireName, uint8 year)
func (_DnsSubName *DnsSubNameFilterer) ParseEvMintSubName(log types.Log) (*DnsSubNameEvMintSubName, error) {
	event := new(DnsSubNameEvMintSubName)
	if err := _DnsSubName.contract.UnpackLog(event, "EvMintSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSubNameTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DnsSubName contract.
type DnsSubNameTransferIterator struct {
	Event *DnsSubNameTransfer // Event containing the contract specifics and raw log

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
func (it *DnsSubNameTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSubNameTransfer)
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
		it.Event = new(DnsSubNameTransfer)
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
func (it *DnsSubNameTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSubNameTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSubNameTransfer represents a Transfer event raised by the DnsSubName contract.
type DnsSubNameTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DnsSubNameTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsSubName.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsSubNameTransferIterator{contract: _DnsSubName.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DnsSubNameTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DnsSubName.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSubNameTransfer)
				if err := _DnsSubName.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsSubName *DnsSubNameFilterer) ParseTransfer(log types.Log) (*DnsSubNameTransfer, error) {
	event := new(DnsSubNameTransfer)
	if err := _DnsSubName.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
