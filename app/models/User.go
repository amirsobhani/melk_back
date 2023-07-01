package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string  `json:"name" validate:"required,min=3,max=150"`
	Family string  `json:"family" validate:"required"`
	Type   int8    `json:"type"`
	Email  *string `json:"email" validate:"email" gorm:"DEFAULT NULL;UNIQUE"`
	Mobile string  `json:"mobile" gorm:"size:15;UNIQUE" validate:"required,len=10"`
}
