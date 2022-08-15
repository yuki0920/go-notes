package infra

import (
	"fmt"
	"time"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type CategoryRepository struct {
	SqlHandler
}

func NewCategoryRepository(sqlHandler SqlHandler) repository.CategoryRepository {
	return &CategoryRepository{
		SqlHandler: sqlHandler,
	}
}

func (categoryRepository *CategoryRepository) Create(category *model.Category) (int64, error) {
	now := time.Now()
	category.Created = now
	category.Updated = now

	query := `INSERT INTO categories (title, created, updated) VALUES (:title, :created, :updated);`

	tx := categoryRepository.SqlHandler.Conn.MustBegin()
	res, err := tx.NamedExec(query, category)
	if err != nil {
		err = fmt.Errorf("failed to create category: %w", err)
		_ = tx.Rollback()

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("failed to get last insert id: %w", err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (categoryRepository *CategoryRepository) List() ([]*model.Category, error) {
	query := `SELECT * FROM categories ORDER BY title;`

	var categories []*model.Category
	err := categoryRepository.SqlHandler.Conn.Select(&categories, query)
	if err != nil {
		err = fmt.Errorf("failed to list categories: %w", err)
		return nil, err
	}

	return categories, nil
}
