package route

import (
	"github.com/amirsobhani/melk_back/app/controllers/clientUserController"
	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	app.Post("/api/signup", clientUserController.Signup)
}
