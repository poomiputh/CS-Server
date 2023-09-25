package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
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

	app.Post("/login", h.Login)

	app.Get("/users", h.GetUsers)
	app.Get("/users/:id", h.GetUser)

	app.Get("/reservationtimes/:id", h.GetReservationTime)
	app.Get("/reservationtimes", h.GetReservationTimes)
	
	app.Get("/rooms", h.GetRooms)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Post("/users", h.AddUser)
	app.Delete("/users/:id", h.DeleteUser)

	app.Post("/reservationtimes", h.AddReservationTime)
	app.Put("/reservationtimes/:id", h.UpdateReservationTime)
	app.Delete("/reservationtimes/:id", h.DeleteReservationTime)

	app.Post("/rooms", h.AddRoom)

}
