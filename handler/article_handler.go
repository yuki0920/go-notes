package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"yuki0920/go-blog/model"
	"yuki0920/go-blog/repository"

	"github.com/labstack/echo/v4"
)

// ハンドラ関数という MVCにおけるコントローラーのアクションの位置づけ
// HTTP リクエストの情報（リクエストの送信元や各種パラメータ等）は、 echo.Context という構造体でハンドラ関数に渡ってくる
func ArticleIndex(c echo.Context) error {
	articles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message":  "Article Index",
		"Now":      time.Now(),
		"Articles": articles,
	}
	return render(c, "article/index.html", data)
}

func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func ArticleShow(c echo.Context) error {
	// パスパラメータを抽出
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

func ArticleCreate(c echo.Context) error {
	// 送信されてくるフォームの内容を格納する構造体を宣言
	var article model.Article

	// レスポンスとして返却する構造体を宣言
	var out ArticleCreateOutput

	// フォームの内容を構造体にバインド
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, out)
	}

	if err := c.Validate(&article); err != nil {
		c.Logger().Error(err.Error())

		out.Message = err.Error()

		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// repository を呼び出して保存処理を実行
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, out)
	}

	// SQL 実行結果から作成されたレコードの ID を取得
	id, _ := res.LastInsertId()

	// 構造体に ID をセット
	article.ID = int(id)

	// レスポンスの構造体に保存した記事のデータを格納
	out.Article = &article

	// JSONにパースしてレスポンスを返却
	return c.JSON(http.StatusOK, out)
}
