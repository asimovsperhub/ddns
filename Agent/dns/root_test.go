package dns

import (
	"encoding/json"
	udidc "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv2"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
	"testing"
)

func WriteHash(writeJson string, cont interface{}) {
	// os.O_TRUNC 覆盖  os.O_APPEND 追加
	if distFile, err := os.OpenFile(writeJson, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
		log.Println("create file failed", err)
	} else {
		enc := json.NewEncoder(distFile)
		if err1 := enc.Encode(cont); err1 != nil {
			log.Println("write hash.json failed", err1)
		} else {
			log.Println("write hash.json successful")
		}
	}
}
func TestRoot(t *testing.T) {
	BatchNewRoot(0, 11062079440000)
	//db := ldb.GetLdb()
	//_, err := db.GetContractList()
	//fmt.Println(err)
	//fmt.Println(strings.ToLower("33Gopher"))
}

func TestRootPassCard(t *testing.T) {
	rootC, cli, err := NewRootClient()
	defer cli.Close()
	if err != nil {
		log.Println("BatchNewRoot", err)
		return
	}
	// mint root域名
	var mintiter *udidc.DnsNameEvMintDnsNameIterator
	mintiter, err = rootC.FilterEvMintDnsName(nil)
	if err != nil {
		log.Println("FilterEvMintDnsName", err)
		return
	}
	dnsowner, err := NewDnsOwnerClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	name := map[string][]string{}
	if mintiter != nil {
		defer mintiter.Close()
		for mintiter.Next() {
			ev := mintiter.Event
			nameHash := byte32(crypto.Keccak256([]byte(ev.EntireName)))
			// nameHashStr := hex.EncodeToString(nameHash[:])
			owner, err1 := dnsowner.DnsOwners(nil, *nameHash)
			if err1 != nil {
				log.Println("BatchNewRoot DnsOwners", err1)
			}
			name[owner.DnsOwner.String()] = append(name[owner.DnsOwner.String()], ev.EntireName)
			log.Println(len(name), name)
			//isused, _ := rootC.PassCardUsed(nil, 3775)
			//log.Println(isused)
			//if isused || false {
			//	log.Println("yes")
			//}
		}
	}
	log.Println(name)
	WriteHash("root.txt", name)
}
