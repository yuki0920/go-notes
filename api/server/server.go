package server

import (
	"yuki0920/go-blog/handler"

	"github.com/labstack/echo/v4"
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
	e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)
	e.GET("/api/articles/:articleID", handler.ArticleShow)

	return e
}
