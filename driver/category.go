package driver

import (
	"github.com/aniyama/thesis-go/adapters/controllers"
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/adapters/presenter"
	"github.com/aniyama/thesis-go/usecases/interactor"
	"github.com/gin-gonic/gin"
)

type Category interface {
	ServeCategory()
}

type CategoryDriver struct {
	gin        *gin.Engine
	controller *controllers.CategoryController
	apiV1      *gin.RouterGroup
}

func NewCategoryDriver(dbHandler gateways.DbHandler, ginEngine *gin.Engine, apiV1 *gin.RouterGroup) Category {
	outputFactory := NewCategoryOutputFactory()
	repositoryFactory := NewCategoryRepositoryFactory()
	inputFactory := NewCategoryInputFactory()
	return &CategoryDriver{
		gin:        ginEngine,
		controller: controllers.NewCategoryController(dbHandler, outputFactory, repositoryFactory, inputFactory),
		apiV1:      apiV1,
	}
}

func NewCategoryRepositoryFactory() controllers.CategoryRepositoryFactory {
	return gateways.NewCategoryRepository
}

func NewCategoryOutputFactory() controllers.CategoryOutputFactory {
	return presenter.NewCategoryOutputPort
}

func NewCategoryInputFactory() controllers.CategoryInputFactory {
	return interactor.NewCategoryInputPort
}

func (driver *CategoryDriver) ServeCategory() {
	driver.apiV1.GET("/category", driver.controller.CategoryGets())

}
