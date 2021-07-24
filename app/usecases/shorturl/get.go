package shorturl

import (
	"context"
	"errors"
	"github.com/AlekSi/pointer"
	"rabbit-test/app/env"
	"rabbit-test/app/repositories/database"
	"time"
)

func (u *UseCase) Get(ctx context.Context, shortCode string) (*string, error) {

	url, err := getURL(u, ctx, shortCode)
	if err != nil {
		return nil, err
	}

	if url.DeletedAt != nil || (url.Expiry != nil && time.Now().After(*url.Expiry)) {
		return nil, errors.New(ErrorURLExpired)
	}

	hits := url.Hits + 1
	err = u.DatabaseRepo.UpdateURL(url.ID, database.UpdateShortURLRequest{Hits: hits})
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	if hits >= env.PopularURLHits {
		url.Hits = hits
		err = setPopularURL(u, ctx, *url)
		if err != nil {
			return nil, err
		}
	}

	fullURL := pointer.ToString(url.FullUrl)

	return fullURL, nil
}
