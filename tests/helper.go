package tests

import (
	"bytes"
	"encoding/json"
	"gin_example/app"
	"gin_example/initializer"
	"gin_example/lib/auth"
	"gin_example/model"
	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
	"net/http"
	"net/http/httptest"
	"path"
	"text/template"
)

func TestApplication() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return initializer.Application()
}

type TestEngine struct {
	router *gin.Engine
	currentUser *model.User
}

func (t *TestEngine) Login(user *model.User) *TestEngine {
	t.currentUser = user
	return t
}

func (t *TestEngine) Logout() {
	t.currentUser = nil
}

func (t *TestEngine) http(method string, path string, data gin.H) *httptest.ResponseRecorder {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	if t.currentUser != nil && t.currentUser.Jti != "" {
		token, err := t.currentUser.GenerateToken()
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

func (t *TestEngine) Post(path string, json gin.H) *httptest.ResponseRecorder {
	return t.http("POST", path, json)
}

func (t *TestEngine) Put(path string, json gin.H) *httptest.ResponseRecorder {
	return t.http("PUT", path, json)
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
		testfixtures.TemplateFuncs(template.FuncMap{
			"EncryptPassword": auth.EncryptPassword,
		}),
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(path.Join(app.Root, "tests/fixtures")),
	)

	if err != nil {
		panic(err)
	}

	err = fixtures.Load()
	if err != nil {
		panic(err)
	}
}