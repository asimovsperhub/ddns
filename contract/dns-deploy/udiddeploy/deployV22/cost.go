package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/cost"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

//type CostContract struct {
//	LibAddressArray common.Address
//	Cost            common.Address
//	Tx              *types.Transaction
//}
//
//func (cc *CostContract) String() string {
//	msg := ""
//	msg += fmt.Sprintf("libAddressArray contract address: %s\r\n", cc.LibAddressArray.String())
//	msg += fmt.Sprintf("Cost contract address: %s\r\n", cc.Cost.String())
//	msg += fmt.Sprintf("Cost contract success tx: %s\r\n", cc.Tx.Hash().String())
//	return msg
//}

func (c *Contract) DeployCost(priv *ecdsa.PrivateKey) error {
	log.Println("begin to deploy cost")
	parsed, err := cost.DnsCostMetaData.GetAbi()
	if err != nil {
		return err
	}
	if parsed == nil {
		return errors.New("GetABI returned nil")
	}

	if AddrIsZero(c.TopName) {
		if err = c.DeployDnsTopLevelName(priv); err != nil {
			return err
		}
	}
	if AddrIsZero(c.LibAddressArray) {
		if c.LibAddressArray, err = DeployLibAddressArray(priv); err != nil {
			return err
		}
	}
	cost.DnsCostMetaData.Bin = strings.ReplaceAll(cost.DnsCostMetaData.Bin, "__$ae7d58307b48f566e6dd130bee0f44ebb6$__", c.LibAddressArray.String()[2:])
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return err
	}
	defer ds.Destroy()

	c.Cost, c.CostTx, _, err = bind.DeployContract(ds.auth, *parsed, common.FromHex(cost.DnsCostMetaData.Bin), ds.cli, c.TopName)
	if err != nil {
		return err
	}

	c.Cost, err = bind.WaitDeployed(context.TODO(), ds.cli, c.CostTx)
	if err != nil {
		return err
	}
	log.Println("end to deploy cost")
	return nil
}

func (c *Contract) SetUsdt(priv *ecdsa.PrivateKey, usdt common.Address, decimal uint8) (r *types.Receipt, err error) {
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}
	var tx *types.Transaction
	if tx, err = t.SetUsdtAddr(ds.auth, usdt, decimal); err != nil {

		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}
