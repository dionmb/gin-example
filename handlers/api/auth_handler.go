package api

import (
	"gin_example/app"
	"gin_example/libs/auth"
	"gin_example/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Params Invalid"})
		return
	}

	var user models.User
	err := app.DB.Where(models.User{
		Username:       req.Username,
		PasswordDigest: req.Password,
	}).First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or Password Invalid"})
		return
	}

	token, err := auth.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err })
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"token":  token,
		//"expire": expire.Format(time.RFC3339),
	})
}