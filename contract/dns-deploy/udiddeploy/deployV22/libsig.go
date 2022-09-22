package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/top"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func DeployLibSig(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
	log.Println("begin to deploy Lib signature ")
	ds := NewDeploySession(priv)
	if err := ds.GetSession(0, nil); err != nil {
		return common.Address{}, err
	}
	defer ds.Destroy()

	return ds.DeployLibSignature()
}

func (ds *DeploySession) DeployLibSignature() (addr common.Address, err error) {
	var tx *types.Transaction
	_, tx, _, err = top.DeployLibDnsSignature(ds.auth, ds.cli)
	if err != nil {
		return common.Address{}, err
	}
	addr, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)
	log.Println("end to deploy Lib signature ")
	return
}
