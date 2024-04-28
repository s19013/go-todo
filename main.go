package main

import (
	"go-todo/db"
	"go-todo/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// envファイル読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("環境変数を読み込めませんでした: %v", err)
	}

	// サーバー用意
	todoDB := db.NewDB()
	defer todoDB.Close()

	// ルーター用意
	mux := router.NewRouter(todoDB)

	//サーバーを立てる
	err = http.ListenAndServe(os.Getenv("SERVER_PORT"), mux)
	if err != nil {
		log.Fatalf("サーバーが立ち上がりませんでした: %v", err)
	}
	log.Println("Server starting on port", os.Getenv("SERVER_PORT"))
}
