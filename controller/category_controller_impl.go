package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"encoding/json"
	"net/http"
	"strconv"

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

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)

}

func (controller *CategoryControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	decoder := json.NewDecoder(request.Body)

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	decoder.Decode(&categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id , err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.categoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data : categoryResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.categoryService.Delete(request.Context(), id)
		webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *CategoryControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	
	findByIdResponse := controller.categoryService.FindById(request.Context(), id)

		controller.categoryService.Delete(request.Context(), id)
		webResponse := web.WebResponse{
			Code : 200,
			Status: "OK",
			Data: findByIdResponse,
		}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
