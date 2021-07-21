package redis_test

import (
	"context"
	"errors"
	"github.com/AlekSi/pointer"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/env"
	"rabbit-test/app/repositories"
	"testing"
)

func TestRepository_Get(t *testing.T) {

	var ctx = context.TODO()
	env.RedisItemTTL = 10

	mockRedisResponse := pointer.ToString("2")
	mockItemKey := "someKey"

	t.Run("Happy", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()

		defer func() {
			_ = rdb.Close()
		}()

		repository := repositories.InitRedis(rdb)

		mock.ExpectGet(mockItemKey).SetVal(*mockRedisResponse)
		result, err := repository.Get(ctx, mockItemKey)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockRedisResponse, result)
	})

	t.Run("Happy - with no value", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()
		defer func() {
			_ = rdb.Close()
		}()

		repository := repositories.InitRedis(rdb)
		mock.ExpectGet(mockItemKey).RedisNil()
		result, err := repository.Get(ctx, mockItemKey)

		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("Fail - Error while get data", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()
		defer func() {
			_ = rdb.Close()
		}()

		expectedError := errors.New("error")
		repository := repositories.InitRedis(rdb)
		mock.ExpectGet(mockItemKey).SetErr(expectedError)
		result, err := repository.Get(ctx, mockItemKey)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

}
