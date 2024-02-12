package controllers

import (
	"log"

	"github.com/gabrielCastilloSerafim/AdWizardAPI/models"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gofiber/fiber/v2"
)

func HandlePing(c *fiber.Ctx, db storage.StorageInterface) error {
	// Get IP "If localhost this field will arrive empty"
	ip := c.GetReqHeaders()["X-Forwarded-For"]
	log.Default().Printf("IP RECEIVED -> %v\n", ip[0])
	// Perform match from ip
	appUser := new(models.AppUser)
	if len(ip) > 0 {
		err := db.GetAppUserWithIp(ip[0], appUser)
		if err != nil {
			return err
		}
	} else {
		err := db.GetAppUserWithIp("localhost", appUser)
		if err != nil {
			return err
		}
	}
	// Create event
	downloadEvent := models.Event{
		CampaignId: appUser.CampaignId,
		UserId:     appUser.Id.Hex(),
		Name:       "download",
	}
	// Save event to database
	err := db.CreateEvent(&downloadEvent)
	if err != nil {
		return err
	}
	// Send response
	response := fiber.Map{
		"userId":     appUser.Id.Hex(),
		"campaignId": appUser.CampaignId,
	}
	return c.JSON(response)
}

func HandleCreateEvent(c *fiber.Ctx, db storage.StorageInterface) error {
	// Parse request body
	event := new(models.Event)
	err := c.BodyParser(event)
	if err != nil {
		return err
	}
	// Save to database
	err = db.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func HandleGetAllEvents(c *fiber.Ctx, db storage.StorageInterface) error {
	events, err := db.GetAllEvents()
	if err != nil {
		return err
	}
	return c.JSON(events)
}

func HandleGetEventByCampaignId(c *fiber.Ctx, db storage.StorageInterface) error {
	campaignId := c.Params("campaignId")
	event, err := db.GetEventByCampaignId(campaignId)
	if err != nil {
		return err
	}
	return c.JSON(event)
}

func HandleDeleteAllEvents(c *fiber.Ctx, db storage.StorageInterface) error {
	err := db.DeleteAllEvents()
	if err != nil {
		return err
	}
	return c.SendString("All Events Deleted")
}
