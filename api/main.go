package main

import (
	"net/http"
	"os"

	"yuki0920/go-blog/handler"
	"yuki0920/go-blog/infra"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := createMux()

	db, err := infra.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Info("db connection established")
	}

	infra.SetDB(db)
	handler.Router(e)

	// echoのインスタンスにカスタムバリデーターを登録する
	e.Validator = &handler.CustomValidator{Validator: validator.New()}

	// Webサーバーをポート番号 8080 で起動する
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func createMux() *echo.Echo {
	// インスタンス生成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONT_URL")},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// インスタンス返却
	return e
}
