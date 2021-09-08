package repository

import (
	"database/sql"
	"math"
	"time"

	"yuki0920/go-blog/model"
)

func ArticleListByCursor(cursor int) ([]*model.Article, error) {
	// 引数で渡されたカーソルの値が 0 以下の場合は、代わりに int 型の最大値で置き換える
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	query := `SELECT *
	FROM articles
	WHERE id < ?
	ORDER BY id desc
	LIMIT 10`

	// クエリ結果を格納するスライスを初期化 10 件取得のため、サイズとキャパシティを指定
	articles := make([]*model.Article, 0, 10)
	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}

	return articles, nil
}

func ArticleCreate(article *model.Article) (sql.Result, error) {
	now := time.Now()

	article.Created = now
	article.Updated = now

	query := `INSERT INTO articles (title, body, created, updated)
	VALUES (:title, :body, :created, :updated);`

	// トランザクションを開始
	tx := db.MustBegin()

	// クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換される
	// 構造体タグで指定してあるフィールドが対象となる ex)`db:"title"`
	res, err := tx.NamedExec(query, article)
	if err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return res, nil
}

func ArticleDelete(id int) error {
	query := "DELETE FROM articles WHERE id = ?"

	tx := db.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
