package reservation

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type AddReservationBody struct {
	RoomID      string `json:"room_id"`
	Instructor  string `json:"instructor"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Date        string `json:"date"`
	TimeS       string `json:"time_start"`
	TimeE       string `json:"time_end"`
	Status      string `json:"status"`
}

func (h handler) addReservation(c *fiber.Ctx) error {

	body := AddReservationBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var reservation models.Reservation
	
	reservation.RoomID = body.RoomID
	reservation.Instructor = body.Instructor
	reservation.Phone = body.Phone
	reservation.Description = body.Description
	reservation.Date = body.Date
	reservation.TimeS = body.TimeS
	reservation.TimeE = body.TimeE
	reservation.Status = body.Status

	if result := h.DB.Create(&reservation); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&reservation)
}
