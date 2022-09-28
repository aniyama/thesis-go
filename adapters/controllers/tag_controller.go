package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TagController struct {
	output     TagOutputFactory
	input      TagInputFactory
	repository TagRepositoryFactory
	dbHandler  gateways.DbHandler
	validator  TagValidatorFactory
}

type TagRepositoryFactory func(gateways.DbHandler) port.TagRepository
type TagOutputFactory func(*gin.Context) port.TagOutputPort
type TagInputFactory func(port.TagOutputPort, port.TagRepository, gateways.DbHandler) port.TagInputPort
type TagValidatorFactory func(*gin.Context, gateways.DbHandler) port.TagValidator

func NewTagController(dbHandler gateways.DbHandler, output TagOutputFactory, repository TagRepositoryFactory, input TagInputFactory, validator TagValidatorFactory) *TagController {
	return &TagController{
		output:     output,
		repository: repository,
		input:      input,
		dbHandler:  dbHandler,
		validator:  validator,
	}
}

func (u *TagController) TagGets() func(c *gin.Context) {
	return func(c *gin.Context) {

		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		var tags []entities.Tag

		u.newTagInputPort(c).TagGets(c.Request.Context(), &tags, userID)
	}
}

func (u *TagController) TagCreate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		Tag := new(entities.Tag)
		Tag.UserId = userID
		err := c.BindJSON(&Tag)
		if err != nil {
			u.output(c).Output400Error(err)
		}

		u.newTagInputPort(c).TagCreate(c.Request.Context(), Tag)
	}
}

func (u *TagController) TagUpdate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		id := c.Param("id")
		Tag := new(entities.Tag)
		Tag.UserId = userID
		err := c.BindJSON(&Tag)
		if err != nil {
			u.output(c).Output400Error(err)
		}
		u.newTagInputPort(c).TagUpdate(c, Tag, id)
	}
}

func (u *TagController) TagDelete() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		Tag := new(entities.Tag)
		fmt.Println("tagtagtagtagtag", Tag, id)
		u.newTagInputPort(c).TagDelete(c.Request.Context(), Tag, id)
	}
}

func (u *TagController) newTagInputPort(c *gin.Context) port.TagInputPort {
	outputPort := u.output(c)
	repository := u.repository(u.dbHandler)
	return u.input(outputPort, repository, u.dbHandler)
}

func (u *TagController) newTagOutputPort(c *gin.Context) port.TagOutputPort {
	return u.output(c)
}

func (u *TagController) newTagValidator(c *gin.Context) port.TagValidator {
	return u.validator(c, u.dbHandler)
}
