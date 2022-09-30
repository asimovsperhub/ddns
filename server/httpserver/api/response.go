package api

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
	{
    code:1,  // 1 服务正常返回，0 服务端异常
    message:'ok',// 异常时 描述
    data:{  // 异常时 data null 或 无此字段
        total:32,
        "page_size": 100,
        "page_number": 0,
        items:[
            {
                name:"did",
                owner:'0x...865',
                erc_721_addr:'0x67...98Fd', //
                token_id:1,
                "work": false,
                "erc_20_addr": "0x0000000000000000000000000000000000000000",
                income:[
                  {
                    erc20_addr:'0x000...00', // eth address, required
                    withdraw_wei:'12000500000000000000000000',  // eth withdraw balance, wei  required
                    decimals:18, // optional
                    symbol:'ETH',// optional
                    withdraw_value: 12.0005,// optional show value
                  },

*/

type GetEarningsByAddressIncome struct {
	Erc20Addr     common.Address `json:"erc_20_addr"`
	WithDrawWei   *big.Int       `json:"with_draw_wei"`
	Decimals      int            `json:"decimals"`
	Symbol        string         `json:"symbol"`
	WithdrawValue *big.Int       `json:"withdraw_value"`
}

type GetEarningsByAddressItems struct {
	Name       string                             `json:"name"`
	Owner      common.Address                     `json:"owner"`
	Erc721Addr common.Address                     `json:"erc_721_addr"`
	TokenId    *big.Int                           `json:"token_id"`
	Work       bool                               `json:"work"`
	Erc20Addr  common.Address                     `json:"erc_20_addr"`
	Income     []*GetEarningsByAddressIncome      `json:"income"`
	Signers    []*GetSignTldListByDidNameSignList `json:"signers"`
}
type GetEarningsByAddressRes struct {
	Total      int                          `json:"total"`
	PageSize   int                          `json:"page_size"`
	PageNumber int                          `json:"page_number"`
	Items      []*GetEarningsByAddressItems `json:"items"`
}

/*
	data:{  // 异常时 data null 或 无此字段
      total:4,
      "page_size": 100,
      "page_number": 0,
      items:[
        {
          namehash: '0x....abs',        //多签任务 对应域名 的 namehash,如 `did` hash
          task_hash:'0x....abc', //
          max_sig:3,                    // 最少签名数量
          work:true,                    // 对应合约 contracts\dnsProject\multisig\LibMultiSig.sol work
          lock:true,                    //
          signer_list:[                 // 已参与签名的 list
            {
              signer: '0x...ac',        // 操作签名的地址
              sign_type:1,              // 操作类型 ，Agree 1 ，Disagree： 0
            },
            {

*/
type GetSignTldListByDidNameSignList struct {
	Signer   common.Address `json:"signer"`
	SignType int            `json:"sign_type"`
}
type GetSignTldListByDidNameItems struct {
	NameHash   string                             `json:"name_hash"`
	TaskHash   []byte                             `json:"task_hash"`
	MaxSig     *big.Int                           `json:"max_sig"`
	Work       bool                               `json:"work"`
	Lock       bool                               `json:"lock"`
	SignerList []*GetSignTldListByDidNameSignList `json:"signer_list"`
}
type GetSignTldListByDidNameRes struct {
	Total      int                             `json:"total"`
	PageSize   int                             `json:"page_size"`
	PageNumber int                             `json:"page_number"`
	Items      []*GetSignTldListByDidNameItems `json:"items"`
}

type Sign struct {
	Name       string           `json:"name"`
	Work       bool             `json:"work"`
	Erc721Addr string           `json:"erc_721_addr"`
	Erc20Addr  string           `json:"erc_20_addr"`
	Income     *big.Int         `json:"income"`
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
	Name       string         `json:"name"`
	Erc721Addr common.Address `json:"erc_721_addr"`
	TokenId    *big.Int       `json:"token_id"`
	//OpenToReg   bool           `json:"open_to_reg"`
	//HasPrice    bool           `json:"has_price"`
	ExpireTime *big.Int       `json:"expire_time"`
	Owner      common.Address `json:"owner"`
	UsdtPrices [8]*big.Int    `json:"usdt_prices"`
}

type AddrTopListData struct {
	Total      int         `json:"total"`
	PageSize   int         `json:"page_size"`
	PageNumber int         `json:"page_number"`
	Items      interface{} `json:"items"`
}
type ResData struct {
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
	DomainName  string         `json:"domain_name"`
	Year        int32          `json:"year"`
	Erc20Addr   common.Address `json:"erc_20_addr"`
	Price       string         `json:"price"`
	MsgSender   common.Address `json:"msg_sender"`
	TokenId     int32          `json:"token_id"`
	Nonce       string         `json:"nonce"`
	LastPriceId string         `json:"lastPriceId"`
}
type PostSignMintData struct {
	Signature string `json:"signature"`
	// PAddr     common.Address      `json:"p_addr"`
	Params *PostSignMintParams `json:"params"`
}

func NotDataRes(msg string) string {
	res := &Res{
		Code:    0,
		Message: msg,
		Data:    &ResData{Total: 0, PageNumber: 0, PageSize: 0},
	}
	resbyte, err := json.Marshal(res)
	if err == nil {
		return string(resbyte)
	}
	return ""
}

type TotalBySubLenResData struct {
	All   int64              `json:"all"`
	Rules []*TotalBySubRules `json:"rules"`
}
type TotalBySubRules struct {
	Rule  string `json:"rule"`
	Min   int64  `json:"min"`
	Max   int64  `json:"max"`
	Total int64  `json:"total"`
}
