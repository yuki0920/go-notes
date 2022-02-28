package handler

import (
	"yuki0920/go-notes/middleware"

	"github.com/labstack/echo/v4"
)

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
	e.GET("/api/auth", authHandler.Get())
	e.POST("/api/login", authHandler.Create())
	e.POST("/api/logout", authHandler.Delete(), middleware.IsAuthenticated)
}

func InitCategoryRouting(e *echo.Echo, categoryHandler CategoryHandler) {
	e.GET("/api/categories", categoryHandler.List())
	e.POST("/api/categories", categoryHandler.Create(), middleware.IsAuthenticated)
}
