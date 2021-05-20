package initializers

import (
	"gin_example/handlers"
	"gin_example/handlers/api"
	"github.com/gin-gonic/gin"
)

func Group(app *gin.Engine, relativePath string, handler func(group *gin.RouterGroup))  {
	handler(app.Group("/api"))
}

func Route(app *gin.Engine)  {
	app.GET("/", handlers.HomeIndex)

	Group(app, "/api", func(group *gin.RouterGroup) {
		group.GET("/repos", api.ReposIndex)
		group.POST("/repos", api.ReposCreate)
		group.GET("/repos/:id", api.ReposShow)
		group.PUT("/repos/:id", api.ReposUpdate)
		group.DELETE("/repos/:id", api.ReposDestroy)
	})

}
