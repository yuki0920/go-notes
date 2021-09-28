package handler

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

func TestGetAuthWithoutCookie(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/auth", nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("ioutil.ReadAll body failed: %s", err)
	}

	authJSON := `{"IsAuthenticated":false}`

	assert.JSONEq(t, authJSON, string(body))
}

func TestGetAuthWithCookie(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/auth", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("ioutil.ReadAll body failed: %s", err)
	}

	authJSON := `{"IsAuthenticated":true}`

	assert.JSONEq(t, authJSON, string(body))
}

func TestCreateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Post failed: %s", err)
	}

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestCreateArticleWithUnknownType(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":["ABC", "DEF"]}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Post failed: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateArticleWithoutTitle(t *testing.T) {
	e := echo.New()
	// アプリケーションサーバーの設定をテスト用サーバーでも設定しないと未定義で落ちる
	e.Validator = &CustomValidator{Validator: validator.New()}
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Post failed: %s", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestUpdateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestUpdateArticleWithUnknownType(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":["ABC", "DEF"]}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestUpdateArticleWithoutTitle(t *testing.T) {
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
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestDeleteArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	ts := httptest.NewServer(Router(e))
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/api/articles/1", nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
