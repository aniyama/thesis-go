package auth

import (
	"os"
	"strconv"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthCheckUser struct {
	ctx *gin.Context
}

func NewAuthCheckUser(ctx *gin.Context, dbHandler gateways.DbHandler) *AuthCheckUser {
	return &AuthCheckUser{
		ctx: ctx,
	}
}

// ユーザ情報取得
func (u *AuthCheckUser) GetUserId() (userID int) {
	cookie, _ := u.ctx.Cookie("jwt")

	token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte((os.Getenv("SECRET_KEY"))), nil
	})

	claims := token.Claims.(*Claims)
	// User IDを取得
	userID, _ = strconv.Atoi(claims.Issuer)
	// userID := int64(claims["user_id"].(float64))
	return userID
}
