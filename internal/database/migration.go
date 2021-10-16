package database

import (
	"github.com/f0ntana/go-begin/internal/comment"
	"github.com/jinzhu/gorm"
)

func MigrateDb(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result != nil {
		return result.Error
	}

	return nil
}
