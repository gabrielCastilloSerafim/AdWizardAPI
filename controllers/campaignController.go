package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleCreateCampaign(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Parse req body into campaign object
	campaign := new(models.Campaign)
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
	campaignURL := fmt.Sprintf("https://adwizardapi-production.up.railway.app/campaign/redirect/%v", campaignId)
	response := fiber.Map{
		"campaignURL": campaignURL,
	}
	return c.JSON(response)
}

func HandleStoreRedirect(c *fiber.Ctx, mongoClient *mongo.Client) error {
	// Get IP "If localhost this field will arrive empty" and campaign Id
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	campaignId := c.Params("campaignId")
	// Generate user
	appUser := new(models.AppUser)
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
	campaign := new(models.Campaign)
	campaignMatch.Decode(campaign)
	return c.Redirect(campaign.AppStoreLink)
}

func HandleGetAllCampaigns(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("Campaign")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var campaings []models.Campaign
	err = cursor.All(ctx, &campaings)
	if err != nil {
		return err
	}
	return c.JSON(campaings)
}

func HandleDeleteAllCampaigns(c *fiber.Ctx, mongoClient *mongo.Client) error {
	ctx := context.Background()
	collection := mongoClient.Database("production").Collection("AppUser")
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}
	return c.SendString("All AppUsers Deleted")
}
