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

	coll := client.Database("sample_mflix").Collection("movies")

	title := "Back to the Future"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)
}
