package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/sample")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("ioutil.ReadAll body failed: %s", err)
	}

	articleJSON := `{
        "id": 1,
        "title": "Sample Article",
        "body": "Sample Article Body",
        "created":"0001-01-01T00:00:00Z",
        "updated":"0001-01-01T00:00:00Z"
    }`

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.JSONEq(t, articleJSON, string(body))
}
