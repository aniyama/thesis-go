package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type ThesisController struct {
	output     ThesisOutputFactory
	input      ThesisInputFactory
	repository ThesisRepositoryFactory
	dbHandler  gateways.DbHandler
	validator  ThesisValidatorFactory
}

type ThesisRepositoryFactory func(gateways.DbHandler) port.ThesisRepository
type ThesisOutputFactory func(*gin.Context) port.ThesisOutputPort
type ThesisInputFactory func(port.ThesisOutputPort, port.ThesisRepository, gateways.DbHandler) port.ThesisInputPort
type ThesisValidatorFactory func(*gin.Context, gateways.DbHandler) port.ThesisValidator

func NewThesisController(dbHandler gateways.DbHandler, output ThesisOutputFactory, repository ThesisRepositoryFactory, input ThesisInputFactory, validator ThesisValidatorFactory) *ThesisController {
	return &ThesisController{
		output:     output,
		repository: repository,
		input:      input,
		dbHandler:  dbHandler,
		validator:  validator,
	}
}

func (u *ThesisController) ThesisGets() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		var theses []entities.Thesis
		u.newThesisInputPort(c).ThesisGets(c.Request.Context(), &theses, userID)
	}
}

func (u *ThesisController) ThesisCreate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		thesis := new(entities.Thesis)
		thesis.UserId = userID
		jst, _ := time.LoadLocation("Asia/Tokyo")
		time.Local = jst
		now := time.Now()
		thesis.CreatedAt = now
		thesis.UpdatedAt = now
		err := c.BindJSON(&thesis)
		if err != nil {
			u.output(c).Output400Error(err)
		}

		u.newThesisInputPort(c).ThesisCreate(c.Request.Context(), thesis)
	}
}

func (u *ThesisController) ThesisUpdate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		id := c.Param("id")
		thesis := new(entities.Thesis)
		thesis.UserId = userID
		thesis.UpdatedAt = time.Now()

		err := c.BindJSON(&thesis)
		if err != nil {
			u.output(c).Output400Error(err)
		}
		u.newThesisInputPort(c).ThesisUpdate(c, thesis, id)
	}
}

func (u *ThesisController) ThesisDelete() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		thesis := new(entities.Thesis)
		u.newThesisInputPort(c).ThesisDelete(c.Request.Context(), thesis, id)
	}
}

func (u *ThesisController) newThesisInputPort(c *gin.Context) port.ThesisInputPort {
	outputPort := u.output(c)
	repository := u.repository(u.dbHandler)
	return u.input(outputPort, repository, u.dbHandler)
}

func (u *ThesisController) newThesisOutputPort(c *gin.Context) port.ThesisOutputPort {
	return u.output(c)
}

func (u *ThesisController) newThesisValidator(c *gin.Context) port.ThesisValidator {
	return u.validator(c, u.dbHandler)
}
