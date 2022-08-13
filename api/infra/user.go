package infra

import (
	"database/sql"
	"time"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
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
		_ = tx.Rollback()

		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

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

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	return UserRepository{SqlHandler: sqlHandler}
}

func (userRepository UserRepository) GetByName(name string) (*model.User, error) {
	query := `SELECT * FROM users WHERE name = ?;`

	var user model.User
	if err := userRepository.SqlHandler.Conn.Get(&user, query, name); err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (name, password) VALUES (?, ?);`

	if _, err := userRepository.SqlHandler.Conn.Exec(query, user.Name, user.Password); err != nil {
		return err
	}

	return nil
}
