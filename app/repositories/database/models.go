package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CreateShortURLRequest struct {
	ShortCode string
	FullURL   string
	Expiry    *time.Time
}

type URL struct {
	gorm.Model
	ShortCode string
	FullUrl   string
	Expiry    *time.Time `json:",omitempty"`
}

func (URL) TableName() string {
	return "urls"
}
