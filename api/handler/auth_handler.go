package handler

import (
	"net/http"
	"os"
	"time"
	"yuki0920/go-blog/repository"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// レスポンスの型は明示しなくて良い
func Login(c echo.Context) error {
	name := c.FormValue("name")

	user, err := repository.UserGetByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password := c.FormValue("password")
	if user.ComparePassword(password); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: 署名用キーは環境変数から取得する
	token, err := generateJwtToken(name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func generateJwtToken(issuer string) (string, error) {
	claim := &jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	// TODO: 署名用キーは環境変数から取得する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secret := os.Getenv("JWT_SECRET_KEY")
	return token.SignedString([]byte(secret))
}
