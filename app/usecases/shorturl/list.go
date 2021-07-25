package shorturl

import (
	"errors"
	"github.com/jinzhu/copier"
	"rabbit-test/app/repositories/database"
)

func (u *UseCase) List(filter ListURLRequest) ([]URL, error) {

	dbFilter := database.ListURLFilterRequest{}
	_ = copier.Copy(&dbFilter, &filter)

	urls, err := u.DatabaseRepo.ListURL(dbFilter)
	if err != nil {
		return nil, errors.New(ErrorGeneric)
	}

	var listResult []URL
	_ = copier.Copy(&listResult, &urls)

	return listResult, nil
}
