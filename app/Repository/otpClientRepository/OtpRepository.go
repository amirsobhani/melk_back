package otpClientRepository

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/database"
	"gorm.io/datatypes"
	"math/rand"
	"time"
)

func CountValidation(Mobile string) int64 {
	timeQuery := time.Now().Add(time.Minute * -4)
	var countOtp int64
	database.DB.Model(&models.Otp{
		Mobile: Mobile,
	}).Where("created_at >= ?", timeQuery).Count(&countOtp)

	return countOtp
}

func GenerateOtp(user *models.User) int {
	min := 1000
	max := 8999
	token := rand.Intn(rand.Intn(max)) + min

	otp := models.Otp{
		Token:    token,
		Mobile:   user.Mobile,
		TempData: datatypes.NewJSONType(*user),
	}

	database.DB.Create(&otp)

	return token
}

func CheckValidOtp(otp int, mobile string) models.Otp {
	var otpModel models.Otp

	timeQuery := time.Now().Add(time.Minute * -2)

	database.DB.Model(&otpModel).
		Where("token = ? AND mobile = ? AND created_at >= ?", otp, mobile, timeQuery).First(&otpModel)

	return otpModel
}

func RemoveAllMobileOtp(mobile string) {
	var res interface{}
	database.DB.Model(&models.Otp{}).Where("mobile = ?", mobile).Delete(&res)
}
