package ldb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type Mdb struct {
	db *mongo.Database
}

var (
	mdb     *Mdb
	mdbOnce sync.Once
)

func GetMdb() *Mdb {
	if mdb == nil {
		mdbOnce.Do(func() {
			var err error
			clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
			// 连接到MongoDB
			mgoCli, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatal(err)
			}
			// 检查连接
			err = mgoCli.Ping(context.TODO(), nil)
			if err != nil {
				log.Fatal(err)
			}
			db := mgoCli.Database("mintsign")
			mdb = &Mdb{
				db: db,
			}
		})
	}

	return mdb
}
