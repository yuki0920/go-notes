package handler

import (
	"net/http"
	"time"
	"yuki0920/go-blog/repository"
	"yuki0920/go-blog/util"

	"github.com/labstack/echo/v4"
)

type UserParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// レスポンスの型は明示しなくて良い
func Login(c echo.Context) error {
	var userParam UserParam
	if err := c.Bind(&userParam); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := repository.UserGetByName(userParam.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if user.ComparePassword(userParam.Password); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO: 署名用キーは環境変数から取得する
	token, err := util.GenerateJwtToken(user.Name)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
