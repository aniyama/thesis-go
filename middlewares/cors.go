package middlewares

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct {
	e *gin.Engine
}

func NewCorsMiddleware(e *gin.Engine) CorsMiddleware {
	return CorsMiddleware{
		e: e,
	}
}

func (m CorsMiddleware) Setup() {

	// 	env := os.Getenv("ENV")
	// 	allowOrigin := ""
	// if env == "prod" {
	// 	allowOrigins = []string{
	// 		"https://issue.com"
	// 	}
	// }else {
	// 	allowOrigins = []string{
	// 		"https://issue.com"
	// 	}
	// }

	m.e.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			os.Getenv("ALLOW_CORS_PORT"),
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

}
