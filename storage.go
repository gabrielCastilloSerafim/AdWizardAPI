package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Campaign struct {
	Id           string    `json:"_id"`
	Name         string    `json:"name"`
	Source       string    `json:"source"`
	AppStoreLink string    `json:"appStoreLink"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Event struct {
	Id         string    `json:"_id"`
	Name       string    `json:"name"`
	CampaignId string    `json:"campaignId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdateAt   time.Time `json:"updateAt"`
	UserId     string    `json:"userId"`
}

type AppUser struct {
	Id         string    `json:"_id"`
	CampaignId string    `json:"campaignId"`
	UserIp     string    `json:"userIp"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func setupDatabase() {

	uri := "mongodb+srv://adwizard-admin:3HOOQRUQLl7RQpQk@adwizardcluster.uuyu7mt.mongodb.net/admin?authSource=admin&replicaSet=atlas-4e258e-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"

	// if os.Getenv("RAILWAY") == "TRUE" {

	// 	uri = os.Getenv("MONGO_PRIVATE_URL")
	// 	if uri == "" {
	// 		log.Fatal("You must set your 'MONGO_PRIVATE_URL'")
	// 	}
	// } else {

	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatalf("Error loading environment variables file")
	// 	}

	// 	uri = os.Getenv("MONGO_DB_URI_FROM_LOCAL")
	// 	if uri == "" {
	// 		log.Fatal("You must set your 'MONGO_PRIVATE_URL'")
	// 	}
	// }

	// Get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Test retreiving some data
	ctx := context.Background()
	collection := client.Database("production").Collection("Campaign")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
	}

	var campaings []Campaign

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
