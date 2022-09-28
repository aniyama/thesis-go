package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CheckToken interface {
	CheckTokenMiddleware()
}

type Claims struct {
	jwt.StandardClaims
}

type CheckTokenMiddleware struct {
	apiV1 *gin.RouterGroup
}

func NewCheckTokenMiddleware(apiV1 *gin.RouterGroup) *CheckTokenMiddleware {
	return &CheckTokenMiddleware{
		apiV1: apiV1,
	}
}

func (u *CheckTokenMiddleware) CheckTokenMiddleware() {
	u.apiV1.Use(func(c *gin.Context) {
		// CookieからJWTを取得
		tokenString, _ := c.Cookie("jwt")

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		fmt.Print("86868686868")

		// 有効期限
		// claims := token.Claims.(*Claims)
		// expiresAt := claims.ExpiresAt
		// now := time.Now().Unix()
		// fmt.Print("#####", expiresAt, now)
		// if now > expiresAt {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	})
}
