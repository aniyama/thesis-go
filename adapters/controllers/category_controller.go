package controllers

import (
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	output     CategoryOutputFactory
	input      CategoryInputFactory
	repository CategoryRepositoryFactory
	dbHandler  gateways.DbHandler
}

type CategoryRepositoryFactory func(gateways.DbHandler) port.CategoryRepository
type CategoryOutputFactory func(*gin.Context) port.CategoryOutputPort
type CategoryInputFactory func(port.CategoryOutputPort, port.CategoryRepository, gateways.DbHandler) port.CategoryInputPort

func NewCategoryController(dbHandler gateways.DbHandler, output CategoryOutputFactory, repository CategoryRepositoryFactory, input CategoryInputFactory) *CategoryController {
	return &CategoryController{
		output:     output,
		repository: repository,
		input:      input,
		dbHandler:  dbHandler,
	}
}

func (u *CategoryController) CategoryGets() func(c *gin.Context) {
	return func(c *gin.Context) {
		var categories []entities.Category
		u.newCategoryInputPort(c).CategoryGets(c.Request.Context(), &categories)
	}
}

func (u *CategoryController) newCategoryInputPort(c *gin.Context) port.CategoryInputPort {
	outputPort := u.output(c)
	repository := u.repository(u.dbHandler)
	return u.input(outputPort, repository, u.dbHandler)
}
