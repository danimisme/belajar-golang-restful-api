package exeption

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}){
	if notFoundError(writter, request, err) {
		return
	}
	InternalServerError(writter, request, err)
	
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