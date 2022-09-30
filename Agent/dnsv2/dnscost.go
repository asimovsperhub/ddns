package dnsv2

import (
	"github.com/dnsdao/dnsdao.resolver/config"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

func NewCostClient() (*udidc22.DnsCost, *ethclient.Client, error) {
	var (
		cli  *ethclient.Client
		root *udidc22.DnsCost
		err  error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewCostClient EthClient conn err", err)
		return nil, nil, err
	}
	root, err = udidc22.NewDnsCost(common.HexToAddress(config.GetRConf().Cconf.Cost), cli)
	if err != nil {
		cli.Close()
		log.Println("NewCostClient  err", err)
		return nil, nil, err
	}
	return root, cli, nil
}

type Payment struct {
	db *mongo.Database
}

var (
	payment *Payment
	Once    sync.Once
)

func GetPayment() *Payment {
	if payment == nil {
		Once.Do(func() {
			//var err error

		})
	}

	return nil
}
