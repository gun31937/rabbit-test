package database_test

import (
	"errors"
	"github.com/AlekSi/pointer"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"rabbit-test/app/repositories/database"
	"testing"
	"time"
)

func TestRepository_ListURL(t *testing.T) {

	mockFilter := database.ListURLFilterRequest{
		ShortCode: pointer.ToString("abc"),
		Keyword:   pointer.ToString("face"),
	}

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		mockedData := listMockDataFromDB()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().Attach([]*gomocket.FakeResponse{
			{
				Pattern:  `SELECT * FROM "urls"  WHERE`,
				Response: mockedData,
				Once:     true,
			},
		})

		result, err := repository.ListURL(mockFilter)

		assert.NoError(t, err)
		assert.Equal(t, len(mockedData), len(result))
		assert.Equal(t, mockedData[0]["id"], int(result[0].ID))
		assert.Equal(t, mockedData[0]["short_code"], result[0].ShortCode)
		assert.Equal(t, mockedData[0]["full_url"], result[0].FullUrl)
		assert.Equal(t, mockedData[0]["expiry"], result[0].Expiry)
		assert.Equal(t, mockedData[0]["hits"], result[0].Hits)

		assert.Equal(t, mockedData[1]["id"], int(result[1].ID))
		assert.Equal(t, mockedData[1]["short_code"], result[1].ShortCode)
		assert.Equal(t, mockedData[1]["full_url"], result[1].FullUrl)
		assert.Equal(t, mockedData[1]["expiry"], result[1].Expiry)
		assert.Equal(t, mockedData[1]["hits"], result[1].Hits)
	})

	t.Run("Error - unexpected", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		expectedError := errors.New("error")

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		result, err := repository.ListURL(mockFilter)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, result)
	})

}

func listMockDataFromDB() []map[string]interface{} {
	mock := map[string]interface{}{
		"id":         1,
		"short_code": "abc",
		"full_url":   "https://wwww.facebook.com",
		"expiry":     pointer.ToTime(time.Now()),
		"hits":       0,
	}
	return []map[string]interface{}{mock, mock}
}
