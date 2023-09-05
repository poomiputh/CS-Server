package handler

import (
	"fmt"
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type UserBody struct {
	User_ID  uint   `json:"user_id"`
	Fname    string `json:"firstname"`
	Lname    string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}

func (h handler) addUser(c *fiber.Ctx) error {
	body := UserBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fmt.Println("Adding User")

	var user models.User
	user.User_ID = body.User_ID
	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone
	user.Position = body.Position

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h handler) getUsers(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Preload("Reservations").Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.User

	if result := h.DB.Preload("Reservations").Find(&users, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UserBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone

	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(&user)
}

func (h handler) deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&user)

	return c.SendStatus(fiber.StatusOK)
}
