package users

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) getUsers(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}
