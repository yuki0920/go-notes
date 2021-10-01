package usecase

import (
	"testing"
	"yuki0920/go-blog/domain/model"
)

type mockRepository struct {}

func (mockRepo *mockRepository) GetById(id int) (*model.Article, error) {
	article := model.Article{
		ID: id,
		Title: "test",
		Body: "test",
	}
	return &article, nil
}

func TestGetById(t *testing.T) {
	mockRepo := &mockRepository{}
	articleUsecase := NewArticleUsecase(mockRepo)

	article, err := articleUsecase.GetById(10)
	if err != nil {
		t.Error(err)
	}

	if article.ID != 10 {
		t.Error("title is not 10")
	}
}
