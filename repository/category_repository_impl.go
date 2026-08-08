package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepositoryImpl struct {

}

	func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
		panic("Implement me")
	}

	func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
		panic("Implement me")
	}

	func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
		panic("Implement me")
	}

	func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
		panic("Implement me")
	}

	func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
		panic("Implement me")
	}
