package repository

import "yuki0920/go-notes/domain/model"

type CategoryRepository interface {
	Create(cateogory *model.Category) error
	List() ([]*model.Category, error)
}
