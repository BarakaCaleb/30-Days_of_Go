package models

import (
	"gorm.io/gorm"
)

type Song struct {
	ID          uint   `gorm:"primaryKey"`
	Group       string `gorm:"not null"`
	Title       string `gorm:"not null"`
	ReleaseDate string
	Lyrics      string `gorm:"type:text"`
	Link        string `gorm:"type:text"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Song{})
}
