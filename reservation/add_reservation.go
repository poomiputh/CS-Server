package reservation

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type AddReservationBody struct {
	Id          int    `json:"id_room" gorm:"primarykey"`
	Rname       string `json:"roomname"`
	Instructor  string `json:"instructor"`
	Phone       string `json:"phone"`
	Description string `json:"Description"`
	Date        string `json:"date"`
	TimeS       string `json:"time-start"`
	TimeE       string `json:"time-end"`
	Status      string `json:"status"`
}

func (h handler) AddReservation(c *fiber.Ctx) error {

	body := AddReservationBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var reservation models.Reservation
	reservation.Rname = body.Rname
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
