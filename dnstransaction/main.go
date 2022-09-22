package main

import (
	"fmt"
	udidc "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv2"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	cli, err := ethclient.Dial(ethacct.InfuraPoint)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var hashset map[[32]byte]int64
	hashset = make(map[[32]byte]int64)
	var dsn *udidc.DnsSubName
	dsn, err = udidc.NewDnsSubName(common.HexToAddress("0x111cFbc15FC56AD0aC8C7536BcA3347a558a0012"), cli)
	if err != nil {
		panic(err)
	}

	b, _ := dsn.PassCardUsed(nil, 1662)
	fmt.Println(b)

	for i := 0; i < 1500; i++ {
		id := big.NewInt(int64(i))
		hash, err1 := dsn.Id2Hash(nil, id)
		//var zeroHash [32]byte
		if err1 != nil || hash == [32]byte{0} {
			fmt.Println(err1)
			break
		}
		if _, ok := hashset[hash]; !ok {
			hashset[hash] = int64(i)
		} else {
			var n struct {
				EntireName []byte
				ExpireTime *big.Int
				TokenId    *big.Int
			}
			n, _ = dsn.NameStore(nil, hash)
			a, _ := dsn.OwnerOf(nil, id)
			b, _ := dsn.OwnerOf(nil, big.NewInt(int64(hashset[hash])))
			//fmt.Println("0x"+hex.EncodeToString(hash[:]), hashset[hash], i, string(n.EntireName), n.TokenId.String())
			fmt.Println(hashset[hash], b.String(), i, a.String(), string(n.EntireName), n.TokenId.String())
		}
	}
}
