package controllers

import (
	"fmt"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gofiber/fiber/v2"
)

func HandleCreateCampaign(c *fiber.Ctx, db storage.StorageInterface) error {
	// Parse req body into campaign object
	campaign := new(models.Campaign)
	err := c.BodyParser(campaign)
	if err != nil {
		return err
	}
	// Save to database
	campaignId, err := db.CreateCampaign(campaign)
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

func HandleStoreRedirect(c *fiber.Ctx, db storage.StorageInterface) error {
	// Get IP "If localhost this field will arrive empty" and campaign Id
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	campaignId := c.Params("campaignId")
	// Generate user
	appUser := new(models.AppUser)
	appUser.CampaignId = campaignId
	if len(ip) > 0 {
		appUser.UserIp = ip[0]
	} else {
		appUser.UserIp = "localhost"
	}
	// Store user
	err := db.CreateAppUser(appUser)
	if err != nil {
		return err
	}
	// Find store link and redirect
	storeLink, err := db.GetCampaignById(campaignId)
	if err != nil {
		return err
	}
	return c.Redirect(storeLink)
}

func HandleGetAllCampaigns(c *fiber.Ctx, db storage.StorageInterface) error {
	campaings, err := db.GetAllCampaigns()
	if err != nil {
		return err
	}
	return c.JSON(campaings)
}

func HandleDeleteAllCampaigns(c *fiber.Ctx, db storage.StorageInterface) error {
	err := db.DeleteAllCampaigns()
	if err != nil {
		return err
	}
	return c.SendString("All campaigns deleted")
}
