package repository

import "yuki0920/go-notes/domain/model"

type ArticleRepository interface {
	GetById(id int) (*model.Article, error)
}
