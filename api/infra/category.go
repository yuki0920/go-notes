package infra

import (
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

func (categoryRepository *CategoryRepository) Create(category *model.Category) error {
	now := time.Now()
	category.Created = now
	category.Updated = now

	query := `INSERT INTO categories (title, created, updated) VALUES (:title, :created, :updated);`

	tx := categoryRepository.SqlHandler.Conn.MustBegin()
	_, err := tx.NamedExec(query, category)
	if err != nil {
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (categoryRepository *CategoryRepository) List() ([]*model.Category, error) {
	query := `SELECT * FROM categories ORDER BY title;`

	var categories []*model.Category
	err := categoryRepository.SqlHandler.Conn.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
