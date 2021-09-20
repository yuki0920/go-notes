package handler

import (
	"log"
	"net/http"
	"yuki0920/go-blog/repository"

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

	log.Println("password is correct")
	return c.JSON(http.StatusOK, user)
}
