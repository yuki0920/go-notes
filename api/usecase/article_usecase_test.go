package usecase_test

import (
	"reflect"
	"testing"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/usecase"
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
	_ = faker.FakeData(&mockArticle)
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
	_ = faker.FakeData(&article)

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
	_ = faker.FakeData(&article)

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

// Here is table driven test version
func TestNewArticleUsecase(t *testing.T) {
	type args struct {
		articleRepo repository.ArticleRepository
	}
	tests := []struct {
		name string
		args args
		want usecase.ArticleUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := usecase.NewArticleUsecase(tt.args.articleRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticleUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleUsecase_GetById(t *testing.T) {
	type args struct {
		id int
	}

	article := model.Article{}
	_ = faker.FakeData(&article)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		args        args
		wantArticle *model.Article
		wantErr     bool
		setUsecase  func(*mocks.MockArticleUsecase)
	}{
		{
			name: "success",
			args: args{
				id: 10,
			},
			wantArticle: &article,
			wantErr:     false,
			setUsecase: func(mock *mocks.MockArticleUsecase) {
				mock.EXPECT().GetById(10).Return(&article, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockArticleUsecase(ctrl)
			tt.setUsecase(mock)
			gotArticle, err := mock.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleUsecase.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArticle, tt.wantArticle) {
				t.Errorf("articleUsecase.GetById() = %v, want %v", gotArticle, tt.wantArticle)
			}
		})
	}
}

func Test_articleUsecase_ListByPage(t *testing.T) {
	type args struct {
		page int
	}

	var mockArticle model.Article
	_ = faker.FakeData(&mockArticle)
	mockArticles := make([]*model.Article, 0)
	mockArticles = append(mockArticles, &mockArticle)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		args          args
		wantArticles  []*model.Article
		wantTotalPage int
		wantErr       bool
		setUsecase    func(*mocks.MockArticleUsecase)
	}{
		{
			name: "success",
			args: args{
				page: 2,
			},
			wantArticles:  mockArticles,
			wantTotalPage: 3,
			wantErr:       false,
			setUsecase: func(mock *mocks.MockArticleUsecase) {
				mock.EXPECT().ListByPage(2).Return(mockArticles, 3, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockArticleUsecase(ctrl)
			tt.setUsecase(mock)
			gotArticles, gotTotalPage, err := mock.ListByPage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleUsecase.ListByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArticles, tt.wantArticles) {
				t.Errorf("articleUsecase.ListByPage() gotArticles = %v, want %v", gotArticles, tt.wantArticles)
			}
			if gotTotalPage != tt.wantTotalPage {
				t.Errorf("articleUsecase.ListByPage() gotTotalPage = %v, want %v", gotTotalPage, tt.wantTotalPage)
			}
		})
	}
}

func Test_articleUsecase_Create(t *testing.T) {
	type args struct {
		article *model.Article
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	article := model.Article{}
	_ = faker.FakeData(&article)

	tests := []struct {
		name       string
		args       args
		want       int64
		wantErr    bool
		setUsecase func(*mocks.MockArticleUsecase)
	}{
		{
			name: "success",
			args: args{
				article: &article,
			},
			want:    int64(article.ID),
			wantErr: false,
			setUsecase: func(mock *mocks.MockArticleUsecase) {
				mock.EXPECT().Create(&article).Return(int64(article.ID), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockArticleUsecase(ctrl)
			tt.setUsecase(mock)
			got, err := mock.Create(tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("articleUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("articleUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_articleUsecase_Update(t *testing.T) {
	type args struct {
		article *model.Article
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	article := model.Article{}
	_ = faker.FakeData(&article)

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		setUsecase func(*mocks.MockArticleUsecase)
	}{
		{
			name: "success",
			args: args{
				article: &article,
			},
			wantErr: false,
			setUsecase: func(mock *mocks.MockArticleUsecase) {
				mock.EXPECT().Update(&article).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockArticleUsecase(ctrl)
			tt.setUsecase(mock)
			if err := mock.Update(tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("articleUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_articleUsecase_Delete(t *testing.T) {
	type args struct {
		id int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		setUsecase func(*mocks.MockArticleUsecase)
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantErr: false,
			setUsecase: func(mock *mocks.MockArticleUsecase) {
				mock.EXPECT().Delete(1).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := mocks.NewMockArticleUsecase(ctrl)
			tt.setUsecase(mock)
			if err := mock.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("articleUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
