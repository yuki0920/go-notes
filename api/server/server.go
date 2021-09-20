package server

import (
	"os"
	"yuki0920/go-blog/handler"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// TODO: 署名用キーは環境変数から取得する
	// restricted
	secret := os.Getenv("JWT_SECRET_KEY")
	config := middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte(secret),
	}

	// TODO: ルーティングにPOSTも含めたい
	g := e.Group("/api/articles/:articleID")
	g.Use(middleware.JWTWithConfig(config))
	e.POST("/api/articles", handler.ArticleCreate)
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete)
	e.PUT("/api/articles/:articleID", handler.ArticleUpdate)

	return e
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
