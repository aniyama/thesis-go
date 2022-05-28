package service

import (
	"fmt"
	"time"

	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
)

type Thesis entity.Thesis

// 検索
func GetAllThesis(id uint) ([]Thesis, error) {
	// DB接続
	db := db.GetDB()

	var theses []Thesis

	result := db.Where("user_id = ?", id).Find(&theses)
	fmt.Printf("%T\n", theses)

	if result.Error != nil {
		return theses, result.Error
	}

	return theses, nil
}

func CreateThesis(thesis Thesis) (Thesis, error) {
	// DB接続
	db := db.GetDB()
	fmt.Println(thesis)

	result := db.Create(&thesis)

	if result.Error != nil {
		fmt.Println("fail")
		return thesis, result.Error
	}

	return thesis, nil
}

func UpdateThesis(id string, thesis Thesis, userId int, c *gin.Context) (Thesis, error) {
	// DB接続
	db := db.GetDB()

	result := db.First(&thesis, id)

	thesis.UserId = userId

	// thesis.TagId = nil
	// thesis.Id = id.(uint)
	thesis.UpdatedAt = time.Now()

	err := c.BindJSON(&thesis)
	if err != nil {
		panic("unMarchal")
	}

	db.Save(&thesis)

	if result.Error != nil {
		fmt.Println("faiddd################l")
		return thesis, result.Error
	}
	fmt.Println("failx", thesis)

	return thesis, nil
}

func DeleteThesis(id string) error {

	// DB接続
	db := db.GetDB()
	var thesis Thesis

	result := db.Where("id = ?", id).Delete(&thesis)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
