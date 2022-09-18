package driver

import (
	"github.com/aniyama/thesis-go/adapters/controllers"
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/adapters/presenter"
	"github.com/aniyama/thesis-go/adapters/validator"
	"github.com/aniyama/thesis-go/usecases/interactor"
	"github.com/gin-gonic/gin"
)

type Thesis interface {
	ServeThesis()
}

type ThesisDriver struct {
	gin        *gin.Engine
	controller *controllers.ThesisController
	apiV1      *gin.RouterGroup
}

func NewThesisDriver(dbHandler gateways.DbHandler, ginEngine *gin.Engine, apiV1 *gin.RouterGroup) Thesis {
	outputFactory := NewThesisOutputFactory()
	repositoryFactory := NewThesisRepositoryFactory()
	inputFactory := NewThesisInputFactory()
	validatorFactory := NewThesisValidatorFactory()
	return &ThesisDriver{
		gin:        ginEngine,
		controller: controllers.NewThesisController(dbHandler, outputFactory, repositoryFactory, inputFactory, validatorFactory),
		apiV1:      apiV1,
	}
}

func NewThesisRepositoryFactory() controllers.ThesisRepositoryFactory {
	return gateways.NewThesisRepository
}

func NewThesisOutputFactory() controllers.ThesisOutputFactory {
	return presenter.NewThesisOutputPort
}

func NewThesisInputFactory() controllers.ThesisInputFactory {
	return interactor.NewThesisInputPort
}

func NewThesisValidatorFactory() controllers.ThesisValidatorFactory {
	return validator.NewThesisValidator
}

func (driver *ThesisDriver) ServeThesis() {
	driver.apiV1.GET("/thesis", driver.controller.ThesisGets())
	driver.apiV1.PUT("/thesis/create", driver.controller.ThesisCreate())
	driver.apiV1.PUT("/thesis/update/:id", driver.controller.ThesisUpdate())
	driver.apiV1.DELETE("/thesis/delete/:id", driver.controller.ThesisDelete())

}
