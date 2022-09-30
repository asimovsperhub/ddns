package dnsv2

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/big"
	"strconv"
	"strings"
)

type SubNameInfo struct {
	Name       string              `json:"name"`
	Hash       string              `json:"hash"`
	FatherName string              `json:"father_name"`
	Contract   common.Address      `json:"contract"`
	SubName    map[string][]string `json:"sub_name"`
	Price      map[string]*big.Int `json:"price"`
	Owner      common.Address      `json:"owner"`
	ExpireTime uint32              `json:"expireTime"`
	TokenId    *big.Int            `json:"token_id"`
	Conf       map[string]string   `json:"conf"`
}

func GetLenPrice(Rootname *RootNameInfo, namelen int, year int) (*big.Int, float64) {
	log.Println("GetlenPrice namelen year", namelen, year)
	// name price
	priceL, _ := strconv.Atoi(Rootname.Price[fmt.Sprintf("%d", namelen)].String())
	if priceL != 0 {
		return Rootname.Price[fmt.Sprintf("%d", namelen)], (float64(priceL) / 1000000000000000000) * float64(year)
	} else {
		priceD, _ := strconv.Atoi(Rootname.Price["default"].String())
		return Rootname.Price["default"], (float64(priceD) / 1000000000000000000) * float64(year)
	}
}
func NewSubClient(contract string) (*udidc22.DnsSecondLevelName, *ethclient.Client, error) {
	var (
		cli *ethclient.Client
		sub *udidc22.DnsSecondLevelName
		err error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewSubClient EthClient conn err", err)
		return nil, nil, err
	}
	sub, err = udidc22.NewDnsSecondLevelName(common.HexToAddress(contract), cli)
	if err != nil {
		cli.Close()
		log.Println("NewSubClient NewDnsSecondLevelName  err", err)
		return nil, nil, err
	}
	return sub, cli, nil
}

func BatchNewSub(start, end uint64) {
	db := ldb.GetLdb()
	cli, erreth := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if erreth != nil {
		cli.Close()
		log.Println("BatchNewSub EthClient conn err", erreth)
	}
	defer cli.Close()
	var err error
	subC, _, err := NewSubClient(config.GetRConf().Cconf.SecondName)
	if err != nil {
		log.Println("NewSubClient", err)
		return
	}
	b, e := start, end

	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	//var mintiter *udidc22.DnsSecondLevelNameEvMintSecondLevelNameIterator
	//mintiter, err = subC.FilterEvMintSecondLevelName(op)
	//if err != nil {
	//	log.Println("FilterEvMintDnsName", err)
	//	return
	//}
	//if mintiter != nil {
	//	defer mintiter.Close()
	//	// mint 二级域名
	//	for mintiter.Next() {
	//		ev := mintiter.Event
	//		log.Println(ev.EntireName, ev.Year, ev.TokenId)
	//		nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
	//		nameHashStr := hex.EncodeToString(nameHash[:])
	//		c := strings.Split(ev.EntireName, ".")
	//		rootname := c[len(c)-1]
	//		rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte(rootname)))
	//		topnamehashbyte := byte32(crypto.Keccak256([]byte(rootname)))
	//		nameStore, _ := subC.NameStore(nil, *topnamehashbyte, *nameHash)
	//		rnf, _ := db.GetRootName(rootnamehash)
	//		Rootname := new(RootNameInfo)
	//		err = json.Unmarshal([]byte(rnf), Rootname)
	//		if err != nil {
	//			log.Println("json Unmarshal ", err)
	//			continue
	//		} else {
	//			if Rootname.SubNameCount == nil {
	//				Rootname.SubNameCount = map[string]int64{fmt.Sprint(len(c[0])): 1}
	//			} else {
	//				if len(c[0]) >= 8 {
	//					Rootname.SubNameCount["8"] += 1
	//				} else {
	//					Rootname.SubNameCount[fmt.Sprint(len(c[0]))] += 1
	//				}
	//			}
	//			if Rootname.SubName == nil {
	//				Rootname.SubName = map[string][]string{fmt.Sprint(len(c[0])): []string{ev.EntireName}}
	//			} else {
	//				if Rootname.SubName[fmt.Sprint(len(c[0]))] == nil {
	//					Rootname.SubName[fmt.Sprint(len(c[0]))] = []string{ev.EntireName}
	//				} else {
	//					Rootname.SubName[fmt.Sprint(len(c[0]))] = append(Rootname.SubName[fmt.Sprint(len(c[0]))], ev.EntireName)
	//				}
	//			}
	//			j, _ := json.Marshal(Rootname)
	//
	//			err = db.SaveRootName(Rootname.Hash, string(j))
	//			if err != nil {
	//				fmt.Println("ChargeDnsName update ExpireTime ", "save to db error")
	//				continue
	//			}
	//
	//		}
	//		log.Println("Rootname.Contract.String()", Rootname.Contract.String())
	//		dnsnameerc, _, errown := NewDnsNameErcClient(Rootname.Contract.String())
	//		if errown != nil {
	//			log.Println("NewDnsNameErcClient", errown)
	//			return
	//		}
	//		owner, _ := dnsnameerc.OwnerOf(nil, ev.TokenId)
	//		log.Println("Owner", owner)
	//		// coinbase:[namehash]
	//		addrnamehash := strings.ToLower(owner.String())
	//		addrL, _ := db.GetAddressList(addrnamehash)
	//		if addrL == nil {
	//			log.Println("BatchNewSub New Sub AddressList ", addrnamehash, ev.EntireName)
	//			err = db.SaveAddressList(addrnamehash, []string{fmt.Sprintf("snH_1_%s", nameHashStr)})
	//			if err != nil {
	//				fmt.Println("BatchNewRoot SaveAddressList", "save to db error")
	//				continue
	//			}
	//		} else {
	//			log.Println("BatchNewSub Append Sub AddressList ", addrnamehash, ev.EntireName)
	//			addrL = append(addrL, fmt.Sprintf("snH_1_%s", nameHashStr))
	//			err = db.SaveAddressList(addrnamehash, addrL)
	//			if err != nil {
	//				fmt.Println("BatchNewSub SaveAddressList", "save to db error")
	//				continue
	//			}
	//		}
	//
	//		// subname
	//		Subname := new(SubNameInfo)
	//		Subname.Name = ev.EntireName
	//		Subname.Hash = nameHashStr
	//		Subname.ExpireTime = nameStore.ExpireTime
	//		Subname.TokenId = ev.TokenId
	//		Subname.Contract = common.HexToAddress(Rootname.Contract.String())
	//		Subname.Owner = owner
	//		Subname.FatherName = rootname
	//		j, _ := json.Marshal(Subname)
	//		errsub := db.SaveSubName(Subname.Hash, string(j))
	//		if errsub != nil {
	//			log.Println("BatchNewSub save err", "save to db error")
	//			continue
	//		}
	//		log.Println("Register sub name ", string(j))
	//		// SaveContractTokenIdName
	//		contractkey := strings.ToLower(Rootname.Contract.String())
	//		tokenIdName, _ := db.GetContractTokenIdName(contractkey)
	//		if tokenIdName == nil {
	//			tokenIdName = &ldb.ContractTokenIdName{
	//				TokenName: map[string]string{ev.TokenId.String(): ev.EntireName},
	//			}
	//		} else {
	//			tokenIdName.TokenName[ev.TokenId.String()] = ev.EntireName
	//		}
	//		log.Println("BatchNewSub SaveContractTokenIdName ", ev.TokenId.String(), ev.EntireName)
	//		err = db.SaveContractTokenIdName(contractkey, tokenIdName)
	//		if err != nil {
	//			log.Println("BatchNewSub SaveContractTokenIdName", "save to db error")
	//			continue
	//		}
	//	}
	//}

	var mintitersign *udidc22.DnsSecondLevelNameEvMintSecondLevelNameBySigIterator
	mintitersign, err = subC.FilterEvMintSecondLevelNameBySig(op)
	if err != nil {
		log.Println("FilterEvMintDnsName", err)
		return
	}
	if mintitersign != nil {
		defer mintitersign.Close()
		// mint 二级域名
		for mintitersign.Next() {
			ev := mintitersign.Event
			log.Println(ev.EntireName, ev.Year, ev.TokenId)
			nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
			nameHashStr := hex.EncodeToString(nameHash[:])
			c := strings.Split(ev.EntireName, ".")
			rootname := c[len(c)-1]
			rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte(rootname)))
			topnamehashbyte := byte32(crypto.Keccak256([]byte(rootname)))
			nameStore, _ := subC.NameStore(nil, *topnamehashbyte, *nameHash)
			log.Println(nameStore.EntireName, nameStore.ExpireTime, nameStore.TokenId)
			rnf, _ := db.GetRootName(rootnamehash)
			Rootname := new(RootNameInfo)
			err = json.Unmarshal([]byte(rnf), Rootname)
			if err != nil {
				log.Println("json Unmarshal ", err)
				continue
			} else {
				if Rootname.SubNameCount == nil {
					Rootname.SubNameCount = map[string]int64{fmt.Sprint(len(c[0])): 1}
				} else {
					if len(c[0]) >= 8 {
						Rootname.SubNameCount["8"] += 1
					} else {
						Rootname.SubNameCount[fmt.Sprint(len(c[0]))] += 1
					}
				}
				if Rootname.SubName == nil {
					Rootname.SubName = map[string][]string{fmt.Sprint(len(c[0])): []string{ev.EntireName}}
				} else {
					if Rootname.SubName[fmt.Sprint(len(c[0]))] == nil {
						Rootname.SubName[fmt.Sprint(len(c[0]))] = []string{ev.EntireName}
					} else {
						Rootname.SubName[fmt.Sprint(len(c[0]))] = append(Rootname.SubName[fmt.Sprint(len(c[0]))], ev.EntireName)
					}
				}
				j, _ := json.Marshal(Rootname)

				err = db.SaveRootName(Rootname.Hash, string(j))
				if err != nil {
					fmt.Println("ChargeDnsName update ExpireTime ", "save to db error")
					continue
				}

			}
			log.Println("Rootname.Contract.String()", Rootname.Contract.String())
			dnsnameerc, _, errown := NewDnsNameErcClient(Rootname.Contract.String())
			if errown != nil {
				log.Println("NewDnsNameErcClient", errown)
				return
			}
			owner, _ := dnsnameerc.OwnerOf(nil, ev.TokenId)
			log.Println("Owner", owner)
			// coinbase:[namehash]
			addrnamehash := strings.ToLower(owner.String())
			addrL, _ := db.GetAddressList(addrnamehash)
			if addrL == nil {
				log.Println("BatchNewSub New Sub AddressList ", addrnamehash, ev.EntireName)
				err = db.SaveAddressList(addrnamehash, []string{fmt.Sprintf("snH_1_%s", nameHashStr)})
				if err != nil {
					fmt.Println("BatchNewRoot SaveAddressList", "save to db error")
					continue
				}
			} else {
				log.Println("BatchNewSub Append Sub AddressList ", addrnamehash, ev.EntireName)
				addrL = append(addrL, fmt.Sprintf("snH_1_%s", nameHashStr))
				err = db.SaveAddressList(addrnamehash, addrL)
				if err != nil {
					fmt.Println("BatchNewSub SaveAddressList", "save to db error")
					continue
				}
			}

			// subname
			Subname := new(SubNameInfo)
			Subname.Name = ev.EntireName
			Subname.Hash = nameHashStr
			Subname.ExpireTime = nameStore.ExpireTime
			Subname.TokenId = ev.TokenId
			Subname.Contract = common.HexToAddress(Rootname.Contract.String())
			Subname.Owner = owner
			Subname.FatherName = rootname
			j, _ := json.Marshal(Subname)
			errsub := db.SaveSubName(Subname.Hash, string(j))
			if errsub != nil {
				log.Println("BatchNewSub save err", "save to db error")
				continue
			}
			log.Println("Register sub name ", string(j))
			// SaveContractTokenIdName
			contractkey := strings.ToLower(Rootname.Contract.String())
			tokenIdName, _ := db.GetContractTokenIdName(contractkey)
			if tokenIdName == nil {
				tokenIdName = &ldb.ContractTokenIdName{
					TokenName: map[string]string{ev.TokenId.String(): ev.EntireName},
				}
			} else {
				tokenIdName.TokenName[ev.TokenId.String()] = ev.EntireName
			}
			log.Println("BatchNewSub SaveContractTokenIdName ", ev.TokenId.String(), ev.EntireName)
			err = db.SaveContractTokenIdName(contractkey, tokenIdName)
			if err != nil {
				log.Println("BatchNewSub SaveContractTokenIdName", "save to db error")
				continue
			}
			// 验证pass卡签名接口注册的域名
			filter := bson.M{"name": ev.EntireName, "msg_sender": strings.ToLower(owner.String())}
			opm := &options.FindOptions{Skip: func(i int64) *int64 { return &i }(0), Limit: func(i int64) *int64 { return &i }(1), Sort: bson.M{"_id": -1}}
			selectres, _ := ldb.SELECT("signcoll", filter, opm)
			if selectres == nil {
				log.Println("The signed interface has not passed. Procedure", ev.EntireName)
			} else {
				ldb.UPDATEOne("signcoll", filter, bson.M{"status": "2"})
			}
		}
	}

}
