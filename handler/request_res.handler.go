package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type Request_ResBody struct {
	RequestRefer         uint `json:"id_request"`
	ReservationTimeRefer uint `json:"id_reservationtime"`
}

func (h handler) AddRequest_Res(c *fiber.Ctx) error {
	body := Request_ResBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var request_res models.Request_Res
	request_res.RequestRefer = body.RequestRefer
	request_res.ReservationTimeRefer = body.ReservationTimeRefer

	if result := h.DB.Create(&request_res); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&request_res)
}

func (h handler) GetRequest_Res(c *fiber.Ctx) error {
	var Requests_Res []models.Request_Res

	if result := h.DB.Preload("Data_Request").Preload("Data_ReservationTime").Preload("Data_User").Preload("Data_Admin").Find(&Requests_Res); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&Requests_Res)
}
