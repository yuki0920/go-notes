package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main 実行前にグローバル定数を宣言する
const tmplPath = "src/template/"

// main 実行前にグローバル変数を宣言する
var e = createMux()

func main() {
	// ルーティングの設定
	e.GET("/", articleIndex)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	// インスタンス生成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// インスタンス返却
	return e
}

// ハンドラ関数という MVCにおけるコントローラーのアクションの位置づけ
// HTTP リクエストの情報（リクエストの送信元や各種パラメータ等）は、 echo.Context という構造体でハンドラ関数に渡ってくる
func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

// pongo2を利用して、テンプレートファイルとデータからHTMLを生成しバイトデータを返却する
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// バイトデータをHTMLに変換しレスポンスを返す
	return c.HTMLBlob(http.StatusOK, b)
}
