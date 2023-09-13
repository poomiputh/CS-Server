package handler

import (
	"fmt"
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type UserBody struct {
	User_ID uint   `json:"user_id"`
	Fname   string `json:"firstname"`
	Lname   string `json:"lastname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func (h handler) addUser(c *fiber.Ctx) error {
	body := UserBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fmt.Println("Adding User")

	var user models.Data_User
	user.User_ID = body.User_ID
	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h handler) getUsers(c *fiber.Ctx) error {
	var users []models.Data_User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

type AdminBody struct {
	UserRefer uint `json:"id"`
}

func (h handler) addAdmin(c *fiber.Ctx) error {
	body := AdminBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fmt.Println("Adding User")

	var admin models.Admin
	admin.UserRefer = body.UserRefer

	if result := h.DB.Create(&admin); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&admin)
}

func (h handler) getAdmin(c *fiber.Ctx) error {
	var admin []models.Admin

	if result := h.DB.Preload("Data_Admin").Find(&admin); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&admin)
}
