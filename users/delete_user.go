package users

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteUser(c *fiber.Ctx) error{
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil{
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&user)

	return c.SendStatus(fiber.StatusOK)
}