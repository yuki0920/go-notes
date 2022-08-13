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

type mockUserUsecase struct{}

func (usecase *mockUserUsecase) GetByName(name string) (user *model.User, err error) {
	var mockUser model.User
	_ = faker.FakeData((&mockUser))

	return &mockUser, err
}

func (usecase *mockUserUsecase) Create(user *model.User) error {
	return nil
}

func TestLogin(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := AuthHandler{userUsecase: &mockUserUsecase{}}
	InitAuthRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	jsonStr := `{"id":"test","password":"test_password"}`
	// http.NewRequestのの第3引数にはio.Readerを指定するため、バイト列を渡す
	paramsJson := bytes.NewBuffer([]byte(jsonStr))
	req, _ := http.NewRequest(echo.POST, ts.URL+"/api/login", paramsJson)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestLogout(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := AuthHandler{userUsecase: &mockUserUsecase{}}
	InitAuthRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest(echo.POST, ts.URL+"/api/logout", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")

	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestLogoutWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := AuthHandler{userUsecase: &mockUserUsecase{}}
	InitAuthRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest(echo.POST, ts.URL+"/api/logout", nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestGetAuthWithoutCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := AuthHandler{userUsecase: &mockUserUsecase{}}
	InitAuthRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/auth", nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)
	authJSON := `{"IsAuthenticated":false}`

	assert.JSONEq(t, authJSON, string(body))
}

func TestGetAuthWithCookie(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	handler := AuthHandler{userUsecase: &mockUserUsecase{}}
	InitAuthRouting(e, handler)

	ts := httptest.NewServer(e)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/api/auth", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ;")
	client := &http.Client{}

	res, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err)
	authJSON := `{"IsAuthenticated":true}`

	assert.JSONEq(t, authJSON, string(body))
}
