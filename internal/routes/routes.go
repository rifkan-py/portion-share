package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/rifkan-py/portion-share/internal/handlers"
)

type ServerConfig struct {
	DB *sql.DB
}

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func SetupRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		templatePath := filepath.Join(basePath, "../templates/index.html")
		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			panic(err)
		}

		fmt.Println(templatePath)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, nil)
	})
	r.HandleFunc("GET /upload", handlers.UploadHandler)
}
