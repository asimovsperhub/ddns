package ldb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"testing"
)

func createIndexes() error {
	var err error
	db := GetMdb().db
	log.Println(db.Name())
	coll := db.Collection("signcoll")
	indexes := map[string]mongo.IndexModel{
		"_tokenid_name_": {
			Keys: bsonx.Doc{
				// 联合索引
				// 1 为指定按升序创建索引
				{"tokenid", bsonx.Int32(1)},
				{"name", bsonx.Int32(1)},
			},
			//_tokenid_name_ 索引名称
			Options: options.Index().SetBackground(true).SetName("_tokenid_name_"),
		},
	}
	//查询已经存在的索引
	cursor, err := coll.Indexes().List(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for cursor.Next(context.Background()) {
		var idx []mongo.IndexModel
		err = cursor.All(context.TODO(), &idx)
		//err = cursor.Decode(&idx)
		if err != nil {
			log.Println("Decode Index error", err)
			continue
		}
		log.Println(idx)
		//if idx != nil {
		//	// 删除添加的重复索引
		//	//delete(indexes, *idx.Options.Name)
		//}
	}
	// 添加索引
	models := []mongo.IndexModel{}
	for _, v := range indexes {
		models = append(models, v)
	}
	if len(models) > 0 {
		var names []string
		names, err = coll.Indexes().CreateMany(context.Background(), models)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Create Indexes", names)
	}
	return err
}

func TestIndex(t *testing.T) {
	createIndexes()
}
func TestMongo(t *testing.T) {
	var err error
	db := GetMdb().db
	log.Println(db.Name())
	coll := db.Collection("signcoll")

	res, _ := coll.InsertOne(context.TODO(), map[string]string{"tokenid": "1", "msg_sender": "0x000", "name": "x.did"})
	log.Println(res)
	filter := bson.M{"tokenid": "1", "name": "x.did"}
	var data []map[string]string
	// options.Find().SetSkip(0), options.Find().SetLimit(1) 0页  1每页数据量
	findres, _ := coll.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(10))
	defer findres.Close(context.TODO())
	//log.Println(findres)
	if findres != nil {
		err = findres.All(context.TODO(), &data)
		if err != nil {
			log.Println("mongo decode err", err)
		}
		log.Println(findres, data)
	} else {
		log.Println("not find data")
	}
	// coll.Indexes().CreateOne(context.TODO())

}
