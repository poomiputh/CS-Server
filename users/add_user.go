package users

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type AddUSerBody struct {
	ID_user int `json:"id_user"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Email string `json:"email"`
	Phone string  `json:"phone"`
	Role string `json:"role"`

}

func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUSerBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User
	user.ID_user = body.ID_user
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
