package usecase

import (
	"testing"
	"yuki0920/go-notes/domain/model"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {}

func (mockRepo *mockRepository) GetById(id int) (*model.Article, error) {
	var mockArticle model.Article
	faker.FakeData(&mockArticle)

	return &mockArticle, nil
}

func TestGetById(t *testing.T) {
	mockRepo := &mockRepository{}
	articleUsecase := NewArticleUsecase(mockRepo)

	_, err := articleUsecase.GetById(10)
	assert.NoError(t, err)
}
