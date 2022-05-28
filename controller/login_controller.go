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

type Claims struct {
	jwt.StandardClaims
}

// 検索 GET /books
func (pc Controller) Login(c *gin.Context) {
	// パラメータ取得
	var form map[string]string
	c.BindJSON(&form)
	name, _ := form["name"]
	password, _ := form["password"]

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
		// cookie := new(http.Cookie)
    	// cookie.Value = token //Cookieに入れる値

		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("jwt", token, 3600*24, "/", "http://localhost:3000", true, true)

		fmt.Println("userName")
		fmt.Println(user)
		c.JSON(http.StatusOK, user)
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
	cookie := new(http.Cookie)
    cookie.Value = token //Cookieに入れる値

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", token, 3600*24, "/", "http://localhost:3000", true, true)
	userApi, err := service.CheckJWTToken(int(id))
	if err != nil {
		panic(err)
	}
	fmt.Println("userApi")
	fmt.Println(userApi)

	c.JSON(http.StatusOK, userApi)
}

func Logout(c *gin.Context) {

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("jwt", "", -1, "/", "http://localhost:3000", true, true)
	c.JSON(http.StatusOK, "")
}

func verifyJWTToken(c *gin.Context) (user service.User, err error) {
	// CookieからJWTを取得
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return user, err
	}
	claims := token.Claims.(*Claims)
	// User IDを取得
	id, _ := strconv.Atoi(claims.Issuer)
	user, err = service.CheckJWTToken(id)
	fmt.Println("id user", id)

	return user, err
}
