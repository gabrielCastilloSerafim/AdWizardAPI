package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupCampaignRoutes(app *fiber.App, db storage.StorageInterface) {

	app.Post("/campaign", utils.MakeHTTPHandleFunc(controllers.HandleCreateCampaign, db))
	app.Get("/campaign", utils.MakeHTTPHandleFunc(controllers.HandleGetAllCampaigns, db))
	app.Delete("/campaign", utils.MakeHTTPHandleFunc(controllers.HandleDeleteAllCampaigns, db))
	app.Get("/campaign/redirect/:campaignId", utils.MakeHTTPHandleFunc(controllers.HandleStoreRedirect, db))
}
