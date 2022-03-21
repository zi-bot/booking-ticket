package utility

import (
	"net/http"
	"strconv"
	"strings"
)

type WebResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func (r *WebResponse) Default(result interface{}, code int) *WebResponse {
	var status bool
	var err interface{}
	var data interface{}

	if strings.HasPrefix(strconv.Itoa(code), "2") {
		status = true
	} else {
		status = false
	}

	_, ok := result.(error)
	if ok {
		err = result
		data = nil
	} else {
		err = nil
		data = result
	}

	return &WebResponse{
		Status:  status,
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
		Error:   err,
	}
}
