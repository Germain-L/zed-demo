package storage

import (
	"zed-demo/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
