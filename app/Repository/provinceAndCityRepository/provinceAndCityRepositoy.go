package provinceAndCityRepository

import (
	"errors"
	"github.com/amirsobhani/melk_back/app/models"
	"github.com/amirsobhani/melk_back/database"
)

func CheckBeCityExists(CityId int64) (bool, error) {
	var provinceAndCity = models.ProvinceAndCity{}

	database.DB.Where("province_id is not null and id = ?", CityId).Find(&provinceAndCity)

	if provinceAndCity.ID == 0 {
		return false, errors.New("record not exists")
	}

	if provinceAndCity.ProvinceId == nil {
		return false, nil
	}

	return true, nil
}
