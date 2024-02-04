package controllers

import (
	"context"
	"log"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlePing(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Get IP "If localhost this field will arrive empty"
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	// Perform match from ip
	appUser := new(models.AppUser)
	ctx := context.Background()
	if len(ip) > 0 {
		userCollection := mongoClient.Database("production").Collection("AppUser")
		userMatch := userCollection.FindOne(ctx, bson.M{"userIp": ip[0]})
		userMatch.Decode(appUser)
	} else {
		log.Default().Println("Could not find user ip from request header")
		// only used to test with localhost, `remove later`
		// userCollection := mongoClient.Database("production").Collection("AppUser")
		// userMatch := userCollection.FindOne(ctx, bson.M{"userIp": "localhost"})
		// userMatch.Decode(appUser)
	}
	// Create and store an download event
	downloadEvent := new(models.Event)
	downloadEvent.CampaignId = appUser.CampaignId
	downloadEvent.UserId = appUser.Id.Hex()
	downloadEvent.Name = "download"
	eventCollection := mongoClient.Database("production").Collection("Event")
	_, err := eventCollection.InsertOne(ctx, downloadEvent)
	if err != nil {
		log.Default().Printf("Error inserting user into db: %v", err)
		return err
	}
	// Send back the userId
	response := fiber.Map{
		"userId":     appUser.Id.Hex(),
		"campaignId": appUser.CampaignId,
	}
	return c.JSON(response)
}

func HandleCreateEvent(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Parse request body
	event := new(models.Event)
	err := c.BodyParser(event)
	if err != nil {
		return err
	}
	// Save to database
	ctx := context.Background()
	eventCollection := mongoClient.Database("production").Collection("Event")
	_, err = eventCollection.InsertOne(ctx, event)
	if err != nil {
		log.Default().Printf("Error inserting user into db: %v", err)
		return err
	}
	return nil
}

func HandleGetAllEvents(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Event")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var events []models.Event
	err = cursor.All(ctx, &events)
	if err != nil {
		return err
	}
	return c.JSON(events)
}

func HandleGetEventByCampaignId(c *fiber.Ctx, mongoClient *mongo.Client) error {
	campaignId := c.Params("campaignId")
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Event")
	cursor, err := collection.Find(ctx, bson.M{"campaignId": campaignId})
	if err != nil {
		return err
	}
	var events []models.Event
	err = cursor.All(ctx, &events)
	if err != nil {
		return err
	}
	return c.JSON(events)
}

func HandleDeleteAllEvents(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("AppUser")
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}
	return c.SendString("All AppUsers Deleted")
}
