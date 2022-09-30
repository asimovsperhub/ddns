package dnsconf

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
	"time"
)

type MapValue struct {
	Value string `json:"value"`
	Auth  bool   `json:"auth"`
}

type NameMap map[string]string

type NameMapDb map[string]*MapValue

const (
	AddConf = iota + 1
	DelConf
	ReplaceConf
	AuthEthConf
)

var confTypes []string = []string{
	"resolve", "eth", "btc", "sol", "etc", "trx", "eos", "bnb", "fil", "ipv4", "ipv6", "cname"}

var confTypesSet map[string]struct{}

func init() {
	confTypesSet = make(map[string]struct{})
	for i := 0; i < len(confTypes); i++ {
		confTypesSet[confTypes[i]] = struct{}{}
	}
}

func isConfTypes(conf map[string]string) bool {
	for k, _ := range conf {
		if _, ok := confTypesSet[k]; !ok {
			return false
		}
	}
	return true
}

func NameConfKey(name string) string {
	h := crypto.Keccak256Hash([]byte(name))
	hash := strings.ToLower(h.String())

	return fmt.Sprintf("nameconfkey_1_%s", hash)
}

type NameConf struct {
	Name      string  `json:"name"`
	Operation int     `json:"operation"` //AddConf DelConf ReplaceConf
	Conf      NameMap `json:"conf"`
	TimeStamp int64   `json:"timeStamp"`
	UserAddr  string  `json:"userAddr"`
}

func (nc *NameConf) String() string {
	j, _ := json.MarshalIndent(*nc, "\t", " ")
	return string(j)
}

type NameConfDb struct {
	Name      string    `json:"name"`
	Conf      NameMapDb `json:"conf"`
	TimeStamp int64     `json:"timeStamp"`
	UserAddr  string    `json:"userAddr"`
}

func (nc *NameConfDb) String() string {
	j, _ := json.MarshalIndent(*nc, "\t", " ")
	return string(j)
}

type NameAuthItem struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	TimeStamp int64  `json:"timeStamp"`
}

func (nc *NameAuthItem) String() string {
	j, _ := json.MarshalIndent(*nc, "\t", " ")
	return string(j)
}

type NameAuthConf struct {
	ItemContent string        `json:"itemContent"` //json unmarshal from NameAuthItem
	Content     *NameAuthItem `json:"-"`
	ItemSig     string        `json:"itemSig"`
	TimeStamp   int64         `json:"timeStamp"`
	UserAddr    string        `json:"userAddr"`
	Operation   int           `json:"operation"`
}

func (nc *NameAuthConf) String() string {
	j, _ := json.MarshalIndent(*nc, "\t", " ")
	return string(j)
}

type NameConfWithSig struct {
	Content string `json:"content"`
	Sig     string `json:"sig"`
}

func (nc *NameConfWithSig) String() string {
	j, _ := json.MarshalIndent(*nc, "\t", " ")
	return string(j)
}

func SupportConfItem() []string {
	return confTypes
}

func reverKey(k string, v string) string {
	return fmt.Sprintf("reverKey_1_%s_%s", k, v)
}

func getRevertValue(key string) []*MapValue {
	if v, err := GetLdb().Get([]byte(key)); err != nil {
		return nil
	} else {
		var r []*MapValue
		if err = json.Unmarshal(v, &r); err != nil {
			return nil
		}
		return r
	}
}

func GetRevertValue(typ, key string) []*MapValue {
	k := reverKey(typ, key)

	return getRevertValue(k)
}

func deleteRevertValue(key string, name string, oldvalue []*MapValue) {
	idx := -1
	for i := 0; i < len(oldvalue); i++ {
		if name == oldvalue[i].Value {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}

	if len(oldvalue) == 1 {
		var ss []*MapValue
		data, _ := json.Marshal(ss)
		GetLdb().Save([]byte(key), data)
	} else {
		oldvalue[idx] = oldvalue[len(oldvalue)-1]
		oldvalue = oldvalue[:len(oldvalue)-1]
		data, _ := json.Marshal(oldvalue)
		GetLdb().Save([]byte(key), data)
	}
}

func saveRevertValue(key string, name string, auth bool, oldvalue []*MapValue) error {
	for i := 0; i < len(oldvalue); i++ {
		if name == oldvalue[i].Value {
			return nil
		}
	}

	oldvalue = append(oldvalue, &MapValue{Value: name, Auth: auth})

	if data, err := json.Marshal(oldvalue); err != nil {
		return err
	} else {
		return GetLdb().Save([]byte(key), data)
	}

}

func saveRevertMapping(newconf map[string]*MapValue, oldconf map[string]*MapValue, name string) {
	saveNewRevertMapping(newconf, name)
	for k, v := range oldconf {
		if vv, ok := newconf[k]; !ok || vv.Value == "" || vv.Value != v.Value {
			key := reverKey(k, v.Value)
			oldv := getRevertValue(key)
			deleteRevertValue(key, name, oldv)
		}
	}
}

func saveNewRevertMapping(newConf map[string]*MapValue, name string) error {
	for k, v := range newConf {
		if v.Value == "" {
			continue
		}
		key := reverKey(k, v.Value)
		oldv := getRevertValue(key)
		saveRevertValue(key, name, v.Auth, oldv)
	}
	return nil
}

func (nc *NameConf) Save() error {

	ncdb := &NameConfDb{
		Name:      nc.Name,
		Conf:      make(map[string]*MapValue),
		TimeStamp: nc.TimeStamp,
		UserAddr:  nc.UserAddr,
	}

	for k, v := range nc.Conf {
		ncdb.Conf[k] = &MapValue{Value: v}
	}

	var (
		oldConf *NameConfDb
	)
	oldConf, _ = GetNameConfFromDB(nc.Name)
	if oldConf != nil {
		for k, v := range oldConf.Conf {
			if v.Auth {
				if _, ok := ncdb.Conf[k]; ok && v.Value == ncdb.Conf[k].Value {
					ncdb.Conf[k].Auth = true
				}
			}
		}
	}

	//fmt.Println("oldConf", oldConf.Name, oldConf.Conf)
	if v, err := json.Marshal(*ncdb); err != nil {
		return err
	} else {
		if err = GetLdb().Save([]byte(NameConfKey(ncdb.Name)), v); err != nil {
			return err
		}
	}

	if oldConf == nil || len(oldConf.Conf) == 0 {
		saveNewRevertMapping(ncdb.Conf, ncdb.Name)
	} else {
		saveRevertMapping(ncdb.Conf, oldConf.Conf, ncdb.Name)
	}

	return nil
}

func (ncdb *NameConfDb) Get(key []byte) error {
	if data, err := GetLdb().Get(key); err != nil {
		return err
	} else {
		//fmt.Println(string(key), "11122")
		return json.Unmarshal(data, ncdb)
	}
}

func GetNameConfFromDB(name string) (*NameConfDb, error) {
	ncdb := &NameConfDb{
		Conf: make(map[string]*MapValue),
	}

	if err := ncdb.Get([]byte(NameConfKey(name))); err != nil {
		return nil, err
	}

	return ncdb, nil

}

func hashMsg(content string) (common.Hash, error) {

	hash := crypto.Keccak256Hash([]byte(content))

	fmt.Println("first hash", hex.EncodeToString(hash[:]))

	var sigdata []byte
	sigdata = append(sigdata, []byte("\x19Ethereum Signed Message:\n32")...)
	sigdata = append(sigdata, hash[:]...)

	hashsig := crypto.Keccak256Hash(sigdata)

	return hashsig, nil
}

func (ncs *NameConfWithSig) SaveChanges() error {

	content := &NameConf{
		Conf: make(map[string]string),
	}

	if err := json.Unmarshal([]byte(ncs.Content), content); err != nil {
		return err
	}

	if content.Operation != ReplaceConf {
		return errors.New("operator not support")
	}

	if b := isConfTypes(content.Conf); !b {
		return errors.New("type not support")
	}

	if content.TimeStamp+10 > time.Now().Unix() {
		return errors.New("signature expired")
	}

	if err := ncs.CheckSig(common.HexToAddress(content.UserAddr)); err != nil {
		return err
	}

	if expire, owner, err := FindFromContract(content.Name); err != nil {
		return err
	} else if owner != common.HexToAddress(content.UserAddr) {
		return errors.New("name not belong to user")
	} else if expire {
		return errors.New("name is expired")
	}

	if err := content.Save(); err != nil {
		return err
	}

	return nil
}

func (ncs *NameConfWithSig) CheckSig(userAddr common.Address) error {
	if hash, err := hashMsg(ncs.Content); err != nil {
		return err
	} else {
		var sig []byte
		if sig, err = hex.DecodeString(ncs.Sig); err != nil {
			return err
		}

		if sig[len(sig)-1] >= 27 {
			sig[len(sig)-1] = sig[len(sig)-1] - 27
		}
		var pub *ecdsa.PublicKey
		if pub, err = crypto.SigToPub(hash[:], sig); err != nil {
			return err
		}
		if addr := crypto.PubkeyToAddress(*pub); addr == userAddr {
			fmt.Println("sig matched")
		} else {
			fmt.Println("sig not matched")
			return errors.New("sig not matched")
		}
	}
	return nil
}

func (ncs *NameConfWithSig) SaveAuthEthChanges() error {
	nac := &NameAuthConf{}

	if err := json.Unmarshal([]byte(ncs.Content), nac); err != nil {
		return err
	}

	if nac.Operation != AuthEthConf {
		return errors.New("operator not support")
	}

	if nac.TimeStamp+10 > time.Now().Unix() {
		return errors.New("sig expired")
	}

	if err := ncs.CheckSig(common.HexToAddress(nac.UserAddr)); err != nil {
		return err
	}

	nai := &NameAuthItem{}
	if err := json.Unmarshal([]byte(nac.ItemContent), nai); err != nil {
		return err
	}
	if _, ok := confTypesSet[nai.Type]; !ok {
		return errors.New("not exists")
	}

	if nai.TimeStamp+300 > time.Now().Unix() {
		return errors.New("conf sig expired")
	}

	ncs1 := &NameConfWithSig{
		Content: nac.ItemContent,
		Sig:     nac.ItemSig,
	}
	if err := ncs1.CheckSig(common.HexToAddress(nai.Value)); err != nil {
		return errors.New(err.Error() + " conf sig error")
	}

	if expire, owner, err := FindFromContract(nai.Name); err != nil {
		return err
	} else if owner != common.HexToAddress(nac.UserAddr) {
		return errors.New("name not belong to user")
	} else if expire {
		return errors.New("name is expired")
	}

	//begin to save
	nac.Content = nai
	if err := nac.Save(); err != nil {
		return err
	}

	return nil
}

func (nac *NameAuthConf) Save() error {

	//fmt.Println("begin to save \r\n", nac.String(), "\r\n", nac.Content.String())

	conf, err := GetNameConfFromDB(nac.Content.Name)
	if err != nil {
		return err
	}

	//fmt.Println("from db:\r\n", conf.String())

	if v, ok := conf.Conf[nac.Content.Type]; !ok {
		return errors.New("no config item")
	} else if v.Value != nac.Content.Value {
		return errors.New("value not correct")
	}

	conf.Conf[nac.Content.Type].Auth = true

	//change revermapping
	//todo...
	key := reverKey(nac.Content.Type, nac.Content.Value)
	oldv := getRevertValue(key)
	for i := 0; i < len(oldv); i++ {
		if oldv[i].Value == nac.Content.Name {
			oldv[i].Auth = true
			break
		}
	}

	var data []byte
	if data, err = json.Marshal(oldv); err != nil {
		return err
	}
	if err = GetLdb().Save([]byte(key), data); err != nil {
		return err
	}
	//save
	if data, err = json.Marshal(*conf); err != nil {
		log.Println("fatal error, marshal auth ", nac.Content.Type, nac.Content.Value, nac.Content.Name)
	} else {
		if err = GetLdb().Save([]byte(NameConfKey(nac.Content.Name)), data); err != nil {
			log.Println("fatal error, save auth ", nac.Content.Type, nac.Content.Value, nac.Content.Name)
		}
	}

	return nil

}
