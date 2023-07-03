package main

import (
	"errors"
	"github.com/amirsobhani/melk_back/database"
	"github.com/amirsobhani/melk_back/queue"
	"github.com/amirsobhani/melk_back/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("problem with load .env file: %s", err)
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("problem with connect to database: %s", err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatalf("problem with migrating to database: %s", err)
	}

	queue.Connect()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"data":    nil,
					"message": err.Error(),
				})
			}

			// Return from handler
			return nil
		},
	})

	route.Api(app)

	err := app.Listen("127.0.0.1:8080")
	if err != nil {
		log.Fatalf("problem with run server: %s", err)
	}
}
