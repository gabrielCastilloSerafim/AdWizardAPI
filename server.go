package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Campaing struct {
	CampaingId string `json:"campaingId"`
}

func startServer() {

	app := fiber.New()

	setupMiddleWares(app)
	setupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func setupMiddleWares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(requestid.New())
}

func setupRoutes(app *fiber.App) {
	app.Get("/helloLosPibes", handleHelloLosPibes)
	app.Get("/ping", handlePing)
	app.Get("/store", handleStoreRedirect)
	app.Post("/campaing", handleCampaingEvent)
}

func handleCampaingEvent(c *fiber.Ctx) error {

	campaing := new(Campaing)
	err := c.BodyParser(campaing)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	log.Default().Println(campaing.CampaingId)

	return c.Status(fiber.StatusOK).JSON(campaing)
}

func handlePing(c *fiber.Ctx) error {
	return nil
}

func handleStoreRedirect(c *fiber.Ctx) error {
	return c.Redirect("https://apps.apple.com/app/group-task-")
}

func handleHelloLosPibes(c *fiber.Ctx) error {
	return c.SendString("Hola los pibes!")
}
