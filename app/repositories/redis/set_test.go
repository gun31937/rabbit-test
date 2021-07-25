package redis_test

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/env"
	"rabbit-test/app/repositories"
	"testing"
	"time"
)

func TestRepository_Set(t *testing.T) {

	var ctx = context.TODO()
	env.RedisItemTTL = 10
	itemTTL := time.Duration(env.RedisItemTTL) * time.Minute

	mockItemKey := "someKey"
	mockItemValue := "someValue"

	t.Run("Happy", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()

		defer func() {
			_ = rdb.Close()
		}()

		repository := repositories.InitRedis(rdb)

		mock.ExpectSet(mockItemKey, mockItemValue, itemTTL).SetVal("")
		err := repository.Set(ctx, mockItemKey, mockItemValue)
		assert.NoError(t, err)

	})

	t.Run("Fail - Error while store data", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()

		defer func() {
			_ = rdb.Close()
		}()

		expectedError := errors.New("error")
		repository := repositories.InitRedis(rdb)

		mock.ExpectSet(mockItemKey, mockItemValue, itemTTL).SetErr(expectedError)
		err := repository.Set(ctx, mockItemKey, mockItemValue)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)

	})

}
