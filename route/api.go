package route

import (
	"github.com/amirsobhani/melk_back/app/controllers/clientRealEstateController"
	"github.com/amirsobhani/melk_back/app/controllers/clientUserController"
	"github.com/amirsobhani/melk_back/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	app.Post("/api/signup", clientUserController.Signup)
	app.Post("/api/otp-validate", clientUserController.OtpValidator)

	app.Use(middlewares.CheckAuth)
	app.Get("/api/check", clientUserController.Check)

	//real state route
	app.Post("/api/real-state", clientRealEstateController.Create)
}
