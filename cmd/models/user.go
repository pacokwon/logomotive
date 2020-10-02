package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(10);unique;primaryKey" json:"username"`
	Password string `gorm:"type:text" json:"password"`
	Email    string `gorm:"type:text" json:"email"`
}

func (User) TableName() string {
	return "users"
}
