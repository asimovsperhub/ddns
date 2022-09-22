package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/second"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

//type DnsSecondContract struct {
//	LibSig     common.Address
//	LibToolKit common.Address
//	SecondName common.Address
//	Tx         *types.Transaction
//}
//
//func (dsc *DnsSecondContract) String() string {
//	msg := ""
//	msg += fmt.Sprintf("LibSignature contract address: %s\r\n", dsc.LibSig.String())
//	msg += fmt.Sprintf("LibToolKit contract address: %s\r\n", dsc.LibToolKit.String())
//	msg += fmt.Sprintf("SecondName contract address: %s\r\n", dsc.SecondName.String())
//	msg += fmt.Sprintf("SecondName contract success tx: %s\r\n", dsc.Tx.Hash().String())
//	return msg
//}

func (c *Contract) DeploySecondContract(priv *ecdsa.PrivateKey) error {
	log.Println("begin to deploy dns second level name")
	parsed, err := second.DnsSecondLevelNameMetaData.GetAbi()
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

	if AddrIsZero(c.Cost) {
		if err = c.DeployCost(priv); err != nil {
			return err
		}
	}

	if AddrIsZero(c.Accountant) {
		if err = c.DeployDnsAccountant(priv); err != nil {
			return err
		}
	}
	if AddrIsZero(c.LibSig) {
		if c.LibSig, err = DeployLibSig(priv); err != nil {
			return err
		}
	}
	second.DnsSecondLevelNameMetaData.Bin = strings.ReplaceAll(second.DnsSecondLevelNameMetaData.Bin, "__$2ed95cc1c04a020bf25a1d27d6730d85d6$__", c.LibSig.String()[2:])
	if AddrIsZero(c.LibToolKit) {
		if c.LibToolKit, err = DeployLibToolKit(priv); err != nil {
			return err
		}
	}
	second.DnsSecondLevelNameMetaData.Bin = strings.ReplaceAll(second.DnsSecondLevelNameMetaData.Bin, "__$61b1050b44c222c225346b0c1857883025$__", c.LibToolKit.String()[2:])
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return err
	}
	defer ds.Destroy()

	c.SecondName, c.SecondTx, _, err = bind.DeployContract(ds.auth, *parsed, common.FromHex(second.DnsSecondLevelNameMetaData.Bin), ds.cli, c.Cost, c.Accountant, c.TopName)
	if err != nil {
		return err
	}

	c.SecondName, err = bind.WaitDeployed(context.TODO(), ds.cli, c.SecondTx)
	if err != nil {
		return err
	}
	log.Println("end to deploy dns second level name")
	return nil

}

func (c *Contract) SetDnsSecondContract(priv *ecdsa.PrivateKey) (r *types.Receipt, err error) {
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return
	}
	defer ds.Destroy()
	var t *second.DnsSecondLevelName
	if t, err = second.NewDnsSecondLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var tx *types.Transaction
	if tx, err = t.SetContract(ds.auth, c.Cost, c.Accountant, c.TopName, 0); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}
