package injector

import (
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/handler"
	"yuki0920/go-notes/infra"
	"yuki0920/go-notes/usecase"
)

func InjectAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(InjectUserUsecase())
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepository := InjectUserRepository()

	return usecase.NewUserUsecase(userRepository)
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}
