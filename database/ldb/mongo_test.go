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
	coll := db.Collection("signcoll")
	indexes := map[string]mongo.IndexModel{
		//"_tokenid_name_": {
		//	Keys: bsonx.Doc{
		//		// 联合索引
		//		// 1 为指定按升序创建索引
		//		{"tokenid", bsonx.Int32(1)},
		//		{"name", bsonx.Int32(1)},
		//	},
		//	//_tokenid_name_ 索引名称
		//	Options: options.Index().SetBackground(true).SetName("_tokenid_name_"),
		//},
		"_name_": {
			Keys: bsonx.Doc{
				{"name", bsonx.Int32(1)},
			},
			//_tokenid_name_ 索引名称
			Options: options.Index().SetBackground(true).SetName("_name_"),
		},
	}
	//查询已经存在的索引
	//cursor, err := coll.Indexes().List(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for cursor.Next(context.Background()) {
	//	var idx []mongo.IndexModel
	//	err = cursor.All(context.TODO(), &idx)
	//	//err = cursor.Decode(&idx)
	//	if err != nil {
	//		log.Println("Decode Index error", err)
	//		continue
	//	}
	//	log.Println(idx)
	//	//if idx != nil {
	//	//	// 删除添加的重复索引
	//	//	//delete(indexes, *idx.Options.Name)
	//	//}
	//}
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
	filter := bson.M{"name": "x.did", "msg_sender": "0x000"}
	var data []map[string]string
	// options.Find().SetSkip(0), options.Find().SetLimit(1) 0页  1每页数据量
	findres, _ := coll.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(1))
	defer findres.Close(context.TODO())
	err = findres.All(context.TODO(), &data)
	if err != nil {
		log.Println("mongo decode err", err)
	}
	if data == nil {
		res, _ := coll.InsertOne(context.TODO(), map[string]string{"name": "x.did", "msg_sender": "0x000", "tokenid": "1"})
		log.Println("InsertOne", res.InsertedID)
	} else {
		up := bson.M{"$set": bson.M{"tokenid": "2"}}
		upres, err := coll.UpdateOne(context.TODO(), filter, up)
		if err != nil {
			log.Println("Mongo Update failed", err)
		} else {
			log.Println("Mongo Update ", upres.UpsertedID)
		}
	}
	// coll.Indexes().CreateOne(context.TODO())

}

func TestSelect(t *testing.T) {
	filter := bson.M{"name": "x.did", "msg_sender": "0x000"}
	options.Find().SetSkip(0)
	options.Find().SetLimit(1)
	op := &options.FindOptions{Skip: func(i int64) *int64 { return &i }(0), Limit: func(i int64) *int64 { return &i }(10)}
	res, err := SELECT("signcoll", filter, op)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}
}
func TestInsertOne(t *testing.T) {
	insert := bson.M{"name": "x.did", "msg_sender": "0x000", "tokenid": "6"}
	INSERTOne("signcoll", insert)
}

func TestUpdate(t *testing.T) {
	filter := bson.M{"name": "x.did", "msg_sender": "0x0002"}
	up := bson.M{"tokenid": "2"}
	state, _ := UPDATEOne("signcoll", filter, up)
	log.Println(state)
}
