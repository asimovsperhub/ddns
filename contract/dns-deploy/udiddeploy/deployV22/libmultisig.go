package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/accountant"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

//
//func DeployLibMultiSig(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
//	log.Println("begin to deploy Lib multisig ")
//	ds := NewDeploySession(priv)
//	if err := ds.GetSession(3000000, big.NewInt(20000000000)); err != nil {
//		return common.Address{}, err
//	}
//	defer ds.Destroy()
//
//	return ds.DeployLibMultiSig()
//}
//
//func (ds *DeploySession) DeployLibMultiSig() (addr common.Address, err error) {
//	var tx *types.Transaction
//	_, tx, _, err = accountant.DeployLibMultiSig(ds.auth, ds.cli)
//	if err != nil {
//		return common.Address{}, err
//	}
//	addr, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)
//	log.Println("end to deploy Lib multisig ")
//	return
//}

func (c *Contract) DeployLibMultiSig(priv *ecdsa.PrivateKey) error {
	log.Println("begin to deploy LibMultiSig")
	parsed, err := accountant.LibMultiSigMetaData.GetAbi()
	if err != nil {
		return err
	}
	if parsed == nil {
		return errors.New("GetABI returned nil")
	}
	if AddrIsZero(c.LibBytes32) {
		if c.LibBytes32, err = DeployLibBytes32(priv); err != nil {
			return err
		}
	}
	accountant.LibMultiSigMetaData.Bin = strings.ReplaceAll(accountant.LibMultiSigMetaData.Bin, "__$85322f0591bc1264ab1fdf757d429c9497$__", c.LibBytes32.String()[2:])
	if AddrIsZero(c.LibAddressArray) {
		if c.LibAddressArray, err = DeployLibAddressArray(priv); err != nil {
			return err
		}
	}

	accountant.LibMultiSigMetaData.Bin = strings.ReplaceAll(accountant.LibMultiSigMetaData.Bin, "__$ae7d58307b48f566e6dd130bee0f44ebb6$__", c.LibAddressArray.String()[2:])

	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return err
	}
	defer ds.Destroy()

	var tx *types.Transaction

	c.LibMultiSig, tx, _, err = bind.DeployContract(ds.auth, *parsed, common.FromHex(accountant.LibMultiSigMetaData.Bin), ds.cli)
	if err != nil {
		return err
	}

	c.LibMultiSig, err = bind.WaitDeployed(context.TODO(), ds.cli, tx)
	if err != nil {
		return err
	}
	log.Println("end to deploy LibMultiSig")
	return nil
}
