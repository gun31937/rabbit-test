package database

import (
	"time"
)

type CreateShortURLRequest struct {
	ShortCode string
	FullURL   string
	Expiry    *time.Time
}

type UpdateShortURLRequest struct {
	Hits int
}

type ListURLFilterRequest struct {
	ShortCode *string
	Keyword   *string
}

type URL struct {
	ID        uint       `gorm:"primary_key"`
	ShortCode string     `gorm:"column:short_code"`
	FullURL   string     `gorm:"column:full_url"`
	Expiry    *time.Time `gorm:"column:expiry" json:",omitempty"`
	Hits      int        `gorm:"column:hits"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (URL) TableName() string {
	return "urls"
}
