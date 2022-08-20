package api

import (
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type GetEarningsByAddress struct {
	IncomeSummary float64            `json:"income_summary"`
	Details       map[string]float64 `json:"details"`
}
type GetEarningsByAddressRes struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    *GetEarningsByAddress `json:"data"`
}

type GetEarningsDetailsByAddress struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    *ldb.AddressEarnings `json:"data"`
}

type GetCashDetailsByAddress struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *ldb.AddressCash `json:"data"`
}

type Sign struct {
	Name       string           `json:"name"`
	Work       bool             `json:"work"`
	Erc721Addr string           `json:"erc_721_addr"`
	Erc20Addr  string           `json:"erc_20_addr"`
	Income     float64          `json:"income"`
	Singners   []common.Address `json:"singners"`
	Owner      common.Address   `json:"owner"`
	TokenId    *big.Int         `json:"token_id"`
}
type QuerySignTldList struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    []*Sign `json:"data"`
}

//type AddrDomains struct {
//	Name string `json:"name"`
//}

type AddrDomainsResolveRes struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Type       string   `json:"type"`
	Addr       string   `json:"addr"`
	TotalCount int      `json:"total_count"`
	Data       []string `json:"data"`
}
