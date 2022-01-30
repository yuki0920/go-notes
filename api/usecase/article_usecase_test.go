package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

type mockArticleRepository struct{}

func (mockRepo *mockArticleRepository) GetById(id int) (*model.Article, error) {
	var mockArticle model.Article
	faker.FakeData(&mockArticle)

	return &mockArticle, nil
}

func (mockRepo *mockArticleRepository) ListByPage(page int) ([]*model.Article, int, error) {
	var mockArticle model.Article
	faker.FakeData(&mockArticle)

	mockArticles := make([]*model.Article, 0)
	mockArticles = append(mockArticles, &mockArticle)

	return mockArticles, page, nil
}

func (mockRepo *mockArticleRepository) Create(article *model.Article) (int64, error) {
	return int64(article.ID), nil
}

func (mockRepo *mockArticleRepository) Update(article *model.Article) error {
	return nil
}

func (mockRepo *mockArticleRepository) Delete(id int) error {
	return nil
}

func (mockRepo *mockArticleRepository) CreateCategories(article *model.Article) error {
	return nil
}

func (mockRepo *mockArticleRepository) DeleteCategories(id int) error {
	return nil
}

func TestArticleGetById(t *testing.T) {
	mockRepo := &mockArticleRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	article, err := articleUsecase.GetById(10)
	assert.NoError(t, err)
	assert.NotNil(t, article)
}

func TestArticleListByPage(t *testing.T) {
	mockRepo := &mockArticleRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	articles, _, err := articleUsecase.ListByPage(2)
	assert.NoError(t, err)
	assert.NotEmpty(t, articles)
}

func TestArticleCreate(t *testing.T) {
	mockRepo := &mockArticleRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	article := model.Article{}
	faker.FakeData(&article)

	id, err := articleUsecase.Create(&article)
	assert.NoError(t, err)
	assert.Equal(t, id, int64(article.ID))
}

func TestArticleUpdate(t *testing.T) {
	mockRepo := &mockArticleRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	article := model.Article{}
	faker.FakeData(&article)

	err := articleUsecase.Update(&article)
	assert.NoError(t, err)
}

func TestArticleDelete(t *testing.T) {
	mockRepo := &mockArticleRepository{}
	articleUsecase := usecase.NewArticleUsecase(mockRepo)

	err := articleUsecase.Delete(1)
	assert.NoError(t, err)
}
