package database

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	Database *gorm.DB
}
