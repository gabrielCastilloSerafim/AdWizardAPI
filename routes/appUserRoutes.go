package routes

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/controllers"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAppUserRoutes(app *fiber.App, mongoClient *mongo.Client) {

	app.Get("/appuser", utils.MakeHTTPHandleFunc(controllers.HandleGetAllAppUsers, mongoClient))
	app.Delete("/appuser", utils.MakeHTTPHandleFunc(controllers.HandleDeleteAllAppUsers, mongoClient))
}
