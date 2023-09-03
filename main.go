package main

import (
	"go-fiber-api-docker/config"
	"go-fiber-api-docker/db"
	"go-fiber-api-docker/reservation"
	"go-fiber-api-docker/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	app.Use(cors.New())

	users.UserRoutes(app, h)
	reservation.ReservationRoutes(app, h)
	app.Listen(c.Port)

}
