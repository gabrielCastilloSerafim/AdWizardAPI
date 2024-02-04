package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupCampaignRoutes(app *fiber.App, mongoClient *mongo.Client) {

	app.Post("/campaign", utils.MakeHTTPHandleFunc(controllers.HandleCreateCampaign, mongoClient))
	app.Get("/campaign", utils.MakeHTTPHandleFunc(controllers.HandleGetAllCampaigns, mongoClient))
	app.Get("/campaign/redirect/:campaignId", utils.MakeHTTPHandleFunc(controllers.HandleStoreRedirect, mongoClient))
}
