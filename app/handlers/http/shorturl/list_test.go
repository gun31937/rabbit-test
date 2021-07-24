package shorturl_test

import (
	"errors"
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"rabbit-test/app/usecases/shorturl"
	"strings"
	"testing"
	"time"
)

func TestHandler_List(t *testing.T) {

	mockShortCode := "abc"
	mockKeyword := "face"
	loginRequestUrl := "/login"

	mockQueryParams := func() map[string]string {
		return map[string]string{
			"shortCode": mockShortCode,
			"keyword":   mockKeyword,
		}
	}

	buildRequestURL := func(params map[string]string) string {
		requestURL := "/?"
		for paramName, paramValue := range params {
			if paramValue != "" {
				requestURL += fmt.Sprintf("&%s=%s", paramName, paramValue)
			}
		}
		requestURL = strings.Replace(requestURL, "&", "", 1)
		return requestURL
	}

	mockUseCaseResponse := []shorturl.URL{
		{
			ID:        1,
			ShortCode: "abc",
			FullURL:   "https://www.facebook.com/",
			Expiry:    nil,
			Hits:      4,
		},
		{
			ID:        2,
			ShortCode: "def",
			FullURL:   "https://www.facebook.com/",
			Expiry:    pointer.ToTime(time.Now()),
			Hits:      19,
		},
	}

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

	t.Run("Success - with records", func(t *testing.T) {

		params := mockQueryParams()
		requestURL := buildRequestURL(params)

		m.ShortURLseCase.EXPECT().List(mock.MatchedBy(func(input shorturl.ListURLRequest) bool {
			assert.Equal(t, mockShortCode, *input.ShortCode)
			assert.Equal(t, mockKeyword, *input.Keyword)
			return true
		})).Return(mockUseCaseResponse, nil)

		response := m.executeWithContextWithAuthToken(requestURL, http.MethodGet, nil, getAuthToken())
		responseJSON := jsoniter.Get(response.Body.Bytes(), "data")

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, len(mockUseCaseResponse), responseJSON.Size())

		location, _ := time.LoadLocation("Asia/Bangkok")
		expiry1 := mockUseCaseResponse[1].Expiry.In(location).Format("2006-01-02 15:04:05")

		assert.Equal(t, mockUseCaseResponse[0].ID, responseJSON.Get(0, "id").ToUint())
		assert.Equal(t, mockUseCaseResponse[0].ShortCode, responseJSON.Get(0, "shortCode").ToString())
		assert.Equal(t, mockUseCaseResponse[0].FullURL, responseJSON.Get(0, "fullURL").ToString())
		assert.Equal(t, mockUseCaseResponse[0].Hits, responseJSON.Get(0, "hits").ToInt())

		assert.Equal(t, mockUseCaseResponse[1].ID, responseJSON.Get(1, "id").ToUint())
		assert.Equal(t, mockUseCaseResponse[1].ShortCode, responseJSON.Get(1, "shortCode").ToString())
		assert.Equal(t, mockUseCaseResponse[1].FullURL, responseJSON.Get(1, "fullURL").ToString())
		assert.Equal(t, expiry1, responseJSON.Get(1, "expiry").ToString())
		assert.Equal(t, mockUseCaseResponse[1].Hits, responseJSON.Get(1, "hits").ToInt())
	})

	t.Run("Success - with no records", func(t *testing.T) {

		params := mockQueryParams()
		requestURL := buildRequestURL(params)

		m.ShortURLseCase.EXPECT().List(gomock.Any()).Return([]shorturl.URL{}, nil)

		response := m.executeWithContextWithAuthToken(requestURL, http.MethodGet, nil, getAuthToken())
		responseJSON := jsoniter.Get(response.Body.Bytes(), "data")

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, 0, responseJSON.Size())

	})

	t.Run("Error - unexpected", func(t *testing.T) {

		errorFromUsecase := errors.New("error")

		params := mockQueryParams()
		requestURL := buildRequestURL(params)

		m.ShortURLseCase.EXPECT().List(gomock.Any()).Return(nil, errorFromUsecase)
		response := m.executeWithContextWithAuthToken(requestURL, http.MethodGet, nil, getAuthToken())

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}
