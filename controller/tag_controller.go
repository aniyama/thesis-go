package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aniyama/thesis-go/entity"
	"github.com/aniyama/thesis-go/service"
)

type Tag entity.Tag

// 検索 GET /books
func GetAllTag(c *gin.Context) {
	user, err := verifyJWTToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	tag := service.Tag{}
	tag.UserId = int(user.Id)

	tagList, err := service.GetAllTag(user.Id)

	fmt.Println(tagList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, tagList)
}

func CreateTag(c *gin.Context) {
	user, errr := verifyJWTToken(c)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errr})
		return
	}
	//TODO
	tag := service.Tag{}
	
	tag.UserId = int(user.Id)
	err := c.BindJSON(&tag)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	createdTag, err := service.CreateTag(tag)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, createdTag)
}

func UpdateTag(c *gin.Context) {

	user, errr := verifyJWTToken(c)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errr})
		return
	}

	//TODO
	tag := service.Tag{}
	id := c.Param("id")

	updatedTag, err := service.UpdateTag(id, tag, int(user.Id), c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, updatedTag)
}

func DeleteTag(c *gin.Context) {
	_, errr := verifyJWTToken(c)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errr})
		return
	}

	id := c.Param("id")

	err := service.DeleteTag(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, "")
}
