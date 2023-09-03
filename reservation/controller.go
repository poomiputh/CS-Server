package reservation

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func ReservationRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := app.Group("/reservations")
	routes.Post("/", h.addReservation)
	routes.Get("/", h.getReservations)
	routes.Delete("/:id", h.deleteReservation)
}
