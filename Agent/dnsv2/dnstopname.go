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
	"strings"
	"unsafe"
)

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}

type RootNameInfo struct {
	Name         string              `json:"name"`
	Hash         string              `json:"hash"`
	Contract     common.Address      `json:"contract"`
	SubName      map[string][]string `json:"sub_name"`
	SubNameCount map[string]int64    `json:"sub_name_count"`
	Price        map[string]*big.Int `json:"price"`
	Owner        common.Address      `json:"owner"`
	OpenToReg    bool                `json:"open_to_reg"`
	ExpireTime   uint32              `json:"expireTime"`
	TokenId      *big.Int            `json:"token_id"`
	Conf         map[string]string   `json:"conf"`
}

func NewRootClient() (*udidc22.DnsTopLevelName, *ethclient.Client, error) {
	var (
		cli  *ethclient.Client
		root *udidc22.DnsTopLevelName
		err  error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
		return nil, nil, err
	}
	root, err = udidc22.NewDnsTopLevelName(common.HexToAddress(config.GetRConf().Cconf.TopName), cli)
	if err != nil {
		cli.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, nil, err
	}
	return root, cli, nil
}

func BatchNewRoot(start, end uint64) {
	var err error
	rootC, _, err := NewRootClient()
	if err != nil {
		log.Println("BatchNewRoot", err)
		return
	}
	erc721, err := rootC.Erc721Store(nil, common.Hash{})
	if err != nil {
		log.Println(err)
	}
	dnsnameerc, _, err := NewDnsNameErcClient(erc721.String())
	if err != nil {
		log.Println("NewDnsNameErcClient", err)
		return
	}
	//defer cli.Close()
	b, e := start, end
	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	db := ldb.GetLdb()
	var mintiter *udidc22.DnsTopLevelNameEvMintTopLevelNameIterator
	mintiter, err = rootC.FilterEvMintTopLevelName(op)
	if err != nil {
		log.Println("FilterEvMintDnsName", err)
		return
	}
	if mintiter != nil {
		defer mintiter.Close()
		for mintiter.Next() {
			ev := mintiter.Event
			log.Println(ev.EntireName, ev.TokenId, ev.Year)
			nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
			nameHashStr := hex.EncodeToString(nameHash[:])
			owner, _ := dnsnameerc.OwnerOf(nil, ev.TokenId)
			nameStore, _ := rootC.NameStore(nil, *nameHash)
			rn := &RootNameInfo{
				Name:       ev.EntireName,
				Hash:       nameHashStr,
				TokenId:    ev.TokenId,
				Owner:      owner,
				ExpireTime: nameStore.ExpireTime,
			}
			j, _ := json.Marshal(rn)
			log.Println("register rootname ", string(j))
			err = db.SaveRootName(rn.Hash, string(j))
			if err != nil {
				log.Println("BatchNewRoot", "save to db error")
				continue
			}
			// coinbase:[namehash]
			addrkey := strings.ToLower(rn.Owner.String())
			addrL, _ := db.GetAddressList(addrkey)
			if addrL == nil {
				log.Println("BatchNewRoot New Root AddressList ", addrL)
				err = db.SaveAddressList(addrkey, []string{fmt.Sprintf("rnH_1_%s", rn.Hash)})
				if err != nil {
					log.Println("BatchNewRoot SaveAddressList", "save to db error")
					continue
				}
			} else {
				addrL = append(addrL, fmt.Sprintf("rnH_1_%s", rn.Hash))
				log.Println("BatchNewRoot Append Root AddressList ", addrL)
				err = db.SaveAddressList(addrkey, addrL)
				if err != nil {
					log.Println("BatchNewRoot SaveAddressList", "save to db error")
					continue
				}
			}
			// SaveContractTokenIdName 给open sea接口提供的 合约地址:{tokenid:name}
			contractkey := strings.ToLower(config.GetRConf().Cconf.DnsName)
			tokenIdName, _ := db.GetContractTokenIdName(contractkey)
			if tokenIdName == nil {
				tokenIdName = &ldb.ContractTokenIdName{
					TokenName: map[string]string{ev.TokenId.String(): ev.EntireName},
				}
			} else {
				tokenIdName.TokenName[ev.TokenId.String()] = ev.EntireName
			}
			log.Println("BatchNewRoot SaveContractTokenIdName ", ev.TokenId.String(), ev.EntireName)
			err = db.SaveContractTokenIdName(contractkey, tokenIdName)
			if err != nil {
				log.Println("BatchNewRoot SaveContractTokenIdName", "save to db error")
				continue
			}
		}
	}

	var mintitersig *udidc22.DnsTopLevelNameEvMintTopLevelNameBySigIterator
	mintitersig, err = rootC.FilterEvMintTopLevelNameBySig(op)
	if err != nil {
		log.Println("FilterEvMintDnsName", err)
		return
	}
	if mintitersig != nil {
		defer mintiter.Close()
		for mintitersig.Next() {
			ev := mintitersig.Event
			log.Println(ev.EntireName, ev.TokenId, ev.Year)
			nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
			nameHashStr := hex.EncodeToString(nameHash[:])
			owner, _ := dnsnameerc.OwnerOf(nil, ev.TokenId)
			nameStore, _ := rootC.NameStore(nil, *nameHash)
			rn := &RootNameInfo{
				Name:       ev.EntireName,
				Hash:       nameHashStr,
				TokenId:    ev.TokenId,
				Owner:      owner,
				ExpireTime: nameStore.ExpireTime,
			}
			j, _ := json.Marshal(rn)
			log.Println("register rootname ", string(j))
			err = db.SaveRootName(rn.Hash, string(j))
			if err != nil {
				log.Println("BatchNewRoot", "save to db error")
				continue
			}
			// coinbase:[namehash]
			addrkey := strings.ToLower(rn.Owner.String())
			addrL, _ := db.GetAddressList(addrkey)
			if addrL == nil {
				log.Println("BatchNewRoot New Root AddressList ", addrL)
				err = db.SaveAddressList(addrkey, []string{fmt.Sprintf("rnH_1_%s", rn.Hash)})
				if err != nil {
					log.Println("BatchNewRoot SaveAddressList", "save to db error")
					continue
				}
			} else {
				addrL = append(addrL, fmt.Sprintf("rnH_1_%s", rn.Hash))
				log.Println("BatchNewRoot Append Root AddressList ", addrL)
				err = db.SaveAddressList(addrkey, addrL)
				if err != nil {
					log.Println("BatchNewRoot SaveAddressList", "save to db error")
					continue
				}
			}
			// SaveContractTokenIdName 给open sea接口提供的 合约地址:{tokenid:name}
			contractkey := strings.ToLower(config.GetRConf().Cconf.DnsName)
			tokenIdName, _ := db.GetContractTokenIdName(contractkey)
			if tokenIdName == nil {
				tokenIdName = &ldb.ContractTokenIdName{
					TokenName: map[string]string{ev.TokenId.String(): ev.EntireName},
				}
			} else {
				tokenIdName.TokenName[ev.TokenId.String()] = ev.EntireName
			}
			log.Println("BatchNewRoot SaveContractTokenIdName ", ev.TokenId.String(), ev.EntireName)
			err = db.SaveContractTokenIdName(contractkey, tokenIdName)
			if err != nil {
				log.Println("BatchNewRoot SaveContractTokenIdName", "save to db error")
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

	// 开启子域名注册记录 添加合约地址
	var openreg *udidc22.DnsTopLevelNameEvOpen2RegIterator
	openreg, err = rootC.FilterEvOpen2Reg(op)
	if err != nil {
		log.Println("FilterEvOpen2Reg err", err)
		return
	}
	if openreg != nil {
		defer openreg.Close()
		for openreg.Next() {
			ev := openreg.Event
			// 添加contract--RootNameHash
			contract, _ := rootC.Erc721Store(nil, ev.NameHash)
			ca := db.SaveContractAddr(contract.String(), hex.EncodeToString(ev.NameHash[:]))
			if ca != nil {
				log.Println("SaveContractAddr err", contract.String())
			}
			// 添加开启注册的合约地址列表
			ct, cterr := db.GetContractList()
			if cterr != nil {
				if strings.Contains(cterr.Error(), "not found") {
					saveerr := db.SaveContractList([]string{contract.String()})
					if saveerr != nil {
						log.Println("SaveContractList err", saveerr)
						continue
					}
				} else {
					log.Println("GetContractList err", cterr)
					continue
				}
			} else {
				ct = append(ct, contract.String())
				saveerr := db.SaveContractList(ct)
				if saveerr != nil {
					log.Println("SaveContractList err", saveerr)
					continue
				}
			}
			// rootName 添加合约地址
			var res string
			res, err = db.GetRootName(hex.EncodeToString(ev.NameHash[:]))
			if err != nil {
				log.Println("GetRootName ", err)
				continue
			}
			rootname := new(RootNameInfo)
			err = json.Unmarshal([]byte(res), rootname)
			if err != nil {
				log.Println("json Unmarshal ", err)
				continue
			} else {
				rootname.Contract = contract
				log.Println("Open2Reg   Name------>Contract", rootname.Name, contract)
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					log.Println("NewSubName add contract ", "save to db error")
					continue
				}

			}

		}

	}
	// ChargeDnsName 更新域名到期时间
	var chargedns *udidc22.DnsTopLevelNameEVChargeTopLevelNameIterator
	chargedns, err = rootC.FilterEVChargeTopLevelName(op)
	if chargedns != nil {
		defer chargedns.Close()
		for chargedns.Next() {
			ev := chargedns.Event
			nameStore, _ := rootC.NameStore(nil, ev.NameHash)
			var res string
			res, err = db.GetRootName(hex.EncodeToString(ev.NameHash[:]))
			if err != nil {
				log.Println("GetRootName err ", err)
			}
			rootname := new(RootNameInfo)
			err = json.Unmarshal([]byte(res), rootname)
			if err != nil {
				log.Println("json Unmarshal err ", err)
				continue
			} else {
				rootname.ExpireTime = nameStore.ExpireTime
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					log.Println("ChargeDnsName update ExpireTime ", "save to db error")
					continue
				}
			}
		}
	}

}
