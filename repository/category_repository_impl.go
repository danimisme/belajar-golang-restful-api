package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {

}

	func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
		SQL := "INSERT INTO categories(name) VALUES (?)"
		result, err := tx.ExecContext(ctx, SQL, category.Name)
		helper.PanicIfError(err)

		id, err := result.LastInsertId()
		helper.PanicIfError(err)

		category.Id = int(id)
		return category
	}

	func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
		SQL := "UPATE categories SET name = ? WHERE id = ?"
		_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
		helper.PanicIfError(err)

		return category
	}

	func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
		SQL := "DELETE FROM categories WHERE id = ?"
		_, err := tx.ExecContext(ctx, SQL, category.Id)
		helper.PanicIfError(err)
	}

	func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
		SQL := "SELECT * FROM categories WHERE id = ?"
		rows, err := tx.QueryContext(ctx, SQL, categoryId)
		helper.PanicIfError(err)

		category := domain.Category{}
		if rows.Next() {
			err := rows.Scan(&category.Id, &category.Name)
			helper.PanicIfError(err)
			return category, nil
		} else {
			return category, errors.New("category not found")
		}
	}

	func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
		panic("Implement me")
	}
