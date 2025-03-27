package database

import (
	"log"
	"simplyChoreo_Backend/config"
	"simplyChoreo_Backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config.LoadEnv()
	dsn := config.GetDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	err = DB.AutoMigrate(models.GetModels()...)
	if err != nil {
		log.Fatal("AutoMigrate failed: ", err)
	}
}
