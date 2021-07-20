package database_test

import (
	"errors"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"rabbit-test/app/repositories/database"
	"testing"
	"time"
)

func TestRepository_CreateUrl(t *testing.T) {

	exp := time.Now().Add(time.Hour)

	mockRequest := database.CreateShortUrlRequest{
		ShortCode: "abcdefg",
		FullUrl:   "https://www.example.com/path",
		Expiry:    &exp,
	}

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().Attach([]*gomocket.FakeResponse{
			{
				Pattern: `INSERT INTO "urls"`,
				Once:    true,
			},
		})

		err := repository.CreateUrl(mockRequest)

		assert.NoError(t, err)
	})

	t.Run("Error unexpected error", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)
		expectedError := errors.New("error")

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		err := repository.CreateUrl(mockRequest)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
	})

}
