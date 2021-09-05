package repository

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// repository パッケージ内でデータベースへグローバル変数dbを通して参照できるようにしている
func SetDB(d *sqlx.DB) {
	db = d
}
