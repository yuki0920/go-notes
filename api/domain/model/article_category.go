package model

type ArticleCategory struct {
	ID         int `db:"id" json:"id"`
	ArticleID  int `db:"article_id"`
	CategoryID int `db:"category_id"`
}
