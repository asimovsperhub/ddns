package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/chainlink"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func DeployUsdcUsdt(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
	log.Println("begin to deploy usdc-usdt ")
	ds := NewDeploySession(priv)
	if err := ds.GetSession(0, nil); err != nil {
		return common.Address{}, err
	}
	defer ds.Destroy()

	return ds.DeployUsdcUsdt()
}

func (ds *DeploySession) DeployUsdcUsdt() (addr common.Address, err error) {
	var tx *types.Transaction
	_, tx, _, err = chainlink.DeployUSDCUSDT(ds.auth, ds.cli)
	if err != nil {
		return common.Address{}, err
	}
	addr, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)

	log.Println("end to deploy usdc-usdt ")

	return
}
