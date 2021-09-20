package server

import (
	"yuki0920/go-blog/handler"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Router(e *echo.Echo) *echo.Echo {
	// ルーティングの設定
	e.GET("/", handler.ArticleIndex)
	e.GET("/articles", handler.ArticleIndex)
	e.GET("/articles/new", handler.ArticleNew)
	e.GET("/articles/:articleID", handler.ArticleShowData)
	e.GET("/articles/:articleID/edit", handler.ArticleEdit)
	e.GET("/api/articles", handler.ArticleList)
	e.POST("/api/articles", handler.ArticleCreate)
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete)
	e.PATCH("/api/articles/:articleID", handler.ArticleUpdateData)
	e.PUT("/api/articles/:articleID", handler.ArticleUpdate)
	e.GET("/api/articles/:articleID", handler.ArticleShow)
	e.POST("/api/login", handler.Login)
	e.GET("/api/sample", handler.ArticleSample)

	return e
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
