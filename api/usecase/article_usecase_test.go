package usecase_test

import (
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase/mocks"

	"github.com/bxcodec/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArticleGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	articleUsecase := mocks.NewMockArticleUsecase(ctrl)
	articleUsecase.EXPECT().GetById(10).Return(&model.Article{}, nil)

	article, err := articleUsecase.GetById(10)
	assert.NoError(t, err)
	assert.NotNil(t, article)
}

func TestArticleListByPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var mockArticle model.Article
	faker.FakeData(&mockArticle)
	mockArticles := make([]*model.Article, 0)
	mockArticles = append(mockArticles, &mockArticle)

	articleUsecase := mocks.NewMockArticleUsecase(ctrl)
	articleUsecase.EXPECT().ListByPage(2).Return(mockArticles, 3, nil)

	articles, totalPage, err := articleUsecase.ListByPage(2)
	assert.NoError(t, err)
	assert.Equal(t, totalPage, 3)
	assert.NotEmpty(t, articles)
}

func TestArticleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	article := model.Article{}
	faker.FakeData(&article)

	articleUsecase := mocks.NewMockArticleUsecase(ctrl)
	articleUsecase.EXPECT().Create(&article).Return(int64(article.ID), nil)

	id, err := articleUsecase.Create(&article)
	assert.NoError(t, err)
	assert.Equal(t, id, int64(article.ID))
}

func TestArticleUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	article := model.Article{}
	faker.FakeData(&article)

	articleUsecase := mocks.NewMockArticleUsecase(ctrl)
	articleUsecase.EXPECT().Update(&article).Return(nil)

	err := articleUsecase.Update(&article)
	assert.NoError(t, err)
}

func TestArticleDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	articleUsecase := mocks.NewMockArticleUsecase(ctrl)
	articleUsecase.EXPECT().Delete(1).Return(nil)

	err := articleUsecase.Delete(1)
	assert.NoError(t, err)
}
