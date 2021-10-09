package injector

import (
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/handler"
	"yuki0920/go-notes/infra"
	"yuki0920/go-notes/usecase"
)

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepository := InjectUserRepository()

	return usecase.NewUserUsecase(userRepository)
}

func InjectAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(InjectUserUsecase())
}
