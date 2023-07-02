package infastructure

import (
	"github.com/amirsobhani/melk_back/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
)

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

func OutputList(c *fiber.Ctx, dataModel interface{}, output *OutputStruct) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page-size", 10)
	var count int64
	database.DB.Model(dataModel).Count(&count)

	lastPage := int(math.Floor(float64(count/int64(pageSize))) + 1)

	if page > lastPage {
		page = lastPage
	}

	return c.Status(output.Status).JSON(fiber.Map{
		"data": fiber.Map{
			"items":     output.Data,
			"total":     count,
			"page_size": pageSize,
			"page":      page,
			"last_page": lastPage,
		},
		"message": output.Message,
	})
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page-size", 10)
		offset := pageSize * (page - 1)

		return db.Offset(offset).Limit(pageSize)
	}
}
