package routes

import (
	"database/sql"
	"net/http"

	"github.com/rifkan-py/portion-share/internal/handlers"
)

type ServerConfig struct {
	DB *sql.DB
}

func SetupRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /upload", handlers.UploadHandler)
}
