package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type DeploySession struct {
	cli  *ethclient.Client
	auth *bind.TransactOpts
	priv *ecdsa.PrivateKey
}

func NewDeploySession(priv *ecdsa.PrivateKey) *DeploySession {
	return &DeploySession{priv: priv}
}

func (ds *DeploySession) Destroy() {
	ds.cli.Close()
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
	z := big.NewInt(int64(auth.GasLimit))
	if gasLimit == 0 {
		log.Println("gas fee system default")
	} else {
		log.Println("gas fee", z.Mul(z, gasPrice).String())
	}
	//log.Println("nonce", nonce)

	ds.auth = auth
	ds.cli = cli
	return nil
}
