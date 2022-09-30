package dnsv2

import (
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func TestCost(t *testing.T) {
	cost, _, _ := NewCostClient()
	// rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte("did")))
	topnamehashbyte := byte32(crypto.Keccak256([]byte("did")))
	cc, _ := cost.GetAllSecondLevelNamePrice(nil, *topnamehashbyte)
	log.Println(cc)
	for _, p := range cc {
		// log.Println(p.String())
		if p.String() == "0" {
			log.Println("t")
		}
	}

}
