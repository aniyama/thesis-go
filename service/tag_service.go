package service

import (
	"fmt"

	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tag entity.Tag

// 検索
func GetAllTag(id uint) ([]Tag, error) {
	// DB接続
	db := db.GetDB()

	var tag []Tag

	result := db.Where("user_id = ?", id).Find(&tag)
	fmt.Printf("##################", tag)

	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

// func CreateTag(Tag Tag) (Tag, error) {
// 	// DB接続
// 	db := db.GetDB()

// 	result := db.Create(&Tag)

// 	if result.Error != nil {
// 		fmt.Println("fail")
// 		return Tag, result.Error
// 	}

// 	return Tag, nil
// }

// func UpdateTag(id string, Tag Tag) (Tag, error) {
// 	// DB接続
// 	db := db.GetDB()

// 	result := db.Where("id = ?", id).First(&Tag).Updates(&Tag)

// 	if result.Error != nil {
// 		fmt.Println("fail")
// 		return Tag, result.Error
// 	}

// 	return Tag, nil
// }

// func DeleteTag(id string) error {

// 	// DB接続
// 	db := db.GetDB()
// 	var Tag Tag

// 	result := db.Where("id = ?", id).Delete(&Tag)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }
