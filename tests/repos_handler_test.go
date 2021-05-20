package tests

import (
	"fmt"
	"gin_example/app"
	"gin_example/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var test TestEngine

func TestMain(m *testing.M) {
	test = TestEngine{
		router: TestApplication(),
	}
	m.Run()
}

func CreateTestRepo() *models.Repo {
	repo := models.Repo{}
	if err := app.DB.Create(&repo).Error; err != nil {
		fmt.Println(err)
	}
	return &repo
}

func TestReposIndex(t *testing.T) {
	CreateTestRepo()

	res := test.Get("/api/repos")

	assert.Equal(t, 200, res.Code)
}

func TestReposCreate(t *testing.T) {

	res := test.Post("/api/repos", `{"name": "name"}`)

	assert.Equal(t, 200, res.Code)
}

func TestReposShow(t *testing.T) {
	repo := CreateTestRepo()

	res := test.Get(fmt.Sprintf("/api/repos/%d", repo.ID.Int64))

	assert.Equal(t, 200, res.Code)
}

func TestReposUpdate(t *testing.T) {
	repo := CreateTestRepo()

	res := test.Put(fmt.Sprintf("/api/repos/%d", repo.ID.Int64), `{"name": "name"}`)

	assert.Equal(t, 200, res.Code)
}

func TestReposDestroy(t *testing.T) {
	repo := CreateTestRepo()

	res := test.Delete(fmt.Sprintf("/api/repos/%d", repo.ID.Int64))

	assert.Equal(t, 200, res.Code)
}