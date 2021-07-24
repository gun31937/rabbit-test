package shorturl_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rabbit-test/app/repositories/database"
	"testing"
	"time"
)

func TestUseCase_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	m := initTest(ctrl)

	mockShortCode := "rQ3Pfb"
	mockFullURL := "https://www.facebook.com/"

	mockGetURLResult := func() database.URL {
		return database.URL{
			ID:        1,
			ShortCode: mockShortCode,
			FullUrl:   mockFullURL,
			Expiry:    nil,
			Hits:      8,
		}
	}

	t.Run("Happy - get item from db, url is not popular", func(t *testing.T) {

		initEnv()

		getURLResult := mockGetURLResult()
		mockHits := getURLResult.Hits + 1
		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)
		m.DatabaseRepo.EXPECT().
			UpdateURL(getURLResult.ID, mock.MatchedBy(func(input database.UpdateShortURLRequest) bool {
				assert.Equal(t, mockHits, input.Hits)
				return true
			})).Return(nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, getURLResult.FullUrl, *result)
	})

	t.Run("Happy - get item from db, url is popular when update hits", func(t *testing.T) {

		initEnv()

		getURLResult := mockGetURLResult()
		getURLResult.Hits = 9

		mockPopularItem := getURLResult
		mockPopularItem.Hits = getURLResult.Hits + 1
		mockPopularItemJSON, _ := json.Marshal(mockPopularItem)

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)
		m.DatabaseRepo.EXPECT().UpdateURL(getURLResult.ID, gomock.Any()).Return(nil)
		m.RedisRepo.EXPECT().Set(ctx, mockShortCode, string(mockPopularItemJSON)).Return(nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Happy - get item from redis, url is already popular", func(t *testing.T) {

		initEnv()

		mockRedisItem := json.RawMessage(`{"ID":1,"ShortCode":"rQ3Pfb","FullUrl":"https://www.facebook.com/","Hits":10,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}`)
		getURLResult := mockGetURLResult()
		getURLResult.Hits = 10

		mockPopularItem := mockGetURLResult()
		mockPopularItem.Hits = getURLResult.Hits + 1
		mockPopularItemJSON, _ := json.Marshal(mockPopularItem)

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(pointer.ToString(string(mockRedisItem)), nil)
		m.DatabaseRepo.EXPECT().UpdateURL(getURLResult.ID, gomock.Any()).Return(nil)
		m.RedisRepo.EXPECT().Set(ctx, mockShortCode, string(mockPopularItemJSON)).Return(nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Error - url expired", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("url expired")

		getURLResult := mockGetURLResult()
		getURLResult.Expiry = pointer.ToTime(time.Now().Add(-10 * time.Minute))

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - url deleted", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("url expired")

		getURLResult := mockGetURLResult()
		getURLResult.DeletedAt = pointer.ToTime(time.Now())

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - when get data from redis", func(t *testing.T) {

		initEnv()

		mockErrorFromRedis := errors.New("error")
		expectedError := errors.New("something went wrong")

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, mockErrorFromRedis)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - when parse redis item", func(t *testing.T) {

		initEnv()

		mockRedisItem := json.RawMessage(`{bad}`)

		expectedError := errors.New("something went wrong")

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(pointer.ToString(string(mockRedisItem)), nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - when get data from db", func(t *testing.T) {

		initEnv()

		mockErrorFromRepo := errors.New("error")
		expectedError := errors.New("something went wrong")

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(nil, mockErrorFromRepo)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - record not found", func(t *testing.T) {

		initEnv()

		expectedError := errors.New("record not found")

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(nil, nil)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Happy - when update url", func(t *testing.T) {

		initEnv()

		mockErrorFromRepo := errors.New("error")
		expectedError := errors.New("something went wrong")

		getURLResult := mockGetURLResult()
		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)
		m.DatabaseRepo.EXPECT().UpdateURL(getURLResult.ID, gomock.Any()).Return(mockErrorFromRepo)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - when set popular url", func(t *testing.T) {

		initEnv()

		mockErrorFromRedis := errors.New("error")
		expectedError := errors.New("something went wrong")

		getURLResult := mockGetURLResult()
		getURLResult.Hits = 9

		mockPopularItem := getURLResult
		mockPopularItem.Hits = getURLResult.Hits + 1
		mockPopularItemJSON, _ := json.Marshal(mockPopularItem)

		m.RedisRepo.EXPECT().Get(ctx, mockShortCode).Return(nil, nil)
		m.DatabaseRepo.EXPECT().GetURL(mockShortCode).Return(&getURLResult, nil)
		m.DatabaseRepo.EXPECT().UpdateURL(getURLResult.ID, gomock.Any()).Return(nil)
		m.RedisRepo.EXPECT().Set(ctx, mockShortCode, string(mockPopularItemJSON)).Return(mockErrorFromRedis)

		result, err := m.UseCase.Get(ctx, mockShortCode)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

}
