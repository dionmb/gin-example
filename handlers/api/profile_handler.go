package api

import (
	"gin_example/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	user := handlers.CurrentUser(c)
	c.JSON(http.StatusOK, gin.H{
		"Username": user.Username,
	})
}
