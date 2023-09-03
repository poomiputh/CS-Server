package reservation

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) getReservations(c *fiber.Ctx) error {
	var reservations []models.Reservation

	if result := h.DB.Find(&reservations); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservations)
}
