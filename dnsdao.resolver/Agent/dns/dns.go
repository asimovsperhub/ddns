package dns

import (
	"context"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
)

type DNSAgent struct {
	quit chan struct{}
}

var (
	currentBlkNumber uint64
	lastDnsLoopTime  int64
	dnsAgent         *DNSAgent
	dnsAgentOnce     sync.Once
)

func GetDNSAgent() *DNSAgent {
	if dnsAgent == nil {
		dnsAgentOnce.Do(func() {
			dnsAgent = &DNSAgent{
				quit: make(chan struct{}, 8),
			}
		})
	}
	return dnsAgent
}

func currentBlk() (uint64, error) {
	var (
		cli *ethclient.Client
		err error
	)

	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		return 0, err
	}

	defer cli.Close()

	return cli.BlockNumber(context.TODO())
}

func loopForNewName() {
	var err error
	db := ldb.GetLdb()
	dbblk := db.GetLatestBlkNum()

	currentBlkNumber, err = currentBlk()
	if err != nil {
		return
	}

	// GetTypesName()
	//BatchNewRootName(dbblk, currentBlkNumber)
	//BatchNewSubName(dbblk, currentBlkNumber)
	//BatchMarketBuyFromSeller(dbblk, currentBlkNumber)
	//BatchDomainConfChanged(dbblk, currentBlkNumber)
	BatchNewRoot(dbblk, currentBlkNumber)
	BatchNewPrice(dbblk, currentBlkNumber)
	BatchNewSub(dbblk, currentBlkNumber)
	BatchNewOwner(dbblk, currentBlkNumber)
	BatchNewConf(dbblk, currentBlkNumber)
	BatchNewAccountant(dbblk, currentBlkNumber)
	//defer db.CloseLdb()
	db.SaveLatestBlkNum(currentBlkNumber)

}

func (ag *DNSAgent) DNSLoopEvent() {

	for {
		now := time.Now().Unix()

		if now-lastDnsLoopTime > 300 {
			lastDnsLoopTime = now
			loopForNewName()
		} else {
			time.Sleep(time.Second)
		}

		select {
		case <-ag.quit:
			return
		default:
			//continue
		}
	}
}

func (ag *DNSAgent) Quit() {
	ag.quit <- struct{}{}
}
