package main

import (
	"errors"
	"fmt"
	"github.com/amirsobhani/melk_back/database"
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

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"data":    nil,
					"message": "Internal Server Error",
				})
			}

			// Return from handler
			return nil
		},
	})

	route.Api(app)

	app.Listen("127.0.0.1:8080")
}
