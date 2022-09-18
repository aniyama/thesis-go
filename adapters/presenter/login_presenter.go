package presenter

import (
	"net/http"

	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type LoginPresenter struct {
	*MyErrorPresenter
	ctx *gin.Context
}

func NewLoginOutputPort(ctx *gin.Context) port.LoginOutputPort {
	return &LoginPresenter{
		ctx: ctx,
	}
}

func (u *LoginPresenter) Output200() {
	u.ctx.JSON(http.StatusOK, nil)
}

func (u *LoginPresenter) RenderError(err error) {
	u.ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	u.ctx.Abort()
}
