package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	UserRefer     uint   `json:"id_user"`
	AdminRefer    uint   `json:"id_admin"`
	StartTime     string `json:"starttime"`
	EndTime       string `json:"endtime"`
	StartDate     string `json:"startdate"`
	EndDate       string `json:"enddate"`
	ReDescription string `json:"requestdescription"`
	Status        string `json:"status"`
}

func (h handler) AddRequest(c *fiber.Ctx) error {
	body := RequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var request models.Request
	request.UserRefer = body.UserRefer
	request.AdminRefer = body.AdminRefer
	request.StartTime = body.StartTime
	request.EndTime = body.EndTime
	request.StartDate = body.StartDate
	request.EndDate = body.EndDate
	request.ReDescription = body.ReDescription
	request.Status = body.Status

	if result := h.DB.Create(&request); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&request)
}

func (h handler) GetRequests(c *fiber.Ctx) error {
	var Requests []models.Request

	if result := h.DB.Preload("Data_User").Preload("Data_Admin").Find(&Requests); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Requests)
}

func (h handler) GetRequest(c *fiber.Ctx) error {
	request := c.Params("id")
	var requests models.Request

	if result := h.DB.Preload("Data_User").Preload("Data_Admin").Find(&requests, request); result.Error != nil {
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
	request.UserRefer = body.UserRefer
	request.AdminRefer = body.AdminRefer
	request.StartTime = body.StartTime
	request.EndTime = body.EndTime
	request.StartDate = body.StartDate
	request.EndDate = body.EndDate
	request.ReDescription = body.ReDescription
	request.Status = body.Status

	if result := h.DB.First(&request, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&request)
	return c.Status(fiber.StatusOK).JSON(&request)
}
