package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(app *fiber.App, db storage.StorageInterface) {

	app.Post("/event", utils.MakeHTTPHandleFunc(controllers.HandleCreateEvent, db))
	app.Get("/event", utils.MakeHTTPHandleFunc(controllers.HandleGetAllEvents, db))
	app.Delete("/event", utils.MakeHTTPHandleFunc(controllers.HandleDeleteAllEvents, db))
	app.Get("/event/ping", utils.MakeHTTPHandleFunc(controllers.HandlePing, db))
	app.Get("/event/:campaignId", utils.MakeHTTPHandleFunc(controllers.HandleGetEventByCampaignId, db))
}
