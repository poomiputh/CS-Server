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
	app.Get("/users/:id", h.getUser)
	app.Put("/users/:id", h.updateUser)
	app.Delete("/users/:id", h.deleteUser)
	app.Post("/reservations", h.addReservation)
	app.Get("/reservations", h.getReservations)
	app.Put("/reservations/:id", h.updateReservation)
	app.Get("/reservations/:id", h.getReservation)
	app.Delete("/reservations/:id", h.deleteReservation)

}
