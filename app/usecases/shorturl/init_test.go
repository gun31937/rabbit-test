package shorturl_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"rabbit-test/app/env"
	"rabbit-test/app/mocks/repositories"
	"rabbit-test/app/usecases"
)

type mocks struct {
	DatabaseRepo *mock_repositories.MockDatabase
	RedisRepo    *mock_repositories.MockRedis
	UseCase      usecases.ShortURL
	Context      context.Context
}

func initTest(ctrl *gomock.Controller) *mocks {
	m := mocks{
		DatabaseRepo: mock_repositories.NewMockDatabase(ctrl),
		RedisRepo:    mock_repositories.NewMockRedis(ctrl),
		Context:      context.TODO(),
	}
	m.UseCase = usecases.InitShortURL(m.DatabaseRepo, m.RedisRepo)

	return &m
}

func initEnv() {
	env.BlacklistURL = "google.*"
	env.BaseURL = "http://localhost:8080/"
	env.PopularURLHits = 10
}
