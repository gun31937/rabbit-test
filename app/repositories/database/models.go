package database

import (
	"time"
)

type CreateShortUrlRequest struct {
	ShortCode string
	FullUrl   string
	Expiry    *time.Time
}
