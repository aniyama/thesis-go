package driver

import (
	"github.com/aniyama/thesis-go/adapters/controllers"
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/adapters/presenter"
	"github.com/aniyama/thesis-go/adapters/validator"
	"github.com/aniyama/thesis-go/usecases/interactor"
	"github.com/gin-gonic/gin"
)

type Tag interface {
	ServeTag()
}

type TagDriver struct {
	gin        *gin.Engine
	controller *controllers.TagController
	apiV1      *gin.RouterGroup
}

func NewTagDriver(dbHandler gateways.DbHandler, ginEngine *gin.Engine, apiV1 *gin.RouterGroup) Tag {
	outputFactory := NewTagOutputFactory()
	repositoryFactory := NewTagRepositoryFactory()
	inputFactory := NewTagInputFactory()
	validatorFactory := NewTagValidatorFactory()
	return &TagDriver{
		gin:        ginEngine,
		controller: controllers.NewTagController(dbHandler, outputFactory, repositoryFactory, inputFactory, validatorFactory),
		apiV1:      apiV1,
	}
}

func NewTagRepositoryFactory() controllers.TagRepositoryFactory {
	return gateways.NewTagRepository
}

func NewTagOutputFactory() controllers.TagOutputFactory {
	return presenter.NewTagOutputPort
}

func NewTagInputFactory() controllers.TagInputFactory {
	return interactor.NewTagInputPort
}

func NewTagValidatorFactory() controllers.TagValidatorFactory {
	return validator.NewTagValidator
}

func (driver *TagDriver) ServeTag() {
	driver.apiV1.GET("/tag", driver.controller.TagGets())
	driver.apiV1.PUT("/tag/create", driver.controller.TagCreate())
	driver.apiV1.PUT("/tag/update/:id", driver.controller.TagUpdate())
	driver.apiV1.DELETE("/tag/delete/:id", driver.controller.TagDelete())

}
