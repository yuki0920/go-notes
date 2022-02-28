package injector

import (
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/handler"
	"yuki0920/go-notes/infra"
	"yuki0920/go-notes/usecase"
)

func InjectCategoryHandler() handler.CategoryHandler {
	return handler.NewCategoryHandler(InjectUsecase())
}

func InjectUsecase() usecase.CategoryUsecase {
	categoryRepository := InjectCategoryRepository()

	return usecase.NewCategoryUsecase(categoryRepository)
}

func InjectCategoryRepository() repository.CategoryRepository {
	sqlHandler := InjectDB()

	return infra.NewCategoryRepository(sqlHandler)
}
