package model

import "time"

// タグによってメタ情報を付与することで、 sqlxがsqlxがSQL実行結果やフォームのname属性と紐付ける
type Article struct {
	ID      int       `db:"id" form:"id"`
	Title   string    `db:"title" form:"title"`
	Body    string    `db:"body" form:"body"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
