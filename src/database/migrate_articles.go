package database

import (
	"app/src/model"
	"gorm.io/gorm"
)

func autoMigrateCMS(db *gorm.DB) {
	_ = db.AutoMigrate(&model.Article{})
}
