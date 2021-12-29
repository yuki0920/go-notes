package model

import "time"

type Category struct {
	ID      int       `db:"id" json:"id"`
	Title   string    `db:"title" json:"title"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}
