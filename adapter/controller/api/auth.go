package api

import (
	"gin_example/app"
	"gin_example/domain/vo"
	"gin_example/lib/resp"
	"gin_example/model"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.ParamsInvalid(c, err)
		return
	}

	var user model.User
	err := app.DB.Where(model.User{
		Username:       req.Username,
	}).First(&user).Error

	if err != nil || !user.VerifyPassword(req.Password) {
		resp.Unauthorized(c, "Username or Password Invalid")
		return
	}

	token, err := user.GenerateToken()

	if err != nil {
		resp.Unauthorized(c, err)
		return
	}

	resp.JSON(c, vo.TokenRes{
		Token: token,
		//"expire": expire.Format(time.RFC3339),
	})
}