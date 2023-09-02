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
	routes.Post("/", h.AddUser)
	routes.Get("/", h.GetUsers)
	routes.Put("/:id", h.UpdateUser)
	routes.Delete("/:id", h.DeleteUser)
}
