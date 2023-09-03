package users

import (
	"go-fiber-api-docker/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserReqestBody struct {
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (h handler) updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateUserReqestBody{}

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
