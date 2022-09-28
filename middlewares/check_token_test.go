package middlewares

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestCheckTokenMiddleware(t *testing.T) {
	t.Parallel()

	//token作成
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(1),                       // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	claims2 := jwt.StandardClaims{
		Issuer:    strconv.Itoa(2),                       // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	//署名
	tokenString1, err := jwtToken.SignedString([]byte((os.Getenv("SECRET_KEY"))))
	if err != nil {
		t.Fatalf("TestCreateToken failed: %s", err.Error())
	}
	//改ざん
	tokenString2, err := jwtToken.SignedString([]byte("dammy"))
	if err != nil {
		t.Fatalf("TestCreateToken failed: %s", err.Error())
	}
	//ユーザー
	tokenString3, err := jwtToken2.SignedString([]byte((os.Getenv("SECRET_KEY"))))
	if err != nil {
		t.Fatalf("TestCreateToken failed: %s", err.Error())
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString1, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !parsedToken.Valid {
		t.Fatalf("TestParsedToken failed: %s", err.Error())
	}

	// 改ざん
	parsedToken2, err := jwt.ParseWithClaims(tokenString2, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err == nil || parsedToken2.Valid {
		t.Fatalf("TestParsedToken failed: %s", err.Error())
	}

	// ユーザー
	parsedToken3, err := jwt.ParseWithClaims(tokenString3, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !parsedToken3.Valid {
		t.Fatalf("TestParsedToken failed: %s", err.Error())
	}
	if parsedToken == parsedToken3 {
		t.Fatalf("TestParsedToken failed: %s", err.Error())
	}

}
