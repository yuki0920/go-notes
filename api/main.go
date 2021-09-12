package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"yuki0920/go-blog/repository"
	"yuki0920/go-blog/server"

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

	router := server.Router(e)

	router.Validator = &CustomValidator{validator: validator.New()}

	// Webサーバーをポート番号 8080 で起動する
	port := os.Getenv("CONTAINER_PORT")
	router.Logger.Fatal(router.Start(":" + port))
}

func connectDB() *sqlx.DB {
	// DSN(Data Source Name)は環境変数として定義している
	dsn, err := dsn()
	if err != nil {
		e.Logger.Fatal(err)
	}

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

func dsn() (string, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("$DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return "", errors.New("$DB_PASSWORD is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("$DB_PORT is not set")
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("$DB_HOST is not set")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return "", errors.New("$DB_NAME is not set")
	}

	options := "charset=utf8mb4&parseTime=True&loc=Local"

	// "user:password@host:port/dbname?options"
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
		user, password, host, port, name, options), nil
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
