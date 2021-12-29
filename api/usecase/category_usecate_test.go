package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/stretchr/testify/assert"
)

type mockCategoryRepository struct{}

func (mockRepo *mockCategoryRepository) Create(category *model.Category) error {
	return nil
}

func TestCategoryCreate(t *testing.T) {
	mockRepo := &mockCategoryRepository{}
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	err := categoryUsecase.Create(&model.Category{})
	assert.NoError(t, err)
}
