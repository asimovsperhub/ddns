package main

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kprc/nbsnetwork/tools"
	"math/big"
)

type EthAccount struct {
	PrivKey *ecdsa.PrivateKey `json:"-"`
	PubKey  *ecdsa.PublicKey  `json:"-"`
	EAddr   common.Address    `json:"-"`
	SAddr   string            `json:"s_addr"`
}

type AccountJson struct {
	Acct *EthAccount          `json:"acct"`
	Cj   *keystore.CryptoJSON `json:"cj"`
}

func (acct *EthAccount) Unmarshal(data []byte, auth string) error {

	aj := &AccountJson{}
	if err := json.Unmarshal(data, aj); err != nil {
		return err
	}

	acct.SAddr = aj.Acct.SAddr

	if keyBytes, err := keystore.DecryptDataV3(*aj.Cj, auth); err != nil {
		return err
	} else {
		acct.PrivKey = crypto.ToECDSAUnsafe(keyBytes)
		acct.PubKey = (acct.PrivKey.Public()).(*ecdsa.PublicKey)
		acct.EAddr = crypto.PubkeyToAddress(*acct.PubKey)
	}

	return nil
}

func EthUnmarshal(data []byte) (*AccountJson, error) {
	aj := &AccountJson{}
	if err := json.Unmarshal(data, aj); err != nil {
		return nil, err
	}

	return aj, nil
}

func (acct *EthAccount) ImportFromMetaMask(hexString string) (err error) {
	var priKey *ecdsa.PrivateKey
	if priKey, err = crypto.HexToECDSA(hexString); err != nil {
		return err
	}

	acct.PrivKey = priKey
	acct.PubKey = (acct.PrivKey.Public()).(*ecdsa.PublicKey)
	acct.EAddr = crypto.PubkeyToAddress(*acct.PubKey)
	acct.SAddr = acct.EAddr.String()

	return nil
}

func (acct *EthAccount) Marshal(auth string) ([]byte, error) {
	keyBytes := math.PaddedBigBytes(acct.PrivKey.D, 32)
	aj := &AccountJson{}
	if cs, err := keystore.EncryptDataV3(keyBytes, []byte(auth), keystore.StandardScryptN, keystore.StandardScryptP); err != nil {
		return nil, err
	} else {
		aj.Acct = acct
		aj.Cj = &cs
	}

	return json.Marshal(*aj)
}

func NewEthAccount() (acct *EthAccount, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	acct = &EthAccount{}
	acct.PrivKey = key
	acct.PubKey = (key.Public()).(*ecdsa.PublicKey)
	acct.EAddr = crypto.PubkeyToAddress(*acct.PubKey)
	acct.SAddr = acct.EAddr.String()

	return acct, nil
}

const (
	ethsavefile = "./ethaccount"
)

func (acct *EthAccount) Save(auth string) error {
	if ethAcct, err := acct.Marshal(auth); err != nil {
		return err
	} else {
		return tools.Save2File(ethAcct, ethsavefile)
	}
}

func LoadAcct(auth string) (*EthAccount, error) {
	account := &EthAccount{}
	if data, err := tools.OpenAndReadAll(ethsavefile); err != nil {
		return nil, err
	} else {
		if err = account.Unmarshal(data, auth); err != nil {
			return nil, err
		}
		return account, nil
	}
}

//type T struct {
//	Code    int    `json:"code"`
//	Message string `json:"message"`
//	Data    struct {
//		Signature string `json:"signature"`
//		PAddr     string `json:"p_addr"`
//		Params    struct {
//			DomainName string `json:"domain_name"`
//			Year       int    `json:"year"`
//			Erc20Addr  string `json:"erc_20_addr"`
//			Price      string `json:"price"`
//			MsgSender  string `json:"msg_sender"`
//			TokenId    int    `json:"token_id"`
//		} `json:"params"`
//	} `json:"data"`
//}

//{
//"code": 1,
//"message": "ok",
//"data": {
//"signature": "32b9b43e20c60aa71fc339be674372cb08875ff27f66f01c5cc9354ee288c74436ea7311c44b55602b9d0736329040d1728fb373b014ee2da0d9d9eeb6c765f001",
//"p_addr": "0x0000000000000000000000000000000000000000",
//"params": {
//"domain_name": "rickey.did",
//"year": 2,
//"erc_20_addr": "0x6043bfe64a866c7fb17d1855fe3ebc4342ca9c15",
//"price": "3000000000000000000000",
//"msg_sender": "0x5da2a8ec74a62089c8678e9db3ea6d3e8d265ede",
//"token_id": 1
//}
//}
//}
//abi.encodePacked(entireName_,year_,erc20Addr_,price_,msg.sender,passId_
func pack(domainName string, year uint8, erc20Addr common.Address, price *big.Int, address common.Address, tokenId uint32) []byte {
	var concatBytes []byte

	concatBytes = append(concatBytes, []byte(domainName)...)
	concatBytes = append(concatBytes, year)
	concatBytes = append(concatBytes, erc20Addr[:]...)
	var pb [32]byte
	price.FillBytes(pb[:])

	concatBytes = append(concatBytes, pb[:]...)

	concatBytes = append(concatBytes, address[:]...)
	var t [4]byte
	binary.BigEndian.PutUint32(t[:], tokenId)
	concatBytes = append(concatBytes, t[:]...)

	return concatBytes

}

//udid did plur debox id dns dnsdao org edu gov int mil eth sol bit bnb lens
func main() {
	acct, err := LoadAcct("asimov")
	if err != nil {
		fmt.Println(err.Error())
		acct, err = NewEthAccount()
		if err != nil {
			panic(err)
		}
		acct.Save("asimov")
	}

	fmt.Println(acct.SAddr)

	price := &big.Int{}
	price.SetInt64(100)
	dec := &big.Int{}
	dec.SetInt64(1000000000000000000)

	price = price.Mul(price, dec)

	cb := pack("rickey.did", 1, common.HexToAddress("0x6043bfe64a866c7fb17d1855fe3ebc4342ca9c15"),
		price, common.HexToAddress("0x8D95f294D667884FCe823ACa4E7f2A65E4664916"), 4734)

	fmt.Println(hex.EncodeToString(cb))

	h1 := crypto.Keccak256Hash(cb)

	fmt.Println("h1", hex.EncodeToString(h1[:]))

	//var h2b []byte
	//
	//h2b = append(h2b, []byte("\x19Ethereum Signed Message:\n32")...)
	//h2b = append(h2b, h1[:]...)

	//h2 := crypto.Keccak256Hash(h2b)

	//fmt.Println("h2", hex.EncodeToString(h2[:]))

	var sig []byte
	sig, err = crypto.Sign(h1[:], acct.PrivKey)
	if err != nil {
		panic(err)
	}

	if sig[len(sig)-1] <= 1 {
		sig[len(sig)-1] = sig[len(sig)-1] + 27
	}

	fmt.Println(hex.EncodeToString(sig))
	//
	h2 := crypto.Keccak256Hash([]byte("007.did"))
	fmt.Println("h2", hex.EncodeToString(h2[:]))

}
