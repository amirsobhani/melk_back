package realStateClientRepository

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/database"
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/gofiber/fiber/v2"
)

var realStateModel models.RealState

func Create(realState *models.RealState) {
	database.DB.Create(realState)
}

func Records(c *fiber.Ctx) models.RealState {
	database.DB.Model(realStateModel).Preload("City").Preload("User").
		Scopes(infastructure.Paginate(c)).
		Find(&realStateModel)

	return realStateModel
}
