package shorturl

import (
	"rabbit-test/app/repositories"
)

type UseCase struct {
	DatabaseRepo repositories.Database
	RedisRepo    repositories.Redis
}
