package api

import (
	"encoding/json"
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

type AddrDomains struct {
	Name      string `json:"name"`
	ConfType  string `json:"conf_type"`
	ConfValue string `json:"conf_value"`
}

type AddrDomainsResolveRes struct {
	Code       int            `json:"code"`
	Message    string         `json:"message"`
	TotalCount int            `json:"total_count"`
	Data       []*AddrDomains `json:"data"`
}

type AddrTopListDataItems struct {
	Name        string         `json:"name"`
	Erc721_Addr common.Address `json:"erc_721_addr"`
	TokenId     *big.Int       `json:"token_id"`
	OpenToReg   bool           `json:"open_to_reg"`
	ExpireTime  *big.Int       `json:"expire_time"`
	Owner       common.Address `json:"owner"`
	PayTokens   []string       `json:"pay_tokens"`
}

type AddrTopListData struct {
	Total      int         `json:"total"`
	PageSize   int         `json:"page_size"`
	PageNumber int         `json:"page_number"`
	Items      interface{} `json:"items"`
}

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PostSignMintParams struct {
	DomainName string         `json:"domain_name"`
	Year       int32          `json:"year"`
	Erc20Addr  common.Address `json:"erc_20_addr"`
	Price      string         `json:"price"`
	MsgSender  common.Address `json:"msg_sender"`
	TokenId    int32          `json:"token_id"`
}
type PostSignMintData struct {
	Signature string              `json:"signature"`
	PAddr     common.Address      `json:"p_addr"`
	Params    *PostSignMintParams `json:"params"`
}

func NotDataRes(msg string) string {
	res := &Res{
		Code:    0,
		Message: msg,
	}
	resbyte, err := json.Marshal(res)
	if err == nil {
		return string(resbyte)
	}
	return ""
}
