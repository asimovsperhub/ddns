// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlink

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

// FakeSolMetaData contains all meta data concerning the FakeSol contract.
var FakeSolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200004733601260ff16600a6200002a9190620002ae565b631908b1006200003b9190620003eb565b6200004d60201b60201c565b62000492565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200008857600080fd5b620000a481600254620001b060201b620010011790919060201c565b60028190555062000102816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620001b060201b620010011790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620001a49190620001d9565b60405180910390a35050565b60008183620001c09190620001f6565b905092915050565b620001d3816200044c565b82525050565b6000602082019050620001f06000830184620001c8565b92915050565b600062000203826200044c565b915062000210836200044c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000248576200024762000456565b5b828201905092915050565b6000808291508390505b6001851115620002a5578086048111156200027d576200027c62000456565b5b60018516156200028d5780820291505b80810290506200029d8562000485565b94506200025d565b94509492505050565b6000620002bb826200044c565b9150620002c8836200044c565b9250620002f77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484620002ff565b905092915050565b600082620003115760019050620003e4565b81620003215760009050620003e4565b81600181146200033a576002811462000345576200037b565b6001915050620003e4565b60ff8411156200035a576200035962000456565b5b8360020a91508482111562000374576200037362000456565b5b50620003e4565b5060208310610133831016604e8410600b8410161715620003b55782820a905083811115620003af57620003ae62000456565b5b620003e4565b620003c4848484600162000253565b92509050818404811115620003de57620003dd62000456565b5b81810290505b9392505050565b6000620003f8826200044c565b915062000405836200044c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000441576200044062000456565b5b828202905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b61194f80620004a26000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806342966c681161008c57806395d89b411161006657806395d89b411461025f578063a457c2d71461027d578063a9059cbb146102ad578063dd62ed3e146102dd576100ea565b806342966c68146101f757806370a082311461021357806379cc679014610243576100ea565b806323b872dd116100c857806323b872dd1461015b5780632ff2e9dc1461018b578063313ce567146101a957806339509351146101c7576100ea565b806306fdde03146100ef578063095ea7b31461010d57806318160ddd1461013d575b600080fd5b6100f761030d565b6040516101049190611548565b60405180910390f35b6101276004803603810190610122919061145a565b610346565b604051610134919061152d565b60405180910390f35b610145610471565b604051610152919061156a565b60405180910390f35b61017560048036038101906101709190611407565b61047b565b604051610182919061152d565b60405180910390f35b610193610830565b6040516101a0919061156a565b60405180910390f35b6101b1610853565b6040516101be9190611585565b60405180910390f35b6101e160048036038101906101dc919061145a565b610858565b6040516101ee919061152d565b60405180910390f35b610211600480360381019061020c919061149a565b610a8d565b005b61022d6004803603810190610228919061139a565b610a9a565b60405161023a919061156a565b60405180910390f35b61025d6004803603810190610258919061145a565b610ae2565b005b610267610af0565b6040516102749190611548565b60405180910390f35b6102976004803603810190610292919061145a565b610b29565b6040516102a4919061152d565b60405180910390f35b6102c760048036038101906102c2919061145a565b610d5e565b6040516102d4919061152d565b60405180910390f35b6102f760048036038101906102f291906113c7565b610f7a565b604051610304919061156a565b60405180910390f35b6040518060400160405280600881526020017f46616b6520534f4c00000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561038157600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161045f919061156a565b60405180910390a36001905092915050565b6000600254905090565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211156104c857600080fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561055157600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561058b57600080fd5b6105dc826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061066f826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100190919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061074082600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161081d919061156a565b60405180910390a3600190509392505050565b601260ff16600a6108419190611665565b631908b1006108509190611783565b81565b601281565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561089357600080fd5b61092282600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100190919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051610a7b919061156a565b60405180910390a36001905092915050565b610a97338261102d565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610aec82826111ca565b5050565b6040518060400160405280600481526020017f46534f4c0000000000000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610b6457600080fd5b610bf382600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051610d4c919061156a565b60405180910390a36001905092915050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115610dab57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610de557600080fd5b610e36826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610ec9826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461100190919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f68919061156a565b60405180910390a36001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000818361100f91906115bc565b905092915050565b6000818361102591906117dd565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561106757600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111156110b257600080fd5b6110c78160025461101790919063ffffffff16565b60028190555061111e816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111be919061156a565b60405180910390a35050565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481111561125357600080fd5b6112e281600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461101790919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061136c828261102d565b5050565b60008135905061137f816118eb565b92915050565b60008135905061139481611902565b92915050565b6000602082840312156113b0576113af6118c8565b5b60006113be84828501611370565b91505092915050565b600080604083850312156113de576113dd6118c8565b5b60006113ec85828601611370565b92505060206113fd85828601611370565b9150509250929050565b6000806000606084860312156114205761141f6118c8565b5b600061142e86828701611370565b935050602061143f86828701611370565b925050604061145086828701611385565b9150509250925092565b60008060408385031215611471576114706118c8565b5b600061147f85828601611370565b925050602061149085828601611385565b9150509250929050565b6000602082840312156114b0576114af6118c8565b5b60006114be84828501611385565b91505092915050565b6114d081611823565b82525050565b60006114e1826115a0565b6114eb81856115ab565b93506114fb818560208601611866565b611504816118cd565b840191505092915050565b6115188161184f565b82525050565b61152781611859565b82525050565b600060208201905061154260008301846114c7565b92915050565b6000602082019050818103600083015261156281846114d6565b905092915050565b600060208201905061157f600083018461150f565b92915050565b600060208201905061159a600083018461151e565b92915050565b600081519050919050565b600082825260208201905092915050565b60006115c78261184f565b91506115d28361184f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561160757611606611899565b5b828201905092915050565b6000808291508390505b600185111561165c5780860481111561163857611637611899565b5b60018516156116475780820291505b8081029050611655856118de565b945061161c565b94509492505050565b60006116708261184f565b915061167b8361184f565b92506116a87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846116b0565b905092915050565b6000826116c0576001905061177c565b816116ce576000905061177c565b81600181146116e457600281146116ee5761171d565b600191505061177c565b60ff841115611700576116ff611899565b5b8360020a91508482111561171757611716611899565b5b5061177c565b5060208310610133831016604e8410600b84101617156117525782820a90508381111561174d5761174c611899565b5b61177c565b61175f8484846001611612565b9250905081840481111561177657611775611899565b5b81810290505b9392505050565b600061178e8261184f565b91506117998361184f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156117d2576117d1611899565b5b828202905092915050565b60006117e88261184f565b91506117f38361184f565b92508282101561180657611805611899565b5b828203905092915050565b600061181c8261182f565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611884578082015181840152602081019050611869565b83811115611893576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b6118f481611811565b81146118ff57600080fd5b50565b61190b8161184f565b811461191657600080fd5b5056fea26469706673582212209b7fc0c3f7daa75427913dfa2f012237069d061e9dfdde9daf566e5e167da32764736f6c63430008060033",
}

// FakeSolABI is the input ABI used to generate the binding from.
// Deprecated: Use FakeSolMetaData.ABI instead.
var FakeSolABI = FakeSolMetaData.ABI

// FakeSolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FakeSolMetaData.Bin instead.
var FakeSolBin = FakeSolMetaData.Bin

// DeployFakeSol deploys a new Ethereum contract, binding an instance of FakeSol to it.
func DeployFakeSol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FakeSol, error) {
	parsed, err := FakeSolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FakeSolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FakeSol{FakeSolCaller: FakeSolCaller{contract: contract}, FakeSolTransactor: FakeSolTransactor{contract: contract}, FakeSolFilterer: FakeSolFilterer{contract: contract}}, nil
}

// FakeSol is an auto generated Go binding around an Ethereum contract.
type FakeSol struct {
	FakeSolCaller     // Read-only binding to the contract
	FakeSolTransactor // Write-only binding to the contract
	FakeSolFilterer   // Log filterer for contract events
}

// FakeSolCaller is an auto generated read-only Go binding around an Ethereum contract.
type FakeSolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeSolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FakeSolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeSolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FakeSolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FakeSolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FakeSolSession struct {
	Contract     *FakeSol          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FakeSolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FakeSolCallerSession struct {
	Contract *FakeSolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FakeSolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FakeSolTransactorSession struct {
	Contract     *FakeSolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FakeSolRaw is an auto generated low-level Go binding around an Ethereum contract.
type FakeSolRaw struct {
	Contract *FakeSol // Generic contract binding to access the raw methods on
}

// FakeSolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FakeSolCallerRaw struct {
	Contract *FakeSolCaller // Generic read-only contract binding to access the raw methods on
}

// FakeSolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FakeSolTransactorRaw struct {
	Contract *FakeSolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFakeSol creates a new instance of FakeSol, bound to a specific deployed contract.
func NewFakeSol(address common.Address, backend bind.ContractBackend) (*FakeSol, error) {
	contract, err := bindFakeSol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FakeSol{FakeSolCaller: FakeSolCaller{contract: contract}, FakeSolTransactor: FakeSolTransactor{contract: contract}, FakeSolFilterer: FakeSolFilterer{contract: contract}}, nil
}

// NewFakeSolCaller creates a new read-only instance of FakeSol, bound to a specific deployed contract.
func NewFakeSolCaller(address common.Address, caller bind.ContractCaller) (*FakeSolCaller, error) {
	contract, err := bindFakeSol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FakeSolCaller{contract: contract}, nil
}

// NewFakeSolTransactor creates a new write-only instance of FakeSol, bound to a specific deployed contract.
func NewFakeSolTransactor(address common.Address, transactor bind.ContractTransactor) (*FakeSolTransactor, error) {
	contract, err := bindFakeSol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FakeSolTransactor{contract: contract}, nil
}

// NewFakeSolFilterer creates a new log filterer instance of FakeSol, bound to a specific deployed contract.
func NewFakeSolFilterer(address common.Address, filterer bind.ContractFilterer) (*FakeSolFilterer, error) {
	contract, err := bindFakeSol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FakeSolFilterer{contract: contract}, nil
}

// bindFakeSol binds a generic wrapper to an already deployed contract.
func bindFakeSol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FakeSolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FakeSol *FakeSolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FakeSol.Contract.FakeSolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FakeSol *FakeSolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FakeSol.Contract.FakeSolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FakeSol *FakeSolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FakeSol.Contract.FakeSolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FakeSol *FakeSolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FakeSol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FakeSol *FakeSolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FakeSol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FakeSol *FakeSolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FakeSol.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_FakeSol *FakeSolCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_FakeSol *FakeSolSession) INITIALSUPPLY() (*big.Int, error) {
	return _FakeSol.Contract.INITIALSUPPLY(&_FakeSol.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_FakeSol *FakeSolCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _FakeSol.Contract.INITIALSUPPLY(&_FakeSol.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FakeSol *FakeSolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FakeSol *FakeSolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FakeSol.Contract.Allowance(&_FakeSol.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FakeSol *FakeSolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FakeSol.Contract.Allowance(&_FakeSol.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FakeSol *FakeSolCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FakeSol *FakeSolSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FakeSol.Contract.BalanceOf(&_FakeSol.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FakeSol *FakeSolCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FakeSol.Contract.BalanceOf(&_FakeSol.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FakeSol *FakeSolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FakeSol *FakeSolSession) Decimals() (uint8, error) {
	return _FakeSol.Contract.Decimals(&_FakeSol.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FakeSol *FakeSolCallerSession) Decimals() (uint8, error) {
	return _FakeSol.Contract.Decimals(&_FakeSol.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FakeSol *FakeSolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FakeSol *FakeSolSession) Name() (string, error) {
	return _FakeSol.Contract.Name(&_FakeSol.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FakeSol *FakeSolCallerSession) Name() (string, error) {
	return _FakeSol.Contract.Name(&_FakeSol.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FakeSol *FakeSolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FakeSol *FakeSolSession) Symbol() (string, error) {
	return _FakeSol.Contract.Symbol(&_FakeSol.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FakeSol *FakeSolCallerSession) Symbol() (string, error) {
	return _FakeSol.Contract.Symbol(&_FakeSol.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FakeSol *FakeSolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FakeSol.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FakeSol *FakeSolSession) TotalSupply() (*big.Int, error) {
	return _FakeSol.Contract.TotalSupply(&_FakeSol.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FakeSol *FakeSolCallerSession) TotalSupply() (*big.Int, error) {
	return _FakeSol.Contract.TotalSupply(&_FakeSol.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FakeSol *FakeSolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Approve(&_FakeSol.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Approve(&_FakeSol.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_FakeSol *FakeSolTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_FakeSol *FakeSolSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Burn(&_FakeSol.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_FakeSol *FakeSolTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Burn(&_FakeSol.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_FakeSol *FakeSolTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_FakeSol *FakeSolSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.BurnFrom(&_FakeSol.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_FakeSol *FakeSolTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.BurnFrom(&_FakeSol.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FakeSol *FakeSolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FakeSol *FakeSolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.DecreaseAllowance(&_FakeSol.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FakeSol *FakeSolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.DecreaseAllowance(&_FakeSol.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FakeSol *FakeSolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FakeSol *FakeSolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.IncreaseAllowance(&_FakeSol.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FakeSol *FakeSolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.IncreaseAllowance(&_FakeSol.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Transfer(&_FakeSol.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.Transfer(&_FakeSol.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.TransferFrom(&_FakeSol.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FakeSol *FakeSolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FakeSol.Contract.TransferFrom(&_FakeSol.TransactOpts, from, to, value)
}

// FakeSolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FakeSol contract.
type FakeSolApprovalIterator struct {
	Event *FakeSolApproval // Event containing the contract specifics and raw log

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
func (it *FakeSolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FakeSolApproval)
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
		it.Event = new(FakeSolApproval)
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
func (it *FakeSolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FakeSolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FakeSolApproval represents a Approval event raised by the FakeSol contract.
type FakeSolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FakeSol *FakeSolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FakeSolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FakeSol.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FakeSolApprovalIterator{contract: _FakeSol.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FakeSol *FakeSolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FakeSolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FakeSol.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FakeSolApproval)
				if err := _FakeSol.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FakeSol *FakeSolFilterer) ParseApproval(log types.Log) (*FakeSolApproval, error) {
	event := new(FakeSolApproval)
	if err := _FakeSol.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FakeSolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FakeSol contract.
type FakeSolTransferIterator struct {
	Event *FakeSolTransfer // Event containing the contract specifics and raw log

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
func (it *FakeSolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FakeSolTransfer)
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
		it.Event = new(FakeSolTransfer)
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
func (it *FakeSolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FakeSolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FakeSolTransfer represents a Transfer event raised by the FakeSol contract.
type FakeSolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FakeSol *FakeSolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FakeSolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FakeSol.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FakeSolTransferIterator{contract: _FakeSol.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FakeSol *FakeSolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FakeSolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FakeSol.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FakeSolTransfer)
				if err := _FakeSol.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FakeSol *FakeSolFilterer) ParseTransfer(log types.Log) (*FakeSolTransfer, error) {
	event := new(FakeSolTransfer)
	if err := _FakeSol.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
