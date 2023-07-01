package models

import "gorm.io/gorm"

type ProvinceAndCity struct {
	gorm.Model
	Name       string `json:"name" gorm:"NOT NULL;"`
	Province   *ProvinceAndCity
	ProvinceId *int64 `json:"province_id" gorm:"TYPE:integer REFERENCES ProvinceAndCities"`
}
