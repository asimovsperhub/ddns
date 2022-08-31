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
	ExpireTime   *big.Int            `json:"expireTime"`
	TokenId      *big.Int            `json:"token_id"`
	Conf         map[string]string   `json:"conf"`
}

func NewRootClient() (*udidc.DnsName, *ethclient.Client, error) {
	var (
		cli  *ethclient.Client
		root *udidc.DnsName
		err  error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
		return nil, nil, err
	}
	root, err = udidc.NewDnsName(common.HexToAddress(config.GetRConf().Cconf.DnsName), cli)
	if err != nil {
		cli.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, nil, err
	}
	return root, cli, nil
}

func BatchNewRoot(start, end uint64) {
	var err error
	rootC, cli, err := NewRootClient()
	defer cli.Close()
	if err != nil {
		log.Println("BatchNewRoot", err)
		return
	}
	dnsowner, err := NewDnsOwnerClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	b, e := start, end

	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	db := ldb.GetLdb()
	// mint root域名
	var mintiter *udidc.DnsNameEvMintDnsNameIterator
	mintiter, err = rootC.FilterEvMintDnsName(op)
	if err != nil {
		log.Println("FilterEvMintDnsName", err)
		return
	}
	if mintiter != nil {
		defer mintiter.Close()
		for mintiter.Next() {
			ev := mintiter.Event
			nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
			nameHashStr := hex.EncodeToString(nameHash[:])
			owner, err1 := dnsowner.DnsOwners(nil, *nameHash)
			if err1 != nil {
				log.Println("BatchNewRoot DnsOwners", err1)
			}
			// log.Println(nameHashStr, ev.EntireName, owner.DnsOwner, owner.TokenId)
			nameStore, _ := rootC.NameStore(nil, *nameHash)
			rn := &RootNameInfo{
				Name:       ev.EntireName,
				Hash:       nameHashStr,
				TokenId:    owner.TokenId,
				Owner:      owner.DnsOwner,
				ExpireTime: nameStore.ExpireTime,
				OpenToReg:  nameStore.OpenToReg,
			}
			j, _ := json.Marshal(rn)

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
			// SaveContractTokenIdName
			contractkey := strings.ToLower(config.GetRConf().Cconf.DnsName)
			tokenIdName, _ := db.GetContractTokenIdName(contractkey)
			if tokenIdName == nil {
				tokenIdName = &ldb.ContractTokenIdName{
					TokenName: map[string]string{owner.TokenId.String(): ev.EntireName},
				}
			} else {
				tokenIdName.TokenName[owner.TokenId.String()] = ev.EntireName
			}
			log.Println("BatchNewRoot SaveContractTokenIdName ", owner.TokenId.String(), ev.EntireName)
			err = db.SaveContractTokenIdName(contractkey, tokenIdName)
			if err != nil {
				log.Println("BatchNewRoot SaveContractTokenIdName", "save to db error")
				continue
			}

		}
	}
	// 开启子域名注册记录 添加合约地址
	var newsubname *udidc.DnsNameEvNewSubNameIterator
	newsubname, err = rootC.FilterEvNewSubName(op)
	if err != nil {
		log.Println("FilterEvNewSubName", err)
		return
	}
	if newsubname != nil {
		defer newsubname.Close()
		for newsubname.Next() {
			ev := newsubname.Event
			// 添加contract--RootNameHash
			ca := db.SaveContractAddr(ev.Erc721Addr.String(), hex.EncodeToString(ev.Hash[:]))
			if ca != nil {
				log.Println("SaveContractAddr err", ev.Erc721Addr.String())
			}
			// rootName 添加合约地址
			ct, cterr := db.GetContractList()
			if cterr != nil {
				if strings.Contains(cterr.Error(), "not found") {
					saveerr := db.SaveContractList([]string{ev.Erc721Addr.String()})
					if saveerr != nil {
						log.Println("SaveContractList err", saveerr)
						continue
					}
				}
			} else {
				ct = append(ct, ev.Erc721Addr.String())
				saveerr := db.SaveContractList(ct)
				if saveerr != nil {
					log.Println("SaveContractList err", saveerr)
					continue
				}
			}
			var res string
			res, err = db.GetRootName(hex.EncodeToString(ev.Hash[:]))
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
				rootname.Contract = ev.Erc721Addr
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
	var chargedns *udidc.DnsNameEvChargeDnsNameIterator
	chargedns, err = rootC.FilterEvChargeDnsName(op)
	if chargedns != nil {
		defer chargedns.Close()
		for chargedns.Next() {
			ev := chargedns.Event
			nameStore, _ := rootC.NameStore(nil, ev.NameHash)
			var res string
			res, err = db.GetRootName(hex.EncodeToString(ev.NameHash[:]))
			if err != nil {
				log.Println("GetRootName ", err)
			}
			rootname := new(RootNameInfo)
			err = json.Unmarshal([]byte(res), rootname)
			if err != nil {
				log.Println("json Unmarshal ", err)
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
