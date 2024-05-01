package database

import (
	"context"
	"log"
	"web/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
}

func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func GetByName(name string) models.SCP {
	var result models.SCP
	var err error
	collection := GetCollection("SCPs")
	if err != nil {
		log.Fatal(err)
	}

	err = collection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
	if err != nil {
		log.Default()
		return models.SCP{}
	}

	return result
}

func GetAll() []models.SCP {
	var results []models.SCP
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
		var elem models.SCP
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	return results
}

func CreateUser(user models.User) {
	collection := GetCollection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateSCP(SCP models.SCP) {
	collection := GetCollection("SCPs")
	_, err := collection.InsertOne(ctx, SCP)
	if err != nil {
		log.Fatal(err)
	}
}
