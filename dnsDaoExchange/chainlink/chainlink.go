package chainlink

import (
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// sol-usd 24h 8 0x4ffC43a60e009B551865A93d232E33Fce9f01507
// eth-usd 1h  8 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419
// usdc-usd 24h 8 0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6

const (
	SolUSD  = "0x4ffC43a60e009B551865A93d232E33Fce9f01507"
	EthUSD  = "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"
	USDCUSD = "0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6"
)

func GetEthLatestRoundInfo() (roundId *big.Int, answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int, err error) {
	var client *ethclient.Client
	client, err = ethclient.Dial(ethacct.InfuraMainPoint)
	if err != nil {
		return
	}
	defer client.Close()
	var agv *udidc22.AggregatorV3Interface
	agv, err = udidc22.NewAggregatorV3Interface(common.HexToAddress(EthUSD), client)
	if err != nil {
		return
	}
	var r struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}
	r, err = agv.LatestRoundData(nil)

	return r.RoundId, r.Answer, r.StartedAt, r.UpdatedAt, r.AnsweredInRound, err
}

func GetSolLatestRoundInfo() (roundId *big.Int, answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int, err error) {
	var client *ethclient.Client
	client, err = ethclient.Dial(ethacct.InfuraMainPoint)
	if err != nil {
		return
	}
	defer client.Close()
	var agv *udidc22.AggregatorV3Interface
	agv, err = udidc22.NewAggregatorV3Interface(common.HexToAddress(SolUSD), client)
	if err != nil {
		return
	}
	var r struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}
	r, err = agv.LatestRoundData(nil)

	return r.RoundId, r.Answer, r.StartedAt, r.UpdatedAt, r.AnsweredInRound, err
}

func GetUsdcLatestRoundInfo() (roundId *big.Int, answer *big.Int,
	startedAt *big.Int,
	updatedAt *big.Int,
	answeredInRound *big.Int, err error) {
	var client *ethclient.Client
	client, err = ethclient.Dial(ethacct.InfuraMainPoint)
	if err != nil {
		return
	}
	defer client.Close()
	var agv *udidc22.AggregatorV3Interface
	agv, err = udidc22.NewAggregatorV3Interface(common.HexToAddress(USDCUSD), client)
	if err != nil {
		return
	}
	var r struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}
	r, err = agv.LatestRoundData(nil)

	return r.RoundId, r.Answer, r.StartedAt, r.UpdatedAt, r.AnsweredInRound, err
}
