package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type UserBody struct {
	User_ID uint   `json:"user_id"`
	Email   string `json:"email" `
	Fname   string `json:"firstname"`
	Lname   string `json:"lastname"`
	Phone   string `json:"phone"`
	Role    string `json:"role"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := UserBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	user.User_ID = body.User_ID
	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone
	user.Role = body.Role

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) GetUser(c *fiber.Ctx) error {
	user := c.Params("id")
	var users models.User

	if result := h.DB.Find(&users, user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}


