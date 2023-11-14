package db

import (
	"fmt"
	"go-fiber-api-docker/models"
	"log"
	"go-fiber-api-docker/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ReservationTime{})
	db.AutoMigrate(&models.Room{})

	return db
}
