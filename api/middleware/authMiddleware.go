package middleware

import (
	"yuki0920/go-blog/util"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err == nil {
			c.Echo().Logger.Debug("cookie is empty")
			next(c)
		}
		// FIXME: `runtime error: invalid memory address or nil pointer dereference goroutine`
		if err := util.ParseJwt(cookie.Value); err != nil {
			c.Echo().Logger.Debug("cookie parse failed %s", err)
			next(c)
		}

		return next(c)
	}
}
