package shorturl_test

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler_Delete(t *testing.T) {

	mockShortCode := "abc"
	requestUrl := fmt.Sprintf("%s%s", "/short-url/", mockShortCode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := initTest(ctrl)

	t.Run("Success", func(t *testing.T) {

		m.ShortURLseCase.EXPECT().Delete(gomock.Any(), mockShortCode).Return(nil)
		response := m.executeWithContext(requestUrl, http.MethodDelete, nil, nil)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Error - unexpected", func(t *testing.T) {

		errorFromUsecase := errors.New("error")

		m.ShortURLseCase.EXPECT().Delete(gomock.Any(), mockShortCode).Return(errorFromUsecase)
		response := m.executeWithContext(requestUrl, http.MethodDelete, nil, nil)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
