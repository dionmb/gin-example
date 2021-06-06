package initializer

import (
	"gin_example/app"
	"gin_example/lib/configuration"
	"gin_example/model"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"path"
	"runtime"
)

func loadApplicationConfig() app.ApplicationConfig {
	var config app.ApplicationConfig
	configuration.LoadConfig("application", &config)
	return config
}

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
	app.Config = loadApplicationConfig()

	app.DB = Database(model.Repo{}, model.User{}, model.Dashboard{})
}