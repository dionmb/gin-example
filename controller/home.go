package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeIndex(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}