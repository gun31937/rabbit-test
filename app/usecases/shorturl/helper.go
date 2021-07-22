package shorturl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func mapCreateShortURLResponse(shortCode string, expiredTime *time.Time) *CreateShortURLResponse {

	response := new(CreateShortURLResponse)
	response.ShortURL = fmt.Sprintf("%s%s", env.BaseURL, shortCode)

	if expiredTime != nil {
		location, _ := time.LoadLocation("Asia/Bangkok")
		response.ExpiredTime = pointer.ToString(expiredTime.In(location).Format(TimeFormat))
	}

	return response
}

func validateCreateURLRequest(fullURL string) error {

	_, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return errors.New(ErrorURLFormat)
	}

	blacklistUrl := env.BlacklistURL

	matched, err := regexp.MatchString(blacklistUrl, fullURL)
	if err != nil {
		return errors.New(ErrorGeneric)
	}

	if matched {
		return errors.New(ErrorMatchBlacklist)
	}

	return nil
}

func getCurrentURLID(u *UseCase, ctx context.Context) (*uint64, error) {

	redisCurrentID, err := u.RedisRepo.Get(ctx, CurrentURLID)
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	if redisCurrentID != nil {
		i, err := strconv.Atoi(*redisCurrentID)
		if err != nil {
			return nil, errors.New(ErrorGeneric)
		}
		currentID := uint64(i)
		return pointer.ToUint64(currentID), nil
	}

	countAllURL, err := u.DatabaseRepo.CountAllURL()
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	return countAllURL, nil

}

func getURL(u *UseCase, ctx context.Context, shortCode string) (*database.URL, error) {

	getPopularURL, err := u.RedisRepo.Get(ctx, shortCode)
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	if getPopularURL != nil {
		response := new(database.URL)
		err = json.Unmarshal([]byte(*getPopularURL), response)
		if err != nil {
			return nil, errors.New(ErrorGeneric)
		}
		return response, nil
	}

	response, err := u.DatabaseRepo.GetURL(shortCode)
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	if response == nil {
		return nil, errors.New(ErrorRecordNotFound)
	}

	return response, nil

}

func setPopularURL(u *UseCase, ctx context.Context, popularURL database.URL) error {

	popularURLJSON, _ := json.Marshal(popularURL)

	err := u.RedisRepo.Set(ctx, popularURL.ShortCode, string(popularURLJSON))
	if err != nil {
		return errors.New(ErrorGeneric)
	}

	return nil

}
