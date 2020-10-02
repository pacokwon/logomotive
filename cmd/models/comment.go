package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ArticleID int16
	Article   Article `gorm:"foreignKey:ArticleID;references:Id"`
	Content   string  `gorm:"varchar(200)"`
	Posted    time.Time
	Author    string `gorm:"varchar(10)"`
}

func (Comment) TableName() string {
	return "comments"
}
