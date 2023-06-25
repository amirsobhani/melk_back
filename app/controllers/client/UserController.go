package client

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/app/requests"
	"github.com/amirsobhani/melk_back/database"
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"math/rand"
	"time"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
		})
	}

	if errors := requests.ValidateStruct(*user); errors != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: errors,
			Status:  fiber.StatusBadRequest,
		})
	}

	timeQuery := time.Now().Add(time.Minute * -4)
	var countOtp int64
	database.DB.Model(&models.Otp{
		Mobile: user.Mobile,
	}).Where("created_at >= ?", timeQuery).Count(&countOtp)

	if countOtp > 4 {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "too many request please later",
			Status:  fiber.StatusTooManyRequests,
		})
	}

	min := 1000
	max := 8999
	token := rand.Intn(rand.Intn(max)) + min

	otp := models.Otp{
		Token:    token,
		Mobile:   user.Mobile,
		TempData: datatypes.NewJSONType(*user),
	}

	database.DB.Create(&otp)

	return infastructure.Output(c, &infastructure.OutputStruct{
		Data: token,
	})
}
