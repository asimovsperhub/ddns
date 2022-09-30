package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/cost"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
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

func (c *Contract) DisablePayment(priv *ecdsa.PrivateKey,
	pamentErc20 common.Address, decimal uint8) (r *types.Receipt, err error) {
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
	if tx, err = t.SetPayment(ds.auth, pamentErc20, decimal, false); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return

}

func (c *Contract) AddPayment(priv *ecdsa.PrivateKey,
	pamentErc20 common.Address,
	decimal uint8,
	chainlinkAddr common.Address) (r *types.Receipt, err error) {
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
	if tx, err = t.AddPayment(ds.auth, pamentErc20, decimal, chainlinkAddr); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return

}

func (c *Contract) SetTopLevelPrice(priv *ecdsa.PrivateKey, price [8]*big.Int) (r *types.Receipt, err error) {
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
	if tx, err = t.SetTopLevelNamePrice(ds.auth, price); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) SetSecondTax(priv *ecdsa.PrivateKey, price [8]*big.Int) (r *types.Receipt, err error) {
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
	if tx, err = t.SetSecondLevelNameTax(ds.auth, price); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) SetRelation(priv *ecdsa.PrivateKey) (r *types.Receipt, err error) {

	if AddrIsZero(c.TopName) {
		return nil, errors.New("no top name contract")
	}

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
	if tx, err = t.SetRelation(ds.auth, c.TopName); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) SetSecondPrice(priv *ecdsa.PrivateKey, hash common.Hash, price [8]*big.Int) (r *types.Receipt, err error) {
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
	if tx, err = t.SetSecondLevelNamePrice(ds.auth, hash, price); err != nil {
		return
	}

	if r, err = bind.WaitMined(context.Background(), ds.cli, tx); err != nil {
		return
	}

	return
}

func (c *Contract) GetLatestPrice(erc20Addr common.Address, price *big.Int) (exPrice *big.Int, roundId *big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}
	fmt.Println(erc20Addr.String(), price.String())
	return t.GetLatestExchangeValue(nil, erc20Addr, price)
}

func (c *Contract) GetTopLevelNameUsdtPrice(nameLen uint8, year uint8) (price *big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}
	var b bool
	if price, b, err = t.GetTopLevelNamePrice(nil, c.UsdtAddr, big.NewInt(0), nameLen, year); err != nil {
		return nil, err
	}
	if !b {
		return nil, errors.New("not usdt erc20 address")
	}

	return price, nil

}

func (c *Contract) GetSecondLevelNameUsdtPrice(nameLen uint8, year uint8, topHash common.Hash) (tax *big.Int, price *big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}
	var b bool
	if tax, price, b, err = t.GetSecondLevelNamePrice(nil, topHash, c.UsdtAddr, big.NewInt(0), nameLen, year); err != nil {
		return nil, nil, err
	}

	if !b {
		return nil, nil, errors.New("not usdt erc20 address")
	}

	return

}

func (c *Contract) GetPriceConf(nameHash common.Hash) (price [8]*big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}

	if nameHash == common.Hash(common.Hash{}) {
		if price, err = t.GetAllTopLevelNamePrice(nil); err != nil {
			return
		}
	} else {
		if price, err = t.GetAllSecondLevelNamePrice(nil, nameHash); err != nil {
			return
		}
	}

	return price, nil
}

func (c *Contract) GetTaxConf() (tax [8]*big.Int, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()
	var t *cost.DnsCost
	if t, err = cost.NewDnsCost(c.Cost, ds.cli); err != nil {
		return
	}

	if tax, err = t.GetAllSecondLevelNameTax(nil); err != nil {
		return
	}
	return
}
