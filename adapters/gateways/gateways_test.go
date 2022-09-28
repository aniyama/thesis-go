package gateways_test

import (
	"net/http/httptest"

	"github.com/aniyama/thesis-go/infrastructure/database"
	"github.com/gin-gonic/gin"
)

type TestContext struct {
	Ctx *gin.Context
	Db  *database.SQLHandler
}

func NewTestContext(name string) TestContext {
	database.InitTestSQLHandler(name)
	db := database.GetDBHandler()
	response := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(response)
	return TestContext{
		Ctx: ginContext,
		Db:  db,
	}
}
