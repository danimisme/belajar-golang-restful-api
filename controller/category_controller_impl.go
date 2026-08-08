package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}


func (controller *CategoryControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequest(request, &categoryCreateRequest)
	categoryResponse := controller.categoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data : categoryResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)

}

func (controller *CategoryControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params){

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequest(request, &categoryUpdateRequest)

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

	helper.WriteToResponseBody(writter, webResponse)
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

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	
	findByIdResponse := controller.categoryService.FindById(request.Context(), id)

		controller.categoryService.FindById(request.Context(), id)
		webResponse := web.WebResponse{
			Code : 200,
			Status: "OK",
			Data: findByIdResponse,
		}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params){
	
	categoryResponses := controller.categoryService.FindAll(request.Context())

		controller.categoryService.FindAll(request.Context())
		webResponse := web.WebResponse{
			Code : 200,
			Status: "OK",
			Data: categoryResponses,
		}

	helper.WriteToResponseBody(writter, webResponse)
}