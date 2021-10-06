package repository

import "yuki0920/go-notes/domain/model"

type ArticleRepository interface {
	GetById(id int) (*model.Article, error)
	ListByCursor(cursor int) ([]*model.Article, error)
	Create(article *model.Article) (int64, error)
	Update(article *model.Article) error
}
