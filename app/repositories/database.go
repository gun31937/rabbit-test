package repositories

import (
	"github.com/jinzhu/gorm"
	"rabbit-test/app/repositories/database"
)

type Database interface {
	CreateURL(request database.CreateShortURLRequest) (*uint, error)
	CountAllURL() (*uint64, error)
}

func InitDatabase(db *gorm.DB) Database {
	return &database.Repository{
		Database: db,
	}
}
