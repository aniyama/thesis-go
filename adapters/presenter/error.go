package presenter

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MyErrorPresenter struct {
	ctx *gin.Context
}

func (m *MyErrorPresenter) errorResponse(code int, err error, msg string) {
	log.Println(err)
	m.ctx.JSON(code, map[string]interface{}{
		"code":            code,
		"message":         err.Error(),
		"display_message": msg,
	})
}

func (m *MyErrorPresenter) Output500Error(err error) {
	m.errorResponse(http.StatusInternalServerError, err, http.StatusText(m.ctx.Writer.Status()))
}

func (m *MyErrorPresenter) Output403Error(err error) {
	m.errorResponse(http.StatusForbidden, err, http.StatusText(http.StatusForbidden))
}

func (m *MyErrorPresenter) OutputValidationError(err error) {
	m.errorResponse(http.StatusUnprocessableEntity, err, http.StatusText(http.StatusUnprocessableEntity))
}

func (m *MyErrorPresenter) Output400Error(err error) {
	m.errorResponse(http.StatusBadRequest, err, http.StatusText(http.StatusBadRequest))
}

func (m *MyErrorPresenter) Output404Error(err error) {
	m.errorResponse(http.StatusNotFound, err, http.StatusText(http.StatusNotFound))
}
