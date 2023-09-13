package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Routes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	app.Post("/users", h.addUser)
	app.Get("/users", h.getUsers)
	app.Post("/admins", h.addAdmin)
	app.Get("/admins", h.getAdmin)
}
