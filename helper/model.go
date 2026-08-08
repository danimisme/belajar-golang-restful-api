package helper

import (
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id : category.Id,
		Name : category.Name,
	}
}
