package handlers

import (
	"gin_example/app"
	"gin_example/lib/auth"
	"gin_example/lib/resp"
	"gin_example/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserProvider(c *gin.Context, claims *jwt.StandardClaims) interface{} {
	var user model.User
	if err := app.DB.Where("jti = ?", claims.Id).First(&user); err == nil {
		resp.JwtUnauthorized(c, "Jti not found")
		c.Abort()
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
		resp.AuthorizedRequired(c, "Limited User")
 		c.Abort()
		return
	}

	if !user.Activated {
		resp.LimitedLogin(c, "Limited User")
	}
}