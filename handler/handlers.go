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

	api := app.Group("/api")

	login := api.Group("/login")
	login.Post("/", h.Login)

	users := api.Group("/users")
	users.Post("/add", h.AddUser)
	users.Delete("/delete/:id", h.DeleteUser)
	users.Get("/list", h.GetUsers)
	users.Get("/get/:id", h.GetUser)

	rooms := api.Group("/rooms")
	rooms.Post("/add", h.AddRoom)
	rooms.Get("/list", h.GetRooms)

	// https://localhost:3000/api/reservations/add
	// https://localhost:3000/api/reservations/delete/1
	reservations := api.Group("/reservations")
	reservations.Post("/add", h.AddReservation)
	reservations.Delete("/delete_course/:course_id/:course_type", h.DeleteCourseReservations)
	reservations.Delete("/delete/:id", h.DeleteReservation)
	reservations.Get("/get/:id", h.GetReservation)
	reservations.Get("/list", h.GetAllReservations)
	reservations.Get("/list/:type", h.GetAllReservationsByType)
	reservations.Get("get_course/:course_id/:course_type", h.GetCourseReservations)
	reservations.Put("/update/:id", h.UpdateReservation)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

}
