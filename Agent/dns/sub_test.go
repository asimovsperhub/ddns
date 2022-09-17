package dns

import (
	"log"
	"testing"
)

func TestSub(t *testing.T) {
	BatchNewSub(0, 1106207944000)
	//type CC struct {
	//	Price map[string]*big.Int `json:"price"`
	//}
	//cc := new(CC)
	//if cc.Price["c"] == nil {
	//	fmt.Println("1", cc.Price["c"])
	//	fmt.Println("2", cc.Price["c"].String())
	//	price, _ := strconv.Atoi(cc.Price["c"].String())
	//	fmt.Println("CCC", price)
	//}
}
func TestSubPassCard(t *testing.T) {
	subC, _, err := NewSubClient("0x111cfbc15fc56ad0ac8c7536bca3347a558a0012")
	if err != nil {
		log.Println(err)
	}
	isuseds, err := subC.PassCardUsed(nil, 1662)
	if err != nil {
		log.Println(err)
	}
	log.Println(isuseds)
	// dnsowner, errown := NewDnsOwnerClient(cli)
	//defer cli.Close()
	//if errown != nil {
	//	log.Println("NewDnsOwnerClient", errown)
	//	return
	//}
	//item, _ := subC.FilterEvMintSubName(nil)
	//data := map[string][]string{}
	//for item.Next() {
	//	ev := item.Event
	//	// nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
	//	//nameHashStr := hex.EncodeToString(nameHash[:])
	//	if data[ev.EntireName] == nil {
	//		data[ev.EntireName] = []string{}
	//	} else {
	//		data[ev.EntireName] = append(data[ev.EntireName], ev.EntireName)
	//	}
	//}
	//for k, v := range data {
	//	if len(v) > 1 {
	//		log.Println(k)
	//	}
	//}
}
