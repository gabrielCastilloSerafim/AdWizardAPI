package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDatabase() {

	var uri string

	if os.Getenv("RAILWAY") == "TRUE" {

		uri = os.Getenv("MONGO_PRIVATE_URL")
		if uri == "" {
			log.Fatal("You must set your 'MONGO_PRIVATE_URL'")
		}
	} else {

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading environment variables file")
		}

		uri = os.Getenv("MONGO_DB_URI_FROM_LOCAL")
		if uri == "" {
			log.Fatal("You must set your 'MONGO_PRIVATE_URL'")
		}
	}

	// Get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Test retreiving some data
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

	jsonData, err := json.MarshalIndent(campaings, "", "    ")
	if err != nil {
		panic(err)
	}

	log.Default().Printf("%s\n", jsonData)
}
