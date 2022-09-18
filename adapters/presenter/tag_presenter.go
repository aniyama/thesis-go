package presenter

import (
	"net/http"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
)

type TagPresenter struct {
	*MyErrorPresenter
	ctx *gin.Context
}

func NewTagOutputPort(ctx *gin.Context) port.TagOutputPort {
	return &TagPresenter{
		ctx: ctx,
	}
}

func (u *TagPresenter) OutputTags(tags *[]entities.Tag) {
	u.ctx.JSON(http.StatusOK, tags)
}

func (u *TagPresenter) OutputTag(Tag *entities.Tag) {
	u.ctx.JSON(http.StatusOK, Tag)
}

func (u *TagPresenter) Output200() {
	u.ctx.JSON(http.StatusOK, nil)
}

func (u *TagPresenter) RenderError(err error) {
	u.ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
}
