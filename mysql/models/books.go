package models

import (
	"gorm.io/gorm"
)

type Books struct {
	Id        uint64 `gorm:"primaryKey;autoIncreament" json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
