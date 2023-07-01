package realStateClientRepository

import (
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/database"
)

func Create(realState *models.RealState) {
	database.DB.Create(realState)
}
