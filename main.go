package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", handleGetRequest)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
}

func handleGetRequest(c *fiber.Ctx) error {
	return c.SendString("Hello, Word!")
}
