package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		return nil
	}

	return db
}
