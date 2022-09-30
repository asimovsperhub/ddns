package ethacct

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kprc/nbsnetwork/tools"
)

const (
	ethsavefile     = "./ethaccount.acct"
	signfile        = "./ethaccount"
	InfuraPoint     = "https://goerli.infura.io/v3/cbc6ef92e0714bf39f4397904f3f15d2"
	InfuraMainPoint = "https://mainnet.infura.io/v3/cbc6ef92e0714bf39f4397904f3f15d2"
	//rinkeby
	//EthUsdtContract  = "0x13AFb3753858E4Ff786402db4032f474a8BF03C8"
	//SolUsdtContract  = "0x038aF4CE2a7DEB23Ca8330D4446F1ff408C0FF45"
	//UsdcUsdtContract = "0x4Acde6eD17183f06F451343D5062d502E8bAF251"
	EthUsdtContract  = "0xe0005488e0d9be7b24cce51f722d24749391327f"
	SolUsdtContract  = "0x136a734040f258054da7ebf42bd25b73e913839a"
	UsdcUsdtContract = "0x8857aaea5c663a4408963e8d964d22e438d596f2"
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

func LoadAcctFromFile(auth string) (*EthAccount, error) {
	account := &EthAccount{}
	if data, err := tools.OpenAndReadAll(signfile); err != nil {
		return nil, err
	} else {
		if err = account.Unmarshal(data, auth); err != nil {
			return nil, err
		}
		return account, nil
	}
}
