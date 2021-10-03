package handler

import (
	"yuki0920/go-notes/middleware"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Router(e *echo.Echo) {
	e.POST("/api/login", Login)
	e.GET("/api/auth", Auth)
	e.GET("/api/sample", ArticleSample)
	e.GET("/api/articles", ArticleIndex)
	e.GET("/api/articles/:articleID", ArticleShow)

	// NOTE: IsAuthenticatedのカスタムミドルウェアを利用してクッキー内のJWTトークンの検証をしている
	//       検証が失敗したら、エラーを返して実行されないようにする
	e.POST("/api/logout", Logout, middleware.IsAuthenticated)
	e.POST("/api/articles", ArticleCreate, middleware.IsAuthenticated)
	e.DELETE("/api/articles/:articleID", ArticleDelete, middleware.IsAuthenticated)
	e.PUT("/api/articles/:articleID", ArticleUpdate, middleware.IsAuthenticated)
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func InitRouting(e *echo.Echo, articleHandler ArticleHandler) {
	e.GET("/api/v2/articles/:articleID", articleHandler.Show())
	e.GET("/api/v2/articles", articleHandler.Index())
}
