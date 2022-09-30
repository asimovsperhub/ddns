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
	isuseds, err := subC.PassCardUsed(nil, 10407)
	if err != nil {
		log.Println(err)
	}
	log.Println(isuseds)
	//dnsowner, errown := NewDnsOwnerClient(cli)
	//defer cli.Close()
	//if errown != nil {
	//	log.Println("NewDnsOwnerClient", errown)
	//	return
	//}
	//item, _ := subC.FilterEvMintSubName(nil)
	//data := map[string][]string{}
	//count := 0
	//for item.Next() {
	//	ev := item.Event
	//	count += 1
	//	nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
	//	//nameHashStr := hex.EncodeToString(nameHash[:])
	//	owner, err1 := dnsowner.DnsOwners(nil, *nameHash)
	//	if err1 != nil {
	//		log.Println("BatchNewRoot DnsOwners", err1)
	//	}
	//	data[owner.DnsOwner.String()] = append(data[owner.DnsOwner.String()], ev.EntireName)
	//}
	//log.Println(len(data), data, count)

	// WriteHash("sub.txt", data)
	//for k, v := range data {
	//	if len(v) > 1 {
	//		log.Println(k)
	//	}
	//}
}
