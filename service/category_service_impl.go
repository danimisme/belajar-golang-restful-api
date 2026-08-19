package service

import (
	"belajar-golang-restful-api/exeption"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		validate: validate,
	}
}

func (categoryService *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := categoryService.validate.Struct(request)
	helper.PanicIfError(err)
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
	err := categoryService.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}
	category.Name = request.Name

	category = categoryService.CategoryRepository.Update(ctx, tx, category)

	return  helper.ToCategoryResponse(category)
}

func (categoryService *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, categoryId)
 	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}


	categoryService.CategoryRepository.Delete(ctx, tx, category)

}

func (categoryService *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, categoryId)
 	if err != nil {
		panic(exeption.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)	
}

func (categoryService *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)


	categories := categoryService.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}