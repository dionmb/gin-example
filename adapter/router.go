package adapter

import (
	"gin_example/adapter/controller"
	"gin_example/adapter/controller/api"
	"gin_example/lib/auth"
	"github.com/gin-gonic/gin"
)

func Route(app *gin.Engine)  {

	app.GET("/", handlers.HomeIndex)

	apiGroup := app.Group("/api")
	{
		apiGroup.POST("/login", api.Login)

		apiGroup.POST("/delay_job", api.DelaysCreate)
		apiGroup.GET("/delay_job/:id", api.DelaysShow)

		apiGroup.GET("/repos", api.ReposIndex)
		apiGroup.POST("/repos", api.ReposCreate)
		apiGroup.GET("/repos/:id", api.ReposShow)
		apiGroup.PUT("/repos/:id", api.ReposUpdate)
		apiGroup.DELETE("/repos/:id", api.ReposDestroy)

		authGroup := apiGroup.Group("")
		authGroup.Use(auth.JwtMiddleware(handlers.UserProvider))
		{
			authGroup.GET("/profile", handlers.Required, api.Profile)
		}
	}

}
