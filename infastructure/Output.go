package infastructure

import "github.com/gofiber/fiber/v2"

type OutputStruct struct {
	Data    interface{}
	Message interface{}
	Status  int
}

func Output(c *fiber.Ctx, output *OutputStruct) error {
	return c.Status(output.Status).JSON(fiber.Map{
		"data":    output.Data,
		"message": output.Message,
	})
}
