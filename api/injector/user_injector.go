package injector

import (
	"yuki0920/go-notes/domain/repository"
	"yuki0920/go-notes/infra"
)

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()

	return infra.NewUserRepository(sqlHandler)
}
