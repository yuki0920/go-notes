//go:generate mockgen -source=$GOFILE -package=mocks -destination=mocks/category_repository.go
package repository

import "yuki0920/go-notes/domain/model"

type CategoryRepository interface {
	Create(cateogory *model.Category) (int64, error)
	List() ([]*model.Category, error)
}
