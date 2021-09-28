package injector

import (
	"yuki0920/go-blog/handler"
	"yuki0920/go-blog/infra"
	"yuki0920/go-blog/repository"
	"yuki0920/go-blog/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()

	return *sqlhandler
}

func InjectArticleRepository() repository.ArticleRepository {
	sqlHandler := InjectDB()

	return infra.NewArticleRepository(sqlHandler)
}

func InjectArticleUsecase() usecase.ArticleUsecase {
	articleRepository := InjectArticleRepository()

	return usecase.NewArticleUsecase(articleRepository)
}

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}
