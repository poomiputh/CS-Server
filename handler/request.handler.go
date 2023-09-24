package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

type RequestBody struct {
	RoomID       string `json:"room_id"`
	UserRefer     uint   `json:"user_refer"`
	AdminRefer    uint   `json:"admin_refer"`
	Instructor    string `json:"instructor"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Date          string `json:"date"`
	ReDescription string `json:"request_description"`
	Status        string `json:"status"`
}

func (h handler) AddRequest(c *fiber.Ctx) error {
	body := RequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var request models.Request
	request.RoomID = body.RoomID
	request.UserRefer = body.UserRefer
	request.AdminRefer = body.AdminRefer
	request.Instructor = body.Instructor
	request.StartTime = body.StartTime
	request.EndTime = body.EndTime
	request.Date = body.Date
	request.ReDescription = body.ReDescription
	request.Status = body.Status

	if result := h.DB.Create(&request); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&request)
}

func (h handler) GetRequests(c *fiber.Ctx) error {
	var Requests []models.Request

	if result := h.DB.Preload(clause.Associations).Find(&Requests); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Requests)
}

func (h handler) GetRequest(c *fiber.Ctx) error {
	request := c.Params("id")
	var requests models.Request

	if result := h.DB.Preload(clause.Associations).Find(&requests, request); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&requests)
}

func (h handler) UpdateRequest(c *fiber.Ctx) error {
	id := c.Params("id")
	body := RequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var request models.Request

	if result := h.DB.First(&request, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	request.RoomID = body.RoomID
	request.UserRefer = body.UserRefer
	request.AdminRefer = body.AdminRefer
	request.Instructor = body.Instructor
	request.StartTime = body.StartTime
	request.EndTime = body.EndTime
	request.Date = body.Date
	request.ReDescription = body.ReDescription
	request.Status = body.Status

	h.DB.Save(&request)

	return c.Status(fiber.StatusOK).JSON(&request)
}

func (h handler) DeleteRequest(c *fiber.Ctx) error {
	id := c.Params("id")

	var Requests models.Request

	if result := h.DB.First(&Requests, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&Requests)

	return c.SendStatus(fiber.StatusOK)
}
