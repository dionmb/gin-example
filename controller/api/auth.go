package api

import (
	"gin_example/app"
	"gin_example/model"
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

	var user model.User
	err := app.DB.Where(model.User{
		Username:       req.Username,
	}).First(&user).Error

	if err != nil || !user.VerifyPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username or Password Invalid"})
		return
	}



	token, err := user.GenerateToken()

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