package database

import (
	"github.com/amirsobhani/melk_back/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USERNAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(&models.User{}, &models.Otp{})

	if err != nil {
		return err
	}

	return nil
}
