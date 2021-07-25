package database_test

import (
	"errors"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"rabbit-test/app/repositories/database"
	"testing"
)

func TestRepository_UpdateURL(t *testing.T) {

	id := uint(2)

	mockRequest := database.UpdateShortURLRequest{
		Hits: 2,
	}

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().NewMock().WithQuery(`UPDATE "urls" SET`)

		err := repository.UpdateURL(id, mockRequest)

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

		err := repository.UpdateURL(id, mockRequest)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)

	})

}
