package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(issuer string) (string, error) {
	// クレームはエンティティ(通常はユーザー)に関するデータ
	claim := &jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	// 秘密鍵を使ってクレームを暗号化する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secret))
}

func ParseJwt(claim string) error {
	// 秘密鍵を使ってクッキーのValueを復号化する
	_, err := jwt.ParseWithClaims(claim, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	fmt.Println(claim)
	if err != nil {
		return err
	}

	return nil
}
