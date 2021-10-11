package usecase

import (
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type ArticleUsecase interface {
	GetById(id int) (article *model.Article, err error)
	ListByCursor(cursor int) (articles []*model.Article, err error)
	Create(article *model.Article) (int64, error)
	Update(article *model.Article) error
	Delete(id int) error
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

func (usecase *articleUsecase) ListByCursor(cursor int) (articles []*model.Article, err error) {
	articles, err = usecase.articleRepo.ListByCursor(cursor)
	if err != nil {
		return nil, err
	}

	return articles, err
}

func (usecase *articleUsecase) Create(article *model.Article) (int64, error) {
	id, err := usecase.articleRepo.Create(article)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (usecase *articleUsecase) Update(article *model.Article) error {
	err := usecase.articleRepo.Update(article)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *articleUsecase) Delete(id int) error {
	err := usecase.articleRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}