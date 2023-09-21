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
	app.Post("/users", h.AddUser)
	app.Get("/users", h.GetUsers)
	app.Get("/users/:id", h.GetUser)
	app.Post("/requests", h.AddRequest)
	app.Get("/requests", h.GetRequests)
	app.Get("/requests/:id", h.GetRequest)
	app.Put("/requests/:id", h.UpdateRequest)
	app.Post("/courses", h.AddCourse)
	app.Get("/courses", h.GetCourses)
	app.Get("/courses/:id", h.GetCourse)
	app.Put("/courses/:id", h.UpdateCourse)
	app.Post("/requests_res", h.AddRequest_Res)
	app.Get("/requests_res", h.GetRequest_Res)
	app.Post("/reservationtimes", h.AddReservationTime)
	app.Get("/reservationtimes", h.GetReservationTimes)
	app.Put("/reservationtimes/:id", h.GetReservationTime)

}
