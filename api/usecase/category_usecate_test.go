package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type mockCategoryRepository struct{}

func (mockRepo *mockCategoryRepository) Create(category *model.Category) (int64, error) {
	return 0, nil
}

func (mockRepo *mockCategoryRepository) List() ([]*model.Category, error) {
	var mockCategory model.Category
	_ = faker.FakeData(&mockCategory)

	mockCategories := make([]*model.Category, 0)
	mockCategories = append(mockCategories, &mockCategory)

	return mockCategories, nil
}

func TestCategoryCreate(t *testing.T) {
	mockRepo := &mockCategoryRepository{}
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	_, err := categoryUsecase.Create(&model.Category{})
	assert.NoError(t, err)
}

func TestCategoryList(t *testing.T) {
	mockRepo := &mockCategoryRepository{}
	categoryUsecase := usecase.NewCategoryUsecase(mockRepo)

	_, err := categoryUsecase.List()
	assert.NoError(t, err)
}
