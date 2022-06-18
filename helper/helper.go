package helper

import (
	"transaction-service-v2/util"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	RequestTime string      `json:"request_time"`
	StatusCode  int         `json:"status_code"`
	Message     string      `json:"message"`
	Content     interface{} `json:"content"`
}

func APIResponse(message string, code int, content interface{}) Response {
	jsonResponse := Response{}
	jsonResponse.RequestTime = util.CTimeNow().Format("2006-01-02 15:04:05")
	jsonResponse.Message = message
	jsonResponse.StatusCode = code
	jsonResponse.Content = content

	return jsonResponse
}

func ErrorValidationResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
