package shorturl

import (
	"context"
	"errors"
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
		return nil, errors.New(ErrorGeneric)
	}

	err = u.RedisRepo.Set(ctx, CurrentURLID, strconv.Itoa(int(*insertID)))
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	response := mapCreateShortURLResponse(shortCode, createURLRequest.Expiry)

	return response, nil
}
