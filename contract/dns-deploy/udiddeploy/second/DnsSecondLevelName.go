// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package second

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

// DnsSecondLevelNameMetaData contains all meta data concerning the DnsSecondLevelName contract.
var DnsSecondLevelNameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cost_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"acct_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsTop_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireSubName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"level2FatherHash\",\"type\":\"bytes32\"}],\"name\":\"EvAddSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeSecondLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"}],\"name\":\"EvChargeSecondLevelNameBySig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"level2FatherHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"topHash\",\"type\":\"bytes32\"}],\"name\":\"EvDelSubName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintSecondLevelName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"year\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EvMintSecondLevelNameBySig\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireSubName_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"level2FatherHash_\",\"type\":\"bytes32\"}],\"name\":\"AddSubName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeSecondLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isTransfer_\",\"type\":\"bool\"}],\"name\":\"ChargeSecondLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"level2FatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"topHash_\",\"type\":\"bytes32\"}],\"name\":\"DelSubName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"MintSecondLevelName\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"entireName_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"MintSecondLevelNameBySig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountantC\",\"outputs\":[{\"internalType\":\"contractIDnsAccountant\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costContractAddr\",\"outputs\":[{\"internalType\":\"contractIcost\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsTop\",\"outputs\":[{\"internalType\":\"contractIDnsTopLevelName\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintSwitch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"entireName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"expireTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cost_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"acct_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dnsTop_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mintSw_\",\"type\":\"uint8\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subNameStore\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600260146101000a81548160ff021916908360ff1602179055503480156200002d57600080fd5b5060405162004c7c38038062004c7c833981810160405281019062000053919062000135565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001e4565b6000815190506200012f81620001ca565b92915050565b600080600060608486031215620001515762000150620001c5565b5b600062000161868287016200011e565b935050602062000174868287016200011e565b925050604062000187868287016200011e565b9150509250925092565b60006200019e82620001a5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001d58162000191565b8114620001e157600080fd5b50565b614a8880620001f46000396000f3fe6080604052600436106100c25760003560e01c80639150cf8f1161007f578063c66aee2d11610059578063c66aee2d1461023f578063eacb912d1461025b578063f3c8bb0c14610286578063fddd23a8146102b1576100c2565b80639150cf8f146101ca578063b353963014610207578063bcfb550514610223576100c2565b806306475e73146100c757806309d52623146100f057806312014f011461012f5780633752770e1461015a5780636130457d14610183578063752ab03d146101ae575b600080fd5b3480156100d357600080fd5b506100ee60048036038101906100e99190613689565b6102da565b005b3480156100fc57600080fd5b506101176004803603810190610112919061343a565b6106f1565b6040516101269392919061410d565b60405180910390f35b34801561013b57600080fd5b506101446107c0565b6040516101519190614048565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c919061347a565b6107e6565b005b34801561018f57600080fd5b50610198610a1d565b6040516101a5919061407e565b60405180910390f35b6101c860048036038101906101c391906134cd565b610a41565b005b3480156101d657600080fd5b506101f160048036038101906101ec919061343a565b610cea565b6040516101fe91906140bb565b60405180910390f35b610221600480360381019061021c919061355a565b610d97565b005b61023d60048036038101906102389190613768565b6112ee565b005b610259600480360381019061025491906136e5565b611bf1565b005b34801561026757600080fd5b506102706121c7565b60405161027d9190614383565b60405180910390f35b34801561029257600080fd5b5061029b6121da565b6040516102a89190614063565b60405180910390f35b3480156102bd57600080fd5b506102d860048036038101906102d39190613379565b612200565b005b8073__$61b1050b44c222c225346b0c1857883025$__63af207c58846040518263ffffffff1660e01b81526004016103129190614004565b60206040518083038186803b15801561032a57600080fd5b505af415801561033e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610362919061340d565b146103a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610399906142bf565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__630b434d59846040518263ffffffff1660e01b81526004016103db9190614004565b60206040518083038186803b1580156103f357600080fd5b505af4158015610407573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042b919061340d565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76826040518263ffffffff1660e01b81526004016104889190613e36565b60206040518083038186803b1580156104a057600080fd5b505afa1580156104b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d8919061331f565b73ffffffffffffffffffffffffffffffffffffffff16636352211e600460008481526020019081526020016000206000858152602001908152602001600020600201546040518263ffffffff1660e01b8152600401610537919061433f565b60206040518083038186803b15801561054f57600080fd5b505afa158015610563573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610587919061331f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105eb9061431f565b60405180910390fd5b8260036000848152602001908152602001600020600073__$61b1050b44c222c225346b0c1857883025$__63d7eafff3876040518263ffffffff1660e01b81526004016106419190614099565b60206040518083038186803b15801561065957600080fd5b505af415801561066d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610691919061340d565b815260200190815260200160002090805190602001906106b2929190613005565b507f3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b883836040516106e49291906140dd565b60405180910390a1505050565b6004602052816000526040600020602052806000526040600020600091509150508060000180546107219061467f565b80601f016020809104026020016040519081016040528092919081815260200182805461074d9061467f565b801561079a5780601f1061076f5761010080835404028352916020019161079a565b820191906000526020600020905b81548152906001019060200180831161077d57829003601f168201915b5050505050908060010160009054906101000a900463ffffffff16908060020154905083565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76826040518263ffffffff1660e01b81526004016108419190613e36565b60206040518083038186803b15801561085957600080fd5b505afa15801561086d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610891919061331f565b73ffffffffffffffffffffffffffffffffffffffff16636352211e600460008481526020019081526020016000206000858152602001908152602001600020600201546040518263ffffffff1660e01b81526004016108f0919061433f565b60206040518083038186803b15801561090857600080fd5b505afa15801561091c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610940919061331f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a49061431f565b60405180910390fd5b60036000838152602001908152602001600020600084815260200190815260200160002060006109dd919061308b565b7f1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a838383604051610a1093929190613ea4565b60405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060028060149054906101000a900460ff161660ff1614610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f906141ff565b60405180910390fd5b600060046000888152602001908152602001600020600087815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff1611610b18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0f9061427f565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63a052a819600460008a815260200190815260200160002060008981526020019081526020016000206000016040518263ffffffff1660e01b8152600401610b779190614026565b60006040518083038186803b158015610b8f57600080fd5b505af4158015610ba3573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610bcc9190613640565b905060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f33b2d88a888887518c6040518663ffffffff1660e01b8152600401610c33959493929190613e51565b60606040518083038186803b158015610c4b57600080fd5b505afa158015610c5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c83919061386f565b5091509150610c94898784846123f0565b610ca089898987612c26565b7f920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe516408989898988604051610cd7959493929190613edb565b60405180910390a1505050505050505050565b6003602052816000526040600020602052806000526040600020600091509150508054610d169061467f565b80601f0160208091040260200160405190810160405280929190818152602001828054610d429061467f565b8015610d8f5780601f10610d6457610100808354040283529160200191610d8f565b820191906000526020600020905b815481529060010190602001808311610d7257829003601f168201915b505050505081565b60006008600260149054906101000a900460ff161660ff1614610def576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de6906141ff565b60405180910390fd5b6000600460008b815260200190815260200160002060008a815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff1611610e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e669061427f565b60405180910390fd5b73__$2ed95cc1c04a020bf25a1d27d6730d85d6$__6354ed93b28a898989898933604051602001610ea69796959493929190613c94565b60405160208183030381529060405280519060200120846040518363ffffffff1660e01b8152600401610eda929190613f9d565b60206040518083038186803b158015610ef257600080fd5b505af4158015610f06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2a919061331f565b73ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be768b6040518263ffffffff1660e01b8152600401610f9b9190613e36565b60206040518083038186803b158015610fb357600080fd5b505afa158015610fc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610feb919061331f565b73ffffffffffffffffffffffffffffffffffffffff16633cf8baa26040518163ffffffff1660e01b815260040160206040518083038186803b15801561103057600080fd5b505afa158015611044573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611068919061331f565b73ffffffffffffffffffffffffffffffffffffffff16146110be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b5906142df565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63a052a819600460008d815260200190815260200160002060008c81526020019081526020016000206000016040518263ffffffff1660e01b815260040161111d9190614026565b60006040518083038186803b15801561113557600080fd5b505af4158015611149573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906111729190613640565b905060008411156112935760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f33b2d88c8a8a86518e6040518663ffffffff1660e01b81526004016111e2959493929190613e51565b60606040518083038186803b1580156111fa57600080fd5b505afa15801561120e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611232919061386f565b5050905080851015611279576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611270906142ff565b60405180910390fd5b80851015611285578490505b6112918b8983886123f0565b505b61129f8a8a8a85612c26565b7f07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a78a8a8a8a8989886040516112da9796959493929190613f2e565b60405180910390a150505050505050505050565b60006004600260149054906101000a900460ff161660ff1614611346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133d9061429f565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__6399912b8a896040518263ffffffff1660e01b815260040161137f9190614004565b60206040518083038186803b15801561139757600080fd5b505af41580156113ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113cf919061340d565b9050600088805190602001209050600060046000848152602001908152602001600020600083815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff161461145d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114549061425f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76846040518263ffffffff1660e01b81526004016114d09190613e36565b60206040518083038186803b1580156114e857600080fd5b505afa1580156114fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611520919061331f565b73ffffffffffffffffffffffffffffffffffffffff161415611577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156e9061423f565b60405180910390fd5b73__$2ed95cc1c04a020bf25a1d27d6730d85d6$__6354ed93b28a8a8a8a8a8a336040516020016115ae9796959493929190613d15565b60405160208183030381529060405280519060200120856040518363ffffffff1660e01b81526004016115e2929190613f9d565b60206040518083038186803b1580156115fa57600080fd5b505af415801561160e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611632919061331f565b73ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76846040518263ffffffff1660e01b81526004016116a39190613e36565b60206040518083038186803b1580156116bb57600080fd5b505afa1580156116cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f3919061331f565b73ffffffffffffffffffffffffffffffffffffffff16633cf8baa26040518163ffffffff1660e01b815260040160206040518083038186803b15801561173857600080fd5b505afa15801561174c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611770919061331f565b73ffffffffffffffffffffffffffffffffffffffff16146117c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117bd906142df565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63a052a8198b6040518263ffffffff1660e01b81526004016117ff9190614004565b60006040518083038186803b15801561181757600080fd5b505af415801561182b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906118549190613640565b905060008511156119755760008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f33b2d8858b8b86518f6040518663ffffffff1660e01b81526004016118c4959493929190613e51565b60606040518083038186803b1580156118dc57600080fd5b505afa1580156118f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611914919061386f565b505090508086101561195b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611952906142ff565b60405180910390fd5b80861015611967578590505b611973848a83896123f0565b505b60405180606001604052808b81526020016301e133808b60ff1661199991906144c8565b426119a4919061448e565b63ffffffff168152602001600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76866040518263ffffffff1660e01b8152600401611a0a9190613e36565b60206040518083038186803b158015611a2257600080fd5b505afa158015611a36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5a919061331f565b73ffffffffffffffffffffffffffffffffffffffff1663adfd5f91856301e133808e60ff16611a8991906144c8565b42611a94919061448e565b336040518463ffffffff1660e01b8152600401611ab393929190613fcd565b602060405180830381600087803b158015611acd57600080fd5b505af1158015611ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b059190613842565b8152506004600085815260200190815260200160002060008481526020019081526020016000206000820151816000019080519060200190611b48929190613005565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff160217905550604082015181600201559050507fd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d68a8a8a8989600460008a8152602001908152602001600020600089815260200190815260200160002060020154604051611bdd96959493929190614197565b60405180910390a150505050505050505050565b60006001600260149054906101000a900460ff161660ff1614611c49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c409061429f565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__6399912b8a866040518263ffffffff1660e01b8152600401611c829190614004565b60206040518083038186803b158015611c9a57600080fd5b505af4158015611cae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd2919061340d565b9050600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76836040518263ffffffff1660e01b8152600401611d479190613e36565b60206040518083038186803b158015611d5f57600080fd5b505afa158015611d73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d97919061331f565b73ffffffffffffffffffffffffffffffffffffffff161415611dee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611de59061423f565b60405180910390fd5b600073__$61b1050b44c222c225346b0c1857883025$__63a052a819876040518263ffffffff1660e01b8152600401611e279190614004565b60006040518083038186803b158015611e3f57600080fd5b505af4158015611e53573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611e7c9190613640565b905060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f33b2d885888887518c6040518663ffffffff1660e01b8152600401611ee3959493929190613e51565b60606040518083038186803b158015611efb57600080fd5b505afa158015611f0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f33919061386f565b5091509150611f44848784846123f0565b60008880519060200120905060405180606001604052808a81526020016301e133808a60ff16611f7491906144c8565b42611f7f919061448e565b63ffffffff168152602001600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76886040518263ffffffff1660e01b8152600401611fe59190613e36565b60206040518083038186803b158015611ffd57600080fd5b505afa158015612011573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612035919061331f565b73ffffffffffffffffffffffffffffffffffffffff1663adfd5f91846301e133808d60ff1661206491906144c8565b4261206f919061448e565b336040518463ffffffff1660e01b815260040161208e93929190613fcd565b602060405180830381600087803b1580156120a857600080fd5b505af11580156120bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120e09190613842565b8152506004600087815260200190815260200160002060008381526020019081526020016000206000820151816000019080519060200190612123929190613005565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff160217905550604082015181600201559050507f500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9898989600460008a81526020019081526020016000206000868152602001908152602001600020600201546040516121b4949392919061414b565b60405180910390a1505050505050505050565b600260149054906101000a900460ff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561226857600080fd5b505afa15801561227c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a0919061334c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461230d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123049061431f565b60405180910390fd5b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260146101000a81548160ff021916908360ff16021790555050505050565b600081905081831115612401578290505b6000838261240f9190614506565b9050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614156127875781341015612489576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124809061421f565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be766000801b6040518263ffffffff1660e01b81526004016125259190613e36565b60206040518083038186803b15801561253d57600080fd5b505afa158015612551573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612575919061331f565b87876040518463ffffffff1660e01b815260040161259593929190613dd6565b600060405180830381600087803b1580156125af57600080fd5b505af11580156125c3573d6000803e3d6000fd5b50505050813411156125de5783346125db9190614506565b90505b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76896040518263ffffffff1660e01b81526004016126779190613e36565b60206040518083038186803b15801561268f57600080fd5b505afa1580156126a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c7919061331f565b87876040518463ffffffff1660e01b81526004016126e793929190613dd6565b600060405180830381600087803b15801561270157600080fd5b505af1158015612715573d6000803e3d6000fd5b50505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015612781573d6000803e3d6000fd5b50612c1e565b818573ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016127c19190613d92565b60206040518083038186803b1580156127d957600080fd5b505afa1580156127ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128119190613842565b101580156128a95750818573ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401612856929190613dad565b60206040518083038186803b15801561286e57600080fd5b505afa158015612882573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128a69190613842565b10155b6128e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128df9061421f565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be766000801b6040518263ffffffff1660e01b81526004016129849190613e36565b60206040518083038186803b15801561299c57600080fd5b505afa1580156129b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129d4919061331f565b87876040518463ffffffff1660e01b81526004016129f493929190613dd6565b600060405180830381600087803b158015612a0e57600080fd5b505af1158015612a22573d6000803e3d6000fd5b505050506000811115612b6b57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638340f549600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76896040518263ffffffff1660e01b8152600401612ac89190613e36565b60206040518083038186803b158015612ae057600080fd5b505afa158015612af4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b18919061331f565b87846040518463ffffffff1660e01b8152600401612b3893929190613dd6565b600060405180830381600087803b158015612b5257600080fd5b505af1158015612b66573d6000803e3d6000fd5b505050505b8473ffffffffffffffffffffffffffffffffffffffff166323b872dd33600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518463ffffffff1660e01b8152600401612bca93929190613dd6565b602060405180830381600087803b158015612be457600080fd5b505af1158015612bf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c1c91906133e0565b505b505050505050565b60046000858152602001908152602001600020600084815260200190815260200160002060010160009054906101000a900463ffffffff1663ffffffff16421115612e14578015612db057600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76856040518263ffffffff1660e01b8152600401612ccc9190613e36565b60206040518083038186803b158015612ce457600080fd5b505afa158015612cf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d1c919061331f565b73ffffffffffffffffffffffffffffffffffffffff1663c8eba76033600460008881526020019081526020016000206000878152602001908152602001600020600201546040518363ffffffff1660e01b8152600401612d7d929190613e0d565b600060405180830381600087803b158015612d9757600080fd5b505af1158015612dab573d6000803e3d6000fd5b505050505b6301e133808260ff16612dc391906144c8565b42612dce919061448e565b60046000868152602001908152602001600020600085815260200190815260200160002060010160006101000a81548163ffffffff021916908363ffffffff1602179055505b6301e133808260ff16612e2791906144c8565b60046000868152602001908152602001600020600085815260200190815260200160002060010160008282829054906101000a900463ffffffff16612e6c919061448e565b92506101000a81548163ffffffff021916908363ffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76856040518263ffffffff1660e01b8152600401612ee59190613e36565b60206040518083038186803b158015612efd57600080fd5b505afa158015612f11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f35919061331f565b73ffffffffffffffffffffffffffffffffffffffff1663e767f8fd6004600087815260200190815260200160002060008681526020019081526020016000206002015460046000888152602001908152602001600020600087815260200190815260200160002060010160009054906101000a900463ffffffff166040518363ffffffff1660e01b8152600401612fcd92919061435a565b600060405180830381600087803b158015612fe757600080fd5b505af1158015612ffb573d6000803e3d6000fd5b5050505050505050565b8280546130119061467f565b90600052602060002090601f016020900481019282613033576000855561307a565b82601f1061304c57805160ff191683800117855561307a565b8280016001018555821561307a579182015b8281111561307957825182559160200191906001019061305e565b5b50905061308791906130cb565b5090565b5080546130979061467f565b6000825580601f106130a957506130c8565b601f0160209004906000526020600020908101906130c791906130cb565b5b50565b5b808211156130e45760008160009055506001016130cc565b5090565b60006130fb6130f6846143c3565b61439e565b905082815260208101848484011115613117576131166147d0565b5b61312284828561463d565b509392505050565b600061313d613138846143c3565b61439e565b905082815260208101848484011115613159576131586147d0565b5b61316484828561464c565b509392505050565b600061317f61317a846143f4565b61439e565b90508281526020810184848401111561319b5761319a6147d0565b5b6131a684828561463d565b509392505050565b6000813590506131bd816149b1565b92915050565b6000815190506131d2816149b1565b92915050565b6000815190506131e7816149c8565b92915050565b6000813590506131fc816149df565b92915050565b600081519050613211816149df565b92915050565b600081359050613226816149f6565b92915050565b60008151905061323b816149f6565b92915050565b600082601f830112613256576132556147cb565b5b81356132668482602086016130e8565b91505092915050565b600082601f830112613284576132836147cb565b5b815161329484826020860161312a565b91505092915050565b600082601f8301126132b2576132b16147cb565b5b81356132c284826020860161316c565b91505092915050565b6000813590506132da81614a0d565b92915050565b6000815190506132ef81614a0d565b92915050565b60008135905061330481614a24565b92915050565b60008135905061331981614a3b565b92915050565b600060208284031215613335576133346147da565b5b6000613343848285016131c3565b91505092915050565b600060208284031215613362576133616147da565b5b6000613370848285016131d8565b91505092915050565b60008060008060808587031215613393576133926147da565b5b60006133a1878288016131ae565b94505060206133b2878288016131ae565b93505060406133c3878288016131ae565b92505060606133d4878288016132f5565b91505092959194509250565b6000602082840312156133f6576133f56147da565b5b600061340484828501613202565b91505092915050565b600060208284031215613423576134226147da565b5b60006134318482850161322c565b91505092915050565b60008060408385031215613451576134506147da565b5b600061345f85828601613217565b925050602061347085828601613217565b9150509250929050565b600080600060608486031215613493576134926147da565b5b60006134a186828701613217565b93505060206134b286828701613217565b92505060406134c386828701613217565b9150509250925092565b60008060008060008060c087890312156134ea576134e96147da565b5b60006134f889828a01613217565b965050602061350989828a01613217565b955050604061351a89828a016132f5565b945050606061352b89828a016131ae565b935050608061353c89828a0161330a565b92505060a061354d89828a016131ed565b9150509295509295509295565b60008060008060008060008060006101208a8c03121561357d5761357c6147da565b5b600061358b8c828d01613217565b995050602061359c8c828d01613217565b98505060406135ad8c828d016132f5565b97505060606135be8c828d016131ae565b96505060806135cf8c828d0161330a565b95505060a06135e08c828d016132cb565b94505060c06135f18c828d016132cb565b93505060e08a013567ffffffffffffffff811115613612576136116147d5565b5b61361e8c828d01613241565b9250506101006136308c828d016131ed565b9150509295985092959850929598565b600060208284031215613656576136556147da565b5b600082015167ffffffffffffffff811115613674576136736147d5565b5b6136808482850161326f565b91505092915050565b600080604083850312156136a05761369f6147da565b5b600083013567ffffffffffffffff8111156136be576136bd6147d5565b5b6136ca8582860161329d565b92505060206136db85828601613217565b9150509250929050565b600080600080608085870312156136ff576136fe6147da565b5b600085013567ffffffffffffffff81111561371d5761371c6147d5565b5b6137298782880161329d565b945050602061373a878288016132f5565b935050604061374b878288016131ae565b925050606061375c8782880161330a565b91505092959194509250565b600080600080600080600060e0888a031215613787576137866147da565b5b600088013567ffffffffffffffff8111156137a5576137a46147d5565b5b6137b18a828b0161329d565b97505060206137c28a828b016132f5565b96505060406137d38a828b016131ae565b95505060606137e48a828b0161330a565b94505060806137f58a828b016132cb565b93505060a06138068a828b016132cb565b92505060c088013567ffffffffffffffff811115613827576138266147d5565b5b6138338a828b01613241565b91505092959891949750929550565b600060208284031215613858576138576147da565b5b6000613866848285016132e0565b91505092915050565b600080600060608486031215613888576138876147da565b5b6000613896868287016132e0565b93505060206138a7868287016132e0565b92505060406138b886828701613202565b9150509250925092565b6138cb8161453a565b82525050565b6138e26138dd8261453a565b6146e2565b82525050565b6138f18161455e565b82525050565b6139008161456a565b82525050565b61390f8161456a565b82525050565b6139266139218261456a565b6146f4565b82525050565b60006139378261443a565b6139418185614450565b935061395181856020860161464c565b61395a816147df565b840191505092915050565b600081546139728161467f565b61397c8186614450565b9450600182166000811461399757600181146139a9576139dc565b60ff19831686526020860193506139dc565b6139b285614425565b60005b838110156139d4578154818901526001820191506020810190506139b5565b808801955050505b50505092915050565b6139ee816145d1565b82525050565b6139fd816145f5565b82525050565b613a0c81614619565b82525050565b6000613a1d82614445565b613a278185614461565b9350613a3781856020860161464c565b613a40816147df565b840191505092915050565b6000613a5682614445565b613a608185614472565b9350613a7081856020860161464c565b613a79816147df565b840191505092915050565b6000613a8f82614445565b613a998185614483565b9350613aa981856020860161464c565b80840191505092915050565b6000613ac2600c83614461565b9150613acd82614817565b602082019050919050565b6000613ae5601483614461565b9150613af082614840565b602082019050919050565b6000613b08600f83614461565b9150613b1382614869565b602082019050919050565b6000613b2b601383614461565b9150613b3682614892565b602082019050919050565b6000613b4e600e83614461565b9150613b59826148bb565b602082019050919050565b6000613b71600a83614461565b9150613b7c826148e4565b602082019050919050565b6000613b94601783614461565b9150613b9f8261490d565b602082019050919050565b6000613bb7600d83614461565b9150613bc282614936565b602082019050919050565b6000613bda601183614461565b9150613be58261495f565b602082019050919050565b6000613bfd600983614461565b9150613c0882614988565b602082019050919050565b613c1c81614594565b82525050565b613c33613c2e82614594565b614710565b82525050565b613c428161459e565b82525050565b613c51816145bb565b82525050565b613c68613c63826145bb565b61472c565b82525050565b613c77816145ae565b82525050565b613c8e613c89826145ae565b61471a565b82525050565b6000613ca0828a613915565b602082019150613cb08289613c7d565b600182019150613cc082886138d1565b601482019150613cd08287613c57565b600a82019150613ce08286613c22565b602082019150613cf08285613c22565b602082019150613d0082846138d1565b60148201915081905098975050505050505050565b6000613d21828a613a84565b9150613d2d8289613c7d565b600182019150613d3d82886138d1565b601482019150613d4d8287613c57565b600a82019150613d5d8286613c22565b602082019150613d6d8285613c22565b602082019150613d7d82846138d1565b60148201915081905098975050505050505050565b6000602082019050613da760008301846138c2565b92915050565b6000604082019050613dc260008301856138c2565b613dcf60208301846138c2565b9392505050565b6000606082019050613deb60008301866138c2565b613df860208301856138c2565b613e056040830184613c13565b949350505050565b6000604082019050613e2260008301856138c2565b613e2f6020830184613c13565b9392505050565b6000602082019050613e4b60008301846138f7565b92915050565b600060a082019050613e6660008301886138f7565b613e7360208301876138c2565b613e806040830186613c48565b613e8d6060830185613c6e565b613e9a6080830184613c6e565b9695505050505050565b6000606082019050613eb960008301866138f7565b613ec660208301856138f7565b613ed360408301846138f7565b949350505050565b600060a082019050613ef060008301886138f7565b613efd60208301876138f7565b613f0a6040830186613c6e565b613f1760608301856138c2565b613f2460808301846138e8565b9695505050505050565b600060e082019050613f43600083018a6138f7565b613f5060208301896138f7565b613f5d6040830188613c6e565b613f6a60608301876138c2565b613f776080830186613c13565b613f8460a0830185613c13565b613f9160c08301846138e8565b98975050505050505050565b6000604082019050613fb26000830185613906565b8181036020830152613fc4818461392c565b90509392505050565b6000606082019050613fe260008301866138f7565b613fef6020830185613c39565b613ffc60408301846138c2565b949350505050565b6000602082019050818103600083015261401e818461392c565b905092915050565b600060208201905081810360008301526140408184613965565b905092915050565b600060208201905061405d60008301846139e5565b92915050565b600060208201905061407860008301846139f4565b92915050565b60006020820190506140936000830184613a03565b92915050565b600060208201905081810360008301526140b38184613a4b565b905092915050565b600060208201905081810360008301526140d58184613a12565b905092915050565b600060408201905081810360008301526140f78185613a12565b905061410660208301846138f7565b9392505050565b600060608201905081810360008301526141278186613a12565b90506141366020830185613c39565b6141436040830184613c13565b949350505050565b600060808201905081810360008301526141658187613a12565b90506141746020830186613c6e565b61418160408301856138c2565b61418e6060830184613c13565b95945050505050565b600060c08201905081810360008301526141b18189613a12565b90506141c06020830188613c6e565b6141cd60408301876138c2565b6141da6060830186613c13565b6141e76080830185613c13565b6141f460a0830184613c13565b979650505050505050565b6000602082019050818103600083015261421881613ab5565b9050919050565b6000602082019050818103600083015261423881613ad8565b9050919050565b6000602082019050818103600083015261425881613afb565b9050919050565b6000602082019050818103600083015261427881613b1e565b9050919050565b6000602082019050818103600083015261429881613b41565b9050919050565b600060208201905081810360008301526142b881613b64565b9050919050565b600060208201905081810360008301526142d881613b87565b9050919050565b600060208201905081810360008301526142f881613baa565b9050919050565b6000602082019050818103600083015261431881613bcd565b9050919050565b6000602082019050818103600083015261433881613bf0565b9050919050565b60006020820190506143546000830184613c13565b92915050565b600060408201905061436f6000830185613c13565b61437c6020830184613c39565b9392505050565b60006020820190506143986000830184613c6e565b92915050565b60006143a86143b9565b90506143b482826146b1565b919050565b6000604051905090565b600067ffffffffffffffff8211156143de576143dd61479c565b5b6143e7826147df565b9050602081019050919050565b600067ffffffffffffffff82111561440f5761440e61479c565b5b614418826147df565b9050602081019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006144998261459e565b91506144a48361459e565b92508263ffffffff038211156144bd576144bc61473e565b5b828201905092915050565b60006144d38261459e565b91506144de8361459e565b92508163ffffffff04831182151516156144fb576144fa61473e565b5b828202905092915050565b600061451182614594565b915061451c83614594565b92508282101561452f5761452e61473e565b5b828203905092915050565b600061454582614574565b9050919050565b600061455782614574565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b60006145dc826145e3565b9050919050565b60006145ee82614574565b9050919050565b600061460082614607565b9050919050565b600061461282614574565b9050919050565b60006146248261462b565b9050919050565b600061463682614574565b9050919050565b82818337600083830152505050565b60005b8381101561466a57808201518184015260208101905061464f565b83811115614679576000848401525b50505050565b6000600282049050600182168061469757607f821691505b602082108114156146ab576146aa61476d565b5b50919050565b6146ba826147df565b810181811067ffffffffffffffff821117156146d9576146d861479c565b5b80604052505050565b60006146ed826146fe565b9050919050565b6000819050919050565b60006147098261480a565b9050919050565b6000819050919050565b6000614725826147fd565b9050919050565b6000614737826147f0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160b01b9050919050565b60008160f81b9050919050565b60008160601b9050919050565b7f63616e2774206368617267650000000000000000000000000000000000000000600082015250565b7f7061796f7574206973206e6f7420656e6f756768000000000000000000000000600082015250565b7f6e6f74206f70656e20746f207265670000000000000000000000000000000000600082015250565b7f6e616d6520776173207265676973746572656400000000000000000000000000600082015250565b7f6e616d65206e6f74206578697374000000000000000000000000000000000000600082015250565b7f63616e2774206d696e7400000000000000000000000000000000000000000000600082015250565b7f666174686572206e616d65206e6f7420636f7272656374000000000000000000600082015250565b7f736967206e6f74206d6174636800000000000000000000000000000000000000600082015250565b7f6e6f7420612076616c6964207072696365000000000000000000000000000000600082015250565b7f6e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b6149ba8161453a565b81146149c557600080fd5b50565b6149d18161454c565b81146149dc57600080fd5b50565b6149e88161455e565b81146149f357600080fd5b50565b6149ff8161456a565b8114614a0a57600080fd5b50565b614a1681614594565b8114614a2157600080fd5b50565b614a2d816145ae565b8114614a3857600080fd5b50565b614a44816145bb565b8114614a4f57600080fd5b5056fea2646970667358221220a343b9686a54ec9b164801482fc04fa0f16dcbb9db3430acc094a690d998c37c64736f6c63430008060033",
}

// DnsSecondLevelNameABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsSecondLevelNameMetaData.ABI instead.
var DnsSecondLevelNameABI = DnsSecondLevelNameMetaData.ABI

// DnsSecondLevelNameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DnsSecondLevelNameMetaData.Bin instead.
var DnsSecondLevelNameBin = DnsSecondLevelNameMetaData.Bin

// DeployDnsSecondLevelName deploys a new Ethereum contract, binding an instance of DnsSecondLevelName to it.
func DeployDnsSecondLevelName(auth *bind.TransactOpts, backend bind.ContractBackend, cost_ common.Address, acct_ common.Address, dnsTop_ common.Address) (common.Address, *types.Transaction, *DnsSecondLevelName, error) {
	parsed, err := DnsSecondLevelNameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	libDnsSignatureAddr, _, _, _ := DeployLibDnsSignature(auth, backend)
	DnsSecondLevelNameBin = strings.ReplaceAll(DnsSecondLevelNameBin, "__$2ed95cc1c04a020bf25a1d27d6730d85d6$__", libDnsSignatureAddr.String()[2:])

	libDnsToolKitAddr, _, _, _ := DeployLibDnsToolKit(auth, backend)
	DnsSecondLevelNameBin = strings.ReplaceAll(DnsSecondLevelNameBin, "__$61b1050b44c222c225346b0c1857883025$__", libDnsToolKitAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DnsSecondLevelNameBin), backend, cost_, acct_, dnsTop_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DnsSecondLevelName{DnsSecondLevelNameCaller: DnsSecondLevelNameCaller{contract: contract}, DnsSecondLevelNameTransactor: DnsSecondLevelNameTransactor{contract: contract}, DnsSecondLevelNameFilterer: DnsSecondLevelNameFilterer{contract: contract}}, nil
}

// DnsSecondLevelName is an auto generated Go binding around an Ethereum contract.
type DnsSecondLevelName struct {
	DnsSecondLevelNameCaller     // Read-only binding to the contract
	DnsSecondLevelNameTransactor // Write-only binding to the contract
	DnsSecondLevelNameFilterer   // Log filterer for contract events
}

// DnsSecondLevelNameCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsSecondLevelNameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsSecondLevelNameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsSecondLevelNameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsSecondLevelNameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsSecondLevelNameSession struct {
	Contract     *DnsSecondLevelName // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DnsSecondLevelNameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsSecondLevelNameCallerSession struct {
	Contract *DnsSecondLevelNameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DnsSecondLevelNameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsSecondLevelNameTransactorSession struct {
	Contract     *DnsSecondLevelNameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DnsSecondLevelNameRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsSecondLevelNameRaw struct {
	Contract *DnsSecondLevelName // Generic contract binding to access the raw methods on
}

// DnsSecondLevelNameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsSecondLevelNameCallerRaw struct {
	Contract *DnsSecondLevelNameCaller // Generic read-only contract binding to access the raw methods on
}

// DnsSecondLevelNameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsSecondLevelNameTransactorRaw struct {
	Contract *DnsSecondLevelNameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsSecondLevelName creates a new instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelName(address common.Address, backend bind.ContractBackend) (*DnsSecondLevelName, error) {
	contract, err := bindDnsSecondLevelName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelName{DnsSecondLevelNameCaller: DnsSecondLevelNameCaller{contract: contract}, DnsSecondLevelNameTransactor: DnsSecondLevelNameTransactor{contract: contract}, DnsSecondLevelNameFilterer: DnsSecondLevelNameFilterer{contract: contract}}, nil
}

// NewDnsSecondLevelNameCaller creates a new read-only instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameCaller(address common.Address, caller bind.ContractCaller) (*DnsSecondLevelNameCaller, error) {
	contract, err := bindDnsSecondLevelName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameCaller{contract: contract}, nil
}

// NewDnsSecondLevelNameTransactor creates a new write-only instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsSecondLevelNameTransactor, error) {
	contract, err := bindDnsSecondLevelName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameTransactor{contract: contract}, nil
}

// NewDnsSecondLevelNameFilterer creates a new log filterer instance of DnsSecondLevelName, bound to a specific deployed contract.
func NewDnsSecondLevelNameFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsSecondLevelNameFilterer, error) {
	contract, err := bindDnsSecondLevelName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameFilterer{contract: contract}, nil
}

// bindDnsSecondLevelName binds a generic wrapper to an already deployed contract.
func bindDnsSecondLevelName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsSecondLevelNameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSecondLevelName *DnsSecondLevelNameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DnsSecondLevelNameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsSecondLevelName *DnsSecondLevelNameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsSecondLevelName.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.contract.Transact(opts, method, params...)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) AccountantC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "accountantC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) AccountantC() (common.Address, error) {
	return _DnsSecondLevelName.Contract.AccountantC(&_DnsSecondLevelName.CallOpts)
}

// AccountantC is a free data retrieval call binding the contract method 0x12014f01.
//
// Solidity: function accountantC() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) AccountantC() (common.Address, error) {
	return _DnsSecondLevelName.Contract.AccountantC(&_DnsSecondLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) CostContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "costContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) CostContractAddr() (common.Address, error) {
	return _DnsSecondLevelName.Contract.CostContractAddr(&_DnsSecondLevelName.CallOpts)
}

// CostContractAddr is a free data retrieval call binding the contract method 0x6130457d.
//
// Solidity: function costContractAddr() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) CostContractAddr() (common.Address, error) {
	return _DnsSecondLevelName.Contract.CostContractAddr(&_DnsSecondLevelName.CallOpts)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) DnsTop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "dnsTop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) DnsTop() (common.Address, error) {
	return _DnsSecondLevelName.Contract.DnsTop(&_DnsSecondLevelName.CallOpts)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) DnsTop() (common.Address, error) {
	return _DnsSecondLevelName.Contract.DnsTop(&_DnsSecondLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) MintSwitch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "mintSwitch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSwitch() (uint8, error) {
	return _DnsSecondLevelName.Contract.MintSwitch(&_DnsSecondLevelName.CallOpts)
}

// MintSwitch is a free data retrieval call binding the contract method 0xeacb912d.
//
// Solidity: function mintSwitch() view returns(uint8)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) MintSwitch() (uint8, error) {
	return _DnsSecondLevelName.Contract.MintSwitch(&_DnsSecondLevelName.CallOpts)
}

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) NameStore(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "nameStore", arg0, arg1)

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

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) NameStore(arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsSecondLevelName.Contract.NameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// NameStore is a free data retrieval call binding the contract method 0x09d52623.
//
// Solidity: function nameStore(bytes32 , bytes32 ) view returns(string entireName, uint32 expireTime, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) NameStore(arg0 [32]byte, arg1 [32]byte) (struct {
	EntireName string
	ExpireTime uint32
	TokenId    *big.Int
}, error) {
	return _DnsSecondLevelName.Contract.NameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameCaller) SubNameStore(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (string, error) {
	var out []interface{}
	err := _DnsSecondLevelName.contract.Call(opts, &out, "subNameStore", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameSession) SubNameStore(arg0 [32]byte, arg1 [32]byte) (string, error) {
	return _DnsSecondLevelName.Contract.SubNameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// SubNameStore is a free data retrieval call binding the contract method 0x9150cf8f.
//
// Solidity: function subNameStore(bytes32 , bytes32 ) view returns(string)
func (_DnsSecondLevelName *DnsSecondLevelNameCallerSession) SubNameStore(arg0 [32]byte, arg1 [32]byte) (string, error) {
	return _DnsSecondLevelName.Contract.SubNameStore(&_DnsSecondLevelName.CallOpts, arg0, arg1)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) AddSubName(opts *bind.TransactOpts, entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "AddSubName", entireSubName_, level2FatherHash_)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) AddSubName(entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.AddSubName(&_DnsSecondLevelName.TransactOpts, entireSubName_, level2FatherHash_)
}

// AddSubName is a paid mutator transaction binding the contract method 0x06475e73.
//
// Solidity: function AddSubName(string entireSubName_, bytes32 level2FatherHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) AddSubName(entireSubName_ string, level2FatherHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.AddSubName(&_DnsSecondLevelName.TransactOpts, entireSubName_, level2FatherHash_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) ChargeSecondLevelName(opts *bind.TransactOpts, fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "ChargeSecondLevelName", fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) ChargeSecondLevelName(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelName(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelName is a paid mutator transaction binding the contract method 0x752ab03d.
//
// Solidity: function ChargeSecondLevelName(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) ChargeSecondLevelName(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelName(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) ChargeSecondLevelNameBySig(opts *bind.TransactOpts, fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "ChargeSecondLevelNameBySig", fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) ChargeSecondLevelNameBySig(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// ChargeSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xb3539630.
//
// Solidity: function ChargeSecondLevelNameBySig(bytes32 fatherHash_, bytes32 nameHash_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig, bool isTransfer_) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) ChargeSecondLevelNameBySig(fatherHash_ [32]byte, nameHash_ [32]byte, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte, isTransfer_ bool) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.ChargeSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, fatherHash_, nameHash_, year_, erc20Addr_, lastPriceId, nonce, price_, sig, isTransfer_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) DelSubName(opts *bind.TransactOpts, nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "DelSubName", nameHash_, level2FatherHash_, topHash_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) DelSubName(nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DelSubName(&_DnsSecondLevelName.TransactOpts, nameHash_, level2FatherHash_, topHash_)
}

// DelSubName is a paid mutator transaction binding the contract method 0x3752770e.
//
// Solidity: function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) DelSubName(nameHash_ [32]byte, level2FatherHash_ [32]byte, topHash_ [32]byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.DelSubName(&_DnsSecondLevelName.TransactOpts, nameHash_, level2FatherHash_, topHash_)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) MintSecondLevelName(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "MintSecondLevelName", entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSecondLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelName(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelName is a paid mutator transaction binding the contract method 0xc66aee2d.
//
// Solidity: function MintSecondLevelName(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) MintSecondLevelName(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelName(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) MintSecondLevelNameBySig(opts *bind.TransactOpts, entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "MintSecondLevelNameBySig", entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) MintSecondLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// MintSecondLevelNameBySig is a paid mutator transaction binding the contract method 0xbcfb5505.
//
// Solidity: function MintSecondLevelNameBySig(string entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId, uint256 nonce, uint256 price_, bytes sig) payable returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) MintSecondLevelNameBySig(entireName_ string, year_ uint8, erc20Addr_ common.Address, lastPriceId *big.Int, nonce *big.Int, price_ *big.Int, sig []byte) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.MintSecondLevelNameBySig(&_DnsSecondLevelName.TransactOpts, entireName_, year_, erc20Addr_, lastPriceId, nonce, price_, sig)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactor) SetContract(opts *bind.TransactOpts, cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.contract.Transact(opts, "setContract", cost_, acct_, dnsTop_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameSession) SetContract(cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.SetContract(&_DnsSecondLevelName.TransactOpts, cost_, acct_, dnsTop_, mintSw_)
}

// SetContract is a paid mutator transaction binding the contract method 0xfddd23a8.
//
// Solidity: function setContract(address cost_, address acct_, address dnsTop_, uint8 mintSw_) returns()
func (_DnsSecondLevelName *DnsSecondLevelNameTransactorSession) SetContract(cost_ common.Address, acct_ common.Address, dnsTop_ common.Address, mintSw_ uint8) (*types.Transaction, error) {
	return _DnsSecondLevelName.Contract.SetContract(&_DnsSecondLevelName.TransactOpts, cost_, acct_, dnsTop_, mintSw_)
}

// DnsSecondLevelNameEvAddSubNameIterator is returned from FilterEvAddSubName and is used to iterate over the raw logs and unpacked data for EvAddSubName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvAddSubNameIterator struct {
	Event *DnsSecondLevelNameEvAddSubName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvAddSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvAddSubName)
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
		it.Event = new(DnsSecondLevelNameEvAddSubName)
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
func (it *DnsSecondLevelNameEvAddSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvAddSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvAddSubName represents a EvAddSubName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvAddSubName struct {
	EntireSubName    string
	Level2FatherHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEvAddSubName is a free log retrieval operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvAddSubName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvAddSubNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvAddSubNameIterator{contract: _DnsSecondLevelName.contract, event: "EvAddSubName", logs: logs, sub: sub}, nil
}

// WatchEvAddSubName is a free log subscription operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvAddSubName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvAddSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvAddSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvAddSubName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
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

// ParseEvAddSubName is a log parse operation binding the contract event 0x3c9525b151e4fb56f1695beb201da54c7d7138da2e3b3b6f144d6942ce1215b8.
//
// Solidity: event EvAddSubName(string entireSubName, bytes32 level2FatherHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvAddSubName(log types.Log) (*DnsSecondLevelNameEvAddSubName, error) {
	event := new(DnsSecondLevelNameEvAddSubName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvAddSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameIterator is returned from FilterEvChargeSecondLevelName and is used to iterate over the raw logs and unpacked data for EvChargeSecondLevelName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameIterator struct {
	Event *DnsSecondLevelNameEvChargeSecondLevelName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvChargeSecondLevelName)
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
		it.Event = new(DnsSecondLevelNameEvChargeSecondLevelName)
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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvChargeSecondLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvChargeSecondLevelName represents a EvChargeSecondLevelName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelName struct {
	FatherHash [32]byte
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeSecondLevelName is a free log retrieval operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvChargeSecondLevelName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvChargeSecondLevelNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvChargeSecondLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvChargeSecondLevelNameIterator{contract: _DnsSecondLevelName.contract, event: "EvChargeSecondLevelName", logs: logs, sub: sub}, nil
}

// WatchEvChargeSecondLevelName is a free log subscription operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvChargeSecondLevelName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvChargeSecondLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvChargeSecondLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvChargeSecondLevelName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelName", log); err != nil {
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

// ParseEvChargeSecondLevelName is a log parse operation binding the contract event 0x920d2a7be40a3b09552193f59134b5ab441a2147c7f7ea33a92aa7e62fe51640.
//
// Solidity: event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvChargeSecondLevelName(log types.Log) (*DnsSecondLevelNameEvChargeSecondLevelName, error) {
	event := new(DnsSecondLevelNameEvChargeSecondLevelName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator is returned from FilterEvChargeSecondLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvChargeSecondLevelNameBySig events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator struct {
	Event *DnsSecondLevelNameEvChargeSecondLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
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
		it.Event = new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
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
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvChargeSecondLevelNameBySig represents a EvChargeSecondLevelNameBySig event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvChargeSecondLevelNameBySig struct {
	FatherHash [32]byte
	NameHash   [32]byte
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	IsTransfer bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvChargeSecondLevelNameBySig is a free log retrieval operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvChargeSecondLevelNameBySig(opts *bind.FilterOpts) (*DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvChargeSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvChargeSecondLevelNameBySigIterator{contract: _DnsSecondLevelName.contract, event: "EvChargeSecondLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvChargeSecondLevelNameBySig is a free log subscription operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvChargeSecondLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvChargeSecondLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvChargeSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelNameBySig", log); err != nil {
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

// ParseEvChargeSecondLevelNameBySig is a log parse operation binding the contract event 0x07e1ec45292c9a9eaee63fac915434fe52b48b78097f99b43091754e210ff5a7.
//
// Solidity: event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvChargeSecondLevelNameBySig(log types.Log) (*DnsSecondLevelNameEvChargeSecondLevelNameBySig, error) {
	event := new(DnsSecondLevelNameEvChargeSecondLevelNameBySig)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvChargeSecondLevelNameBySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvDelSubNameIterator is returned from FilterEvDelSubName and is used to iterate over the raw logs and unpacked data for EvDelSubName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvDelSubNameIterator struct {
	Event *DnsSecondLevelNameEvDelSubName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvDelSubNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvDelSubName)
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
		it.Event = new(DnsSecondLevelNameEvDelSubName)
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
func (it *DnsSecondLevelNameEvDelSubNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvDelSubNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvDelSubName represents a EvDelSubName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvDelSubName struct {
	NameHash         [32]byte
	Level2FatherHash [32]byte
	TopHash          [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEvDelSubName is a free log retrieval operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvDelSubName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvDelSubNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvDelSubNameIterator{contract: _DnsSecondLevelName.contract, event: "EvDelSubName", logs: logs, sub: sub}, nil
}

// WatchEvDelSubName is a free log subscription operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvDelSubName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvDelSubName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvDelSubName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvDelSubName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
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

// ParseEvDelSubName is a log parse operation binding the contract event 0x1b02e665d84f2a8d31cdf5ecf8fef1a617a40ad108d47282f45551a86da8ed4a.
//
// Solidity: event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvDelSubName(log types.Log) (*DnsSecondLevelNameEvDelSubName, error) {
	event := new(DnsSecondLevelNameEvDelSubName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvDelSubName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvMintSecondLevelNameIterator is returned from FilterEvMintSecondLevelName and is used to iterate over the raw logs and unpacked data for EvMintSecondLevelName events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameIterator struct {
	Event *DnsSecondLevelNameEvMintSecondLevelName // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvMintSecondLevelName)
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
		it.Event = new(DnsSecondLevelNameEvMintSecondLevelName)
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
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvMintSecondLevelNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvMintSecondLevelName represents a EvMintSecondLevelName event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelName struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintSecondLevelName is a free log retrieval operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvMintSecondLevelName(opts *bind.FilterOpts) (*DnsSecondLevelNameEvMintSecondLevelNameIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvMintSecondLevelName")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvMintSecondLevelNameIterator{contract: _DnsSecondLevelName.contract, event: "EvMintSecondLevelName", logs: logs, sub: sub}, nil
}

// WatchEvMintSecondLevelName is a free log subscription operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvMintSecondLevelName(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvMintSecondLevelName) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvMintSecondLevelName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvMintSecondLevelName)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelName", log); err != nil {
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

// ParseEvMintSecondLevelName is a log parse operation binding the contract event 0x500ebb53308b99f82ed59d05e3865086a18f7125c8dded3a3328df8feec0eff9.
//
// Solidity: event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvMintSecondLevelName(log types.Log) (*DnsSecondLevelNameEvMintSecondLevelName, error) {
	event := new(DnsSecondLevelNameEvMintSecondLevelName)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DnsSecondLevelNameEvMintSecondLevelNameBySigIterator is returned from FilterEvMintSecondLevelNameBySig and is used to iterate over the raw logs and unpacked data for EvMintSecondLevelNameBySig events raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameBySigIterator struct {
	Event *DnsSecondLevelNameEvMintSecondLevelNameBySig // Event containing the contract specifics and raw log

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
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
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
		it.Event = new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
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
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnsSecondLevelNameEvMintSecondLevelNameBySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnsSecondLevelNameEvMintSecondLevelNameBySig represents a EvMintSecondLevelNameBySig event raised by the DnsSecondLevelName contract.
type DnsSecondLevelNameEvMintSecondLevelNameBySig struct {
	EntireName string
	Year       uint8
	Erc20Addr  common.Address
	Nonce      *big.Int
	Price      *big.Int
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEvMintSecondLevelNameBySig is a free log retrieval operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) FilterEvMintSecondLevelNameBySig(opts *bind.FilterOpts) (*DnsSecondLevelNameEvMintSecondLevelNameBySigIterator, error) {

	logs, sub, err := _DnsSecondLevelName.contract.FilterLogs(opts, "EvMintSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return &DnsSecondLevelNameEvMintSecondLevelNameBySigIterator{contract: _DnsSecondLevelName.contract, event: "EvMintSecondLevelNameBySig", logs: logs, sub: sub}, nil
}

// WatchEvMintSecondLevelNameBySig is a free log subscription operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) WatchEvMintSecondLevelNameBySig(opts *bind.WatchOpts, sink chan<- *DnsSecondLevelNameEvMintSecondLevelNameBySig) (event.Subscription, error) {

	logs, sub, err := _DnsSecondLevelName.contract.WatchLogs(opts, "EvMintSecondLevelNameBySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
				if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelNameBySig", log); err != nil {
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

// ParseEvMintSecondLevelNameBySig is a log parse operation binding the contract event 0xd96e0aac740ffa8e9c0796d9c5934aeb8b1088f40cf9e5655677374ce908f5d6.
//
// Solidity: event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId)
func (_DnsSecondLevelName *DnsSecondLevelNameFilterer) ParseEvMintSecondLevelNameBySig(log types.Log) (*DnsSecondLevelNameEvMintSecondLevelNameBySig, error) {
	event := new(DnsSecondLevelNameEvMintSecondLevelNameBySig)
	if err := _DnsSecondLevelName.contract.UnpackLog(event, "EvMintSecondLevelNameBySig", log); err != nil {
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
