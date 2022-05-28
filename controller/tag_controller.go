package controller

import (
	"fmt"
	"log"
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
		panic("aaaaaad")
	}

	tagList, err := service.GetAllTag(user.Id)

	fmt.Println(tagList)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, tagList)
}

// func CreateTag(c *gin.Context) {
// 	user, errr := verifyJWTToken(c)
// 	if errr != nil {
// 		panic("aaaaaad")
// 	}

// 	//TODO
// 	Tag := service.Tag{}

// 	Tag.UserId = int(user.Id)
// 	err := c.BindJSON(&Tag)
// 	if err != nil {
// 		panic("unMarchal")
// 	}

// 	createdTag, err := service.CreateTag(Tag)

// 	fmt.Println(createdTag)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	c.JSON(http.StatusOK, createdTag)
// }

// func UpdateTag(c *gin.Context) {

// 	user, errr := verifyJWTToken(c)
// 	if errr != nil {
// 		panic("aaaaaad")
// 	}

// 	//TODO
// 	Tag := service.Tag{}

// 	Tag.UserId = int(user.Id)
// 	id := c.Param("id")

// 	err := c.BindJSON(&Tag)
// 	if err != nil {
// 		panic("unMarchal")
// 	}

// 	updatedTag, err := service.UpdateTag(id, Tag)

// 	fmt.Println(updatedTag)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.JSON(http.StatusOK, updatedTag)
// }

// func DeleteTag(c *gin.Context) {
// 	_, errr := verifyJWTToken(c)
// 	if errr != nil {
// 		panic("aaaaaad")
// 	}

// 	id := c.Param("id")

// 	err := service.DeleteTag(id)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.JSON(http.StatusOK, "")
// }
