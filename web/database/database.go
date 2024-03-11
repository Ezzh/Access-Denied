package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SCP struct {
	Name            string
	DescryptionPath string
	ImagePath       string
}

var (
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
)

func InitDataBase() {
	var err error
	ctx = context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err = mongo.Connect(ctx, clientOptions)
	db = client.Database("Archive")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
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

func GetByName(name string) SCP {
	var result SCP
	var err error
	collection := GetCollection("SCPs")
	if err != nil {
		log.Fatal(err)
	}

	err = collection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetAll() []SCP {
	var result []SCP
	var err error
	collection := GetCollection("SCPs")
	if err != nil {
		log.Fatal(err)
	}

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result
}
