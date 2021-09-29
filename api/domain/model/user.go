package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint      `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Password []byte    `json:"-"`
	Created  time.Time `json:"created" db:"created"`
	Updated  time.Time `json:"updated" db:"updated"`
}

func (user *User) SetPassword(password string) error {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePassword(password string) error {
	// ハッシュ化されたuser.passwordと入力の文字列のパスワードを比較している
	return bcrypt.CompareHashAndPassword((user.Password), []byte(password))
}
