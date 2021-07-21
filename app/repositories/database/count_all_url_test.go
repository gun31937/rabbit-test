package database_test

import (
	"errors"
	gomocket "github.com/Selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"rabbit-test/app/repositories"
	"testing"
)

func TestRepository_CountAllURL(t *testing.T) {

	expectedCount := uint64(2)

	t.Run("Happy", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)

		gomocket.Catcher.NewMock().
			WithQuery(`SELECT count(*) FROM "urls"`).WithReply([]map[string]interface{}{
			{
				"count": expectedCount,
			},
		})

		count, err := repository.CountAllURL()

		assert.NoError(t, err)
		assert.Equal(t, expectedCount, *count)
	})

	t.Run("Error unexpected error", func(t *testing.T) {
		db := mockingDB()
		defer func() {
			_ = db.Close()
		}()

		repository := repositories.InitDatabase(db)
		expectedError := errors.New("error")

		gomocket.Catcher.Reset().NewMock().WithError(expectedError)

		count, err := repository.CountAllURL()

		assert.Error(t, err)
		assert.Exactly(t, expectedError, err)
		assert.Nil(t, count)

	})

}
