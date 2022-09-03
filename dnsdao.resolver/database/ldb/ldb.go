package ldb

import (
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strconv"
	"sync"
)

const (
	dnsResolveBlKnumber    = "dnsResolveBlKnumber"
	dnsResolveTypes        = "dnsResolveTypes"
	dnsRootNameHead        = "rnH_1_%s"
	dnsSubNameHead         = "snH_1_%s"
	dnsContract            = "dnsOpenContract"
	dnsAddress             = "dnsAddress"
	dnsRootEarnings        = "dnsRootEarnings"
	dnsAddressCash         = "dnsAddressCash"
	dnsContractAddr        = "dnsContractAddr"
	dnsConf                = "dnsConf"
	dnsContractTokenIdName = "dnsContractTokenIdName"
	dnsNftPass             = "dnsNftPass"
)

type Ldb struct {
	db *leveldb.DB
}

var (
	ldb     *Ldb
	ldbOnce sync.Once
)

func GetLdb() *Ldb {
	if ldb == nil {
		ldbOnce.Do(func() {
			opts := opt.Options{
				Strict:      opt.DefaultStrict,
				Compression: opt.NoCompression,
				Filter:      filter.NewBloomFilter(10),
			}
			db, err := leveldb.OpenFile(config.GetRConf().GetDBPath(), &opts)
			if err != nil {
				panic(err)
			}
			ldb = &Ldb{
				db: db,
			}

		})
	}

	return ldb
}

func (db *Ldb) CloseLdb() {
	if ldb != nil {
		db.db.Close()
	}
	ldb = nil
}

func (db *Ldb) GetLatestBlkNum(network string) (blk uint64) {
	if v, err := db.db.Get([]byte(dnsResolveBlKnumber+network), nil); err != nil {
		return 0
	} else {
		if blk, err = strconv.ParseUint(string(v), 10, 64); err != nil {
			return 0
		}
	}
	return
}

func (db *Ldb) SaveLatestBlkNum(network string, blk uint64) error {

	if err := db.db.Put([]byte(dnsResolveBlKnumber+network), []byte(strconv.FormatUint(blk, 10)), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}

	return nil
}

func (db *Ldb) SaveRootName(name string, val string) error {
	k := fmt.Sprintf(dnsRootNameHead, name)
	if err := db.db.Put([]byte(k), []byte(val), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}

	return nil
}

func (db *Ldb) GetRootName(name string) (string, error) {
	k := fmt.Sprintf(dnsRootNameHead, name)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) SaveSubName(name string, val string) error {
	k := fmt.Sprintf(dnsSubNameHead, name)
	if err := db.db.Put([]byte(k), []byte(val), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}

	return nil
}

func (db *Ldb) GetSubName(name string) (string, error) {
	k := fmt.Sprintf(dnsSubNameHead, name)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) SaveTypeNames(ts []string) error {
	v, _ := json.Marshal(ts)
	if err := db.db.Put([]byte(dnsResolveTypes), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetTypeNames() ([]string, error) {
	var ret []string
	if v, err := db.db.Get([]byte(dnsResolveTypes), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveTypeNameHash(conf_data string, name string, name_type string) error {
	var val string
	if name_type == "root" {
		val = fmt.Sprintf(dnsRootNameHead, name)
	} else {
		val = fmt.Sprintf(dnsSubNameHead, name)
	}
	if err := db.db.Put([]byte(conf_data), []byte(val), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil

}

// doamin conf get namehash
func (db *Ldb) GetTypeNameHash(conf_data string) (string, error) {
	if v, err := db.db.Get([]byte(conf_data), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) SavaOwnerNameHash(addr string, val []string) error {
	j, _ := json.Marshal(val)
	if err := db.db.Put([]byte(addr), []byte(j), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

// addr  get namehashlist
func (db *Ldb) GetOwnerNameHash(addr string) ([]string, error) {
	var ret []string
	if v, err := db.db.Get([]byte(addr), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}
		return ret, nil
	}
}

// get key
func (db *Ldb) GetKey(key string) (string, error) {
	if v, err := db.db.Get([]byte(key), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) GetContractList() ([]string, error) {
	var ret []string
	if v, err := db.db.Get([]byte(dnsContract), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveContractList(ct []string) error {
	v, _ := json.Marshal(ct)
	if err := db.db.Put([]byte(dnsContract), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetAddressList(address string) ([]string, error) {
	var ret []string
	k := fmt.Sprintf(dnsAddress, address)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveAddressList(address string, ct []string) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsAddress, address)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetRootEarnings(address string) (*AddressEarnings, error) {
	var ret *AddressEarnings
	k := fmt.Sprintf(dnsRootEarnings, address)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveRootEarnings(address string, ct *AddressEarnings) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsRootEarnings, address)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetAddressCash(address string) (*AddressCash, error) {
	var ret *AddressCash
	k := fmt.Sprintf(dnsAddressCash, address)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveAddressCash(address string, ct *AddressCash) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsAddressCash, address)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetContractAddr(contract string) (string, error) {
	k := fmt.Sprintf(dnsContractAddr, contract)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) SaveContractAddr(contract string, addr string) error {
	k := fmt.Sprintf(dnsContractAddr, contract)
	if err := db.db.Put([]byte(k), []byte(addr), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	} else {
		return nil
	}
}

func (db *Ldb) GetConfNameHashList(conftype string) ([]string, error) {
	var ret []string
	k := fmt.Sprintf(dnsConf, conftype)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveConfNameHashList(conftype string, ct []string) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsConf, conftype)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetContractTokenIdName(contract string) (*ContractTokenIdName, error) {
	var ret *ContractTokenIdName
	k := fmt.Sprintf(dnsContractTokenIdName, contract)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveContractTokenIdName(contract string, ct *ContractTokenIdName) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsContractTokenIdName, contract)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetS3() (*S3, error) {
	var ret *S3
	k := fmt.Sprintf("S3")
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveS3(ct *S3) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf("S3")
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetNftPass(owner string) ([]*NftPass, error) {
	var ret []*NftPass
	k := fmt.Sprintf(dnsNftPass, owner)
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveNftPass(owner string, ct []*NftPass) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf(dnsNftPass, owner)
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}
