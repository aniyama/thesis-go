package auth

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	ctx       *gin.Context
	dbHandler gateways.DbHandler
}

type Claims struct {
	jwt.StandardClaims
}

func NewAuth(ctx *gin.Context, dbHandler gateways.DbHandler) port.AuthService {
	return &Auth{
		ctx:       ctx,
		dbHandler: dbHandler,
	}
}

func (u *Auth) CreateToken(userId int) (tokenString string, err error) {
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(userId),                  // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//署名
	tokenString, err = jwtToken.SignedString([]byte((os.Getenv("SECRET_KEY"))))
	if err != nil {
		return tokenString, err
	}
	return tokenString, err
}

func (u *Auth) SetCookie(tokenString string) {
	u.ctx.SetSameSite(http.SameSiteNoneMode)
	u.ctx.SetCookie("jwt", tokenString, 3600*12, "/", "http://localhost:3000", true, true)
}
