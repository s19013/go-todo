package router

import (
	"database/sql"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
