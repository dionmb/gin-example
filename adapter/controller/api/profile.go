package api

import (
	"gin_example/adapter/controller"
	"gin_example/domain/vo"
	"gin_example/lib/resp"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	user := handlers.CurrentUser(c)
	resp.JSON(c, vo.ProfileRes{
		Username: user.Username,
	})
}
