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

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

// @Summary Create category
// @Description Create a new category
// @Tags Categories API
// @Accept json
// @Produce json
// @Param category body web.CategoryCreateRequest true "Category to create"
// @Success 200 {object} web.WebResponse{data=web.CategoryResponse}
// @Router /categories [post]
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

// @Summary Update category
// @Description Update an existing category by id
// @Tags Categories API
// @Accept json
// @Produce json
// @Param categoryId path int true "Category Id"
// @Param category body web.CategoryUpdateRequest true "Category to update"
// @Success 200 {object} web.WebResponse{data=web.CategoryResponse}
// @Router /categories/{categoryId} [put]
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

// @Summary Delete category
// @Description Delete a category by id
// @Tags Categories API
// @Produce json
// @Param categoryId path int true "Category Id"
// @Success 200 {object} web.WebResponse
// @Router /categories/{categoryId} [delete]
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

// @Summary Get category by id
// @Description Get a single category by id
// @Tags Categories API
// @Produce json
// @Param categoryId path int true "Category Id"
// @Success 200 {object} web.WebResponse{data=web.CategoryResponse}
// @Router /categories/{categoryId} [get]
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

// @Summary Get all categories
// @Description List all categories
// @Tags Categories API
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.CategoryResponse}
// @Router /categories [get]
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