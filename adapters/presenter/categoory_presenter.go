package presenter

import (
	"net/http"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type CategoryPresenter struct {
	*MyErrorPresenter
	ctx *gin.Context
}

func NewCategoryOutputPort(ctx *gin.Context) port.CategoryOutputPort {
	return &CategoryPresenter{
		ctx: ctx,
	}
}

func (u *CategoryPresenter) OutputCategories(categories *[]entities.Category) {
	u.ctx.JSON(http.StatusOK, categories)
}

func (u *CategoryPresenter) RenderError(err error) {
	u.ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
}
