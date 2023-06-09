package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type IMongoClient interface {
	GetObject(filter interface{}) interface{}
	GetBulkObject(collection string, limit int, offset int) (*mongo.Cursor, int)
	GetRandomObjects(collection string, size int) (*mongo.Cursor, int)
	CreateObject(collection string, model interface{}) *mongo.SingleResult
	UpdateObjectCounter(collection string, uuid string) *mongo.SingleResult
	UpdateObject(collection string, uuid string) *mongo.SingleResult
	DeleteObject() interface{}
}

type Client struct {
	Worker *mongo.Client
	DBName string
}

var ClientInstance = &Client{}

func InitDB(mongoUrl string, dbName string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database.\n%+v", err))
	}
	if ClientInstance.Worker == nil {
		ClientInstance.Worker = client
		ClientInstance.DBName = dbName
	}
}

func GetMongoClient() IMongoClient {
	return ClientInstance
}

func (mc *Client) GetObject(filter interface{}) interface{} {
	return nil
}

func (mc *Client) GetBulkObject(collection string, limit int, offset int) (*mongo.Cursor, int) {
	params := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	response, err := mc.Worker.Database(mc.DBName).Collection(collection).Find(context.TODO(), bson.D{}, params)

	if err != nil {
		log.Println("Error happened during getting list of objects.")
	}
	count, err := mc.Worker.Database(mc.DBName).Collection(collection).CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		log.Println("Error happened during counting objects.")
	}
	return response, int(count)
}

func (mc *Client) GetRandomObjects(collection string, size int) (*mongo.Cursor, int) {
	pipeline := []bson.D{{{"$sample", bson.D{{"size", size}}}}}
	response, err := mc.Worker.Database(mc.DBName).Collection(collection).Aggregate(context.TODO(), pipeline)

	if err != nil {
		log.Println("Error happened during random objects get.")
	}

	return response, size
}

func (mc *Client) CreateObject(collection string, model interface{}) *mongo.SingleResult {
	insertResponse, err := mc.Worker.Database(mc.DBName).Collection(collection).InsertOne(context.TODO(), model)

	if err != nil {
		log.Println("Error happened during inserting object.")
	}

	response := mc.Worker.Database(mc.DBName).Collection(collection).FindOne(context.TODO(), bson.D{{"_id", insertResponse.InsertedID}})

	return response
}

func (mc *Client) UpdateObject(collection string, uuid string) *mongo.SingleResult {
	//TODO implement me
	panic("implement me")
}

func (mc *Client) UpdateObjectCounter(collection string, uuid string) *mongo.SingleResult {
	// TODO: refactor this shitty request
	response := mc.Worker.Database(mc.DBName).Collection(collection).FindOneAndUpdate(
		context.TODO(), bson.M{"uuid": uuid}, bson.D{
			{"$inc", bson.D{{"counter", 1}}},
		}, options.FindOneAndUpdate().SetReturnDocument(options.After))

	return response
}

func (mc *Client) DeleteObject() interface{} {
	//TODO implement me
	panic("implement me")
}
