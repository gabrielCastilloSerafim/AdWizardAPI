package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupEventRoutes(app *fiber.App, mongoClient *mongo.Client) {

	app.Post("/event", utils.MakeHTTPHandleFunc(controllers.HandleCreateEvent, mongoClient))
	app.Get("/event", utils.MakeHTTPHandleFunc(controllers.HandleGetAllEvents, mongoClient))
	app.Delete("/event", utils.MakeHTTPHandleFunc(controllers.HandleDeleteAllEvents, mongoClient))
	app.Get("/event/ping", utils.MakeHTTPHandleFunc(controllers.HandlePing, mongoClient))
	app.Get("/event/:campaignId", utils.MakeHTTPHandleFunc(controllers.HandleGetEventByCampaignId, mongoClient))
}
