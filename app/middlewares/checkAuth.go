package middlewares

import (
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	userId, err := infastructure.VerifyJWT(c)

	// if token not valid return unauthorized status
	if err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusUnauthorized,
		})
	}

	// set user id in fiber ctx
	c.Locals("user_id", userId)

	return c.Next()
}
