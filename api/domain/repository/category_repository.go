package repository

import "yuki0920/go-notes/domain/model"

type CategoryRepository interface {
	Create(cateogory *model.Category) (int64, error)
	List() ([]*model.Category, error)
}
