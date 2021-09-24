package server

import (
	"yuki0920/go-blog/handler"
	"yuki0920/go-blog/middleware"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Router(e *echo.Echo) *echo.Echo {
	e.POST("/api/login", handler.Login)
	e.GET("/api/auth", handler.Auth)
	e.GET("/api/sample", handler.ArticleSample)
	e.GET("/api/articles", handler.ArticleIndex)
	e.GET("/api/articles/:articleID", handler.ArticleShow)

	// NOTE: IsAuthenticatedのカスタムミドルウェアを利用してクッキー内のJWTトークンの検証をしている
	//       検証が失敗したら、エラーを返してhandlerが実行されないようにする
	e.POST("/api/logout", handler.Logout, middleware.IsAuthenticated)
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
