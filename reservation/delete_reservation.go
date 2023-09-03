package reservation

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) deleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")

	var reservation models.Reservation

	if result := h.DB.First(&reservation, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&reservation)

	return c.SendStatus(fiber.StatusOK)
}
