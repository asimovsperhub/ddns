package dns

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/contract/dnscontract"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func BatchMarketBuyFromSeller(start, end uint64) {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		log.Println("BatchMarketBuyFromSeller EthClient conn err", err)
		return
	}
	root, err := dnscontract.NewBasRootDomain(common.HexToAddress(config.GetRConf().Cconf.BasRootContract), cli)
	if err != nil {
		log.Println("BatchMarketBuyFromSeller newBasRootClient", err)
		return
	}
	sub, err := dnscontract.NewBasSubDomain(common.HexToAddress(config.GetRConf().Cconf.BasSubContract), cli)
	if err != nil {
		log.Println("BatchMarketBuyFromSeller newBasSubClient", err)
		return
	}
	tr, err := dnscontract.NewBasTradableOwnership(common.HexToAddress(config.GetRConf().Cconf.BasTradableOwnership), cli)
	if err != nil {
		log.Println("BatchMarketBuyFromSeller newBasTRClient  err", err)
		return
	}
	defer cli.Close()
	b, e := start, end
	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	iter, _ := tr.FilterTransferFrom(op)
	db := ldb.GetLdb()
	for iter.Next() {
		ev := iter.Event
		addr := ev.To
		if root_info, _ := root.Root(nil, ev.NameHash); len(root_info.RootName) != 0 {
			// parsing root domain
			name := hex.EncodeToString(ev.NameHash[:])
			val, _ := db.GetRootName(name)
			rootc := new(RootName)
			json.Unmarshal([]byte(val), rootc)
			if rootc.Owner != addr.String()[2:] {
				rootc.Owner = addr.String()[2:]
				j, _ := json.Marshal(rootc)
				err = db.SaveRootName(name, string(j))
				if err != nil {
					log.Println("BatchMarketRootBuyFromSeller", "save to db error")
					continue
				}
				// update owner:namehash
				UpdateOwner(ev.From.String(), ev.To.String(), "rnH_1_"+name, db)
			}
		} else {
			// parsing sub domain
			info, err := sub.Sub(nil, ev.NameHash)
			if err != nil {
				log.Println("BatchMarketRootBuyFromSeller parsing sub err ", err)
			}
			namehash := crypto.Keccak256(info.TotalName)
			name := hex.EncodeToString(namehash)
			val, _ := db.GetSubName(name)
			subc := new(SubName)
			json.Unmarshal([]byte(val), subc)
			if subc.Owner != addr.String()[2:] {
				subc.Owner = addr.String()[2:]
				j, _ := json.Marshal(subc)
				err = db.SaveSubName(name, string(j))
				if err != nil {
					log.Println("BatchMarketSubBuyFromSeller", "save to db error")
					continue
				}
				// update owner:namehash
				UpdateOwner(ev.From.String(), ev.To.String(), "snH_1_"+name, db)
			}

		}
	}
}

func UpdateOwner(from string, to string, namehash string, db *ldb.Ldb) {
	from_data, _ := db.GetOwnerNameHash(from)
	var newfrom []string
	for _, v := range from_data {
		if v != namehash {
			newfrom = append(newfrom, v)
		}
	}
	db.SavaOwnerNameHash(from, newfrom)

	to_data, _ := db.GetOwnerNameHash(to)
	to_data = append(to_data, namehash)
	db.SavaOwnerNameHash(to, to_data)

}

func NewOwner(addr string, namehash string, db *ldb.Ldb) {
	names, _ := db.GetOwnerNameHash(addr)
	names = append(names, namehash)
	db.SavaOwnerNameHash(addr, names)
}
