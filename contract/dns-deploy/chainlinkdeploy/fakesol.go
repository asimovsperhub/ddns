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

func DeployFakeSol(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
	log.Println("begin to deploy fake-sol ")
	ds := NewDeploySession(priv)
	if err := ds.GetSession(0, nil); err != nil {
		return common.Address{}, err
	}
	defer ds.Destroy()

	return ds.DeployFakeSol()
}

func (ds *DeploySession) DeployFakeSol() (addr common.Address, err error) {
	var tx *types.Transaction
	_, tx, _, err = chainlink.DeployFakeSol(ds.auth, ds.cli)
	if err != nil {
		return common.Address{}, err
	}
	addr, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)

	log.Println("end to deploy fake-sol ")

	return
}
