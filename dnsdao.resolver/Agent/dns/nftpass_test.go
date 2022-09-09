package dns

import (
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"log"
	"strings"
	"testing"
)

func TestPassCard(t *testing.T) {
	// BatchNewColdBootClient(0, 11062079440000)
	db := ldb.GetLdb()
	data, _ := db.GetNftPassTokenIdName(strings.ToLower("0x5120B54E64c1fBA0a94eF2771DaB84A97B7D7CC4"))
	log.Println(data.TokenName)
	//fmt.Println(err)
	//fmt.Println(strings.ToLower("33Gopher"))
}
