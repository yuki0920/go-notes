package middleware

import (
	"fmt"

	"yuki0920/go-blog/util"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err == nil {
			c.Echo().Logger.Debug("cookie is empty")
			err := fmt.Errorf("cookie is empty")
			c.Error(err)
		}
		// FIXME: `runtime error: invalid memory address or nil pointer dereference goroutine`
		if err := util.ParseJwt(cookie.Value); err != nil {
			c.Echo().Logger.Debug("cookie parse failed %s", err)
			c.Error(err)
		}

		return next(c)
	}
}
