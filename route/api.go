package route

import (
	"github.com/amirsobhani/melk_back/app/controllers/clientRealEstateController"
	"github.com/amirsobhani/melk_back/app/controllers/clientUserController"
	"github.com/amirsobhani/melk_back/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("/api/")

	api.Post("signup", clientUserController.Signup)
	api.Post("signin", clientUserController.SignIn)
	api.Post("otp-validate", clientUserController.OtpValidator)

	api.Use(middlewares.CheckAuth)
	api.Get("check", clientUserController.Check)

	//real state route
	api.Post("real-state", clientRealEstateController.Create)
}
