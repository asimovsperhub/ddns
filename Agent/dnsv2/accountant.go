package dnsv2

import (
	"context"
	"github.com/dnsdao/dnsdao.resolver/config"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strconv"
	"strings"
	"time"
)

type Deposit struct {
	FromContract common.Address `json:"from_contract"`
	ToContract   common.Address `json:"to_contract"`
	Erc20        common.Address `json:"erc_20"`
	Amount       common.Address `json:"amount"`
}

func NewDnsAccountantClient(client *ethclient.Client) (*udidc22.DnsAccountant, error) {
	var (
		dnsaccountant *udidc22.DnsAccountant
		err           error
	)
	dnsaccountant, err = udidc22.NewDnsAccountant(common.HexToAddress(config.GetRConf().Cconf.Accountant), client)
	if err != nil {
		client.Close()
		log.Println("NewRootClient NewDnsName  err", err)
		return nil, err
	}
	return dnsaccountant, nil
}

func BatchNewAccountant(start, end uint64) {
	var (
		cli           *ethclient.Client
		dnsaccountant *udidc22.DnsAccountant
		err           error
	)
	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsaccountant, err = NewDnsAccountantClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	db := ldb.GetLdb()
	b, e := start, end
	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	//var deposit *udidc22.DnsAccountantEvDepositIterator
	//deposit, err = dnsaccountant.FilterEvDeposit(op)
	//for deposit.Next() {
	//	ev := deposit.Event
	//	fmt.Println(ev.Operator, ev.ContractAddr, ev.Erc20Addr, ev.Amount)
	//}
	var drawcash *udidc22.DnsAccountantEvWithdrawCashIterator
	drawcash, err = dnsaccountant.FilterEvWithdrawCash(op)
	if drawcash != nil {
		defer drawcash.Close()
		for drawcash.Next() {
			ev := drawcash.Event
			log.Println("Cash", ev.Operator, ev.ContractAddr, ev.Amount, ev.Out, ev.Erc20Addr)
			addrkey := strings.ToLower(ev.Operator.String())
			cash, _ := db.GetAddressCash(addrkey)
			amount, _ := strconv.Atoi(ev.Amount.String())
			if cash != nil {
				cash.Amount += float64(amount) / 1000000000000000000
				cash.Details = append(cash.Details, &ldb.AddressCashDetails{Amount: ev.Amount, Operator: ev.Operator, Contract: ev.ContractAddr, To: ev.Out, CashTime: time.Now().Unix()})
				errsave := db.SaveAddressCash(addrkey, cash)
				if errsave != nil {
					log.Println("SaveAddressCash err", errsave)
				}
			} else {
				cashN := new(ldb.AddressCash)
				cashN.Amount = float64(amount) / 1000000000000000000
				cashN.Details = []*ldb.AddressCashDetails{&ldb.AddressCashDetails{
					Amount: ev.Amount, Operator: ev.Operator, Contract: ev.ContractAddr, To: ev.Out, CashTime: time.Now().Unix()},
				}
				errsave := db.SaveAddressCash(addrkey, cashN)
				if errsave != nil {
					log.Println("SaveAddressCash err", errsave)
				}
			}
		}
	}
}
