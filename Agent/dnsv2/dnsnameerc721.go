package dnsv2

import (
	"github.com/dnsdao/dnsdao.resolver/config"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func NewDnsNameErcClient() (*udidc22.DnsNameErc721, *ethclient.Client, error) {
	var (
		cli  *ethclient.Client
		root *udidc22.DnsNameErc721
		err  error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewDnsNameErcClient EthClient conn err", err)
		return nil, nil, err
	}
	root, err = udidc22.NewDnsNameErc721(common.HexToAddress(config.GetRConf().Cconf.DnsName), cli)
	if err != nil {
		cli.Close()
		log.Println("NewDnsNameErcClient NewDnsNameErc721  err", err)
		return nil, nil, err
	}
	return root, cli, nil
}
