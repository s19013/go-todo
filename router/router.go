package router

import (
	"database/sql"
	"go-todo/handler"
	"go-todo/service"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/todo", func(writer http.ResponseWriter, request *http.Request) {
		todoService := service.NewTODOService(todoDB)
		todoHandler := handler.NewTODOHandler(todoService)

		if request.Method == http.MethodPost {
			todoHandler.Create(writer, request)
		}
	})

	return mux
}
