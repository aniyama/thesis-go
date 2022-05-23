package main

import (
	"github.com/aniyama/thesis-go/server"
	_ "github.com/go-sql-driver/mysql"
)

// go mod = thesis-go

func main() {
	// r := gin.Default()
	// r.GET("/greeting", func(c *gin.Context) {
	// 	greeting := "hello world"
	// 	c.JSON(200, gin.H{
	// 		"message": greeting,
	// 	})
	// })
	// r.Run()
	// fmt.Println("Hello world")

	server.Init()

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "クイックスタートしました",
	// 	})
	// })
	// r.Run()

}

// func sqlConnect() (database *gorm.DB) {
// 	DBMS := "mysql"
// 	USER := "aniyama"
// 	PASS := "aniyama"
// 	PROTOCOL := "tcp(db_container:3306)"
// 	DBNAME := "thesis_db"

// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"

// 	count := 0
// 	db, err := gorm.Open(DBMS, CONNECT)
// 	if err != nil {
// 	  for {
// 		if err == nil {
// 		  fmt.Println("")
// 		  break
// 		}
// 		fmt.Print(".")
// 		time.Sleep(time.Second)
// 		count++
// 		if count > 180 {
// 		  fmt.Println("")
// 		  fmt.Println("DB接続失敗")
// 		  panic(err)
// 		}
// 		db, err = gorm.Open(DBMS, CONNECT)
// 	  }
// 	}
// 	fmt.Println("DB接続成功")

// 	return db
//   }
