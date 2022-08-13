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

type mockCategoryUsecase struct{}

func (usecase *mockCategoryUsecase) Create(category *model.Category) (id int64, err error) {
	return 0, err
}

func (usecase *mockCategoryUsecase) List() (categories []*model.Category, err error) {
	var mockCategory model.Category
	_ = faker.FakeData((&mockCategory))
	categories = make([]*model.Category, 0)
	categories = append(categories, &mockCategory)

	return categories, err
}

func TestCategoryCreate(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := CategoryHandler{categoryUsecase: &mockCategoryUsecase{}}
	InitCategoryRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"title":"タイトル"}`
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest("POST", ts.URL+"/api/categories", paramsJson)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCategoryList(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := CategoryHandler{categoryUsecase: &mockCategoryUsecase{}}
	InitCategoryRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/categories", nil)
	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
