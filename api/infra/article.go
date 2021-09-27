package infra

import (
	"yuki0920/go-blog/model"
	"yuki0920/go-blog/repository"
)

type ArticleRepository struct {
	SqlHandler
}

func NewArticleRepository(sqlHandler SqlHandler) repository.ArticleRepository {
	articleRepository := ArticleRepository{SqlHandler: sqlHandler}

	return &articleRepository
}

func (arRepo *ArticleRepository) GetByID(id int) (*model.Article, error) {
	query := `SELECT *
	FROM articles
	WHERE id = ?;`

	var article model.Article
	if err := arRepo.SqlHandler.Conn.Get(&article, query, id); err != nil {
		return nil, err
	}

	return &article, nil
}
