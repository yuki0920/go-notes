package handler

import (
	"yuki0920/go-notes/middleware"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Router(e *echo.Echo) {
	e.GET("/api/auth", Auth)

	// NOTE: IsAuthenticatedのカスタムミドルウェアを利用してクッキー内のJWTトークンの検証をしている
	//       検証が失敗したら、エラーを返して実行されないようにする
	e.POST("/api/logout", Logout, middleware.IsAuthenticated)
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func InitArticleRouting(e *echo.Echo, articleHandler ArticleHandler) {
	e.GET("/api/articles/:articleID", articleHandler.Show())
	e.GET("/api/articles", articleHandler.Index())

	// NOTE: IsAuthenticatedのカスタムミドルウェアを利用してクッキー内のJWTトークンの検証をしている
	//       検証が失敗したら、エラーを返して実行されないようにする
	e.POST("/api/articles", articleHandler.Create(), middleware.IsAuthenticated)
	e.PUT("/api/articles/:articleID", articleHandler.Update(), middleware.IsAuthenticated)
	e.DELETE("/api/articles/:articleID", articleHandler.Delete(), middleware.IsAuthenticated)
}

func InitAuthRouting(e *echo.Echo, authHandler AuthHandler) {
	e.POST("/api/login", authHandler.Create())
	e.POST("/api/logout", authHandler.Delete(), middleware.IsAuthenticated)
}
