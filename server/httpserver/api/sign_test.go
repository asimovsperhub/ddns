package api

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func Test(t *testing.T) {
	//cc := big.NewInt(1)
	//fmt.Println(cc.String() == "1")
	result := encodePacked(
		encodeBytesString(hex.EncodeToString([]byte("rickey.did"))),
		// encodeUint256Array(signatures),
		encodeUint256("2")[32-1:],
		encodeBytesString("6043bfe64a866c7fb17D1855fe3eBC4342Ca9c15"),
		encodeUint256("3000000000000000000000"),
		encodeBytesString("5da2a8ec74a62089c8678e9db3ea6d3e8d265ede"),
		encodeUint256("1")[32-4:],
	)
	// cc, _ := NewWallet("a")
	hash := crypto.Keccak256Hash(result)
	fmt.Println("result", hex.EncodeToString(result[:]))
	fmt.Println("hash", hex.EncodeToString(hash[:]))
	hash2 := encodePacked(
		hash[:],
	)
	hash3 := crypto.Keccak256Hash(hash2)
	fmt.Println("hash3", hex.EncodeToString(hash3[:]))
	// 签名
	signer, err := Signer(hash[:])
	if err != nil {
		log.Println(err)
	}
	// log.Println(signer)
	// 验签
	log.Println("signer", hex.EncodeToString(signer), len(signer))
	f := VerifySig(hash[:], signer)
	fmt.Println("VerifySig", f)
}
