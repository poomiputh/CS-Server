package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func UserRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := app.Group("/users")
	routes.Post("/", h.addUser)
	routes.Get("/", h.getUsers)
	routes.Get("/:id", h.getUser)
	routes.Put("/:id", h.updateUser)
	routes.Delete("/:id", h.deleteUser)
}
