package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(10) primaryKey"`
	Password string `gorm:"type:text"`
	Email    string `gorm:"type:citext"`
}

func (User) TableName() string {
	return "users"
}
