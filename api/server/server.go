package server

import (
	"yuki0920/go-blog/handler"
	"yuki0920/go-blog/middleware"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Router(e *echo.Echo) *echo.Echo {
	// unused
	e.GET("/", handler.ArticleIndex)
	e.GET("/articles", handler.ArticleIndex)
	e.GET("/articles/new", handler.ArticleNew)
	e.GET("/articles/:articleID", handler.ArticleShowData)
	e.GET("/articles/:articleID/edit", handler.ArticleEdit)
	e.GET("/api/articles", handler.ArticleList)
	e.PATCH("/api/articles/:articleID", handler.ArticleUpdateData)

	e.GET("/api/articles/:articleID", handler.ArticleShow)
	e.POST("/api/login", handler.Login)
	e.GET("/api/sample", handler.ArticleSample)

	// NOTE: IsAuthenticatedのカスタムミドルウェアを利用してクッキー内のJWTトークンの検証をしている
	e.POST("/api/articles", handler.ArticleCreate, middleware.IsAuthenticated)
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete, middleware.IsAuthenticated)
	e.PUT("/api/articles/:articleID", handler.ArticleUpdate, middleware.IsAuthenticated)

	return e
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
