package dns

import (
	"log"
	"testing"
)

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
	isused, _ := rootC.PassCardUsed(nil, 3775)
	log.Println(isused)
	//if isused || false {
	//	log.Println("yes")
	//}
}
