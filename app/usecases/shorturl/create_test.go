package shorturl_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rabbit-test/app/env"
	"rabbit-test/app/repositories/database"
	"strconv"
	"testing"
	"time"
)

func TestUseCase_Create(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	m := initTest(ctrl)

	mockShortCode := "rQ3Pfb"
	mockFullURL := "https://www.facebook.com/"
	mockExpiry := pointer.ToInt(50)
	var mockExpiredTime *time.Time

	redisItemKey := "currentURLID"
	mockCurrentID := pointer.ToString("1000000000")
	mockInsertID := pointer.ToUint(1000000000)

	t.Run("Happy - with current id from redis", func(t *testing.T) {

		initEnv()
		mockShortURL := fmt.Sprintf("%s%s", env.BaseURL, mockShortCode)

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(mockCurrentID, nil)
		m.DatabaseRepo.EXPECT().
			CreateURL(mock.MatchedBy(func(input database.CreateShortURLRequest) bool {
				assert.Equal(t, mockShortCode, input.ShortCode)
				assert.Equal(t, mockFullURL, input.FullURL)
				mockExpiredTime = input.Expiry
				return true
			})).Return(mockInsertID, nil)
		m.RedisRepo.EXPECT().Set(ctx, redisItemKey, strconv.Itoa(int(*mockInsertID))).Return(nil)

		result, err := m.UseCase.Create(ctx, mockFullURL, mockExpiry)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockShortURL, result.ShortURL)
		assert.Equal(t, *mockExpiredTime, *result.ExpiredTime)
	})

	t.Run("Happy - with current id from db", func(t *testing.T) {

		initEnv()
		mockShortURL := fmt.Sprintf("%s%s", env.BaseURL, mockShortCode)

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(nil, nil)
		m.DatabaseRepo.EXPECT().CountAllURL().Return(pointer.ToUint64(uint64(*mockInsertID)), nil)
		m.DatabaseRepo.EXPECT().
			CreateURL(mock.MatchedBy(func(input database.CreateShortURLRequest) bool {
				assert.Equal(t, mockShortCode, input.ShortCode)
				assert.Equal(t, mockFullURL, input.FullURL)
				mockExpiredTime = input.Expiry
				return true
			})).Return(mockInsertID, nil)
		m.RedisRepo.EXPECT().Set(ctx, redisItemKey, strconv.Itoa(int(*mockInsertID))).Return(nil)

		result, err := m.UseCase.Create(ctx, mockFullURL, mockExpiry)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockShortURL, result.ShortURL)
		assert.Equal(t, *mockExpiredTime, *result.ExpiredTime)
	})

	t.Run("Error - when parse url", func(t *testing.T) {

		initEnv()

		mockBadFullURL := "abc"
		expectedError := errors.New("full url is not in the right format")

		result, err := m.UseCase.Create(ctx, mockBadFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - when matching blacklist", func(t *testing.T) {

		initEnv()

		env.BlacklistURL = "(a"
		expectedError := errors.New("something went wrong while create short url")

		result, err := m.UseCase.Create(ctx, mockFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - when matched blacklist", func(t *testing.T) {

		initEnv()

		env.BlacklistURL = "facebook"
		expectedError := errors.New("full url is matched in blacklist")

		result, err := m.UseCase.Create(ctx, mockFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - get current id from redis", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("something went wrong while create short url")

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(nil, expectedError)
		result, err := m.UseCase.Create(ctx, mockFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - convert current id from redis to int", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("something went wrong while create short url")

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(pointer.ToString("bad"), nil)
		result, err := m.UseCase.Create(ctx, mockFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - count all rows in url db", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("something went wrong while create short url")

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(nil, nil)
		m.DatabaseRepo.EXPECT().CountAllURL().Return(nil, expectedError)
		result, err := m.UseCase.Create(ctx, mockFullURL, nil)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - when create short url", func(t *testing.T) {

		initEnv()

		dbRepoError := errors.New("error")
		expectedError := errors.New("something went wrong while create short url")

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(mockCurrentID, nil)
		m.DatabaseRepo.EXPECT().CreateURL(gomock.Any()).Return(nil, dbRepoError)

		result, err := m.UseCase.Create(ctx, mockFullURL, mockExpiry)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

	t.Run("Error - when create short url", func(t *testing.T) {

		initEnv()

		dbRepoError := errors.New("error")
		expectedError := errors.New("something went wrong while create short url")

		m.RedisRepo.EXPECT().Get(ctx, redisItemKey).Return(mockCurrentID, nil)
		m.DatabaseRepo.EXPECT().CreateURL(gomock.Any()).Return(mockInsertID, nil)
		m.RedisRepo.EXPECT().Set(ctx, redisItemKey, strconv.Itoa(int(*mockInsertID))).Return(dbRepoError)

		result, err := m.UseCase.Create(ctx, mockFullURL, mockExpiry)
		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

}
