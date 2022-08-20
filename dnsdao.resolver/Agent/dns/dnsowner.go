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

type AddressNameHashList struct {
	NameHashList []string `json:"name_hash_list"`
}

func NewDnsOwnerClient(client *ethclient.Client) (*udidc.DnsOwners, error) {
	var (
		dnsowner *udidc.DnsOwners
		err      error
	)
	dnsowner, err = udidc.NewDnsOwners(common.HexToAddress(config.GetRConf().Cconf.DnsOwner), client)
	if err != nil {
		client.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, err
	}
	return dnsowner, nil
}
func BatchNewOwner(start, end uint64) {
	var (
		cli      *ethclient.Client
		dnsowner *udidc.DnsOwners
		err      error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewDnsOwnerClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsowner, err = NewDnsOwnerClient(cli)
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
	var updateowner *udidc.DnsOwnersEvUpdateOwnerIterator
	updateowner, _ = dnsowner.FilterEvUpdateOwner(op)
	if updateowner != nil {
		defer updateowner.Close()
		for updateowner.Next() {
			ev := updateowner.Event
			nameHashStr := hex.EncodeToString(ev.NameHash[:])
			log.Println("BatchNewOwner Update", nameHashStr, ev.NameOwner)
			rootinfo, _ := db.GetRootName(nameHashStr)
			// 是否root的更新
			if rootinfo != "" {
				err = json.Unmarshal([]byte(rootinfo), rootname)
				if err != nil {
					log.Println("json Unmarshal ", err)
					continue
				} else {
					// 当owner改变的时候去更新
					log.Println("BatchNewOwner Root Owner is ", rootname.Owner, ev.NameOwner)
					if rootname.Owner != ev.NameOwner {
						//  更新old owner 数据
						addressL, _ := db.GetAddressList(strings.ToLower(rootname.Owner.String()))
						if addressL != nil {
							newNameHashList := []string{}
							for _, data := range addressL {
								if data != fmt.Sprintf("rnH_1_%s", hex.EncodeToString(ev.NameHash[:])) {
									newNameHashList = append(newNameHashList, data)
								}
							}
							db.SaveAddressList(strings.ToLower(rootname.Owner.String()), newNameHashList)
							log.Println("BatchNewOwner Root Update Old Owner ", rootname.Owner, newNameHashList)
						}
						// 更新new owner数据
						addrkey := strings.ToLower(ev.NameOwner.String())
						addressN, _ := db.GetAddressList(addrkey)
						if addressN == nil {
							db.SaveAddressList(addrkey, []string{fmt.Sprintf("rnH_1_%s", hex.EncodeToString(ev.NameHash[:]))})
							log.Println("BatchNewOwner Root Update New Owner", ev.NameOwner, []string{fmt.Sprintf("rnH_1_%s", hex.EncodeToString(ev.NameHash[:]))})
						} else {
							addressN = append(addressN, fmt.Sprintf("rnH_1_%s", hex.EncodeToString(ev.NameHash[:])))
							db.SaveAddressList(addrkey, addressN)
							log.Println("BatchNewOwner Root Update New Owner", addressN)
						}
						// 更新root name info
						rootname.Owner = ev.NameOwner
						j, _ := json.Marshal(rootname)

						err = db.SaveRootName(rootname.Hash, string(j))
						if err != nil {
							log.Println("updateowner ", "save to db error")
							continue
						}
					}

				}

			} else {
				subinfo, _ := db.GetSubName(nameHashStr)
				if subinfo != "" {
					err = json.Unmarshal([]byte(subinfo), subname)
					if err != nil {
						log.Println("json Unmarshal ", err)
						continue
					} else {
						log.Println("BatchNewOwner SubName Owner is ", rootname.Owner, ev.NameOwner)
						// Old
						if subname.Owner != ev.NameOwner {
							addressL, _ := db.GetAddressList(strings.ToLower(subname.Owner.String()))
							if addressL != nil {
								newNameHashList := []string{}
								for _, data := range addressL {
									if data != fmt.Sprintf("snH_1_%s", hex.EncodeToString(ev.NameHash[:])) {
										newNameHashList = append(newNameHashList, data)
									}
								}
								db.SaveAddressList(strings.ToLower(subname.Owner.String()), newNameHashList)
							}
							// New
							addrkey := strings.ToLower(ev.NameOwner.String())
							addressN, _ := db.GetAddressList(addrkey)
							if addressN == nil {
								db.SaveAddressList(addrkey, []string{fmt.Sprintf("snH_1_%s", hex.EncodeToString(ev.NameHash[:]))})
							} else {
								addressN = append(addressN, fmt.Sprintf("snH_1_%s", hex.EncodeToString(ev.NameHash[:])))
								db.SaveAddressList(addrkey, addressN)
							}

							// 更新subname
							subname.Owner = ev.NameOwner
							j, _ := json.Marshal(subname)
							err = db.SaveSubName(subname.Hash, string(j))
							if err != nil {
								fmt.Println("updateowner", "save to db error")
								continue
							}
						}
					}
				}
			}
		}
	}

}
