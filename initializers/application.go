package initializers

import (
	"gin_example/app"
	"gin_example/models"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"path"
	"runtime"
)

func Application() *gin.Engine {
	r := gin.Default()

	setup()

	r.Use(cors.Default())

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