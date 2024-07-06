package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rifkan-py/portion-share/internal/routes"
)

func main() {
	r := http.NewServeMux()

	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:              ":4200",
		Handler:           r,
		MaxHeaderBytes:    10 << 20,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
