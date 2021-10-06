package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"yuki0920/go-notes/domain/model"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestSample(t *testing.T) {
	e := echo.New()
	Router(e)
	ts := httptest.NewServer(e)
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
	Router(e)
	ts := httptest.NewServer(e)
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
	Router(e)
	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/auth", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
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

func TestDeleteArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	Router(e)
	ts := httptest.NewServer(e)
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

type mockArticleUsecase struct{}

func (usecase *mockArticleUsecase) GetById(id int) (article *model.Article, err error) {
	var mockArticle model.Article
	faker.FakeData((&mockArticle))

	return &mockArticle, err
}

func (usecase *mockArticleUsecase) ListByCursor(cursor int) (articles []*model.Article, err error) {
	var mockArticle model.Article
	faker.FakeData((&mockArticle))
	articles = make([]*model.Article, 0)
	articles = append(articles, &mockArticle)

	return articles, err
}

func (usecase *mockArticleUsecase) Create(article *model.Article) (id int64, err error) {
	return int64(article.ID), err
}

func (usecase *mockArticleUsecase) Update(article *model.Article) (err error) {
	return err
}

func TestShow(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, err := http.NewRequest(echo.GET, ts.URL+"/api/articles/1", nil)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestIndex(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, err := http.NewRequest(echo.GET, ts.URL+"/api/articles", nil)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreate(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	InitRouting(e, handler)
	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
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
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":["ABC", "DEF"]}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Post failed: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateArticleWithoutTitle(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Post failed: %s", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestUpdate(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUpdateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
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
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":["ABC", "DEF"]}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestUpdateArticleWithoutTitle(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("http.Put failed: %s", err)
	}

	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}
