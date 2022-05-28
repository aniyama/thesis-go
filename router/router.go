package router

import (
	"time"

	"github.com/aniyama/thesis-go/controller"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Router() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
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

	apiV1 := r.Group("/api")

	apiV1.GET("/thesis", controller.GetAllThesis)
	apiV1.PUT("/thesis/create", controller.CreateThesis)
	apiV1.PUT("/thesis/update/:id", controller.UpdateThesis)
	apiV1.DELETE("/thesis/delete/:id", controller.DeleteThesis)

	apiV1.GET("/tag", controller.GetAllTag)
	// apiV1.PUT("/tag/create", controller.Createtag)
	// apiV1.PUT("/tag/update:id", controller.Updatetag)
	// apiV1.DELETE("/tag/delete/id", controller.Deletetag)

	apiV1.POST("/login", controller.Controller{}.Login)
	apiV1.POST("/register", controller.Controller{}.Register)
	apiV1.POST("/logout", controller.Logout)

	r.Run()
}
