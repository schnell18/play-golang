package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cred := options.Credential{
		Username: "localenv",
		Password: "localenv",
	}
	options := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(cred)
	client, err := mongo.Connect(ctx, options)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Unable to ping mongodb due to:%v", err)
	}

	collection := client.Database("testing").Collection("numbers")
	document := bson.D{
		{Key: "name", Value: "pi"},
		{Key: "value", Value: 3.14159},
	}
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatalf("Fail to insert document %v due to %v", document, err)
	}
	fmt.Printf("Inserted document: %v with _id: %v", document, res.InsertedID)
}
