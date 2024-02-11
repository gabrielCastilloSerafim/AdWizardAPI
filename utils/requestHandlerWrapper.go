package utils

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gofiber/fiber/v2"
)

type serverFuncWithMongoClient func(c *fiber.Ctx, db storage.StorageInterface) error
type fiberHandlerFunction func(c *fiber.Ctx) error

// Handler functions wrapper
func MakeHTTPHandleFunc(myServerFunc serverFuncWithMongoClient, db storage.StorageInterface) fiberHandlerFunction {
	return func(c *fiber.Ctx) error {
		err := myServerFunc(c, db)
		if err != nil {
			response := fiber.Map{
				"error": err,
			}
			return c.JSON(response)
		}
		return nil
	}
}
