package handler

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

type UserBody struct {
	CollegeID  uint   `json:"college_id"`
	Email    string `json:"email" `
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := UserBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	user.CollegeID = body.CollegeID
	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone
	user.Role = body.Role
	user.Password = body.Password

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

func (h handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Preload(clause.Associations).Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) GetUser(c *fiber.Ctx) error {
	user := c.Params("id")
	var users models.User

	if result := h.DB.Preload(clause.Associations).Find(&users, user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&user)

	return c.SendStatus(fiber.StatusOK)
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UserBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user models.User
	user.CollegeID = body.CollegeID
	user.Fname = body.Fname
	user.Lname = body.Lname
	user.Email = body.Email
	user.Phone = body.Phone
	user.Role = body.Role
	user.Password = body.Password

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(&user)
}
