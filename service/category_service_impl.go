package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
}

func (categoryService *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name : request.Name,
	}

	category = categoryService.CategoryRepository.Save(ctx, tx, category)
	
	return helper.ToCategoryResponse(category)
}

func (categoryService *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, request.Id)
  helper.PanicIfError(err)
	category.Name = request.Name

	category = categoryService.CategoryRepository.Update(ctx, tx, category)

	return  helper.ToCategoryResponse(category)
}