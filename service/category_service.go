package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
}