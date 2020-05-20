package errors

import "net/http"

type RestError struct {

	Message   string `json:"message"`
	Code 	  int 	 `json:"code"`
	Error     interface{} `json:"error"`
}

func NewBadRequest(msg string) *RestError{
	return &RestError{
		Message	: msg,
		Code	:http.StatusBadRequest,
		Error   : "bad request" ,
	}
}

func NotFoundError(msg string) *RestError{
	return &RestError{
		Message	: msg,
		Code	:http.StatusNotFound,
		Error   : "not found request" ,
	}
}



func NewInteenalServerError(msg string) * RestError  {
	return &RestError{
		Message	: msg,
		Code	:http.StatusInternalServerError,
		Error   : "error in server" ,
	}
}

func ErrorValidation(errors interface{})  * RestError {
	return &RestError{
		Message	: "erro validation",
		Code	:http.StatusBadRequest,
		Error   : errors ,
	}
}