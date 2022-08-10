//go:generate mockgen -source=$GOFILE -package=mocks -destination=mocks/user_usecase.go
package usecase

import (
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type UserUsecase interface {
	GetByName(name string) (user *model.User, err error)
	Create(user *model.User) error
}

// リポジトリを参照するのはユースケースからのみのため、小文字始まりのstructを作成
type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (usecase *userUsecase) GetByName(name string) (user *model.User, err error) {
	user, err = usecase.userRepo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (usecase *userUsecase) Create(user *model.User) error {
	err := usecase.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
