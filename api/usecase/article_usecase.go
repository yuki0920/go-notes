package usecase

import (
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type ArticleUsecase interface {
	GetById(id int) (article *model.Article, err error)
}

// リポジトリを参照するのはユースケースからのみのため、小文字始まりのstructを作成
type articleUsecase struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUsecase(articleRepo repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepo: articleRepo,
	}
}

func (usecase *articleUsecase) GetById(id int) (article *model.Article, err error) {
	article, err = usecase.articleRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return article, err
}
