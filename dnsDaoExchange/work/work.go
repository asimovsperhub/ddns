package work

import (
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

type wGet func() (roundId *big.Int, answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int, err error)
type wSet func(acct *ethacct.EthAccount,
	roundId *big.Int,
	answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int) (tx *types.Transaction, err error)

type wInit func(acct *ethacct.EthAccount) (roundId *big.Int, err error)

type WorkItem struct {
	name   string
	lastId *big.Int
	get    wGet
	set    wSet
	init   wInit
}

type Work struct {
	Acct *ethacct.EthAccount
	Item []*WorkItem
	quit chan struct{}
}

func (w *Work) AddWorkItem(name string, init wInit, get wGet, set wSet) error {
	for i := 0; i < len(w.Item); i++ {
		if name == w.Item[i].name {
			return errors.New("duplicate name")
		}
	}

	wi := &WorkItem{
		name: name,
		get:  get,
		set:  set,
		init: init,
	}

	w.Item = append(w.Item, wi)
	return nil
}

func (w *Work) run() {
	for i := 0; i < len(w.Item); i++ {
		item := w.Item[i]
		fmt.Println(w.Item[i].name, "start get data from main net")
		roundId, answer, startedAt, updatedAt, answeredInRound, err := item.get()
		if err != nil || roundId == nil || roundId.Cmp(item.lastId) == 0 {
			fmt.Println(w.Item[i].name, "get data from main net failed")
			if err != nil {
				fmt.Println(err.Error())
			}
			continue
		}
		var tx *types.Transaction
		tx, err = item.set(w.Acct, roundId, answer, startedAt, updatedAt, answeredInRound)
		if err != nil {
			log.Println(item.name, "set data from main net failed", err.Error())
			continue
		}
		item.lastId = roundId
		fmt.Println(item.name, roundId, tx.Hash().String())
	}
}

func (w *Work) initWork() {
	for i := 0; i < len(w.Item); i++ {
		item := w.Item[i]
		if roundId, err := item.init(w.Acct); err != nil {
			panic(err)
		} else {
			item.lastId = roundId
		}
	}
}

func (w *Work) Run() {
	lastTime := int64(0)
	w.initWork()
	for {
		select {
		case <-w.quit:
			return
		default:
			if time.Now().Unix()-lastTime < 300 {
				time.Sleep(time.Second)
				continue
			}
			//log.Println("begin work")
			lastTime = time.Now().Unix()
			w.run()
		}
	}

}
