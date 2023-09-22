package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type ReservationTimeBody struct {
	RoomRefer   uint   `json:"id_room"`
	CourseRefer uint   `json:"id_course"`
	StartTime   string `json:"starttime"`
	EndTime     string `json:"endtime"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Type        string `json:"type"`
}

func (h handler) AddReservationTime(c *fiber.Ctx) error {
	body := ReservationTimeBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var resertime models.ReservationTime
	resertime.RoomRefer = body.RoomRefer
	resertime.CourseRefer = body.CourseRefer
	resertime.StartTime = body.StartTime
	resertime.EndTime = body.EndTime
	resertime.EndTime = body.EndTime
	resertime.StartDate = body.StartDate
	resertime.EndDate = body.EndDate
	resertime.Type = body.Type

	if result := h.DB.Create(&resertime); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&resertime)
}

func (h handler) GetReservationTimes(c *fiber.Ctx) error {
	var ReservationTimes []models.ReservationTime

	if result := h.DB.Preload("Data_Room").Find(&ReservationTimes); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&ReservationTimes)
}

func (h handler) GetReservationTime(c *fiber.Ctx) error {
	reservationtime := c.Params("id")
	var reservationtimes models.User

	if result := h.DB.Preload("Data_Room").Find(&reservationtimes, reservationtime); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&reservationtimes)
}

func (h handler) UpdateReservationTime(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ReservationTimeBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var resertime models.ReservationTime
	resertime.RoomRefer = body.RoomRefer
	resertime.StartTime = body.StartTime
	resertime.EndTime = body.EndTime
	resertime.EndTime = body.EndTime
	resertime.StartDate = body.StartDate
	resertime.EndDate = body.EndDate
	resertime.Type = body.Type
	resertime.CourseRefer = body.CourseRefer

	if result := h.DB.First(&resertime, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&resertime)
	return c.Status(fiber.StatusOK).JSON(&resertime)
}

func (h handler) DeleteReservationTime(c *fiber.Ctx) error {
	id := c.Params("id")

	var ReservationTimes models.ReservationTime

	if result := h.DB.First(&ReservationTimes, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&id)

	return c.SendStatus(fiber.StatusOK)
}
