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
)

func newBasSubClient() (*dnscontract.BasSubDomain, *ethclient.Client, error) {
	var (
		cli *ethclient.Client
		sub *dnscontract.BasSubDomain
		err error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		return nil, nil, err
	}

	sub, err = dnscontract.NewBasSubDomain(common.HexToAddress(config.GetRConf().Cconf.BasSubContract), cli)
	if err != nil {
		cli.Close()
		return nil, nil, err
	}

	return sub, cli, nil
}

type SubName struct {
	Name       string            `json:"name"`
	NameHash   string            `json:"nameHash"`
	RootHash   string            `json:"rootHash"`
	Owner      string            `json:"owner"`
	ExpireTime int64             `json:"expireTime"`
	Conf       map[string][]byte `json:"conf"`
}

func BatchNewSubName(start, end uint64) {
	subC, cli, err := newBasSubClient()
	if err != nil {
		log.Println("BatchNewSubName", err)
		return
	}

	defer cli.Close()

	var tr *BASTRC
	tr, err = newBasTRClient(cli)
	if err != nil {
		log.Println("BatchNewSubName", err)
		return
	}

	var bascfg *BASConfC
	bascfg, err = newBasConfClient(cli)
	if err != nil {
		log.Println("BatchNewSubName", err)
		return
	}

	b, e := start, end

	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	var iter *dnscontract.BasSubDomainNewSubDomainIterator
	iter, err = subC.FilterNewSubDomain(op, nil)
	if err != nil {
		fmt.Println("BatchNewSubName", err)
		return
	}
	defer iter.Close()
	db := ldb.GetLdb()
	//defer db.CloseLdb()
	var ts []string
	ts, err = db.GetTypeNames()
	if err != nil {
		log.Println("BatchNewSubName", err)
		return
	}

	for iter.Next() {
		ev := iter.Event

		sn := &SubName{
			Name:     string(ev.TotalName),
			NameHash: hex.EncodeToString(ev.NameHash[:]),
			RootHash: hex.EncodeToString(ev.RootHash[:]),
		}

		var oi *OwnerInfo
		oi, err = tr.GetOwnerInfo(ev.NameHash)
		if err == nil {
			sn.Owner = hex.EncodeToString(oi.Owner[:])
			sn.ExpireTime = oi.ExpireTime
		}

		sn.Conf = bascfg.GetDnsConf(ev.NameHash, ts)

		j, _ := json.Marshal(sn)

		err = db.SaveSubName(sn.NameHash, string(j))
		if err != nil {
			fmt.Println("BatchNewSubName", "save to db error")
			continue
		}
		// add new owner:namehash
		NewOwner(sn.Owner, "snH_1_"+sn.NameHash, db)
	}
}
