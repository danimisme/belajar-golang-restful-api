package exeption

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}){
	if notFoundError(writter, request, err) {
		return
	}

	if validationError(writter, request, err) {
		return
	}
	
	InternalServerError(writter, request, err)
	
}

func validationError(writter http.ResponseWriter, request *http.Request, err interface{}) bool{
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "Bad Request",
			Data: exception.Error(),
		}
		helper.WriteToResponseBody(writter, webResponse)
		return true
	}
	return false
}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool{
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Data: exception.Error,
		}
		helper.WriteToResponseBody(writter, webResponse)
		return true
	}
	return false
}

func InternalServerError(writter http.ResponseWriter, request *http.Request, err interface{}){
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)
	
	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "Internal Server Error",
	}

	helper.WriteToResponseBody(writter, webResponse)
}