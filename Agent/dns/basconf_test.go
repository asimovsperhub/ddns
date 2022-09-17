package dns

import (
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestBASConfC_GetTypeNames(t *testing.T) {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	cfg, err1 := newBasConfClient(cli)
	if err1 != nil {
		panic(err)
	}

	ts, err2 := cfg.GetTypeNames()
	if err2 != nil {
		panic(err)
	}

	fmt.Println(ts)
}

func TestBASConfC_GetDnsConf(t *testing.T) {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	cfg, err1 := newBasConfClient(cli)
	if err1 != nil {
		panic(err)
	}

	ts, err2 := cfg.GetTypeNames()
	if err2 != nil {
		panic(err)
	}

	fmt.Println(ts)

	h := "0x396a90cccb4f6e0685a47ab9969963c145b1763812e807783f685797854f6da7"

	hash := common.HexToHash(h)

	m := cfg.GetDnsConf(hash, ts)

	type MapJson struct {
		A string
		B map[string][]byte
	}

	var mj MapJson

	mj.A = "aa"
	mj.B = m

	j, _ := json.Marshal(mj)

	fmt.Println(m)

	fmt.Println(string(j))
}
