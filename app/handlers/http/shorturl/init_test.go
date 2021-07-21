package shorturl_test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	handlersHttp "rabbit-test/app/handlers/http"
	"rabbit-test/app/mocks/usecases"
	"strings"
)

type mocks struct {
	ShortURLseCase *mock_usecases.MockShortURL
}

func initTest(ctrl *gomock.Controller) *mocks {
	gin.SetMode(gin.TestMode)
	m := mocks{
		ShortURLseCase: mock_usecases.NewMockShortURL(ctrl),
	}
	return &m
}

func (mocks *mocks) executeWithContext(requestUrl string, httpMethod string, jsonRequestBody *string, middleware gin.HandlerFunc) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	_, ginEngine := gin.CreateTestContext(response)
	if middleware != nil {
		ginEngine.Use(middleware)
	}

	var httpRequest *http.Request
	var reader io.Reader

	if jsonRequestBody != nil {
		reader = strings.NewReader(*jsonRequestBody)
	}

	httpRequest, _ = http.NewRequest(httpMethod, requestUrl, reader)

	handlersHttp.NewRouterShortURL(ginEngine, mocks.ShortURLseCase)
	ginEngine.ServeHTTP(response, httpRequest)
	return response
}
