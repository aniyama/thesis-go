package controllers

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	output     LoginOutputFactory
	input      LoginInputFactory
	repository LoginRepositoryFactory
	dbHandler  gateways.DbHandler
	validator  LoginValidatorFactory
	auth       AuthFactory
}

type Claims struct {
	jwt.StandardClaims
}

type LoginRepositoryFactory func(gateways.DbHandler) port.LoginRepository
type LoginOutputFactory func(*gin.Context) port.LoginOutputPort
type LoginInputFactory func(port.LoginOutputPort, port.LoginRepository, gateways.DbHandler, port.AuthService) port.LoginInputPort
type LoginValidatorFactory func(*gin.Context, gateways.DbHandler) port.LoginValidator
type AuthFactory func(*gin.Context, gateways.DbHandler) port.AuthService

func NewLoginController(dbHandler gateways.DbHandler, output LoginOutputFactory, repository LoginRepositoryFactory, input LoginInputFactory, validator LoginValidatorFactory, auth AuthFactory) *LoginController {
	return &LoginController{
		output:     output,
		repository: repository,
		input:      input,
		dbHandler:  dbHandler,
		validator:  validator,
		auth:       auth,
	}
}

func (u *LoginController) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		// パラメータ取得
		var form map[string]string
		if err := c.BindJSON(&form); err != nil {
			u.newLoginOutputPort(c).Output400Error(err)
			return
		}

		name, _ := form["name"]
		password, _ := form["password"]
		user := &entities.User{
			Name:     name,
			Password: encrypt(password),
		}
		err := u.newLoginValidator(c).LoginValidator(*user)
		if err != nil {
			u.newLoginOutputPort(c).OutputValidationError(err)
			return
		}

		u.newLoginInputPort(c).Login(c.Request.Context(), user)

	}
}

func (u *LoginController) Register() func(c *gin.Context) {
	return func(c *gin.Context) {
		// パラメータ取得
		var form map[string]string
		if err := c.BindJSON(&form); err != nil {
			u.newLoginOutputPort(c).Output400Error(err)
			return
		}

		name, _ := form["name"]
		password, _ := form["password"]
		user := &entities.User{
			Name:     name,
			Password: encrypt(password),
		}

		//同じユーザー名かどうかチェック
		err := u.newLoginValidator(c).ExistsUserByUserName(name)
		if err != nil {
			u.newLoginOutputPort(c).OutputValidationError(err)
			return
		}

		err = u.newLoginValidator(c).RegisterValidator(*user)
		if err != nil {
			u.newLoginOutputPort(c).OutputValidationError(err)
			return
		}

		u.newLoginInputPort(c).Register(c.Request.Context(), user)
	}
}

func (u *LoginController) Logout() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("jwt", "", -1, "/", "http://localhost:3000", true, true)
		u.newLoginInputPort(c).Logout(c.Request.Context())
	}
}

func (u *LoginController) VerifyJWTToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		// CookieからJWTを取得
		tokenString, _ := c.Cookie("jwt")

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			u.newLoginOutputPort(c).Output403Error(err)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (u *LoginController) newLoginInputPort(c *gin.Context) port.LoginInputPort {
	outputPort := u.output(c)
	repository := u.repository(u.dbHandler)
	auth := u.auth(c, u.dbHandler)
	return u.input(outputPort, repository, u.dbHandler, auth)
}

func (u *LoginController) newLoginOutputPort(c *gin.Context) port.LoginOutputPort {
	return u.output(c)
}

func (u *LoginController) newLoginValidator(c *gin.Context) port.LoginValidator {
	return u.validator(c, u.dbHandler)
}

func (u *LoginController) newAuth(c *gin.Context) port.AuthService {
	return u.auth(c, u.dbHandler)
}

func encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
