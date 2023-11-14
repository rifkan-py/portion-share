package server

import (
	"fmt"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API is running ...")
	}))

	return r
}
