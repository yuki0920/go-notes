package model

// タグによってメタ情報を付与し、sqlsがsqlxがSQL実行結果と構造体を紐付けてくれる
type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
