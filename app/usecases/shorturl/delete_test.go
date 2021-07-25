package shorturl_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCase_Delete(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	m := initTest(ctrl)

	mockShortCode := "rQ3Pfb"

	t.Run("Happy", func(t *testing.T) {

		initEnv()

		m.DatabaseRepo.EXPECT().
			DeleteURL(mock.MatchedBy(func(shortCode string) bool {
				assert.Equal(t, mockShortCode, shortCode)
				return true
			})).Return(nil)
		m.RedisRepo.EXPECT().Delete(ctx, mockShortCode).Return(nil)

		err := m.UseCase.Delete(ctx, mockShortCode)
		assert.NoError(t, err)
	})

	t.Run("Error - from db repo", func(t *testing.T) {

		initEnv()

		errorFromDB := errors.New("error")
		expectedError := errors.New("something went wrong")

		m.DatabaseRepo.EXPECT().DeleteURL(mockShortCode).Return(errorFromDB)

		err := m.UseCase.Delete(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
	})

	t.Run("Error - when delete data from redis", func(t *testing.T) {

		initEnv()

		mockErrorFromRedis := errors.New("error")
		expectedError := errors.New("something went wrong")

		m.DatabaseRepo.EXPECT().DeleteURL(mockShortCode).Return(nil)
		m.RedisRepo.EXPECT().Delete(ctx, mockShortCode).Return(mockErrorFromRedis)

		err := m.UseCase.Delete(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
	})

}
