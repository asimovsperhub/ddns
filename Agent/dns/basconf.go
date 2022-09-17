package dns

import (
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/contract/dnscontract"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func newBasConfClient(cli *ethclient.Client) (*BASConfC, error) {

	cfg, err := dnscontract.NewBasDomainConf(common.HexToAddress(config.GetRConf().Cconf.BasDomainConf), cli)
	if err != nil {
		cli.Close()
		return nil, err
	}
	return &BASConfC{cfg: cfg}, nil
}

type BASConfC struct {
	cfg *dnscontract.BasDomainConf
}

func (cc *BASConfC) GetDnsConf(nameHash [32]byte, ts []string) map[string][]byte {

	ret := make(map[string][]byte)

	for i := 0; i < len(ts); i++ {
		r, err := cc.cfg.Query(nil, nameHash, ts[i])
		if err != nil {
			log.Println("GetDnsConf", err)
			continue
		}

		ret[ts[i]] = r
	}

	return ret
}

func (cc *BASConfC) GetTypeNames() ([]string, error) {
	var ts []string
	i := int64(0)
	for {
		n := big.NewInt(i)
		s, err := cc.cfg.TypeNames(nil, n)
		if err != nil {
			if strings.Contains(err.Error(), "invalid opcode") {
				return ts, nil
			} else {
				return nil, err
			}
		}

		i++
		ts = append(ts, s)
	}
	return ts, nil
}

func GetTypesName() {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		log.Println("GetTypesName", err)
		return
	}
	defer cli.Close()

	cfg, err1 := newBasConfClient(cli)
	if err1 != nil {
		log.Println("GetTypesName", err1)
		return
	}

	ts, err2 := cfg.GetTypeNames()
	if err2 != nil {
		log.Println("GetTypesName", err2)
		return
	}

	db := ldb.GetLdb()
	db.SaveTypeNames(ts)
}
