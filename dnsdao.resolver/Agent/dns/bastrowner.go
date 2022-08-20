package dns

import (
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/contract/dnscontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
)

func newBasTRClient(cli *ethclient.Client) (*BASTRC, error) {

	tr, err := dnscontract.NewBasTradableOwnership(common.HexToAddress(config.GetRConf().Cconf.BasTradableOwnership), cli)
	if err != nil {
		cli.Close()
		return nil, err
	}
	return &BASTRC{tr: tr}, nil
}

type BASTRC struct {
	tr *dnscontract.BasTradableOwnership
}

type OwnerInfo struct {
	Owner      common.Address
	ExpireTime int64
}

func (tr *BASTRC) GetOwnerInfo(nameHash [32]byte) (*OwnerInfo, error) {

	var (
		addr common.Address
		t    *big.Int
		err  error
	)

	addr, t, err = tr.tr.OwnerOfWithExpire(nil, nameHash)
	if err != nil {
		return nil, err
	}

	//fmt.Println(hex.EncodeToString(addr[:]))
	//fmt.Println(t.String())

	return &OwnerInfo{
		Owner:      addr,
		ExpireTime: t.Int64(),
	}, nil
}
