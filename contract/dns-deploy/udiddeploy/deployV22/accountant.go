package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/accountant"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

//type AccountantContract struct {
//	LibAddrArray  common.Address
//	LibMultiSig   common.Address
//	LibAccountant common.Address
//	Tx            *types.Transaction
//	Accountant    common.Address
//}
//
//func (ac *AccountantContract) String() string {
//	msg := ""
//	msg += fmt.Sprintf("libMultiSig contract address: %s\r\n", ac.LibMultiSig.String())
//	msg += fmt.Sprintf("LibAddressArray contract address: %s\r\n", ac.LibAddrArray.String())
//	msg += fmt.Sprintf("LibAccountant contract address: %s\r\n", ac.LibAccountant.String())
//	msg += fmt.Sprintf("Accountant contract address: %s\r\n", ac.Accountant.String())
//	msg += fmt.Sprintf("Accountant contract success tx: %s\r\n", ac.Tx.Hash().String())
//	return msg
//}

func (c *Contract) DeployDnsAccountant(priv *ecdsa.PrivateKey) error {
	log.Println("begin to deploy accountant")
	parsed, err := accountant.DnsAccountantMetaData.GetAbi()
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
	if AddrIsZero(c.LibMultiSig) {
		if err = c.DeployLibMultiSig(priv); err != nil {
			return err
		}
	}
	accountant.DnsAccountantMetaData.Bin = strings.ReplaceAll(accountant.DnsAccountantMetaData.Bin, "__$75b7a25587b53564bca376d1b99e97afd3$__", c.LibMultiSig.String()[2:])

	if AddrIsZero(c.LibAddressArray) {
		if c.LibAddressArray, err = DeployLibAddressArray(priv); err != nil {
			return err
		}
	}
	accountant.DnsAccountantMetaData.Bin = strings.ReplaceAll(accountant.DnsAccountantMetaData.Bin, "__$ae7d58307b48f566e6dd130bee0f44ebb6$__", c.LibAddressArray.String()[2:])

	if AddrIsZero(c.LibAccountant) {
		if c.LibAccountant, err = DeployLibAccountant(priv); err != nil {
			return err
		}
	}
	accountant.DnsAccountantMetaData.Bin = strings.ReplaceAll(accountant.DnsAccountantMetaData.Bin, "__$f1fbeac8e6c305d33caaca44e2a1f51440$__", c.LibAccountant.String()[2:])

	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return err
	}
	defer ds.Destroy()

	c.Accountant, c.AcctTx, _, err = bind.DeployContract(ds.auth, *parsed, common.FromHex(accountant.DnsAccountantMetaData.Bin), ds.cli, c.TopName)
	if err != nil {
		return err
	}

	c.Accountant, err = bind.WaitDeployed(context.TODO(), ds.cli, c.AcctTx)
	if err != nil {
		return err
	}
	log.Println("end to deploy accountant")

	return nil

}
