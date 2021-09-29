package repository

import (
	"database/sql"
	"time"
	"yuki0920/go-blog/domain/model"
)

func UserCreate(user *model.User) (sql.Result, error) {
	now := time.Now()

	user.Created = now
	user.Updated = now

	query := `INSERT INTO users (name, password, created, updated)
    VALUES (:name, :password, :created, :updated)`

	tx := db.MustBegin()
	res, err := tx.NamedExec(query, user)
	if err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return res, nil
}

func UserGetByName(name string) (*model.User, error) {
	query := `SELECT * FROM users WHERE name = ?;`

	var user model.User
	if err := db.Get(&user, query, name); err != nil {
		return nil, err
	}

	return &user, nil
}
