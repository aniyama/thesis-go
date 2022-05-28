package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aniyama/thesis-go/entity"
	"github.com/aniyama/thesis-go/service"
)

type Thesis entity.Thesis

// 検索 GET /books
func GetAllThesis(c *gin.Context) {

	user, err := verifyJWTToken(c)
	if err != nil {
		fmt.Println("err")
		c.String(http.StatusBadRequest, "Don't through validate")
		return
	}

	thesisList, err := service.GetAllThesis(user.Id)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, thesisList)
}

func CreateThesis(c *gin.Context) {
	user, errr := verifyJWTToken(c)
	if errr != nil {
		panic("valudatre")
	}

	//TODO
	thesis := service.Thesis{}
	
	thesis.UserId = int(user.Id)
	now := time.Now()
	// thesis.TagId = nil
	thesis.CreatedAt = now
	thesis.UpdatedAt = now
	err := c.BindJSON(&thesis)
	

	if err != nil {
		panic("unMarchal")
	}

	createdThesis, err := service.CreateThesis(thesis)

	fmt.Println(createdThesis)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, createdThesis)
}

func UpdateThesis(c *gin.Context) {

	user, errr := verifyJWTToken(c)
	if errr != nil {
		panic("aaaaaad")
	}

	//TODO
	thesis := service.Thesis{}
	id := c.Param("id")

	updatedThesis, err := service.UpdateThesis(id, thesis, int(user.Id), c)

	fmt.Println("updated", updatedThesis)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, updatedThesis)
}

func DeleteThesis(c *gin.Context) {
	_, errr := verifyJWTToken(c)
	if errr != nil {
		panic("aaaaaad")
	}

	id := c.Param("id")

	fmt.Println("###########delete", id)
	err := service.DeleteThesis(id)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, "")
}
