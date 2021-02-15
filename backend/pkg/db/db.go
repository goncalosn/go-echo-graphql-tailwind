package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}

func Open() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	Database := client.Database("menu")

	list, err := Database.ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("lista: ", list)
}