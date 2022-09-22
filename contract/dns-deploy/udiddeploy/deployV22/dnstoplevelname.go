package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/top"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
)

//type TopLevelContract struct {
//	LibSig     common.Address
//	LibToolKit common.Address
//	LibErc721  common.Address
//	TopName    common.Address
//	Tx         *types.Transaction
//}

//func (tc *TopLevelContract) String() string {
//	msg := ""
//	msg += fmt.Sprintf("libsig contract address: %s\r\n", tc.LibSig.String())
//	msg += fmt.Sprintf("libToolKit contract address: %s\r\n", tc.LibToolKit.String())
//	msg += fmt.Sprintf("libErc721 contract address: %s\r\n", tc.LibErc721.String())
//	msg += fmt.Sprintf("TopLevelName contract address: %s\r\n", tc.TopName.String())
//	msg += fmt.Sprintf("TopLevelName contract success tx: %s\r\n", tc.Tx.Hash().String())
//	return msg
//}

func (c *Contract) DeployDnsTopLevelName(priv *ecdsa.PrivateKey) error {
	log.Println("begin to deploy dns top level name")
	parsed, err := top.DnsTopLevelNameMetaData.GetAbi()
	if err != nil {
		return err
	}
	if parsed == nil {
		return errors.New("GetABI returned nil")
	}

	if AddrIsZero(c.LibSig) {
		if c.LibSig, err = DeployLibSig(priv); err != nil {
			return err
		}
	}
	top.DnsTopLevelNameMetaData.Bin = strings.ReplaceAll(top.DnsTopLevelNameMetaData.Bin, "__$2ed95cc1c04a020bf25a1d27d6730d85d6$__", c.LibSig.String()[2:])
	if AddrIsZero(c.LibToolKit) {
		if c.LibToolKit, err = DeployLibToolKit(priv); err != nil {
			return err
		}
	}

	top.DnsTopLevelNameMetaData.Bin = strings.ReplaceAll(top.DnsTopLevelNameMetaData.Bin, "__$61b1050b44c222c225346b0c1857883025$__", c.LibToolKit.String()[2:])

	if AddrIsZero(c.LibErc721) {
		if c.LibErc721, err = DeployLibErc721(priv); err != nil {
			return err
		}
	}

	top.DnsTopLevelNameMetaData.Bin = strings.ReplaceAll(top.DnsTopLevelNameMetaData.Bin, "__$d074600fcdd3cf12194b4b963c2fdc2769$__", c.LibErc721.String()[2:])

	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return err
	}
	defer ds.Destroy()

	c.TopName, c.TopTx, _, err = bind.DeployContract(ds.auth, *parsed, common.FromHex(top.DnsTopLevelNameMetaData.Bin), ds.cli, ERCName, ERCSymbol)
	if err != nil {
		return err
	}

	c.TopName, err = bind.WaitDeployed(context.TODO(), ds.cli, c.TopTx)
	if err != nil {
		return err
	}

	log.Println("end to deploy dns top level name")

	return nil

}

func (c *Contract) SetDnsTopContract(priv *ecdsa.PrivateKey) (r *types.Receipt, err error) {
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return
	}
	defer ds.Destroy()
	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var tx *types.Transaction
	if tx, err = t.SetContract(ds.auth, c.Cost, c.Accountant, c.SecondName, 0); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) GetOwner(priv *ecdsa.PrivateKey) (addr common.Address, err error) {
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return
	}
	defer ds.Destroy()
	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}

	return t.Owner(nil)
}
