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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type SubNameInfo struct {
	Name       string              `json:"name"`
	Hash       string              `json:"hash"`
	FatherName string              `json:"father_name"`
	Contract   common.Address      `json:"contract"`
	SubName    map[string][]string `json:"sub_name"`
	Price      map[string]*big.Int `json:"price"`
	Owner      common.Address      `json:"owner"`
	ExpireTime *big.Int            `json:"expireTime"`
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
func NewSubClient(contract string) (*udidc.DnsSubName, *ethclient.Client, error) {
	var (
		cli *ethclient.Client
		sub *udidc.DnsSubName
		err error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
		return nil, nil, err
	}
	sub, err = udidc.NewDnsSubName(common.HexToAddress(contract), cli)
	if err != nil {
		cli.Close()
		log.Println("NewRootClient NewDnsName  err", err)
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
	contractL, _ := db.GetContractList()
	dnsowner, errown := NewDnsOwnerClient(cli)
	if errown != nil {
		log.Println("NewDnsOwnerClient", errown)
		return
	}
	for _, contract := range contractL {
		var err error
		subC, _, err := NewSubClient(contract)
		defer cli.Close()
		if err != nil {
			log.Println("BatchNewSub", err)
			return
		}
		b, e := start, end

		op := &bind.FilterOpts{
			Start:   b,
			End:     &e,
			Context: context.TODO(),
		}
		var mintiter *udidc.DnsSubNameEvMintSubNameIterator
		mintiter, err = subC.FilterEvMintSubName(op)
		if err != nil {
			log.Println("FilterEvMintDnsName", err)
			return
		}
		if mintiter != nil {
			defer mintiter.Close()
			// mint 二级域名
			for mintiter.Next() {
				ev := mintiter.Event
				nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
				nameHashStr := hex.EncodeToString(nameHash[:])
				owner, _ := dnsowner.DnsOwners(nil, *nameHash)
				nameStore, _ := subC.NameStore(nil, *nameHash)
				c := strings.Split(ev.EntireName, ".")
				rootname := c[len(c)-1]
				rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte(rootname)))

				rnf, _ := db.GetRootName(rootnamehash)
				if len(rnf) > 0 {
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
					// root收益
					addrkey := strings.ToLower(Rootname.Owner.String())
					earnings, _ := db.GetRootEarnings(addrkey)
					// uint8 不会超过int
					price, earn := GetLenPrice(Rootname, len(c[0]), int(ev.Year))
					log.Println("price,earn", price, earn)
					if earnings != nil {
						earnings.Earnings += earn
						rootearn := earnings.RootNameMap[Rootname.Name]
						if rootearn != nil {
							rootearn.Earnings += earn
							rootearn.Details = append(rootearn.Details, &ldb.RootEarningsDetails{Price: price, Contract: Rootname.Contract, RegisterName: ev.EntireName,
								RegisterOwner: owner.DnsOwner, RegisterTime: time.Now().Unix()})
						} else {
							earnings.RootNameMap[Rootname.Name] = &ldb.RootEarnings{earn, []*ldb.RootEarningsDetails{&ldb.RootEarningsDetails{Price: price, Contract: Rootname.Contract, RegisterName: ev.EntireName,
								RegisterOwner: owner.DnsOwner, RegisterTime: time.Now().Unix()}}}
						}
						errsave := db.SaveRootEarnings(addrkey, earnings)
						log.Println("BatchNewSub SaveRootEarnings", addrkey, earnings)
						if errsave != nil {
							log.Println("BatchNewSub SaveRootEarnings err", errsave)
						}
					} else {
						earningsN := new(ldb.AddressEarnings)
						earningsN.Earnings = earn
						earningsN.RootNameMap = map[string]*ldb.RootEarnings{rootname: &ldb.RootEarnings{Earnings: earn, Details: []*ldb.RootEarningsDetails{&ldb.RootEarningsDetails{Price: price, Contract: Rootname.Contract, RegisterName: ev.EntireName,
							RegisterOwner: owner.DnsOwner, RegisterTime: time.Now().Unix()}}}}
						errsave := db.SaveRootEarnings(addrkey, earningsN)
						log.Println("BatchNewSub SaveRootEarnings", addrkey, earningsN)
						if errsave != nil {
							log.Println("SaveRootEarnings err", errsave)
						}
					}

					// coinbase:[namehash]
					addrnamehash := strings.ToLower(Rootname.Owner.String())
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
							fmt.Println("BatchNewRoot SaveAddressList", "save to db error")
							continue
						}
					}

				}

				// subname
				Subname := new(SubNameInfo)
				Subname.Name = ev.EntireName
				Subname.Hash = nameHashStr
				Subname.ExpireTime = nameStore.ExpireTime
				Subname.TokenId = owner.TokenId
				Subname.Contract = common.HexToAddress(contract)
				Subname.Owner = owner.DnsOwner
				Subname.FatherName = rootname
				j, _ := json.Marshal(Subname)
				errsub := db.SaveSubName(Subname.Hash, string(j))
				if errsub != nil {
					log.Println("BatchNewSub save err", "save to db error")
					continue
				}
				// SaveContractTokenIdName
				contractkey := strings.ToLower(contract)
				tokenIdName, _ := db.GetContractTokenIdName(contractkey)
				if tokenIdName == nil {
					tokenIdName = &ldb.ContractTokenIdName{
						TokenName: map[string]string{owner.TokenId.String(): ev.EntireName},
					}
				} else {
					tokenIdName.TokenName[owner.TokenId.String()] = ev.EntireName
				}
				log.Println("BatchNewSub SaveContractTokenIdName ", owner.TokenId.String(), ev.EntireName)
				err = db.SaveContractTokenIdName(contractkey, tokenIdName)
				if err != nil {
					log.Println("BatchNewSub SaveContractTokenIdName", "save to db error")
					continue
				}
			}
		}
		//var newsubname *udidc.DnsSubNameEvAddSubNameIterator
		//newsubname, err = subC.FilterEvAddSubName(op)
		//if err != nil {
		//	log.Println("FilterEvAddSubName", err)
		//	return
		//}
		//defer newsubname.Close()
		//// 添加域名三级及以上域名
		//for newsubname.Next() {
		//	ev := newsubname.Event
		//	fmt.Println(ev.Hash, ev.EntireName)
		//}
	}

}
