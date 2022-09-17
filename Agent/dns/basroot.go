package dns

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/contract/dnscontract"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func newBasRootClient() (*dnscontract.BasRootDomain, *ethclient.Client, error) {
	var (
		cli  *ethclient.Client
		root *dnscontract.BasRootDomain
		err  error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		return nil, nil, err
	}

	root, err = dnscontract.NewBasRootDomain(common.HexToAddress(config.GetRConf().Cconf.BasRootContract), cli)
	if err != nil {
		cli.Close()
		return nil, nil, err
	}

	return root, cli, nil
}

type RootName struct {
	Name        string            `json:"name"`
	Hash        string            `json:"hash"`
	IsCustom    bool              `json:"isCustom"`
	CustomPrice *big.Int          `json:"customPrice"`
	IsOpen      bool              `json:"isOpen"`
	Owner       string            `json:"owner"`
	ExpireTime  int64             `json:"expireTime"`
	Conf        map[string][]byte `json:"conf"`
}

func BatchNewRootName(start, end uint64) {
	rootC, cli, err := newBasRootClient()
	if err != nil {
		log.Println("BatchNewRootName", err)
		return
	}

	defer cli.Close()
	var tr *BASTRC
	tr, err = newBasTRClient(cli)
	if err != nil {
		log.Println("BatchNewRootName", err)
		return
	}

	var bascfg *BASConfC
	bascfg, err = newBasConfClient(cli)
	if err != nil {
		log.Println("BatchNewRootName", err)
		return
	}

	b, e := start, end

	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	var iter *dnscontract.BasRootDomainNewRootDomainIterator
	iter, err = rootC.FilterNewRootDomain(op, nil)
	if err != nil {
		log.Println("BatchNewRootName", err)
		return
	}
	defer iter.Close()
	db := ldb.GetLdb()
	//defer db.CloseLdb()
	var ts []string
	ts, err = db.GetTypeNames()
	if err != nil {
		log.Println("BatchNewRootName", err)
		return
	}

	for iter.Next() {
		ev := iter.Event
		rn := &RootName{
			Name:        string(ev.RootName),
			Hash:        hex.EncodeToString(ev.NameHash[:]),
			IsOpen:      ev.OpenToPublic,
			IsCustom:    ev.IsCustom,
			CustomPrice: ev.CustomPrice,
		}

		var oi *OwnerInfo
		oi, err = tr.GetOwnerInfo(ev.NameHash)
		if err == nil {
			rn.Owner = hex.EncodeToString(oi.Owner[:])
			rn.ExpireTime = oi.ExpireTime

		}

		rn.Conf = bascfg.GetDnsConf(ev.NameHash, ts)

		j, _ := json.Marshal(rn)

		err = db.SaveRootName(rn.Hash, string(j))
		if err != nil {
			fmt.Println("BatchNewRootName", "save to db error")
			continue
		}
		// add new owner:namehash
		NewOwner(rn.Owner, "rnH_1_"+rn.Hash, db)
	}

}
