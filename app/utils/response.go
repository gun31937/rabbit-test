package utils

import (
	"net/http"
	"rabbit-test/app/usecases/shorturl"
)

const (
	StatusFail    string = "fail"
	StatusSuccess string = "success"
)

type BaseSuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) BaseSuccessResponse {
	r := new(BaseSuccessResponse)
	r.Status = StatusSuccess
	r.Data = data
	return *r
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	errorResponse := ErrorResponse{}
	errorResponse.Status = StatusFail
	errorResponse.Message = message
	return errorResponse
}

// GetHTTPStatusCodeWithMessage Catch error message and return with status
func GetHTTPStatusCodeWithMessage(err error) (int, string) {
	msg := err.Error()
	switch msg {

	//bad request
	case shorturl.ErrorURLFormat,
		shorturl.ErrorMatchBlacklist:
		return http.StatusBadRequest, msg
	//gone
	case shorturl.ErrorURLExpired:
		return http.StatusGone, msg
	//not found
	case shorturl.ErrorRecordNotFound:
		return http.StatusNotFound, msg
	default:
		return http.StatusInternalServerError, msg
	}
}
