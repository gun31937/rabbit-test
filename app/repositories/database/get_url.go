package database

import (
	"github.com/jinzhu/gorm"
)

func (repository *Repository) GetURL(shortCode string) (*URL, error) {

	var url URL

	// Will found deleted item for response status
	result := repository.Database.Unscoped().Where(&URL{ShortCode: shortCode}).First(&url)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}
