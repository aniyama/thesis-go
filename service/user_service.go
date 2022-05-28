package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Service struct{}

type User entity.User

type AUser struct {
	Name string
	Password string
}

// 検索
func (s Service) Login(name string, password string) (User, error) {

	// DB接続
	db := db.GetDB()
	// DBMS := "mysql"
	// USER := "aniyama"
	// PASS := "aniyama"
	// PROTOCOL := "tcp(db_container:3306)"
	// DBNAME := "thesis_db"

	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"


	// db, _ := gorm.Open(DBMS, CONNECT)

	// 本モデルから作成
	var user User

	result := db.Where("name = ?", name).Where("password = ?", Encrypt(password)).First(&user)
	fmt.Println(Encrypt(password))

	if result.Error != nil {
		fmt.Println("fail")
		return user, result.Error
	}

	return user, nil
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func GetUserByName(name string) (user User, err error) {
	// DB接続
	db := db.GetDB()

	result := db.Where(&User{Name: name}).First(&user)
	
	if result.Error != nil {
		return user, result.Error
	}
	return user, err
}

func (u *User) CreateUser() (id int64) {
	// DB接続
	db := db.GetDB()

	res := db.Create(&u)
	if res.Error != nil {
		panic(res.Error)
	}
	id = int64(u.Id)

	return id
}

func CheckJWTToken(userId int) (user User, err error) {
	// DB接続
	db := db.GetDB()

	result := db.Where(&User{Id: uint(userId)}).Find(&user)
	fmt.Println("check", user)

	if result.Error != nil {
		return user, err
	}

	return user, err
}
