package presenter

import (
	"net/http"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type ModulePresenter struct {
	*MyErrorPresenter
	ctx *gin.Context
}

func NewModuleOutputPort(ctx *gin.Context) port.ModuleOutputPort {
	return &ModulePresenter{
		ctx: ctx,
	}
}

func (u *ModulePresenter) OutputModules(modules *[]entities.Module) {
	u.ctx.JSON(http.StatusOK, modules)
}

func (u *ModulePresenter) OutputModule(Module *entities.Module) {
	u.ctx.JSON(http.StatusOK, Module)
}

func (u *ModulePresenter) Output200() {
	u.ctx.JSON(http.StatusOK, nil)
}
