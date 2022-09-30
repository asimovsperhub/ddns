// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cost

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

// AggregatorV3InterfaceMetaData contains all meta data concerning the AggregatorV3Interface contract.
var AggregatorV3InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AggregatorV3InterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorV3InterfaceMetaData.ABI instead.
var AggregatorV3InterfaceABI = AggregatorV3InterfaceMetaData.ABI

// AggregatorV3Interface is an auto generated Go binding around an Ethereum contract.
type AggregatorV3Interface struct {
	AggregatorV3InterfaceCaller     // Read-only binding to the contract
	AggregatorV3InterfaceTransactor // Write-only binding to the contract
	AggregatorV3InterfaceFilterer   // Log filterer for contract events
}

// AggregatorV3InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorV3InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorV3InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorV3InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorV3InterfaceSession struct {
	Contract     *AggregatorV3Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AggregatorV3InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorV3InterfaceCallerSession struct {
	Contract *AggregatorV3InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AggregatorV3InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorV3InterfaceTransactorSession struct {
	Contract     *AggregatorV3InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AggregatorV3InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorV3InterfaceRaw struct {
	Contract *AggregatorV3Interface // Generic contract binding to access the raw methods on
}

// AggregatorV3InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorV3InterfaceCallerRaw struct {
	Contract *AggregatorV3InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorV3InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorV3InterfaceTransactorRaw struct {
	Contract *AggregatorV3InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorV3Interface creates a new instance of AggregatorV3Interface, bound to a specific deployed contract.
func NewAggregatorV3Interface(address common.Address, backend bind.ContractBackend) (*AggregatorV3Interface, error) {
	contract, err := bindAggregatorV3Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3Interface{AggregatorV3InterfaceCaller: AggregatorV3InterfaceCaller{contract: contract}, AggregatorV3InterfaceTransactor: AggregatorV3InterfaceTransactor{contract: contract}, AggregatorV3InterfaceFilterer: AggregatorV3InterfaceFilterer{contract: contract}}, nil
}

// NewAggregatorV3InterfaceCaller creates a new read-only instance of AggregatorV3Interface, bound to a specific deployed contract.
func NewAggregatorV3InterfaceCaller(address common.Address, caller bind.ContractCaller) (*AggregatorV3InterfaceCaller, error) {
	contract, err := bindAggregatorV3Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceCaller{contract: contract}, nil
}

// NewAggregatorV3InterfaceTransactor creates a new write-only instance of AggregatorV3Interface, bound to a specific deployed contract.
func NewAggregatorV3InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorV3InterfaceTransactor, error) {
	contract, err := bindAggregatorV3Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceTransactor{contract: contract}, nil
}

// NewAggregatorV3InterfaceFilterer creates a new log filterer instance of AggregatorV3Interface, bound to a specific deployed contract.
func NewAggregatorV3InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorV3InterfaceFilterer, error) {
	contract, err := bindAggregatorV3Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceFilterer{contract: contract}, nil
}

// bindAggregatorV3Interface binds a generic wrapper to an already deployed contract.
func bindAggregatorV3Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorV3InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorV3Interface *AggregatorV3InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorV3Interface *AggregatorV3InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Decimals() (uint8, error) {
	return _AggregatorV3Interface.Contract.Decimals(&_AggregatorV3Interface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Decimals() (uint8, error) {
	return _AggregatorV3Interface.Contract.Decimals(&_AggregatorV3Interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Description() (string, error) {
	return _AggregatorV3Interface.Contract.Description(&_AggregatorV3Interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Description() (string, error) {
	return _AggregatorV3Interface.Contract.Description(&_AggregatorV3Interface.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.GetRoundData(&_AggregatorV3Interface.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.GetRoundData(&_AggregatorV3Interface.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.LatestRoundData(&_AggregatorV3Interface.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.LatestRoundData(&_AggregatorV3Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Version() (*big.Int, error) {
	return _AggregatorV3Interface.Contract.Version(&_AggregatorV3Interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Version() (*big.Int, error) {
	return _AggregatorV3Interface.Contract.Version(&_AggregatorV3Interface.CallOpts)
}

// DnsCostMetaData contains all meta data concerning the DnsCost contract.
var DnsCostMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"top_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"tax_\",\"type\":\"uint256[8]\"}],\"name\":\"SetSecondLevelNameTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"price_\",\"type\":\"uint256[8]\"}],\"name\":\"SetTopLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"chainLinkContact_\",\"type\":\"address\"}],\"name\":\"addPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseDecimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chainLinkRateAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dnsTop\",\"outputs\":[{\"internalType\":\"contractIDnsTopLevelName\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPayments\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"}],\"name\":\"getAllSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSecondLevelNameTax\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getExchangeWithId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getLatestExchangeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"paymentIsExist\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"priceIdIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dnstop_\",\"type\":\"address\"}],\"name\":\"setRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"prices_\",\"type\":\"uint256[8]\"}],\"name\":\"setSecondLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdtTokenAddr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimal_\",\"type\":\"uint8\"}],\"name\":\"setUsdtAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003d5438038062003d54833981810160405281019062000037919062000782565b60005b60088110156200007b5763b2d05e006009826008811062000060576200005f62000907565b5b0181905550808062000072906200088a565b9150506200003a565b5060005b6008811015620000bf57624c4b4060018260088110620000a457620000a362000907565b5b01819055508080620000b6906200088a565b9150506200007f565b50734f4f07917e13785bfd55cd3485b254bde6f092656000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506006600060146101000a81548160ff021916908360ff1602179055506040518060400160405280601260ff16815260200160011515815250601260008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d8909160006040518363ffffffff1660e01b81526004016200020e92919062000813565b604080518083038186803b1580156200022657600080fd5b505af41580156200023b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002619190620007b4565b50506040518060400160405280601260ff168152602001600115158152506012600073571787b5e033bf8bb2a1e5930265828ef3ffca0073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d8909173571787b5e033bf8bb2a1e5930265828ef3ffca006040518363ffffffff1660e01b81526004016200036892919062000813565b604080518083038186803b1580156200038057600080fd5b505af415801562000395573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003bb9190620007b4565b50506040518060400160405280600660ff16815260200160011515815250601260007365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d890917365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e6040518363ffffffff1660e01b8152600401620004c292919062000813565b604080518083038186803b158015620004da57600080fd5b505af4158015620004ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005159190620007b4565b505080601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073e0005488e0d9be7b24cce51f722d24749391327f601460008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073136a734040f258054da7ebf42bd25b73e913839a6014600073571787b5e033bf8bb2a1e5930265828ef3ffca0073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550738857aaea5c663a4408963e8d964d22e438d596f2601460007365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000989565b6000815190506200074e816200093b565b92915050565b600081519050620007658162000955565b92915050565b6000815190506200077c816200096f565b92915050565b6000602082840312156200079b576200079a62000936565b5b6000620007ab848285016200073d565b91505092915050565b60008060408385031215620007ce57620007cd62000936565b5b6000620007de8582860162000754565b9250506020620007f1858286016200076b565b9150509250929050565b620008068162000840565b82525050565b8082525050565b60006040820190506200082a60008301856200080c565b620008396020830184620007fb565b9392505050565b60006200084d8262000860565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000620008978262000880565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620008cd57620008cc620008d8565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b620009468162000840565b81146200095257600080fd5b50565b620009608162000854565b81146200096c57600080fd5b50565b6200097a8162000880565b81146200098657600080fd5b50565b6133bb80620009996000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80639fa001ef116100f9578063d34f422211610097578063f32ec51811610071578063f32ec5181461052f578063f37e6f1714610560578063f3c8bb0c1461057e578063f8defa3c1461059c576101a9565b8063d34f4222146104b1578063dabff59a146104e1578063db7b373e14610511576101a9565b8063cacf1e0e116100d3578063cacf1e0e1461042b578063ccc7905814610449578063cfdbee1114610479578063d3443c6214610495576101a9565b80639fa001ef146103ae578063b84ff547146103de578063bba4095c146103fa576101a9565b80635f3dda0211610166578063793908931161014057806379390893146102eb5780638412de081461031c57806395297fb01461034c5780639f33b2d81461037c576101a9565b80635f3dda021461028357806372be125d1461029f57806373260cb9146102bb576101a9565b80630868f7b1146101ae5780631554e1ce146101ca5780631d869ca1146101e85780633325d51c146102185780633dd45f1314610234578063546fa0be14610265575b600080fd5b6101c860048036038101906101c39190612302565b6105ba565b005b6101d261070b565b6040516101df9190612a6a565b60405180910390f35b61020260048036038101906101fd9190612302565b61072f565b60405161020f9190612a6a565b60405180910390f35b610232600480360381019061022d9190612503565b610762565b005b61024e6004803603810190610249919061249c565b610a1d565b60405161025c929190612bf8565b60405180910390f35b61026d610b1f565b60405161027a9190612aae565b60405180910390f35b61029d600480360381019061029891906125a9565b610b6a565b005b6102b960048036038101906102b491906126bf565b610ccb565b005b6102d560048036038101906102d09190612740565b610ed5565b6040516102e29190612bdd565b60405180910390f35b61030560048036038101906103009190612449565b610ef0565b604051610313929190612bf8565b60405180910390f35b61033660048036038101906103319190612700565b610f0a565b6040516103439190612bdd565b60405180910390f35b61036660048036038101906103619190612409565b610f32565b6040516103739190612aca565b60405180910390f35b61039660048036038101906103919190612644565b611147565b6040516103a593929190612c21565b60405180910390f35b6103c860048036038101906103c39190612740565b611339565b6040516103d59190612bdd565b60405180910390f35b6103f860048036038101906103f391906125a9565b611354565b005b610414600480360381019061040f9190612389565b6114b5565b604051610422929190612c58565b60405180910390f35b6104336116c1565b6040516104409190612bdd565b60405180910390f35b610463600480360381019061045e9190612740565b6116ce565b6040516104709190612a6a565b60405180910390f35b610493600480360381019061048e91906123c9565b61170d565b005b6104af60048036038101906104aa9190612556565b611879565b005b6104cb60048036038101906104c69190612302565b611abd565b6040516104d89190612c9c565b60405180910390f35b6104fb60048036038101906104f69190612617565b611b16565b6040516105089190612aae565b60405180910390f35b610519611b74565b6040516105269190612c9c565b60405180910390f35b61054960048036038101906105449190612302565b611b87565b604051610557929190612cb7565b60405180910390f35b610568611bc5565b6040516105759190612aae565b60405180910390f35b610586611c10565b6040516105939190612b22565b60405180910390f35b6105a4611c36565b6040516105b19190612b00565b60405180910390f35b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062257600080fd5b505afa158015610636573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065a919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106be90612b7d565b60405180910390fd5b80601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60146020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107ca57600080fd5b505afa1580156107de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610802919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461086f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086690612b7d565b60405180910390fd5b60405180604001604052808360ff16815260200160011515815250601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d89091856040518363ffffffff1660e01b8152600401610949929190612a85565b604080518083038186803b15801561096057600080fd5b505af4158015610974573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099891906125d7565b505080601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008060008360ff16610a2f86611cc6565b610a399190612f27565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161415610a9c578060019250925050610b16565b60001515601260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff1615151415610b05576000809250925050610b16565b610b10878783611d04565b92509250505b94509492505050565b610b276121ac565b6001600880602002604051908101604052809291908260088015610b60576020028201915b815481526020019060010190808311610b4c575b5050505050905090565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bd257600080fd5b505afa158015610be6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0a919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90612b7d565b60405180910390fd5b60005b6008811015610cc757818160088110610c9657610c9561319c565b5b602002013560018260088110610caf57610cae61319c565b5b01819055508080610cbf906130d1565b915050610c7a565b5050565b6000801b8214158015610e3057503373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76846040518263ffffffff1660e01b8152600401610d4b9190612ae5565b60206040518083038186803b158015610d6357600080fd5b505afa158015610d77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9b919061232f565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610de057600080fd5b505afa158015610df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e18919061235c565b73ffffffffffffffffffffffffffffffffffffffff16145b610e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6690612b7d565b60405180910390fd5b60005b6008811015610ed057818160088110610e8e57610e8d61319c565b5b6020020135601160008581526020019081526020016000208260088110610eb857610eb761319c565b5b01819055508080610ec8906130d1565b915050610e72565b505050565b60098160088110610ee557600080fd5b016000915090505481565b600080610efe858585611d04565b91509150935093915050565b60116020528160005260406000208160088110610f2657600080fd5b01600091509150505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610f925760019050611141565b60001515601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff1615151415610ff75760009050611141565b6000601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b1580156110a357600080fd5b505afa1580156110b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110db919061276d565b5050505090508069ffffffffffffffffffff168469ffffffffffffffffffff16148061112a57508069ffffffffffffffffffff1660018561111c9190612d45565b69ffffffffffffffffffff16145b1561113a57600192505050611141565b6000925050505b92915050565b6000806000808460ff1661115e8a8860ff1661207f565b6111689190612f27565b905060008560ff16611179886120f1565b6111839190612f27565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1614156111ea5780826001945094509450505061132e565b601260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff16611279576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161127090612b9d565b60405180910390fd5b6000806112878b8b86611d04565b91509150806112cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c290612bbd565b60405180910390fd5b6000806112d98d8d87611d04565b915091508061131d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131490612b3d565b60405180910390fd5b818460019850985098505050505050505b955095509592505050565b6001816008811061134957600080fd5b016000915090505481565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156113bc57600080fd5b505afa1580156113d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f4919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145890612b7d565b60405180910390fd5b60005b60088110156114b1578181600881106114805761147f61319c565b5b6020020135600982600881106114995761149861319c565b5b018190555080806114a9906130d1565b915050611464565b5050565b6000806000601460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561156557600080fd5b505afa158015611579573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159d919061276d565b505050915091506000818473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156115ed57600080fd5b505afa158015611601573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162591906127e8565b600a6116319190612e09565b8861163c9190612f27565b6116469190612d85565b90506116b1600060149054906101000a900460ff16601260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff168361212f565b8395509550505050509250929050565b6000601380549050905090565b601381815481106116de57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561177557600080fd5b505afa158015611789573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ad919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461181a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161181190612b7d565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060146101000a81548160ff021916908360ff1602179055505050565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156118e157600080fd5b505afa1580156118f5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611919919061235c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611986576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161197d90612b7d565b60405180910390fd5b6000601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1660ff1611611a1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1290612b5d565b60405180910390fd5b60405180604001604052808360ff168152602001821515815250601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050505050565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169050919050565b611b1e6121ac565b60116000838152602001908152602001600020600880602002604051908101604052809291908260088015611b68576020028201915b815481526020019060010190808311611b54575b50505050509050919050565b600060149054906101000a900460ff1681565b60126020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900460ff16905082565b611bcd6121ac565b6009600880602002604051908101604052809291908260088015611c06576020028201915b815481526020019060010190808311611bf2575b5050505050905090565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60608060005b601380549050811015611cbe578160138281548110611c5e57611c5d61319c565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001611c9a929190612a42565b60405160208183030381529060405291508080611cb6906130d1565b915050611c3c565b508091505090565b600060088260ff1610611ce45760016008611ce19190612fb5565b91505b60098260ff1660088110611cfb57611cfa61319c565b5b01549050919050565b6000806000601460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b158015611db457600080fd5b505afa158015611dc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dec919061276d565b505050915091506000818473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015611e3c57600080fd5b505afa158015611e50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e7491906127e8565b600a611e809190612e09565b88611e8b9190612f27565b611e959190612d85565b90508769ffffffffffffffffffff168369ffffffffffffffffffff161415611f2f57611f20600060149054906101000a900460ff16601260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff168361212f565b60019550955050505050612077565b600188611f3c9190612d45565b69ffffffffffffffffffff168369ffffffffffffffffffff16141561206b578373ffffffffffffffffffffffffffffffffffffffff16639a6fc8f5896040518263ffffffff1660e01b8152600401611f949190612c81565b60a06040518083038186803b158015611fac57600080fd5b505afa158015611fc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe4919061276d565b90919293509091509050508092505061205c600060149054906101000a900460ff16601260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff168361212f565b60019550955050505050612077565b60008095509550505050505b935093915050565b60006011600084815260200190815260200160002050600882106120c2576001601160008581526020019081526020016000205060086120bf9190612f81565b91505b6011600084815260200190815260200160002082600881106120e7576120e661319c565b5b0154905092915050565b600060088260ff161061210f576001600861210c9190612f81565b91505b60018260ff16600881106121265761212561319c565b5b01549050919050565b60008260ff168460ff161415612147578190506121a5565b8260ff168460ff16111561217f5782846121619190612fb5565b600a61216d9190612e09565b826121789190612d85565b90506121a5565b838361218b9190612fb5565b600a6121979190612e09565b826121a29190612f27565b90505b9392505050565b604051806101000160405280600890602082028036833780820191505090505090565b6000813590506121de816132cd565b92915050565b6000815190506121f3816132cd565b92915050565b600081519050612208816132e4565b92915050565b60008190508260206008028201111561222a576122296131cb565b5b92915050565b60008135905061223f816132fb565b92915050565b600081519050612254816132fb565b92915050565b60008135905061226981613312565b92915050565b60008151905061227e81613329565b92915050565b60008135905061229381613340565b92915050565b6000815190506122a881613340565b92915050565b6000813590506122bd81613357565b92915050565b6000813590506122d28161336e565b92915050565b6000815190506122e78161336e565b92915050565b6000815190506122fc81613357565b92915050565b600060208284031215612318576123176131d0565b5b6000612326848285016121cf565b91505092915050565b600060208284031215612345576123446131d0565b5b6000612353848285016121e4565b91505092915050565b600060208284031215612372576123716131d0565b5b6000612380848285016121f9565b91505092915050565b600080604083850312156123a05761239f6131d0565b5b60006123ae858286016121cf565b92505060206123bf85828601612284565b9150509250929050565b600080604083850312156123e0576123df6131d0565b5b60006123ee858286016121cf565b92505060206123ff858286016122ae565b9150509250929050565b600080604083850312156124205761241f6131d0565b5b600061242e858286016121cf565b925050602061243f858286016122c3565b9150509250929050565b600080600060608486031215612462576124616131d0565b5b6000612470868287016121cf565b9350506020612481868287016122c3565b925050604061249286828701612284565b9150509250925092565b600080600080608085870312156124b6576124b56131d0565b5b60006124c4878288016121cf565b94505060206124d5878288016122c3565b93505060406124e6878288016122ae565b92505060606124f7878288016122ae565b91505092959194509250565b60008060006060848603121561251c5761251b6131d0565b5b600061252a868287016121cf565b935050602061253b868287016122ae565b925050604061254c868287016121cf565b9150509250925092565b60008060006060848603121561256f5761256e6131d0565b5b600061257d868287016121cf565b935050602061258e868287016122ae565b925050604061259f86828701612230565b9150509250925092565b600061010082840312156125c0576125bf6131d0565b5b60006125ce8482850161220e565b91505092915050565b600080604083850312156125ee576125ed6131d0565b5b60006125fc85828601612245565b925050602061260d85828601612299565b9150509250929050565b60006020828403121561262d5761262c6131d0565b5b600061263b8482850161225a565b91505092915050565b600080600080600060a086880312156126605761265f6131d0565b5b600061266e8882890161225a565b955050602061267f888289016121cf565b9450506040612690888289016122c3565b93505060606126a1888289016122ae565b92505060806126b2888289016122ae565b9150509295509295909350565b60008061012083850312156126d7576126d66131d0565b5b60006126e58582860161225a565b92505060206126f68582860161220e565b9150509250929050565b60008060408385031215612717576127166131d0565b5b60006127258582860161225a565b925050602061273685828601612284565b9150509250929050565b600060208284031215612756576127556131d0565b5b600061276484828501612284565b91505092915050565b600080600080600060a08688031215612789576127886131d0565b5b6000612797888289016122d8565b95505060206127a88882890161226f565b94505060406127b988828901612299565b93505060606127ca88828901612299565b92505060806127db888289016122d8565b9150509295509295909350565b6000602082840312156127fe576127fd6131d0565b5b600061280c848285016122ed565b91505092915050565b60006128218383612a06565b60208301905092915050565b61283681612fe9565b82525050565b61284581612fe9565b82525050565b61285c61285782612fe9565b61311a565b82525050565b8082525050565b61287281612cea565b61287c8184612d0d565b925061288782612ce0565b8060005b838110156128b857815161289f8782612815565b96506128aa83612d00565b92505060018101905061288b565b505050505050565b6128c98161300d565b82525050565b6128d881613019565b82525050565b60006128e982612cf5565b6128f38185612d18565b935061290381856020860161309e565b61290c816131d5565b840191505092915050565b600061292282612cf5565b61292c8185612d29565b935061293c81856020860161309e565b80840191505092915050565b6129518161307a565b82525050565b6000612964600f83612d34565b915061296f82613200565b602082019050919050565b6000612987601583612d34565b915061299282613229565b602082019050919050565b60006129aa600b83612d34565b91506129b582613252565b602082019050919050565b60006129cd601083612d34565b91506129d88261327b565b602082019050919050565b60006129f0601183612d34565b91506129fb826132a4565b602082019050919050565b612a0f8161304d565b82525050565b612a1e8161304d565b82525050565b612a2d81613064565b82525050565b612a3c81613057565b82525050565b6000612a4e8285612917565b9150612a5a828461284b565b6014820191508190509392505050565b6000602082019050612a7f600083018461282d565b92915050565b6000604082019050612a9a6000830185612862565b612aa7602083018461283c565b9392505050565b600061010082019050612ac46000830184612869565b92915050565b6000602082019050612adf60008301846128c0565b92915050565b6000602082019050612afa60008301846128cf565b92915050565b60006020820190508181036000830152612b1a81846128de565b905092915050565b6000602082019050612b376000830184612948565b92915050565b60006020820190508181036000830152612b5681612957565b9050919050565b60006020820190508181036000830152612b768161297a565b9050919050565b60006020820190508181036000830152612b968161299d565b9050919050565b60006020820190508181036000830152612bb6816129c0565b9050919050565b60006020820190508181036000830152612bd6816129e3565b9050919050565b6000602082019050612bf26000830184612a15565b92915050565b6000604082019050612c0d6000830185612a15565b612c1a60208301846128c0565b9392505050565b6000606082019050612c366000830186612a15565b612c436020830185612a15565b612c5060408301846128c0565b949350505050565b6000604082019050612c6d6000830185612a15565b612c7a6020830184612a24565b9392505050565b6000602082019050612c966000830184612a24565b92915050565b6000602082019050612cb16000830184612a33565b92915050565b6000604082019050612ccc6000830185612a33565b612cd960208301846128c0565b9392505050565b6000819050919050565b600060089050919050565b600081519050919050565b6000602082019050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612d5082613064565b9150612d5b83613064565b92508269ffffffffffffffffffff03821115612d7a57612d7961313e565b5b828201905092915050565b6000612d908261304d565b9150612d9b8361304d565b925082612dab57612daa61316d565b5b828204905092915050565b6000808291508390505b6001851115612e0057808604811115612ddc57612ddb61313e565b5b6001851615612deb5780820291505b8081029050612df9856131f3565b9450612dc0565b94509492505050565b6000612e148261304d565b9150612e1f83613057565b9250612e4c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612e54565b905092915050565b600082612e645760019050612f20565b81612e725760009050612f20565b8160018114612e885760028114612e9257612ec1565b6001915050612f20565b60ff841115612ea457612ea361313e565b5b8360020a915084821115612ebb57612eba61313e565b5b50612f20565b5060208310610133831016604e8410600b8410161715612ef65782820a905083811115612ef157612ef061313e565b5b612f20565b612f038484846001612db6565b92509050818404811115612f1a57612f1961313e565b5b81810290505b9392505050565b6000612f328261304d565b9150612f3d8361304d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612f7657612f7561313e565b5b828202905092915050565b6000612f8c8261304d565b9150612f978361304d565b925082821015612faa57612fa961313e565b5b828203905092915050565b6000612fc082613057565b9150612fcb83613057565b925082821015612fde57612fdd61313e565b5b828203905092915050565b6000612ff48261302d565b9050919050565b60006130068261302d565b9050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b60006130858261308c565b9050919050565b60006130978261302d565b9050919050565b60005b838110156130bc5780820151818401526020810190506130a1565b838111156130cb576000848401525b50505050565b60006130dc8261304d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561310f5761310e61313e565b5b600182019050919050565b60006131258261312c565b9050919050565b6000613137826131e6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b60008160011c9050919050565b7f746178206e6f7420636f72726563740000000000000000000000000000000000600082015250565b7f7061796d656e74206973206e6f74206578697374730000000000000000000000600082015250565b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b7f7061796d656e742064697361626c656400000000000000000000000000000000600082015250565b7f7072696365206e6f7420636f7272656374000000000000000000000000000000600082015250565b6132d681612fe9565b81146132e157600080fd5b50565b6132ed81612ffb565b81146132f857600080fd5b50565b6133048161300d565b811461330f57600080fd5b50565b61331b81613019565b811461332657600080fd5b50565b61333281613023565b811461333d57600080fd5b50565b6133498161304d565b811461335457600080fd5b50565b61336081613057565b811461336b57600080fd5b50565b61337781613064565b811461338257600080fd5b5056fea26469706673582212203dc6915e35c862a344417bea3462aba7c252a71cfbc9e1d27fec4da18dcc998364736f6c63430008060033",
}

// DnsCostABI is the input ABI used to generate the binding from.
// Deprecated: Use DnsCostMetaData.ABI instead.
var DnsCostABI = DnsCostMetaData.ABI

// DnsCostBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DnsCostMetaData.Bin instead.
var DnsCostBin = DnsCostMetaData.Bin

// DeployDnsCost deploys a new Ethereum contract, binding an instance of DnsCost to it.
func DeployDnsCost(auth *bind.TransactOpts, backend bind.ContractBackend, top_ common.Address) (common.Address, *types.Transaction, *DnsCost, error) {
	parsed, err := DnsCostMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	libAddressArrayAddr, _, _, _ := DeployLibAddressArray(auth, backend)
	DnsCostBin = strings.ReplaceAll(DnsCostBin, "__$ae7d58307b48f566e6dd130bee0f44ebb6$__", libAddressArrayAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DnsCostBin), backend, top_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DnsCost{DnsCostCaller: DnsCostCaller{contract: contract}, DnsCostTransactor: DnsCostTransactor{contract: contract}, DnsCostFilterer: DnsCostFilterer{contract: contract}}, nil
}

// DnsCost is an auto generated Go binding around an Ethereum contract.
type DnsCost struct {
	DnsCostCaller     // Read-only binding to the contract
	DnsCostTransactor // Write-only binding to the contract
	DnsCostFilterer   // Log filterer for contract events
}

// DnsCostCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnsCostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnsCostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnsCostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnsCostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnsCostSession struct {
	Contract     *DnsCost          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnsCostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnsCostCallerSession struct {
	Contract *DnsCostCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DnsCostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnsCostTransactorSession struct {
	Contract     *DnsCostTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DnsCostRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnsCostRaw struct {
	Contract *DnsCost // Generic contract binding to access the raw methods on
}

// DnsCostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnsCostCallerRaw struct {
	Contract *DnsCostCaller // Generic read-only contract binding to access the raw methods on
}

// DnsCostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnsCostTransactorRaw struct {
	Contract *DnsCostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnsCost creates a new instance of DnsCost, bound to a specific deployed contract.
func NewDnsCost(address common.Address, backend bind.ContractBackend) (*DnsCost, error) {
	contract, err := bindDnsCost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DnsCost{DnsCostCaller: DnsCostCaller{contract: contract}, DnsCostTransactor: DnsCostTransactor{contract: contract}, DnsCostFilterer: DnsCostFilterer{contract: contract}}, nil
}

// NewDnsCostCaller creates a new read-only instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostCaller(address common.Address, caller bind.ContractCaller) (*DnsCostCaller, error) {
	contract, err := bindDnsCost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnsCostCaller{contract: contract}, nil
}

// NewDnsCostTransactor creates a new write-only instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostTransactor(address common.Address, transactor bind.ContractTransactor) (*DnsCostTransactor, error) {
	contract, err := bindDnsCost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnsCostTransactor{contract: contract}, nil
}

// NewDnsCostFilterer creates a new log filterer instance of DnsCost, bound to a specific deployed contract.
func NewDnsCostFilterer(address common.Address, filterer bind.ContractFilterer) (*DnsCostFilterer, error) {
	contract, err := bindDnsCost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnsCostFilterer{contract: contract}, nil
}

// bindDnsCost binds a generic wrapper to an already deployed contract.
func bindDnsCost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnsCostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsCost *DnsCostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsCost.Contract.DnsCostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsCost *DnsCostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsCost.Contract.DnsCostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsCost *DnsCostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsCost.Contract.DnsCostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DnsCost *DnsCostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DnsCost.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DnsCost *DnsCostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DnsCost.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DnsCost *DnsCostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DnsCost.Contract.contract.Transact(opts, method, params...)
}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) SLPrice(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "SLPrice", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) SLPrice(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLPrice(&_DnsCost.CallOpts, arg0, arg1)
}

// SLPrice is a free data retrieval call binding the contract method 0x8412de08.
//
// Solidity: function SLPrice(bytes32 , uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) SLPrice(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLPrice(&_DnsCost.CallOpts, arg0, arg1)
}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) SLTax(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "SLTax", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) SLTax(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLTax(&_DnsCost.CallOpts, arg0)
}

// SLTax is a free data retrieval call binding the contract method 0x9fa001ef.
//
// Solidity: function SLTax(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) SLTax(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.SLTax(&_DnsCost.CallOpts, arg0)
}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCaller) TLPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "TLPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostSession) TLPrice(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.TLPrice(&_DnsCost.CallOpts, arg0)
}

// TLPrice is a free data retrieval call binding the contract method 0x73260cb9.
//
// Solidity: function TLPrice(uint256 ) view returns(uint256)
func (_DnsCost *DnsCostCallerSession) TLPrice(arg0 *big.Int) (*big.Int, error) {
	return _DnsCost.Contract.TLPrice(&_DnsCost.CallOpts, arg0)
}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostCaller) BaseDecimal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "baseDecimal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostSession) BaseDecimal() (uint8, error) {
	return _DnsCost.Contract.BaseDecimal(&_DnsCost.CallOpts)
}

// BaseDecimal is a free data retrieval call binding the contract method 0xdb7b373e.
//
// Solidity: function baseDecimal() view returns(uint8)
func (_DnsCost *DnsCostCallerSession) BaseDecimal() (uint8, error) {
	return _DnsCost.Contract.BaseDecimal(&_DnsCost.CallOpts)
}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostCaller) ChainLinkRateAddr(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "chainLinkRateAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostSession) ChainLinkRateAddr(arg0 common.Address) (common.Address, error) {
	return _DnsCost.Contract.ChainLinkRateAddr(&_DnsCost.CallOpts, arg0)
}

// ChainLinkRateAddr is a free data retrieval call binding the contract method 0x1d869ca1.
//
// Solidity: function chainLinkRateAddr(address ) view returns(address)
func (_DnsCost *DnsCostCallerSession) ChainLinkRateAddr(arg0 common.Address) (common.Address, error) {
	return _DnsCost.Contract.ChainLinkRateAddr(&_DnsCost.CallOpts, arg0)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsCost *DnsCostCaller) DnsTop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "dnsTop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsCost *DnsCostSession) DnsTop() (common.Address, error) {
	return _DnsCost.Contract.DnsTop(&_DnsCost.CallOpts)
}

// DnsTop is a free data retrieval call binding the contract method 0xf3c8bb0c.
//
// Solidity: function dnsTop() view returns(address)
func (_DnsCost *DnsCostCallerSession) DnsTop() (common.Address, error) {
	return _DnsCost.Contract.DnsTop(&_DnsCost.CallOpts)
}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostCaller) GetAllPayments(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllPayments")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostSession) GetAllPayments() ([]byte, error) {
	return _DnsCost.Contract.GetAllPayments(&_DnsCost.CallOpts)
}

// GetAllPayments is a free data retrieval call binding the contract method 0xf8defa3c.
//
// Solidity: function getAllPayments() view returns(bytes)
func (_DnsCost *DnsCostCallerSession) GetAllPayments() ([]byte, error) {
	return _DnsCost.Contract.GetAllPayments(&_DnsCost.CallOpts)
}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllSecondLevelNamePrice(opts *bind.CallOpts, ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllSecondLevelNamePrice", ownerNameHash_)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllSecondLevelNamePrice(ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNamePrice(&_DnsCost.CallOpts, ownerNameHash_)
}

// GetAllSecondLevelNamePrice is a free data retrieval call binding the contract method 0xdabff59a.
//
// Solidity: function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllSecondLevelNamePrice(ownerNameHash_ [32]byte) ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNamePrice(&_DnsCost.CallOpts, ownerNameHash_)
}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllSecondLevelNameTax(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllSecondLevelNameTax")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllSecondLevelNameTax() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNameTax(&_DnsCost.CallOpts)
}

// GetAllSecondLevelNameTax is a free data retrieval call binding the contract method 0x546fa0be.
//
// Solidity: function getAllSecondLevelNameTax() view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllSecondLevelNameTax() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllSecondLevelNameTax(&_DnsCost.CallOpts)
}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostCaller) GetAllTopLevelNamePrice(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getAllTopLevelNamePrice")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostSession) GetAllTopLevelNamePrice() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllTopLevelNamePrice(&_DnsCost.CallOpts)
}

// GetAllTopLevelNamePrice is a free data retrieval call binding the contract method 0xf37e6f17.
//
// Solidity: function getAllTopLevelNamePrice() view returns(uint256[8])
func (_DnsCost *DnsCostCallerSession) GetAllTopLevelNamePrice() ([8]*big.Int, error) {
	return _DnsCost.Contract.GetAllTopLevelNamePrice(&_DnsCost.CallOpts)
}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostCaller) GetExchangeWithId(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getExchangeWithId", erc20Addr_, lastPriceId_, price_)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostSession) GetExchangeWithId(erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetExchangeWithId(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, price_)
}

// GetExchangeWithId is a free data retrieval call binding the contract method 0x79390893.
//
// Solidity: function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) view returns(uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetExchangeWithId(erc20Addr_ common.Address, lastPriceId_ *big.Int, price_ *big.Int) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetExchangeWithId(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, price_)
}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostCaller) GetLatestExchangeValue(opts *bind.CallOpts, erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getLatestExchangeValue", erc20Addr_, price_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostSession) GetLatestExchangeValue(erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	return _DnsCost.Contract.GetLatestExchangeValue(&_DnsCost.CallOpts, erc20Addr_, price_)
}

// GetLatestExchangeValue is a free data retrieval call binding the contract method 0xbba4095c.
//
// Solidity: function getLatestExchangeValue(address erc20Addr_, uint256 price_) view returns(uint256, uint80)
func (_DnsCost *DnsCostCallerSession) GetLatestExchangeValue(erc20Addr_ common.Address, price_ *big.Int) (*big.Int, *big.Int, error) {
	return _DnsCost.Contract.GetLatestExchangeValue(&_DnsCost.CallOpts, erc20Addr_, price_)
}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostCaller) GetPaymentsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getPaymentsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostSession) GetPaymentsCount() (*big.Int, error) {
	return _DnsCost.Contract.GetPaymentsCount(&_DnsCost.CallOpts)
}

// GetPaymentsCount is a free data retrieval call binding the contract method 0xcacf1e0e.
//
// Solidity: function getPaymentsCount() view returns(uint256)
func (_DnsCost *DnsCostCallerSession) GetPaymentsCount() (*big.Int, error) {
	return _DnsCost.Contract.GetPaymentsCount(&_DnsCost.CallOpts)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_DnsCost *DnsCostCaller) GetSecondLevelNamePrice(opts *bind.CallOpts, fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getSecondLevelNamePrice", fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)

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
func (_DnsCost *DnsCostSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _DnsCost.Contract.GetSecondLevelNamePrice(&_DnsCost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetSecondLevelNamePrice is a free data retrieval call binding the contract method 0x9f33b2d8.
//
// Solidity: function getSecondLevelNamePrice(bytes32 fatherHash_, address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetSecondLevelNamePrice(fatherHash_ [32]byte, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, *big.Int, bool, error) {
	return _DnsCost.Contract.GetSecondLevelNamePrice(&_DnsCost.CallOpts, fatherHash_, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_DnsCost *DnsCostCaller) GetTopLevelNamePrice(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "getTopLevelNamePrice", erc20Addr_, lastPriceId_, nameLength_, year_)

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
func (_DnsCost *DnsCostSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetTopLevelNamePrice(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// GetTopLevelNamePrice is a free data retrieval call binding the contract method 0x3dd45f13.
//
// Solidity: function getTopLevelNamePrice(address erc20Addr_, uint80 lastPriceId_, uint8 nameLength_, uint8 year_) view returns(uint256, bool)
func (_DnsCost *DnsCostCallerSession) GetTopLevelNamePrice(erc20Addr_ common.Address, lastPriceId_ *big.Int, nameLength_ uint8, year_ uint8) (*big.Int, bool, error) {
	return _DnsCost.Contract.GetTopLevelNamePrice(&_DnsCost.CallOpts, erc20Addr_, lastPriceId_, nameLength_, year_)
}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostCaller) PaymentIsExist(opts *bind.CallOpts, erc20Addr_ common.Address) (uint8, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentIsExist", erc20Addr_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostSession) PaymentIsExist(erc20Addr_ common.Address) (uint8, error) {
	return _DnsCost.Contract.PaymentIsExist(&_DnsCost.CallOpts, erc20Addr_)
}

// PaymentIsExist is a free data retrieval call binding the contract method 0xd34f4222.
//
// Solidity: function paymentIsExist(address erc20Addr_) view returns(uint8)
func (_DnsCost *DnsCostCallerSession) PaymentIsExist(erc20Addr_ common.Address) (uint8, error) {
	return _DnsCost.Contract.PaymentIsExist(&_DnsCost.CallOpts, erc20Addr_)
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostCaller) PaymentList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostSession) PaymentList(arg0 *big.Int) (common.Address, error) {
	return _DnsCost.Contract.PaymentList(&_DnsCost.CallOpts, arg0)
}

// PaymentList is a free data retrieval call binding the contract method 0xccc79058.
//
// Solidity: function paymentList(uint256 ) view returns(address)
func (_DnsCost *DnsCostCallerSession) PaymentList(arg0 *big.Int) (common.Address, error) {
	return _DnsCost.Contract.PaymentList(&_DnsCost.CallOpts, arg0)
}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostCaller) PaymentMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "paymentMap", arg0)

	outstruct := new(struct {
		Decimal uint8
		Enable  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimal = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Enable = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostSession) PaymentMap(arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	return _DnsCost.Contract.PaymentMap(&_DnsCost.CallOpts, arg0)
}

// PaymentMap is a free data retrieval call binding the contract method 0xf32ec518.
//
// Solidity: function paymentMap(address ) view returns(uint8 decimal, bool enable)
func (_DnsCost *DnsCostCallerSession) PaymentMap(arg0 common.Address) (struct {
	Decimal uint8
	Enable  bool
}, error) {
	return _DnsCost.Contract.PaymentMap(&_DnsCost.CallOpts, arg0)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostCaller) PriceIdIsValid(opts *bind.CallOpts, erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "priceIdIsValid", erc20Addr_, lastPriceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _DnsCost.Contract.PriceIdIsValid(&_DnsCost.CallOpts, erc20Addr_, lastPriceId)
}

// PriceIdIsValid is a free data retrieval call binding the contract method 0x95297fb0.
//
// Solidity: function priceIdIsValid(address erc20Addr_, uint80 lastPriceId) view returns(bool)
func (_DnsCost *DnsCostCallerSession) PriceIdIsValid(erc20Addr_ common.Address, lastPriceId *big.Int) (bool, error) {
	return _DnsCost.Contract.PriceIdIsValid(&_DnsCost.CallOpts, erc20Addr_, lastPriceId)
}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostCaller) UsdtAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DnsCost.contract.Call(opts, &out, "usdtAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostSession) UsdtAddr() (common.Address, error) {
	return _DnsCost.Contract.UsdtAddr(&_DnsCost.CallOpts)
}

// UsdtAddr is a free data retrieval call binding the contract method 0x1554e1ce.
//
// Solidity: function usdtAddr() view returns(address)
func (_DnsCost *DnsCostCallerSession) UsdtAddr() (common.Address, error) {
	return _DnsCost.Contract.UsdtAddr(&_DnsCost.CallOpts)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostTransactor) SetSecondLevelNameTax(opts *bind.TransactOpts, tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "SetSecondLevelNameTax", tax_)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostSession) SetSecondLevelNameTax(tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNameTax(&_DnsCost.TransactOpts, tax_)
}

// SetSecondLevelNameTax is a paid mutator transaction binding the contract method 0x5f3dda02.
//
// Solidity: function SetSecondLevelNameTax(uint256[8] tax_) returns()
func (_DnsCost *DnsCostTransactorSession) SetSecondLevelNameTax(tax_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNameTax(&_DnsCost.TransactOpts, tax_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostTransactor) SetTopLevelNamePrice(opts *bind.TransactOpts, price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "SetTopLevelNamePrice", price_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostSession) SetTopLevelNamePrice(price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetTopLevelNamePrice(&_DnsCost.TransactOpts, price_)
}

// SetTopLevelNamePrice is a paid mutator transaction binding the contract method 0xb84ff547.
//
// Solidity: function SetTopLevelNamePrice(uint256[8] price_) returns()
func (_DnsCost *DnsCostTransactorSession) SetTopLevelNamePrice(price_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetTopLevelNamePrice(&_DnsCost.TransactOpts, price_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3325d51c.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostTransactor) AddPayment(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal_ uint8, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "addPayment", erc20Addr_, decimal_, chainLinkContact_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3325d51c.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.AddPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, chainLinkContact_)
}

// AddPayment is a paid mutator transaction binding the contract method 0x3325d51c.
//
// Solidity: function addPayment(address erc20Addr_, uint8 decimal_, address chainLinkContact_) returns()
func (_DnsCost *DnsCostTransactorSession) AddPayment(erc20Addr_ common.Address, decimal_ uint8, chainLinkContact_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.AddPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, chainLinkContact_)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostTransactor) SetPayment(opts *bind.TransactOpts, erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setPayment", erc20Addr_, decimal_, enable)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostSession) SetPayment(erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.Contract.SetPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, enable)
}

// SetPayment is a paid mutator transaction binding the contract method 0xd3443c62.
//
// Solidity: function setPayment(address erc20Addr_, uint8 decimal_, bool enable) returns()
func (_DnsCost *DnsCostTransactorSession) SetPayment(erc20Addr_ common.Address, decimal_ uint8, enable bool) (*types.Transaction, error) {
	return _DnsCost.Contract.SetPayment(&_DnsCost.TransactOpts, erc20Addr_, decimal_, enable)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostTransactor) SetRelation(opts *bind.TransactOpts, dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setRelation", dnstop_)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostSession) SetRelation(dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.SetRelation(&_DnsCost.TransactOpts, dnstop_)
}

// SetRelation is a paid mutator transaction binding the contract method 0x0868f7b1.
//
// Solidity: function setRelation(address dnstop_) returns()
func (_DnsCost *DnsCostTransactorSession) SetRelation(dnstop_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.SetRelation(&_DnsCost.TransactOpts, dnstop_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostTransactor) SetSecondLevelNamePrice(opts *bind.TransactOpts, ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setSecondLevelNamePrice", ownerNameHash_, prices_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostSession) SetSecondLevelNamePrice(ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNamePrice(&_DnsCost.TransactOpts, ownerNameHash_, prices_)
}

// SetSecondLevelNamePrice is a paid mutator transaction binding the contract method 0x72be125d.
//
// Solidity: function setSecondLevelNamePrice(bytes32 ownerNameHash_, uint256[8] prices_) returns()
func (_DnsCost *DnsCostTransactorSession) SetSecondLevelNamePrice(ownerNameHash_ [32]byte, prices_ [8]*big.Int) (*types.Transaction, error) {
	return _DnsCost.Contract.SetSecondLevelNamePrice(&_DnsCost.TransactOpts, ownerNameHash_, prices_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostTransactor) SetUsdtAddr(opts *bind.TransactOpts, usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "setUsdtAddr", usdtTokenAddr_, baseDecimal_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostSession) SetUsdtAddr(usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.SetUsdtAddr(&_DnsCost.TransactOpts, usdtTokenAddr_, baseDecimal_)
}

// SetUsdtAddr is a paid mutator transaction binding the contract method 0xcfdbee11.
//
// Solidity: function setUsdtAddr(address usdtTokenAddr_, uint8 baseDecimal_) returns()
func (_DnsCost *DnsCostTransactorSession) SetUsdtAddr(usdtTokenAddr_ common.Address, baseDecimal_ uint8) (*types.Transaction, error) {
	return _DnsCost.Contract.SetUsdtAddr(&_DnsCost.TransactOpts, usdtTokenAddr_, baseDecimal_)
}

// Erc721OwnerMetaData contains all meta data concerning the Erc721Owner contract.
var Erc721OwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721Owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"erc721Owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"erc721Owner_\",\"type\":\"address\"}],\"name\":\"transferErc721Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516105f93803806105f9833981810160405281019061003291906100cf565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b6000815190506100c981610146565b92915050565b600080604083850312156100e6576100e5610141565b5b60006100f4858286016100ba565b9250506020610105858286016100ba565b9150509250929050565b600061011a82610121565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b61014f8161010f565b811461015a57600080fd5b50565b61048d8061016c6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806332cdee7b146100515780638da5cb5b1461006f578063c93f5e261461008d578063f2fde38b146100a9575b600080fd5b6100596100c5565b604051610066919061034b565b60405180910390f35b6100776100e9565b604051610084919061034b565b60405180910390f35b6100a760048036038101906100a291906102c9565b61010f565b005b6100c360048036038101906100be91906102c9565b6101e2565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461019f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019690610366565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610270576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026790610386565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000813590506102c381610440565b92915050565b6000602082840312156102df576102de6103e9565b5b60006102ed848285016102b4565b91505092915050565b6102ff816103b7565b82525050565b6000610312600d836103a6565b915061031d826103ee565b602082019050919050565b6000610335600b836103a6565b915061034082610417565b602082019050919050565b600060208201905061036060008301846102f6565b92915050565b6000602082019050818103600083015261037f81610305565b9050919050565b6000602082019050818103600083015261039f81610328565b9050919050565b600082825260208201905092915050565b60006103c2826103c9565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b7f6e6f7420746865206f776e657200000000000000000000000000000000000000600082015250565b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b610449816103b7565b811461045457600080fd5b5056fea264697066735822122070e40e7296b29e7ecc7b1d1664dc66ebc2117e6fa0da6d16498637c51748c51d64736f6c63430008060033",
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

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_Erc721Owner *Erc721OwnerTransactor) TransferErc721Owner(opts *bind.TransactOpts, erc721Owner_ common.Address) (*types.Transaction, error) {
	return _Erc721Owner.contract.Transact(opts, "transferErc721Owner", erc721Owner_)
}

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_Erc721Owner *Erc721OwnerSession) TransferErc721Owner(erc721Owner_ common.Address) (*types.Transaction, error) {
	return _Erc721Owner.Contract.TransferErc721Owner(&_Erc721Owner.TransactOpts, erc721Owner_)
}

// TransferErc721Owner is a paid mutator transaction binding the contract method 0xc93f5e26.
//
// Solidity: function transferErc721Owner(address erc721Owner_) returns()
func (_Erc721Owner *Erc721OwnerTransactorSession) TransferErc721Owner(erc721Owner_ common.Address) (*types.Transaction, error) {
	return _Erc721Owner.Contract.TransferErc721Owner(&_Erc721Owner.TransactOpts, erc721Owner_)
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
