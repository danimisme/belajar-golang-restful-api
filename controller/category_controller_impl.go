package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}


func (controller *CategoryControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	decoder := json.NewDecoder(request.Body)

	categoryCreateRequest := web.CategoryCreateRequest{}
	decoder.Decode(&categoryCreateRequest)

	categoryResponse := controller.categoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data : categoryResponse,
	}

	writter.Header().Add("Content-Type", "aplication/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)

}