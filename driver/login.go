package driver

import (
	auth_login "github.com/aniyama/thesis-go/adapters/auth"
	"github.com/aniyama/thesis-go/adapters/controllers"
	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/adapters/presenter"
	"github.com/aniyama/thesis-go/adapters/validator"
	"github.com/aniyama/thesis-go/usecases/interactor"
	"github.com/gin-gonic/gin"
)

type Login interface {
	ServeLogin()
}

type LoginDriver struct {
	gin        *gin.Engine
	controller *controllers.LoginController
	auth       *gin.RouterGroup
}

func NewLoginDriver(dbHandler gateways.DbHandler, ginEngine *gin.Engine, auth *gin.RouterGroup) Login {
	outputFactory := NewLoginOutputFactory()
	repositoryFactory := NewLoginRepositoryFactory()
	inputFactory := NewLoginInputFactory()
	validatorFactory := NewLoginValidatorFactory()
	authFactory := NewAuthFactory()
	return &LoginDriver{
		gin:        ginEngine,
		controller: controllers.NewLoginController(dbHandler, outputFactory, repositoryFactory, inputFactory, validatorFactory, authFactory),
		auth:       auth,
	}
}

func NewLoginRepositoryFactory() controllers.LoginRepositoryFactory {
	return gateways.NewLoginRepository
}

func NewLoginOutputFactory() controllers.LoginOutputFactory {
	return presenter.NewLoginOutputPort
}

func NewLoginInputFactory() controllers.LoginInputFactory {
	return interactor.NewLoginInputPort
}

func NewLoginValidatorFactory() controllers.LoginValidatorFactory {
	return validator.NewLoginValidator
}

func NewAuthFactory() controllers.AuthFactory {
	return auth_login.NewAuth
}

func (driver *LoginDriver) ServeLogin() {
	driver.auth.POST("/login", driver.controller.Login())
	driver.auth.POST("/register", driver.controller.Register())
	driver.auth.POST("/logout", driver.controller.Logout())
}
