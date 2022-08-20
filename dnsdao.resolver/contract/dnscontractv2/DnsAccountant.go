// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package udidc

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

// DnsAccountantMetaData contains all meta data concerning the DnsAccountant contract.
var DnsAccountantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerC_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EvDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"out\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"EvWithdrawCash\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"closeMultiSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"delSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"getAllTask\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"idx_\",\"type\":\"uint256\"}],\"name\":\"getContractList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[8]\",\"name\":\"\",\"type\":\"address[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"candidate_\",\"type\":\"address\"}],\"name\":\"getSafeSig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"getSignerSetAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cnt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"getTaskSignerSetAddress\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"multiSignerStore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"work\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"lock\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"SignerCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"openMultiSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"name\":\"removeSafeSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferDaoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"out_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"name\":\"withdrawCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DnsAccountantABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsAccountantMetaData.ABI instead.
var DnsAccountantABI = DnsAccountantMetaData.ABI

// DnsAccountant is an auto generated Go binding around an Ethereum contract.
type DnsAccountant struct {
	DnsAccountantCaller     // Read-only binding to the contract
	DnsAccountantTransactor // Write-only binding to the contract
	DnsAccountantFilterer   // Log filterer for contract events
}

// DnsAccountantCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsAccountantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsAccountantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsAccountantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsAccountantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsAccountantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsAccountantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsAccountantSession struct {
	Contract     *DnsAccountant    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsAccountantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsAccountantCallerSession struct {
	Contract *DnsAccountantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DnsAccountantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsAccountantTransactorSession struct {
	Contract     *DnsAccountantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DnsAccountantRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsAccountantRaw struct {
	Contract *DnsAccountant // Generic contract binding to access the raw methods on
}

// DnsAccountantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsAccountantCallerRaw struct {
	Contract *DnsAccountantCaller // Generic read-only contract binding to access the raw methods on
}

// DnsAccountantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsAccountantTransactorRaw struct {
	Contract *DnsAccountantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsAccountant creates a new instance of DnsAccountant, bound to a specific deployed contract.
func NewDnsAccountant(address common.Address, backend bind.ContractBackend) (*DnsAccountant, error) {
	contract, err := bindDnsAccountant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsAccountant{DnsAccountantCaller: DnsAccountantCaller{contract: contract}, DnsAccountantTransactor: DnsAccountantTransactor{contract: contract}, DnsAccountantFilterer: DnsAccountantFilterer{contract: contract}}, nil
}

// NewDnsAccountantCaller creates a new read-only instance of DnsAccountant, bound to a specific deployed contract.
func NewDnsAccountantCaller(address common.Address, caller bind.ContractCaller) (*DnsAccountantCaller, error) {
	contract, err := bindDnsAccountant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsAccountantCaller{contract: contract}, nil
}

// NewDnsAccountantTransactor creates a new write-only instance of DnsAccountant, bound to a specific deployed contract.
func NewDnsAccountantTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsAccountantTransactor, error) {
	contract, err := bindDnsAccountant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsAccountantTransactor{contract: contract}, nil
}

// NewDnsAccountantFilterer creates a new log filterer instance of DnsAccountant, bound to a specific deployed contract.
func NewDnsAccountantFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsAccountantFilterer, error) {
	contract, err := bindDnsAccountant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsAccountantFilterer{contract: contract}, nil
}

// bindDnsAccountant binds a generic wrapper to an already deployed contract.
func bindDnsAccountant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsAccountantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsAccountant *DnsAccountantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsAccountant.Contract.DnsAccountantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsAccountant *DnsAccountantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsAccountant.Contract.DnsAccountantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsAccountant *DnsAccountantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsAccountant.Contract.DnsAccountantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsAccountant *DnsAccountantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsAccountant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsAccountant *DnsAccountantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsAccountant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsAccountant *DnsAccountantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsAccountant.Contract.contract.Transact(opts, method, params...)
}

// ContractList is a free data retrieval call binding the contract method 0x741489b7.
//
// Solidity: function contractList(address , uint256 ) view returns(address)
func (_DnsAccountant *DnsAccountantCaller) ContractList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "contractList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractList is a free data retrieval call binding the contract method 0x741489b7.
//
// Solidity: function contractList(address , uint256 ) view returns(address)
func (_DnsAccountant *DnsAccountantSession) ContractList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DnsAccountant.Contract.ContractList(&_DnsAccountant.CallOpts, arg0, arg1)
}

// ContractList is a free data retrieval call binding the contract method 0x741489b7.
//
// Solidity: function contractList(address , uint256 ) view returns(address)
func (_DnsAccountant *DnsAccountantCallerSession) ContractList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DnsAccountant.Contract.ContractList(&_DnsAccountant.CallOpts, arg0, arg1)
}

// DefaultCount is a free data retrieval call binding the contract method 0xe8fe7bff.
//
// Solidity: function defaultCount() view returns(uint256)
func (_DnsAccountant *DnsAccountantCaller) DefaultCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "defaultCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultCount is a free data retrieval call binding the contract method 0xe8fe7bff.
//
// Solidity: function defaultCount() view returns(uint256)
func (_DnsAccountant *DnsAccountantSession) DefaultCount() (*big.Int, error) {
	return _DnsAccountant.Contract.DefaultCount(&_DnsAccountant.CallOpts)
}

// DefaultCount is a free data retrieval call binding the contract method 0xe8fe7bff.
//
// Solidity: function defaultCount() view returns(uint256)
func (_DnsAccountant *DnsAccountantCallerSession) DefaultCount() (*big.Int, error) {
	return _DnsAccountant.Contract.DefaultCount(&_DnsAccountant.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address account_, address erc20Addr_) view returns(uint256)
func (_DnsAccountant *DnsAccountantCaller) Get(opts *bind.CallOpts, account_ common.Address, erc20Addr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "get", account_, erc20Addr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address account_, address erc20Addr_) view returns(uint256)
func (_DnsAccountant *DnsAccountantSession) Get(account_ common.Address, erc20Addr_ common.Address) (*big.Int, error) {
	return _DnsAccountant.Contract.Get(&_DnsAccountant.CallOpts, account_, erc20Addr_)
}

// Get is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address account_, address erc20Addr_) view returns(uint256)
func (_DnsAccountant *DnsAccountantCallerSession) Get(account_ common.Address, erc20Addr_ common.Address) (*big.Int, error) {
	return _DnsAccountant.Contract.Get(&_DnsAccountant.CallOpts, account_, erc20Addr_)
}

// GetAllTask is a free data retrieval call binding the contract method 0x3db98b04.
//
// Solidity: function getAllTask(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCaller) GetAllTask(opts *bind.CallOpts, contractAddr_ common.Address) ([]byte, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getAllTask", contractAddr_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllTask is a free data retrieval call binding the contract method 0x3db98b04.
//
// Solidity: function getAllTask(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantSession) GetAllTask(contractAddr_ common.Address) ([]byte, error) {
	return _DnsAccountant.Contract.GetAllTask(&_DnsAccountant.CallOpts, contractAddr_)
}

// GetAllTask is a free data retrieval call binding the contract method 0x3db98b04.
//
// Solidity: function getAllTask(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCallerSession) GetAllTask(contractAddr_ common.Address) ([]byte, error) {
	return _DnsAccountant.Contract.GetAllTask(&_DnsAccountant.CallOpts, contractAddr_)
}

// GetContractList is a free data retrieval call binding the contract method 0x3410c6f9.
//
// Solidity: function getContractList(address signer_, uint256 idx_) view returns(uint256, uint256, address[8])
func (_DnsAccountant *DnsAccountantCaller) GetContractList(opts *bind.CallOpts, signer_ common.Address, idx_ *big.Int) (*big.Int, *big.Int, [8]common.Address, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getContractList", signer_, idx_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([8]common.Address)).(*[8]common.Address)

	return out0, out1, out2, err

}

// GetContractList is a free data retrieval call binding the contract method 0x3410c6f9.
//
// Solidity: function getContractList(address signer_, uint256 idx_) view returns(uint256, uint256, address[8])
func (_DnsAccountant *DnsAccountantSession) GetContractList(signer_ common.Address, idx_ *big.Int) (*big.Int, *big.Int, [8]common.Address, error) {
	return _DnsAccountant.Contract.GetContractList(&_DnsAccountant.CallOpts, signer_, idx_)
}

// GetContractList is a free data retrieval call binding the contract method 0x3410c6f9.
//
// Solidity: function getContractList(address signer_, uint256 idx_) view returns(uint256, uint256, address[8])
func (_DnsAccountant *DnsAccountantCallerSession) GetContractList(signer_ common.Address, idx_ *big.Int) (*big.Int, *big.Int, [8]common.Address, error) {
	return _DnsAccountant.Contract.GetContractList(&_DnsAccountant.CallOpts, signer_, idx_)
}

// GetSafeSig is a free data retrieval call binding the contract method 0x74fa74f8.
//
// Solidity: function getSafeSig(address contractAddr_, bytes32 hash_, address candidate_) view returns(bool)
func (_DnsAccountant *DnsAccountantCaller) GetSafeSig(opts *bind.CallOpts, contractAddr_ common.Address, hash_ [32]byte, candidate_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getSafeSig", contractAddr_, hash_, candidate_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSafeSig is a free data retrieval call binding the contract method 0x74fa74f8.
//
// Solidity: function getSafeSig(address contractAddr_, bytes32 hash_, address candidate_) view returns(bool)
func (_DnsAccountant *DnsAccountantSession) GetSafeSig(contractAddr_ common.Address, hash_ [32]byte, candidate_ common.Address) (bool, error) {
	return _DnsAccountant.Contract.GetSafeSig(&_DnsAccountant.CallOpts, contractAddr_, hash_, candidate_)
}

// GetSafeSig is a free data retrieval call binding the contract method 0x74fa74f8.
//
// Solidity: function getSafeSig(address contractAddr_, bytes32 hash_, address candidate_) view returns(bool)
func (_DnsAccountant *DnsAccountantCallerSession) GetSafeSig(contractAddr_ common.Address, hash_ [32]byte, candidate_ common.Address) (bool, error) {
	return _DnsAccountant.Contract.GetSafeSig(&_DnsAccountant.CallOpts, contractAddr_, hash_, candidate_)
}

// GetSignerSetAddress is a free data retrieval call binding the contract method 0x65945587.
//
// Solidity: function getSignerSetAddress(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCaller) GetSignerSetAddress(opts *bind.CallOpts, contractAddr_ common.Address) ([]byte, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getSignerSetAddress", contractAddr_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSignerSetAddress is a free data retrieval call binding the contract method 0x65945587.
//
// Solidity: function getSignerSetAddress(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantSession) GetSignerSetAddress(contractAddr_ common.Address) ([]byte, error) {
	return _DnsAccountant.Contract.GetSignerSetAddress(&_DnsAccountant.CallOpts, contractAddr_)
}

// GetSignerSetAddress is a free data retrieval call binding the contract method 0x65945587.
//
// Solidity: function getSignerSetAddress(address contractAddr_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCallerSession) GetSignerSetAddress(contractAddr_ common.Address) ([]byte, error) {
	return _DnsAccountant.Contract.GetSignerSetAddress(&_DnsAccountant.CallOpts, contractAddr_)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x328d4d8c.
//
// Solidity: function getTaskInfo(address contractAddr_, bytes32 hash_) view returns(uint256 max, uint256 cnt, bytes)
func (_DnsAccountant *DnsAccountantCaller) GetTaskInfo(opts *bind.CallOpts, contractAddr_ common.Address, hash_ [32]byte) (*big.Int, *big.Int, []byte, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getTaskInfo", contractAddr_, hash_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x328d4d8c.
//
// Solidity: function getTaskInfo(address contractAddr_, bytes32 hash_) view returns(uint256 max, uint256 cnt, bytes)
func (_DnsAccountant *DnsAccountantSession) GetTaskInfo(contractAddr_ common.Address, hash_ [32]byte) (*big.Int, *big.Int, []byte, error) {
	return _DnsAccountant.Contract.GetTaskInfo(&_DnsAccountant.CallOpts, contractAddr_, hash_)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x328d4d8c.
//
// Solidity: function getTaskInfo(address contractAddr_, bytes32 hash_) view returns(uint256 max, uint256 cnt, bytes)
func (_DnsAccountant *DnsAccountantCallerSession) GetTaskInfo(contractAddr_ common.Address, hash_ [32]byte) (*big.Int, *big.Int, []byte, error) {
	return _DnsAccountant.Contract.GetTaskInfo(&_DnsAccountant.CallOpts, contractAddr_, hash_)
}

// GetTaskSignerSetAddress is a free data retrieval call binding the contract method 0xea6e0c0c.
//
// Solidity: function getTaskSignerSetAddress(address contractAddr_, bytes32 hash_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCaller) GetTaskSignerSetAddress(opts *bind.CallOpts, contractAddr_ common.Address, hash_ [32]byte) ([]byte, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "getTaskSignerSetAddress", contractAddr_, hash_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTaskSignerSetAddress is a free data retrieval call binding the contract method 0xea6e0c0c.
//
// Solidity: function getTaskSignerSetAddress(address contractAddr_, bytes32 hash_) view returns(bytes)
func (_DnsAccountant *DnsAccountantSession) GetTaskSignerSetAddress(contractAddr_ common.Address, hash_ [32]byte) ([]byte, error) {
	return _DnsAccountant.Contract.GetTaskSignerSetAddress(&_DnsAccountant.CallOpts, contractAddr_, hash_)
}

// GetTaskSignerSetAddress is a free data retrieval call binding the contract method 0xea6e0c0c.
//
// Solidity: function getTaskSignerSetAddress(address contractAddr_, bytes32 hash_) view returns(bytes)
func (_DnsAccountant *DnsAccountantCallerSession) GetTaskSignerSetAddress(contractAddr_ common.Address, hash_ [32]byte) ([]byte, error) {
	return _DnsAccountant.Contract.GetTaskSignerSetAddress(&_DnsAccountant.CallOpts, contractAddr_, hash_)
}

// MultiSignerStore is a free data retrieval call binding the contract method 0x2696145b.
//
// Solidity: function multiSignerStore(address ) view returns(bool work, bool lock, uint256 SignerCount)
func (_DnsAccountant *DnsAccountantCaller) MultiSignerStore(opts *bind.CallOpts, arg0 common.Address) (struct {
	Work        bool
	Lock        bool
	SignerCount *big.Int
}, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "multiSignerStore", arg0)

	outstruct := new(struct {
		Work        bool
		Lock        bool
		SignerCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Work = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Lock = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.SignerCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MultiSignerStore is a free data retrieval call binding the contract method 0x2696145b.
//
// Solidity: function multiSignerStore(address ) view returns(bool work, bool lock, uint256 SignerCount)
func (_DnsAccountant *DnsAccountantSession) MultiSignerStore(arg0 common.Address) (struct {
	Work        bool
	Lock        bool
	SignerCount *big.Int
}, error) {
	return _DnsAccountant.Contract.MultiSignerStore(&_DnsAccountant.CallOpts, arg0)
}

// MultiSignerStore is a free data retrieval call binding the contract method 0x2696145b.
//
// Solidity: function multiSignerStore(address ) view returns(bool work, bool lock, uint256 SignerCount)
func (_DnsAccountant *DnsAccountantCallerSession) MultiSignerStore(arg0 common.Address) (struct {
	Work        bool
	Lock        bool
	SignerCount *big.Int
}, error) {
	return _DnsAccountant.Contract.MultiSignerStore(&_DnsAccountant.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsAccountant *DnsAccountantCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsAccountant.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsAccountant *DnsAccountantSession) Owner() (common.Address, error) {
	return _DnsAccountant.Contract.Owner(&_DnsAccountant.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsAccountant *DnsAccountantCallerSession) Owner() (common.Address, error) {
	return _DnsAccountant.Contract.Owner(&_DnsAccountant.CallOpts)
}

// AddSigner is a paid mutator transaction binding the contract method 0x2239f556.
//
// Solidity: function addSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantTransactor) AddSigner(opts *bind.TransactOpts, contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "addSigner", contractAddr_, signer_)
}

// AddSigner is a paid mutator transaction binding the contract method 0x2239f556.
//
// Solidity: function addSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantSession) AddSigner(contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.AddSigner(&_DnsAccountant.TransactOpts, contractAddr_, signer_)
}

// AddSigner is a paid mutator transaction binding the contract method 0x2239f556.
//
// Solidity: function addSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) AddSigner(contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.AddSigner(&_DnsAccountant.TransactOpts, contractAddr_, signer_)
}

// CloseMultiSig is a paid mutator transaction binding the contract method 0x56a77bfe.
//
// Solidity: function closeMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactor) CloseMultiSig(opts *bind.TransactOpts, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "closeMultiSig", contractAddr_)
}

// CloseMultiSig is a paid mutator transaction binding the contract method 0x56a77bfe.
//
// Solidity: function closeMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantSession) CloseMultiSig(contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.CloseMultiSig(&_DnsAccountant.TransactOpts, contractAddr_)
}

// CloseMultiSig is a paid mutator transaction binding the contract method 0x56a77bfe.
//
// Solidity: function closeMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) CloseMultiSig(contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.CloseMultiSig(&_DnsAccountant.TransactOpts, contractAddr_)
}

// DelSigner is a paid mutator transaction binding the contract method 0x4f93e17f.
//
// Solidity: function delSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantTransactor) DelSigner(opts *bind.TransactOpts, contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "delSigner", contractAddr_, signer_)
}

// DelSigner is a paid mutator transaction binding the contract method 0x4f93e17f.
//
// Solidity: function delSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantSession) DelSigner(contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.DelSigner(&_DnsAccountant.TransactOpts, contractAddr_, signer_)
}

// DelSigner is a paid mutator transaction binding the contract method 0x4f93e17f.
//
// Solidity: function delSigner(address contractAddr_, address signer_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) DelSigner(contractAddr_ common.Address, signer_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.DelSigner(&_DnsAccountant.TransactOpts, contractAddr_, signer_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantTransactor) Deposit(opts *bind.TransactOpts, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "deposit", erc20Addr_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantSession) Deposit(erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Deposit(&_DnsAccountant.TransactOpts, erc20Addr_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) Deposit(erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Deposit(&_DnsAccountant.TransactOpts, erc20Addr_, amount_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantTransactor) Deposit0(opts *bind.TransactOpts, contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "deposit0", contractAddr_, erc20Addr_, amount_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantSession) Deposit0(contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Deposit0(&_DnsAccountant.TransactOpts, contractAddr_, erc20Addr_, amount_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) Deposit0(contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Deposit0(&_DnsAccountant.TransactOpts, contractAddr_, erc20Addr_, amount_)
}

// OpenMultiSig is a paid mutator transaction binding the contract method 0xba272e03.
//
// Solidity: function openMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactor) OpenMultiSig(opts *bind.TransactOpts, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "openMultiSig", contractAddr_)
}

// OpenMultiSig is a paid mutator transaction binding the contract method 0xba272e03.
//
// Solidity: function openMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantSession) OpenMultiSig(contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.OpenMultiSig(&_DnsAccountant.TransactOpts, contractAddr_)
}

// OpenMultiSig is a paid mutator transaction binding the contract method 0xba272e03.
//
// Solidity: function openMultiSig(address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) OpenMultiSig(contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.OpenMultiSig(&_DnsAccountant.TransactOpts, contractAddr_)
}

// RemoveSafeSig is a paid mutator transaction binding the contract method 0x13f1ff1b.
//
// Solidity: function removeSafeSig(address contractAddr_, bytes32 hash_) returns()
func (_DnsAccountant *DnsAccountantTransactor) RemoveSafeSig(opts *bind.TransactOpts, contractAddr_ common.Address, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "removeSafeSig", contractAddr_, hash_)
}

// RemoveSafeSig is a paid mutator transaction binding the contract method 0x13f1ff1b.
//
// Solidity: function removeSafeSig(address contractAddr_, bytes32 hash_) returns()
func (_DnsAccountant *DnsAccountantSession) RemoveSafeSig(contractAddr_ common.Address, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsAccountant.Contract.RemoveSafeSig(&_DnsAccountant.TransactOpts, contractAddr_, hash_)
}

// RemoveSafeSig is a paid mutator transaction binding the contract method 0x13f1ff1b.
//
// Solidity: function removeSafeSig(address contractAddr_, bytes32 hash_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) RemoveSafeSig(contractAddr_ common.Address, hash_ [32]byte) (*types.Transaction, error) {
	return _DnsAccountant.Contract.RemoveSafeSig(&_DnsAccountant.TransactOpts, contractAddr_, hash_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsAccountant *DnsAccountantTransactor) TransferDaoOwner(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "transferDaoOwner", newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsAccountant *DnsAccountantSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.TransferDaoOwner(&_DnsAccountant.TransactOpts, newOwner_)
}

// TransferDaoOwner is a paid mutator transaction binding the contract method 0x028bff95.
//
// Solidity: function transferDaoOwner(address newOwner_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) TransferDaoOwner(newOwner_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.TransferDaoOwner(&_DnsAccountant.TransactOpts, newOwner_)
}

// WithdrawCash is a paid mutator transaction binding the contract method 0x47fd8e7d.
//
// Solidity: function withdrawCash(address erc20Addr_, uint256 amount_, address out_, address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactor) WithdrawCash(opts *bind.TransactOpts, erc20Addr_ common.Address, amount_ *big.Int, out_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.contract.Transact(opts, "withdrawCash", erc20Addr_, amount_, out_, contractAddr_)
}

// WithdrawCash is a paid mutator transaction binding the contract method 0x47fd8e7d.
//
// Solidity: function withdrawCash(address erc20Addr_, uint256 amount_, address out_, address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantSession) WithdrawCash(erc20Addr_ common.Address, amount_ *big.Int, out_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.WithdrawCash(&_DnsAccountant.TransactOpts, erc20Addr_, amount_, out_, contractAddr_)
}

// WithdrawCash is a paid mutator transaction binding the contract method 0x47fd8e7d.
//
// Solidity: function withdrawCash(address erc20Addr_, uint256 amount_, address out_, address contractAddr_) returns()
func (_DnsAccountant *DnsAccountantTransactorSession) WithdrawCash(erc20Addr_ common.Address, amount_ *big.Int, out_ common.Address, contractAddr_ common.Address) (*types.Transaction, error) {
	return _DnsAccountant.Contract.WithdrawCash(&_DnsAccountant.TransactOpts, erc20Addr_, amount_, out_, contractAddr_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DnsAccountant *DnsAccountantTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DnsAccountant.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DnsAccountant *DnsAccountantSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Fallback(&_DnsAccountant.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DnsAccountant *DnsAccountantTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DnsAccountant.Contract.Fallback(&_DnsAccountant.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DnsAccountant *DnsAccountantTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsAccountant.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DnsAccountant *DnsAccountantSession) Receive() (*types.Transaction, error) {
	return _DnsAccountant.Contract.Receive(&_DnsAccountant.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DnsAccountant *DnsAccountantTransactorSession) Receive() (*types.Transaction, error) {
	return _DnsAccountant.Contract.Receive(&_DnsAccountant.TransactOpts)
}

// DnsAccountantEvDepositIterator is returned from FilterEvDeposit and is used to iterate over the raw logs and unpacked data for EvDeposit events raised by the DnsAccountant contract.
type DnsAccountantEvDepositIterator struct {
	Event *DnsAccountantEvDeposit // Event containing the contract specifics and raw log

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
func (it *DnsAccountantEvDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsAccountantEvDeposit)
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
		it.Event = new(DnsAccountantEvDeposit)
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
func (it *DnsAccountantEvDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsAccountantEvDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsAccountantEvDeposit represents a EvDeposit event raised by the DnsAccountant contract.
type DnsAccountantEvDeposit struct {
	Operator     common.Address
	ContractAddr common.Address
	Erc20Addr    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvDeposit is a free log retrieval operation binding the contract event 0x5e8655bc31e2c8e8d0e80adca4302fafb5ec9320198d3221616d9214c1304601.
//
// Solidity: event EvDeposit(address operator, address contractAddr, address erc20Addr, uint256 amount)
func (_DnsAccountant *DnsAccountantFilterer) FilterEvDeposit(opts *bind.FilterOpts) (*DnsAccountantEvDepositIterator, error) {

	logs, sub, err := _DnsAccountant.contract.FilterLogs(opts, "EvDeposit")
	if err != nil {
		return nil, err
	}
	return &DnsAccountantEvDepositIterator{contract: _DnsAccountant.contract, event: "EvDeposit", logs: logs, sub: sub}, nil
}

// WatchEvDeposit is a free log subscription operation binding the contract event 0x5e8655bc31e2c8e8d0e80adca4302fafb5ec9320198d3221616d9214c1304601.
//
// Solidity: event EvDeposit(address operator, address contractAddr, address erc20Addr, uint256 amount)
func (_DnsAccountant *DnsAccountantFilterer) WatchEvDeposit(opts *bind.WatchOpts, sink chan<- *DnsAccountantEvDeposit) (event.Subscription, error) {

	logs, sub, err := _DnsAccountant.contract.WatchLogs(opts, "EvDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsAccountantEvDeposit)
				if err := _DnsAccountant.contract.UnpackLog(event, "EvDeposit", log); err != nil {
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

// ParseEvDeposit is a log parse operation binding the contract event 0x5e8655bc31e2c8e8d0e80adca4302fafb5ec9320198d3221616d9214c1304601.
//
// Solidity: event EvDeposit(address operator, address contractAddr, address erc20Addr, uint256 amount)
func (_DnsAccountant *DnsAccountantFilterer) ParseEvDeposit(log types.Log) (*DnsAccountantEvDeposit, error) {
	event := new(DnsAccountantEvDeposit)
	if err := _DnsAccountant.contract.UnpackLog(event, "EvDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsAccountantEvWithdrawCashIterator is returned from FilterEvWithdrawCash and is used to iterate over the raw logs and unpacked data for EvWithdrawCash events raised by the DnsAccountant contract.
type DnsAccountantEvWithdrawCashIterator struct {
	Event *DnsAccountantEvWithdrawCash // Event containing the contract specifics and raw log

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
func (it *DnsAccountantEvWithdrawCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsAccountantEvWithdrawCash)
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
		it.Event = new(DnsAccountantEvWithdrawCash)
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
func (it *DnsAccountantEvWithdrawCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsAccountantEvWithdrawCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsAccountantEvWithdrawCash represents a EvWithdrawCash event raised by the DnsAccountant contract.
type DnsAccountantEvWithdrawCash struct {
	Operator     common.Address
	Erc20Addr    common.Address
	Amount       *big.Int
	Out          common.Address
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEvWithdrawCash is a free log retrieval operation binding the contract event 0x0a9c814dc547650550ef7ab0ca47351429be4f6d9556f7039a7652799f098d78.
//
// Solidity: event EvWithdrawCash(address operator, address erc20Addr, uint256 amount, address out, address contractAddr)
func (_DnsAccountant *DnsAccountantFilterer) FilterEvWithdrawCash(opts *bind.FilterOpts) (*DnsAccountantEvWithdrawCashIterator, error) {

	logs, sub, err := _DnsAccountant.contract.FilterLogs(opts, "EvWithdrawCash")
	if err != nil {
		return nil, err
	}
	return &DnsAccountantEvWithdrawCashIterator{contract: _DnsAccountant.contract, event: "EvWithdrawCash", logs: logs, sub: sub}, nil
}

// WatchEvWithdrawCash is a free log subscription operation binding the contract event 0x0a9c814dc547650550ef7ab0ca47351429be4f6d9556f7039a7652799f098d78.
//
// Solidity: event EvWithdrawCash(address operator, address erc20Addr, uint256 amount, address out, address contractAddr)
func (_DnsAccountant *DnsAccountantFilterer) WatchEvWithdrawCash(opts *bind.WatchOpts, sink chan<- *DnsAccountantEvWithdrawCash) (event.Subscription, error) {

	logs, sub, err := _DnsAccountant.contract.WatchLogs(opts, "EvWithdrawCash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsAccountantEvWithdrawCash)
				if err := _DnsAccountant.contract.UnpackLog(event, "EvWithdrawCash", log); err != nil {
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

// ParseEvWithdrawCash is a log parse operation binding the contract event 0x0a9c814dc547650550ef7ab0ca47351429be4f6d9556f7039a7652799f098d78.
//
// Solidity: event EvWithdrawCash(address operator, address erc20Addr, uint256 amount, address out, address contractAddr)
func (_DnsAccountant *DnsAccountantFilterer) ParseEvWithdrawCash(log types.Log) (*DnsAccountantEvWithdrawCash, error) {
	event := new(DnsAccountantEvWithdrawCash)
	if err := _DnsAccountant.contract.UnpackLog(event, "EvWithdrawCash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
