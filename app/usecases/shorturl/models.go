package shorturl

import "time"

type CreateShortURLResponse struct {
	ShortURL    string
	ExpiredTime *time.Time
}
