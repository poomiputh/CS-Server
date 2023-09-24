package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type Room struct {
	RoomID string `json:"room_id"`
}

func (h handler) AddRoom(c *fiber.Ctx) error {
	body := Room{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var room models.Room
	room.RoomID = body.RoomID

	if result := h.DB.Create(&room); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&room)
}

func (h handler) GetRooms(c *fiber.Ctx) error {
	var Room []models.Room

	if result := h.DB.Preload("DataReservation").Find(&Room); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Room)
}
