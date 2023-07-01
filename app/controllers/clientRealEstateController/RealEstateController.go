package clientRealEstateController

import (
	"github.com/amirsobhani/melk_back/app/Repository/provinceAndCityRepository"
	"github.com/amirsobhani/melk_back/app/Repository/realStateClientRepository"
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/app/requests"
	"github.com/amirsobhani/melk_back/infastructure"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var realState = new(models.RealState)

	if err := c.BodyParser(&realState); err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err.Error(),
			Status:  fiber.StatusInternalServerError,
		})
	}

	realState.UserId = c.Locals("user_id").(int64)

	if err := requests.ValidateStruct(*realState); err != nil {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: err,
			Status:  fiber.StatusBadRequest,
		})
	}

	if check, _ := provinceAndCityRepository.CheckBeCityExists(realState.CityId); !check {
		return infastructure.Output(c, &infastructure.OutputStruct{
			Message: "city id not valid",
			Status:  fiber.StatusBadRequest,
		})
	}

	realStateClientRepository.Create(realState)

	return infastructure.Output(c, &infastructure.OutputStruct{
		Message: "create real state successfully",
		Status:  fiber.StatusCreated,
	})
}
