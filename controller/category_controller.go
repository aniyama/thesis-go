package controller

import (
	"fmt"
	"net/http"

	"github.com/aniyama/thesis-go/entity"
	"github.com/aniyama/thesis-go/service"
	"github.com/gin-gonic/gin"
)

type Category entity.Tag

// 検索 GET /books
func GetAllCategory(c *gin.Context) {
	_, err := verifyJWTToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	categoryList, err := service.GetAllCategory()

	fmt.Println(categoryList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, categoryList)
}
