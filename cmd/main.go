package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		html := `
            <div class="bg-gray-300 rounded-full w-full h-2.5 relative max-w-4xl mx-auto mt-8">
				<div class="w-1/5 h-full rounded-full bg-green-500"></div>
				<p class="text-sm text-green-500 font-bold absolute right-0 -top-6">20%</p>
			</div>
        `
		w.Write([]byte(html))
	})

	if err := http.ListenAndServe(":4040", nil); err != nil {
		log.Fatal("failed to run the server")
	}
}
