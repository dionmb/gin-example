package initializers

import (
	"gin_example/app"
	"gin_example/models"
	"github.com/gin-gonic/gin"
	"path"
	"runtime"
)

func Application() *gin.Engine {
	r := gin.Default()

	setup()

	Route(r)

	return r
}

func setup() {
	switch gin.Mode() {
	case gin.DebugMode:
		app.Env = "development"
	case gin.ReleaseMode:
		app.Env = "production"
	case gin.TestMode:
		app.Env = "test"
	}

	_, filename, _, _ := runtime.Caller(1)
	app.Root = path.Join(path.Dir(filename), "..")

	app.DB = Database(models.Repo{})
}