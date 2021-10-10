package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"yuki0920/go-notes/domain/model"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

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

func (usecase *mockArticleUsecase) Delete(id int) (err error) {
	return err
}

func TestArticleShow(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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

func TestArticleIndex(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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

func TestArticleCreate(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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

func TestArticleCreateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/articles", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestArticleCreateArticleWithUnknownType(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestArticleCreateArticleWithoutTitle(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestArticleUpdate(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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

func TestArticleUpdateArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル","body":"ボディ"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("PUT", ts.URL+"/api/articles/1", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestArticleUpdateArticleWithUnknownType(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestArticleUpdateArticleWithoutTitle(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

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
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}

func TestArticleDelete(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/api/articles/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestArticleDeleteArticleWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := ArticleHandler{articleUsecase: &mockArticleUsecase{}}
	InitArticleRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/api/articles/1", nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
