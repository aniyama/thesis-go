package driver

import (
	"github.com/aniyama/thesis-go/adapters/controllers"
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/adapters/presenter"
	"github.com/aniyama/thesis-go/adapters/validator"
	"github.com/aniyama/thesis-go/usecases/interactor"
	"github.com/gin-gonic/gin"
)

type Module interface {
	ServeModule()
}

type ModuleDriver struct {
	gin        *gin.Engine
	controller *controllers.ModuleController
	apiV1      *gin.RouterGroup
}

func NewModuleDriver(dbHandler gateways.DbHandler, ginEngine *gin.Engine, apiV1 *gin.RouterGroup) Module {
	outputFactory := NewModuleOutputFactory()
	repositoryFactory := NewModuleRepositoryFactory()
	inputFactory := NewModuleInputFactory()
	validatorFactory := NewModuleValidatorFactory()
	return &ModuleDriver{
		gin:        ginEngine,
		controller: controllers.NewModuleController(dbHandler, outputFactory, repositoryFactory, inputFactory, validatorFactory),
		apiV1:      apiV1,
	}
}

func NewModuleRepositoryFactory() controllers.ModuleRepositoryFactory {
	return gateways.NewModuleRepository
}

func NewModuleOutputFactory() controllers.ModuleOutputFactory {
	return presenter.NewModuleOutputPort
}

func NewModuleInputFactory() controllers.ModuleInputFactory {
	return interactor.NewModuleInputPort
}

func NewModuleValidatorFactory() controllers.ModuleValidatorFactory {
	return validator.NewModuleValidator
}

func (driver *ModuleDriver) ServeModule() {
	driver.apiV1.GET("/module", driver.controller.ModuleGets())
	driver.apiV1.PUT("/module/create", driver.controller.ModuleCreate())
	driver.apiV1.PUT("/module/update/:id", driver.controller.ModuleUpdate())
	driver.apiV1.DELETE("/module/delete/:id", driver.controller.ModuleDelete())

}
