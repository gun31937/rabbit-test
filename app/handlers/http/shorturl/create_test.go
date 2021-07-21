package shorturl_test

import (
	"errors"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"rabbit-test/app/usecases/shorturl"
	"testing"
	"time"
)

func TestHandler_Create(t *testing.T) {

	requestUrl := "/short-url/create"
	mockFullURL := "https://www.facebook.com"
	mockExpiry := pointer.ToInt(30)
	timeFormat := "2006-01-02 15:04:05"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := initTest(ctrl)

	t.Run("Success", func(t *testing.T) {

		jsonRequestBody := `{
			"fullURL": "` + mockFullURL + `",
			"expiry": 30
		}`

		mockUseCaseResponse := shorturl.CreateShortURLResponse{
			ShortURL:    "http://localhost:8080/e",
			ExpiredTime: pointer.ToString(time.Now().Add(time.Duration(*mockExpiry) * time.Minute).Format(timeFormat)),
		}

		m.ShortURLseCase.EXPECT().Create(gomock.Any(), mockFullURL, mockExpiry).Return(&mockUseCaseResponse, nil)
		response := m.executeWithContext(requestUrl, http.MethodPost, pointer.ToString(jsonRequestBody), nil)
		responseJSON := jsoniter.Get(response.Body.Bytes(), "data")

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, mockUseCaseResponse.ShortURL, responseJSON.Get("shortURL").ToString())
		assert.Equal(t, *mockUseCaseResponse.ExpiredTime, responseJSON.Get("expiredTime").ToString())
	})

	t.Run("Success - with no expiry", func(t *testing.T) {

		jsonRequestBody := `{
			"fullURL": "` + mockFullURL + `"
		}`

		mockUseCaseResponse := shorturl.CreateShortURLResponse{
			ShortURL:    "http://localhost:8080/e",
			ExpiredTime: nil,
		}

		m.ShortURLseCase.EXPECT().Create(gomock.Any(), mockFullURL, nil).Return(&mockUseCaseResponse, nil)
		response := m.executeWithContext(requestUrl, http.MethodPost, pointer.ToString(jsonRequestBody), nil)
		responseJSON := jsoniter.Get(response.Body.Bytes(), "data")

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, mockUseCaseResponse.ShortURL, responseJSON.Get("shortURL").ToString())
		assert.Equal(t, "", responseJSON.Get("expiredTime").ToString())
	})

	t.Run("Error - required field", func(t *testing.T) {

		jsonRequestBody := `{
			"fullURL": ""
		}`

		response := m.executeWithContext(requestUrl, http.MethodPost, pointer.ToString(jsonRequestBody), nil)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Error - usecase return error", func(t *testing.T) {

		jsonRequestBody := `{
			"fullURL": "` + mockFullURL + `",
			"expiry": 30
		}`

		mockUsecaseError := errors.New("error")

		m.ShortURLseCase.EXPECT().Create(gomock.Any(), mockFullURL, mockExpiry).Return(nil, mockUsecaseError)
		response := m.executeWithContext(requestUrl, http.MethodPost, pointer.ToString(jsonRequestBody), nil)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
