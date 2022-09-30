package dnsconf

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"testing"
)

func TestNameIsNormal(t *testing.T) {
	name := "aac"

	if IsRootName(name) {
		fmt.Println("root")
	}
	if IsNormalName(name) {
		fmt.Println("normal")
	}

	arrDot := strings.Split(name, ".")
	if len(arrDot) >= 2 {
		fmt.Println("root", arrDot[len(arrDot)-1])
		fmt.Println("second", arrDot[len(arrDot)-2])
	}

}

func TestTopOwner(t *testing.T) {
	name := "did"

	if b, owner, err := findTop(crypto.Keccak256Hash([]byte(name))); err != nil {
		panic(err)
	} else {
		fmt.Println(b, owner.String())
	}

}
