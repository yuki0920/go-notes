package middleware

import (
	"net/http"
	"yuki0920/go-blog/util"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "cookie is empty",
			})
		}

		if err := util.ParseJwt(cookie.Value); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "cookie parse failed",
			})
		}

		return next(c)
	}
}
