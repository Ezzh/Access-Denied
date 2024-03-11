package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

var (
	client *mongo.Client
	db     *mongo.Database
)

func InitDataBase() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

}

func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func Get() Person {
	var result Person
	var err error
	collection := client.Database("test").Collection("people")

	// Вставка документа
	_, err = collection.InsertOne(context.Background(), Person{"John Doe", 30})
	if err != nil {
		log.Fatal(err)
	}

	// Поиск вставленного документа
	err = collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
