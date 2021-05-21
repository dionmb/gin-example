package auth

import (
	"gin_example/app"
	"gin_example/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var identityKey = "currentUser"

func CurrentUser(c *gin.Context) *models.User {
	v, ok := c.Get(identityKey)

	if !ok || v == nil {
		return nil
	}

	return v.(*models.User)
}

func GenerateToken(user * models.User) (string, error) {
	key := []byte(app.Config.JwtSecret)
	expire := time.Now().Add(time.Hour)

	claims := &jwt.StandardClaims{
		Id: user.Jti,
		ExpiresAt: expire.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return ss, nil
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


func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := getTokenFromHeader(c)

		if token == "" {
			return
		}

		time.Sleep(time.Second)
		claims, err := ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"error": err,
			})
			return
		}

		var user models.User
		if err := app.DB.Where("jti = ?", claims.Id).First(&user); err == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"error": "User not found",
			})
			return
		}

		c.Set(identityKey, &user)
	}
}

func getTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}

	return parts[1]
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.Config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)

	if !(ok && token.Valid) {
		return nil, jwt.NewValidationError("", jwt.ValidationErrorClaimsInvalid)
	}

	return claims, nil
}