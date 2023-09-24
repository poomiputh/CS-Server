package handler

import (
	// jwtware "github.com/gofiber/contrib/jwt"
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
	app.Get("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)

	app.Post("/requests", h.AddRequest)
	app.Get("/requests", h.GetRequests)
	app.Get("/requests/:id", h.GetRequest)
	app.Put("/requests/:id", h.UpdateRequest)
	app.Delete("/requests/:id", h.DeleteRequest)

	app.Post("/courses", h.AddCourse)
	app.Get("/courses", h.GetCourses)
	app.Get("/courses/:id", h.GetCourse)
	app.Put("/courses/:id", h.UpdateCourse)
	app.Delete("/courses/:id", h.DeleteCourse)

	app.Post("/reservationtimes", h.AddReservationTime)
	app.Get("/reservationtimes", h.GetReservationTimes)
	app.Get("/reservationtimes/:id", h.GetReservationTime)
	app.Put("/reservationtimes/:id", h.UpdateReservationTime)
	app.Delete("/reservationtimes/:id", h.DeleteReservationTime)

	app.Post("/requests_res", h.AddRequest_Res)
	app.Get("/requests_res", h.GetRequest_Res)

	app.Post("/rooms", h.AddRoom)
	app.Get("/rooms", h.GetRooms)

	// app.Post("/login", h.Login)

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))

	// app.Post("/users", h.AddUser)
	// app.Get("/users", h.GetUsers)
	// app.Get("/users/:id", h.GetUser)
	// app.Delete("/users/:id", h.DeleteUser)

	// app.Post("/requests", h.AddRequest)
	// app.Get("/requests", h.GetRequests)
	// app.Get("/requests/:id", h.GetRequest)
	// app.Put("/requests/:id", h.UpdateRequest)
	// app.Delete("/requests/:id", h.DeleteRequest)

	// app.Post("/courses", h.AddCourse)
	// app.Get("/courses", h.GetCourses)
	// app.Get("/courses/:id", h.GetCourse)
	// app.Put("/courses/:id", h.UpdateCourse)
	// app.Delete("/courses/:id", h.DeleteCourse)

	// app.Post("/reservationtimes", h.AddReservationTime)
	// app.Get("/reservationtimes", h.GetReservationTimes)
	// app.Get("/reservationtimes/:id", h.GetReservationTime)
	// app.Put("/reservationtimes/:id", h.UpdateReservationTime)
	// app.Delete("/reservationtimes/:id", h.DeleteReservationTime)

	// app.Post("/requests_res", h.AddRequest_Res)
	// app.Get("/requests_res", h.GetRequest_Res)

	// app.Post("/rooms", h.AddRoom)
	// app.Get("/rooms", h.GetRooms)

}
