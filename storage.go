package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createDatabase() *mongo.Client {

	uri := "mongodb+srv://adwizard-admin:3HOOQRUQLl7RQpQk@adwizardcluster.uuyu7mt.mongodb.net/admin?authSource=admin&replicaSet=atlas-4e258e-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"

	// Get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
