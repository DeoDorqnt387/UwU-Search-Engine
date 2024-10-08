package main

import (
	"html/template"
	"net/http"
	"os"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("search.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", searchHandler)

	println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println("Failed to start server:", err)
		os.Exit(1)
	}
}
