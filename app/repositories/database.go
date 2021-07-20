package repositories

import (
	"github.com/jinzhu/gorm"
	"rabbit-test/app/repositories/database"
)

type Database interface {
	CreateUrl(request database.CreateShortUrlRequest) error
}

func InitDatabase(db *gorm.DB) Database {
	return &database.Repository{
		Database: db,
	}
}
