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

func DeployLibErc721(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
	log.Println("begin to deploy Lib erc721 ")
	ds := NewDeploySession(priv)
	if err := ds.GetSession(0, nil); err != nil {
		return common.Address{}, err
	}
	defer ds.Destroy()

	return ds.DeployLibErc721()
}

func (ds *DeploySession) DeployLibErc721() (addr common.Address, err error) {
	var tx *types.Transaction
	_, tx, _, err = top.DeployLibDnsNameErc721(ds.auth, ds.cli)
	if err != nil {
		return common.Address{}, err
	}
	addr, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)

	log.Println("end to deploy Lib erc721 ")

	return
}
