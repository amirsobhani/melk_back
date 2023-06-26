package clientUserController

import (
	"github.com/amirsobhani/melk_back/app/Repository/otpClientRepository"
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/app/requests"
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/gofiber/fiber/v2"
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

	countOtp := otpClientRepository.CountValidation(user.Mobile)

	if countOtp > 4 {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "too many request please later",
			Status:  fiber.StatusTooManyRequests,
		})
	}

	token := otpClientRepository.GenerateOtp(user)

	return infastructure.Output(c, &infastructure.OutputStruct{
		Data: token,
	})
}
