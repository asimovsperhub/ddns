package main

import (
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/chainlink"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethusdt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/solusdt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/usdcusdt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/work"
	"log"
)

func main() {
	acct, err := ethacct.LoadAcct("test")
	if err != nil {
		fmt.Println(err.Error())
		acct, err = ethacct.NewEthAccount()
		if err != nil {
			panic(err)
		}
		acct.Save("test")
	}

	log.Println(acct.SAddr)

	w := &work.Work{
		Acct: acct,
	}
	w.AddWorkItem("eth-usd", ethusdt.GetLastestRound, chainlink.GetEthLatestRoundInfo, ethusdt.SetEthUsdt)
	w.AddWorkItem("sol-usdt", solusdt.GetLastestRound, chainlink.GetSolLatestRoundInfo, solusdt.SetSolUsdt)
	w.AddWorkItem("usdc-usdt", usdcusdt.GetLastestRound, chainlink.GetUsdcLatestRoundInfo, usdcusdt.SetSolUsdt)

	w.Run()

}
