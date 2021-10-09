package injector

import (
	"yuki0920/go-notes/infra"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()

	return *sqlhandler
}
