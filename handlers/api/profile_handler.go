package api

import (
	"gin_example/libs/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	user := auth.CurrentUser(c)
	c.JSON(http.StatusOK, gin.H{
		"Username": user.Username,
	})
}
