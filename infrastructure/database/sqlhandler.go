package database

import (
	"fmt"

	"os"

	"github.com/DATA-DOG/go-txdb"
	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		fmt.Println("DB接続失敗!!")
		panic(err)
	}

	fmt.Println("DB接続成功!!")
	sqlHandler := new(SQLHandler)
	sqlHandler.Db = db
	handler = sqlHandler

}

func InitTestSQLHandler(name string) {

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := fmt.Sprintf("tcp(%s:%s)", os.Getenv("TEST_MYSQL_CONTAINER_NAME"), os.Getenv("MYSQL_PORT"))
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"

	txdb.Register(name, DBMS, CONNECT)
	fmt.Println(DBMS, DBNAME, USER, PASS, PROTOCOL)

	dialector := mysql.New(mysql.Config{
		DriverName: name,
		DSN:        CONNECT,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Println("testDB接続失敗!!")
		panic(err)
	}

	fmt.Println("testDB接続成功!!")
	sqlHandler := new(SQLHandler)
	sqlHandler.Db = db
	handler = sqlHandler

}

// DB取得
func GetDBHandler() *SQLHandler {
	return handler
}
