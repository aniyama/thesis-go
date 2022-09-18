package routes

import (
	"github.com/aniyama/thesis-go/driver"
	"github.com/aniyama/thesis-go/infrastructure/database"

	"github.com/aniyama/thesis-go/middlewares"
	"github.com/gin-gonic/gin"
)

func Init() {

	database.InitSQLHandler()

	ginEngine := gin.Default()
	// CORS設定
	middlewares.NewCorsMiddleware(ginEngine).Setup()

	dbHandler := database.GetDBHandler()

	// 認証関係
	auth := ginEngine.Group("/auth")
	login := driver.NewLoginDriver(dbHandler, ginEngine, auth)
	login.ServeLogin()

	apiV1 := ginEngine.Group("/api")
	middlewares.NewCheckTokenMiddleware(apiV1).CheckTokenMiddleware()
	thesis := driver.NewThesisDriver(dbHandler, ginEngine, apiV1)
	thesis.ServeThesis()

	module := driver.NewModuleDriver(dbHandler, ginEngine, apiV1)
	module.ServeModule()

	tag := driver.NewTagDriver(dbHandler, ginEngine, apiV1)
	tag.ServeTag()

	category := driver.NewCategoryDriver(dbHandler, ginEngine, apiV1)
	category.ServeCategory()

	ginEngine.Run()
}
