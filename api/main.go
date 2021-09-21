package main

import (
	"net/http"
	"os"

	database "yuki0920/go-blog/db"
	"yuki0920/go-blog/repository"
	"yuki0920/go-blog/server"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := createMux()

	db, err := database.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Info("db connection established")
	}

	repository.SetDB(db)

	router := server.Router(e)

	// echoのインスタンスにカスタムバリデーターを登録する
	router.Validator = &server.CustomValidator{Validator: validator.New()}

	// Webサーバーをポート番号 8080 で起動する
	port := os.Getenv("PORT")
	router.Logger.Fatal(router.Start(":" + port))
}

func createMux() *echo.Echo {
	// インスタンス生成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8008"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// インスタンス返却
	return e
}
