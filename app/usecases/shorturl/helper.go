package shorturl

import (
	"context"
	"errors"
	"github.com/AlekSi/pointer"
	"net/url"
	"rabbit-test/app/env"
	"rabbit-test/app/repositories/database"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func generateShortCode(number uint64) string {

	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func mapCreateURLRequest(shortCode string, fullURL string, expiry *int) database.CreateShortURLRequest {

	request := new(database.CreateShortURLRequest)
	request.ShortCode = shortCode
	request.FullURL = fullURL

	if expiry != nil {
		exp := time.Now().Add(time.Duration(*expiry) * time.Minute)
		request.Expiry = pointer.ToTime(exp)
	}

	return *request
}

func validateCreateURLRequest(fullURL string) error {

	_, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return errors.New(errorURLFormat)
	}

	blacklistUrl := env.BlacklistURL

	matched, err := regexp.MatchString(blacklistUrl, fullURL)
	if err != nil {
		return errors.New(errorGeneric)
	}

	if matched {
		return errors.New(errorMatchBlacklist)
	}

	return nil
}

func getCurrentURLID(u *UseCase, ctx context.Context) (*uint64, error) {

	redisCurrentID, err := u.RedisRepo.Get(ctx, currentURLID)
	if err != nil {
		return nil, errors.New(errorGeneric)
	}

	if redisCurrentID != nil {
		i, err := strconv.Atoi(*redisCurrentID)
		if err != nil {
			return nil, errors.New(errorGeneric)
		}
		currentID := uint64(i)
		return pointer.ToUint64(currentID), nil
	}

	countAllURL, err := u.DatabaseRepo.CountAllURL()
	if err != nil {
		return nil, errors.New(errorGeneric)
	}

	return countAllURL, nil

}
