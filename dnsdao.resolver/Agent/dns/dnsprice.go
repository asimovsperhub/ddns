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
	"math/big"
)

func NewDnsPriceClient(client *ethclient.Client) (*udidc.DnsPrice, error) {
	var (
		dnsprice *udidc.DnsPrice
		err      error
	)
	dnsprice, err = udidc.NewDnsPrice(common.HexToAddress(config.GetRConf().Cconf.DnsPrice), client)
	if err != nil {
		client.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, err
	}
	return dnsprice, nil
}

func BatchNewPrice(start, end uint64) {
	var (
		cli      *ethclient.Client
		dnsprice *udidc.DnsPrice
		err      error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsprice, err = NewDnsPriceClient(cli)
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
	// 默认价格
	var priceadditer *udidc.DnsPriceEvAddPriceIterator
	priceadditer, err = dnsprice.FilterEvAddPrice(op)
	if priceadditer != nil {
		defer priceadditer.Close()
		for priceadditer.Next() {
			ev := priceadditer.Event
			// fmt.Println("add default", hex.EncodeToString(ev.NameHash[:]), ev.DefaultPrice)
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
				rootname.Price = map[string]*big.Int{"default": ev.DefaultPrice}
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					fmt.Println("BatchNewPrice add", "save to db error")
					continue
				} else {
					log.Println("add DefaultPrice", rootname.Name, rootname.Price)
				}
			}
		}

	}
	// 长度 价格
	var priceaddleniter *udidc.DnsPriceEvAddPriceLenIterator
	priceaddleniter, err = dnsprice.FilterEvAddPriceLen(op)
	if priceaddleniter != nil {
		defer priceaddleniter.Close()
		for priceaddleniter.Next() {
			ev := priceaddleniter.Event
			//fmt.Println("add len ", hex.EncodeToString(ev.NameHash[:]), ev.Len, ev.Price)
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
				if rootname.Price == nil {
					rootname.Price = map[string]*big.Int{fmt.Sprintf("%d", ev.Len): ev.Price}
				}
				rootname.Price[fmt.Sprintf("%d", ev.Len)] = ev.Price
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					fmt.Println("BatchNewPrice add", "save to db error")
					continue
				} else {
					log.Println("add LenPrice", rootname.Name, rootname.Price)
				}
			}
		}

	}
	// 数组添加
	var priceaddlenarriter *udidc.DnsPriceEvAddPriceLenArrayIterator
	priceaddlenarriter, err = dnsprice.FilterEvAddPriceLenArray(op)
	if priceaddlenarriter != nil {
		defer priceaddlenarriter.Close()
		for priceaddlenarriter.Next() {
			ev := priceaddlenarriter.Event
			// fmt.Println("add len arr", hex.EncodeToString(ev.NameHash[:]), ev.Len, ev.Price, ev.DefaultPrice)
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
				if rootname.Price == nil {
					rootname.Price = map[string]*big.Int{"default": ev.DefaultPrice}
				}
				rootname.Price["defaultp"] = ev.DefaultPrice
				rootname.Price[fmt.Sprintf("%d", ev.Len)] = ev.Price
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					log.Println("BatchNewPrice priceaddlenarriter", "save to db error")
					continue
				}
				log.Println("add Arr Add Price", rootname.Name, rootname.Price)

			}
		}
	}
	// 删除某个长度的价格
	var delprice *udidc.DnsPriceEvDelPriceLenIterator
	delprice, err = dnsprice.FilterEvDelPriceLen(op)
	if delprice != nil {
		defer delprice.Close()
		for delprice.Next() {
			ev := delprice.Event
			// fmt.Println("del", hex.EncodeToString(ev.NameHash[:]), ev.Len)
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
				delete(rootname.Price, fmt.Sprintf("%d", ev.Len))
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					fmt.Println("BatchNewPrice delprice", "save to db error")
					continue
				}
			}
		}
	}
	// 移除价格
	var removeprice *udidc.DnsPriceEvRemovePriceIterator
	removeprice, err = dnsprice.FilterEvRemovePrice(op)
	if removeprice != nil {
		defer removeprice.Close()
		for removeprice.Next() {
			ev := removeprice.Event
			//fmt.Println("remove", hex.EncodeToString(ev.NameHash[:]))
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
				rootname.Price = nil
				j, _ := json.Marshal(rootname)

				err = db.SaveRootName(rootname.Hash, string(j))
				if err != nil {
					fmt.Println("BatchNewPrice removeprice", "save to db error")
					continue
				}
			}
		}
	}

}
