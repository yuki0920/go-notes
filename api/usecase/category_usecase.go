package usecase

import (
	"fmt"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type CategoryUsecase interface {
	Create(category *model.Category) (int64, error)
	List() ([]*model.Category, error)
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (usecase *categoryUsecase) Create(category *model.Category) (int64, error) {
	id, err := usecase.categoryRepo.Create(category)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usecase *categoryUsecase) List() ([]*model.Category, error) {
	fmt.Println("usecase called start")
	categories, err := usecase.categoryRepo.List()
	fmt.Println("usecase called finished")
	if err != nil {
		return nil, err
	}

	return categories, nil
}
