package repository

import "yuki0920/go-blog/domain/model"

type ArticleRepository interface {
	GetByID(id int) (*model.Article, error)
}
