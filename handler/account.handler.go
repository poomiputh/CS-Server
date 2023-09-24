package handler

import (
	"go-fiber-api-docker/models"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

func (h handler) Login(c *fiber.Ctx) error {

	login_user := c.FormValue("user")
	login_pass := c.FormValue("pass")

	// เช็คว่ามีการส่งค่ามาครบไหม
	if login_user == "" || login_pass == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing login credentials")
	}

	var users models.User

	if result := h.DB.Where("email = ?", login_user).First(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// Throws Unauthorized error
	if users.Password != login_pass {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":         users.ID,
		"college_id": users.CollegeID,
		"first_name": users.Fname,
		"last_name":  users.Lname,
		"email":      users.Email,
		"phone":      users.Phone,
		"role":       users.Role,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
