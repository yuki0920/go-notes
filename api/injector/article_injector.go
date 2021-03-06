package injector

import (
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/handler"
	"yuki0920/go-notes/infra"
	"yuki0920/go-notes/usecase"
)

func InjectArticleHandler() handler.ArticleHandler {
	return handler.NewArticleHandler(InjectArticleUsecase())
}

func InjectArticleUsecase() usecase.ArticleUsecase {
	articleRepository := InjectArticleRepository()

	return usecase.NewArticleUsecase(articleRepository)
}

func InjectArticleRepository() repository.ArticleRepository {
	sqlHandler := InjectDB()

	return infra.NewArticleRepository(sqlHandler)
}
