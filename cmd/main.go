package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rifkan-py/portion-share/internal/server"
)

func main() {
	r := server.SetupRoutes()

	srv := &http.Server{
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		IdleTimeout:       2 * time.Hour,
		MaxHeaderBytes:    10 << 20,
		Handler:           r,
		Addr:              ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
