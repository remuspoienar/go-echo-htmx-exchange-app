package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goro/model"
	"log"
	"os"
)

func New() *gorm.DB {
	dbString := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbString))

	if err != nil {
		log.Fatal("Could not connect to DB")
	}

	db.AutoMigrate(&model.User{}, &model.Asset{}, &model.Network{})

	return db
}
