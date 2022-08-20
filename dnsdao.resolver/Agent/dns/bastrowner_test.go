package dns

import (
	"encoding/hex"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestBASEOC_GetOwnerInfo(t *testing.T) {
	nameh := "0x7dd481eb4b63b94bb55e6b98aabb06c3b8484f82a4d656d6bca0b0cf9b446be0"

	h := common.HexToHash(nameh)

	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	eo, err1 := newBasTRClient(cli)
	if err1 != nil {
		panic(err)
	}

	var oi *OwnerInfo
	oi, err = eo.GetOwnerInfo(h)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(oi.Owner[:]))
	fmt.Println(oi.ExpireTime)

}

func TestBASEOC_GetOwnerInfo1(t *testing.T) {
	nameh := "abc.com"

	h := crypto.Keccak256Hash([]byte(nameh))
	fmt.Println(hex.EncodeToString(h[:]))

	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	eo, err1 := newBasTRClient(cli)
	if err1 != nil {
		panic(err)
	}

	var oi *OwnerInfo
	oi, err = eo.GetOwnerInfo(h)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(oi.Owner[:]))
	fmt.Println(oi.ExpireTime)

}
