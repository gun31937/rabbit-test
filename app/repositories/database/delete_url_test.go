package database_test

import (
	"errors"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"testing"
)

func TestRepository_DeleteURL(t *testing.T) {

	mockShortcode := "abc"

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.Reset().NewMock().WithQuery(`UPDATE "urls" SET`)

		err := repository.DeleteURL(mockShortcode)

		assert.NoError(t, err)
	})

	t.Run("Error - unexpected", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)
		expectedError := errors.New("error")

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		err := repository.DeleteURL(mockShortcode)

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)

	})

}
