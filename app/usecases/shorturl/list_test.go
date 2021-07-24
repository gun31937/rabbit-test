package shorturl_test

import (
	"errors"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rabbit-test/app/repositories/database"
	"rabbit-test/app/usecases/shorturl"
	"testing"
	"time"
)

func TestUseCase_List(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := initTest(ctrl)

	mockShortCode := "rQ3Pfb"
	mockFullURL := "https://www.facebook.com/"

	mockFilter := shorturl.ListURLRequest{
		ShortCode: pointer.ToString(mockShortCode),
		Keyword:   pointer.ToString("face"),
	}

	mockListURLResult := func() []database.URL {
		return []database.URL{
			{
				ID:        1,
				ShortCode: mockShortCode,
				FullURL:   mockFullURL,
				Expiry:    nil,
				Hits:      8,
			},
			{
				ID:        2,
				ShortCode: mockShortCode,
				FullURL:   mockFullURL,
				Expiry:    pointer.ToTime(time.Now()),
				Hits:      5,
			},
		}
	}

	t.Run("Happy - with records", func(t *testing.T) {

		initEnv()

		listURLResult := mockListURLResult()
		m.DatabaseRepo.EXPECT().ListURL(mock.MatchedBy(func(input database.ListURLFilterRequest) bool {
			assert.Equal(t, *mockFilter.ShortCode, *input.ShortCode)
			assert.Equal(t, *mockFilter.Keyword, *input.Keyword)
			return true
		})).Return(listURLResult, nil)

		result, err := m.UseCase.List(mockFilter)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(listURLResult), len(result))

		assert.Equal(t, listURLResult[0].ID, result[0].ID)
		assert.Equal(t, listURLResult[0].FullURL, result[0].FullURL)
		assert.Equal(t, listURLResult[0].Expiry, result[0].Expiry)
		assert.Equal(t, listURLResult[0].Hits, result[0].Hits)

		assert.Equal(t, listURLResult[1].ID, result[1].ID)
		assert.Equal(t, listURLResult[1].FullURL, result[1].FullURL)
		assert.Equal(t, listURLResult[1].Expiry, result[1].Expiry)
		assert.Equal(t, listURLResult[1].Hits, result[1].Hits)
	})

	t.Run("Happy - with no records", func(t *testing.T) {

		initEnv()

		m.DatabaseRepo.EXPECT().ListURL(gomock.Any()).Return([]database.URL{}, nil)

		result, err := m.UseCase.List(mockFilter)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, len(result))
	})

	t.Run("Error - repo return error", func(t *testing.T) {

		initEnv()

		mockRepoError := errors.New("error")
		expectedError := errors.New("something went wrong")

		m.DatabaseRepo.EXPECT().ListURL(gomock.Any()).Return(nil, mockRepoError)

		result, err := m.UseCase.List(mockFilter)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedError, err)
	})

}
