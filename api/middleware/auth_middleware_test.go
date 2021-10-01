package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"yuki0920/go-blog/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsAuthenticatedWithoutCookie(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(echo.GET, "/", nil)
    res := httptest.NewRecorder()
    c := e.NewContext(req, res)
    hundleFunc := middleware.IsAuthenticated(func(c echo.Context) error {
        return c.String(http.StatusOK, "OK")
    })

    err := hundleFunc(c)
    require.NoError(t, err)
    assert.Equal(t, http.StatusUnauthorized, res.Code)
}
