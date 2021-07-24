package shorturl

import (
	"context"
	"errors"
)

func (u *UseCase) Delete(ctx context.Context, shortCode string) error {

	err := u.DatabaseRepo.DeleteURL(shortCode)
	if err != nil {
		return errors.New(ErrorGeneric)
	}

	err = u.RedisRepo.Delete(ctx, shortCode)
	if err != nil {
		return errors.New(ErrorGeneric)
	}

	return nil
}
