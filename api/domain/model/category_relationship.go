package model

type CategoryRelationship struct {
	ID       int      `db:"id" json:"id"`
	Article  Article  `db:"article_id" json:"article_id"`
	Category Category `db:"category_id" json:"category_id"`
}
