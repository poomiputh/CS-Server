package users

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.User

	if result := h.DB.Find(&users, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}
