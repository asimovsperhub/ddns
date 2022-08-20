package dns

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	udidc "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv2"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func AddConfNameHash(conftype string, confval string, NameHash string, db *ldb.Ldb) {
	var err error
	confkey := strings.ToLower(conftype) + "_" + confval
	addrL, _ := db.GetConfNameHashList(confkey)

	if addrL == nil {
		err = db.SaveConfNameHashList(confkey, []string{NameHash})
		if err != nil {
			log.Println("SaveConfNameHashList", "save to db error")
		}
		log.Println("SaveConfNameHashList ", []string{NameHash})
	} else {
		addrL = append(addrL, NameHash)
		err = db.SaveAddressList(confkey, addrL)
		if err != nil {
			log.Println("SaveConfNameHashList", "save to db error")
		}
		log.Println("SaveConfNameHashList ", addrL)
	}
}

func UpdateConfNameHash(conftype string, newconfval string, oldconfval string, NameHash string, db *ldb.Ldb) {
	var err error
	confkey := strings.ToLower(conftype) + "_" + oldconfval
	addrL, _ := db.GetConfNameHashList(confkey)

	if addrL != nil {
		newaddrl := []string{}
		for _, val := range addrL {
			if val != NameHash {
				newaddrl = append(newaddrl, val)
			}
		}
		log.Println("SaveConfNameHashList ", newaddrl)
		err = db.SaveAddressList(confkey, newaddrl)
		if err != nil {
			log.Println("SaveConfNameHashList", "save to db error")
		}
	}
	AddConfNameHash(conftype, newconfval, NameHash, db)
}
func NewDnsConfClient(client *ethclient.Client) (*udidc.DnsConf, error) {
	var (
		dnsprice *udidc.DnsConf
		err      error
	)
	dnsprice, err = udidc.NewDnsConf(common.HexToAddress(config.GetRConf().Cconf.DnsConf), client)
	if err != nil {
		client.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, err
	}
	return dnsprice, nil
}
func BatchNewConf(start, end uint64) {
	var (
		cli     *ethclient.Client
		dnsconf *udidc.DnsConf
		err     error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
		return
	}
	defer cli.Close()
	dnsconf, err = NewDnsConfClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	db := ldb.GetLdb()
	b, e := start, end

	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	rootname := new(RootNameInfo)
	subname := new(SubNameInfo)
	var addmap *udidc.DnsConfEvAddMapIterator
	addmap, _ = dnsconf.FilterEvAddMap(op)
	if addmap != nil {
		defer addmap.Close()
	}
	for addmap.Next() {
		ev := addmap.Event
		nameHashStr := hex.EncodeToString(ev.NameHash[:])
		log.Println("BatchNewConf DnsConfEvAddMapIterator", nameHashStr, string(ev.TName), string(ev.Val))
		rootinfo, _ := db.GetRootName(nameHashStr)
		if rootinfo != "" {
			err = json.Unmarshal([]byte(rootinfo), rootname)
			if err != nil {
				log.Println("json Unmarshal ", err)
				continue
			} else {
				if rootname.Conf == nil {
					rootname.Conf = map[string]string{string(ev.TName): string(ev.Val)}
				} else {
					rootname.Conf[string(ev.TName)] = string(ev.Val)
				}
			}
			log.Println("BatchNewConf AddMap ", rootname.Conf)
			j, _ := json.Marshal(rootname)

			err = db.SaveRootName(rootname.Hash, string(j))
			if err != nil {
				fmt.Println("Conf root addmap", "save to db error")
				continue
			}

			// AddConfNameHash
			AddConfNameHash(strings.ToLower(string(ev.TName)), string(ev.Val), fmt.Sprintf("rnH_1_%s", nameHashStr), db)

		} else {
			subinfo, _ := db.GetSubName(nameHashStr)
			if subinfo != "" {
				err = json.Unmarshal([]byte(subinfo), subname)
				if err != nil {
					log.Println("json Unmarshal ", err)
					continue
				} else {
					if subname.Conf == nil {
						subname.Conf = map[string]string{string(ev.TName): string(ev.Val)}
					} else {
						subname.Conf[string(ev.TName)] = string(ev.Val)
					}
				}
			}
			j, _ := json.Marshal(subname)
			err = db.SaveSubName(subname.Hash, string(j))
			if err != nil {
				log.Println("Conf sub addmap", "save to db error")
				continue
			}

			AddConfNameHash(strings.ToLower(string(ev.TName)), string(ev.Val), fmt.Sprintf("snH_1_%s", nameHashStr), db)
		}
	}
	var updatemap *udidc.DnsConfEvUpdateMapIterator
	updatemap, _ = dnsconf.FilterEvUpdateMap(op)
	if updatemap != nil {
		defer updatemap.Close()
		for updatemap.Next() {
			ev := updatemap.Event
			fmt.Println(ev.NameHash, ev.TName, ev.Val)
			nameHashStr := hex.EncodeToString(ev.NameHash[:])
			rootinfo, _ := db.GetRootName(nameHashStr)
			if rootinfo != "" {
				err = json.Unmarshal([]byte(rootinfo), rootname)
				if err != nil {
					log.Println("json Unmarshal ", err)
					continue
				} else {
					UpdateConfNameHash(strings.ToLower(string(ev.TName)), string(ev.Val), rootname.Conf[string(ev.TName)], fmt.Sprintf("rnH_1_%s", nameHashStr), db)
					rootname.Conf[string(ev.TName)] = string(ev.Val)
				}
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					fmt.Println("Conf root updatemap", "save to db error")
					continue
				}
			} else {
				subinfo, _ := db.GetSubName(nameHashStr)
				if subinfo != "" {
					err = json.Unmarshal([]byte(subinfo), subname)
					if err != nil {
						log.Println("json Unmarshal ", err)
						continue
					} else {
						UpdateConfNameHash(strings.ToLower(string(ev.TName)), string(ev.Val), subname.Conf[string(ev.TName)], fmt.Sprintf("snH_1_%s", nameHashStr), db)
						subname.Conf[string(ev.TName)] = string(ev.Val)
					}
				}
				j, _ := json.Marshal(subname)
				err = db.SaveSubName(subname.Hash, string(j))
				if err != nil {
					fmt.Println("Conf sub updatemap", "save to db error")
					continue
				}
			}
		}
	}
	//var removemap *udidc.DnsConfRemoveMapIterator
	//removemap, _ = dnsconf.FilterRemoveMap(op)
	//if removemap != nil {
	//	defer removemap.Close()
	//	for removemap.Next() {
	//		ev := removemap.Event
	//		nameHashStr := hex.EncodeToString(ev.NameHash[:])
	//		rootinfo, _ := db.GetRootName(nameHashStr)
	//		if rootinfo != "" {
	//			err = json.Unmarshal([]byte(rootinfo), rootname)
	//			if err != nil {
	//				log.Println("json Unmarshal ", err)
	//				continue
	//			} else {
	//				delete(rootname.Conf, fmt.Sprintf("%s", string(ev.TName)))
	//			}
	//			j, _ := json.Marshal(rootname)
	//
	//			err = db.SaveRootName(rootname.Hash, string(j))
	//			if err != nil {
	//				fmt.Println("Conf root removemap", "save to db error")
	//				continue
	//			}
	//		} else {
	//			subinfo, _ := db.GetSubName(nameHashStr)
	//			if subinfo != "" {
	//				err = json.Unmarshal([]byte(subinfo), subname)
	//				if err != nil {
	//					log.Println("json Unmarshal ", err)
	//					continue
	//				} else {
	//					delete(subname.Conf, fmt.Sprintf("%s", string(ev.TName)))
	//				}
	//			}
	//			j, _ := json.Marshal(subname)
	//			err = db.SaveRootName(subname.Hash, string(j))
	//			if err != nil {
	//				fmt.Println("Conf sub removemap", "save to db error")
	//				continue
	//			}
	//		}
	//	}
	//}

}
