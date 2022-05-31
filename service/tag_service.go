package service

import (
	"fmt"
	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/entity"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tag entity.Tag

// 検索
func GetAllTag(id uint) ([]Tag, error) {
	// DB接続
	db := db.GetDB()

	var tag []Tag

	result := db.Where("user_id = ?", id).Order("id desc").Find(&tag)

	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

func CreateTag(tag Tag) (Tag, error) {
	// DB接続
	db := db.GetDB()

	result := db.Create(&tag)

	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

func UpdateTag(id string, tag Tag, userId int, c *gin.Context) (Tag, error) {
	// DB接続
	db := db.GetDB()

	result := db.First(&tag, id)

	tag.UserId = userId
	err := c.BindJSON(&tag)
	fmt.Println(tag, err)
	if err != nil {
		return tag, err
	}

	db.Save(&tag)

	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

func DeleteTag(id string) error {

	// DB接続
	db := db.GetDB()
	var Tag Tag

	result := db.Where("id = ?", id).Delete(&Tag)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
