package main

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/routes"
)

func main() {
	mongoClient := connectToDatabase()
	server := createServer()
	server.addMiddlewares()
	routes.SetupAppUserRoutes(server.App, mongoClient)
	routes.SetupCampaignRoutes(server.App, mongoClient)
	routes.SetupEventRoutes(server.App, mongoClient)
	server.start()
}
