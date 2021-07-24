package shorturl

import "time"

type CreateShortURLResponse struct {
	ShortURL    string
	ExpiredTime *string
}

type URL struct {
	ID        uint
	ShortCode string
	FullUrl   string
	Expiry    *time.Time
	Hits      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ListURLRequest struct {
	ShortCode *string
	Keyword   *string
}
