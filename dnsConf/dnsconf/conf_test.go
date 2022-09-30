package dnsconf

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
	"time"
)

func TestPrintSig(t *testing.T) {
	sig := &NameConfWithSig{}

	content := &NameConf{
		Conf: make(map[string]string),
	}

	//content.Conf["eth"] = "0x9737100D2F42a196DE56ED0d1f6fF598a250E7E4"
	content.Conf["eth"] = "0x780f881b5f0EDf0CbA1F4baa4da697D0Bc3531b1"
	content.Conf["btc"] = "bc1q9x30z7rz52c97jwc2j79w76y7l3ny54nlvd4ew"

	content.Conf["ipv4"] = "127.0.0.1"
	content.Conf["cname"] = "rickey.did"

	content.Name = "xx.did"
	content.UserAddr = "0x9737100D2F42a196DE56ED0d1f6fF598a250E7E4"
	content.TimeStamp = time.Now().Unix() - 1000000
	content.Operation = 3

	if data, err := json.Marshal(*content); err != nil {
		panic(err)
	} else {
		var sigdata []byte
		if sigdata, err = sign(data); err != nil {
			panic(err)
		}
		sig.Sig = hex.EncodeToString(sigdata)
		sig.Content = string(data)
	}

	if data, err := json.Marshal(*sig); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}

}

func sign(data []byte) ([]byte, error) {
	var acct *ethacct.EthAccount
	var err error
	if acct, err = ethacct.LoadAcct("test"); err != nil {
		if acct, err = ethacct.NewEthAccount(); err != nil {
			panic(err)
		}
		acct.Save("test")
	}

	hash := crypto.Keccak256Hash(data)

	var sigdata []byte
	sigdata = append(sigdata, []byte("\x19Ethereum Signed Message:\n32")...)
	sigdata = append(sigdata, hash[:]...)

	hashsig := crypto.Keccak256Hash(sigdata)

	var sig []byte
	if sig, err = crypto.Sign(hashsig[:], acct.PrivKey); err != nil {
		return nil, err
	} else {
		return sig, nil
	}

}

func signAuth(data []byte) ([]byte, error) {
	var acct *ethacct.EthAccount
	var err error
	if acct, err = ethacct.LoadAcctFromFile("test"); err != nil {
		fmt.Println("load eth account failed")
		if acct, err = ethacct.NewEthAccount(); err != nil {
			panic(err)
		}
		acct.Save("test")
	}

	hash := crypto.Keccak256Hash(data)

	var sigdata []byte
	sigdata = append(sigdata, []byte("\x19Ethereum Signed Message:\n32")...)
	sigdata = append(sigdata, hash[:]...)

	hashsig := crypto.Keccak256Hash(sigdata)

	var sig []byte
	if sig, err = crypto.Sign(hashsig[:], acct.PrivKey); err != nil {
		return nil, err
	} else {
		return sig, nil
	}
}

//func TestHash(t *testing.T)  {
//
//}
func TestKey(t *testing.T) {
	k := reverKey("k4", "0x0011223344")
	fmt.Println(k)
	oldv := getRevertValue(k)
	defer GetLdb().CloseLdb()
	fmt.Println(oldv)
	saveRevertValue(k, "n.did", false, oldv)

	oldv1 := getRevertValue(k)
	fmt.Println(oldv1)
	saveRevertValue(k, "n1.did", false, oldv1)
	oldv1 = getRevertValue(k)
	fmt.Println(oldv1)
	var test []string

	if data, err := json.Marshal(test); err != nil {
		fmt.Println("11", err)
		return
	} else {
		GetLdb().Save([]byte("hello"), data)
	}
	if data, err := GetLdb().Get([]byte("hello")); err != nil {
		fmt.Println("222", err)
	} else {
		fmt.Println("ss", string(data))
	}

}

func TestAuthSig(t *testing.T) {
	nai := &NameAuthItem{
		Name:      "xx.did",
		TimeStamp: time.Now().Unix() - 1000000,
		Type:      "eth",
		Value:     "0x780f881b5f0EDf0CbA1F4baa4da697D0Bc3531b1",
	}

	data, err := json.Marshal(*nai)
	if err != nil {
		panic(err)
	}
	var sigauth []byte
	sigauth, err = signAuth(data)
	if err != nil {
		panic(err)
	}
	nac := &NameAuthConf{
		ItemContent: string(data),
		Content:     nai,
		ItemSig:     hex.EncodeToString(sigauth),
		TimeStamp:   time.Now().Unix() - 1000000,
		UserAddr:    "0x9737100D2F42a196DE56ED0d1f6fF598a250E7E4",
		Operation:   AuthEthConf,
	}

	data, err = json.Marshal(*nac)
	if err != nil {
		panic(err)
	}

	var sigbyte []byte
	sigbyte, err = sign(data)
	if err != nil {
		panic(err)
	}

	sig := &NameConfWithSig{
		Content: string(data),
		Sig:     hex.EncodeToString(sigbyte),
	}

	data, err = json.Marshal(*sig)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
