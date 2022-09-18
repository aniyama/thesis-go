package database

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var handler *SQLHandler

type SQLHandler struct {
	Db *gorm.DB
}

func (h *SQLHandler) GetDB() *gorm.DB {
	return h.Db
}

func InitSQLHandler() {

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := fmt.Sprintf("tcp(%s:%s)", os.Getenv("MYSQL_CONTAINER_NAME"), os.Getenv("MYSQL_PORT"))
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"

	fmt.Println(DBMS, DBNAME, USER, PASS, PROTOCOL)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println("DB接続失敗!!")
		panic(err)
	}

	fmt.Println("DB接続成功!!")
	sqlHandler := new(SQLHandler)
	sqlHandler.Db = db
	handler = sqlHandler

}

// DB取得
func GetDBHandler() *SQLHandler {
	return handler
}
