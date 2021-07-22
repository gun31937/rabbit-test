package shorturl_test

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler_Get(t *testing.T) {

	mockShortCode := "abc"
	requestUrl := fmt.Sprintf("%s%s", "/", mockShortCode)
	mockFullURL := "https://www.facebook.com"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := initTest(ctrl)

	t.Run("Success", func(t *testing.T) {

		m.ShortURLseCase.EXPECT().Get(gomock.Any(), mockShortCode).Return(&mockFullURL, nil)
		response := m.executeWithContext(requestUrl, http.MethodGet, nil, nil)

		assert.Equal(t, http.StatusFound, response.Code)
	})

	t.Run("Error - not found", func(t *testing.T) {

		errorFromUsecase := errors.New("record not found")

		m.ShortURLseCase.EXPECT().Get(gomock.Any(), mockShortCode).Return(nil, errorFromUsecase)
		response := m.executeWithContext(requestUrl, http.MethodGet, nil, nil)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Error - url expired", func(t *testing.T) {

		errorFromUsecase := errors.New("url expired")

		m.ShortURLseCase.EXPECT().Get(gomock.Any(), mockShortCode).Return(nil, errorFromUsecase)
		response := m.executeWithContext(requestUrl, http.MethodGet, nil, nil)

		assert.Equal(t, http.StatusGone, response.Code)
	})

	t.Run("Error - unexpected", func(t *testing.T) {

		errorFromUsecase := errors.New("error")

		m.ShortURLseCase.EXPECT().Get(gomock.Any(), mockShortCode).Return(nil, errorFromUsecase)
		response := m.executeWithContext(requestUrl, http.MethodGet, nil, nil)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
