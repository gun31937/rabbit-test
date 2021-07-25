package database_test

import (
	"errors"
	"github.com/AlekSi/pointer"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"testing"
	"time"
)

func TestRepository_GetURL(t *testing.T) {

	mockedShortCode := "abc"

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		mockedData := getMockDataFromDB()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().Attach([]*gomocket.FakeResponse{
			{
				Pattern:  `SELECT * FROM "urls"  WHERE`,
				Response: mockedData,
				Once:     true,
			},
		})

		result, err := repository.GetURL(mockedShortCode)

		assert.NoError(t, err)
		assert.Equal(t, mockedData[0]["id"], int(result.ID))
		assert.Equal(t, mockedData[0]["short_code"], result.ShortCode)
		assert.Equal(t, mockedData[0]["full_url"], result.FullURL)
		assert.Equal(t, mockedData[0]["expiry"], result.Expiry)
		assert.Equal(t, mockedData[0]["hits"], result.Hits)
	})

	t.Run("Happy - no record", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().Attach([]*gomocket.FakeResponse{
			{
				Pattern:  `SELECT * FROM "urls"  WHERE`,
				Response: nil,
				Once:     true,
			},
		})

		result, err := repository.GetURL(mockedShortCode)

		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("Error - unexpected", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		expectedError := errors.New("error")

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		result, err := repository.GetURL(mockedShortCode)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

}

func getMockDataFromDB() []map[string]interface{} {
	mock := map[string]interface{}{
		"id":         1,
		"short_code": "abc",
		"full_url":   "https://wwww.facebook.com",
		"expiry":     pointer.ToTime(time.Now()),
		"hits":       0,
	}
	return []map[string]interface{}{mock}
}
