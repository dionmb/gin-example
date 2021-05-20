package tests

import (
	"bytes"
	"gin_example/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func TestApplication() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return initializers.Application()
}

type TestEngine struct {
	router *gin.Engine
}

func (t *TestEngine) Get(path string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("Content-Type", "application/json")
	t.router.ServeHTTP(res, req)
	return res
}

func (t *TestEngine) Post(path string, json string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", path, bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")
	t.router.ServeHTTP(res, req)
	return res
}

func (t *TestEngine) Put(path string, json string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", path, bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")
	t.router.ServeHTTP(res, req)
	return res
}

func (t *TestEngine) Delete(path string) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", path, nil)
	req.Header.Set("Content-Type", "application/json")
	t.router.ServeHTTP(res, req)
	return res
}