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
	CreateObject() interface{}
	UpdateObject() interface{}
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
	//TODO implement me
	panic("implement me")
}

func (mc *Client) GetBulkObject(collection string, limit int, offset int) (*mongo.Cursor, int) {
	response, err := mc.Worker.Database(mc.DBName).Collection(collection).Find(context.TODO(), bson.D{})

	if err != nil {
		log.Println("Error happened during getting list of objects")
	}
	count, err := mc.Worker.Database(mc.DBName).Collection(collection).CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		log.Println("Error happened during counting objects")
	}
	return response, int(count)
}

func (mc *Client) CreateObject() interface{} {
	//TODO implement me
	panic("implement me")
}

func (mc *Client) UpdateObject() interface{} {
	//TODO implement me
	panic("implement me")
}

func (mc *Client) DeleteObject() interface{} {
	//TODO implement me
	panic("implement me")
}
