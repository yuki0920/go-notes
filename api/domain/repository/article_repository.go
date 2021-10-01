package repository

import "yuki0920/go-blog/domain/model"

type ArticleRepository interface {
	GetById(id int) (*model.Article, error)
}
