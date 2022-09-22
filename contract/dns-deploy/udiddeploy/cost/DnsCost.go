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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"top_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SLTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"tax_\",\"type\":\"uint256[8]\"}],\"name\":\"SetSecondLevelNameTax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"price_\",\"type\":\"uint256[8]\"}],\"name\":\"SetTopLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TLPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"chainLinkContact_\",\"type\":\"address\"}],\"name\":\"addPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseDecimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chainLinkRateAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"delPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPayments\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"}],\"name\":\"getAllSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSecondLevelNameTax\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getExchangeWithId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"getLatestExchangeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fatherHash_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getSecondLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId_\",\"type\":\"uint80\"},{\"internalType\":\"uint8\",\"name\":\"nameLength_\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"year_\",\"type\":\"uint8\"}],\"name\":\"getTopLevelNamePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"}],\"name\":\"paymentIsExist\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentMap\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimal\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"lastPriceId\",\"type\":\"uint80\"}],\"name\":\"priceIdIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimal_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dnstop_\",\"type\":\"address\"}],\"name\":\"setRelation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownerNameHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"prices_\",\"type\":\"uint256[8]\"}],\"name\":\"setSecondLevelNamePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdtTokenAddr_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"baseDecimal_\",\"type\":\"uint8\"}],\"name\":\"setUsdtAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003d8238038062003d8283398181016040528101906200003791906200078c565b60005b6008811015620000805768a2a15d09519be000006009826008811062000065576200006462000911565b5b01819055508080620000779062000894565b9150506200003a565b5060005b6008811015620000c957674563918244f4000060018260088110620000ae57620000ad62000911565b5b01819055508080620000c09062000894565b91505062000084565b50734f4f07917e13785bfd55cd3485b254bde6f092656000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506006600060146101000a81548160ff021916908360ff1602179055506040518060400160405280601260ff16815260200160011515815250601260008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d8909160006040518363ffffffff1660e01b8152600401620002189291906200081d565b604080518083038186803b1580156200023057600080fd5b505af415801562000245573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200026b9190620007be565b50506040518060400160405280601260ff168152602001600115158152506012600073571787b5e033bf8bb2a1e5930265828ef3ffca0073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d8909173571787b5e033bf8bb2a1e5930265828ef3ffca006040518363ffffffff1660e01b8152600401620003729291906200081d565b604080518083038186803b1580156200038a57600080fd5b505af41580156200039f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003c59190620007be565b50506040518060400160405280600660ff16815260200160011515815250601260007365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d890917365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e6040518363ffffffff1660e01b8152600401620004cc9291906200081d565b604080518083038186803b158015620004e457600080fd5b505af4158015620004f9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200051f9190620007be565b505080601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073e0005488e0d9be7b24cce51f722d24749391327f601460008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073136a734040f258054da7ebf42bd25b73e913839a6014600073571787b5e033bf8bb2a1e5930265828ef3ffca0073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550738857aaea5c663a4408963e8d964d22e438d596f2601460007365eedd3b0b0a7c5cfbe97b22bbf19cd492f2237e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000993565b600081519050620007588162000945565b92915050565b6000815190506200076f816200095f565b92915050565b600081519050620007868162000979565b92915050565b600060208284031215620007a557620007a462000940565b5b6000620007b58482850162000747565b91505092915050565b60008060408385031215620007d857620007d762000940565b5b6000620007e8858286016200075e565b9250506020620007fb8582860162000775565b9150509250929050565b62000810816200084a565b82525050565b8082525050565b600060408201905062000834600083018562000816565b62000843602083018462000805565b9392505050565b600062000857826200086a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000620008a1826200088a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620008d757620008d6620008e2565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b62000950816200084a565b81146200095c57600080fd5b50565b6200096a816200085e565b81146200097657600080fd5b50565b62000984816200088a565b81146200099057600080fd5b50565b6133df80620009a36000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80639f33b2d8116100f9578063d3443c6211610097578063db7b373e11610071578063db7b373e1461052d578063f32ec5181461054b578063f37e6f171461057c578063f8defa3c1461059a576101a9565b8063d3443c62146104b1578063d34f4222146104cd578063dabff59a146104fd576101a9565b8063bba4095c116100d3578063bba4095c14610416578063cacf1e0e14610447578063ccc7905814610465578063cfdbee1114610495576101a9565b80639f33b2d8146103985780639fa001ef146103ca578063b84ff547146103fa576101a9565b8063546fa0be1161016657806373260cb91161014057806373260cb9146102d757806379390893146103075780638412de081461033857806395297fb014610368576101a9565b8063546fa0be146102815780635f3dda021461029f57806372be125d146102bb576101a9565b80630868f7b1146101ae5780631554e1ce146101ca5780631d869ca1146101e85780633325d51c146102185780633dd45f13146102345780633ed2e3e814610265575b600080fd5b6101c860048036038101906101c391906123b3565b6105b8565b005b6101d2610709565b6040516101df9190612b16565b60405180910390f35b61020260048036038101906101fd91906123b3565b61072d565b60405161020f9190612b16565b60405180910390f35b610232600480360381019061022d91906125b4565b610760565b005b61024e6004803603810190610249919061254d565b610a1b565b60405161025c929190612c69565b60405180910390f35b61027f600480360381019061027a91906123b3565b610acc565b005b610289610cd3565b6040516102969190612b5a565b60405180910390f35b6102b960048036038101906102b4919061265a565b610d1e565b005b6102d560048036038101906102d0919061279d565b610e7f565b005b6102f160048036038101906102ec919061281e565b611089565b6040516102fe9190612c4e565b60405180910390f35b610321600480360381019061031c91906124fa565b6110a4565b60405161032f929190612c69565b60405180910390f35b610352600480360381019061034d91906127de565b6110be565b60405161035f9190612c4e565b60405180910390f35b610382600480360381019061037d91906124ba565b6110e6565b60405161038f9190612b76565b60405180910390f35b6103b260048036038101906103ad9190612722565b611296565b6040516103c193929190612c92565b60405180910390f35b6103e460048036038101906103df919061281e565b611410565b6040516103f19190612c4e565b60405180910390f35b610414600480360381019061040f919061265a565b61142b565b005b610430600480360381019061042b919061243a565b61158c565b60405161043e929190612cc9565b60405180910390f35b61044f611798565b60405161045c9190612c4e565b60405180910390f35b61047f600480360381019061047a919061281e565b6117a5565b60405161048c9190612b16565b60405180910390f35b6104af60048036038101906104aa919061247a565b6117e4565b005b6104cb60048036038101906104c69190612607565b611950565b005b6104e760048036038101906104e291906123b3565b611b94565b6040516104f49190612d0d565b60405180910390f35b610517600480360381019061051291906126f5565b611bed565b6040516105249190612b5a565b60405180910390f35b610535611c4b565b6040516105429190612d0d565b60405180910390f35b610565600480360381019061056091906123b3565b611c5e565b604051610573929190612d28565b60405180910390f35b610584611c9c565b6040516105919190612b5a565b60405180910390f35b6105a2611ce7565b6040516105af9190612bac565b60405180910390f35b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062057600080fd5b505afa158015610634573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610658919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bc90612c0e565b60405180910390fd5b80601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60146020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107c857600080fd5b505afa1580156107dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610800919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461086d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086490612c0e565b60405180910390fd5b60405180604001604052808360ff16815260200160011515815250601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__6310b142d89091856040518363ffffffff1660e01b8152600401610947929190612b31565b604080518083038186803b15801561095e57600080fd5b505af4158015610972573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099691906126b5565b505080601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008060008360ff16610a2d86611d77565b610a379190612f98565b905060008669ffffffffffffffffffff16148015610aa0575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16145b15610ab2578060019250925050610ac3565b610abd878783611db5565b92509250505b94509492505050565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610b3457600080fd5b505afa158015610b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6c919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bd9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd090612c0e565b60405180910390fd5b601260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549060ff02191690556000820160016101000a81549060ff02191690555050601373__$ae7d58307b48f566e6dd130bee0f44ebb6$__63ec3770099091836040518363ffffffff1660e01b8152600401610c7f929190612b31565b60206040518083038186803b158015610c9757600080fd5b505af4158015610cab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ccf9190612688565b5050565b610cdb61225d565b6001600880602002604051908101604052809291908260088015610d14576020028201915b815481526020019060010190808311610d00575b5050505050905090565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d8657600080fd5b505afa158015610d9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbe919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2290612c0e565b60405180910390fd5b60005b6008811015610e7b57818160088110610e4a57610e496131e9565b5b602002013560018260088110610e6357610e626131e9565b5b01819055508080610e739061311e565b915050610e2e565b5050565b6000801b8214158015610fe457503373ffffffffffffffffffffffffffffffffffffffff16601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632c19be76846040518263ffffffff1660e01b8152600401610eff9190612b91565b60206040518083038186803b158015610f1757600080fd5b505afa158015610f2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4f91906123e0565b73ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f9457600080fd5b505afa158015610fa8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fcc919061240d565b73ffffffffffffffffffffffffffffffffffffffff16145b611023576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101a90612c0e565b60405180910390fd5b60005b600881101561108457818160088110611042576110416131e9565b5b602002013560116000858152602001908152602001600020826008811061106c5761106b6131e9565b5b0181905550808061107c9061311e565b915050611026565b505050565b6009816008811061109957600080fd5b016000915090505481565b6000806110b2858585611db5565b91509150935093915050565b601160205281600052604060002081600881106110da57600080fd5b01600091509150505481565b600080601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561119357600080fd5b505afa1580156111a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111cb919061284b565b50505050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141561123057600192505050611290565b8069ffffffffffffffffffff168469ffffffffffffffffffff16148061127957508069ffffffffffffffffffff1660018561126b9190612db6565b69ffffffffffffffffffff16145b1561128957600192505050611290565b6000925050505b92915050565b6000806000808460ff166112ad8a8860ff16612130565b6112b79190612f98565b905060008560ff166112c8886121a2565b6112d29190612f98565b905060008869ffffffffffffffffffff1614801561133b575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff16145b1561135157808260019450945094505050611405565b60008061135f8b8b86611db5565b91509150806113a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139a90612c2e565b60405180910390fd5b6000806113b18d8d87611db5565b91509150806113f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ec90612bce565b60405180910390fd5b8184849850985098505050505050505b955095509592505050565b6001816008811061142057600080fd5b016000915090505481565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561149357600080fd5b505afa1580156114a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114cb919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611538576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152f90612c0e565b60405180910390fd5b60005b600881101561158857818160088110611557576115566131e9565b5b6020020135600982600881106115705761156f6131e9565b5b018190555080806115809061311e565b91505061153b565b5050565b6000806000601460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561163c57600080fd5b505afa158015611650573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611674919061284b565b505050915091506000818473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156116c457600080fd5b505afa1580156116d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116fc91906128c6565b600a6117089190612e7a565b886117139190612f98565b61171d9190612df6565b9050611788600060149054906101000a900460ff16601260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16836121e0565b8395509550505050509250929050565b6000601380549050905090565b601381815481106117b557600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561184c57600080fd5b505afa158015611860573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611884919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146118f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118e890612c0e565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060146101000a81548160ff021916908360ff1602179055505050565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156119b857600080fd5b505afa1580156119cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119f0919061240d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a5490612c0e565b60405180910390fd5b6000601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1660ff1611611af2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ae990612bee565b60405180910390fd5b60405180604001604052808360ff168152602001821515815250601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908315150217905550905050505050565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff169050919050565b611bf561225d565b60116000838152602001908152602001600020600880602002604051908101604052809291908260088015611c3f576020028201915b815481526020019060010190808311611c2b575b50505050509050919050565b600060149054906101000a900460ff1681565b60126020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900460ff16905082565b611ca461225d565b6009600880602002604051908101604052809291908260088015611cdd576020028201915b815481526020019060010190808311611cc9575b5050505050905090565b60608060005b601380549050811015611d6f578160138281548110611d0f57611d0e6131e9565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001611d4b929190612aee565b60405160208183030381529060405291508080611d679061311e565b915050611ced565b508091505090565b600060088260ff1610611d955760016008611d929190613026565b91505b60098260ff1660088110611dac57611dab6131e9565b5b01549050919050565b6000806000601460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000808273ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b158015611e6557600080fd5b505afa158015611e79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e9d919061284b565b505050915091506000818473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015611eed57600080fd5b505afa158015611f01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2591906128c6565b600a611f319190612e7a565b88611f3c9190612f98565b611f469190612df6565b90508769ffffffffffffffffffff168369ffffffffffffffffffff161415611fe057611fd1600060149054906101000a900460ff16601260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16836121e0565b60019550955050505050612128565b600188611fed9190612db6565b69ffffffffffffffffffff168369ffffffffffffffffffff16141561211c578373ffffffffffffffffffffffffffffffffffffffff16639a6fc8f5896040518263ffffffff1660e01b81526004016120459190612cf2565b60a06040518083038186803b15801561205d57600080fd5b505afa158015612071573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612095919061284b565b90919293509091509050508092505061210d600060149054906101000a900460ff16601260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16836121e0565b60019550955050505050612128565b60008095509550505050505b935093915050565b6000601160008481526020019081526020016000205060088210612173576001601160008581526020019081526020016000205060086121709190612ff2565b91505b601160008481526020019081526020016000208260088110612198576121976131e9565b5b0154905092915050565b600060088260ff16106121c057600160086121bd9190612ff2565b91505b60018260ff16600881106121d7576121d66131e9565b5b01549050919050565b60008260ff168460ff1614156121f857819050612256565b8260ff168460ff1611156122305782846122129190613026565b600a61221e9190612e7a565b826122299190612df6565b9050612256565b838361223c9190613026565b600a6122489190612e7a565b826122539190612f98565b90505b9392505050565b604051806101000160405280600890602082028036833780820191505090505090565b60008135905061228f816132f1565b92915050565b6000815190506122a4816132f1565b92915050565b6000815190506122b981613308565b92915050565b6000819050826020600802820111156122db576122da613218565b5b92915050565b6000813590506122f08161331f565b92915050565b6000815190506123058161331f565b92915050565b60008135905061231a81613336565b92915050565b60008151905061232f8161334d565b92915050565b60008135905061234481613364565b92915050565b60008151905061235981613364565b92915050565b60008135905061236e8161337b565b92915050565b60008135905061238381613392565b92915050565b60008151905061239881613392565b92915050565b6000815190506123ad8161337b565b92915050565b6000602082840312156123c9576123c861321d565b5b60006123d784828501612280565b91505092915050565b6000602082840312156123f6576123f561321d565b5b600061240484828501612295565b91505092915050565b6000602082840312156124235761242261321d565b5b6000612431848285016122aa565b91505092915050565b600080604083850312156124515761245061321d565b5b600061245f85828601612280565b925050602061247085828601612335565b9150509250929050565b600080604083850312156124915761249061321d565b5b600061249f85828601612280565b92505060206124b08582860161235f565b9150509250929050565b600080604083850312156124d1576124d061321d565b5b60006124df85828601612280565b92505060206124f085828601612374565b9150509250929050565b6000806000606084860312156125135761251261321d565b5b600061252186828701612280565b935050602061253286828701612374565b925050604061254386828701612335565b9150509250925092565b600080600080608085870312156125675761256661321d565b5b600061257587828801612280565b945050602061258687828801612374565b93505060406125978782880161235f565b92505060606125a88782880161235f565b91505092959194509250565b6000806000606084860312156125cd576125cc61321d565b5b60006125db86828701612280565b93505060206125ec8682870161235f565b92505060406125fd86828701612280565b9150509250925092565b6000806000606084860312156126205761261f61321d565b5b600061262e86828701612280565b935050602061263f8682870161235f565b9250506040612650868287016122e1565b9150509250925092565b600061010082840312156126715761267061321d565b5b600061267f848285016122bf565b91505092915050565b60006020828403121561269e5761269d61321d565b5b60006126ac848285016122f6565b91505092915050565b600080604083850312156126cc576126cb61321d565b5b60006126da858286016122f6565b92505060206126eb8582860161234a565b9150509250929050565b60006020828403121561270b5761270a61321d565b5b60006127198482850161230b565b91505092915050565b600080600080600060a0868803121561273e5761273d61321d565b5b600061274c8882890161230b565b955050602061275d88828901612280565b945050604061276e88828901612374565b935050606061277f8882890161235f565b92505060806127908882890161235f565b9150509295509295909350565b60008061012083850312156127b5576127b461321d565b5b60006127c38582860161230b565b92505060206127d4858286016122bf565b9150509250929050565b600080604083850312156127f5576127f461321d565b5b60006128038582860161230b565b925050602061281485828601612335565b9150509250929050565b6000602082840312156128345761283361321d565b5b600061284284828501612335565b91505092915050565b600080600080600060a086880312156128675761286661321d565b5b600061287588828901612389565b955050602061288688828901612320565b94505060406128978882890161234a565b93505060606128a88882890161234a565b92505060806128b988828901612389565b9150509295509295909350565b6000602082840312156128dc576128db61321d565b5b60006128ea8482850161239e565b91505092915050565b60006128ff8383612ab2565b60208301905092915050565b6129148161305a565b82525050565b6129238161305a565b82525050565b61293a6129358261305a565b613167565b82525050565b8082525050565b61295081612d5b565b61295a8184612d7e565b925061296582612d51565b8060005b8381101561299657815161297d87826128f3565b965061298883612d71565b925050600181019050612969565b505050505050565b6129a78161307e565b82525050565b6129b68161308a565b82525050565b60006129c782612d66565b6129d18185612d89565b93506129e18185602086016130eb565b6129ea81613222565b840191505092915050565b6000612a0082612d66565b612a0a8185612d9a565b9350612a1a8185602086016130eb565b80840191505092915050565b6000612a33600f83612da5565b9150612a3e8261324d565b602082019050919050565b6000612a56601583612da5565b9150612a6182613276565b602082019050919050565b6000612a79600b83612da5565b9150612a848261329f565b602082019050919050565b6000612a9c601183612da5565b9150612aa7826132c8565b602082019050919050565b612abb816130be565b82525050565b612aca816130be565b82525050565b612ad9816130d5565b82525050565b612ae8816130c8565b82525050565b6000612afa82856129f5565b9150612b068284612929565b6014820191508190509392505050565b6000602082019050612b2b600083018461290b565b92915050565b6000604082019050612b466000830185612940565b612b53602083018461291a565b9392505050565b600061010082019050612b706000830184612947565b92915050565b6000602082019050612b8b600083018461299e565b92915050565b6000602082019050612ba660008301846129ad565b92915050565b60006020820190508181036000830152612bc681846129bc565b905092915050565b60006020820190508181036000830152612be781612a26565b9050919050565b60006020820190508181036000830152612c0781612a49565b9050919050565b60006020820190508181036000830152612c2781612a6c565b9050919050565b60006020820190508181036000830152612c4781612a8f565b9050919050565b6000602082019050612c636000830184612ac1565b92915050565b6000604082019050612c7e6000830185612ac1565b612c8b602083018461299e565b9392505050565b6000606082019050612ca76000830186612ac1565b612cb46020830185612ac1565b612cc1604083018461299e565b949350505050565b6000604082019050612cde6000830185612ac1565b612ceb6020830184612ad0565b9392505050565b6000602082019050612d076000830184612ad0565b92915050565b6000602082019050612d226000830184612adf565b92915050565b6000604082019050612d3d6000830185612adf565b612d4a602083018461299e565b9392505050565b6000819050919050565b600060089050919050565b600081519050919050565b6000602082019050919050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612dc1826130d5565b9150612dcc836130d5565b92508269ffffffffffffffffffff03821115612deb57612dea61318b565b5b828201905092915050565b6000612e01826130be565b9150612e0c836130be565b925082612e1c57612e1b6131ba565b5b828204905092915050565b6000808291508390505b6001851115612e7157808604811115612e4d57612e4c61318b565b5b6001851615612e5c5780820291505b8081029050612e6a85613240565b9450612e31565b94509492505050565b6000612e85826130be565b9150612e90836130c8565b9250612ebd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612ec5565b905092915050565b600082612ed55760019050612f91565b81612ee35760009050612f91565b8160018114612ef95760028114612f0357612f32565b6001915050612f91565b60ff841115612f1557612f1461318b565b5b8360020a915084821115612f2c57612f2b61318b565b5b50612f91565b5060208310610133831016604e8410600b8410161715612f675782820a905083811115612f6257612f6161318b565b5b612f91565b612f748484846001612e27565b92509050818404811115612f8b57612f8a61318b565b5b81810290505b9392505050565b6000612fa3826130be565b9150612fae836130be565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612fe757612fe661318b565b5b828202905092915050565b6000612ffd826130be565b9150613008836130be565b92508282101561301b5761301a61318b565b5b828203905092915050565b6000613031826130c8565b915061303c836130c8565b92508282101561304f5761304e61318b565b5b828203905092915050565b60006130658261309e565b9050919050565b60006130778261309e565b9050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600069ffffffffffffffffffff82169050919050565b60005b838110156131095780820151818401526020810190506130ee565b83811115613118576000848401525b50505050565b6000613129826130be565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561315c5761315b61318b565b5b600182019050919050565b600061317282613179565b9050919050565b600061318482613233565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b60008160011c9050919050565b7f746178206e6f7420636f72726563740000000000000000000000000000000000600082015250565b7f7061796d656e74206973206e6f74206578697374730000000000000000000000600082015250565b7f6e6f7420616c6c6f776564000000000000000000000000000000000000000000600082015250565b7f7072696365206e6f7420636f7272656374000000000000000000000000000000600082015250565b6132fa8161305a565b811461330557600080fd5b50565b6133118161306c565b811461331c57600080fd5b50565b6133288161307e565b811461333357600080fd5b50565b61333f8161308a565b811461334a57600080fd5b50565b61335681613094565b811461336157600080fd5b50565b61336d816130be565b811461337857600080fd5b50565b613384816130c8565b811461338f57600080fd5b50565b61339b816130d5565b81146133a657600080fd5b5056fea2646970667358221220ad2194e972c0427858cba6399636146235d01897535cca99e31456b30566efc464736f6c63430008060033",
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

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostTransactor) DelPayment(opts *bind.TransactOpts, erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.contract.Transact(opts, "delPayment", erc20Addr_)
}

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostSession) DelPayment(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.DelPayment(&_DnsCost.TransactOpts, erc20Addr_)
}

// DelPayment is a paid mutator transaction binding the contract method 0x3ed2e3e8.
//
// Solidity: function delPayment(address erc20Addr_) returns()
func (_DnsCost *DnsCostTransactorSession) DelPayment(erc20Addr_ common.Address) (*types.Transaction, error) {
	return _DnsCost.Contract.DelPayment(&_DnsCost.TransactOpts, erc20Addr_)
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
