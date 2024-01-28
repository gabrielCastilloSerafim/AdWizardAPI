package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server start ..

func startServer(mongoClient *mongo.Client) {

	app := fiber.New()
	port := os.Getenv("PORT")

	// Middlewares
	app.Use(logger.New())
	app.Use(requestid.New())

	// Routes
	app.Post("/campaign", makeHTTPHandleFunc(handleCreateCampaign, mongoClient))
	app.Get("/campaign", makeHTTPHandleFunc(handleGetAllCampaigns, mongoClient))
	app.Get("/store/:campaignId", makeHTTPHandleFunc(handleStoreRedirect, mongoClient))
	app.Post("/event", makeHTTPHandleFunc(handleCreateEvent, mongoClient))
	app.Get("/event", makeHTTPHandleFunc(handleGetAllEvents, mongoClient))
	app.Get("/appuser", makeHTTPHandleFunc(handleGetAllAppUsers, mongoClient))
	app.Get("/ping", makeHTTPHandleFunc(handlePing, mongoClient))

	// Listen
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

// Handlers ..

func handleCreateCampaign(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Parse req body into campaign object
	campaign := new(Campaign)
	err := c.BodyParser(campaign)
	if err != nil {
		return err
	}
	// Save to database
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Campaign")
	insertResult, err := collection.InsertOne(ctx, campaign)
	campaignId := (insertResult.InsertedID.(primitive.ObjectID)).Hex()
	if err != nil {
		return err
	}
	// Generate and send URL to be used in campaign
	campaignURL := fmt.Sprintf("https://adwizardapi-production.up.railway.app/store/%v", campaignId)
	response := fiber.Map{
		"campaignURL": campaignURL,
	}
	return c.JSON(response)
}

func handleStoreRedirect(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Get IP "If localhost this field will arrive empty" and campaign Id
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	campaignId := c.Params("campaignId")
	// Generate user
	appUser := new(AppUser)
	appUser.CampaignId = campaignId
	if len(ip) > 0 {
		appUser.UserIp = ip[0]
	} else {
		appUser.UserIp = "localhost" // Only using this to test in localhost, `remove later``
		log.Default().Println("Could not find user ip from request header")
	}
	// Store user
	ctx := context.Background()
	appUserCollection := mongoClient.Database("production").Collection("AppUser")
	_, err := appUserCollection.InsertOne(ctx, appUser)
	if err != nil {
		log.Default().Printf("Error inserting user into db: %v", err)
		return err
	}
	// Find store link and redirect
	objId, err := primitive.ObjectIDFromHex(campaignId)
	if err != nil {
		return err
	}
	campaignCollection := mongoClient.Database("production").Collection("Campaign")
	campaignMatch := campaignCollection.FindOne(ctx, bson.M{"_id": objId})
	campaign := new(Campaign)
	campaignMatch.Decode(campaign)
	return c.Redirect(campaign.AppStoreLink)
}

func handlePing(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Get IP "If localhost this field will arrive empty"
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	// Perform match from ip
	appUser := new(AppUser)
	ctx := context.Background()
	if len(ip) > 0 {
		userCollection := mongoClient.Database("production").Collection("AppUser")
		userMatch := userCollection.FindOne(ctx, bson.M{"userIp": ip[0]})
		userMatch.Decode(appUser)
		log.Default().Println(userMatch)
	} else {
		log.Default().Println("Could not find user ip from request header")
		// only used to test with localhost, `remove later``
		// userCollection := mongoClient.Database("production").Collection("AppUser")
		// userMatch := userCollection.FindOne(ctx, bson.M{"userIp": "localhost"})
		// userMatch.Decode(user)
	}
	// Create and store an download event
	downloadEvent := new(Event)
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

func handleCreateEvent(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Parse request body
	event := new(Event)
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

func handleGetAllCampaigns(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Campaign")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var campaings []Campaign
	err = cursor.All(ctx, &campaings)
	if err != nil {
		return err
	}
	return c.JSON(campaings)
}

func handleGetAllAppUsers(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("AppUser")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var appUsers []AppUser
	err = cursor.All(ctx, &appUsers)
	if err != nil {
		return err
	}
	return c.JSON(appUsers)
}

func handleGetAllEvents(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Event")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var events []Event
	err = cursor.All(ctx, &events)
	if err != nil {
		return err
	}
	return c.JSON(events)
}
