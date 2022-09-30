package dnsconf

import (
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/filter"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/config"

	"sync"
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
			db, err := leveldb.OpenFile(config.GetDnsConfSetting().GetDbPath(), &opts)
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

func (db *Ldb) Save(key, value []byte) error {
	return db.db.Put(key, value, &opt.WriteOptions{
		Sync: true,
	})
}

func (db *Ldb) Del(key []byte) error {
	return db.db.Delete(key, &opt.WriteOptions{Sync: true})
}

func (db *Ldb) Get(key []byte) (value []byte, err error) {
	return db.db.Get(key, nil)
}
