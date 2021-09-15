package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestSample(t *testing.T) {
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

func TestPutUnknownType(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":["ABC", "DEF"]}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestPutNonTitle(t *testing.T) {
	e := echo.New()
	// アプリケーションサーバーの設定をテスト用サーバーでも設定しないと未定義で落ちる
	e.Validator = &CustomValidator{Validator: validator.New()}
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}
