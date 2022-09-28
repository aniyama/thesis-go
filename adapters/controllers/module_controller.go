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

type ModuleController struct {
	output     ModuleOutputFactory
	input      ModuleInputFactory
	repository ModuleRepositoryFactory
	dbHandler  gateways.DbHandler
	validator  ModuleValidatorFactory
}

type ModuleRepositoryFactory func(gateways.DbHandler) port.ModuleRepository
type ModuleOutputFactory func(*gin.Context) port.ModuleOutputPort
type ModuleInputFactory func(port.ModuleOutputPort, port.ModuleRepository, gateways.DbHandler) port.ModuleInputPort
type ModuleValidatorFactory func(*gin.Context, gateways.DbHandler) port.ModuleValidator

func NewModuleController(dbHandler gateways.DbHandler, output ModuleOutputFactory, repository ModuleRepositoryFactory, input ModuleInputFactory, validator ModuleValidatorFactory) *ModuleController {
	return &ModuleController{
		output:     output,
		repository: repository,
		input:      input,
		dbHandler:  dbHandler,
		validator:  validator,
	}
}

func (u *ModuleController) ModuleGets() func(c *gin.Context) {
	return func(c *gin.Context) {

		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		var theses []entities.Module

		u.newModuleInputPort(c).ModuleGets(c.Request.Context(), &theses, userID)
	}
}

func (u *ModuleController) ModuleCreate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		Module := new(entities.Module)
		Module.UserId = userID
		jst, _ := time.LoadLocation("Asia/Tokyo")
		time.Local = jst
		now := time.Now()
		Module.CreatedAt = now
		Module.UpdatedAt = now
		err := c.BindJSON(&Module)
		if err != nil {
			u.output(c).Output400Error(err)
		}

		u.newModuleInputPort(c).ModuleCreate(c.Request.Context(), Module)
	}
}

func (u *ModuleController) ModuleUpdate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		token, _ := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		claims := token.Claims.(*Claims)
		userID, _ := strconv.Atoi(claims.Issuer)

		id := c.Param("id")
		module := new(entities.Module)
		module.UserId = userID
		module.UpdatedAt = time.Now()

		err := c.BindJSON(&module)
		if err != nil {
			u.output(c).Output400Error(err)
		}
		u.newModuleInputPort(c).ModuleUpdate(c, module, id)
	}
}

func (u *ModuleController) ModuleDelete() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		module := new(entities.Module)
		u.newModuleInputPort(c).ModuleDelete(c.Request.Context(), module, id)
	}
}

func (u *ModuleController) newModuleInputPort(c *gin.Context) port.ModuleInputPort {
	outputPort := u.output(c)
	repository := u.repository(u.dbHandler)
	return u.input(outputPort, repository, u.dbHandler)
}

func (u *ModuleController) newModuleOutputPort(c *gin.Context) port.ModuleOutputPort {
	return u.output(c)
}

func (u *ModuleController) newModuleValidator(c *gin.Context) port.ModuleValidator {
	return u.validator(c, u.dbHandler)
}
