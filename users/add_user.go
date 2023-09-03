package users

import (
	"go-fiber-api-docker/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type AddUserBody struct {
	User_ID uint `json:"user_id"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Email string `json:"email"`
	Phone string  `json:"phone"`
	Position string `json:"position"`

}

func (h handler) addUser(c *fiber.Ctx) error {
	body := AddUserBody{}

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