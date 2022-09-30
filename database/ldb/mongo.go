package ldb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
			// 建索引

		})
	}

	return mdb
}
func SELECT(Table string, Filter bson.M, OptionsFind *options.FindOptions) ([]map[string]string, error) {
	var err error
	db := GetMdb().db
	coll := db.Collection(Table)
	var data []map[string]string
	findres, err := coll.Find(context.TODO(), Filter, OptionsFind)
	defer findres.Close(context.TODO())
	if err != nil {
		log.Println(fmt.Sprintf("Find table:%s Filter %s  err:%s", Table, Filter, err))
		return nil, err
	}
	defer findres.Close(context.TODO())
	err = findres.All(context.TODO(), &data)
	if err != nil {
		log.Println("mongo decode err", err)
		return nil, err
	} else {
		return data, nil
	}
}

func UPDATEOne(Table string, Filter bson.M, UpdateData bson.M) (bool, error) {
	var err error
	db := GetMdb().db
	coll := db.Collection(Table)
	var upres *mongo.UpdateResult
	up := bson.M{"$set": UpdateData}
	upres, err = coll.UpdateOne(context.TODO(), Filter, up)
	if err != nil {
		log.Println("Mongo Update failed", err)
		return false, err
	} else {
		log.Println("Mongo Update success", upres.UpsertedID)
		return true, nil
	}
}

func INSERTOne(Table string, InsertData bson.M) {
	var err error
	db := GetMdb().db
	coll := db.Collection(Table)
	var res *mongo.InsertOneResult
	res, err = coll.InsertOne(context.TODO(), InsertData)
	if err != nil {
		log.Println(fmt.Sprintf("Table %s InsertData %v err %v", Table, InsertData, err))
	} else {
		log.Println(fmt.Sprintf("Table %s InsertData %v success %v", Table, InsertData, res.InsertedID))
	}
}
