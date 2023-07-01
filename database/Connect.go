package database

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() error {
	driver := os.Getenv("DB_DRIVER")

	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}

	var err error

	switch utils.ToLower(driver) {
	case "mysql":
		DB, err = MysqlDriver(&config)
	case "postgres":
		DB, err = PostgresqlDriver(&config)
	default:
		DB, err = PostgresqlDriver(&config)
	}

	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(&models.User{}, &models.Otp{}, &models.ProvinceAndCity{}, &models.RealState{})

	if err != nil {
		return err
	}

	return nil
}
