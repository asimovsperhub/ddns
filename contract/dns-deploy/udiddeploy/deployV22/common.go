package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/top"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type DeploySession struct {
	cli  *ethclient.Client
	auth *bind.TransactOpts
	priv *ecdsa.PrivateKey
}

type Contract struct {
	LibSig          common.Address     `json:"lib_sig"`
	LibToolKit      common.Address     `json:"lib_tool_kit"`
	LibErc721       common.Address     `json:"lib_erc_721"`
	LibAddressArray common.Address     `json:"lib_address_array"`
	LibBytes32      common.Address     `json:"lib_bytes_32"`
	LibMultiSig     common.Address     `json:"lib_multi_sig"`
	LibAccountant   common.Address     `json:"lib_accountant"`
	TopName         common.Address     `json:"top_name"`
	Cost            common.Address     `json:"cost"`
	Accountant      common.Address     `json:"accountant"`
	SecondName      common.Address     `json:"second_name"`
	TopTx           *types.Transaction `json:"-"`
	CostTx          *types.Transaction `json:"-"`
	AcctTx          *types.Transaction `json:"-"`
	SecondTx        *types.Transaction `json:"-"`
	UsdtAddr        common.Address     `json:"usdt_addr"`
}

func (c *Contract) String() string {
	msg := ""
	msg += fmt.Sprintf("libsig contract address: %s\r\n", c.LibSig.String())
	msg += fmt.Sprintf("libToolKit contract address: %s\r\n", c.LibToolKit.String())
	msg += fmt.Sprintf("libErc721 contract address: %s\r\n", c.LibErc721.String())
	msg += fmt.Sprintf("LibAdddressArray contract address: %s\r\n", c.LibAddressArray.String())
	msg += fmt.Sprintf("LibBytes32 contract address: %s\r\n", c.LibBytes32.String())
	msg += fmt.Sprintf("LibMultiSig contract address: %s\r\n", c.LibMultiSig.String())
	msg += fmt.Sprintf("LibAccountant contract address: %s\r\n", c.LibAccountant.String())
	msg += fmt.Sprintf("TopName contract address: %s\r\n", c.TopName.String())
	msg += fmt.Sprintf("Cost contract address: %s\r\n", c.Cost.String())
	msg += fmt.Sprintf("Accountant contract address: %s\r\n", c.Accountant.String())
	msg += fmt.Sprintf("SecondName contract address: %s\r\n", c.SecondName.String())
	msg += fmt.Sprintf("Usdt Address:%s\r\n", c.UsdtAddr.String())
	if c.TopTx != nil {
		msg += fmt.Sprintf("TopLevelName contract success tx: %s\r\n", c.TopTx.Hash().String())
	}
	if c.CostTx != nil {
		msg += fmt.Sprintf("Cost contract success tx: %s\r\n", c.CostTx.Hash().String())
	}
	if c.AcctTx != nil {
		msg += fmt.Sprintf("Accountant contract success tx: %s\r\n", c.AcctTx.Hash().String())
	}
	if c.SecondTx != nil {
		msg += fmt.Sprintf("Second contract success tx: %s\r\n", c.SecondTx.Hash().String())
	}

	return msg
}

const (
	ERCName   = "UDID Network"
	ERCSymbol = "UDIDN"
)

func AddrIsZero(addr common.Address) bool {
	var r common.Address
	if r == addr {
		return true
	}

	return false

}

func NewDeploySession(priv *ecdsa.PrivateKey) *DeploySession {
	return &DeploySession{priv: priv}
}

func (ds *DeploySession) Destroy() {
	ds.cli.Close()
}

func (ds *DeploySession) GetClient() (err error) {
	var cli *ethclient.Client
	cli, err = ethclient.Dial(ethacct.InfuraPoint)
	if err != nil {
		return err
	}
	ds.cli = cli

	return nil
}

func (ds *DeploySession) GetSession(gasLimit uint64, gasPrice *big.Int) (err error) {
	var cli *ethclient.Client
	cli, err = ethclient.Dial(ethacct.InfuraPoint)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			cli.Close()
		}
	}()
	var nonce uint64
	if nonce, err = cli.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(ds.priv.PublicKey)); err != nil {
		return err
	}
	if gasPrice == nil || gasPrice.Cmp(&big.Int{}) == 0 {
		if gasPrice, err = cli.SuggestGasPrice(context.Background()); err != nil {
			return err
		}
	}

	var chainID *big.Int
	if chainID, err = cli.ChainID(context.Background()); err != nil {
		return err
	}
	var auth *bind.TransactOpts

	if auth, err = bind.NewKeyedTransactorWithChainID(ds.priv, chainID); err != nil {
		return err
	}
	auth.From = crypto.PubkeyToAddress(ds.priv.PublicKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	if gasLimit > 0 {
		auth.GasLimit = uint64(gasLimit)
	}
	auth.GasPrice = gasPrice
	//z := big.NewInt(int64(auth.GasLimit))
	//if gasLimit == 0 {
	//	log.Println("gas fee system default")
	//} else {
	//	log.Println("gas fee", z.Mul(z, gasPrice).String())
	//}
	//log.Println("nonce", nonce)

	ds.auth = auth
	ds.cli = cli
	return nil
}

func (ds *DeploySession) GetSessionWithEthValue(gasLimit uint64, gasPrice *big.Int, value *big.Int) (err error) {
	var cli *ethclient.Client
	cli, err = ethclient.Dial(ethacct.InfuraPoint)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			cli.Close()
		}
	}()
	var nonce uint64
	if nonce, err = cli.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(ds.priv.PublicKey)); err != nil {
		return err
	}
	if gasPrice == nil || gasPrice.Cmp(&big.Int{}) == 0 {
		if gasPrice, err = cli.SuggestGasPrice(context.Background()); err != nil {
			return err
		}
	}

	var chainID *big.Int
	if chainID, err = cli.ChainID(context.Background()); err != nil {
		return err
	}
	var auth *bind.TransactOpts

	if auth, err = bind.NewKeyedTransactorWithChainID(ds.priv, chainID); err != nil {
		return err
	}
	auth.From = crypto.PubkeyToAddress(ds.priv.PublicKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	if gasLimit > 0 {
		auth.GasLimit = uint64(gasLimit)
	}
	auth.GasPrice = gasPrice

	ds.auth = auth
	ds.cli = cli
	return nil
}

func (ds *DeploySession) Approve(erc20addr, spender common.Address, amount *big.Int) (r *types.Receipt, err error) {
	var erc20 *top.IERC20
	if erc20, err = top.NewIERC20(erc20addr, ds.cli); err != nil {
		return nil, err
	}
	var balance *big.Int
	if balance, err = erc20.BalanceOf(nil, crypto.PubkeyToAddress(ds.priv.PublicKey)); err != nil {
		return nil, err
	}

	if balance.Cmp(amount) < 0 {
		return nil, errors.New("erc20 balance not enough")
	}

	var tx *types.Transaction
	if tx, err = erc20.Approve(ds.auth, spender, amount); err != nil {
		return nil, err
	}

	return bind.WaitMined(context.Background(), ds.cli, tx)
}

func Approve(priv *ecdsa.PrivateKey, erc20addr, spender common.Address, amount *big.Int) (r *types.Receipt, err error) {
	ds := NewDeploySession(priv)
	if err := ds.GetSession(0, nil); err != nil {
		return nil, err
	}
	defer ds.Destroy()

	return ds.Approve(erc20addr, spender, amount)
}
