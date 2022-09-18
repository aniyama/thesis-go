package presenter

import (
	"net/http"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type ThesisPresenter struct {
	*MyErrorPresenter
	ctx *gin.Context
}

func NewThesisOutputPort(ctx *gin.Context) port.ThesisOutputPort {
	return &ThesisPresenter{
		ctx: ctx,
	}
}

func (u *ThesisPresenter) OutputTheses(theses *[]entities.Thesis) {
	u.ctx.JSON(http.StatusOK, theses)
}

func (u *ThesisPresenter) OutputThesis(thesis *entities.Thesis) {
	u.ctx.JSON(http.StatusOK, thesis)
}

func (u *ThesisPresenter) Output200() {
	u.ctx.JSON(http.StatusOK, nil)
}
