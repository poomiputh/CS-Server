package reservation

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) getReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation models.Reservation

	if result := h.DB.Find(&reservation, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservation)
}
