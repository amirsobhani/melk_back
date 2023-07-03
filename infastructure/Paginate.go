package infastructure

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Paginate use gorm scope for pagination -> .Scope(Paginate(c)).
func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page-size", 10)
		offset := pageSize * (page - 1)

		return db.Offset(offset).Limit(pageSize)
	}
}
