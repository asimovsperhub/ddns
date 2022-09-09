package dns

import (
	"testing"
)

//func TestCurrentBlk(t *testing.T) {
//	if blk, err := currentBlk(); err != nil {
//		panic(err)
//	} else {
//		fmt.Println(blk)
//	}
//
//}

func TestBatchNewRootName(t *testing.T) {
	BatchNewRootName(0, 12324651)
}
