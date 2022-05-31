package service

import (
	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Category entity.Category

// 検索
func GetAllCategory() ([]Category, error) {
	// DB接続
	db := db.GetDB()

	var categoryList []Category

	result := db.Find(&categoryList)

	if result.Error != nil {
		return categoryList, result.Error
	}

	return categoryList, nil
}
