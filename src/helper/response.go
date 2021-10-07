package helper

import "strings"

type Response struct {
	Success  bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(success bool, message string, data interface{}) Response {
	res := Response{
		Success:  success,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Success:  false,
		Message: "Uh oh, error happened :<",
		Errors:  splittedError,
		Data:    data,
	}
	return res
}