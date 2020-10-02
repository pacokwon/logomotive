package store

import (
	"github.com/pacokwon/logomotive/cmd/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
}
