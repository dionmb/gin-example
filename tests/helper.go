package tests

import (
	"bytes"
	"gin_example/app"
	"gin_example/initializers"
	"gin_example/libs/auth"
	"gin_example/models"
	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
)

func TestApplication() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return initializers.Application()
}

type TestEngine struct {
	router *gin.Engine
	currentUser *models.User
}

func (t *TestEngine) Login(user *models.User) *TestEngine {
	t.currentUser = user
	return t
}

func (t *TestEngine) Logout() {
	t.currentUser = nil
}

func (t *TestEngine) http(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	if t.currentUser != nil && t.currentUser.Jti != "" {
		token, err := auth.GenerateToken(t.currentUser)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Authorization", "Bearer " + token)
	}
	t.router.ServeHTTP(res, req)
	return res
}

func (t *TestEngine) Get(path string) *httptest.ResponseRecorder {
	return t.http("GET", path, nil)
}

func (t *TestEngine) Post(path string, json string) *httptest.ResponseRecorder {
	return t.http("POST", path, bytes.NewBufferString(json))
}

func (t *TestEngine) Put(path string, json string) *httptest.ResponseRecorder {
	return t.http("PUT", path, bytes.NewBufferString(json))
}

func (t *TestEngine) Delete(path string) *httptest.ResponseRecorder {
	return t.http("DELETE", path, nil)
}

func loadFixtures() {
	db, err := app.DB.DB()
	if err != nil {
		panic(err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(path.Join(app.Root, "tests/fixtures")),
	)

	if err != nil {
		panic(err)
	}

	fixtures.Load()
}