package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {}

func (mockRepo *mockRepository) GetById(id int) (*model.Article, error) {
	var mockArticle model.Article
	faker.FakeData(&mockArticle)

	return &mockArticle, nil
}

func (mockRepo *mockRepository) ListByCursor(cursor int) ([]*model.Article, error) {
	var mockArticle model.Article
	faker.FakeData(&mockArticle)

	mockArticles := make([]*model.Article, 0)
	mockArticles = append(mockArticles, &mockArticle)

	return mockArticles, nil
}

func TestGetById(t *testing.T) {
	mockRepo := &mockRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	article, err := articleUsecase.GetById(10)
	assert.NoError(t, err)
	assert.NotNil(t, article)
}

func TestListByCursor(t *testing.T) {
	mockRepo := &mockRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	articles, err := articleUsecase.ListByCursor(10)
	assert.NoError(t, err)
	assert.NotEmpty(t, articles)
}
