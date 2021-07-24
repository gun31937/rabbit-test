package usecases

import (
	"context"
	"rabbit-test/app/repositories"
	"rabbit-test/app/usecases/shorturl"
)

type ShortURL interface {
	Create(ctx context.Context, fullUrl string, expiry *int) (*shorturl.CreateShortURLResponse, error)
	Get(ctx context.Context, shortCode string) (*string, error)
	Delete(ctx context.Context, shortCode string) error
}

func InitShortURL(databaseRepo repositories.Database, redisRepo repositories.Redis) ShortURL {
	return &shorturl.UseCase{
		DatabaseRepo: databaseRepo,
		RedisRepo:    redisRepo,
	}
}
