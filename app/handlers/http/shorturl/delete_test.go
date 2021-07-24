package shorturl_test

import (
	"errors"
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler_Delete(t *testing.T) {

	mockShortCode := "abc"
	loginRequestUrl := "/login"
	requestUrl := fmt.Sprintf("%s%s", "/short-url/", mockShortCode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := initTest(ctrl)

	getAuthToken := func() string {
		jsonRequestBody := `{
			"username": "admin",
			"password": "admin"
		}`
		responseToken := m.executeWithContext(loginRequestUrl, http.MethodPost, pointer.ToString(jsonRequestBody), nil)
		return jsoniter.Get(responseToken.Body.Bytes(), "token").ToString()
	}

	t.Run("Success", func(t *testing.T) {

		m.ShortURLseCase.EXPECT().Delete(gomock.Any(), mockShortCode).Return(nil)
		response := m.executeWithContextWithAuthToken(requestUrl, http.MethodDelete, nil, getAuthToken())

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Error - unexpected", func(t *testing.T) {

		errorFromUsecase := errors.New("error")

		m.ShortURLseCase.EXPECT().Delete(gomock.Any(), mockShortCode).Return(errorFromUsecase)
		response := m.executeWithContextWithAuthToken(requestUrl, http.MethodDelete, nil, getAuthToken())

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
