package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/aniyama/thesis-go/entity"
	"github.com/aniyama/thesis-go/service"
)

type Controller struct{}

type User entity.User

// 検索 GET /books
func (pc Controller) Login(c *gin.Context) {
	// パラメータ取得
	var form map[string]string
	c.BindJSON(&form)
	name, _ := form["name"]
	password, _ := form["password"]
	fmt.Println(name)
	fmt.Println("name")

	// 検索処理
	user, err := service.Service{}.Login(name, password)


	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		// fmt.Fatalln(err)
	} else {
		fmt.Println("name233")
		claims := jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.Id)),            // stringに型変換
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//署名
		token, err := jwtToken.SignedString([]byte("secret"))
		if err != nil {
			panic(err)
		}

		c.SetCookie("jwt", token, 3600*24, "/", "http://localhost:3000", true, true)

		// cookie := gin.Context.Cookie{
		// 	Name:     "jwt",
		// 	Value:    token,
		// 	Expires:  time.Now().Add(time.Hour * 24),
		// 	Secure:   true,
		// 	HttpOnly: true, //xss対策
		// }
		//LaxはサードパーティCookieを基本的にブロックされ、Cookieがクライアント側で許可されない
		// c.SameSite = http.SameSiteNoneMode
		// c.SetCookie(&cookie)
		fmt.Println("userName")
		fmt.Println(user)
		c.JSON(http.StatusOK, user.Name)
	}
}

func (pc Controller) Register(c *gin.Context) {
	// パラメータ取得
	var form map[string]string
	c.BindJSON(&form)
	name, _ := form["name"]
	password, _ := form["password"]
	fmt.Println(name)
	fmt.Println("name")

	//同じユーザー名が存在してるか判定
	_, err := service.GetUserByName(name)
	if err == nil {
		panic(err)
	}
	err = nil

	user := service.User{
		Name:     name,
		Password: service.Encrypt(password),
	}

	id := user.CreateUser()


	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(id)),                 // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//署名
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	c.SetCookie("jwt", token, 3600*24, "/", "http://localhost:3000", true, true)

	userApi, err := service.CheckJWTToken(int(id))
	if err != nil {
		panic(err)
	}
	fmt.Println("userApi")
	fmt.Println(userApi)

	c.JSON(http.StatusOK, userApi.Name)
}

func Logout(c *gin.Context) {

	c.SetCookie("jwt", "", 0, "/", "http://localhost:3000", true, true)
	c.JSON(http.StatusOK, "")
}
