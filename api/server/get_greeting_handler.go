package server

import (
	"yuki0920/go-notes/server/gen/restapi/factory"

	"github.com/go-openapi/runtime/middleware"
)

func GetGreeting(p factory.GetGreetingParams) middleware.Responder {
	payload := *p.Name
	return factory.NewGetGreetingOK().WithPayload(payload)
}
