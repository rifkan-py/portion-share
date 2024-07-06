package handlers

import (
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Upload Handler ran")
}
