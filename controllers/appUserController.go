package controllers

import (
	"github.com/gabrielCastilloSerafim/AdWizardAPI/storage"
	"github.com/gofiber/fiber/v2"
)

func HandleGetAllAppUsers(c *fiber.Ctx, db storage.StorageInterface) error {
	appUsers, err := db.GetAllAppUsers()
	if err != nil {
		return err
	}
	return c.JSON(appUsers)
}

func HandleDeleteAllAppUsers(c *fiber.Ctx, db storage.StorageInterface) error {
	err := db.DeleteAllAppUsers()
	if err != nil {
		return err
	}
	return c.SendString("All AppUsers Deleted")
}
