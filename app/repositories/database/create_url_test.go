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

func TestRepository_CreateURL(t *testing.T) {

	exp := time.Now().Add(time.Hour)
	expectedId := uint(2)

	mockRequest := database.CreateShortURLRequest{
		ShortCode: "abcdefg",
		FullURL:   "https://www.example.com/path",
		Expiry:    pointer.ToTime(exp),
	}

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().NewMock().WithQuery(`INSERT INTO "urls"`).WithID(int64(expectedId))

		id, err := repository.CreateURL(mockRequest)

		assert.NoError(t, err)
		assert.Equal(t, expectedId, *id)
	})

	t.Run("Error unexpected error", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)
		expectedError := errors.New("error")

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		id, err := repository.CreateURL(mockRequest)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, id)

	})

}
