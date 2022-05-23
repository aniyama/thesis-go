package server

import (
	"github.com/aniyama/thesis-go/db"
	"github.com/aniyama/thesis-go/router"
)

// 初期化
func Init() {
	db.Init()
	router.Router()

	// r.Run()
}
