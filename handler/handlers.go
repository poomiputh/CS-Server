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

	// http://localhost:3000/api/reservations/add
	// http://localhost:3000/api/reservations/delete/1
	reservations := api.Group("/reservations")
	reservations.Post("/add", h.AddReservation)                                               // สำหรับเพิ่ม Reservation ทั้งแบบเดี่ยวและเป็นชุด
	reservations.Delete("/delete_course/:course_id/:course_type", h.DeleteCourseReservations) // สำหรับลบ Course ทั้งแบบเดี่ยวและเป็นชุด
	reservations.Delete("/delete/:id", h.DeleteReservation)                                   // สำหรับลบ Reservation แบบเดี่ยว
	reservations.Get("/get/:id", h.GetReservation)                                            // สำหรับดึงค่า Reservation แบบเดี่ยว
	// reservations.Get("/list", h.GetAllReservations)
	reservations.Get("/list/:type?/:status?", h.GetAllReservationsByFilter)         // สำหรับดึงค่า Reservation ทั้งหมด หรือทั้งหมดที่มี Type และ Status ที่ต้องการ
	reservations.Get("get_course/:course_id/:course_type", h.GetCourseReservations) // สำหรับดึงค่า Course ทั้งชุด
	reservations.Put("/update/:id", h.UpdateReservation)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

}
