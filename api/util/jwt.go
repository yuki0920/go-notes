package util

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(issuer string) (string, error) {
	claim := &jwt.StandardClaims{
		Issuer: issuer,
	}

	// NOTE: 秘密鍵を使ってクレームを暗号化する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := os.Getenv("JWT_SECRET_KEY")

	return token.SignedString([]byte(secret))
}

func ParseJwt(claim string) error {
	// NOTE: 秘密鍵を使ってクッキーのValueを復号化する
	_, err := jwt.ParseWithClaims(claim, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	return nil
}
