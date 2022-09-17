package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kprc/nbsnetwork/tools"
	"log"
	"math/big"
)

type SignMsg struct {
	EntireName string         `json:"entireName"`
	Year       int            `json:"Year"`
	Erc20Addr  common.Address `json:"erc20Addr"`
	Price      *big.Int       `json:"price"`
	MsgSender  common.Address `json:"msgSender"`
	TimeStamp  *big.Int       `json:"timeStamp"`
}
type Wallet struct {
}
type WalletKey struct {
	SubPriKey []byte
	// PrivateKey表示ECDSA私钥。
	MainPriKey *ecdsa.PrivateKey
}
type PWallet struct {
	Version  string
	MainAddr common.Address
	// SubAddr   common.Address
	Crypto    keystore.CryptoJSON
	SubCipher string
	key       *WalletKey
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func encodeBytesString(v string) []byte {
	decoded, err := hex.DecodeString(v)
	if err != nil {
		panic(err)
	}
	return decoded
}

func encodeUint256(v string) []byte {
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}

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
func (acct *EthAccount) Save(auth string) error {
	if ethAcct, err := acct.Marshal(auth); err != nil {
		return err
	} else {
		return tools.Save2File(ethAcct, ethsavefile)
	}
}

// 签名
func Signer(v []byte) ([]byte, error) {
	acct, err := LoadAcct("asimov")
	if err != nil {
		acct, err = NewEthAccount()
		if err != nil {
			panic(err)
		}
		acct.Save("asimov")
	}
	p2 := acct.PrivKey
	if len(v) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(v))
	}
	if p2.Curve != crypto.S256() {
		return nil, fmt.Errorf("private key curve is not secp256k1")
	}
	log.Println(crypto.PubkeyToAddress(p2.PublicKey))
	return crypto.Sign(v, p2)
}

//func VerifySig(hash, signature []byte) bool {
//	//hash := crypto.Keccak256Hash(message)
//	// 验签的时候需要公钥
//	p2, _ := crypto.HexToECDSA("6796b36f5837d48c5f51b1cb6a46c37a7aa57bda0898974c49bdfb0b18e00278")
//	pk := crypto.FromECDSAPub(&p2.PublicKey)
//
//	signature = signature[:len(signature)-1]
//	// 公钥 数据hash 签名
//	return crypto.VerifySignature(pk, hash[:], signature)
//}
