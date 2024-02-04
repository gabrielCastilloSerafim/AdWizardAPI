package controllers

import (
	"context"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleGetAllAppUsers(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("AppUser")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var appUsers []models.AppUser
	err = cursor.All(ctx, &appUsers)
	if err != nil {
		return err
	}
	return c.JSON(appUsers)
}
