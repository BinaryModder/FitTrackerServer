package db

import (
	"github.com/binarymodder/fitness-tracker-server/internal/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDataBase() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Russia/Krasnodar", PreferSimpleProtocol: true}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Todo{})

	return db, nil
}
