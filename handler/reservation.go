package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type ReservationBody struct {
	RoomID      string `json:"room_id"`
	UserRefer   uint   `json:"user_id"`
	Instructor  string `json:"instructor"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Date        string `json:"date"`
	TimeS       string `json:"time_start"`
	TimeE       string `json:"time_end"`
	Status      string `json:"status"`
}

func (h handler) addReservation(c *fiber.Ctx) error {

	body := ReservationBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var reservation models.Reservation

	reservation.RoomID = body.RoomID
	reservation.Instructor = body.Instructor
	reservation.UserRefer = body.UserRefer
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

func (h handler) getReservations(c *fiber.Ctx) error {
	var reservations []models.Reservation

	if result := h.DB.Find(&reservations); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservations)
}

func (h handler) getReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation models.Reservation

	if result := h.DB.Find(&reservation, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservation)
}

func (h handler) updateReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ReservationBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var reservation models.Reservation

	if result := h.DB.First(&reservation, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	reservation.RoomID = body.RoomID
	reservation.Instructor = body.Instructor
	reservation.UserRefer = body.UserRefer
	reservation.Phone = body.Phone
	reservation.Description = body.Description
	reservation.Date = body.Date
	reservation.TimeS = body.TimeS
	reservation.TimeE = body.TimeE
	reservation.Status = body.Status

	h.DB.Save(&reservation)

	return c.Status(fiber.StatusOK).JSON(&reservation)
}

func (h handler) deleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")

	var reservation models.Reservation

	if result := h.DB.First(&reservation, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&reservation)

	return c.SendStatus(fiber.StatusOK)
}
