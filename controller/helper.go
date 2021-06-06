package handlers

import (
	"gin_example/app"
	"gin_example/lib/auth"
	"gin_example/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserProvider(c *gin.Context, claims *jwt.StandardClaims) interface{} {
	var user model.User
	if err := app.DB.Where("jti = ?", claims.Id).First(&user); err == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"error": "User not found",
		})
		return nil
	}
	return &user
}

func CurrentUser(c *gin.Context) *model.User {
	v, ok := c.Get(auth.IdentityKey)

	if !ok || v == nil {
		return nil
	}

	return v.(*model.User)
}


func Required(c *gin.Context) {
	user := CurrentUser(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"error": "Authorized Required",
		})
		c.Abort()
		return
	}

	if !user.Activated {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"error": "Limited User",
		})
	}
}