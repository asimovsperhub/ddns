package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/erc721"
	"github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/top"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (c *Contract) GetSigUserAddr(nameHash common.Hash) (addr common.Address, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var erc721Addr common.Address
	if erc721Addr, err = t.Erc721Store(nil, nameHash); err != nil {
		return [20]byte{}, err
	}

	var e *erc721.DnsNameErc721
	if e, err = erc721.NewDnsNameErc721(erc721Addr, ds.cli); err != nil {
		return [20]byte{}, err
	}
	if addr, err = e.SigUserAddr(nil); err != nil {
		return [20]byte{}, err
	}
	return
}

func (c *Contract) GetIdInfo(nameHash common.Hash, id *big.Int) (expireTime uint32, hash common.Hash, owner common.Address, err error) {
	ds := NewDeploySession(nil)
	if err = ds.GetClient(); err != nil {
		return
	}
	defer ds.Destroy()

	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var erc721Addr common.Address
	if erc721Addr, err = t.Erc721Store(nil, nameHash); err != nil {
		return 0, [32]byte{}, common.Address{}, err
	}

	var e *erc721.DnsNameErc721
	if e, err = erc721.NewDnsNameErc721(erc721Addr, ds.cli); err != nil {
		return 0, [32]byte{}, common.Address{}, err
	}
	var r struct {
		ExpireTime uint32
		NameHash   [32]byte
	}
	if r, err = e.NameIdInfo(nil, id); err != nil {
		return 0, [32]byte{}, common.Address{}, err
	}

	if owner, err = e.OwnerOf(nil, id); err != nil {
		return 0, [32]byte{}, common.Address{}, err
	}

	return r.ExpireTime, r.NameHash, owner, nil
}

func (c *Contract) SetMetaUrl(priv *ecdsa.PrivateKey, nameHash common.Hash, url string) (r *types.Receipt, err error) {
	ds := NewDeploySession(priv)
	if err = ds.GetSession(0, nil); err != nil {
		return
	}
	var t *top.DnsTopLevelName
	if t, err = top.NewDnsTopLevelName(c.TopName, ds.cli); err != nil {
		return
	}
	var erc721Addr common.Address
	if erc721Addr, err = t.Erc721Store(nil, nameHash); err != nil {
		return
	}

	var e *erc721.DnsNameErc721
	if e, err = erc721.NewDnsNameErc721(erc721Addr, ds.cli); err != nil {
		return
	}
	var tx *types.Transaction
	if tx, err = e.SetBaseUri(nil, url); err != nil {
		return
	}
	return bind.WaitMined(context.TODO(), ds.cli, tx)
}
