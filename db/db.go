package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// DB初期化
func Init() {
	// // 実行環境取得
	// env := os.Getenv("ENV")

	// if "production" == env {
	// 	env = "production"
	// } else {
	// 	env = "development"
	// } // 環境変数取得
	// godotenv.Load(".env." + env)
	// godotenv.Load()

	db = sqlConnect()

	// autoMigration()
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "aniyama"
	PASS := "aniyama"
	PROTOCOL := "tcp(db_container:3306)"
	DBNAME := "thesis_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("DB接続成功")

	return db
}

// DB取得
func GetDB() *gorm.DB {
	return db
}

// DB接続終了
// func Close() {
// 	if err := db.Close(); err != nil {
// 		panic(err)
// 	}
// }

// マイグレーション
// func autoMigration() {
// 	db.AutoMigrate(&entity.User{})
//   }
