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
    context := e.NewContext(req, res)
    hundleFunc := middleware.IsAuthenticated(func(c echo.Context) error {
        return c.String(http.StatusOK, "OK")
    })

    err := hundleFunc(context)
    require.NoError(t, err)
    assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestIsAuthenticatedWithCookie(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
    res := httptest.NewRecorder()
    context := e.NewContext(req, res)
    hundleFunc := middleware.IsAuthenticated(func(c echo.Context) error {
        return c.String(http.StatusOK, "OK")
    })

    err := hundleFunc(context)
    require.NoError(t, err)
    assert.Equal(t, http.StatusOK, res.Code)
}
