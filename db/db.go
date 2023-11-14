package db

import (
	"fmt"
	"go-fiber-api-docker/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gofiber/storage/postgres/v3"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://postgres:G2gFF36dced5b-EcAG-1DGGF-aGe2aAF@roundhouse.proxy.rlwy.net:26379/railway")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ReservationTime{})
	db.AutoMigrate(&models.Room{})

	return db
}
