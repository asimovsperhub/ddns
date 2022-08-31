package ldb

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type RootEarningsDetails struct {
	Contract      common.Address `json:"contract"`
	RegisterOwner common.Address `json:"register_owner"`
	RegisterName  string         `json:"register_name"`
	Price         *big.Int       `json:"price"`
	RegisterTime  int64          `json:"register_time"`
}

type RootEarnings struct {
	Earnings float64                `json:"earnings"`
	Details  []*RootEarningsDetails `json:"details"`
}
type AddressEarnings struct {
	Earnings    float64                  `json:"earnings"`
	RootNameMap map[string]*RootEarnings `json:"root_name_map"`
}

type AddressCashDetails struct {
	Operator common.Address `json:"operator"`
	// RootName string         `json:"root_name"`
	Contract common.Address `json:"contract"`
	Amount   *big.Int       `json:"amount"`
	To       common.Address `json:"to"`
	CashTime int64          `json:"cash_time"`
}
type AddressCash struct {
	Amount  float64               `json:"amount"`
	Details []*AddressCashDetails `json:"details"`
}

type ContractTokenIdName struct {
	TokenName map[string]string `json:"token_name"`
}

type S3 struct {
	Upload map[string]string `json:"upload"`
}