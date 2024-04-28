package router

import (
	"database/sql"
	"go-todo/handler"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", handler.NewHelloWorldHandler().ServeHTTP)

	return mux
}
