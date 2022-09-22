// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package top

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200921c3e9992238c355345b8fee4626612e3ebdf646b2f393c024e2e15285ca6364736f6c63430008060033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DnsNameErc721MetaData contains all meta data concerning the DnsNameErc721 contract.
var DnsNameErc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"}],\"name\":\"DnsExtendExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DnsTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"idOwner_\",\"type\":\"address\"}],\"name\":\"MintId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SigUserAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc721Owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gNameId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nameIdInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseUri_\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sigUser_\",\"type\":\"address\"}],\"name\":\"setSigUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003e0438038062003e04833981810160405281019062000037919062000268565b813385858160009080519060200190620000539291906200010c565b5080600190805190602001906200006c9291906200010c565b50505080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600060088190555080600981905550505050506200050e565b8280546200011a90620003eb565b90600052602060002090601f0160209004810192826200013e57600085556200018a565b82601f106200015957805160ff19168380011785556200018a565b828001600101855582156200018a579182015b82811115620001895782518255916020019190600101906200016c565b5b5090506200019991906200019d565b5090565b5b80821115620001b85760008160009055506001016200019e565b5090565b6000620001d3620001cd8462000341565b62000318565b905082815260208101848484011115620001f257620001f1620004ba565b5b620001ff848285620003b5565b509392505050565b6000815190506200021881620004da565b92915050565b6000815190506200022f81620004f4565b92915050565b600082601f8301126200024d576200024c620004b5565b5b81516200025f848260208601620001bc565b91505092915050565b60008060008060808587031215620002855762000284620004c4565b5b600085015167ffffffffffffffff811115620002a657620002a5620004bf565b5b620002b48782880162000235565b945050602085015167ffffffffffffffff811115620002d857620002d7620004bf565b5b620002e68782880162000235565b9350506040620002f98782880162000207565b92505060606200030c878288016200021e565b91505092959194509250565b60006200032462000337565b905062000332828262000421565b919050565b6000604051905090565b600067ffffffffffffffff8211156200035f576200035e62000486565b5b6200036a82620004c9565b9050602081019050919050565b6000620003848262000395565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b83811015620003d5578082015181840152602081019050620003b8565b83811115620003e5576000848401525b50505050565b600060028204905060018216806200040457607f821691505b602082108114156200041b576200041a62000457565b5b50919050565b6200042c82620004c9565b810181811067ffffffffffffffff821117156200044e576200044d62000486565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b620004e58162000377565b8114620004f157600080fd5b50565b620004ff816200038b565b81146200050b57600080fd5b50565b6138e6806200051e6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c806395d89b41116100de578063c87b56dd11610097578063d575f4ef11610071578063d575f4ef14610494578063e767f8fd146104b2578063e985e9c5146104ce578063f2fde38b146104fe5761018e565b8063c87b56dd1461042c578063c8eba7601461045c578063cb50fe76146104785761018e565b806395d89b411461036c5780639abc83201461038a578063a0bcfc7f146103a8578063a22cb465146103c4578063adfd5f91146103e0578063b88d4fde146104105761018e565b806332cdee7b1161014b5780635c707f07116101255780635c707f07146102d25780636352211e146102ee57806370a082311461031e5780638da5cb5b1461034e5761018e565b806332cdee7b1461027a5780633cf8baa21461029857806342842e0e146102b65761018e565b806301ffc9a7146101935780630636a797146101c357806306fdde03146101f4578063081812fc14610212578063095ea7b31461024257806323b872dd1461025e575b600080fd5b6101ad60048036038101906101a8919061294c565b61051a565b6040516101ba9190612e3e565b60405180910390f35b6101dd60048036038101906101d89190612a67565b61052c565b6040516101eb929190613073565b60405180910390f35b6101fc610560565b6040516102099190612e96565b60405180910390f35b61022c60048036038101906102279190612a67565b6105f2565b6040516102399190612dbc565b60405180910390f35b61025c600480360381019061025791906128b9565b610604565b005b610278600480360381019061027391906127a3565b610612565b005b61028261062c565b60405161028f9190612dd7565b60405180910390f35b6102a0610652565b6040516102ad9190612dbc565b60405180910390f35b6102d060048036038101906102cb91906127a3565b61067c565b005b6102ec60048036038101906102e791906129ef565b61069c565b005b61030860048036038101906103039190612a67565b6106ce565b6040516103159190612dbc565b60405180910390f35b610338600480360381019061033391906126dc565b610780565b6040516103459190613058565b60405180910390f35b610356610838565b6040516103639190612dd7565b60405180910390f35b61037461085e565b6040516103819190612e96565b60405180910390f35b6103926108f0565b60405161039f9190612e74565b60405180910390f35b6103c260048036038101906103bd91906129a6565b61097e565b005b6103de60048036038101906103d99190612879565b6109f2565b005b6103fa60048036038101906103f591906128f9565b610a00565b6040516104079190613058565b60405180910390f35b61042a600480360381019061042591906127f6565b610c1c565b005b61044660048036038101906104419190612a67565b610c38565b6040516104539190612e96565b60405180910390f35b610476600480360381019061047191906128b9565b610ca0565b005b610492600480360381019061048d91906126dc565b610e47565b005b61049c610ee5565b6040516104a99190613058565b60405180910390f35b6104cc60048036038101906104c79190612a94565b610eeb565b005b6104e860048036038101906104e39190612763565b6110aa565b6040516104f59190612e3e565b60405180910390f35b61051860048036038101906105139190612736565b6110be565b005b600061052582611192565b9050919050565b600c6020528060005260406000206000915090508060000160009054906101000a900463ffffffff16908060010154905082565b60606000805461056f90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461059b90613303565b80156105e85780601f106105bd576101008083540402835291602001916105e8565b820191906000526020600020905b8154815290600101906020018083116105cb57829003601f168201915b5050505050905090565b60006105fd82611274565b9050919050565b61060e82826112ba565b5050565b61061c82826113d2565b610627838383611672565b505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61069783838360405180602001604052806000815250610c1c565b505050565b81600090805190602001906106b2929190612416565b5080600190805190602001906106c9929190612416565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076e90612fd8565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e890612f78565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606001805461086d90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461089990613303565b80156108e65780601f106108bb576101008083540402835291602001916108e6565b820191906000526020600020905b8154815290600101906020018083116108c957829003601f168201915b5050505050905090565b600a80546108fd90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461092990613303565b80156109765780601f1061094b57610100808354040283529160200191610976565b820191906000526020600020905b81548152906001019060200180831161095957829003601f168201915b505050505081565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d857600080fd5b80600a90805190602001906109ee92919061249c565b5050565b6109fc82826116d2565b5050565b60008060001b600954148015610a6357503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80610b4957506000801b60095414158015610b4857503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610af857600080fd5b505afa158015610b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b309190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b610b88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7f90612f58565b60405180910390fd5b60086000815480929190610b9b90613366565b9190505550610bac826008546116e8565b60405180604001604052808463ffffffff16815260200185815250600c6000600854815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff1602179055506020820151816001015590505060085490509392505050565b610c2683836113d2565b610c32848484846118c2565b50505050565b6060610c4382611924565b6000610c4d61196f565b90506000815111610c6d5760405180602001604052806000815250610c98565b80610c7784611a01565b604051602001610c88929190612d98565b6040516020818303038152906040525b915050919050565b6000801b600954148015610d0157503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80610de757506000801b60095414158015610de657503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610d9657600080fd5b505afa158015610daa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dce9190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b610e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1d90612f58565b60405180910390fd5b610e39610e32826106ce565b8383611b62565b610e4382826113d2565b5050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ea157600080fd5b80600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60085481565b6000801b600954148015610f4c57503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b8061103257506000801b6009541415801561103157503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610fe157600080fd5b505afa158015610ff5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110199190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b611071576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106890612f58565b60405180910390fd5b80600c600084815260200190815260200160002060000160006101000a81548163ffffffff021916908363ffffffff1602179055505050565b60006110b68383611dc9565b905092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461114e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114590612f58565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061125d57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061126d575061126c82611e5d565b5b9050919050565b600061127f82611924565b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006112c5826106ce565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132d90612ff8565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16611355611ec7565b73ffffffffffffffffffffffffffffffffffffffff16148061138457506113838161137e611ec7565b6110aa565b5b6113c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ba90612f98565b60405180910390fd5b6113cd8383611ecf565b505050565b42600c600083815260200190815260200160002060000160009054906101000a900463ffffffff1663ffffffff1611611440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143790613018565b60405180910390fd5b6000801b600954141561166e57600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76600c6000858152602001908152602001600020600101546040518263ffffffff1660e01b81526004016114d69190612e59565b60206040518083038186803b1580156114ee57600080fd5b505afa158015611502573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115269190612709565b73ffffffffffffffffffffffffffffffffffffffff161461166d57600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76600c6000848152602001908152602001600020600101546040518263ffffffff1660e01b81526004016115b29190612e59565b60206040518083038186803b1580156115ca57600080fd5b505afa1580156115de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116029190612709565b73ffffffffffffffffffffffffffffffffffffffff1663f2fde38b836040518263ffffffff1660e01b815260040161163a9190612dd7565b600060405180830381600087803b15801561165457600080fd5b505af1158015611668573d6000803e3d6000fd5b505050505b5b5050565b61168361167d611ec7565b82611f88565b6116c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b990613038565b60405180910390fd5b6116cd838383611b62565b505050565b6116e46116dd611ec7565b838361201d565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f90612fb8565b60405180910390fd5b6117618161218a565b156117a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179890612ef8565b60405180910390fd5b6117ad600083836121f6565b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117fd9190613166565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a46118be600083836121fb565b5050565b6118d36118cd611ec7565b83611f88565b611912576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190990613038565b60405180910390fd5b61191e84848484612200565b50505050565b61192d8161218a565b61196c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196390612fd8565b60405180910390fd5b50565b6060600a805461197e90613303565b80601f01602080910402602001604051908101604052809291908181526020018280546119aa90613303565b80156119f75780601f106119cc576101008083540402835291602001916119f7565b820191906000526020600020905b8154815290600101906020018083116119da57829003601f168201915b5050505050905090565b60606000821415611a49576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611b5d565b600082905060005b60008214611a7b578080611a6490613366565b915050600a82611a7491906131bc565b9150611a51565b60008167ffffffffffffffff811115611a9757611a9661349c565b5b6040519080825280601f01601f191660200182016040528015611ac95781602001600182028036833780820191505090505b5090505b60008514611b5657600182611ae291906131ed565b9150600a85611af191906133af565b6030611afd9190613166565b60f81b818381518110611b1357611b1261346d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a85611b4f91906131bc565b9450611acd565b8093505050505b919050565b8273ffffffffffffffffffffffffffffffffffffffff16611b82826106ce565b73ffffffffffffffffffffffffffffffffffffffff1614611bd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bcf90612ed8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3f90612f18565b60405180910390fd5b611c538383836121f6565b611c5e600082611ecf565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611cae91906131ed565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d059190613166565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4611dc48383836121fb565b505050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16611f42836106ce565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080611f94836106ce565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611fd65750611fd581856110aa565b5b8061201457508373ffffffffffffffffffffffffffffffffffffffff16611ffc846105f2565b73ffffffffffffffffffffffffffffffffffffffff16145b91505092915050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561208c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161208390612f38565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161217d9190612e3e565b60405180910390a3505050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b505050565b505050565b61220b848484611b62565b6122178484848461225c565b612256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161224d90612eb8565b60405180910390fd5b50505050565b600061227d8473ffffffffffffffffffffffffffffffffffffffff166123f3565b156123e6578373ffffffffffffffffffffffffffffffffffffffff1663150b7a026122a6611ec7565b8786866040518563ffffffff1660e01b81526004016122c89493929190612df2565b602060405180830381600087803b1580156122e257600080fd5b505af192505050801561231357506040513d601f19601f820116820180604052508101906123109190612979565b60015b612396573d8060008114612343576040519150601f19603f3d011682016040523d82523d6000602084013e612348565b606091505b5060008151141561238e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161238590612eb8565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149150506123eb565b600190505b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b82805461242290613303565b90600052602060002090601f016020900481019282612444576000855561248b565b82601f1061245d57805160ff191683800117855561248b565b8280016001018555821561248b579182015b8281111561248a57825182559160200191906001019061246f565b5b5090506124989190612522565b5090565b8280546124a890613303565b90600052602060002090601f0160209004810192826124ca5760008555612511565b82601f106124e357805160ff1916838001178555612511565b82800160010185558215612511579182015b828111156125105782518255916020019190600101906124f5565b5b50905061251e9190612522565b5090565b5b8082111561253b576000816000905550600101612523565b5090565b600061255261254d846130c1565b61309c565b90508281526020810184848401111561256e5761256d6134d0565b5b6125798482856132c1565b509392505050565b600061259461258f846130f2565b61309c565b9050828152602081018484840111156125b0576125af6134d0565b5b6125bb8482856132c1565b509392505050565b6000813590506125d28161380f565b92915050565b6000815190506125e78161380f565b92915050565b6000813590506125fc81613826565b92915050565b6000813590506126118161383d565b92915050565b60008135905061262681613854565b92915050565b60008135905061263b8161386b565b92915050565b6000815190506126508161386b565b92915050565b600082601f83011261266b5761266a6134cb565b5b813561267b84826020860161253f565b91505092915050565b600082601f830112612699576126986134cb565b5b81356126a9848260208601612581565b91505092915050565b6000813590506126c181613882565b92915050565b6000813590506126d681613899565b92915050565b6000602082840312156126f2576126f16134da565b5b6000612700848285016125c3565b91505092915050565b60006020828403121561271f5761271e6134da565b5b600061272d848285016125d8565b91505092915050565b60006020828403121561274c5761274b6134da565b5b600061275a848285016125ed565b91505092915050565b6000806040838503121561277a576127796134da565b5b6000612788858286016125c3565b9250506020612799858286016125c3565b9150509250929050565b6000806000606084860312156127bc576127bb6134da565b5b60006127ca868287016125c3565b93505060206127db868287016125c3565b92505060406127ec868287016126b2565b9150509250925092565b600080600080608085870312156128105761280f6134da565b5b600061281e878288016125c3565b945050602061282f878288016125c3565b9350506040612840878288016126b2565b925050606085013567ffffffffffffffff811115612861576128606134d5565b5b61286d87828801612656565b91505092959194509250565b600080604083850312156128905761288f6134da565b5b600061289e858286016125c3565b92505060206128af85828601612602565b9150509250929050565b600080604083850312156128d0576128cf6134da565b5b60006128de858286016125c3565b92505060206128ef858286016126b2565b9150509250929050565b600080600060608486031215612912576129116134da565b5b600061292086828701612617565b9350506020612931868287016126c7565b9250506040612942868287016125c3565b9150509250925092565b600060208284031215612962576129616134da565b5b60006129708482850161262c565b91505092915050565b60006020828403121561298f5761298e6134da565b5b600061299d84828501612641565b91505092915050565b6000602082840312156129bc576129bb6134da565b5b600082013567ffffffffffffffff8111156129da576129d96134d5565b5b6129e684828501612684565b91505092915050565b60008060408385031215612a0657612a056134da565b5b600083013567ffffffffffffffff811115612a2457612a236134d5565b5b612a3085828601612684565b925050602083013567ffffffffffffffff811115612a5157612a506134d5565b5b612a5d85828601612684565b9150509250929050565b600060208284031215612a7d57612a7c6134da565b5b6000612a8b848285016126b2565b91505092915050565b60008060408385031215612aab57612aaa6134da565b5b6000612ab9858286016126b2565b9250506020612aca858286016126c7565b9150509250929050565b612add81613233565b82525050565b612aec81613221565b82525050565b612afb81613245565b82525050565b612b0a81613251565b82525050565b6000612b1b82613123565b612b258185613139565b9350612b358185602086016132d0565b612b3e816134df565b840191505092915050565b6000612b548261312e565b612b5e818561314a565b9350612b6e8185602086016132d0565b612b77816134df565b840191505092915050565b6000612b8d8261312e565b612b97818561315b565b9350612ba78185602086016132d0565b80840191505092915050565b6000612bc060328361314a565b9150612bcb826134f0565b604082019050919050565b6000612be360258361314a565b9150612bee8261353f565b604082019050919050565b6000612c06601c8361314a565b9150612c118261358e565b602082019050919050565b6000612c2960248361314a565b9150612c34826135b7565b604082019050919050565b6000612c4c60198361314a565b9150612c5782613606565b602082019050919050565b6000612c6f600b8361314a565b9150612c7a8261362f565b602082019050919050565b6000612c9260298361314a565b9150612c9d82613658565b604082019050919050565b6000612cb5603e8361314a565b9150612cc0826136a7565b604082019050919050565b6000612cd860208361314a565b9150612ce3826136f6565b602082019050919050565b6000612cfb60188361314a565b9150612d068261371f565b602082019050919050565b6000612d1e60218361314a565b9150612d2982613748565b604082019050919050565b6000612d41600c8361314a565b9150612d4c82613797565b602082019050919050565b6000612d64602e8361314a565b9150612d6f826137c0565b604082019050919050565b612d83816132a7565b82525050565b612d92816132b1565b82525050565b6000612da48285612b82565b9150612db08284612b82565b91508190509392505050565b6000602082019050612dd16000830184612ae3565b92915050565b6000602082019050612dec6000830184612ad4565b92915050565b6000608082019050612e076000830187612ae3565b612e146020830186612ae3565b612e216040830185612d7a565b8181036060830152612e338184612b10565b905095945050505050565b6000602082019050612e536000830184612af2565b92915050565b6000602082019050612e6e6000830184612b01565b92915050565b60006020820190508181036000830152612e8e8184612b10565b905092915050565b60006020820190508181036000830152612eb08184612b49565b905092915050565b60006020820190508181036000830152612ed181612bb3565b9050919050565b60006020820190508181036000830152612ef181612bd6565b9050919050565b60006020820190508181036000830152612f1181612bf9565b9050919050565b60006020820190508181036000830152612f3181612c1c565b9050919050565b60006020820190508181036000830152612f5181612c3f565b9050919050565b60006020820190508181036000830152612f7181612c62565b9050919050565b60006020820190508181036000830152612f9181612c85565b9050919050565b60006020820190508181036000830152612fb181612ca8565b9050919050565b60006020820190508181036000830152612fd181612ccb565b9050919050565b60006020820190508181036000830152612ff181612cee565b9050919050565b6000602082019050818103600083015261301181612d11565b9050919050565b6000602082019050818103600083015261303181612d34565b9050919050565b6000602082019050818103600083015261305181612d57565b9050919050565b600060208201905061306d6000830184612d7a565b92915050565b60006040820190506130886000830185612d89565b6130956020830184612b01565b9392505050565b60006130a66130b7565b90506130b28282613335565b919050565b6000604051905090565b600067ffffffffffffffff8211156130dc576130db61349c565b5b6130e5826134df565b9050602081019050919050565b600067ffffffffffffffff82111561310d5761310c61349c565b5b613116826134df565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000613171826132a7565b915061317c836132a7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156131b1576131b06133e0565b5b828201905092915050565b60006131c7826132a7565b91506131d2836132a7565b9250826131e2576131e161340f565b5b828204905092915050565b60006131f8826132a7565b9150613203836132a7565b925082821015613216576132156133e0565b5b828203905092915050565b600061322c82613287565b9050919050565b600061323e82613287565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b82818337600083830152505050565b60005b838110156132ee5780820151818401526020810190506132d3565b838111156132fd576000848401525b50505050565b6000600282049050600182168061331b57607f821691505b6020821081141561332f5761332e61343e565b5b50919050565b61333e826134df565b810181811067ffffffffffffffff8211171561335d5761335c61349c565b5b80604052505050565b6000613371826132a7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156133a4576133a36133e0565b5b600182019050919050565b60006133ba826132a7565b91506133c5836132a7565b9250826133d5576133d461340f565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b7f4552433732313a207472616e736665722066726f6d20696e636f72726563742060008201527f6f776e6572000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b7f4552433732313a2061646472657373207a65726f206973206e6f74206120766160008201527f6c6964206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60008201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000602082015250565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b7f4552433732313a20696e76616c696420746f6b656e2049440000000000000000600082015250565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b7f6e616d6520657870697265640000000000000000000000000000000000000000600082015250565b7f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560008201527f72206e6f7220617070726f766564000000000000000000000000000000000000602082015250565b61381881613221565b811461382357600080fd5b50565b61382f81613233565b811461383a57600080fd5b50565b61384681613245565b811461385157600080fd5b50565b61385d81613251565b811461386857600080fd5b50565b6138748161325b565b811461387f57600080fd5b50565b61388b816132a7565b811461389657600080fd5b50565b6138a2816132b1565b81146138ad57600080fd5b5056fea264697066735822122076dc0689d8c8fdbe1daa082df3c00f7140c193dd5874b241af35ef3fa1301eac64736f6c63430008060033",
}

// DnsNameErc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsNameErc721MetaData.ABI instead.
var DnsNameErc721ABI = DnsNameErc721MetaData.ABI

// DnsNameErc721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DnsNameErc721MetaData.Bin instead.
var DnsNameErc721Bin = DnsNameErc721MetaData.Bin

// DeployDnsNameErc721 deploys a new Ethereum contract, binding an instance of DnsNameErc721 to it.
func DeployDnsNameErc721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, owner_ common.Address, hash_ [32]byte) (common.Address, *types.Transaction, *DnsNameErc721, error) {
	parsed, err := DnsNameErc721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DnsNameErc721Bin), backend, name_, symbol_, owner_, hash_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DnsNameErc721{DnsNameErc721Caller: DnsNameErc721Caller{contract: contract}, DnsNameErc721Transactor: DnsNameErc721Transactor{contract: contract}, DnsNameErc721Filterer: DnsNameErc721Filterer{contract: contract}}, nil
}

// DnsNameErc721 is an auto generated Go binding around an Ethereum contract.
type DnsNameErc721 struct {
	DnsNameErc721Caller     // Read-only binding to the contract
	DnsNameErc721Transactor // Write-only binding to the contract
	DnsNameErc721Filterer   // Log filterer for contract events
}

// DnsNameErc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type DnsNameErc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsNameErc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsNameErc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsNameErc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsNameErc721Session struct {
	Contract     *DnsNameErc721    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsNameErc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsNameErc721CallerSession struct {
	Contract *DnsNameErc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DnsNameErc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsNameErc721TransactorSession struct {
	Contract     *DnsNameErc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DnsNameErc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type DnsNameErc721Raw struct {
	Contract *DnsNameErc721 // Generic contract binding to access the raw methods on
}

// DnsNameErc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsNameErc721CallerRaw struct {
	Contract *DnsNameErc721Caller // Generic read-only contract binding to access the raw methods on
}

// DnsNameErc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsNameErc721TransactorRaw struct {
	Contract *DnsNameErc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsNameErc721 creates a new instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721(address common.Address, backend bind.ContractBackend) (*DnsNameErc721, error) {
	contract, err := bindDnsNameErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721{DnsNameErc721Caller: DnsNameErc721Caller{contract: contract}, DnsNameErc721Transactor: DnsNameErc721Transactor{contract: contract}, DnsNameErc721Filterer: DnsNameErc721Filterer{contract: contract}}, nil
}

// NewDnsNameErc721Caller creates a new read-only instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Caller(address common.Address, caller bind.ContractCaller) (*DnsNameErc721Caller, error) {
	contract, err := bindDnsNameErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Caller{contract: contract}, nil
}

// NewDnsNameErc721Transactor creates a new write-only instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*DnsNameErc721Transactor, error) {
	contract, err := bindDnsNameErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Transactor{contract: contract}, nil
}

// NewDnsNameErc721Filterer creates a new log filterer instance of DnsNameErc721, bound to a specific deployed contract.
func NewDnsNameErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*DnsNameErc721Filterer, error) {
	contract, err := bindDnsNameErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721Filterer{contract: contract}, nil
}

// bindDnsNameErc721 binds a generic wrapper to an already deployed contract.
func bindDnsNameErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsNameErc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsNameErc721 *DnsNameErc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsNameErc721.Contract.DnsNameErc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsNameErc721 *DnsNameErc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsNameErc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsNameErc721 *DnsNameErc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsNameErc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsNameErc721 *DnsNameErc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsNameErc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsNameErc721 *DnsNameErc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsNameErc721 *DnsNameErc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.contract.Transact(opts, method, params...)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) SigUserAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "SigUserAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) SigUserAddr() (common.Address, error) {
	return _DnsNameErc721.Contract.SigUserAddr(&_DnsNameErc721.CallOpts)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) SigUserAddr() (common.Address, error) {
	return _DnsNameErc721.Contract.SigUserAddr(&_DnsNameErc721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsNameErc721.Contract.BalanceOf(&_DnsNameErc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DnsNameErc721.Contract.BalanceOf(&_DnsNameErc721.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721Caller) BaseUri(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721Session) BaseUri() ([]byte, error) {
	return _DnsNameErc721.Contract.BaseUri(&_DnsNameErc721.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(bytes)
func (_DnsNameErc721 *DnsNameErc721CallerSession) BaseUri() ([]byte, error) {
	return _DnsNameErc721.Contract.BaseUri(&_DnsNameErc721.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) Erc721Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "erc721Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) Erc721Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Erc721Owner(&_DnsNameErc721.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Erc721Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Erc721Owner(&_DnsNameErc721.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Caller) GNameId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "gNameId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) GNameId() (*big.Int, error) {
	return _DnsNameErc721.Contract.GNameId(&_DnsNameErc721.CallOpts)
}

// GNameId is a free data retrieval call binding the contract method 0xd575f4ef.
//
// Solidity: function gNameId() view returns(uint256)
func (_DnsNameErc721 *DnsNameErc721CallerSession) GNameId() (*big.Int, error) {
	return _DnsNameErc721.Contract.GNameId(&_DnsNameErc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.GetApproved(&_DnsNameErc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.GetApproved(&_DnsNameErc721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator_ common.Address) (bool, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Session) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsNameErc721.Contract.IsApprovedForAll(&_DnsNameErc721.CallOpts, owner, operator_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator_) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721CallerSession) IsApprovedForAll(owner common.Address, operator_ common.Address) (bool, error) {
	return _DnsNameErc721.Contract.IsApprovedForAll(&_DnsNameErc721.CallOpts, owner, operator_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) Name() (string, error) {
	return _DnsNameErc721.Contract.Name(&_DnsNameErc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Name() (string, error) {
	return _DnsNameErc721.Contract.Name(&_DnsNameErc721.CallOpts)
}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721Caller) NameIdInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "nameIdInfo", arg0)

	outstruct := new(struct {
		ExpireTime uint32
		NameHash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExpireTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.NameHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721Session) NameIdInfo(arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	return _DnsNameErc721.Contract.NameIdInfo(&_DnsNameErc721.CallOpts, arg0)
}

// NameIdInfo is a free data retrieval call binding the contract method 0x0636a797.
//
// Solidity: function nameIdInfo(uint256 ) view returns(uint32 expireTime, bytes32 nameHash)
func (_DnsNameErc721 *DnsNameErc721CallerSession) NameIdInfo(arg0 *big.Int) (struct {
	ExpireTime uint32
	NameHash   [32]byte
}, error) {
	return _DnsNameErc721.Contract.NameIdInfo(&_DnsNameErc721.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Owner(&_DnsNameErc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Owner() (common.Address, error) {
	return _DnsNameErc721.Contract.Owner(&_DnsNameErc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.OwnerOf(&_DnsNameErc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DnsNameErc721 *DnsNameErc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DnsNameErc721.Contract.OwnerOf(&_DnsNameErc721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsNameErc721.Contract.SupportsInterface(&_DnsNameErc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DnsNameErc721 *DnsNameErc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DnsNameErc721.Contract.SupportsInterface(&_DnsNameErc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) Symbol() (string, error) {
	return _DnsNameErc721.Contract.Symbol(&_DnsNameErc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) Symbol() (string, error) {
	return _DnsNameErc721.Contract.Symbol(&_DnsNameErc721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DnsNameErc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsNameErc721.Contract.TokenURI(&_DnsNameErc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DnsNameErc721 *DnsNameErc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DnsNameErc721.Contract.TokenURI(&_DnsNameErc721.CallOpts, tokenId)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) DnsExtendExpire(opts *bind.TransactOpts, tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "DnsExtendExpire", tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsExtendExpire(&_DnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsExtendExpire(&_DnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) DnsTransfer(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "DnsTransfer", to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsTransfer(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.DnsTransfer(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Transactor) MintId(opts *bind.TransactOpts, hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "MintId", hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721Session) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.MintId(&_DnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_DnsNameErc721 *DnsNameErc721TransactorSession) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.MintId(&_DnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.Approve(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.Approve(&_DnsNameErc721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom0(&_DnsNameErc721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SafeTransferFrom0(&_DnsNameErc721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setApprovalForAll", operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetApprovalForAll(&_DnsNameErc721.TransactOpts, operator_, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator_, bool approved) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetApprovalForAll(operator_ common.Address, approved bool) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetApprovalForAll(&_DnsNameErc721.TransactOpts, operator_, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetBaseUri(opts *bind.TransactOpts, baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setBaseUri", baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetBaseUri(&_DnsNameErc721.TransactOpts, baseUri_)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string baseUri_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetBaseUri(baseUri_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetBaseUri(&_DnsNameErc721.TransactOpts, baseUri_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetName(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setName", name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetName(&_DnsNameErc721.TransactOpts, name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetName(&_DnsNameErc721.TransactOpts, name_, symbol_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) SetSigUser(opts *bind.TransactOpts, sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "setSigUser", sigUser_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721Session) SetSigUser(sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetSigUser(&_DnsNameErc721.TransactOpts, sigUser_)
}

// SetSigUser is a paid mutator transaction binding the contract method 0xcb50fe76.
//
// Solidity: function setSigUser(address sigUser_) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) SetSigUser(sigUser_ common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.SetSigUser(&_DnsNameErc721.TransactOpts, sigUser_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferFrom(&_DnsNameErc721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferOwnership(&_DnsNameErc721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsNameErc721 *DnsNameErc721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsNameErc721.Contract.TransferOwnership(&_DnsNameErc721.TransactOpts, newOwner)
}

// DnsNameErc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalIterator struct {
	Event *DnsNameErc721Approval // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721Approval)
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
		it.Event = new(DnsNameErc721Approval)
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
func (it *DnsNameErc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721Approval represents a Approval event raised by the DnsNameErc721 contract.
type DnsNameErc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DnsNameErc721ApprovalIterator, error) {

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

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721ApprovalIterator{contract: _DnsNameErc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DnsNameErc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721Approval)
				if err := _DnsNameErc721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseApproval(log types.Log) (*DnsNameErc721Approval, error) {
	event := new(DnsNameErc721Approval)
	if err := _DnsNameErc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameErc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalForAllIterator struct {
	Event *DnsNameErc721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721ApprovalForAll)
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
		it.Event = new(DnsNameErc721ApprovalForAll)
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
func (it *DnsNameErc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721ApprovalForAll represents a ApprovalForAll event raised by the DnsNameErc721 contract.
type DnsNameErc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DnsNameErc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721ApprovalForAllIterator{contract: _DnsNameErc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DnsNameErc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721ApprovalForAll)
				if err := _DnsNameErc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseApprovalForAll(log types.Log) (*DnsNameErc721ApprovalForAll, error) {
	event := new(DnsNameErc721ApprovalForAll)
	if err := _DnsNameErc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsNameErc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DnsNameErc721 contract.
type DnsNameErc721TransferIterator struct {
	Event *DnsNameErc721Transfer // Event containing the contract specifics and raw log

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
func (it *DnsNameErc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsNameErc721Transfer)
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
		it.Event = new(DnsNameErc721Transfer)
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
func (it *DnsNameErc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsNameErc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsNameErc721Transfer represents a Transfer event raised by the DnsNameErc721 contract.
type DnsNameErc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DnsNameErc721TransferIterator, error) {

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

	logs, sub, err := _DnsNameErc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DnsNameErc721TransferIterator{contract: _DnsNameErc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DnsNameErc721 *DnsNameErc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DnsNameErc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DnsNameErc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsNameErc721Transfer)
				if err := _DnsNameErc721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DnsNameErc721 *DnsNameErc721Filterer) ParseTransfer(log types.Log) (*DnsNameErc721Transfer, error) {
	event := new(DnsNameErc721Transfer)
	if err := _DnsNameErc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameMetaData contains all meta data concerning the DnsTopLevelName contract.
var DnsTopLevelNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"EVChargeTopLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeTopLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintTopLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintTopLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"name\":\"EvOpen2Reg\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeTopLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeTopLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"MintTopLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintTopLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"}],\"name\":\"Open2Reg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"contractIDnsAccountant\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costContractAddr\",\"outputs\":[{\"internalType\":\"contractIcost\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsSecond\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"erc721Store\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"}],\"name\":\"getErc721Contract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintSwitch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountantAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsSecond_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mintSw_\",\"type\":\"uint8\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600360146101000a81548160ff021916908360ff1602179055503480156200002d57600080fd5b50604051620047a5380380620047a583398181016040528101906200005391906200024c565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073__$d074600fcdd3cf12194b4b963c2fdc2769$__63651fed368383336000801b6040518563ffffffff1660e01b8152600401620000d5949392919062000334565b60206040518083038186803b158015620000ee57600080fd5b505af415801562000103573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012991906200021a565b600460008060001b815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000522565b60006200019c6200019684620003b8565b6200038f565b905082815260208101848484011115620001bb57620001ba620004e8565b5b620001c884828562000448565b509392505050565b600081519050620001e18162000508565b92915050565b600082601f830112620001ff57620001fe620004e3565b5b81516200021184826020860162000185565b91505092915050565b600060208284031215620002335762000232620004f2565b5b60006200024384828501620001d0565b91505092915050565b60008060408385031215620002665762000265620004f2565b5b600083015167ffffffffffffffff811115620002875762000286620004ed565b5b6200029585828601620001e7565b925050602083015167ffffffffffffffff811115620002b957620002b8620004ed565b5b620002c785828601620001e7565b9150509250929050565b620002dc816200040a565b82525050565b620002ed816200041e565b82525050565b60006200030082620003ee565b6200030c8185620003f9565b93506200031e81856020860162000448565b6200032981620004f7565b840191505092915050565b60006080820190508181036000830152620003508187620002f3565b90508181036020830152620003668186620002f3565b9050620003776040830185620002d1565b620003866060830184620002e2565b95945050505050565b60006200039b620003ae565b9050620003a982826200047e565b919050565b6000604051905090565b600067ffffffffffffffff821115620003d657620003d5620004b4565b5b620003e182620004f7565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000620004178262000428565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b83811015620004685780820151818401526020810190506200044b565b8381111562000478576000848401525b50505050565b6200048982620004f7565b810181811067ffffffffffffffff82111715620004ab57620004aa620004b4565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b62000513816200040a565b81146200051f57600080fd5b50565b61427380620005326000396000f3fe6080604052600436106100f35760003560e01c8063b28ec9d11161008a578063e7f43c6811610059578063e7f43c68146102f6578063eacb912d14610321578063f2fde38b1461034c578063fddd23a814610375576100f3565b8063b28ec9d11461026a578063bc3f419214610293578063bfc0dbf8146102be578063d84c99eb146102da576100f3565b80638a253b2c116100c65780638a253b2c146101ca5780638da5cb5b146101e657806399456e7a146102115780639b737ea51461024e576100f3565b806312014f01146100f85780632c19be76146101235780636130457d14610160578063767013fe1461018b575b600080fd5b34801561010457600080fd5b5061010d61039e565b60405161011a919061360e565b60405180910390f35b34801561012f57600080fd5b5061014a60048036038101906101459190612a99565b6103c4565b6040516101579190613397565b60405180910390f35b34801561016c57600080fd5b506101756104a3565b6040516101829190613629565b60405180910390f35b34801561019757600080fd5b506101b260048036038101906101ad9190612a99565b6104c9565b6040516101c193929190613666565b60405180910390f35b6101e460048036038101906101df9190612cc3565b61058b565b005b3480156101f257600080fd5b506101fb610cb5565b60405161020891906133b2565b60405180910390f35b34801561021d57600080fd5b5061023860048036038101906102339190612a99565b610cd9565b6040516102459190613397565b60405180910390f35b61026860048036038101906102639190612b6e565b610d0c565b005b34801561027657600080fd5b50610291600480360381019061028c9190612a99565b611157565b005b34801561029f57600080fd5b506102a86114a5565b6040516102b59190613397565b60405180910390f35b6102d860048036038101906102d39190612af3565b6114cb565b005b6102f460048036038101906102ef9190612c40565b611706565b005b34801561030257600080fd5b5061030b611be9565b6040516103189190613397565b60405180910390f35b34801561032d57600080fd5b50610336611c13565b604051610343919061398f565b60405180910390f35b34801561035857600080fd5b50610373600480360381019061036e91906129d8565b611c26565b005b34801561038157600080fd5b5061039c60048036038101906103979190612a05565b611d62565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff166004600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610468576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045f9061384b565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60056020528060005260406000206000915090508060000180546104ec90613cee565b80601f016020809104026020016040519081016040528092919081815260200182805461051890613cee565b80156105655780601f1061053a57610100808354040283529160200191610565565b820191906000526020600020905b81548152906001019060200180831161054857829003601f168201915b5050505050908060010160009054906101000a900463ffffffff16908060020154905083565b60006004600360149054906101000a900460ff161660ff16146105e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105da906138ab565b60405180910390fd5b73__$61b1050b44c222c225346b0c1857883025$__633c1f718c886040518263ffffffff1660e01b815260040161061a91906135ec565b60206040518083038186803b15801561063257600080fd5b505af4158015610646573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066a9190612a6c565b6106a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a09061392b565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63d7eafff3896040518263ffffffff1660e01b81526004016106e29190613644565b60206040518083038186803b1580156106fa57600080fd5b505af415801561070e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107329190612ac6565b905060006005600083815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff16146107a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079a9061386b565b60405180910390fd5b73__$2ed95cc1c04a020bf25a1d27d6730d85d6$__6354ed93b2898989898989336040516020016107da97969594939291906132f8565b60405160208183030381529060405280519060200120846040518363ffffffff1660e01b815260040161080e9291906134df565b60206040518083038186803b15801561082657600080fd5b505af415801561083a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085e91906129ab565b73ffffffffffffffffffffffffffffffffffffffff16600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633cf8baa26040518163ffffffff1660e01b815260040160206040518083038186803b1580156108f057600080fd5b505afa158015610904573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092891906129ab565b73ffffffffffffffffffffffffffffffffffffffff161461097e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610975906138cb565b60405180910390fd5b6000831115610a88576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395297fb088886040518363ffffffff1660e01b81526004016109e6929190613456565b60206040518083038186803b1580156109fe57600080fd5b505afa158015610a12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a369190612a6c565b905080610a78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6f906137eb565b60405180910390fd5b610a828785611e9e565b50610ad1565b8463ffffffff164210610ad0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac79061380b565b60405180910390fd5b5b60405180606001604052808981526020016301e133808960ff16610af59190613b5b565b42610b009190613af0565b63ffffffff168152602001600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663adfd5f91846301e133808c60ff16610b719190613b5b565b42610b7c9190613af0565b336040518463ffffffff1660e01b8152600401610b9b9392919061350f565b602060405180830381600087803b158015610bb557600080fd5b505af1158015610bc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bed9190612d9d565b815250600560008381526020019081526020016000206000820151816000019080519060200190610c1f929190612741565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff160217905550604082015181600201559050507f909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad1807488888887876005600088815260200190815260200160002060020154604051610ca3969594939291906136f0565b60405180910390a15050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006008600360149054906101000a900460ff161660ff1614610d64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5b906137ab565b60405180910390fd5b6000600560008a815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff1611610dd3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dca9061388b565b60405180910390fd5b73__$2ed95cc1c04a020bf25a1d27d6730d85d6$__6354ed93b289898989898933604051602001610e0a9796959493929190613277565b60405160208183030381529060405280519060200120846040518363ffffffff1660e01b8152600401610e3e9291906134df565b60206040518083038186803b158015610e5657600080fd5b505af4158015610e6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8e91906129ab565b73ffffffffffffffffffffffffffffffffffffffff16600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633cf8baa26040518163ffffffff1660e01b815260040160206040518083038186803b158015610f2057600080fd5b505afa158015610f34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5891906129ab565b73ffffffffffffffffffffffffffffffffffffffff1614610fae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa5906138cb565b60405180910390fd5b60008311156110b8576000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395297fb088886040518363ffffffff1660e01b8152600401611016929190613456565b60206040518083038186803b15801561102e57600080fd5b505afa158015611042573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110669190612a6c565b9050806110a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109f906137eb565b60405180910390fd5b6110b28785611e9e565b50611101565b8463ffffffff164210611100576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f79061380b565b60405180910390fd5b5b61110c88888361236d565b7fcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad8888888787866040516111459695949392919061358b565b60405180910390a15050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146111f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f0906138eb565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e60056000858152602001908152602001600020600201546040518263ffffffff1660e01b8152600401611295919061394b565b60206040518083038186803b1580156112ad57600080fd5b505afa1580156112c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e591906129ab565b73ffffffffffffffffffffffffffffffffffffffff161461133b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113329061390b565b60405180910390fd5b73__$d074600fcdd3cf12194b4b963c2fdc2769$__63651fed366005600084815260200190815260200160002060000161138a60056000868152602001908152602001600020600201546125e0565b60405160200161139a9190613375565b60405160208183030381529060405233856040518563ffffffff1660e01b81526004016113ca9493929190613758565b60206040518083038186803b1580156113e257600080fd5b505af41580156113f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141a91906129ab565b6004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a5388160405161149a91906134c4565b60405180910390a150565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006002600360149054906101000a900460ff161660ff1614611523576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151a906137ab565b60405180910390fd5b60006005600087815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff1611611592576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115899061388b565b60405180910390fd5b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633dd45f138686600560008c815260200190815260200160002060000180546115f690613cee565b90508a6040518563ffffffff1660e01b8152600401611618949392919061347f565b604080518083038186803b15801561162f57600080fd5b505afa158015611643573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116679190612dca565b91509150806116ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116a2906137eb565b60405180910390fd5b6116b58583611e9e565b6116c087878561236d565b7fc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5878787866040516116f59493929190613546565b60405180910390a150505050505050565b60006001600360149054906101000a900460ff161660ff161461175e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611755906138ab565b60405180910390fd5b73__$61b1050b44c222c225346b0c1857883025$__633c1f718c856040518263ffffffff1660e01b815260040161179591906135ec565b60206040518083038186803b1580156117ad57600080fd5b505af41580156117c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e59190612a6c565b611824576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161181b9061392b565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63d7eafff3866040518263ffffffff1660e01b815260040161185d9190613644565b60206040518083038186803b15801561187557600080fd5b505af4158015611889573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ad9190612ac6565b90506000600560008381526020019081526020016000206002015414611908576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ff9061386b565b60405180910390fd5b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633dd45f1386868a518a6040518563ffffffff1660e01b815260040161196d949392919061347f565b604080518083038186803b15801561198457600080fd5b505afa158015611998573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119bc9190612dca565b9150915080611a00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119f7906137eb565b60405180910390fd5b611a0a8583611e9e565b60405180606001604052808881526020016301e133808860ff16611a2e9190613b5b565b42611a399190613af0565b63ffffffff168152602001600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663adfd5f91866301e133808b60ff16611aaa9190613b5b565b42611ab59190613af0565b336040518463ffffffff1660e01b8152600401611ad49392919061350f565b602060405180830381600087803b158015611aee57600080fd5b505af1158015611b02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b269190612d9d565b815250600560008581526020019081526020016000206000820151816000019080519060200190611b58929190612741565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff160217905550604082015181600201559050507fc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e8787876005600088815260200190815260200160002060020154604051611bd894939291906136a4565b60405180910390a150505050505050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600360149054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611c7e57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f2fde38b826040518263ffffffff1660e01b8152600401611d2d91906133b2565b600060405180830381600087803b158015611d4757600080fd5b505af1158015611d5b573d6000803e3d6000fd5b5050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611dba57600080fd5b83600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360146101000a81548160ff021916908360ff16021790555050505050565b60008111611ee1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ed89061382b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561208e5780341015611f59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f50906137cb565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684346040518463ffffffff1660e01b8152600401611fee939291906133f6565b600060405180830381600087803b15801561200857600080fd5b505af115801561201c573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015612088573d6000803e3d6000fd5b50612369565b808273ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016120c89190613397565b60206040518083038186803b1580156120e057600080fd5b505afa1580156120f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121189190612d9d565b101580156121b05750808273ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161215d9291906133cd565b60206040518083038186803b15801561217557600080fd5b505afa158015612189573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ad9190612d9d565b10155b6121ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121e6906137cb565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684846040518463ffffffff1660e01b8152600401612284939291906133f6565b600060405180830381600087803b15801561229e57600080fd5b505af11580156122b2573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166323b872dd33600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401612315939291906133f6565b602060405180830381600087803b15801561232f57600080fd5b505af1158015612343573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123679190612a6c565b505b5050565b6005600084815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff1642111561249757801561246157600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c8eba7603360056000878152602001908152602001600020600201546040518363ffffffff1660e01b815260040161242e92919061342d565b600060405180830381600087803b15801561244857600080fd5b505af115801561245c573d6000803e3d6000fd5b505050505b426005600085815260200190815260200160002060010160006101000a81548163ffffffff021916908363ffffffff1602179055505b6301e133808260ff166124aa9190613b5b565b6005600085815260200190815260200160002060010160008282829054906101000a900463ffffffff166124de9190613af0565b92506101000a81548163ffffffff021916908363ffffffff160217905550600460008060001b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e767f8fd60056000868152602001908152602001600020600201546005600087815260200190815260200160002060010160009054906101000a900463ffffffff166040518363ffffffff1660e01b81526004016125a9929190613966565b600060405180830381600087803b1580156125c357600080fd5b505af11580156125d7573d6000803e3d6000fd5b50505050505050565b60606000821415612628576040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815250905061273c565b600082905060005b6000821461265a57808061264390613d51565b915050600a826126539190613b2a565b9150612630565b60008167ffffffffffffffff81111561267657612675613ee3565b5b6040519080825280601f01601f1916602001820160405280156126a85781602001600182028036833780820191505090505b5090505b60008514612735576001826126c19190613b99565b9150600a856126d09190613df6565b60306126dc9190613a9a565b60f81b8183815181106126f2576126f1613eb4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a8561272e9190613b2a565b94506126ac565b8093505050505b919050565b82805461274d90613cee565b90600052602060002090601f01602090048101928261276f57600085556127b6565b82601f1061278857805160ff19168380011785556127b6565b828001600101855582156127b6579182015b828111156127b557825182559160200191906001019061279a565b5b5090506127c391906127c7565b5090565b5b808211156127e05760008160009055506001016127c8565b5090565b60006127f76127f2846139cf565b6139aa565b90508281526020810184848401111561281357612812613f17565b5b61281e848285613cac565b509392505050565b600061283961283484613a00565b6139aa565b90508281526020810184848401111561285557612854613f17565b5b612860848285613cac565b509392505050565b6000813590506128778161419c565b92915050565b60008151905061288c8161419c565b92915050565b6000813590506128a1816141b3565b92915050565b6000813590506128b6816141ca565b92915050565b6000815190506128cb816141ca565b92915050565b6000813590506128e0816141e1565b92915050565b6000815190506128f5816141e1565b92915050565b600082601f8301126129105761290f613f12565b5b81356129208482602086016127e4565b91505092915050565b600082601f83011261293e5761293d613f12565b5b813561294e848260208601612826565b91505092915050565b600081359050612966816141f8565b92915050565b60008151905061297b816141f8565b92915050565b6000813590506129908161420f565b92915050565b6000813590506129a581614226565b92915050565b6000602082840312156129c1576129c0613f21565b5b60006129cf8482850161287d565b91505092915050565b6000602082840312156129ee576129ed613f21565b5b60006129fc84828501612892565b91505092915050565b60008060008060808587031215612a1f57612a1e613f21565b5b6000612a2d87828801612868565b9450506020612a3e87828801612868565b9350506040612a4f87828801612868565b9250506060612a6087828801612981565b91505092959194509250565b600060208284031215612a8257612a81613f21565b5b6000612a90848285016128bc565b91505092915050565b600060208284031215612aaf57612aae613f21565b5b6000612abd848285016128d1565b91505092915050565b600060208284031215612adc57612adb613f21565b5b6000612aea848285016128e6565b91505092915050565b600080600080600060a08688031215612b0f57612b0e613f21565b5b6000612b1d888289016128d1565b9550506020612b2e88828901612981565b9450506040612b3f88828901612868565b9350506060612b5088828901612996565b9250506080612b61888289016128a7565b9150509295509295909350565b600080600080600080600080610100898b031215612b8f57612b8e613f21565b5b6000612b9d8b828c016128d1565b9850506020612bae8b828c01612981565b9750506040612bbf8b828c01612868565b9650506060612bd08b828c01612996565b9550506080612be18b828c01612957565b94505060a0612bf28b828c01612957565b93505060c089013567ffffffffffffffff811115612c1357612c12613f1c565b5b612c1f8b828c016128fb565b92505060e0612c308b828c016128a7565b9150509295985092959890939650565b60008060008060808587031215612c5a57612c59613f21565b5b600085013567ffffffffffffffff811115612c7857612c77613f1c565b5b612c8487828801612929565b9450506020612c9587828801612981565b9350506040612ca687828801612868565b9250506060612cb787828801612996565b91505092959194509250565b600080600080600080600060e0888a031215612ce257612ce1613f21565b5b600088013567ffffffffffffffff811115612d0057612cff613f1c565b5b612d0c8a828b01612929565b9750506020612d1d8a828b01612981565b9650506040612d2e8a828b01612868565b9550506060612d3f8a828b01612996565b9450506080612d508a828b01612957565b93505060a0612d618a828b01612957565b92505060c088013567ffffffffffffffff811115612d8257612d81613f1c565b5b612d8e8a828b016128fb565b91505092959891949750929550565b600060208284031215612db357612db2613f21565b5b6000612dc18482850161296c565b91505092915050565b60008060408385031215612de157612de0613f21565b5b6000612def8582860161296c565b9250506020612e00858286016128bc565b9150509250929050565b612e1381613bdf565b82525050565b612e2281613bcd565b82525050565b612e3181613bcd565b82525050565b612e48612e4382613bcd565b613d9a565b82525050565b612e5781613bf1565b82525050565b612e6681613bfd565b82525050565b612e7581613bfd565b82525050565b612e8c612e8782613bfd565b613dac565b82525050565b6000612e9d82613a46565b612ea78185613a5c565b9350612eb7818560208601613cbb565b612ec081613f26565b840191505092915050565b612ed481613c64565b82525050565b612ee381613c88565b82525050565b6000612ef482613a51565b612efe8185613a6d565b9350612f0e818560208601613cbb565b612f1781613f26565b840191505092915050565b6000612f2d82613a51565b612f378185613a7e565b9350612f47818560208601613cbb565b612f5081613f26565b840191505092915050565b6000612f6682613a51565b612f708185613a8f565b9350612f80818560208601613cbb565b80840191505092915050565b60008154612f9981613cee565b612fa38186613a7e565b94506001821660008114612fbe5760018114612fd057613003565b60ff1983168652602086019350613003565b612fd985613a31565b60005b83811015612ffb57815481890152600182019150602081019050612fdc565b808801955050505b50505092915050565b6000613019600c83613a6d565b915061302482613f5e565b602082019050919050565b600061303c601483613a6d565b915061304782613f87565b602082019050919050565b600061305f600f83613a6d565b915061306a82613fb0565b602082019050919050565b6000613082600b83613a6d565b915061308d82613fd9565b602082019050919050565b60006130a5601083613a6d565b91506130b082614002565b602082019050919050565b60006130c8601083613a6d565b91506130d38261402b565b602082019050919050565b60006130eb601383613a6d565b91506130f682614054565b602082019050919050565b600061310e600e83613a6d565b91506131198261407d565b602082019050919050565b6000613131600a83613a6d565b915061313c826140a6565b602082019050919050565b6000613154600783613a8f565b915061315f826140cf565b600782019050919050565b6000613177600d83613a6d565b9150613182826140f8565b602082019050919050565b600061319a600683613a6d565b91506131a582614121565b602082019050919050565b60006131bd600983613a6d565b91506131c88261414a565b602082019050919050565b60006131e0600f83613a6d565b91506131eb82614173565b602082019050919050565b6131ff81613c27565b82525050565b61321661321182613c27565b613dc8565b82525050565b61322581613c31565b82525050565b61323481613c4e565b82525050565b61324b61324682613c4e565b613de4565b82525050565b61325a81613c41565b82525050565b61327161326c82613c41565b613dd2565b82525050565b6000613283828a612e7b565b6020820191506132938289613260565b6001820191506132a38288612e37565b6014820191506132b3828761323a565b600a820191506132c38286613205565b6020820191506132d38285613205565b6020820191506132e38284612e37565b60148201915081905098975050505050505050565b6000613304828a612f5b565b91506133108289613260565b6001820191506133208288612e37565b601482019150613330828761323a565b600a820191506133408286613205565b6020820191506133508285613205565b6020820191506133608284612e37565b60148201915081905098975050505050505050565b600061338082613147565b915061338c8284612f5b565b915081905092915050565b60006020820190506133ac6000830184612e19565b92915050565b60006020820190506133c76000830184612e0a565b92915050565b60006040820190506133e26000830185612e19565b6133ef6020830184612e19565b9392505050565b600060608201905061340b6000830186612e19565b6134186020830185612e19565b61342560408301846131f6565b949350505050565b60006040820190506134426000830185612e19565b61344f60208301846131f6565b9392505050565b600060408201905061346b6000830185612e19565b613478602083018461322b565b9392505050565b60006080820190506134946000830187612e19565b6134a1602083018661322b565b6134ae6040830185613251565b6134bb6060830184613251565b95945050505050565b60006020820190506134d96000830184612e5d565b92915050565b60006040820190506134f46000830185612e6c565b81810360208301526135068184612e92565b90509392505050565b60006060820190506135246000830186612e5d565b613531602083018561321c565b61353e6040830184612e19565b949350505050565b600060808201905061355b6000830187612e5d565b6135686020830186613251565b6135756040830185612e19565b6135826060830184612e4e565b95945050505050565b600060c0820190506135a06000830189612e5d565b6135ad6020830188613251565b6135ba6040830187612e19565b6135c760608301866131f6565b6135d460808301856131f6565b6135e160a0830184612e4e565b979650505050505050565b600060208201905081810360008301526136068184612e92565b905092915050565b60006020820190506136236000830184612ecb565b92915050565b600060208201905061363e6000830184612eda565b92915050565b6000602082019050818103600083015261365e8184612f22565b905092915050565b600060608201905081810360008301526136808186612ee9565b905061368f602083018561321c565b61369c60408301846131f6565b949350505050565b600060808201905081810360008301526136be8187612ee9565b90506136cd6020830186613251565b6136da6040830185612e19565b6136e760608301846131f6565b95945050505050565b600060c082019050818103600083015261370a8189612ee9565b90506137196020830188613251565b6137266040830187612e19565b61373360608301866131f6565b61374060808301856131f6565b61374d60a08301846131f6565b979650505050505050565b600060808201905081810360008301526137728187612f8c565b905081810360208301526137868186612f22565b90506137956040830185612e28565b6137a26060830184612e6c565b95945050505050565b600060208201905081810360008301526137c48161300c565b9050919050565b600060208201905081810360008301526137e48161302f565b9050919050565b6000602082019050818103600083015261380481613052565b9050919050565b6000602082019050818103600083015261382481613075565b9050919050565b6000602082019050818103600083015261384481613098565b9050919050565b60006020820190508181036000830152613864816130bb565b9050919050565b60006020820190508181036000830152613884816130de565b9050919050565b600060208201905081810360008301526138a481613101565b9050919050565b600060208201905081810360008301526138c481613124565b9050919050565b600060208201905081810360008301526138e48161316a565b9050919050565b600060208201905081810360008301526139048161318d565b9050919050565b60006020820190508181036000830152613924816131b0565b9050919050565b60006020820190508181036000830152613944816131d3565b9050919050565b600060208201905061396060008301846131f6565b92915050565b600060408201905061397b60008301856131f6565b613988602083018461321c565b9392505050565b60006020820190506139a46000830184613251565b92915050565b60006139b46139c5565b90506139c08282613d20565b919050565b6000604051905090565b600067ffffffffffffffff8211156139ea576139e9613ee3565b5b6139f382613f26565b9050602081019050919050565b600067ffffffffffffffff821115613a1b57613a1a613ee3565b5b613a2482613f26565b9050602081019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000613aa582613c27565b9150613ab083613c27565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115613ae557613ae4613e27565b5b828201905092915050565b6000613afb82613c31565b9150613b0683613c31565b92508263ffffffff03821115613b1f57613b1e613e27565b5b828201905092915050565b6000613b3582613c27565b9150613b4083613c27565b925082613b5057613b4f613e56565b5b828204905092915050565b6000613b6682613c31565b9150613b7183613c31565b92508163ffffffff0483118215151615613b8e57613b8d613e27565b5b828202905092915050565b6000613ba482613c27565b9150613baf83613c27565b925082821015613bc257613bc1613e27565b5b828203905092915050565b6000613bd882613c07565b9050919050565b6000613bea82613c07565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b6000613c6f82613c76565b9050919050565b6000613c8182613c07565b9050919050565b6000613c9382613c9a565b9050919050565b6000613ca582613c07565b9050919050565b82818337600083830152505050565b60005b83811015613cd9578082015181840152602081019050613cbe565b83811115613ce8576000848401525b50505050565b60006002820490506001821680613d0657607f821691505b60208210811415613d1a57613d19613e85565b5b50919050565b613d2982613f26565b810181811067ffffffffffffffff82111715613d4857613d47613ee3565b5b80604052505050565b6000613d5c82613c27565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415613d8f57613d8e613e27565b5b600182019050919050565b6000613da582613db6565b9050919050565b6000819050919050565b6000613dc182613f51565b9050919050565b6000819050919050565b6000613ddd82613f44565b9050919050565b6000613def82613f37565b9050919050565b6000613e0182613c27565b9150613e0c83613c27565b925082613e1c57613e1b613e56565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160b01b9050919050565b60008160f81b9050919050565b60008160601b9050919050565b7f63616e2774206368617267650000000000000000000000000000000000000000600082015250565b7f7061796f7574206973206e6f7420656e6f756768000000000000000000000000600082015250565b7f7072696365206e6f742076616c69640000000000000000000000000000000000600082015250565b7f7369672065787069726564000000000000000000000000000000000000000000600082015250565b7f6e6f7420612076616c696420636f737400000000000000000000000000000000600082015250565b7f657263373231206e6f7420666f756e6400000000000000000000000000000000600082015250565b7f6e616d6520776173207265676973746572656400000000000000000000000000600082015250565b7f6e616d65206e6f74206578697374000000000000000000000000000000000000600082015250565b7f63616e2774206d696e7400000000000000000000000000000000000000000000600082015250565b7f646e7344616f2300000000000000000000000000000000000000000000000000600082015250565b7f736967206e6f74206d6174636800000000000000000000000000000000000000600082015250565b7f6f70656e65640000000000000000000000000000000000000000000000000000600082015250565b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b7f6e6f74206120726f6f74204e616d650000000000000000000000000000000000600082015250565b6141a581613bcd565b81146141b057600080fd5b50565b6141bc81613bdf565b81146141c757600080fd5b50565b6141d381613bf1565b81146141de57600080fd5b50565b6141ea81613bfd565b81146141f557600080fd5b50565b61420181613c27565b811461420c57600080fd5b50565b61421881613c41565b811461422357600080fd5b50565b61422f81613c4e565b811461423a57600080fd5b5056fea2646970667358221220e736b0a10b7bc4fb0f086f52ad3e4a7eb9918135485a74bf8296852cf92dbb5164736f6c63430008060033",
}

// DnsTopLevelNameABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsTopLevelNameMetaData.ABI instead.
var DnsTopLevelNameABI = DnsTopLevelNameMetaData.ABI

// DnsTopLevelNameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DnsTopLevelNameMetaData.Bin instead.
var DnsTopLevelNameBin = DnsTopLevelNameMetaData.Bin

// DeployDnsTopLevelName deploys a new Ethereum contract, binding an instance of DnsTopLevelName to it.
func DeployDnsTopLevelName(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *DnsTopLevelName, error) {
	parsed, err := DnsTopLevelNameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	libDnsSignatureAddr, _, _, _ := DeployLibDnsSignature(auth, backend)
	DnsTopLevelNameBin = strings.ReplaceAll(DnsTopLevelNameBin, "__$2ed95cc1c04a020bf25a1d27d6730d85d6$__", libDnsSignatureAddr.String()[2:])

	libDnsToolKitAddr, _, _, _ := DeployLibDnsToolKit(auth, backend)
	DnsTopLevelNameBin = strings.ReplaceAll(DnsTopLevelNameBin, "__$61b1050b44c222c225346b0c1857883025$__", libDnsToolKitAddr.String()[2:])

	libDnsNameErc721Addr, _, _, _ := DeployLibDnsNameErc721(auth, backend)
	DnsTopLevelNameBin = strings.ReplaceAll(DnsTopLevelNameBin, "__$d074600fcdd3cf12194b4b963c2fdc2769$__", libDnsNameErc721Addr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DnsTopLevelNameBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DnsTopLevelName{DnsTopLevelNameCaller: DnsTopLevelNameCaller{contract: contract}, DnsTopLevelNameTransactor: DnsTopLevelNameTransactor{contract: contract}, DnsTopLevelNameFilterer: DnsTopLevelNameFilterer{contract: contract}}, nil
}

// DnsTopLevelName is an auto generated Go binding around an Ethereum contract.
type DnsTopLevelName struct {
	DnsTopLevelNameCaller     // Read-only binding to the contract
	DnsTopLevelNameTransactor // Write-only binding to the contract
	DnsTopLevelNameFilterer   // Log filterer for contract events
}

// DnsTopLevelNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsTopLevelNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsTopLevelNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsTopLevelNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsTopLevelNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsTopLevelNameSession struct {
	Contract     *DnsTopLevelName  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsTopLevelNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsTopLevelNameCallerSession struct {
	Contract *DnsTopLevelNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DnsTopLevelNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsTopLevelNameTransactorSession struct {
	Contract     *DnsTopLevelNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DnsTopLevelNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsTopLevelNameRaw struct {
	Contract *DnsTopLevelName // Generic contract binding to access the raw methods on
}

// DnsTopLevelNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsTopLevelNameCallerRaw struct {
	Contract *DnsTopLevelNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsTopLevelNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsTopLevelNameTransactorRaw struct {
	Contract *DnsTopLevelNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsTopLevelName creates a new instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelName(address common.Address, backend bind.ContractBackend) (*DnsTopLevelName, error) {
	contract, err := bindDnsTopLevelName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelName{DnsTopLevelNameCaller: DnsTopLevelNameCaller{contract: contract}, DnsTopLevelNameTransactor: DnsTopLevelNameTransactor{contract: contract}, DnsTopLevelNameFilterer: DnsTopLevelNameFilterer{contract: contract}}, nil
}

// NewDnsTopLevelNameCaller creates a new read-only instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameCaller(address common.Address, caller bind.ContractCaller) (*DnsTopLevelNameCaller, error) {
	contract, err := bindDnsTopLevelName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameCaller{contract: contract}, nil
}

// NewDnsTopLevelNameTransactor creates a new write-only instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsTopLevelNameTransactor, error) {
	contract, err := bindDnsTopLevelName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameTransactor{contract: contract}, nil
}

// NewDnsTopLevelNameFilterer creates a new log filterer instance of DnsTopLevelName, bound to a specific deployed contract.
func NewDnsTopLevelNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsTopLevelNameFilterer, error) {
	contract, err := bindDnsTopLevelName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameFilterer{contract: contract}, nil
}

// bindDnsTopLevelName binds a generic wrapper to an already deployed contract.
func bindDnsTopLevelName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsTopLevelNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsTopLevelName.Contract.DnsTopLevelNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.DnsTopLevelNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsTopLevelName *DnsTopLevelNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.DnsTopLevelNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsTopLevelName *DnsTopLevelNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsTopLevelName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsTopLevelName *DnsTopLevelNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsTopLevelName *DnsTopLevelNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) AccountantC() (common.Address, error) {
	return _DnsTopLevelName.Contract.AccountantC(&_DnsTopLevelName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsTopLevelName.Contract.AccountantC(&_DnsTopLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) CostContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "costContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) CostContractAddr() (common.Address, error) {
	return _DnsTopLevelName.Contract.CostContractAddr(&_DnsTopLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) CostContractAddr() (common.Address, error) {
	return _DnsTopLevelName.Contract.CostContractAddr(&_DnsTopLevelName.CallOpts)
}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) DnsSecond(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "dnsSecond")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) DnsSecond() (common.Address, error) {
	return _DnsTopLevelName.Contract.DnsSecond(&_DnsTopLevelName.CallOpts)
}

// DnsSecond is a free data retrieval call binding the contract method 0xbc3f4192.
//
// Solidity: function dnsSecond() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) DnsSecond() (common.Address, error) {
	return _DnsTopLevelName.Contract.DnsSecond(&_DnsTopLevelName.CallOpts)
}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) Erc721Store(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "erc721Store", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) Erc721Store(arg0 [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.Erc721Store(&_DnsTopLevelName.CallOpts, arg0)
}

// Erc721Store is a free data retrieval call binding the contract method 0x99456e7a.
//
// Solidity: function erc721Store(bytes32 ) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) Erc721Store(arg0 [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.Erc721Store(&_DnsTopLevelName.CallOpts, arg0)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) GetErc721Contract(opts *bind.CallOpts, fatherHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "getErc721Contract", fatherHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.GetErc721Contract(&_DnsTopLevelName.CallOpts, fatherHash)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _DnsTopLevelName.Contract.GetErc721Contract(&_DnsTopLevelName.CallOpts, fatherHash)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) GetOperator() (common.Address, error) {
	return _DnsTopLevelName.Contract.GetOperator(&_DnsTopLevelName.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) GetOperator() (common.Address, error) {
	return _DnsTopLevelName.Contract.GetOperator(&_DnsTopLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameCaller) MintSwitch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "mintSwitch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameSession) MintSwitch() (uint8, error) {
	return _DnsTopLevelName.Contract.MintSwitch(&_DnsTopLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) MintSwitch() (uint8, error) {
	return _DnsTopLevelName.Contract.MintSwitch(&_DnsTopLevelName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "nameStore", arg0)

	outstruct := new(struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EntireName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ExpireTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameSession) NameStore(arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsTopLevelName.Contract.NameStore(&_DnsTopLevelName.CallOpts, arg0)
}

// NameStore is a free data retrieval call binding the contract method 0x767013fe.
//
// Solidity: function nameStore(bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) NameStore(arg0 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsTopLevelName.Contract.NameStore(&_DnsTopLevelName.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsTopLevelName.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameSession) Owner() (common.Address, error) {
	return _DnsTopLevelName.Contract.Owner(&_DnsTopLevelName.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DnsTopLevelName *DnsTopLevelNameCallerSession) Owner() (common.Address, error) {
	return _DnsTopLevelName.Contract.Owner(&_DnsTopLevelName.CallOpts)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) ChargeTopLevelName(opts *bind.TransactOpts, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "ChargeTopLevelName", nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) ChargeTopLevelName(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelName(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelName is a paid mutator transaction binding the contract method 0xbfc0dbf8.
//
// Solidity: function ChargeTopLevelName(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) ChargeTopLevelName(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelName(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) ChargeTopLevelNameBySig(opts *bind.TransactOpts, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "ChargeTopLevelNameBySig", nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) ChargeTopLevelNameBySig(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeTopLevelNameBySig is a paid mutator transaction binding the contract method 0x9b737ea5.
//
// Solidity: function ChargeTopLevelNameBySig(bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) ChargeTopLevelNameBySig(nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.ChargeTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) MintTopLevelName(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "MintTopLevelName", entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) MintTopLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelName(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelName is a paid mutator transaction binding the contract method 0xd84c99eb.
//
// Solidity: function MintTopLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) MintTopLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelName(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) MintTopLevelNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "MintTopLevelNameBySig", entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) MintTopLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintTopLevelNameBySig is a paid mutator transaction binding the contract method 0x8a253b2c.
//
// Solidity: function MintTopLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) MintTopLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.MintTopLevelNameBySig(&_DnsTopLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) Open2Reg(opts *bind.TransactOpts, nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "Open2Reg", nameHash_)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) Open2Reg(nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.Open2Reg(&_DnsTopLevelName.TransactOpts, nameHash_)
}

// Open2Reg is a paid mutator transaction binding the contract method 0xb28ec9d1.
//
// Solidity: function Open2Reg(bytes32 nameHash_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) Open2Reg(nameHash_ [32]byte) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.Open2Reg(&_DnsTopLevelName.TransactOpts, nameHash_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) SetContract(opts *bind.TransactOpts, costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "setContract", costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) SetContract(costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.SetContract(&_DnsTopLevelName.TransactOpts, costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address costAddr_, address accountantAddr_, address dnsSecond_, uint8 mintSw_) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) SetContract(costAddr_ common.Address, accountantAddr_ common.Address, dnsSecond_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.SetContract(&_DnsTopLevelName.TransactOpts, costAddr_, accountantAddr_, dnsSecond_, mintSw_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.TransferOwnership(&_DnsTopLevelName.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DnsTopLevelName *DnsTopLevelNameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DnsTopLevelName.Contract.TransferOwnership(&_DnsTopLevelName.TransactOpts, newOwner)
}

// DnsTopLevelNameEVChargeTopLevelNameIterator is returned from FilterEVChargeTopLevelName and is used to iterate over the raw logs and unpacked data for EVChargeTopLevelName events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEVChargeTopLevelNameIterator struct {
	Event *DnsTopLevelNameEVChargeTopLevelName // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEVChargeTopLevelName)
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
		it.Event = new(DnsTopLevelNameEVChargeTopLevelName)
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
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEVChargeTopLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEVChargeTopLevelName represents a EVChargeTopLevelName event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEVChargeTopLevelName struct {
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEVChargeTopLevelName is a free log retrieval operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEVChargeTopLevelName(opts *bind.FilterOpts) (*DnsTopLevelNameEVChargeTopLevelNameIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EVChargeTopLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEVChargeTopLevelNameIterator{contract: _DnsTopLevelName.contract, event: "EVChargeTopLevelName", logs: logs, sub: sub}, nil
}

// WatchEVChargeTopLevelName is a free log subscription operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEVChargeTopLevelName(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEVChargeTopLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EVChargeTopLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEVChargeTopLevelName)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EVChargeTopLevelName", log); err != nil {
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

// ParseEVChargeTopLevelName is a log parse operation binding the contract event 0xc0090001001b9a9a116dbf8366eb890e513553c35592a8bcdf048fccae2fe6d5.
//
// Solidity: event EVChargeTopLevelName(bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer_)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEVChargeTopLevelName(log types.Log) (*DnsTopLevelNameEVChargeTopLevelName, error) {
	event := new(DnsTopLevelNameEVChargeTopLevelName)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EVChargeTopLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvChargeTopLevelNameBySigIterator is returned from FilterEvChargeTopLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvChargeTopLevelNameBySig events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvChargeTopLevelNameBySigIterator struct {
	Event *DnsTopLevelNameEvChargeTopLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvChargeTopLevelNameBySig)
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
		it.Event = new(DnsTopLevelNameEvChargeTopLevelNameBySig)
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
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvChargeTopLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvChargeTopLevelNameBySig represents a EvChargeTopLevelNameBySig event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvChargeTopLevelNameBySig struct {
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeTopLevelNameBySig is a free log retrieval operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvChargeTopLevelNameBySig(opts *bind.FilterOpts) (*DnsTopLevelNameEvChargeTopLevelNameBySigIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvChargeTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvChargeTopLevelNameBySigIterator{contract: _DnsTopLevelName.contract, event: "EvChargeTopLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvChargeTopLevelNameBySig is a free log subscription operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvChargeTopLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvChargeTopLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvChargeTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvChargeTopLevelNameBySig)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvChargeTopLevelNameBySig", log); err != nil {
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

// ParseEvChargeTopLevelNameBySig is a log parse operation binding the contract event 0xcca122d40cb0eeec3fa2e8383178d2ecc5932ae99c5d49a28223721251df67ad.
//
// Solidity: event EvChargeTopLevelNameBySig(bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvChargeTopLevelNameBySig(log types.Log) (*DnsTopLevelNameEvChargeTopLevelNameBySig, error) {
	event := new(DnsTopLevelNameEvChargeTopLevelNameBySig)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvChargeTopLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvMintTopLevelNameIterator is returned from FilterEvMintTopLevelName and is used to iterate over the raw logs and unpacked data for EvMintTopLevelName events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameIterator struct {
	Event *DnsTopLevelNameEvMintTopLevelName // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvMintTopLevelName)
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
		it.Event = new(DnsTopLevelNameEvMintTopLevelName)
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
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvMintTopLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvMintTopLevelName represents a EvMintTopLevelName event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelName struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintTopLevelName is a free log retrieval operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvMintTopLevelName(opts *bind.FilterOpts) (*DnsTopLevelNameEvMintTopLevelNameIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvMintTopLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvMintTopLevelNameIterator{contract: _DnsTopLevelName.contract, event: "EvMintTopLevelName", logs: logs, sub: sub}, nil
}

// WatchEvMintTopLevelName is a free log subscription operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvMintTopLevelName(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvMintTopLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvMintTopLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvMintTopLevelName)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelName", log); err != nil {
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

// ParseEvMintTopLevelName is a log parse operation binding the contract event 0xc4a8f2e0e5841ebdb7e2795b609a11f4854c3f11feb7b06ef1e6257352923f1e.
//
// Solidity: event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvMintTopLevelName(log types.Log) (*DnsTopLevelNameEvMintTopLevelName, error) {
	event := new(DnsTopLevelNameEvMintTopLevelName)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvMintTopLevelNameBySigIterator is returned from FilterEvMintTopLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvMintTopLevelNameBySig events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameBySigIterator struct {
	Event *DnsTopLevelNameEvMintTopLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvMintTopLevelNameBySig)
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
		it.Event = new(DnsTopLevelNameEvMintTopLevelNameBySig)
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
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvMintTopLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvMintTopLevelNameBySig represents a EvMintTopLevelNameBySig event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvMintTopLevelNameBySig struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintTopLevelNameBySig is a free log retrieval operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvMintTopLevelNameBySig(opts *bind.FilterOpts) (*DnsTopLevelNameEvMintTopLevelNameBySigIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvMintTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvMintTopLevelNameBySigIterator{contract: _DnsTopLevelName.contract, event: "EvMintTopLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvMintTopLevelNameBySig is a free log subscription operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvMintTopLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvMintTopLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvMintTopLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvMintTopLevelNameBySig)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelNameBySig", log); err != nil {
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

// ParseEvMintTopLevelNameBySig is a log parse operation binding the contract event 0x909b7e0a99499599cc97db3c9357fbbf86a5a5763242ec72c0f641dd4ad18074.
//
// Solidity: event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvMintTopLevelNameBySig(log types.Log) (*DnsTopLevelNameEvMintTopLevelNameBySig, error) {
	event := new(DnsTopLevelNameEvMintTopLevelNameBySig)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvMintTopLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsTopLevelNameEvOpen2RegIterator is returned from FilterEvOpen2Reg and is used to iterate over the raw logs and unpacked data for EvOpen2Reg events raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvOpen2RegIterator struct {
	Event *DnsTopLevelNameEvOpen2Reg // Event containing the contract specifics and raw log

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
func (it *DnsTopLevelNameEvOpen2RegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsTopLevelNameEvOpen2Reg)
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
		it.Event = new(DnsTopLevelNameEvOpen2Reg)
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
func (it *DnsTopLevelNameEvOpen2RegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsTopLevelNameEvOpen2RegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsTopLevelNameEvOpen2Reg represents a EvOpen2Reg event raised by the DnsTopLevelName contract.
type DnsTopLevelNameEvOpen2Reg struct {
	NameHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvOpen2Reg is a free log retrieval operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) FilterEvOpen2Reg(opts *bind.FilterOpts) (*DnsTopLevelNameEvOpen2RegIterator, error) {

	logs, sub, err := _DnsTopLevelName.contract.FilterLogs(opts, "EvOpen2Reg")
	if err != nil {
		return nil, err
	}
	return &DnsTopLevelNameEvOpen2RegIterator{contract: _DnsTopLevelName.contract, event: "EvOpen2Reg", logs: logs, sub: sub}, nil
}

// WatchEvOpen2Reg is a free log subscription operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) WatchEvOpen2Reg(opts *bind.WatchOpts, sink chan<- *DnsTopLevelNameEvOpen2Reg) (event.Subscription, error) {

	logs, sub, err := _DnsTopLevelName.contract.WatchLogs(opts, "EvOpen2Reg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsTopLevelNameEvOpen2Reg)
				if err := _DnsTopLevelName.contract.UnpackLog(event, "EvOpen2Reg", log); err != nil {
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

// ParseEvOpen2Reg is a log parse operation binding the contract event 0x031f11cb7f71bb8ba334405345f935daaf55532f282a4b1d2e15da056d87a538.
//
// Solidity: event EvOpen2Reg(bytes32 nameHash)
func (_DnsTopLevelName *DnsTopLevelNameFilterer) ParseEvOpen2Reg(log types.Log) (*DnsTopLevelNameEvOpen2Reg, error) {
	event := new(DnsTopLevelNameEvOpen2Reg)
	if err := _DnsTopLevelName.contract.UnpackLog(event, "EvOpen2Reg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620027733803806200277383398181016040528101906200003791906200019f565b81600090805190602001906200004f92919062000071565b5080600190805190602001906200006892919062000071565b505050620003a8565b8280546200007f90620002b9565b90600052602060002090601f016020900481019282620000a35760008555620000ef565b82601f10620000be57805160ff1916838001178555620000ef565b82800160010185558215620000ef579182015b82811115620000ee578251825591602001919060010190620000d1565b5b509050620000fe919062000102565b5090565b5b808211156200011d57600081600090555060010162000103565b5090565b60006200013862000132846200024d565b62000224565b90508281526020810184848401111562000157576200015662000388565b5b6200016484828562000283565b509392505050565b600082601f83011262000184576200018362000383565b5b81516200019684826020860162000121565b91505092915050565b60008060408385031215620001b957620001b862000392565b5b600083015167ffffffffffffffff811115620001da57620001d96200038d565b5b620001e8858286016200016c565b925050602083015167ffffffffffffffff8111156200020c576200020b6200038d565b5b6200021a858286016200016c565b9150509250929050565b60006200023062000243565b90506200023e8282620002ef565b919050565b6000604051905090565b600067ffffffffffffffff8211156200026b576200026a62000354565b5b620002768262000397565b9050602081019050919050565b60005b83811015620002a357808201518184015260208101905062000286565b83811115620002b3576000848401525b50505050565b60006002820490506001821680620002d257607f821691505b60208210811415620002e957620002e862000325565b5b50919050565b620002fa8262000397565b810181811067ffffffffffffffff821117156200031c576200031b62000354565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6123bb80620003b86000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb4651461025b578063b88d4fde14610277578063c87b56dd14610293578063e985e9c5146102c3576100ea565b80636352211e146101dd57806370a082311461020d57806395d89b411461023d576100ea565b8063095ea7b3116100c8578063095ea7b31461016d57806323b872dd1461018957806342842e0e146101a55780635c707f07146101c1576100ea565b806301ffc9a7146100ef57806306fdde031461011f578063081812fc1461013d575b600080fd5b61010960048036038101906101049190611779565b6102f3565b6040516101169190611b0e565b60405180910390f35b6101276103d5565b6040516101349190611b29565b60405180910390f35b6101576004803603810190610152919061184b565b610467565b6040516101649190611aa7565b60405180910390f35b61018760048036038101906101829190611739565b6104ad565b005b6101a3600480360381019061019e9190611623565b6105c5565b005b6101bf60048036038101906101ba9190611623565b610625565b005b6101db60048036038101906101d691906117d3565b610645565b005b6101f760048036038101906101f2919061184b565b610677565b6040516102049190611aa7565b60405180910390f35b610227600480360381019061022291906115b6565b610729565b6040516102349190611c6b565b60405180910390f35b6102456107e1565b6040516102529190611b29565b60405180910390f35b610275600480360381019061027091906116f9565b610873565b005b610291600480360381019061028c9190611676565b610889565b005b6102ad60048036038101906102a8919061184b565b6108eb565b6040516102ba9190611b29565b60405180910390f35b6102dd60048036038101906102d891906115e3565b610953565b6040516102ea9190611b0e565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103be57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103ce57506103cd826109e7565b5b9050919050565b6060600080546103e490611ec1565b80601f016020809104026020016040519081016040528092919081815260200182805461041090611ec1565b801561045d5780601f106104325761010080835404028352916020019161045d565b820191906000526020600020905b81548152906001019060200180831161044057829003601f168201915b5050505050905090565b600061047282610a51565b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104b882610677565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610529576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052090611c2b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610548610a9c565b73ffffffffffffffffffffffffffffffffffffffff161480610577575061057681610571610a9c565b610953565b5b6105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad90611beb565b60405180910390fd5b6105c08383610aa4565b505050565b6105d66105d0610a9c565b82610b5d565b610615576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060c90611c4b565b60405180910390fd5b610620838383610bf2565b505050565b61064083838360405180602001604052806000815250610889565b505050565b816000908051906020019061065b9291906113ca565b5080600190805190602001906106729291906113ca565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610720576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071790611c0b565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561079a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079190611bcb565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180546107f090611ec1565b80601f016020809104026020016040519081016040528092919081815260200182805461081c90611ec1565b80156108695780601f1061083e57610100808354040283529160200191610869565b820191906000526020600020905b81548152906001019060200180831161084c57829003601f168201915b5050505050905090565b61088561087e610a9c565b8383610e59565b5050565b61089a610894610a9c565b83610b5d565b6108d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d090611c4b565b60405180910390fd5b6108e584848484610fc6565b50505050565b60606108f682610a51565b6000610900611022565b90506000815111610920576040518060200160405280600081525061094b565b8061092a84611039565b60405160200161093b929190611a83565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610a5a8161119a565b610a99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9090611c0b565b60405180910390fd5b50565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610b1783610677565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610b6983610677565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610bab5750610baa8185610953565b5b80610be957508373ffffffffffffffffffffffffffffffffffffffff16610bd184610467565b73ffffffffffffffffffffffffffffffffffffffff16145b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610c1282610677565b73ffffffffffffffffffffffffffffffffffffffff1614610c68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5f90611b6b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610cd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ccf90611b8b565b60405180910390fd5b610ce3838383611206565b610cee600082610aa4565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d3e9190611dd7565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d959190611d50565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4610e5483838361120b565b505050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ec8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebf90611bab565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610fb99190611b0e565b60405180910390a3505050565b610fd1848484610bf2565b610fdd84848484611210565b61101c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101390611b4b565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b60606000821415611081576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611195565b600082905060005b600082146110b357808061109c90611f24565b915050600a826110ac9190611da6565b9150611089565b60008167ffffffffffffffff8111156110cf576110ce61205a565b5b6040519080825280601f01601f1916602001820160405280156111015781602001600182028036833780820191505090505b5090505b6000851461118e5760018261111a9190611dd7565b9150600a856111299190611f6d565b60306111359190611d50565b60f81b81838151811061114b5761114a61202b565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111879190611da6565b9450611105565b8093505050505b919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b505050565b505050565b60006112318473ffffffffffffffffffffffffffffffffffffffff166113a7565b1561139a578373ffffffffffffffffffffffffffffffffffffffff1663150b7a0261125a610a9c565b8786866040518563ffffffff1660e01b815260040161127c9493929190611ac2565b602060405180830381600087803b15801561129657600080fd5b505af19250505080156112c757506040513d601f19601f820116820180604052508101906112c491906117a6565b60015b61134a573d80600081146112f7576040519150601f19603f3d011682016040523d82523d6000602084013e6112fc565b606091505b50600081511415611342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133990611b4b565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161491505061139f565b600190505b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b8280546113d690611ec1565b90600052602060002090601f0160209004810192826113f8576000855561143f565b82601f1061141157805160ff191683800117855561143f565b8280016001018555821561143f579182015b8281111561143e578251825591602001919060010190611423565b5b50905061144c9190611450565b5090565b5b80821115611469576000816000905550600101611451565b5090565b600061148061147b84611cab565b611c86565b90508281526020810184848401111561149c5761149b61208e565b5b6114a7848285611e7f565b509392505050565b60006114c26114bd84611cdc565b611c86565b9050828152602081018484840111156114de576114dd61208e565b5b6114e9848285611e7f565b509392505050565b60008135905061150081612329565b92915050565b60008135905061151581612340565b92915050565b60008135905061152a81612357565b92915050565b60008151905061153f81612357565b92915050565b600082601f83011261155a57611559612089565b5b813561156a84826020860161146d565b91505092915050565b600082601f83011261158857611587612089565b5b81356115988482602086016114af565b91505092915050565b6000813590506115b08161236e565b92915050565b6000602082840312156115cc576115cb612098565b5b60006115da848285016114f1565b91505092915050565b600080604083850312156115fa576115f9612098565b5b6000611608858286016114f1565b9250506020611619858286016114f1565b9150509250929050565b60008060006060848603121561163c5761163b612098565b5b600061164a868287016114f1565b935050602061165b868287016114f1565b925050604061166c868287016115a1565b9150509250925092565b600080600080608085870312156116905761168f612098565b5b600061169e878288016114f1565b94505060206116af878288016114f1565b93505060406116c0878288016115a1565b925050606085013567ffffffffffffffff8111156116e1576116e0612093565b5b6116ed87828801611545565b91505092959194509250565b600080604083850312156117105761170f612098565b5b600061171e858286016114f1565b925050602061172f85828601611506565b9150509250929050565b600080604083850312156117505761174f612098565b5b600061175e858286016114f1565b925050602061176f858286016115a1565b9150509250929050565b60006020828403121561178f5761178e612098565b5b600061179d8482850161151b565b91505092915050565b6000602082840312156117bc576117bb612098565b5b60006117ca84828501611530565b91505092915050565b600080604083850312156117ea576117e9612098565b5b600083013567ffffffffffffffff81111561180857611807612093565b5b61181485828601611573565b925050602083013567ffffffffffffffff81111561183557611834612093565b5b61184185828601611573565b9150509250929050565b60006020828403121561186157611860612098565b5b600061186f848285016115a1565b91505092915050565b61188181611e0b565b82525050565b61189081611e1d565b82525050565b60006118a182611d0d565b6118ab8185611d23565b93506118bb818560208601611e8e565b6118c48161209d565b840191505092915050565b60006118da82611d18565b6118e48185611d34565b93506118f4818560208601611e8e565b6118fd8161209d565b840191505092915050565b600061191382611d18565b61191d8185611d45565b935061192d818560208601611e8e565b80840191505092915050565b6000611946603283611d34565b9150611951826120ae565b604082019050919050565b6000611969602583611d34565b9150611974826120fd565b604082019050919050565b600061198c602483611d34565b91506119978261214c565b604082019050919050565b60006119af601983611d34565b91506119ba8261219b565b602082019050919050565b60006119d2602983611d34565b91506119dd826121c4565b604082019050919050565b60006119f5603e83611d34565b9150611a0082612213565b604082019050919050565b6000611a18601883611d34565b9150611a2382612262565b602082019050919050565b6000611a3b602183611d34565b9150611a468261228b565b604082019050919050565b6000611a5e602e83611d34565b9150611a69826122da565b604082019050919050565b611a7d81611e75565b82525050565b6000611a8f8285611908565b9150611a9b8284611908565b91508190509392505050565b6000602082019050611abc6000830184611878565b92915050565b6000608082019050611ad76000830187611878565b611ae46020830186611878565b611af16040830185611a74565b8181036060830152611b038184611896565b905095945050505050565b6000602082019050611b236000830184611887565b92915050565b60006020820190508181036000830152611b4381846118cf565b905092915050565b60006020820190508181036000830152611b6481611939565b9050919050565b60006020820190508181036000830152611b848161195c565b9050919050565b60006020820190508181036000830152611ba48161197f565b9050919050565b60006020820190508181036000830152611bc4816119a2565b9050919050565b60006020820190508181036000830152611be4816119c5565b9050919050565b60006020820190508181036000830152611c04816119e8565b9050919050565b60006020820190508181036000830152611c2481611a0b565b9050919050565b60006020820190508181036000830152611c4481611a2e565b9050919050565b60006020820190508181036000830152611c6481611a51565b9050919050565b6000602082019050611c806000830184611a74565b92915050565b6000611c90611ca1565b9050611c9c8282611ef3565b919050565b6000604051905090565b600067ffffffffffffffff821115611cc657611cc561205a565b5b611ccf8261209d565b9050602081019050919050565b600067ffffffffffffffff821115611cf757611cf661205a565b5b611d008261209d565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611d5b82611e75565b9150611d6683611e75565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611d9b57611d9a611f9e565b5b828201905092915050565b6000611db182611e75565b9150611dbc83611e75565b925082611dcc57611dcb611fcd565b5b828204905092915050565b6000611de282611e75565b9150611ded83611e75565b925082821015611e0057611dff611f9e565b5b828203905092915050565b6000611e1682611e55565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611eac578082015181840152602081019050611e91565b83811115611ebb576000848401525b50505050565b60006002820490506001821680611ed957607f821691505b60208210811415611eed57611eec611ffc565b5b50919050565b611efc8261209d565b810181811067ffffffffffffffff82111715611f1b57611f1a61205a565b5b80604052505050565b6000611f2f82611e75565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f6257611f61611f9e565b5b600182019050919050565b6000611f7882611e75565b9150611f8383611e75565b925082611f9357611f92611fcd565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b7f4552433732313a207472616e736665722066726f6d20696e636f72726563742060008201527f6f776e6572000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b7f4552433732313a2061646472657373207a65726f206973206e6f74206120766160008201527f6c6964206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60008201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000602082015250565b7f4552433732313a20696e76616c696420746f6b656e2049440000000000000000600082015250565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560008201527f72206e6f7220617070726f766564000000000000000000000000000000000000602082015250565b61233281611e0b565b811461233d57600080fd5b50565b61234981611e1d565b811461235457600080fd5b50565b61236081611e29565b811461236b57600080fd5b50565b61237781611e75565b811461238257600080fd5b5056fea2646970667358221220cb3fed047124550803098660715ee7cc3f12e1d05c902b254e891f554b27180964736f6c63430008060033",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// ERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721MetaData.Bin instead.
var ERC721Bin = ERC721MetaData.Bin

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := ERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_ERC721 *ERC721Transactor) SetName(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setName", name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_ERC721 *ERC721Session) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _ERC721.Contract.SetName(&_ERC721.TransactOpts, name_, symbol_)
}

// SetName is a paid mutator transaction binding the contract method 0x5c707f07.
//
// Solidity: function setName(string name_, string symbol_) returns()
func (_ERC721 *ERC721TransactorSession) SetName(name_ string, symbol_ string) (*types.Transaction, error) {
	return _ERC721.Contract.SetName(&_ERC721.TransactOpts, name_, symbol_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

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

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

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

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721OwnerMetaData contains all meta data concerning the Erc721Owner contract.
var Erc721OwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721Owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"erc721Owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610493380380610493833981810160405281019061003291906100cf565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b6000815190506100c981610146565b92915050565b600080604083850312156100e6576100e5610141565b5b60006100f4858286016100ba565b9250506020610105858286016100ba565b9150509250929050565b600061011a82610121565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b61014f8161010f565b811461015a57600080fd5b50565b6103278061016c6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806332cdee7b146100465780638da5cb5b14610064578063f2fde38b14610082575b600080fd5b61004e61009e565b60405161005b919061022e565b60405180910390f35b61006c6100c2565b604051610079919061022e565b60405180910390f35b61009c600480360381019061009791906101cf565b6100e8565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610176576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016d90610249565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000813590506101c9816102da565b92915050565b6000602082840312156101e5576101e46102ac565b5b60006101f3848285016101ba565b91505092915050565b6102058161027a565b82525050565b6000610218600b83610269565b9150610223826102b1565b602082019050919050565b600060208201905061024360008301846101fc565b92915050565b600060208201905081810360008301526102628161020b565b9050919050565b600082825260208201905092915050565b60006102858261028c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b6102e38161027a565b81146102ee57600080fd5b5056fea26469706673582212201d8a34347bb267bbd8b0ddc8cbcaf3749c10903c9920f2af36fb64339d58fbcd64736f6c63430008060033",
}

// Erc721OwnerABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721OwnerMetaData.ABI instead.
var Erc721OwnerABI = Erc721OwnerMetaData.ABI

// Erc721OwnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc721OwnerMetaData.Bin instead.
var Erc721OwnerBin = Erc721OwnerMetaData.Bin

// DeployErc721Owner deploys a new Ethereum contract, binding an instance of Erc721Owner to it.
func DeployErc721Owner(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address, erc721Owner_ common.Address) (common.Address, *types.Transaction, *Erc721Owner, error) {
	parsed, err := Erc721OwnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc721OwnerBin), backend, owner_, erc721Owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc721Owner{Erc721OwnerCaller: Erc721OwnerCaller{contract: contract}, Erc721OwnerTransactor: Erc721OwnerTransactor{contract: contract}, Erc721OwnerFilterer: Erc721OwnerFilterer{contract: contract}}, nil
}

// Erc721Owner is an auto generated Go binding around an Ethereum contract.
type Erc721Owner struct {
	Erc721OwnerCaller     // Read-only binding to the contract
	Erc721OwnerTransactor // Write-only binding to the contract
	Erc721OwnerFilterer   // Log filterer for contract events
}

// Erc721OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721OwnerSession struct {
	Contract     *Erc721Owner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721OwnerCallerSession struct {
	Contract *Erc721OwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Erc721OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721OwnerTransactorSession struct {
	Contract     *Erc721OwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Erc721OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721OwnerRaw struct {
	Contract *Erc721Owner // Generic contract binding to access the raw methods on
}

// Erc721OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721OwnerCallerRaw struct {
	Contract *Erc721OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721OwnerTransactorRaw struct {
	Contract *Erc721OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721Owner creates a new instance of Erc721Owner, bound to a specific deployed contract.
func NewErc721Owner(address common.Address, backend bind.ContractBackend) (*Erc721Owner, error) {
	contract, err := bindErc721Owner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721Owner{Erc721OwnerCaller: Erc721OwnerCaller{contract: contract}, Erc721OwnerTransactor: Erc721OwnerTransactor{contract: contract}, Erc721OwnerFilterer: Erc721OwnerFilterer{contract: contract}}, nil
}

// NewErc721OwnerCaller creates a new read-only instance of Erc721Owner, bound to a specific deployed contract.
func NewErc721OwnerCaller(address common.Address, caller bind.ContractCaller) (*Erc721OwnerCaller, error) {
	contract, err := bindErc721Owner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721OwnerCaller{contract: contract}, nil
}

// NewErc721OwnerTransactor creates a new write-only instance of Erc721Owner, bound to a specific deployed contract.
func NewErc721OwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721OwnerTransactor, error) {
	contract, err := bindErc721Owner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721OwnerTransactor{contract: contract}, nil
}

// NewErc721OwnerFilterer creates a new log filterer instance of Erc721Owner, bound to a specific deployed contract.
func NewErc721OwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721OwnerFilterer, error) {
	contract, err := bindErc721Owner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721OwnerFilterer{contract: contract}, nil
}

// bindErc721Owner binds a generic wrapper to an already deployed contract.
func bindErc721Owner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Owner *Erc721OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Owner.Contract.Erc721OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Owner *Erc721OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Owner.Contract.Erc721OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Owner *Erc721OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Owner.Contract.Erc721OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Owner *Erc721OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Owner *Erc721OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Owner *Erc721OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Owner.Contract.contract.Transact(opts, method, params...)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_Erc721Owner *Erc721OwnerCaller) Erc721Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721Owner.contract.Call(opts, &out, "erc721Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_Erc721Owner *Erc721OwnerSession) Erc721Owner() (common.Address, error) {
	return _Erc721Owner.Contract.Erc721Owner(&_Erc721Owner.CallOpts)
}

// Erc721Owner is a free data retrieval call binding the contract method 0x32cdee7b.
//
// Solidity: function erc721Owner() view returns(address)
func (_Erc721Owner *Erc721OwnerCallerSession) Erc721Owner() (common.Address, error) {
	return _Erc721Owner.Contract.Erc721Owner(&_Erc721Owner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Owner *Erc721OwnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721Owner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Owner *Erc721OwnerSession) Owner() (common.Address, error) {
	return _Erc721Owner.Contract.Owner(&_Erc721Owner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Owner *Erc721OwnerCallerSession) Owner() (common.Address, error) {
	return _Erc721Owner.Contract.Owner(&_Erc721Owner.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721Owner *Erc721OwnerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc721Owner.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721Owner *Erc721OwnerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721Owner.Contract.TransferOwnership(&_Erc721Owner.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721Owner *Erc721OwnerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721Owner.Contract.TransferOwnership(&_Erc721Owner.TransactOpts, newOwner)
}

// IDnsAccountantMetaData contains all meta data concerning the IDnsAccountant contract.
var IDnsAccountantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDnsAccountantABI is the input ABI used to generate the binding from.
// Deprecated: Use IDnsAccountantMetaData.ABI instead.
var IDnsAccountantABI = IDnsAccountantMetaData.ABI

// IDnsAccountant is an auto generated Go binding around an Ethereum contract.
type IDnsAccountant struct {
	IDnsAccountantCaller     // Read-only binding to the contract
	IDnsAccountantTransactor // Write-only binding to the contract
	IDnsAccountantFilterer   // Log filterer for contract events
}

// IDnsAccountantCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDnsAccountantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsAccountantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDnsAccountantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsAccountantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDnsAccountantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsAccountantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDnsAccountantSession struct {
	Contract     *IDnsAccountant   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDnsAccountantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDnsAccountantCallerSession struct {
	Contract *IDnsAccountantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IDnsAccountantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDnsAccountantTransactorSession struct {
	Contract     *IDnsAccountantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IDnsAccountantRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDnsAccountantRaw struct {
	Contract *IDnsAccountant // Generic contract binding to access the raw methods on
}

// IDnsAccountantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDnsAccountantCallerRaw struct {
	Contract *IDnsAccountantCaller // Generic read-only contract binding to access the raw methods on
}

// IDnsAccountantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDnsAccountantTransactorRaw struct {
	Contract *IDnsAccountantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDnsAccountant creates a new instance of IDnsAccountant, bound to a specific deployed contract.
func NewIDnsAccountant(address common.Address, backend bind.ContractBackend) (*IDnsAccountant, error) {
	contract, err := bindIDnsAccountant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDnsAccountant{IDnsAccountantCaller: IDnsAccountantCaller{contract: contract}, IDnsAccountantTransactor: IDnsAccountantTransactor{contract: contract}, IDnsAccountantFilterer: IDnsAccountantFilterer{contract: contract}}, nil
}

// NewIDnsAccountantCaller creates a new read-only instance of IDnsAccountant, bound to a specific deployed contract.
func NewIDnsAccountantCaller(address common.Address, caller bind.ContractCaller) (*IDnsAccountantCaller, error) {
	contract, err := bindIDnsAccountant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsAccountantCaller{contract: contract}, nil
}

// NewIDnsAccountantTransactor creates a new write-only instance of IDnsAccountant, bound to a specific deployed contract.
func NewIDnsAccountantTransactor(address common.Address, transactor bind.ContractTransactor) (*IDnsAccountantTransactor, error) {
	contract, err := bindIDnsAccountant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsAccountantTransactor{contract: contract}, nil
}

// NewIDnsAccountantFilterer creates a new log filterer instance of IDnsAccountant, bound to a specific deployed contract.
func NewIDnsAccountantFilterer(address common.Address, filterer bind.ContractFilterer) (*IDnsAccountantFilterer, error) {
	contract, err := bindIDnsAccountant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDnsAccountantFilterer{contract: contract}, nil
}

// bindIDnsAccountant binds a generic wrapper to an already deployed contract.
func bindIDnsAccountant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDnsAccountantABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsAccountant *IDnsAccountantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsAccountant.Contract.IDnsAccountantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsAccountant *IDnsAccountantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.IDnsAccountantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsAccountant *IDnsAccountantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.IDnsAccountantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsAccountant *IDnsAccountantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsAccountant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsAccountant *IDnsAccountantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsAccountant *IDnsAccountantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_IDnsAccountant *IDnsAccountantTransactor) Deposit(opts *bind.TransactOpts, contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IDnsAccountant.contract.Transact(opts, "deposit", contractAddr_, erc20Addr_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_IDnsAccountant *IDnsAccountantSession) Deposit(contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.Deposit(&_IDnsAccountant.TransactOpts, contractAddr_, erc20Addr_, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address contractAddr_, address erc20Addr_, uint256 amount_) returns()
func (_IDnsAccountant *IDnsAccountantTransactorSession) Deposit(contractAddr_ common.Address, erc20Addr_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IDnsAccountant.Contract.Deposit(&_IDnsAccountant.TransactOpts, contractAddr_, erc20Addr_, amount_)
}

// IDnsErc721OwnerMetaData contains all meta data concerning the IDnsErc721Owner contract.
var IDnsErc721OwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDnsErc721OwnerABI is the input ABI used to generate the binding from.
// Deprecated: Use IDnsErc721OwnerMetaData.ABI instead.
var IDnsErc721OwnerABI = IDnsErc721OwnerMetaData.ABI

// IDnsErc721Owner is an auto generated Go binding around an Ethereum contract.
type IDnsErc721Owner struct {
	IDnsErc721OwnerCaller     // Read-only binding to the contract
	IDnsErc721OwnerTransactor // Write-only binding to the contract
	IDnsErc721OwnerFilterer   // Log filterer for contract events
}

// IDnsErc721OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDnsErc721OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsErc721OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDnsErc721OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsErc721OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDnsErc721OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsErc721OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDnsErc721OwnerSession struct {
	Contract     *IDnsErc721Owner  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDnsErc721OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDnsErc721OwnerCallerSession struct {
	Contract *IDnsErc721OwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IDnsErc721OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDnsErc721OwnerTransactorSession struct {
	Contract     *IDnsErc721OwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IDnsErc721OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDnsErc721OwnerRaw struct {
	Contract *IDnsErc721Owner // Generic contract binding to access the raw methods on
}

// IDnsErc721OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDnsErc721OwnerCallerRaw struct {
	Contract *IDnsErc721OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// IDnsErc721OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDnsErc721OwnerTransactorRaw struct {
	Contract *IDnsErc721OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDnsErc721Owner creates a new instance of IDnsErc721Owner, bound to a specific deployed contract.
func NewIDnsErc721Owner(address common.Address, backend bind.ContractBackend) (*IDnsErc721Owner, error) {
	contract, err := bindIDnsErc721Owner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDnsErc721Owner{IDnsErc721OwnerCaller: IDnsErc721OwnerCaller{contract: contract}, IDnsErc721OwnerTransactor: IDnsErc721OwnerTransactor{contract: contract}, IDnsErc721OwnerFilterer: IDnsErc721OwnerFilterer{contract: contract}}, nil
}

// NewIDnsErc721OwnerCaller creates a new read-only instance of IDnsErc721Owner, bound to a specific deployed contract.
func NewIDnsErc721OwnerCaller(address common.Address, caller bind.ContractCaller) (*IDnsErc721OwnerCaller, error) {
	contract, err := bindIDnsErc721Owner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsErc721OwnerCaller{contract: contract}, nil
}

// NewIDnsErc721OwnerTransactor creates a new write-only instance of IDnsErc721Owner, bound to a specific deployed contract.
func NewIDnsErc721OwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IDnsErc721OwnerTransactor, error) {
	contract, err := bindIDnsErc721Owner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsErc721OwnerTransactor{contract: contract}, nil
}

// NewIDnsErc721OwnerFilterer creates a new log filterer instance of IDnsErc721Owner, bound to a specific deployed contract.
func NewIDnsErc721OwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IDnsErc721OwnerFilterer, error) {
	contract, err := bindIDnsErc721Owner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDnsErc721OwnerFilterer{contract: contract}, nil
}

// bindIDnsErc721Owner binds a generic wrapper to an already deployed contract.
func bindIDnsErc721Owner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDnsErc721OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsErc721Owner *IDnsErc721OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsErc721Owner.Contract.IDnsErc721OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsErc721Owner *IDnsErc721OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.IDnsErc721OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsErc721Owner *IDnsErc721OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.IDnsErc721OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsErc721Owner *IDnsErc721OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsErc721Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsErc721Owner *IDnsErc721OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsErc721Owner *IDnsErc721OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.contract.Transact(opts, method, params...)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDnsErc721Owner *IDnsErc721OwnerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IDnsErc721Owner.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDnsErc721Owner *IDnsErc721OwnerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.TransferOwnership(&_IDnsErc721Owner.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDnsErc721Owner *IDnsErc721OwnerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IDnsErc721Owner.Contract.TransferOwnership(&_IDnsErc721Owner.TransactOpts, newOwner)
}

// IDnsNameErc721MetaData contains all meta data concerning the IDnsNameErc721 contract.
var IDnsNameErc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"}],\"name\":\"DnsExtendExpire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DnsTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"expireTime_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"idOwner_\",\"type\":\"address\"}],\"name\":\"MintId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SigUserAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDnsNameErc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use IDnsNameErc721MetaData.ABI instead.
var IDnsNameErc721ABI = IDnsNameErc721MetaData.ABI

// IDnsNameErc721 is an auto generated Go binding around an Ethereum contract.
type IDnsNameErc721 struct {
	IDnsNameErc721Caller     // Read-only binding to the contract
	IDnsNameErc721Transactor // Write-only binding to the contract
	IDnsNameErc721Filterer   // Log filterer for contract events
}

// IDnsNameErc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IDnsNameErc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsNameErc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IDnsNameErc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsNameErc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDnsNameErc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsNameErc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDnsNameErc721Session struct {
	Contract     *IDnsNameErc721   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDnsNameErc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDnsNameErc721CallerSession struct {
	Contract *IDnsNameErc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IDnsNameErc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDnsNameErc721TransactorSession struct {
	Contract     *IDnsNameErc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IDnsNameErc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IDnsNameErc721Raw struct {
	Contract *IDnsNameErc721 // Generic contract binding to access the raw methods on
}

// IDnsNameErc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDnsNameErc721CallerRaw struct {
	Contract *IDnsNameErc721Caller // Generic read-only contract binding to access the raw methods on
}

// IDnsNameErc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDnsNameErc721TransactorRaw struct {
	Contract *IDnsNameErc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIDnsNameErc721 creates a new instance of IDnsNameErc721, bound to a specific deployed contract.
func NewIDnsNameErc721(address common.Address, backend bind.ContractBackend) (*IDnsNameErc721, error) {
	contract, err := bindIDnsNameErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDnsNameErc721{IDnsNameErc721Caller: IDnsNameErc721Caller{contract: contract}, IDnsNameErc721Transactor: IDnsNameErc721Transactor{contract: contract}, IDnsNameErc721Filterer: IDnsNameErc721Filterer{contract: contract}}, nil
}

// NewIDnsNameErc721Caller creates a new read-only instance of IDnsNameErc721, bound to a specific deployed contract.
func NewIDnsNameErc721Caller(address common.Address, caller bind.ContractCaller) (*IDnsNameErc721Caller, error) {
	contract, err := bindIDnsNameErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsNameErc721Caller{contract: contract}, nil
}

// NewIDnsNameErc721Transactor creates a new write-only instance of IDnsNameErc721, bound to a specific deployed contract.
func NewIDnsNameErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*IDnsNameErc721Transactor, error) {
	contract, err := bindIDnsNameErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsNameErc721Transactor{contract: contract}, nil
}

// NewIDnsNameErc721Filterer creates a new log filterer instance of IDnsNameErc721, bound to a specific deployed contract.
func NewIDnsNameErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*IDnsNameErc721Filterer, error) {
	contract, err := bindIDnsNameErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDnsNameErc721Filterer{contract: contract}, nil
}

// bindIDnsNameErc721 binds a generic wrapper to an already deployed contract.
func bindIDnsNameErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDnsNameErc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsNameErc721 *IDnsNameErc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsNameErc721.Contract.IDnsNameErc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsNameErc721 *IDnsNameErc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.IDnsNameErc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsNameErc721 *IDnsNameErc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.IDnsNameErc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsNameErc721 *IDnsNameErc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsNameErc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsNameErc721 *IDnsNameErc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsNameErc721 *IDnsNameErc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.contract.Transact(opts, method, params...)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_IDnsNameErc721 *IDnsNameErc721Caller) SigUserAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDnsNameErc721.contract.Call(opts, &out, "SigUserAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_IDnsNameErc721 *IDnsNameErc721Session) SigUserAddr() (common.Address, error) {
	return _IDnsNameErc721.Contract.SigUserAddr(&_IDnsNameErc721.CallOpts)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x3cf8baa2.
//
// Solidity: function SigUserAddr() view returns(address)
func (_IDnsNameErc721 *IDnsNameErc721CallerSession) SigUserAddr() (common.Address, error) {
	return _IDnsNameErc721.Contract.SigUserAddr(&_IDnsNameErc721.CallOpts)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_IDnsNameErc721 *IDnsNameErc721Transactor) DnsExtendExpire(opts *bind.TransactOpts, tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _IDnsNameErc721.contract.Transact(opts, "DnsExtendExpire", tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_IDnsNameErc721 *IDnsNameErc721Session) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.DnsExtendExpire(&_IDnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsExtendExpire is a paid mutator transaction binding the contract method 0xe767f8fd.
//
// Solidity: function DnsExtendExpire(uint256 tokenId_, uint32 expireTime_) returns()
func (_IDnsNameErc721 *IDnsNameErc721TransactorSession) DnsExtendExpire(tokenId_ *big.Int, expireTime_ uint32) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.DnsExtendExpire(&_IDnsNameErc721.TransactOpts, tokenId_, expireTime_)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_IDnsNameErc721 *IDnsNameErc721Transactor) DnsTransfer(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IDnsNameErc721.contract.Transact(opts, "DnsTransfer", to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_IDnsNameErc721 *IDnsNameErc721Session) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.DnsTransfer(&_IDnsNameErc721.TransactOpts, to, tokenId)
}

// DnsTransfer is a paid mutator transaction binding the contract method 0xc8eba760.
//
// Solidity: function DnsTransfer(address to, uint256 tokenId) returns()
func (_IDnsNameErc721 *IDnsNameErc721TransactorSession) DnsTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.DnsTransfer(&_IDnsNameErc721.TransactOpts, to, tokenId)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_IDnsNameErc721 *IDnsNameErc721Transactor) MintId(opts *bind.TransactOpts, hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _IDnsNameErc721.contract.Transact(opts, "MintId", hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_IDnsNameErc721 *IDnsNameErc721Session) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.MintId(&_IDnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// MintId is a paid mutator transaction binding the contract method 0xadfd5f91.
//
// Solidity: function MintId(bytes32 hash_, uint32 expireTime_, address idOwner_) returns(uint256)
func (_IDnsNameErc721 *IDnsNameErc721TransactorSession) MintId(hash_ [32]byte, expireTime_ uint32, idOwner_ common.Address) (*types.Transaction, error) {
	return _IDnsNameErc721.Contract.MintId(&_IDnsNameErc721.TransactOpts, hash_, expireTime_, idOwner_)
}

// IDnsTopLevelNameMetaData contains all meta data concerning the IDnsTopLevelName contract.
var IDnsTopLevelNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"}],\"name\":\"getErc721Contract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDnsTopLevelNameABI is the input ABI used to generate the binding from.
// Deprecated: Use IDnsTopLevelNameMetaData.ABI instead.
var IDnsTopLevelNameABI = IDnsTopLevelNameMetaData.ABI

// IDnsTopLevelName is an auto generated Go binding around an Ethereum contract.
type IDnsTopLevelName struct {
	IDnsTopLevelNameCaller     // Read-only binding to the contract
	IDnsTopLevelNameTransactor // Write-only binding to the contract
	IDnsTopLevelNameFilterer   // Log filterer for contract events
}

// IDnsTopLevelNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDnsTopLevelNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsTopLevelNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDnsTopLevelNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsTopLevelNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDnsTopLevelNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDnsTopLevelNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDnsTopLevelNameSession struct {
	Contract     *IDnsTopLevelName // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDnsTopLevelNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDnsTopLevelNameCallerSession struct {
	Contract *IDnsTopLevelNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IDnsTopLevelNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDnsTopLevelNameTransactorSession struct {
	Contract     *IDnsTopLevelNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IDnsTopLevelNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDnsTopLevelNameRaw struct {
	Contract *IDnsTopLevelName // Generic contract binding to access the raw methods on
}

// IDnsTopLevelNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDnsTopLevelNameCallerRaw struct {
	Contract *IDnsTopLevelNameCaller // Generic read-only contract binding to access the raw methods on
}

// IDnsTopLevelNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDnsTopLevelNameTransactorRaw struct {
	Contract *IDnsTopLevelNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDnsTopLevelName creates a new instance of IDnsTopLevelName, bound to a specific deployed contract.
func NewIDnsTopLevelName(address common.Address, backend bind.ContractBackend) (*IDnsTopLevelName, error) {
	contract, err := bindIDnsTopLevelName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDnsTopLevelName{IDnsTopLevelNameCaller: IDnsTopLevelNameCaller{contract: contract}, IDnsTopLevelNameTransactor: IDnsTopLevelNameTransactor{contract: contract}, IDnsTopLevelNameFilterer: IDnsTopLevelNameFilterer{contract: contract}}, nil
}

// NewIDnsTopLevelNameCaller creates a new read-only instance of IDnsTopLevelName, bound to a specific deployed contract.
func NewIDnsTopLevelNameCaller(address common.Address, caller bind.ContractCaller) (*IDnsTopLevelNameCaller, error) {
	contract, err := bindIDnsTopLevelName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsTopLevelNameCaller{contract: contract}, nil
}

// NewIDnsTopLevelNameTransactor creates a new write-only instance of IDnsTopLevelName, bound to a specific deployed contract.
func NewIDnsTopLevelNameTransactor(address common.Address, transactor bind.ContractTransactor) (*IDnsTopLevelNameTransactor, error) {
	contract, err := bindIDnsTopLevelName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDnsTopLevelNameTransactor{contract: contract}, nil
}

// NewIDnsTopLevelNameFilterer creates a new log filterer instance of IDnsTopLevelName, bound to a specific deployed contract.
func NewIDnsTopLevelNameFilterer(address common.Address, filterer bind.ContractFilterer) (*IDnsTopLevelNameFilterer, error) {
	contract, err := bindIDnsTopLevelName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDnsTopLevelNameFilterer{contract: contract}, nil
}

// bindIDnsTopLevelName binds a generic wrapper to an already deployed contract.
func bindIDnsTopLevelName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDnsTopLevelNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsTopLevelName *IDnsTopLevelNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsTopLevelName.Contract.IDnsTopLevelNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsTopLevelName *IDnsTopLevelNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsTopLevelName.Contract.IDnsTopLevelNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsTopLevelName *IDnsTopLevelNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsTopLevelName.Contract.IDnsTopLevelNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDnsTopLevelName *IDnsTopLevelNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDnsTopLevelName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDnsTopLevelName *IDnsTopLevelNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDnsTopLevelName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDnsTopLevelName *IDnsTopLevelNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDnsTopLevelName.Contract.contract.Transact(opts, method, params...)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameCaller) GetErc721Contract(opts *bind.CallOpts, fatherHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IDnsTopLevelName.contract.Call(opts, &out, "getErc721Contract", fatherHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _IDnsTopLevelName.Contract.GetErc721Contract(&_IDnsTopLevelName.CallOpts, fatherHash)
}

// GetErc721Contract is a free data retrieval call binding the contract method 0x2c19be76.
//
// Solidity: function getErc721Contract(bytes32 fatherHash) view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameCallerSession) GetErc721Contract(fatherHash [32]byte) (common.Address, error) {
	return _IDnsTopLevelName.Contract.GetErc721Contract(&_IDnsTopLevelName.CallOpts, fatherHash)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDnsTopLevelName.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameSession) GetOperator() (common.Address, error) {
	return _IDnsTopLevelName.Contract.GetOperator(&_IDnsTopLevelName.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_IDnsTopLevelName *IDnsTopLevelNameCallerSession) GetOperator() (common.Address, error) {
	return _IDnsTopLevelName.Contract.GetOperator(&_IDnsTopLevelName.CallOpts)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var IERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetaData.ABI instead.
var IERC721ABI = IERC721MetaData.ABI

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

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

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

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

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataMetaData contains all meta data concerning the IERC721Metadata contract.
var IERC721MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetadataMetaData.ABI instead.
var IERC721MetadataABI = IERC721MetadataMetaData.ABI

// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
type IERC721Metadata struct {
	IERC721MetadataCaller     // Read-only binding to the contract
	IERC721MetadataTransactor // Write-only binding to the contract
	IERC721MetadataFilterer   // Log filterer for contract events
}

// IERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721MetadataSession struct {
	Contract     *IERC721Metadata  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721MetadataCallerSession struct {
	Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721MetadataTransactorSession struct {
	Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721MetadataRaw struct {
	Contract *IERC721Metadata // Generic contract binding to access the raw methods on
}

// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721MetadataCallerRaw struct {
	Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactorRaw struct {
	Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	contract, err := bindIERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Metadata{IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract}}, nil
}

// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	contract, err := bindIERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataCaller{contract: contract}, nil
}

// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransactor{contract: contract}, nil
}

// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
	contract, err := bindIERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataFilterer{contract: contract}, nil
}

// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalIterator struct {
	Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApproval)
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
		it.Event = new(IERC721MetadataApproval)
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
func (it *IERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {

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

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApproval)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
	event := new(IERC721MetadataApproval)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAllIterator struct {
	Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApprovalForAll)
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
		it.Event = new(IERC721MetadataApprovalForAll)
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
func (it *IERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApprovalForAll)
				if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := new(IERC721MetadataApprovalForAll)
	if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
type IERC721MetadataTransferIterator struct {
	Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataTransfer)
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
		it.Event = new(IERC721MetadataTransfer)
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
func (it *IERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {

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

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataTransfer)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
	event := new(IERC721MetadataTransfer)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ReceiverMetaData contains all meta data concerning the IERC721Receiver contract.
var IERC721ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721ReceiverMetaData.ABI instead.
var IERC721ReceiverABI = IERC721ReceiverMetaData.ABI

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// IcostMetaData contains all meta data concerning the Icost contract.
var IcostMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"priceIdIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IcostABI is the input ABI used to generate the binding from.
// Deprecated: Use IcostMetaData.ABI instead.
var IcostABI = IcostMetaData.ABI

// Icost is an auto generated Go binding around an Ethereum contract.
type Icost struct {
	IcostCaller     // Read-only binding to the contract
	IcostTransactor // Write-only binding to the contract
	IcostFilterer   // Log filterer for contract events
}

// IcostCaller is an auto generated read-only Go binding around an Ethereum contract.
type IcostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IcostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IcostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IcostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IcostSession struct {
	Contract     *Icost            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IcostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IcostCallerSession struct {
	Contract *IcostCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IcostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IcostTransactorSession struct {
	Contract     *IcostTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IcostRaw is an auto generated low-level Go binding around an Ethereum contract.
type IcostRaw struct {
	Contract *Icost // Generic contract binding to access the raw methods on
}

// IcostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IcostCallerRaw struct {
	Contract *IcostCaller // Generic read-only contract binding to access the raw methods on
}

// IcostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IcostTransactorRaw struct {
	Contract *IcostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIcost creates a new instance of Icost, bound to a specific deployed contract.
func NewIcost(address common.Address, backend bind.ContractBackend) (*Icost, error) {
	contract, err := bindIcost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Icost{IcostCaller: IcostCaller{contract: contract}, IcostTransactor: IcostTransactor{contract: contract}, IcostFilterer: IcostFilterer{contract: contract}}, nil
}

// NewIcostCaller creates a new read-only instance of Icost, bound to a specific deployed contract.
func NewIcostCaller(address common.Address, caller bind.ContractCaller) (*IcostCaller, error) {
	contract, err := bindIcost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IcostCaller{contract: contract}, nil
}

// NewIcostTransactor creates a new write-only instance of Icost, bound to a specific deployed contract.
func NewIcostTransactor(address common.Address, transactor bind.ContractTransactor) (*IcostTransactor, error) {
	contract, err := bindIcost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IcostTransactor{contract: contract}, nil
}

// NewIcostFilterer creates a new log filterer instance of Icost, bound to a specific deployed contract.
func NewIcostFilterer(address common.Address, filterer bind.ContractFilterer) (*IcostFilterer, error) {
	contract, err := bindIcost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IcostFilterer{contract: contract}, nil
}

// bindIcost binds a generic wrapper to an already deployed contract.
func bindIcost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IcostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Icost *IcostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Icost.Contract.IcostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Icost *IcostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icost.Contract.IcostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Icost *IcostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Icost.Contract.IcostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Icost *IcostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Icost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Icost *IcostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Icost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Icost *IcostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Icost.Contract.contract.Transact(opts, method, params...)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_Icost *IcostCaller) GetSecondLevelNamePrice(opts *bind.CallOpts, fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Icost.contract.Call(opts, &out, "getSecondLevelNamePrice", fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)

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
func (_Icost *IcostSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _Icost.Contract.GetSecondLevelNamePrice(&_Icost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_Icost *IcostCallerSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _Icost.Contract.GetSecondLevelNamePrice(&_Icost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_Icost *IcostCaller) GetTopLevelNamePrice(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	var out []interface{}
	err := _Icost.contract.Call(opts, &out, "getTopLevelNamePrice", erc20Addr_, lastPriceId_, nameLength_, year_)

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
func (_Icost *IcostSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _Icost.Contract.GetTopLevelNamePrice(&_Icost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_Icost *IcostCallerSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _Icost.Contract.GetTopLevelNamePrice(&_Icost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_Icost *IcostCaller) PriceIdIsValid(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	var out []interface{}
	err := _Icost.contract.Call(opts, &out, "priceIdIsValid", erc20Addr_, lastPriceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_Icost *IcostSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _Icost.Contract.PriceIdIsValid(&_Icost.CallOpts, erc20Addr_, lastPriceId)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_Icost *IcostCallerSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _Icost.Contract.PriceIdIsValid(&_Icost.CallOpts, erc20Addr_, lastPriceId)
}

// LibAddressArrayMetaData contains all meta data concerning the LibAddressArray contract.
var LibAddressArrayMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6106f2610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c806310b142d81461005057806395953da51461008e578063ec377009146100cb575b600080fd5b81801561005c57600080fd5b5061007760048036038101906100729190610470565b610108565b6040516100859291906104e9565b60405180910390f35b81801561009a57600080fd5b506100b560048036038101906100b09190610470565b61022c565b6040516100c29190610512565b60405180910390f35b8180156100d757600080fd5b506100f260048036038101906100ed9190610470565b6102a9565b6040516100ff91906104ce565b60405180910390f35b60008060005b84805490508110156101ab578373ffffffffffffffffffffffffffffffffffffffff168582815481106101445761014361065a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610198576000809250925050610225565b80806101a3906105b3565b91505061010e565b5083839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001808580549050610220919061052d565b915091505b9250929050565b600082829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183805490506102a1919061052d565b905092915050565b600080600090505b838054905081101561043a578273ffffffffffffffffffffffffffffffffffffffff168482815481106102e7576102e661065a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610427578360018580549050610340919061052d565b815481106103515761035061065a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684828154811061038f5761038e61065a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838054806103e8576103e761062b565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556001915050610440565b8080610432906105b3565b9150506102b1565b50600090505b92915050565b6000813590506104558161068e565b92915050565b60008135905061046a816106a5565b92915050565b6000806040838503121561048757610486610689565b5b60006104958582860161045b565b92505060206104a685828601610446565b9150509250929050565b6104b98161057d565b82525050565b6104c8816105a9565b82525050565b60006020820190506104e360008301846104b0565b92915050565b60006040820190506104fe60008301856104b0565b61050b60208301846104bf565b9392505050565b600060208201905061052760008301846104bf565b92915050565b6000610538826105a9565b9150610543836105a9565b925082821015610556576105556105fc565b5b828203905092915050565b600061056c82610589565b9050919050565b6000819050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006105be826105a9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156105f1576105f06105fc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b61069781610561565b81146106a257600080fd5b50565b6106ae81610573565b81146106b957600080fd5b5056fea264697066735822122094b62f15ef48e345daafc1dbd9bff2ee96f51b05b628901d35ad53b33f326da364736f6c63430008060033",
}

// LibAddressArrayABI is the input ABI used to generate the binding from.
// Deprecated: Use LibAddressArrayMetaData.ABI instead.
var LibAddressArrayABI = LibAddressArrayMetaData.ABI

// LibAddressArrayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibAddressArrayMetaData.Bin instead.
var LibAddressArrayBin = LibAddressArrayMetaData.Bin

// DeployLibAddressArray deploys a new Ethereum contract, binding an instance of LibAddressArray to it.
func DeployLibAddressArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibAddressArray, error) {
	parsed, err := LibAddressArrayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibAddressArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibAddressArray{LibAddressArrayCaller: LibAddressArrayCaller{contract: contract}, LibAddressArrayTransactor: LibAddressArrayTransactor{contract: contract}, LibAddressArrayFilterer: LibAddressArrayFilterer{contract: contract}}, nil
}

// LibAddressArray is an auto generated Go binding around an Ethereum contract.
type LibAddressArray struct {
	LibAddressArrayCaller     // Read-only binding to the contract
	LibAddressArrayTransactor // Write-only binding to the contract
	LibAddressArrayFilterer   // Log filterer for contract events
}

// LibAddressArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibAddressArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibAddressArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibAddressArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibAddressArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibAddressArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibAddressArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibAddressArraySession struct {
	Contract     *LibAddressArray  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibAddressArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibAddressArrayCallerSession struct {
	Contract *LibAddressArrayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LibAddressArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibAddressArrayTransactorSession struct {
	Contract     *LibAddressArrayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LibAddressArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibAddressArrayRaw struct {
	Contract *LibAddressArray // Generic contract binding to access the raw methods on
}

// LibAddressArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibAddressArrayCallerRaw struct {
	Contract *LibAddressArrayCaller // Generic read-only contract binding to access the raw methods on
}

// LibAddressArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibAddressArrayTransactorRaw struct {
	Contract *LibAddressArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibAddressArray creates a new instance of LibAddressArray, bound to a specific deployed contract.
func NewLibAddressArray(address common.Address, backend bind.ContractBackend) (*LibAddressArray, error) {
	contract, err := bindLibAddressArray(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibAddressArray{LibAddressArrayCaller: LibAddressArrayCaller{contract: contract}, LibAddressArrayTransactor: LibAddressArrayTransactor{contract: contract}, LibAddressArrayFilterer: LibAddressArrayFilterer{contract: contract}}, nil
}

// NewLibAddressArrayCaller creates a new read-only instance of LibAddressArray, bound to a specific deployed contract.
func NewLibAddressArrayCaller(address common.Address, caller bind.ContractCaller) (*LibAddressArrayCaller, error) {
	contract, err := bindLibAddressArray(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibAddressArrayCaller{contract: contract}, nil
}

// NewLibAddressArrayTransactor creates a new write-only instance of LibAddressArray, bound to a specific deployed contract.
func NewLibAddressArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*LibAddressArrayTransactor, error) {
	contract, err := bindLibAddressArray(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibAddressArrayTransactor{contract: contract}, nil
}

// NewLibAddressArrayFilterer creates a new log filterer instance of LibAddressArray, bound to a specific deployed contract.
func NewLibAddressArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*LibAddressArrayFilterer, error) {
	contract, err := bindLibAddressArray(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibAddressArrayFilterer{contract: contract}, nil
}

// bindLibAddressArray binds a generic wrapper to an already deployed contract.
func bindLibAddressArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibAddressArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibAddressArray *LibAddressArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibAddressArray.Contract.LibAddressArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibAddressArray *LibAddressArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibAddressArray.Contract.LibAddressArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibAddressArray *LibAddressArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibAddressArray.Contract.LibAddressArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibAddressArray *LibAddressArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibAddressArray.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibAddressArray *LibAddressArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibAddressArray.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibAddressArray *LibAddressArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibAddressArray.Contract.contract.Transact(opts, method, params...)
}

// LibBytes32ArrayMetaData contains all meta data concerning the LibBytes32Array contract.
var LibBytes32ArrayMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x610545610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063300a068614610050578063e623fe531461008d578063ffdd3bcc146100cb575b600080fd5b81801561005c57600080fd5b50610077600480360381019061007291906102eb565b610108565b604051610084919061038d565b60405180910390f35b81801561009957600080fd5b506100b460048036038101906100af91906102eb565b61014b565b6040516100c2929190610364565b60405180910390f35b8180156100d757600080fd5b506100f260048036038101906100ed91906102eb565b6101e9565b6040516100ff9190610349565b60405180910390f35b6000828290806001815401808255809150506001900390600052602060002001600090919091909150556001838054905061014391906103a8565b905092915050565b60008060005b84805490508110156101a25783858281548110610171576101706104ad565b5b9060005260206000200154141561018f5760008092509250506101e2565b808061019a90610406565b915050610151565b508383908060018154018082558091505060019003906000526020600020016000909190919091505560018085805490506101dd91906103a8565b915091505b9250929050565b600080600090505b83805490508110156102b55782848281548110610211576102106104ad565b5b906000526020600020015414156102a257836001858054905061023491906103a8565b81548110610245576102446104ad565b5b9060005260206000200154848281548110610263576102626104ad565b5b9060005260206000200181905550838054806102825761028161047e565b5b6001900381819060005260206000200160009055905560019150506102bb565b80806102ad90610406565b9150506101f1565b50600090505b92915050565b6000813590506102d0816104e1565b92915050565b6000813590506102e5816104f8565b92915050565b60008060408385031215610302576103016104dc565b5b6000610310858286016102c1565b9250506020610321858286016102d6565b9150509250929050565b610334816103e6565b82525050565b610343816103fc565b82525050565b600060208201905061035e600083018461032b565b92915050565b6000604082019050610379600083018561032b565b610386602083018461033a565b9392505050565b60006020820190506103a2600083018461033a565b92915050565b60006103b3826103fc565b91506103be836103fc565b9250828210156103d1576103d061044f565b5b828203905092915050565b6000819050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b6000610411826103fc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156104445761044361044f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b6104ea816103dc565b81146104f557600080fd5b50565b610501816103f2565b811461050c57600080fd5b5056fea26469706673582212202b4ac7446c61a0aa4791aad3d9fcf32be38ec537919d586221c8cf4f18594e5664736f6c63430008060033",
}

// LibBytes32ArrayABI is the input ABI used to generate the binding from.
// Deprecated: Use LibBytes32ArrayMetaData.ABI instead.
var LibBytes32ArrayABI = LibBytes32ArrayMetaData.ABI

// LibBytes32ArrayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibBytes32ArrayMetaData.Bin instead.
var LibBytes32ArrayBin = LibBytes32ArrayMetaData.Bin

// DeployLibBytes32Array deploys a new Ethereum contract, binding an instance of LibBytes32Array to it.
func DeployLibBytes32Array(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibBytes32Array, error) {
	parsed, err := LibBytes32ArrayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibBytes32ArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibBytes32Array{LibBytes32ArrayCaller: LibBytes32ArrayCaller{contract: contract}, LibBytes32ArrayTransactor: LibBytes32ArrayTransactor{contract: contract}, LibBytes32ArrayFilterer: LibBytes32ArrayFilterer{contract: contract}}, nil
}

// LibBytes32Array is an auto generated Go binding around an Ethereum contract.
type LibBytes32Array struct {
	LibBytes32ArrayCaller     // Read-only binding to the contract
	LibBytes32ArrayTransactor // Write-only binding to the contract
	LibBytes32ArrayFilterer   // Log filterer for contract events
}

// LibBytes32ArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibBytes32ArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibBytes32ArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibBytes32ArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibBytes32ArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibBytes32ArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibBytes32ArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibBytes32ArraySession struct {
	Contract     *LibBytes32Array  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibBytes32ArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibBytes32ArrayCallerSession struct {
	Contract *LibBytes32ArrayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LibBytes32ArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibBytes32ArrayTransactorSession struct {
	Contract     *LibBytes32ArrayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LibBytes32ArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibBytes32ArrayRaw struct {
	Contract *LibBytes32Array // Generic contract binding to access the raw methods on
}

// LibBytes32ArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibBytes32ArrayCallerRaw struct {
	Contract *LibBytes32ArrayCaller // Generic read-only contract binding to access the raw methods on
}

// LibBytes32ArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibBytes32ArrayTransactorRaw struct {
	Contract *LibBytes32ArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibBytes32Array creates a new instance of LibBytes32Array, bound to a specific deployed contract.
func NewLibBytes32Array(address common.Address, backend bind.ContractBackend) (*LibBytes32Array, error) {
	contract, err := bindLibBytes32Array(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibBytes32Array{LibBytes32ArrayCaller: LibBytes32ArrayCaller{contract: contract}, LibBytes32ArrayTransactor: LibBytes32ArrayTransactor{contract: contract}, LibBytes32ArrayFilterer: LibBytes32ArrayFilterer{contract: contract}}, nil
}

// NewLibBytes32ArrayCaller creates a new read-only instance of LibBytes32Array, bound to a specific deployed contract.
func NewLibBytes32ArrayCaller(address common.Address, caller bind.ContractCaller) (*LibBytes32ArrayCaller, error) {
	contract, err := bindLibBytes32Array(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibBytes32ArrayCaller{contract: contract}, nil
}

// NewLibBytes32ArrayTransactor creates a new write-only instance of LibBytes32Array, bound to a specific deployed contract.
func NewLibBytes32ArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*LibBytes32ArrayTransactor, error) {
	contract, err := bindLibBytes32Array(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibBytes32ArrayTransactor{contract: contract}, nil
}

// NewLibBytes32ArrayFilterer creates a new log filterer instance of LibBytes32Array, bound to a specific deployed contract.
func NewLibBytes32ArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*LibBytes32ArrayFilterer, error) {
	contract, err := bindLibBytes32Array(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibBytes32ArrayFilterer{contract: contract}, nil
}

// bindLibBytes32Array binds a generic wrapper to an already deployed contract.
func bindLibBytes32Array(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibBytes32ArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibBytes32Array *LibBytes32ArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibBytes32Array.Contract.LibBytes32ArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibBytes32Array *LibBytes32ArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibBytes32Array.Contract.LibBytes32ArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibBytes32Array *LibBytes32ArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibBytes32Array.Contract.LibBytes32ArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibBytes32Array *LibBytes32ArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibBytes32Array.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibBytes32Array *LibBytes32ArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibBytes32Array.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibBytes32Array *LibBytes32ArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibBytes32Array.Contract.contract.Transact(opts, method, params...)
}

// LibDnsNameErc721MetaData contains all meta data concerning the LibDnsNameErc721 contract.
var LibDnsNameErc721MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x614317610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610620000375760003560e01c8063651fed36146200003c575b600080fd5b8180156200004957600080fd5b5062000068600480360381019062000062919062000185565b62000080565b604051620000779190620002a9565b60405180910390f35b6000848484846040516200009490620000cb565b620000a39493929190620002c6565b604051809103906000f080158015620000c0573d6000803e3d6000fd5b509050949350505050565b613e0480620004de83390190565b6000620000f0620000ea846200034a565b62000321565b9050828152602081018484840111156200010f576200010e62000489565b5b6200011c848285620003da565b509392505050565b6000813590506200013581620004a9565b92915050565b6000813590506200014c81620004c3565b92915050565b600082601f8301126200016a576200016962000484565b5b81356200017c848260208601620000d9565b91505092915050565b60008060008060808587031215620001a257620001a162000493565b5b600085013567ffffffffffffffff811115620001c357620001c26200048e565b5b620001d18782880162000152565b945050602085013567ffffffffffffffff811115620001f557620001f46200048e565b5b620002038782880162000152565b9350506040620002168782880162000124565b925050606062000229878288016200013b565b91505092959194509250565b62000240816200039c565b82525050565b62000251816200039c565b82525050565b6200026281620003b0565b82525050565b6000620002758262000380565b6200028181856200038b565b935062000293818560208601620003e9565b6200029e8162000498565b840191505092915050565b6000602082019050620002c0600083018462000246565b92915050565b60006080820190508181036000830152620002e2818762000268565b90508181036020830152620002f8818662000268565b905062000309604083018562000235565b62000318606083018462000257565b95945050505050565b60006200032d62000340565b90506200033b82826200041f565b919050565b6000604051905090565b600067ffffffffffffffff82111562000368576200036762000455565b5b620003738262000498565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000620003a982620003ba565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101562000409578082015181840152602081019050620003ec565b8381111562000419576000848401525b50505050565b6200042a8262000498565b810181811067ffffffffffffffff821117156200044c576200044b62000455565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b620004b4816200039c565b8114620004c057600080fd5b50565b620004ce81620003b0565b8114620004da57600080fd5b5056fe60806040523480156200001157600080fd5b5060405162003e0438038062003e04833981810160405281019062000037919062000268565b813385858160009080519060200190620000539291906200010c565b5080600190805190602001906200006c9291906200010c565b50505080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600060088190555080600981905550505050506200050e565b8280546200011a90620003eb565b90600052602060002090601f0160209004810192826200013e57600085556200018a565b82601f106200015957805160ff19168380011785556200018a565b828001600101855582156200018a579182015b82811115620001895782518255916020019190600101906200016c565b5b5090506200019991906200019d565b5090565b5b80821115620001b85760008160009055506001016200019e565b5090565b6000620001d3620001cd8462000341565b62000318565b905082815260208101848484011115620001f257620001f1620004ba565b5b620001ff848285620003b5565b509392505050565b6000815190506200021881620004da565b92915050565b6000815190506200022f81620004f4565b92915050565b600082601f8301126200024d576200024c620004b5565b5b81516200025f848260208601620001bc565b91505092915050565b60008060008060808587031215620002855762000284620004c4565b5b600085015167ffffffffffffffff811115620002a657620002a5620004bf565b5b620002b48782880162000235565b945050602085015167ffffffffffffffff811115620002d857620002d7620004bf565b5b620002e68782880162000235565b9350506040620002f98782880162000207565b92505060606200030c878288016200021e565b91505092959194509250565b60006200032462000337565b905062000332828262000421565b919050565b6000604051905090565b600067ffffffffffffffff8211156200035f576200035e62000486565b5b6200036a82620004c9565b9050602081019050919050565b6000620003848262000395565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b83811015620003d5578082015181840152602081019050620003b8565b83811115620003e5576000848401525b50505050565b600060028204905060018216806200040457607f821691505b602082108114156200041b576200041a62000457565b5b50919050565b6200042c82620004c9565b810181811067ffffffffffffffff821117156200044e576200044d62000486565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b620004e58162000377565b8114620004f157600080fd5b50565b620004ff816200038b565b81146200050b57600080fd5b50565b6138e6806200051e6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c806395d89b41116100de578063c87b56dd11610097578063d575f4ef11610071578063d575f4ef14610494578063e767f8fd146104b2578063e985e9c5146104ce578063f2fde38b146104fe5761018e565b8063c87b56dd1461042c578063c8eba7601461045c578063cb50fe76146104785761018e565b806395d89b411461036c5780639abc83201461038a578063a0bcfc7f146103a8578063a22cb465146103c4578063adfd5f91146103e0578063b88d4fde146104105761018e565b806332cdee7b1161014b5780635c707f07116101255780635c707f07146102d25780636352211e146102ee57806370a082311461031e5780638da5cb5b1461034e5761018e565b806332cdee7b1461027a5780633cf8baa21461029857806342842e0e146102b65761018e565b806301ffc9a7146101935780630636a797146101c357806306fdde03146101f4578063081812fc14610212578063095ea7b31461024257806323b872dd1461025e575b600080fd5b6101ad60048036038101906101a8919061294c565b61051a565b6040516101ba9190612e3e565b60405180910390f35b6101dd60048036038101906101d89190612a67565b61052c565b6040516101eb929190613073565b60405180910390f35b6101fc610560565b6040516102099190612e96565b60405180910390f35b61022c60048036038101906102279190612a67565b6105f2565b6040516102399190612dbc565b60405180910390f35b61025c600480360381019061025791906128b9565b610604565b005b610278600480360381019061027391906127a3565b610612565b005b61028261062c565b60405161028f9190612dd7565b60405180910390f35b6102a0610652565b6040516102ad9190612dbc565b60405180910390f35b6102d060048036038101906102cb91906127a3565b61067c565b005b6102ec60048036038101906102e791906129ef565b61069c565b005b61030860048036038101906103039190612a67565b6106ce565b6040516103159190612dbc565b60405180910390f35b610338600480360381019061033391906126dc565b610780565b6040516103459190613058565b60405180910390f35b610356610838565b6040516103639190612dd7565b60405180910390f35b61037461085e565b6040516103819190612e96565b60405180910390f35b6103926108f0565b60405161039f9190612e74565b60405180910390f35b6103c260048036038101906103bd91906129a6565b61097e565b005b6103de60048036038101906103d99190612879565b6109f2565b005b6103fa60048036038101906103f591906128f9565b610a00565b6040516104079190613058565b60405180910390f35b61042a600480360381019061042591906127f6565b610c1c565b005b61044660048036038101906104419190612a67565b610c38565b6040516104539190612e96565b60405180910390f35b610476600480360381019061047191906128b9565b610ca0565b005b610492600480360381019061048d91906126dc565b610e47565b005b61049c610ee5565b6040516104a99190613058565b60405180910390f35b6104cc60048036038101906104c79190612a94565b610eeb565b005b6104e860048036038101906104e39190612763565b6110aa565b6040516104f59190612e3e565b60405180910390f35b61051860048036038101906105139190612736565b6110be565b005b600061052582611192565b9050919050565b600c6020528060005260406000206000915090508060000160009054906101000a900463ffffffff16908060010154905082565b60606000805461056f90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461059b90613303565b80156105e85780601f106105bd576101008083540402835291602001916105e8565b820191906000526020600020905b8154815290600101906020018083116105cb57829003601f168201915b5050505050905090565b60006105fd82611274565b9050919050565b61060e82826112ba565b5050565b61061c82826113d2565b610627838383611672565b505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61069783838360405180602001604052806000815250610c1c565b505050565b81600090805190602001906106b2929190612416565b5080600190805190602001906106c9929190612416565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076e90612fd8565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e890612f78565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606001805461086d90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461089990613303565b80156108e65780601f106108bb576101008083540402835291602001916108e6565b820191906000526020600020905b8154815290600101906020018083116108c957829003601f168201915b5050505050905090565b600a80546108fd90613303565b80601f016020809104026020016040519081016040528092919081815260200182805461092990613303565b80156109765780601f1061094b57610100808354040283529160200191610976565b820191906000526020600020905b81548152906001019060200180831161095957829003601f168201915b505050505081565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d857600080fd5b80600a90805190602001906109ee92919061249c565b5050565b6109fc82826116d2565b5050565b60008060001b600954148015610a6357503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80610b4957506000801b60095414158015610b4857503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610af857600080fd5b505afa158015610b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b309190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b610b88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7f90612f58565b60405180910390fd5b60086000815480929190610b9b90613366565b9190505550610bac826008546116e8565b60405180604001604052808463ffffffff16815260200185815250600c6000600854815260200190815260200160002060008201518160000160006101000a81548163ffffffff021916908363ffffffff1602179055506020820151816001015590505060085490509392505050565b610c2683836113d2565b610c32848484846118c2565b50505050565b6060610c4382611924565b6000610c4d61196f565b90506000815111610c6d5760405180602001604052806000815250610c98565b80610c7784611a01565b604051602001610c88929190612d98565b6040516020818303038152906040525b915050919050565b6000801b600954148015610d0157503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b80610de757506000801b60095414158015610de657503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610d9657600080fd5b505afa158015610daa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dce9190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b610e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1d90612f58565b60405180910390fd5b610e39610e32826106ce565b8383611b62565b610e4382826113d2565b5050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ea157600080fd5b80600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60085481565b6000801b600954148015610f4c57503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b8061103257506000801b6009541415801561103157503373ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7f43c686040518163ffffffff1660e01b815260040160206040518083038186803b158015610fe157600080fd5b505afa158015610ff5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110199190612709565b73ffffffffffffffffffffffffffffffffffffffff16145b5b611071576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106890612f58565b60405180910390fd5b80600c600084815260200190815260200160002060000160006101000a81548163ffffffff021916908363ffffffff1602179055505050565b60006110b68383611dc9565b905092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461114e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114590612f58565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061125d57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061126d575061126c82611e5d565b5b9050919050565b600061127f82611924565b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006112c5826106ce565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132d90612ff8565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16611355611ec7565b73ffffffffffffffffffffffffffffffffffffffff16148061138457506113838161137e611ec7565b6110aa565b5b6113c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ba90612f98565b60405180910390fd5b6113cd8383611ecf565b505050565b42600c600083815260200190815260200160002060000160009054906101000a900463ffffffff1663ffffffff1611611440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161143790613018565b60405180910390fd5b6000801b600954141561166e57600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76600c6000858152602001908152602001600020600101546040518263ffffffff1660e01b81526004016114d69190612e59565b60206040518083038186803b1580156114ee57600080fd5b505afa158015611502573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115269190612709565b73ffffffffffffffffffffffffffffffffffffffff161461166d57600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76600c6000848152602001908152602001600020600101546040518263ffffffff1660e01b81526004016115b29190612e59565b60206040518083038186803b1580156115ca57600080fd5b505afa1580156115de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116029190612709565b73ffffffffffffffffffffffffffffffffffffffff1663f2fde38b836040518263ffffffff1660e01b815260040161163a9190612dd7565b600060405180830381600087803b15801561165457600080fd5b505af1158015611668573d6000803e3d6000fd5b505050505b5b5050565b61168361167d611ec7565b82611f88565b6116c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b990613038565b60405180910390fd5b6116cd838383611b62565b505050565b6116e46116dd611ec7565b838361201d565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f90612fb8565b60405180910390fd5b6117618161218a565b156117a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179890612ef8565b60405180910390fd5b6117ad600083836121f6565b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117fd9190613166565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a46118be600083836121fb565b5050565b6118d36118cd611ec7565b83611f88565b611912576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190990613038565b60405180910390fd5b61191e84848484612200565b50505050565b61192d8161218a565b61196c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196390612fd8565b60405180910390fd5b50565b6060600a805461197e90613303565b80601f01602080910402602001604051908101604052809291908181526020018280546119aa90613303565b80156119f75780601f106119cc576101008083540402835291602001916119f7565b820191906000526020600020905b8154815290600101906020018083116119da57829003601f168201915b5050505050905090565b60606000821415611a49576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611b5d565b600082905060005b60008214611a7b578080611a6490613366565b915050600a82611a7491906131bc565b9150611a51565b60008167ffffffffffffffff811115611a9757611a9661349c565b5b6040519080825280601f01601f191660200182016040528015611ac95781602001600182028036833780820191505090505b5090505b60008514611b5657600182611ae291906131ed565b9150600a85611af191906133af565b6030611afd9190613166565b60f81b818381518110611b1357611b1261346d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a85611b4f91906131bc565b9450611acd565b8093505050505b919050565b8273ffffffffffffffffffffffffffffffffffffffff16611b82826106ce565b73ffffffffffffffffffffffffffffffffffffffff1614611bd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bcf90612ed8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3f90612f18565b60405180910390fd5b611c538383836121f6565b611c5e600082611ecf565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611cae91906131ed565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d059190613166565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4611dc48383836121fb565b505050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16611f42836106ce565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080611f94836106ce565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611fd65750611fd581856110aa565b5b8061201457508373ffffffffffffffffffffffffffffffffffffffff16611ffc846105f2565b73ffffffffffffffffffffffffffffffffffffffff16145b91505092915050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561208c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161208390612f38565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161217d9190612e3e565b60405180910390a3505050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b505050565b505050565b61220b848484611b62565b6122178484848461225c565b612256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161224d90612eb8565b60405180910390fd5b50505050565b600061227d8473ffffffffffffffffffffffffffffffffffffffff166123f3565b156123e6578373ffffffffffffffffffffffffffffffffffffffff1663150b7a026122a6611ec7565b8786866040518563ffffffff1660e01b81526004016122c89493929190612df2565b602060405180830381600087803b1580156122e257600080fd5b505af192505050801561231357506040513d601f19601f820116820180604052508101906123109190612979565b60015b612396573d8060008114612343576040519150601f19603f3d011682016040523d82523d6000602084013e612348565b606091505b5060008151141561238e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161238590612eb8565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149150506123eb565b600190505b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b82805461242290613303565b90600052602060002090601f016020900481019282612444576000855561248b565b82601f1061245d57805160ff191683800117855561248b565b8280016001018555821561248b579182015b8281111561248a57825182559160200191906001019061246f565b5b5090506124989190612522565b5090565b8280546124a890613303565b90600052602060002090601f0160209004810192826124ca5760008555612511565b82601f106124e357805160ff1916838001178555612511565b82800160010185558215612511579182015b828111156125105782518255916020019190600101906124f5565b5b50905061251e9190612522565b5090565b5b8082111561253b576000816000905550600101612523565b5090565b600061255261254d846130c1565b61309c565b90508281526020810184848401111561256e5761256d6134d0565b5b6125798482856132c1565b509392505050565b600061259461258f846130f2565b61309c565b9050828152602081018484840111156125b0576125af6134d0565b5b6125bb8482856132c1565b509392505050565b6000813590506125d28161380f565b92915050565b6000815190506125e78161380f565b92915050565b6000813590506125fc81613826565b92915050565b6000813590506126118161383d565b92915050565b60008135905061262681613854565b92915050565b60008135905061263b8161386b565b92915050565b6000815190506126508161386b565b92915050565b600082601f83011261266b5761266a6134cb565b5b813561267b84826020860161253f565b91505092915050565b600082601f830112612699576126986134cb565b5b81356126a9848260208601612581565b91505092915050565b6000813590506126c181613882565b92915050565b6000813590506126d681613899565b92915050565b6000602082840312156126f2576126f16134da565b5b6000612700848285016125c3565b91505092915050565b60006020828403121561271f5761271e6134da565b5b600061272d848285016125d8565b91505092915050565b60006020828403121561274c5761274b6134da565b5b600061275a848285016125ed565b91505092915050565b6000806040838503121561277a576127796134da565b5b6000612788858286016125c3565b9250506020612799858286016125c3565b9150509250929050565b6000806000606084860312156127bc576127bb6134da565b5b60006127ca868287016125c3565b93505060206127db868287016125c3565b92505060406127ec868287016126b2565b9150509250925092565b600080600080608085870312156128105761280f6134da565b5b600061281e878288016125c3565b945050602061282f878288016125c3565b9350506040612840878288016126b2565b925050606085013567ffffffffffffffff811115612861576128606134d5565b5b61286d87828801612656565b91505092959194509250565b600080604083850312156128905761288f6134da565b5b600061289e858286016125c3565b92505060206128af85828601612602565b9150509250929050565b600080604083850312156128d0576128cf6134da565b5b60006128de858286016125c3565b92505060206128ef858286016126b2565b9150509250929050565b600080600060608486031215612912576129116134da565b5b600061292086828701612617565b9350506020612931868287016126c7565b9250506040612942868287016125c3565b9150509250925092565b600060208284031215612962576129616134da565b5b60006129708482850161262c565b91505092915050565b60006020828403121561298f5761298e6134da565b5b600061299d84828501612641565b91505092915050565b6000602082840312156129bc576129bb6134da565b5b600082013567ffffffffffffffff8111156129da576129d96134d5565b5b6129e684828501612684565b91505092915050565b60008060408385031215612a0657612a056134da565b5b600083013567ffffffffffffffff811115612a2457612a236134d5565b5b612a3085828601612684565b925050602083013567ffffffffffffffff811115612a5157612a506134d5565b5b612a5d85828601612684565b9150509250929050565b600060208284031215612a7d57612a7c6134da565b5b6000612a8b848285016126b2565b91505092915050565b60008060408385031215612aab57612aaa6134da565b5b6000612ab9858286016126b2565b9250506020612aca858286016126c7565b9150509250929050565b612add81613233565b82525050565b612aec81613221565b82525050565b612afb81613245565b82525050565b612b0a81613251565b82525050565b6000612b1b82613123565b612b258185613139565b9350612b358185602086016132d0565b612b3e816134df565b840191505092915050565b6000612b548261312e565b612b5e818561314a565b9350612b6e8185602086016132d0565b612b77816134df565b840191505092915050565b6000612b8d8261312e565b612b97818561315b565b9350612ba78185602086016132d0565b80840191505092915050565b6000612bc060328361314a565b9150612bcb826134f0565b604082019050919050565b6000612be360258361314a565b9150612bee8261353f565b604082019050919050565b6000612c06601c8361314a565b9150612c118261358e565b602082019050919050565b6000612c2960248361314a565b9150612c34826135b7565b604082019050919050565b6000612c4c60198361314a565b9150612c5782613606565b602082019050919050565b6000612c6f600b8361314a565b9150612c7a8261362f565b602082019050919050565b6000612c9260298361314a565b9150612c9d82613658565b604082019050919050565b6000612cb5603e8361314a565b9150612cc0826136a7565b604082019050919050565b6000612cd860208361314a565b9150612ce3826136f6565b602082019050919050565b6000612cfb60188361314a565b9150612d068261371f565b602082019050919050565b6000612d1e60218361314a565b9150612d2982613748565b604082019050919050565b6000612d41600c8361314a565b9150612d4c82613797565b602082019050919050565b6000612d64602e8361314a565b9150612d6f826137c0565b604082019050919050565b612d83816132a7565b82525050565b612d92816132b1565b82525050565b6000612da48285612b82565b9150612db08284612b82565b91508190509392505050565b6000602082019050612dd16000830184612ae3565b92915050565b6000602082019050612dec6000830184612ad4565b92915050565b6000608082019050612e076000830187612ae3565b612e146020830186612ae3565b612e216040830185612d7a565b8181036060830152612e338184612b10565b905095945050505050565b6000602082019050612e536000830184612af2565b92915050565b6000602082019050612e6e6000830184612b01565b92915050565b60006020820190508181036000830152612e8e8184612b10565b905092915050565b60006020820190508181036000830152612eb08184612b49565b905092915050565b60006020820190508181036000830152612ed181612bb3565b9050919050565b60006020820190508181036000830152612ef181612bd6565b9050919050565b60006020820190508181036000830152612f1181612bf9565b9050919050565b60006020820190508181036000830152612f3181612c1c565b9050919050565b60006020820190508181036000830152612f5181612c3f565b9050919050565b60006020820190508181036000830152612f7181612c62565b9050919050565b60006020820190508181036000830152612f9181612c85565b9050919050565b60006020820190508181036000830152612fb181612ca8565b9050919050565b60006020820190508181036000830152612fd181612ccb565b9050919050565b60006020820190508181036000830152612ff181612cee565b9050919050565b6000602082019050818103600083015261301181612d11565b9050919050565b6000602082019050818103600083015261303181612d34565b9050919050565b6000602082019050818103600083015261305181612d57565b9050919050565b600060208201905061306d6000830184612d7a565b92915050565b60006040820190506130886000830185612d89565b6130956020830184612b01565b9392505050565b60006130a66130b7565b90506130b28282613335565b919050565b6000604051905090565b600067ffffffffffffffff8211156130dc576130db61349c565b5b6130e5826134df565b9050602081019050919050565b600067ffffffffffffffff82111561310d5761310c61349c565b5b613116826134df565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000613171826132a7565b915061317c836132a7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156131b1576131b06133e0565b5b828201905092915050565b60006131c7826132a7565b91506131d2836132a7565b9250826131e2576131e161340f565b5b828204905092915050565b60006131f8826132a7565b9150613203836132a7565b925082821015613216576132156133e0565b5b828203905092915050565b600061322c82613287565b9050919050565b600061323e82613287565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b82818337600083830152505050565b60005b838110156132ee5780820151818401526020810190506132d3565b838111156132fd576000848401525b50505050565b6000600282049050600182168061331b57607f821691505b6020821081141561332f5761332e61343e565b5b50919050565b61333e826134df565b810181811067ffffffffffffffff8211171561335d5761335c61349c565b5b80604052505050565b6000613371826132a7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156133a4576133a36133e0565b5b600182019050919050565b60006133ba826132a7565b91506133c5836132a7565b9250826133d5576133d461340f565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b7f4552433732313a207472616e736665722066726f6d20696e636f72726563742060008201527f6f776e6572000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b7f4552433732313a2061646472657373207a65726f206973206e6f74206120766160008201527f6c6964206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60008201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c0000602082015250565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b7f4552433732313a20696e76616c696420746f6b656e2049440000000000000000600082015250565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b7f6e616d6520657870697265640000000000000000000000000000000000000000600082015250565b7f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560008201527f72206e6f7220617070726f766564000000000000000000000000000000000000602082015250565b61381881613221565b811461382357600080fd5b50565b61382f81613233565b811461383a57600080fd5b50565b61384681613245565b811461385157600080fd5b50565b61385d81613251565b811461386857600080fd5b50565b6138748161325b565b811461387f57600080fd5b50565b61388b816132a7565b811461389657600080fd5b50565b6138a2816132b1565b81146138ad57600080fd5b5056fea264697066735822122076dc0689d8c8fdbe1daa082df3c00f7140c193dd5874b241af35ef3fa1301eac64736f6c63430008060033a264697066735822122079a107fa03f5b60b44ac8a8950005bc179a4434d350e013610e61f9a37dd0a7164736f6c63430008060033",
}

// LibDnsNameErc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use LibDnsNameErc721MetaData.ABI instead.
var LibDnsNameErc721ABI = LibDnsNameErc721MetaData.ABI

// LibDnsNameErc721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibDnsNameErc721MetaData.Bin instead.
var LibDnsNameErc721Bin = LibDnsNameErc721MetaData.Bin

// DeployLibDnsNameErc721 deploys a new Ethereum contract, binding an instance of LibDnsNameErc721 to it.
func DeployLibDnsNameErc721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibDnsNameErc721, error) {
	parsed, err := LibDnsNameErc721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibDnsNameErc721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibDnsNameErc721{LibDnsNameErc721Caller: LibDnsNameErc721Caller{contract: contract}, LibDnsNameErc721Transactor: LibDnsNameErc721Transactor{contract: contract}, LibDnsNameErc721Filterer: LibDnsNameErc721Filterer{contract: contract}}, nil
}

// LibDnsNameErc721 is an auto generated Go binding around an Ethereum contract.
type LibDnsNameErc721 struct {
	LibDnsNameErc721Caller     // Read-only binding to the contract
	LibDnsNameErc721Transactor // Write-only binding to the contract
	LibDnsNameErc721Filterer   // Log filterer for contract events
}

// LibDnsNameErc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type LibDnsNameErc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsNameErc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LibDnsNameErc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsNameErc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibDnsNameErc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsNameErc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibDnsNameErc721Session struct {
	Contract     *LibDnsNameErc721 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibDnsNameErc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibDnsNameErc721CallerSession struct {
	Contract *LibDnsNameErc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LibDnsNameErc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibDnsNameErc721TransactorSession struct {
	Contract     *LibDnsNameErc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LibDnsNameErc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type LibDnsNameErc721Raw struct {
	Contract *LibDnsNameErc721 // Generic contract binding to access the raw methods on
}

// LibDnsNameErc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibDnsNameErc721CallerRaw struct {
	Contract *LibDnsNameErc721Caller // Generic read-only contract binding to access the raw methods on
}

// LibDnsNameErc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibDnsNameErc721TransactorRaw struct {
	Contract *LibDnsNameErc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLibDnsNameErc721 creates a new instance of LibDnsNameErc721, bound to a specific deployed contract.
func NewLibDnsNameErc721(address common.Address, backend bind.ContractBackend) (*LibDnsNameErc721, error) {
	contract, err := bindLibDnsNameErc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibDnsNameErc721{LibDnsNameErc721Caller: LibDnsNameErc721Caller{contract: contract}, LibDnsNameErc721Transactor: LibDnsNameErc721Transactor{contract: contract}, LibDnsNameErc721Filterer: LibDnsNameErc721Filterer{contract: contract}}, nil
}

// NewLibDnsNameErc721Caller creates a new read-only instance of LibDnsNameErc721, bound to a specific deployed contract.
func NewLibDnsNameErc721Caller(address common.Address, caller bind.ContractCaller) (*LibDnsNameErc721Caller, error) {
	contract, err := bindLibDnsNameErc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsNameErc721Caller{contract: contract}, nil
}

// NewLibDnsNameErc721Transactor creates a new write-only instance of LibDnsNameErc721, bound to a specific deployed contract.
func NewLibDnsNameErc721Transactor(address common.Address, transactor bind.ContractTransactor) (*LibDnsNameErc721Transactor, error) {
	contract, err := bindLibDnsNameErc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsNameErc721Transactor{contract: contract}, nil
}

// NewLibDnsNameErc721Filterer creates a new log filterer instance of LibDnsNameErc721, bound to a specific deployed contract.
func NewLibDnsNameErc721Filterer(address common.Address, filterer bind.ContractFilterer) (*LibDnsNameErc721Filterer, error) {
	contract, err := bindLibDnsNameErc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibDnsNameErc721Filterer{contract: contract}, nil
}

// bindLibDnsNameErc721 binds a generic wrapper to an already deployed contract.
func bindLibDnsNameErc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibDnsNameErc721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsNameErc721 *LibDnsNameErc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsNameErc721.Contract.LibDnsNameErc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsNameErc721 *LibDnsNameErc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsNameErc721.Contract.LibDnsNameErc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsNameErc721 *LibDnsNameErc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsNameErc721.Contract.LibDnsNameErc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsNameErc721 *LibDnsNameErc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsNameErc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsNameErc721 *LibDnsNameErc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsNameErc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsNameErc721 *LibDnsNameErc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsNameErc721.Contract.contract.Transact(opts, method, params...)
}

// LibDnsSignatureMetaData contains all meta data concerning the LibDnsSignature contract.
var LibDnsSignatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"SigUserAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61040d610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806354ed93b21461003a575b600080fd5b610054600480360381019061004f91906101a4565b61006a565b604051610061919061022d565b60405180910390f35b6000610076838361007e565b905092915050565b60008060008061008d856100ed565b925092509250600186848484604051600081526020016040526040516100b69493929190610248565b6020604051602081039080840390855afa1580156100d8573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000806000604184511461010057600080fd5b6020840151915060408401519050606084015160001a92509193909250565b600061013261012d846102b2565b61028d565b90508281526020810184848401111561014e5761014d6103a0565b5b61015984828561032c565b509392505050565b600081359050610170816103c0565b92915050565b600082601f83011261018b5761018a61039b565b5b813561019b84826020860161011f565b91505092915050565b600080604083850312156101bb576101ba6103aa565b5b60006101c985828601610161565b925050602083013567ffffffffffffffff8111156101ea576101e96103a5565b5b6101f685828601610176565b9150509250929050565b610209816102e3565b82525050565b610218816102f5565b82525050565b6102278161031f565b82525050565b60006020820190506102426000830184610200565b92915050565b600060808201905061025d600083018761020f565b61026a602083018661021e565b610277604083018561020f565b610284606083018461020f565b95945050505050565b60006102976102a8565b90506102a3828261033b565b919050565b6000604051905090565b600067ffffffffffffffff8211156102cd576102cc61036c565b5b6102d6826103af565b9050602081019050919050565b60006102ee826102ff565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b610344826103af565b810181811067ffffffffffffffff821117156103635761036261036c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6103c9816102f5565b81146103d457600080fd5b5056fea2646970667358221220c75386ae9d6165be0ea3032c3b045182ad8ba47fcc0d57bc7e56b59c7beb671964736f6c63430008060033",
}

// LibDnsSignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use LibDnsSignatureMetaData.ABI instead.
var LibDnsSignatureABI = LibDnsSignatureMetaData.ABI

// LibDnsSignatureBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibDnsSignatureMetaData.Bin instead.
var LibDnsSignatureBin = LibDnsSignatureMetaData.Bin

// DeployLibDnsSignature deploys a new Ethereum contract, binding an instance of LibDnsSignature to it.
func DeployLibDnsSignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibDnsSignature, error) {
	parsed, err := LibDnsSignatureMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibDnsSignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibDnsSignature{LibDnsSignatureCaller: LibDnsSignatureCaller{contract: contract}, LibDnsSignatureTransactor: LibDnsSignatureTransactor{contract: contract}, LibDnsSignatureFilterer: LibDnsSignatureFilterer{contract: contract}}, nil
}

// LibDnsSignature is an auto generated Go binding around an Ethereum contract.
type LibDnsSignature struct {
	LibDnsSignatureCaller     // Read-only binding to the contract
	LibDnsSignatureTransactor // Write-only binding to the contract
	LibDnsSignatureFilterer   // Log filterer for contract events
}

// LibDnsSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibDnsSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibDnsSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibDnsSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibDnsSignatureSession struct {
	Contract     *LibDnsSignature  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibDnsSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibDnsSignatureCallerSession struct {
	Contract *LibDnsSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LibDnsSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibDnsSignatureTransactorSession struct {
	Contract     *LibDnsSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LibDnsSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibDnsSignatureRaw struct {
	Contract *LibDnsSignature // Generic contract binding to access the raw methods on
}

// LibDnsSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibDnsSignatureCallerRaw struct {
	Contract *LibDnsSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// LibDnsSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibDnsSignatureTransactorRaw struct {
	Contract *LibDnsSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibDnsSignature creates a new instance of LibDnsSignature, bound to a specific deployed contract.
func NewLibDnsSignature(address common.Address, backend bind.ContractBackend) (*LibDnsSignature, error) {
	contract, err := bindLibDnsSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibDnsSignature{LibDnsSignatureCaller: LibDnsSignatureCaller{contract: contract}, LibDnsSignatureTransactor: LibDnsSignatureTransactor{contract: contract}, LibDnsSignatureFilterer: LibDnsSignatureFilterer{contract: contract}}, nil
}

// NewLibDnsSignatureCaller creates a new read-only instance of LibDnsSignature, bound to a specific deployed contract.
func NewLibDnsSignatureCaller(address common.Address, caller bind.ContractCaller) (*LibDnsSignatureCaller, error) {
	contract, err := bindLibDnsSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsSignatureCaller{contract: contract}, nil
}

// NewLibDnsSignatureTransactor creates a new write-only instance of LibDnsSignature, bound to a specific deployed contract.
func NewLibDnsSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*LibDnsSignatureTransactor, error) {
	contract, err := bindLibDnsSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsSignatureTransactor{contract: contract}, nil
}

// NewLibDnsSignatureFilterer creates a new log filterer instance of LibDnsSignature, bound to a specific deployed contract.
func NewLibDnsSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*LibDnsSignatureFilterer, error) {
	contract, err := bindLibDnsSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibDnsSignatureFilterer{contract: contract}, nil
}

// bindLibDnsSignature binds a generic wrapper to an already deployed contract.
func bindLibDnsSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibDnsSignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsSignature *LibDnsSignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsSignature.Contract.LibDnsSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsSignature *LibDnsSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsSignature.Contract.LibDnsSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsSignature *LibDnsSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsSignature.Contract.LibDnsSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsSignature *LibDnsSignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsSignature *LibDnsSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsSignature *LibDnsSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsSignature.Contract.contract.Transact(opts, method, params...)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x54ed93b2.
//
// Solidity: function SigUserAddr(bytes32 msgHash_, bytes sig) pure returns(address)
func (_LibDnsSignature *LibDnsSignatureCaller) SigUserAddr(opts *bind.CallOpts, msgHash_ [32]byte, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _LibDnsSignature.contract.Call(opts, &out, "SigUserAddr", msgHash_, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigUserAddr is a free data retrieval call binding the contract method 0x54ed93b2.
//
// Solidity: function SigUserAddr(bytes32 msgHash_, bytes sig) pure returns(address)
func (_LibDnsSignature *LibDnsSignatureSession) SigUserAddr(msgHash_ [32]byte, sig []byte) (common.Address, error) {
	return _LibDnsSignature.Contract.SigUserAddr(&_LibDnsSignature.CallOpts, msgHash_, sig)
}

// SigUserAddr is a free data retrieval call binding the contract method 0x54ed93b2.
//
// Solidity: function SigUserAddr(bytes32 msgHash_, bytes sig) pure returns(address)
func (_LibDnsSignature *LibDnsSignatureCallerSession) SigUserAddr(msgHash_ [32]byte, sig []byte) (common.Address, error) {
	return _LibDnsSignature.Contract.SigUserAddr(&_LibDnsSignature.CallOpts, msgHash_, sig)
}

// LibDnsToolKitMetaData contains all meta data concerning the LibDnsToolKit contract.
var LibDnsToolKitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytes2Bytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"entireNameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"fatherNameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getFatherName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getLeve2FatherNameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getLevel2FatherName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getSecondLevelName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getTopFatherNameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name_\",\"type\":\"bytes\"}],\"name\":\"getTopLevelName\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"verifyName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"verifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x611a00610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100be5760003560e01c8063a53a3d641161007b578063a53a3d64146101e3578063af207c5814610213578063c065c73114610243578063ccc1546214610273578063d7eafff3146102a3578063ec8bc798146102d3576100be565b80630b434d59146100c35780633c1f718c146100f35780636d5433e614610123578063937b91ee1461015357806399912b8a14610183578063a052a819146101b3575b600080fd5b6100dd60048036038101906100d891906111b4565b610303565b6040516100ea919061145c565b60405180910390f35b61010d600480360381019061010891906111b4565b610340565b60405161011a9190611441565b60405180910390f35b61013d60048036038101906101389190611246565b6103c8565b60405161014a919061153b565b60405180910390f35b61016d600480360381019061016891906111b4565b6103e4565b60405161017a9190611499565b60405180910390f35b61019d600480360381019061019891906111b4565b610608565b6040516101aa919061145c565b60405180910390f35b6101cd60048036038101906101c891906111b4565b61068f565b6040516101da9190611477565b60405180910390f35b6101fd60048036038101906101f891906111b4565b6107fc565b60405161020a919061145c565b60405180910390f35b61022d600480360381019061022891906111b4565b61088d565b60405161023a919061145c565b60405180910390f35b61025d600480360381019061025891906111b4565b6108ca565b60405161026a9190611441565b60405180910390f35b61028d600480360381019061028891906111b4565b610b29565b60405161029a9190611477565b60405180910390f35b6102bd60048036038101906102b891906111fd565b610da3565b6040516102ca919061145c565b60405180910390f35b6102ed60048036038101906102e891906111b4565b610dd3565b6040516102fa9190611477565b60405180910390f35b60008061030f83610dd3565b9050806040516020016103229190611413565b60405160208183030381529060405280519060200120915050919050565b600080825114806103565750604060ff16825110155b1561036457600090506103c3565b60005b82518160ff1610156103bd5761039c838260ff168151811061038c5761038b61188c565b5b602001015160f81c60f81b610f97565b6103aa5760009150506103c3565b80806103b590611833565b915050610367565b50600190505b919050565b6000818311156103da578290506103de565b8190505b92915050565b60606103ef826108ca565b61042e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610425906114fb565b60405180910390fd5b6060600080600090505b84518110156105b95760008211156104b65784818151811061045d5761045c61188c565b5b602001015160f81c60f81b83838361047591906116ec565b815181106104865761048561188c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b60008214801561052a5750602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168582815181106104fb576104fa61188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b156105a657600181865161053e91906116ec565b61054891906116ec565b67ffffffffffffffff811115610561576105606118bb565b5b6040519080825280601f01601f1916602001820160405280156105935781602001600182028036833780820191505090505b5092506001816105a3919061163c565b91505b80806105b1906117ea565b915050610438565b5060008251116105fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f5906114db565b60405180910390fd5b8192505050919050565b600061061382610340565b15610653576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064a9061151b565b60405180910390fd5b600061065e836103e4565b905080604051602001610671919061142a565b60405160208183030381529060405280519060200120915050919050565b60608060005b835181101561077257602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168482815181106106d8576106d761188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141561075f578067ffffffffffffffff811115610725576107246118bb565b5b6040519080825280601f01601f1916602001820160405280156107575781602001600182028036833780820191505090505b509150610772565b808061076a906117ea565b915050610695565b5060005b81518110156107f2578381815181106107925761079161188c565b5b602001015160f81c60f81b8282815181106107b0576107af61188c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806107ea906117ea565b915050610776565b5080915050919050565b60008060005b8351811080156108125750602081105b15610883576008816108249190611692565b60ff60f81b85838151811061083c5761083b61188c565b5b602001015160f81c60f81b167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c82179150808061087b906117ea565b915050610802565b5080915050919050565b60008061089983610b29565b9050806040516020016108ac9190611413565b60405160208183030381529060405280519060200120915050919050565b600060c061ffff168251106108e25760009050610b24565b602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168260008151811061091d5761091c61188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806109c75750602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916826001845161098791906116ec565b815181106109985761099761188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b156109d55760009050610b24565b6000805b8351811015610b1d5760008482815181106109f7576109f661188c565b5b602001015160f81c60f81b9050610a0d81610f97565b158015610a625750602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b15610a735760009350505050610b24565b602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415610ae2576000831415610ad85760009350505050610b24565b6000925050610b0a565b600183610aef919061163c565b9250604060ff168310610b085760009350505050610b24565b505b8080610b15906117ea565b9150506109d9565b5060019150505b919050565b6060610b34826108ca565b610b73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6a906114fb565b60405180910390fd5b600080600060018551610b8691906116ec565b90505b60008110610c4157602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916858281518110610bcb57610bca61188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415610c0d578280610c0990611833565b9350505b60028360ff1610610c2057809150610c41565b6000811415610c2e57610c41565b8080610c399061178f565b915050610b89565b5060028260ff161015610c89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c80906114bb565b60405180910390fd5b60006001828651610c9a91906116ec565b610ca491906116ec565b67ffffffffffffffff811115610cbd57610cbc6118bb565b5b6040519080825280601f01601f191660200182016040528015610cef5781602001600182028036833780820191505090505b5090506000600183610d01919061163c565b90505b8551811015610d9757858181518110610d2057610d1f61188c565b5b602001015160f81c60f81b8260018584610d3a91906116ec565b610d4491906116ec565b81518110610d5557610d5461188c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610d8f906117ea565b915050610d04565b50809350505050919050565b600081604051602001610db6919061142a565b604051602081830303815290604052805190602001209050919050565b606080600060018451610de691906116ec565b90505b60008110610edd57602e60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916848281518110610e2b57610e2a61188c565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415610eca576001818551610e6d91906116ec565b610e7791906116ec565b67ffffffffffffffff811115610e9057610e8f6118bb565b5b6040519080825280601f01601f191660200182016040528015610ec25781602001600182028036833780820191505090505b509150610edd565b8080610ed59061178f565b915050610de9565b5060005b8151811015610f8d57836001828651610efa91906116ec565b610f0491906116ec565b81518110610f1557610f1461188c565b5b602001015160f81c60f81b828260018551610f3091906116ec565b610f3a91906116ec565b81518110610f4b57610f4a61188c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610f85906117ea565b915050610ee1565b5080915050919050565b6000603060f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191610158015610ff55750603960f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b806110585750606160f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916101580156110575750607a60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191611155b5b806110885750602d60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806110b85750605f60f81b827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050919050565b60006110d26110cd8461157b565b611556565b9050828152602081018484840111156110ee576110ed6118ef565b5b6110f984828561174d565b509392505050565b600061111461110f846115ac565b611556565b9050828152602081018484840111156111305761112f6118ef565b5b61113b84828561174d565b509392505050565b600082601f830112611158576111576118ea565b5b81356111688482602086016110bf565b91505092915050565b600082601f830112611186576111856118ea565b5b8135611196848260208601611101565b91505092915050565b6000813590506111ae816119b3565b92915050565b6000602082840312156111ca576111c96118f9565b5b600082013567ffffffffffffffff8111156111e8576111e76118f4565b5b6111f484828501611143565b91505092915050565b600060208284031215611213576112126118f9565b5b600082013567ffffffffffffffff811115611231576112306118f4565b5b61123d84828501611171565b91505092915050565b6000806040838503121561125d5761125c6118f9565b5b600061126b8582860161119f565b925050602061127c8582860161119f565b9150509250929050565b61128f81611720565b82525050565b61129e8161172c565b82525050565b60006112af826115dd565b6112b981856115f3565b93506112c981856020860161175c565b6112d2816118fe565b840191505092915050565b60006112e8826115dd565b6112f28185611604565b935061130281856020860161175c565b80840191505092915050565b6000611319826115e8565b6113238185611620565b935061133381856020860161175c565b61133c816118fe565b840191505092915050565b6000611352826115e8565b61135c8185611631565b935061136c81856020860161175c565b80840191505092915050565b600061138560198361160f565b91506113908261190f565b602082019050919050565b60006113a8600e8361160f565b91506113b382611938565b602082019050919050565b60006113cb600e8361160f565b91506113d682611961565b602082019050919050565b60006113ee600d8361160f565b91506113f98261198a565b602082019050919050565b61140d81611736565b82525050565b600061141f82846112dd565b915081905092915050565b60006114368284611347565b915081905092915050565b60006020820190506114566000830184611286565b92915050565b60006020820190506114716000830184611295565b92915050565b6000602082019050818103600083015261149181846112a4565b905092915050565b600060208201905081810360008301526114b3818461130e565b905092915050565b600060208201905081810360008301526114d481611378565b9050919050565b600060208201905081810360008301526114f48161139b565b9050919050565b60006020820190508181036000830152611514816113be565b9050919050565b60006020820190508181036000830152611534816113e1565b9050919050565b60006020820190506115506000830184611404565b92915050565b6000611560611571565b905061156c82826117b9565b919050565b6000604051905090565b600067ffffffffffffffff821115611596576115956118bb565b5b61159f826118fe565b9050602081019050919050565b600067ffffffffffffffff8211156115c7576115c66118bb565b5b6115d0826118fe565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061164782611736565b915061165283611736565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156116875761168661185d565b5b828201905092915050565b600061169d82611736565b91506116a883611736565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156116e1576116e061185d565b5b828202905092915050565b60006116f782611736565b915061170283611736565b9250828210156117155761171461185d565b5b828203905092915050565b60008115159050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b8381101561177a57808201518184015260208101905061175f565b83811115611789576000848401525b50505050565b600061179a82611736565b915060008214156117ae576117ad61185d565b5b600182039050919050565b6117c2826118fe565b810181811067ffffffffffffffff821117156117e1576117e06118bb565b5b80604052505050565b60006117f582611736565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156118285761182761185d565b5b600182019050919050565b600061183e82611740565b915060ff8214156118525761185161185d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f6e6f742061206c6576656c203220666174686572206e616d6500000000000000600082015250565b7f6e6f20666174686572206e616d65000000000000000000000000000000000000600082015250565b7f6e616d65206e6f742076616c6964000000000000000000000000000000000000600082015250565b7f6e6f206120737562206e616d6500000000000000000000000000000000000000600082015250565b6119bc81611736565b81146119c757600080fd5b5056fea264697066735822122025c5cccb8cf583e21b1bcdee3b75298bb3da916bfd5e98950bd91688b57f0fe264736f6c63430008060033",
}

// LibDnsToolKitABI is the input ABI used to generate the binding from.
// Deprecated: Use LibDnsToolKitMetaData.ABI instead.
var LibDnsToolKitABI = LibDnsToolKitMetaData.ABI

// LibDnsToolKitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibDnsToolKitMetaData.Bin instead.
var LibDnsToolKitBin = LibDnsToolKitMetaData.Bin

// DeployLibDnsToolKit deploys a new Ethereum contract, binding an instance of LibDnsToolKit to it.
func DeployLibDnsToolKit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibDnsToolKit, error) {
	parsed, err := LibDnsToolKitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibDnsToolKitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibDnsToolKit{LibDnsToolKitCaller: LibDnsToolKitCaller{contract: contract}, LibDnsToolKitTransactor: LibDnsToolKitTransactor{contract: contract}, LibDnsToolKitFilterer: LibDnsToolKitFilterer{contract: contract}}, nil
}

// LibDnsToolKit is an auto generated Go binding around an Ethereum contract.
type LibDnsToolKit struct {
	LibDnsToolKitCaller     // Read-only binding to the contract
	LibDnsToolKitTransactor // Write-only binding to the contract
	LibDnsToolKitFilterer   // Log filterer for contract events
}

// LibDnsToolKitCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibDnsToolKitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsToolKitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibDnsToolKitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsToolKitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibDnsToolKitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDnsToolKitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibDnsToolKitSession struct {
	Contract     *LibDnsToolKit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibDnsToolKitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibDnsToolKitCallerSession struct {
	Contract *LibDnsToolKitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LibDnsToolKitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibDnsToolKitTransactorSession struct {
	Contract     *LibDnsToolKitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LibDnsToolKitRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibDnsToolKitRaw struct {
	Contract *LibDnsToolKit // Generic contract binding to access the raw methods on
}

// LibDnsToolKitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibDnsToolKitCallerRaw struct {
	Contract *LibDnsToolKitCaller // Generic read-only contract binding to access the raw methods on
}

// LibDnsToolKitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibDnsToolKitTransactorRaw struct {
	Contract *LibDnsToolKitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibDnsToolKit creates a new instance of LibDnsToolKit, bound to a specific deployed contract.
func NewLibDnsToolKit(address common.Address, backend bind.ContractBackend) (*LibDnsToolKit, error) {
	contract, err := bindLibDnsToolKit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibDnsToolKit{LibDnsToolKitCaller: LibDnsToolKitCaller{contract: contract}, LibDnsToolKitTransactor: LibDnsToolKitTransactor{contract: contract}, LibDnsToolKitFilterer: LibDnsToolKitFilterer{contract: contract}}, nil
}

// NewLibDnsToolKitCaller creates a new read-only instance of LibDnsToolKit, bound to a specific deployed contract.
func NewLibDnsToolKitCaller(address common.Address, caller bind.ContractCaller) (*LibDnsToolKitCaller, error) {
	contract, err := bindLibDnsToolKit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsToolKitCaller{contract: contract}, nil
}

// NewLibDnsToolKitTransactor creates a new write-only instance of LibDnsToolKit, bound to a specific deployed contract.
func NewLibDnsToolKitTransactor(address common.Address, transactor bind.ContractTransactor) (*LibDnsToolKitTransactor, error) {
	contract, err := bindLibDnsToolKit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibDnsToolKitTransactor{contract: contract}, nil
}

// NewLibDnsToolKitFilterer creates a new log filterer instance of LibDnsToolKit, bound to a specific deployed contract.
func NewLibDnsToolKitFilterer(address common.Address, filterer bind.ContractFilterer) (*LibDnsToolKitFilterer, error) {
	contract, err := bindLibDnsToolKit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibDnsToolKitFilterer{contract: contract}, nil
}

// bindLibDnsToolKit binds a generic wrapper to an already deployed contract.
func bindLibDnsToolKit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibDnsToolKitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsToolKit *LibDnsToolKitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsToolKit.Contract.LibDnsToolKitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsToolKit *LibDnsToolKitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsToolKit.Contract.LibDnsToolKitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsToolKit *LibDnsToolKitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsToolKit.Contract.LibDnsToolKitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDnsToolKit *LibDnsToolKitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDnsToolKit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDnsToolKit *LibDnsToolKitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDnsToolKit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDnsToolKit *LibDnsToolKitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDnsToolKit.Contract.contract.Transact(opts, method, params...)
}

// Bytes2Bytes32 is a free data retrieval call binding the contract method 0xa53a3d64.
//
// Solidity: function bytes2Bytes32(bytes b) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCaller) Bytes2Bytes32(opts *bind.CallOpts, b []byte) ([32]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "bytes2Bytes32", b)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Bytes2Bytes32 is a free data retrieval call binding the contract method 0xa53a3d64.
//
// Solidity: function bytes2Bytes32(bytes b) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitSession) Bytes2Bytes32(b []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.Bytes2Bytes32(&_LibDnsToolKit.CallOpts, b)
}

// Bytes2Bytes32 is a free data retrieval call binding the contract method 0xa53a3d64.
//
// Solidity: function bytes2Bytes32(bytes b) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) Bytes2Bytes32(b []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.Bytes2Bytes32(&_LibDnsToolKit.CallOpts, b)
}

// EntireNameHash is a free data retrieval call binding the contract method 0xd7eafff3.
//
// Solidity: function entireNameHash(string name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCaller) EntireNameHash(opts *bind.CallOpts, name_ string) ([32]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "entireNameHash", name_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EntireNameHash is a free data retrieval call binding the contract method 0xd7eafff3.
//
// Solidity: function entireNameHash(string name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitSession) EntireNameHash(name_ string) ([32]byte, error) {
	return _LibDnsToolKit.Contract.EntireNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// EntireNameHash is a free data retrieval call binding the contract method 0xd7eafff3.
//
// Solidity: function entireNameHash(string name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) EntireNameHash(name_ string) ([32]byte, error) {
	return _LibDnsToolKit.Contract.EntireNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// FatherNameHash is a free data retrieval call binding the contract method 0x99912b8a.
//
// Solidity: function fatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCaller) FatherNameHash(opts *bind.CallOpts, name_ []byte) ([32]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "fatherNameHash", name_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FatherNameHash is a free data retrieval call binding the contract method 0x99912b8a.
//
// Solidity: function fatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitSession) FatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.FatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// FatherNameHash is a free data retrieval call binding the contract method 0x99912b8a.
//
// Solidity: function fatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) FatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.FatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// GetFatherName is a free data retrieval call binding the contract method 0x937b91ee.
//
// Solidity: function getFatherName(bytes name_) pure returns(string)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetFatherName(opts *bind.CallOpts, name_ []byte) (string, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getFatherName", name_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFatherName is a free data retrieval call binding the contract method 0x937b91ee.
//
// Solidity: function getFatherName(bytes name_) pure returns(string)
func (_LibDnsToolKit *LibDnsToolKitSession) GetFatherName(name_ []byte) (string, error) {
	return _LibDnsToolKit.Contract.GetFatherName(&_LibDnsToolKit.CallOpts, name_)
}

// GetFatherName is a free data retrieval call binding the contract method 0x937b91ee.
//
// Solidity: function getFatherName(bytes name_) pure returns(string)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetFatherName(name_ []byte) (string, error) {
	return _LibDnsToolKit.Contract.GetFatherName(&_LibDnsToolKit.CallOpts, name_)
}

// GetLeve2FatherNameHash is a free data retrieval call binding the contract method 0xaf207c58.
//
// Solidity: function getLeve2FatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetLeve2FatherNameHash(opts *bind.CallOpts, name_ []byte) ([32]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getLeve2FatherNameHash", name_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeve2FatherNameHash is a free data retrieval call binding the contract method 0xaf207c58.
//
// Solidity: function getLeve2FatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitSession) GetLeve2FatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.GetLeve2FatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// GetLeve2FatherNameHash is a free data retrieval call binding the contract method 0xaf207c58.
//
// Solidity: function getLeve2FatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetLeve2FatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.GetLeve2FatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// GetLevel2FatherName is a free data retrieval call binding the contract method 0xccc15462.
//
// Solidity: function getLevel2FatherName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetLevel2FatherName(opts *bind.CallOpts, name_ []byte) ([]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getLevel2FatherName", name_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLevel2FatherName is a free data retrieval call binding the contract method 0xccc15462.
//
// Solidity: function getLevel2FatherName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitSession) GetLevel2FatherName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetLevel2FatherName(&_LibDnsToolKit.CallOpts, name_)
}

// GetLevel2FatherName is a free data retrieval call binding the contract method 0xccc15462.
//
// Solidity: function getLevel2FatherName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetLevel2FatherName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetLevel2FatherName(&_LibDnsToolKit.CallOpts, name_)
}

// GetSecondLevelName is a free data retrieval call binding the contract method 0xa052a819.
//
// Solidity: function getSecondLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetSecondLevelName(opts *bind.CallOpts, name_ []byte) ([]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getSecondLevelName", name_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSecondLevelName is a free data retrieval call binding the contract method 0xa052a819.
//
// Solidity: function getSecondLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitSession) GetSecondLevelName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetSecondLevelName(&_LibDnsToolKit.CallOpts, name_)
}

// GetSecondLevelName is a free data retrieval call binding the contract method 0xa052a819.
//
// Solidity: function getSecondLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetSecondLevelName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetSecondLevelName(&_LibDnsToolKit.CallOpts, name_)
}

// GetTopFatherNameHash is a free data retrieval call binding the contract method 0x0b434d59.
//
// Solidity: function getTopFatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetTopFatherNameHash(opts *bind.CallOpts, name_ []byte) ([32]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getTopFatherNameHash", name_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTopFatherNameHash is a free data retrieval call binding the contract method 0x0b434d59.
//
// Solidity: function getTopFatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitSession) GetTopFatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.GetTopFatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// GetTopFatherNameHash is a free data retrieval call binding the contract method 0x0b434d59.
//
// Solidity: function getTopFatherNameHash(bytes name_) pure returns(bytes32)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetTopFatherNameHash(name_ []byte) ([32]byte, error) {
	return _LibDnsToolKit.Contract.GetTopFatherNameHash(&_LibDnsToolKit.CallOpts, name_)
}

// GetTopLevelName is a free data retrieval call binding the contract method 0xec8bc798.
//
// Solidity: function getTopLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCaller) GetTopLevelName(opts *bind.CallOpts, name_ []byte) ([]byte, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "getTopLevelName", name_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTopLevelName is a free data retrieval call binding the contract method 0xec8bc798.
//
// Solidity: function getTopLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitSession) GetTopLevelName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetTopLevelName(&_LibDnsToolKit.CallOpts, name_)
}

// GetTopLevelName is a free data retrieval call binding the contract method 0xec8bc798.
//
// Solidity: function getTopLevelName(bytes name_) pure returns(bytes)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) GetTopLevelName(name_ []byte) ([]byte, error) {
	return _LibDnsToolKit.Contract.GetTopLevelName(&_LibDnsToolKit.CallOpts, name_)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256)
func (_LibDnsToolKit *LibDnsToolKitCaller) Max(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "max", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256)
func (_LibDnsToolKit *LibDnsToolKitSession) Max(x *big.Int, y *big.Int) (*big.Int, error) {
	return _LibDnsToolKit.Contract.Max(&_LibDnsToolKit.CallOpts, x, y)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) Max(x *big.Int, y *big.Int) (*big.Int, error) {
	return _LibDnsToolKit.Contract.Max(&_LibDnsToolKit.CallOpts, x, y)
}

// VerifyName is a free data retrieval call binding the contract method 0xc065c731.
//
// Solidity: function verifyName(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitCaller) VerifyName(opts *bind.CallOpts, name []byte) (bool, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "verifyName", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyName is a free data retrieval call binding the contract method 0xc065c731.
//
// Solidity: function verifyName(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitSession) VerifyName(name []byte) (bool, error) {
	return _LibDnsToolKit.Contract.VerifyName(&_LibDnsToolKit.CallOpts, name)
}

// VerifyName is a free data retrieval call binding the contract method 0xc065c731.
//
// Solidity: function verifyName(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) VerifyName(name []byte) (bool, error) {
	return _LibDnsToolKit.Contract.VerifyName(&_LibDnsToolKit.CallOpts, name)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitCaller) VerifyRoot(opts *bind.CallOpts, name []byte) (bool, error) {
	var out []interface{}
	err := _LibDnsToolKit.contract.Call(opts, &out, "verifyRoot", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitSession) VerifyRoot(name []byte) (bool, error) {
	return _LibDnsToolKit.Contract.VerifyRoot(&_LibDnsToolKit.CallOpts, name)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x3c1f718c.
//
// Solidity: function verifyRoot(bytes name) pure returns(bool)
func (_LibDnsToolKit *LibDnsToolKitCallerSession) VerifyRoot(name []byte) (bool, error) {
	return _LibDnsToolKit.Contract.VerifyRoot(&_LibDnsToolKit.CallOpts, name)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201a5677f50ee2ddddc9e6f5caefb0e303987dc83ff8ba1f0e491977e5532970e964736f6c63430008060033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// OwnedMetaData contains all meta data concerning the Owned contract.
var OwnedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610224806100606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b14610059575b600080fd5b610043610075565b6040516100509190610185565b60405180910390f35b610073600480360381019061006e9190610149565b610099565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100f157600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600081359050610143816101d7565b92915050565b60006020828403121561015f5761015e6101d2565b5b600061016d84828501610134565b91505092915050565b61017f816101a0565b82525050565b600060208201905061019a6000830184610176565b92915050565b60006101ab826101b2565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6101e0816101a0565b81146101eb57600080fd5b5056fea264697066735822122089ce7fbab0e1d575be5c8b6491d531440cd2bc4c80d84c29650909dad7ddaa5c64736f6c63430008060033",
}

// OwnedABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnedMetaData.ABI instead.
var OwnedABI = OwnedMetaData.ABI

// OwnedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnedMetaData.Bin instead.
var OwnedBin = OwnedMetaData.Bin

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := OwnedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owned.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Owned *OwnedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}
