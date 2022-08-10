//go:generate mockgen -source=$GOFILE -package=mocks -destination=mocks/user_repository.go
package repository

import "yuki0920/go-notes/domain/model"

type UserRepository interface {
	GetByName(name string) (*model.User, error)
	Create(user *model.User) error
}
