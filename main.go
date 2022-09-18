package main

import (
	"fmt"
	"log"

	"github.com/aniyama/thesis-go/infrastructure/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("dev.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした...: %v", err)
		log.Fatalln(err)
	}

	//デプロイ時変更するため
	// err := godotenv.Load(fmt.Sprintf("envfiles/%s.env", os.Getenv("GO_ENV")))
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	fmt.Println("プロジェクト開始!!")
	routes.Init()

}
