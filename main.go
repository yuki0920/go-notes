package main

import (
	"log"
	"os"

	"yuki0920/go-blog/handler"
	"yuki0920/go-blog/repository"

	_ "github.com/go-sql-driver/mysql" // MySQLのドライバーを使う
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// main 実行前にグローバル定数を宣言する
const tmplPath = "src/template/"

// main 実行前にグローバル変数を宣言する
var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// ルーティングの設定
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)
	e.POST("/", handler.ArticleCreate)

	e.Validator = &CustomValidator{validator: validator.New()}

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() *sqlx.DB {
	// DSN(Data Source Name)は環境変数として定義している
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMux() *echo.Echo {
	// インスタンス生成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// インスタンス返却
	return e
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
