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

func BatchDomainConfChanged(start, end uint64) {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		log.Println("BatchDomainConfChanged EthClient conn err", err)
		return
	}
	cfg, err := dnscontract.NewBasDomainConf(common.HexToAddress(config.GetRConf().Cconf.BasDomainConf), cli)
	if err != nil {
		log.Println("BatchDomainConfChanged NewBasDomainConf", err)
		return
	}
	root, err := dnscontract.NewBasRootDomain(common.HexToAddress(config.GetRConf().Cconf.BasRootContract), cli)
	if err != nil {
		log.Println("BatchDomainConfChanged NewBasRootDomain", err)
		return
	}
	sub, err := dnscontract.NewBasSubDomain(common.HexToAddress(config.GetRConf().Cconf.BasSubContract), cli)
	if err != nil {
		log.Println("BatchDomainConfChanged NewBasSubDomain", err)
		return
	}
	defer cli.Close()
	b, e := start, end
	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	var cfgitem *dnscontract.BasDomainConfDomainConfChangedIterator
	cfgitem, err = cfg.FilterDomainConfChanged(op)
	if err != nil {
		log.Println("BatchDomainConfChanged FilterDomainConfChanged", err)
		return
	}
	db := ldb.GetLdb()
	for cfgitem.Next() {
		ev := cfgitem.Event
		// parsing root domain
		root_info, err := root.Root(nil, ev.NameHash)
		if err != nil {
			log.Println("BatchDomainConfChanged  parsing root domain ", err)
		}
		if len(root_info.RootName) != 0 {
			name := hex.EncodeToString(ev.NameHash[:])
			val, _ := db.GetRootName(name)
			rootc := new(RootName)
			json.Unmarshal([]byte(val), rootc)
			rootc.Conf[ev.TypeName] = ev.Data
			j, _ := json.Marshal(rootc)
			err = db.SaveRootName(name, string(j))
			if err != nil {
				log.Println("BatchDomainConfChanged", "save to db error")
				continue
			}
			// add new/updata domainconf    domaintype:namehash
			err = db.SaveTypeNameHash(ev.TypeName+"_"+string(ev.Data), name, "root")
			if err != nil {
				log.Println("BatchDomainConfChanged SaveTypeNameHash error", err)
				continue
			}

		} else {
			// parsing sub domain
			info, err := sub.Sub(nil, ev.NameHash)
			if err != nil {
				log.Println("BatchDomainConfChanged parsing sub err ", err)
			}
			namehash := crypto.Keccak256(info.TotalName)
			name := hex.EncodeToString(namehash)
			val, _ := db.GetSubName(name)
			subc := new(SubName)
			json.Unmarshal([]byte(val), subc)
			subc.Conf[ev.TypeName] = ev.Data
			j, _ := json.Marshal(subc)
			err = db.SaveSubName(name, string(j))
			if err != nil {
				log.Println("BatchDomainConfChanged", "save to db error")
				continue
			}
			// add new/update domainconf    domaintype:namehash
			err = db.SaveTypeNameHash(ev.TypeName+"_"+string(ev.Data), name, "sub")
			if err != nil {
				log.Println("BatchDomainConfChanged SaveTypeNameHash error", err)
				continue
			}

		}

	}

}
