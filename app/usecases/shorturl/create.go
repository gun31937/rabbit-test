package shorturl

import (
	"context"
	"errors"
	"fmt"
	"rabbit-test/app/env"
	"strconv"
)

func (u *UseCase) Create(ctx context.Context, fullURL string, expiry *int) (*CreateShortURLResponse, error) {

	err := validateCreateURLRequest(fullURL)
	if err != nil {
		return nil, err
	}

	currentID, err := getCurrentURLID(u, ctx)
	if err != nil {
		return nil, err
	}

	shortCode := generateShortCode(*currentID + 1)
	createURLRequest := mapCreateURLRequest(shortCode, fullURL, expiry)

	insertID, err := u.DatabaseRepo.CreateURL(createURLRequest)

	if err != nil {
		return nil, errors.New(errorGeneric)
	}

	err = u.RedisRepo.Set(ctx, currentURLID, strconv.Itoa(int(*insertID)))
	if err != nil {
		return nil, errors.New(errorGeneric)
	}

	response := CreateShortURLResponse{
		ShortURL:    fmt.Sprintf("%s%s", env.BaseURL, shortCode),
		ExpiredTime: createURLRequest.Expiry,
	}

	return &response, nil
}
