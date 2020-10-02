package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Id         int16  `gorm:"type:serial;primaryKey;unique"`
	Title      string `gorm:"type:varchar(100)"`
	AuthorName string
	Author     User   `gorm:"foreignKey:AuthorName;references:Username"`
	Content    string `gorm:"type:text"`
	Labels     string `gorm:"type:text[]"`
	Posted     time.Time
	Modified   time.Time
}

func (Article) TableName() string {
	return "articles"
}
