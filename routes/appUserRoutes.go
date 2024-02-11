package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupAppUserRoutes(app *fiber.App, db storage.StorageInterface) {

	app.Get("/appuser", utils.MakeHTTPHandleFunc(controllers.HandleGetAllAppUsers, db))
	app.Delete("/appuser", utils.MakeHTTPHandleFunc(controllers.HandleDeleteAllAppUsers, db))
}
