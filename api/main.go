package main

import (
	"fmt"
	"net/http"
	"os"

	"yuki0920/go-notes/handler"
	"yuki0920/go-notes/injector"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := createMux()
	setupRouting(e)

	e.Validator = &handler.CustomValidator{Validator: validator.New()}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_URL")},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\nmethod=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.BodyDump(bodyDumpHandler))

	return e
}

func setupRouting(e *echo.Echo) {
	articleHandler := injector.InjectArticleHandler()
	handler.InitArticleRouting(e, articleHandler)

	authHandler := injector.InjectAuthHandler()
	handler.InitAuthRouting(e, authHandler)

	categoryHandler := injector.InjectCategoryHandler()
	handler.InitCategoryRouting(e, categoryHandler)
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}
