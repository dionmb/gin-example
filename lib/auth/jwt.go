package auth

import (
	"gin_example/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var IdentityKey = "currentUser"

func GenerateToken(jti string) (string, error) {
	key := []byte(app.Config.JwtSecret)
	expire := time.Now().Add(time.Hour)

	claims := &jwt.StandardClaims{
		Id: jti,
		ExpiresAt: expire.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return ss, nil
}

func JwtMiddleware(UserProvider func(c * gin.Context, claims *jwt.StandardClaims) interface{}) func(c *gin.Context) {
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

		user := UserProvider(c, claims)

		c.Set(IdentityKey, user)
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