package handler

import (
	"encoding/json"
	"go-todo/model"
	"log"
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{}
}

func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// レスポンスを生成する
	response := model.HelloWorldResponse{Message: "HelloWorld"}

	// データ構造をJSONにエンコードしてHTTPレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
}
