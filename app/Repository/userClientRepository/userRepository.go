package userClientRepository

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/database"
)

func CheckExists(mobile string) error {
	var user = models.User{
		Mobile: mobile,
	}
	var exists bool

	return database.DB.Model(user).First(&exists).Error
}

func Create(data models.User) models.User {
	database.DB.Create(&data)

	return data
}

func FindByMobile(mobile string) models.User {
	var user models.User
	database.DB.Model(&models.User{}).Where("mobile = ?", mobile).First(&user)
	return user
}
