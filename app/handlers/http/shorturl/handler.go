package shorturl

import (
	"github.com/go-playground/validator/v10"
	"rabbit-test/app/usecases"
)

type Handler struct {
	Validator       *validator.Validate
	ShortURLUseCase usecases.ShortURL
}
