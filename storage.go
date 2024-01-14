package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDatabase() {

	uri := os.Getenv("MONGO_PRIVATE_URL")

	if uri == "" {
		log.Fatal("You must set your 'MONGO_PRIVATE_URL'")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	collection := client.Database("test").Collection("events")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
	}

	var campaings []Campaing

	err = cursor.All(ctx, &campaings)
	if err != nil {
		log.Panic(err)
	}

	out, err := json.MarshalIndent(campaings, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Result:", string(out))

	if err == mongo.ErrNoDocuments {
		fmt.Println("No document named: events")
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(campaings, "", "    ")
	if err != nil {
		panic(err)
	}

	log.Printf("%s\n", jsonData)
}
