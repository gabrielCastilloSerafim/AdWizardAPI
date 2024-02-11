package main

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/routes"
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
)

func main() {
	mongoClient := connectToDatabase()
	mongoStrorage := storage.InitMongoStorage(mongoClient)
	server := createServer()
	server.addMiddlewares()
	routes.SetupAppUserRoutes(server.App, mongoStrorage)
	routes.SetupCampaignRoutes(server.App, mongoStrorage)
	routes.SetupEventRoutes(server.App, mongoStrorage)
	server.start()
}
