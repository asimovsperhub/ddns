package usdcusdt

import (
	"context"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type UsdcUSDTClient struct {
	client  *ethclient.Client
	receipt *types.Receipt
	tx      *types.Transaction
	acct    *ethacct.EthAccount
}

func NewClient(acct *ethacct.EthAccount) (*UsdcUSDTClient, error) {
	cli, err := ethclient.Dial(ethacct.InfuraPoint)
	if err != nil {
		return nil, err
	}
	return &UsdcUSDTClient{
		client: cli,
		acct:   acct,
	}, nil
}

func (euc *UsdcUSDTClient) Close() {
	euc.client.Close()
}

func SetSolUsdt(acct *ethacct.EthAccount, roundId *big.Int,
	answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int) (tx *types.Transaction, err error) {
	var cli *UsdcUSDTClient
	cli, err = NewClient(acct)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	return cli.InsertExchange(roundId, answer, startedAt, updatedAt, answeredInRound)
}

func GetLastestRound(acct *ethacct.EthAccount) (roundId *big.Int, err error) {
	var cli *UsdcUSDTClient
	cli, err = NewClient(acct)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	return cli.getLatest()
}

func (euc *UsdcUSDTClient) getLatest() (roundId *big.Int, err error) {
	var eu *udidc22.EthUsdt
	eu, err = udidc22.NewEthUsdt(common.HexToAddress(ethacct.UsdcUsdtContract), euc.client)
	if err != nil {
		//log.Println(err)
		return
	}
	var n struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}

	n, err = eu.LatestRoundData(nil)
	if err != nil {
		return nil, err
	}
	return n.RoundId, nil
}

func (euc *UsdcUSDTClient) InsertExchange(
	roundId *big.Int,
	answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int) (tx *types.Transaction, err error) {

	var eu *udidc22.SolUSDT
	eu, err = udidc22.NewSolUSDT(common.HexToAddress(ethacct.UsdcUsdtContract), euc.client)
	if err != nil {
		//log.Println(err)
		return
	}

	var cid *big.Int
	cid, err = euc.client.ChainID(context.TODO())
	if err != nil {
		//log.Println(err)
		return
	}
	var opt *bind.TransactOpts
	opt, err = bind.NewKeyedTransactorWithChainID(euc.acct.PrivKey, cid)
	if err != nil {
		//log.Println(err)
		return
	}

	//sign and insert
	//var ;
	tx, err = eu.InsertIntoRoundInfo(opt, roundId, answer, startedAt, updatedAt, answeredInRound)
	if err != nil {
		return
	}

	euc.tx = tx

	return
}

func (euc *UsdcUSDTClient) WaitTxReply() (err error) {
	euc.receipt, err = bind.WaitMined(context.TODO(), euc.client, euc.tx)
	return
}