package tests

import (
	"gin_example/app"
	"gin_example/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var test TestEngine

func TestMain(m *testing.M) {
	r := TestApplication()
	test = TestEngine{
		router: r,
	}

	loadFixtures()

	m.Run()
	test.Logout()
}

func TestReposIndex(t *testing.T) {
	res := test.Get("/api/repos")
	assert.Equal(t, 200, res.Code)
}

func TestReposCreate(t *testing.T) {
	res := test.Post("/api/repos", `{"name": "name"}`)
	assert.Equal(t, 200, res.Code)
}

func TestReposShow(t *testing.T) {
	res := test.Get("/api/repos/1")
	assert.Equal(t, 200, res.Code)
}

func TestReposUpdate(t *testing.T) {
	res := test.Put("/api/repos/1", `{"name": "name"}`)
	assert.Equal(t, 200, res.Code)
}

func TestReposDestroy(t *testing.T) {
	res := test.Delete("/api/repos/1")
	assert.Equal(t, 200, res.Code)
}

func TestUnauthorized(t *testing.T) {
	res := test.Get("/api/profile")
	assert.Equal(t, 401, res.Code)
}

func TestForbidden(t *testing.T) {
	var user models.User
	app.DB.Where("activated != true").First(&user)

	res := test.Login(&user).Get("/api/profile")
	assert.Equal(t, 403, res.Code)
}

func TestProfile(t *testing.T) {
	var user models.User
	app.DB.Where("activated = true").First(&user)

	res := test.Login(&user).Get("/api/profile")
	assert.Equal(t, 200, res.Code)
}