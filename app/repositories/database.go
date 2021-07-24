package repositories

import (
	"github.com/jinzhu/gorm"
	"rabbit-test/app/repositories/database"
)

type Database interface {
	CreateURL(request database.CreateShortURLRequest) (*uint, error)
	CountAllURL() (*uint64, error)
	GetURL(shortCode string) (*database.URL, error)
	ListURL(filter database.ListURLFilterRequest) ([]database.URL, error)
	UpdateURL(id uint, request database.UpdateShortURLRequest) error
	DeleteURL(shortCode string) error
}

func InitDatabase(db *gorm.DB) Database {
	return &database.Repository{
		Database: db,
	}
}
