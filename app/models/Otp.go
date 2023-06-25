package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	UserId   *int64                   `json:"user_id"`
	User     User                     `json:"-"`
	Token    int                      `json:"token"`
	Mobile   string                   `json:"mobile" gorm:"size:15" validate:"len=10"`
	TempData datatypes.JSONType[User] `json:"temp_data" gorm:"DEFAULT NULL"`
}
