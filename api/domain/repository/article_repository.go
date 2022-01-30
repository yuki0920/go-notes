package repository

import "yuki0920/go-notes/domain/model"

type ArticleRepository interface {
	GetById(id int) (*model.Article, error)
	ListByPage(page int) ([]*model.Article, int, error)
	Create(article *model.Article) (int64, error)
	Update(article *model.Article) error
	Delete(id int) error
	CreateCategories(id int, article *model.Article) error
	DeleteCategories(id int) error
}
