package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(issuer string) (string, error) {
	claim := &jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	// TODO: 署名用キーは環境変数から取得する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secret))
}

func ParseJwt(cookie string) error {
	// 秘密鍵を使ってクッキーのValueを復号化する
	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}

	return nil
}
