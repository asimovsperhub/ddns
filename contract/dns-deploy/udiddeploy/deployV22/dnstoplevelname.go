package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/top"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
)

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

func (c *Contract) SetDnsTopContract(priv *ecdsa.PrivateKey, nswitch uint8) (r *types.Receipt, err error) {
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
	if tx, err = t.SetContract(ds.auth, c.Cost, c.Accountant, c.SecondName, nswitch); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) GetOwner() (addr common.Address, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}

	return t.Owner(nil)
}

func (c *Contract) MintName(priv *ecdsa.PrivateKey, entireName string, year uint8, erc20Addr common.Address) (r *types.Receipt, err error) {

	var price *big.Int
	price, err = c.GetTopLevelNameUsdtPrice(uint8(len(param.name)), year)
	if err != nil {
		fmt.Println(err)
		return
	}
	var (
		exPrice, roundId *big.Int
	)

	if erc20Addr != c.UsdtAddr {
		exPrice, roundId, err = c.GetLatestPrice(erc20Addr, price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		exPrice = price
		roundId = &big.Int{}
	}

	ds := NewDeploySession(priv)
	if AddrIsZero(erc20Addr) {
		if err = ds.GetSessionWithEthValue(0, nil, exPrice); err != nil {
			return
		}
		defer ds.Destroy()
		var balance *big.Int
		if balance, err = ds.cli.BalanceAt(context.Background(), crypto.PubkeyToAddress(priv.PublicKey), nil); err != nil {
			fmt.Println(err)

			return nil, err
		}
		if balance.Cmp(exPrice) < 0 {
			fmt.Println("eth not enough")

			return nil, errors.New("eth not enough")
		}

	} else {
		//var r *types.Receipt
		fmt.Println("approve to", erc20Addr.String(), "amount:", exPrice.String())
		if _, err = Approve(priv, erc20Addr, c.TopName, exPrice); err != nil {
			fmt.Println(err)
			return
		}
		if err = ds.GetSession(0, nil); err != nil {
			return
		}

		defer ds.Destroy()

	}

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	fmt.Println("begin mint", entireName, "for", year, "year")
	var tx *types.Transaction
	if tx, err = t.MintTopLevelName(ds.auth, entireName, year, erc20Addr, roundId); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) NameHashIsExists(nameHash common.Hash) (b bool, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}

	n, err1 := t.NameStore(nil, nameHash)
	if err1 != nil {
		return false, err1
	}

	if n.ExpireTime > 0 {
		return true, nil
	}

	return false, nil

}

func (c *Contract) NameIsExists(name string) (bool, error) {
	hash := crypto.Keccak256Hash([]byte(name))
	return c.NameHashIsExists(hash)
}

func (c *Contract) GetErc721Addr(nameHash common.Hash) (erc721Addr common.Address, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	return t.Erc721Store(nil, nameHash)
}

func (c *Contract) GetDnsInfo(nameHash common.Hash) (entireName string, expTime uint32, tokenId *big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var r struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	}
	if r, err = t.NameStore(nil, nameHash); err != nil {
		return
	}

	return r.EntireName, r.ExpireTime, r.TokenId, nil
}

func (c *Contract) Open2Reg(priv *ecdsa.PrivateKey, nameHash common.Hash) (r *types.Receipt, err error) {
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
	if tx, err = t.Open2Reg(ds.auth, nameHash); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}
