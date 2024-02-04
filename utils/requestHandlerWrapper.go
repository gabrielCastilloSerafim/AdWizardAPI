package utils

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type serverFuncWithMongoClient func(c *fiber.Ctx, mongoClient *mongo.Client) error
type fiberHandlerFunction func(c *fiber.Ctx) error

// Handler functions wrapper
func MakeHTTPHandleFunc(myServerFunc serverFuncWithMongoClient, mongoClient *mongo.Client) fiberHandlerFunction {
	return func(c *fiber.Ctx) error {
		err := myServerFunc(c, mongoClient)
		if err != nil {
			response := fiber.Map{
				"error": err,
			}
			return c.JSON(response)
		}
		return nil
	}
}
