package dnsv2

import (
	"log"
	"math/big"
	"testing"
)

func TestDnsS(t *testing.T) {
	z := new(big.Int)
	x := big.NewInt(100)
	y := big.NewInt(2)
	//number := x.String()      //转成string
	//num, _ := strconv.Atoi(number) //string转int
	cc := z.Mul(x, y)
	log.Println(cc)
}
