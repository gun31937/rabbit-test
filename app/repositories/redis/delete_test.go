package redis_test

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"testing"
)

func TestRepository_Delete(t *testing.T) {

	var ctx = context.TODO()

	mockItemKey := "someKey"

	t.Run("Happy", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()

		defer func() {
			_ = rdb.Close()
		}()

		repository := repositories.InitRedis(rdb)

		mock.ExpectDel(mockItemKey).SetVal(1)
		err := repository.Delete(ctx, mockItemKey)
		assert.NoError(t, err)
	})

	t.Run("Fail - unexpected error", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()
		defer func() {
			_ = rdb.Close()
		}()

		expectedError := errors.New("error")
		repository := repositories.InitRedis(rdb)
		mock.ExpectDel(mockItemKey).SetErr(expectedError)
		err := repository.Delete(ctx, mockItemKey)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
	})

}
