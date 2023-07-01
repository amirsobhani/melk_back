package models

import "gorm.io/gorm"

type RealState struct {
	gorm.Model
	Name       string `json:"name" validate:"required,max=100,min=2" gorm:"NOT NULL;"`
	User       User   `validate:"-"`
	UserId     int64  `json:"user_id" validate:"required"`
	City       ProvinceAndCity
	CityId     int64  `json:"city_id" validate:"required"`
	Address    string `json:"address" validate:"required,max=150,min=5"`
	Phone      string `json:"phone" validate:"required"`
	IsActive   bool   `json:"is_active" validate:"boolean" gorm:"DEFAULT TRUE;"`
	IsSelected bool   `json:"is_selected" validate:"boolean" gorm:"DEFAULT FALSE;"`
}
